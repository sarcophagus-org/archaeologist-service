package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Dev43/arweave-go"
	"github.com/Dev43/arweave-go/tx"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	MB = 1 << 20
)

type FileHandler struct {
	AssetDoubleHash [32]byte
	EmbalmerAddress common.Address
	StorageFee      *big.Int
	Archaeologist   *Archaeologist
}

/*
	Note: The CreateTransaction function exists in arweave/transactor
	However, it uses tr.Client.TxAnchor to get the last Tx.
	That does not work locally, but using tr.Client.LastTransaction does, so that's why this function exists below.

	TODO: Test this works on live arweave blockchain (or testnet) when the time comes
*/

func (fileHandler *FileHandler) CreateTransaction(ctx context.Context, w arweave.WalletSigner, amount string, data []byte, target string) (*tx.Transaction, error) {
	tr := fileHandler.Archaeologist.ArweaveTransactor
	lastTx, err := tr.Client.LastTransaction(ctx, w.Address())
	if err != nil {
		return nil, err
	}

	price, err := tr.Client.GetReward(ctx, []byte(data))
	if err != nil {
		return nil, err
	}

	// Non encoded transaction fields
	tx := tx.NewTransaction(
		lastTx,
		w.PubKeyModulus(),
		amount,
		target,
		data,
		price,
	)

	return tx, nil
}

func (fileHandler *FileHandler) uploadFileToArweave(fileBytes []byte) (*tx.Transaction, error) {
	// create a transaction
	ar := fileHandler.Archaeologist.ArweaveTransactor
	w := fileHandler.Archaeologist.ArweaveWallet

	/*
		Arweave Transaction:
		Amount and Target are blank, as we aren't sending arweave tokens to anyone
	*/
	txBuilder, err := fileHandler.CreateTransaction(context.TODO(), w, "0", fileBytes, "")
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		return &tx.Transaction{}, err
	}

	// sign the transaction
	txn, err := txBuilder.Sign(w)
	if err != nil {
		log.Printf("Error signing transaction: %v", err)
		return &tx.Transaction{}, err
	}

	// send the transaction
	resp, err := ar.SendTransaction(context.TODO(), txn)

	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return &tx.Transaction{}, err
	}

	log.Printf("Arweave Transaction Sent: %v", resp)

	// wait for the transaction to get mined
	finalTx, err := ar.WaitMined(context.TODO(), txn)
	if err != nil {
		log.Printf("Error with transaction getting mined: %v", err)
		return &tx.Transaction{}, err
	}

	return finalTx, nil
}

func (fileHandler *FileHandler) doubleHashesMatch(fileBytes []byte) bool {
	assetSingleHash := crypto.Keccak256(fileBytes)
	assetDoubleHash := crypto.Keccak256(assetSingleHash)

	compareVal := bytes.Compare(assetDoubleHash, fileHandler.AssetDoubleHash[:])
	if compareVal != 0 {
		return false
	}

	return true
}

func (fileHandler *FileHandler) fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "There was an error receiving your file:"+err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()
	fileBytes, _ := ioutil.ReadAll(file)
	fileByteLen := len(fileBytes)

	/* Validate Size. */
	if fileByteLen > (3 * MB) {
		http.Error(w, "The file sent is larger than the limit of 3MB.", http.StatusBadRequest)
		return
	}

	/* Validate Storage Fee is sufficient */
	storageExpectation := new(big.Int).Mul(big.NewInt(int64(fileByteLen)), big.NewInt(fileHandler.Archaeologist.FeePerByte))
	if storageExpectation.Cmp(fileHandler.StorageFee) == 1 {
		errMsg := fmt.Sprintf("The storage fee is not enough. Expected storage fee of at least: %v, storage fee was: %v", storageExpectation, fileHandler.StorageFee)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	/* Decrypt the outer layer of file */
	currentPrivateKey := hdw.PrivateKeyFromIndex(fileHandler.Archaeologist.Wallet, fileHandler.Archaeologist.AccountIndex)
	decryptedFileBytes, err := utility.DecryptFile(fileBytes, currentPrivateKey)
	if err != nil {
		log.Printf("Error decrypting file: %v", err)
		http.Error(w, "The file cannot be decrypted by archaeologist. Confirm it was encrypted with the correct public key.", http.StatusBadRequest)
		return
	}

	/* Validate the inner layer double hash matches double hash */
	if !fileHandler.doubleHashesMatch(decryptedFileBytes) {
		http.Error(w, "The double hash of the file does not match the asset double hash provided.", http.StatusBadRequest)
		return
	}

	log.Printf("File %s was validated successfully:", header.Filename)

	arweaveTx, err := fileHandler.uploadFileToArweave(fileBytes)
	if err != nil {
		log.Printf("Arweave Transaction Failed: %v", err)
		http.Error(w, "There was an error with the file. Please try again.", http.StatusBadRequest)
		return
	}

	log.Printf("Transaction from arweave successful: %v", arweaveTx.Hash())

	/* Generate new public and private keys */

	/* Respond to embalmer with:
		New Public Key
	   	Arweave Tx Hash
		Signature of New Public Key + Tx Hash
	*/
	arweaveTxHash := arweaveTx.Hash()
	newPublicKey := hdw.PublicKeyFromIndex(fileHandler.Archaeologist.Wallet, fileHandler.Archaeologist.AccountIndex+1)
	pubKeyConcatTxHash := append(crypto.FromECDSAPub(newPublicKey)[1:], []byte(arweaveTxHash)...)
	hash := crypto.Keccak256Hash(pubKeyConcatTxHash)
	assetIdSig, err := crypto.Sign(hash.Bytes(), fileHandler.Archaeologist.PrivateKey)
	if err != nil {
		log.Printf("Couldnt sign the arweave tx: %v", err)
		http.Error(w, "There was an error with the file. Please try again.", http.StatusBadRequest)
		return
	}

	R, S, V := utility.SigRSV(assetIdSig)

	w.Header().Set("Content-Type", "application/json")
	response := ResponseToEmbalmer{
		NewPublicKey:    crypto.FromECDSAPub(newPublicKey)[1:],
		AssetId:         arweaveTxHash,
		AssetDoubleHash: fileHandler.AssetDoubleHash,
		V:               V,
		R:               R,
		S:               S,
	}

	json.NewEncoder(w).Encode(response)

	/* TODO: Potentially Close Connection if there is an error? */
	// We may not want to close it b/c user may re-try file upload
	// Perhaps limit # of attempts before closing?
	// defer bufrw.Flush()
	// defer conn.Close()
	// Could potentially set KeepAlivesEnabled to false: https://golang.org/src/net/http/server.go?s=96272:96319#L3068
}

func (fileHandler *FileHandler) HandleFileUpload() {
	sm := http.NewServeMux()
	sm.Handle("/file", http.HandlerFunc(fileHandler.fileUploadHandler))

	server := &http.Server{Addr: ":" + fileHandler.Archaeologist.FilePort, Handler: sm}

	log.Printf("Listening for file on port %s:", fileHandler.Archaeologist.FilePort)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("Error starting http server:", err)
		}
	}()

	// Stop the server if the program exits
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error shutting down http server:", err)
	}
}
