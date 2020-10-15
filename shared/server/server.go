package server

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	MB = 1 << 20
)

var assetDoubleHash [32]byte
var embalmerAddress common.Address
var archPrivateKey *ecdsa.PrivateKey

func embalmerSignatureValid(signedAssetDoubleHash string) bool {
	signedAssetDoubleHashBytes, err := hex.DecodeString(signedAssetDoubleHash)
	if err != nil {
		log.Printf("Could not decode signature: %v", err)
		return false
	}

	hopefullyEmbalmerPubKey, err := crypto.SigToPub(assetDoubleHash[:], signedAssetDoubleHashBytes)
	if err != nil {
		log.Printf("Could not derive embalmers public key from hash and signature: %v", err)
		return false
	}

	hopefullyEmbalmerAddress := crypto.PubkeyToAddress(*hopefullyEmbalmerPubKey)
	if hopefullyEmbalmerAddress != embalmerAddress {
		log.Printf("Address derived from embalmers signature does not match")
		return false
	}

	log.Printf("embalmers signature is valid!")

	return true
}

func canDecryptFile(file *os.File) bool {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Printf("Error copying file to buffer: %v", err)
		return false
	}

	fileBytes := buf.Bytes()

	privateKeyBytes := crypto.FromECDSA(archPrivateKey)
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)

	_, decryptError := btcec.Decrypt(privKey, fileBytes)
	if decryptError != nil {
		log.Printf("Error decrypting file: %v", decryptError)
		return false
	}

	log.Printf("file decrypted successfully!")

	return true
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
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

	if header.Size > (3 * MB) {
		http.Error(w, "The file sent is larger than the limit of 3MB.", http.StatusBadRequest)
		return
	}

	/* Validate embalmers signature */
	signedAssetDoubleHash := r.Form.Get("signedAssetDoubleHash")
	if !embalmerSignatureValid(signedAssetDoubleHash) {
		http.Error(w, "The signature from the embalmer could not be verified.", http.StatusBadRequest)
	}

	/* Save temp file and validate the file 2nd layer of file encryption can be decrypted */
	defer file.Close()

	tmpFile, err := os.Create("tmp/file")
	if err != nil {
		fmt.Fprintf(w, "Failed to open the file for writing")
		return
	}
	defer tmpFile.Close()
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	osFile, err := os.Open("tmp/file")

	if !canDecryptFile(osFile) {
		http.Error(w, "The file cannot be decrypted", http.StatusBadRequest)
	}

	/*
		Validations have passed.
		TODO: Handle file upload to arweave
	*/

	fmt.Fprintf(w, "File %s was uploaded and validated successfully.", header.Filename)
}

func HandleFileUpload(filePort string, doubleHash [32]byte, embalmerAddy common.Address, archPrivKey *ecdsa.PrivateKey) {
	assetDoubleHash = doubleHash
	embalmerAddress = embalmerAddy
	archPrivateKey = archPrivKey

	sm := http.NewServeMux()
	sm.Handle("/file", http.HandlerFunc(fileUploadHandler))

	server := &http.Server{Addr: ":" + filePort, Handler: sm}

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