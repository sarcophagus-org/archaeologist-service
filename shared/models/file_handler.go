package models

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/Dev43/arweave-go"
	"github.com/Dev43/arweave-go/transactor"
	"github.com/Dev43/arweave-go/tx"
	"github.com/Dev43/arweave-go/wallet"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
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
	ArchPrivateKey *ecdsa.PrivateKey
	StorageFee *big.Int
	FeePerByte *big.Int
	FilePort string
	ArweaveTransactor *transactor.Transactor
	ArweaveWallet *wallet.Wallet
}

/*
	Note: The CreateTransaction function exists in arweave/transactor
	However, it uses tr.Client.TxAnchor to get the last Tx.
	That does not work locally, but using tr.Client.LastTransaction does, so that's why this function exists below.

	TODO: Test this works on live arweave blockchain (or testnet) when the time comes
 */

func (fileHandler *FileHandler) CreateTransaction(ctx context.Context, w arweave.WalletSigner, amount string, data []byte, target string) (*tx.Transaction, error) {
	tr := fileHandler.ArweaveTransactor
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

func (fileHandler *FileHandler) uploadFileToArweave(file *os.File) (*tx.Transaction, error){
	// create a transaction

	ar := fileHandler.ArweaveTransactor
	w := fileHandler.ArweaveWallet
	fileBytes, _ := utility.FileToBytes(file)
	log.Println("FILE BYTES 2:", fileBytes)
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
		http.Error(w, "There was an error receiving your file:" + err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("File received with header:", header)

	fileBytes, _ := ioutil.ReadAll(file)
	fileByteLen := len(fileBytes)

	/* Validate Size. */
	if fileByteLen > (3 * MB) {
		http.Error(w, "The file sent is larger than the limit of 3MB.", http.StatusBadRequest)
		return
	}

	/* Validate Storage Fee is sufficient */
	storageExpectation := new(big.Int).Mul(big.NewInt(int64(fileByteLen)), fileHandler.FeePerByte)
	if storageExpectation.Cmp(fileHandler.StorageFee) == 1 {
		errMsg := fmt.Sprintf("The storage fee is not enough. Expected storage fee of at least: %v, storage fee was: %v", storageExpectation, fileHandler.StorageFee)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	/* Save temp file needed for decryption validation */
	defer file.Close()

	/* Decrypt the outer layer of file */
	decryptedFileBytes, err := utility.DecryptFile(fileBytes, fileHandler.ArchPrivateKey)
	if err != nil {
		http.Error(w, "The file cannot be decrypted by archaeologist. Confirm it was encrypted with the correct public key.", http.StatusBadRequest)
		return
	}

	/* Validate the inner layer double hash matches double hash */
	if !fileHandler.doubleHashesMatch(decryptedFileBytes) {
		http.Error(w, "The double hash of the file does not match the asset double hash provided.", http.StatusBadRequest)
		return
	}

	log.Printf("File %s was validated successfully:", header.Filename)

	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		if err := os.Mkdir("tmp", 0755); err != nil {
			log.Fatalf("Failed to create tmp directory for file.")
		}
	}

	filePath := fmt.Sprintf("tmp/%s", header.Filename)
	tmpFile, err := os.Create(filePath)
	if err != nil {
		log.Println("Failed to open the file for writing.")
		http.Error(w, "The archaeologist had an issue receiving the file. Please try again.", http.StatusBadRequest)
		return
	}
	defer tmpFile.Close()
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		log.Println("Failed to copy the file to disk.")
		http.Error(w, "The archaeologist had an issue receiving the file. Please try again.", http.StatusBadRequest)
		return
	}
	osFile, err := os.Open(filePath)

	log.Println("FILE BYTES 1:", fileBytes)

	arweaveTx, err := fileHandler.uploadFileToArweave(osFile)
	if err != nil {
		log.Printf("Arweave Transaction Failed: %v", err)
		http.Error(w, "There was an error with the file. Please try again.", http.StatusBadRequest)
		return
	}

	log.Printf("Transaction from arweave successful: %v", arweaveTx.Hash())

	/* Sign Arweave TX and respond to the Embalmer */
	assetIdSig, err := crypto.Sign(arweaveTx.ID(), fileHandler.ArchPrivateKey)
	if err!= nil {
		log.Printf("Couldnt sign the arweave tx: %v", err)
		http.Error(w, "There was an error with the file. Please try again.", http.StatusBadRequest)
		return
	}

	R, S, V := utility.SigRSV(assetIdSig)

	w.Header().Set("Content-Type", "application/json")
	response := ResponseToEmbalmer {
		AssetId: hexutil.Encode(assetIdSig),
		AssetDoubleHash: fileHandler.AssetDoubleHash,
		V: V,
		R: R,
		S: S,
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

	server := &http.Server{Addr: ":" + fileHandler.FilePort, Handler: sm}

	log.Printf("Listening for file on port %s:", fileHandler.FilePort)

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