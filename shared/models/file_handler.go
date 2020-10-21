package models

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/Dev43/arweave-go"
	"github.com/Dev43/arweave-go/transactor"
	"github.com/Dev43/arweave-go/tx"
	"github.com/Dev43/arweave-go/wallet"
	"github.com/btcsuite/btcd/btcec"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
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

func (fileHandler *FileHandler) embalmerSignatureValid(signedAssetDoubleHash string) bool {
	signedAssetDoubleHashBytes, err := hex.DecodeString(signedAssetDoubleHash)
	if err != nil {
		log.Printf("Could not decode signature: %v", err)
		return false
	}

	hopefullyEmbalmerPubKey, err := crypto.SigToPub(fileHandler.AssetDoubleHash[:], signedAssetDoubleHashBytes)
	if err != nil {
		log.Printf("Could not derive embalmers public key from hash and signature: %v", err)
		return false
	}

	hopefullyEmbalmerAddress := crypto.PubkeyToAddress(*hopefullyEmbalmerPubKey)
	if hopefullyEmbalmerAddress != fileHandler.EmbalmerAddress {
		log.Printf("Address derived from the provided signature does not match the address of embalmer that created the Sarcophagus")
		return false
	}

	log.Printf("embalmers signature is valid!")

	return true
}

func (fileHandler *FileHandler) canDecryptFile(file *os.File) bool {
	fileBytes, err := utility.FileToBytes(file)
	if err != nil {
		log.Printf("Error copying file to buffer: %v", err)
		return false
	}

	privateKeyBytes := crypto.FromECDSA(fileHandler.ArchPrivateKey)
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)

	_, decryptError := btcec.Decrypt(privKey, fileBytes)
	if decryptError != nil {
		log.Printf("Error decrypting file with your private key. Error: %v", decryptError)
		return false
	}

	log.Printf("file decrypted successfully!")

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

	/* Validate Size. */
	if header.Size > (3 * MB) {
		http.Error(w, "The file sent is larger than the limit of 3MB.", http.StatusBadRequest)
		return
	}

	/* Validate Storage Fee is sufficient */
	storageExpectation := new(big.Int).Mul(big.NewInt(header.Size), fileHandler.FeePerByte)
	if storageExpectation.Cmp(fileHandler.StorageFee) == 1 {
		errMsg := fmt.Sprintf("The storage fee is not enough. Expected storage fee of at least: %v, storage fee was: %v", storageExpectation, fileHandler.StorageFee)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	/* Validate embalmers signature */
	signedAssetDoubleHash := r.Form.Get("signedAssetDoubleHash")
	if !fileHandler.embalmerSignatureValid(signedAssetDoubleHash) {
		http.Error(w, "The signature from the embalmer could not be verified.", http.StatusBadRequest)
		return
	}

	/* Save temp file needed for decryption validation */
	defer file.Close()
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

	/* Validate the 2nd layer of file encryption can be decrypted */
	if !fileHandler.canDecryptFile(osFile) {
		http.Error(w, "The file cannot be decrypted by archaeologist. Confirm it was encrypted with the correct public key.", http.StatusBadRequest)
		return
	}

	log.Printf("File %s was uploaded and validated successfully:", header.Filename)

	arweaveTx, err := fileHandler.uploadFileToArweave(osFile)
	if err != nil {
		log.Println("Arweave Transaction Failed")
	}

	log.Printf("Transaction from arweave successful: %v", arweaveTx.Hash())

	/* Cleanup local file */
	os.Remove(filePath)
	os.Remove("tmp")

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