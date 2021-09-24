// Archaeologist model is responsible for:
// 1 .Registering and Updating the Archaeologist on the Sarcophagus contract
// 2. Handling Free Bond deposits and withdrawal to the contract
// 3. Handling the file sent from the embalmer, uploading it to Arweave,
//    and responding to the embalmer with the values needed to update a created Sarcophagus
// 4. Stopping and Starting the server that is used to receive the file from embalmer

package models

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Archaeologist struct {
	Client                *ethclient.Client
	ArweaveWallet         *goar.Wallet
	ArweaveClient         *goar.Client
	ArweaveMultiplier     int64
	PrivateKey            *ecdsa.PrivateKey
	CurrentPublicKeyBytes []byte
	CurrentPrivateKey     *ecdsa.PrivateKey
	ArchAddress           common.Address
	PaymentAddress        common.Address
	SarcoAddress          common.Address
	SarcoSession          contracts.SarcophagusSession
	SarcoTokenAddress     common.Address
	TokenSession          contracts.TokenSession
	FreeBond              *big.Int
	FeePerByte            *big.Int
	MinBounty             *big.Int
	MinDiggingFee         *big.Int
	MaxResurectionTime    *big.Int
	Endpoint              string
	FilePort              string
	Mnemonic              string
	Wallet                *hdwallet.Wallet
	AccountIndex          int
	Server                *http.Server
	Sarcophaguses         map[[32]byte]*Sarco
	FileHandlers          map[[32]byte]*big.Int
	RebuildChan           chan string
	ReconnectChan         chan string
}

// MB used for validating file size
const (
	MB = 1 << 20
)

// SarcoBalance returns archaeologists Sarco Balance at the address
// derived from the config value: eth_private_key
func (arch *Archaeologist) SarcoBalance() *big.Int {
	balance, err := arch.TokenSession.BalanceOf(arch.ArchAddress)

	if err != nil {
		log.Fatalf("Could not get sarcophagus balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

// SarcoBalance returns archaeologists Eth Balance at the address
// derived from the config value: eth_private_key
func (arch *Archaeologist) EthBalance() *big.Int {
	balance, err := arch.Client.BalanceAt(context.Background(), arch.ArchAddress, nil)

	if err != nil {
		log.Fatalf("Could not get eth balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

// ApproveFreeBondTransfer approves the value set in the config value: add_to_free_bond
func (arch *Archaeologist) ApproveFreeBondTransfer() {
	archSarcoBalance := arch.SarcoBalance()

	// Check if archSarcoBalance < freeBond
	// Throw error if current balance is less than what is attempting to be withdrawn
	if archSarcoBalance.Cmp(arch.FreeBond) == -1 {
		log.Fatalf("Your balance is too low to cover the free bond transfer. \n Balance Needed: %v \n Current Balance: %v", arch.FreeBond, archSarcoBalance)
	}

	txn, err := arch.TokenSession.Approve(
		arch.SarcoAddress,
		arch.FreeBond,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Approval Transaction for %v Sarco Tokens Submitted. Transaction ID: %v", utility.ToDecimal(arch.FreeBond, 18), txn.Hash().Hex())
	log.Printf("Gas Used for Approval: %v", txn.Gas())

	err = ethereum.WaitMined(arch.Client, txn.Hash(), "Approval of Free Bond")

	if err != nil {
		log.Fatalf("There was an error mining the approval of free bond transaction: %v", err)
	}

	// buffer added to avoid issue with post-approval tx failing
	time.Sleep(10 * time.Second)
}

// WithdrawBond withdraws Sarco tokens from the archaeologist's Free Bond balance
func (arch *Archaeologist) WithdrawBond(bondToWithdraw *big.Int) {
	txn, err := arch.SarcoSession.WithdrawBond(bondToWithdraw)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Withdrawing Bond: %v \n Config value REMOVE_FROM_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Withdrawal of %v Sarco Tokens transaction submitted. Transaction ID: %v", bondToWithdraw, txn.Hash().Hex())
	log.Printf("Gas Used for Withdrawal: %v", txn.Gas())

	err = ethereum.WaitMined(arch.Client, txn.Hash(), "Withdraw Bond")

	if err != nil {
		log.Fatalf("There was an error mining the withdraw bond transaction: %v", err)
	} else {
		log.Printf("Withdrawal of Sarco Tokens transaction successful")
	}
}

// RegisterArchaeologist - registers the archaeologist on the contract
// If the free bond value is > 0 then the arch can accept new jobs
func (arch *Archaeologist) RegisterArchaeologist() {
	log.Println("***REGISTERING ARCHAEOLOGIST***")
	txn, err := arch.SarcoSession.RegisterArchaeologist(
		arch.CurrentPublicKeyBytes,
		arch.Endpoint,
		arch.PaymentAddress,
		arch.FeePerByte,
		arch.MinBounty,
		arch.MinDiggingFee,
		arch.MaxResurectionTime,
		arch.FreeBond,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error registering Archaeologist: %v Config values ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND have been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Register Archaeologist Transaction Submitted. Transaction ID: %s", txn.Hash().Hex())
	log.Printf("Gas Used: %v", txn.Gas())

	err = ethereum.WaitMined(arch.Client, txn.Hash(), "Registering Archaeologist")
	arch.FreeBond = big.NewInt(0)
	if err != nil {
		log.Fatalf("There was an error mining the register archaeologist transaction: %v", err)
	} else {
		log.Printf("Register Archaeologist Transaction Successful")
	}
}

// UpdateArchaeologist .
func (arch *Archaeologist) UpdateArchaeologist() {
	log.Println("***UPDATING ARCHAEOLOGIST***")
	log.Printf("Free Bond: %v", arch.FreeBond)
	log.Printf("Endpoint: %v", arch.Endpoint)
	log.Printf("ArchAddress: %v", arch.ArchAddress)
	log.Printf("FeePerByte: %v", arch.FeePerByte)
	log.Printf("MinBounty: %v", arch.MinBounty)
	log.Printf("MinDiggingFee: %v", arch.MinDiggingFee)
	log.Printf("MaxResurectionTime: %v", arch.MaxResurectionTime)
	txn, err := arch.SarcoSession.UpdateArchaeologist(
		arch.Endpoint,
		arch.CurrentPublicKeyBytes,
		arch.ArchAddress,
		arch.FeePerByte,
		arch.MinBounty,
		arch.MinDiggingFee,
		arch.MaxResurectionTime,
		arch.FreeBond,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error updating Archaeologist: %v Config values ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND have been reset to 0. You will need to reset these.", err)
	}

	log.Printf("Update Archaeologist Transaction Submitted. Transaction ID: %s", txn.Hash().Hex())
	log.Printf("Gas Used: %v", txn.Gas())

	err = ethereum.WaitMined(arch.Client, txn.Hash(), "Update Archaeologist")
	arch.FreeBond = big.NewInt(0)
	if err != nil {
		log.Fatalf("There was an error mining the update archaeologist transaction: %v", err)
	} else {
		log.Printf("Update Archaeologist Transaction Successful")
	}
}

// UploadFileToArweave uploads the double encrypted file bytes to arweave
// Creates and returns an arweave tx
func (arch *Archaeologist) UploadFileToArweave(fileBytes []byte) (string, error) {
	w := arch.ArweaveWallet

	log.Printf("Uploading file bytes to arweave: %v", fileBytes)
	id, err := w.SendDataSpeedUp(fileBytes, []types.Tag{}, arch.ArweaveMultiplier)
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		return "", err
	}

	log.Printf("Arweave Transaction Sent: %v", id)

	return id, nil
}

// fileHandlerCheck resets the file handlers if there is only 1 file handler open
// called either when an error has occurred in the file upload process or a sarcophagus is removed from state
func (arch *Archaeologist) fileHandlerCheck() {
	fileHandlerLen := len(arch.FileHandlers)
	if fileHandlerLen <= 1 {
		arch.FileHandlers = map[[32]byte]*big.Int{}
	}
}

// fileUploadError .
func (arch *Archaeologist) fileUploadError(logMsg string, httpErrMsg string, httpErrType int, w http.ResponseWriter) {
	log.Printf("Error uploading file: %v", logMsg)
	http.Error(w, httpErrMsg, httpErrType)

	// Check if only one file handler exists, and if so, remove it
	arch.fileHandlerCheck()
}

func (arch *Archaeologist) pingHandler(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "true")
}

// fileUploadHandler validates the file sent by the embalmer for:
// 1. Size (< 3MB)
// 2. Can be Decrypted using private key a the current account index from the hd wallet
// 3. Storage Fee sent by embalmer is adequate
func (arch *Archaeologist) fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")

	log.Print("Receiving File...")

	log.Print("Starting Arch State Rebuild in fileHandler")
	arch.RebuildChan <- "start"
	log.Print("Waiting for Rebuild Response in fileHandler...")
	<-arch.RebuildChan
	log.Print("Finished Rebuild in fileHandler...")

	if len(arch.FileHandlers) < 1 {
		http.Error(w, "We are not expecting a file", 406)
		return
	}

	if r.Method != "POST" && r.Method != "OPTIONS" {
		arch.fileUploadError("File handler received non-post method, exiting.", "Method not allowed", http.StatusMethodNotAllowed, w)
		return
	}

	// Double Encrypted file bytes are sent as encoded json
	// Decode the file bytes into a struct
	var sarcoFile SarcoFile

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&sarcoFile)
	if err != nil {
		log.Printf("error decoding: %v", err)
		http.Error(w, err.Error(), 400)
		return
	}

	fileBytes, err := base64.StdEncoding.DecodeString(sarcoFile.FileBytes)
	if err != nil {
		log.Printf("error decoding file: %v", err)
	}

	fileByteLen := len(fileBytes)

	// validate Size
	if fileByteLen > (3 * MB) {
		arch.fileUploadError("File was too large to receive. Size:"+string(fileByteLen), "The file sent is larger than the limit of 3MB.", http.StatusBadRequest, w)
		return
	}

	// validate outer layer of file encryption can be decrypted
	log.Print("Decrypting file...")
	currentPrivateKey := hdw.PrivateKeyFromIndex(arch.Wallet, arch.AccountIndex)
	decryptedFileBytes, err := utility.DecryptFile(fileBytes, currentPrivateKey)

	if err != nil {
		arch.fileUploadError("Error decrypting file:"+err.Error(), "The file cannot be decrypted by archaeologist. Confirm it was encrypted with the correct public key.", http.StatusBadRequest, w)
		return
	}

	// calculate the sarcophagus identifier from the file bytes
	assetDoubleHash := utility.FileBytesToDoubleHashBytes(decryptedFileBytes)
	log.Printf("asset double hash: %v", assetDoubleHash)

	// validate the sarcophagus identifier matches a sarcophagus identifier the archaeologist is expecting to receive a file for
	// this would be an edge case where the embalmer sends a correctly encrypted file for the wrong sarcophagus
	storageFee, ok := arch.FileHandlers[assetDoubleHash]
	if !ok {
		errMsg := "The double hash of the file does not match any open double hashes."
		arch.fileUploadError(errMsg, errMsg, http.StatusBadRequest, w)
		return
	}

	// validate storage fee is sufficient
	storageExpectation := new(big.Int).Mul(big.NewInt(int64(fileByteLen)), arch.FeePerByte)
	if storageExpectation.Cmp(storageFee) == 1 {
		errMsg := fmt.Sprintf("The storage fee is not enough. Expected storage fee of at least: %v, storage fee was: %v", utility.ToDecimal(storageExpectation, 18), utility.ToDecimal(storageFee, 18))
		arch.fileUploadError(errMsg, errMsg, http.StatusBadRequest, w)
		return
	}

	// all validations have passed
	log.Printf("File was validated successfully")

	// create arweave tx
	arweaveTxId, err := arch.UploadFileToArweave(fileBytes)
	if err != nil {
		errMsg := fmt.Sprintf("There was an error with the file. Error: %v", err)
		arch.fileUploadError(errMsg, errMsg, http.StatusBadRequest, w)
		return
	}

	log.Printf("Transaction from arweave successful: %v", arweaveTxId)

	// respond to embalmer with:
	// 1. Public key from the hd wallet at the next index
	// 2. Sarcophagus identifier
	// 3. Arweave Tx Hash
	// 4. Signature of New Public Key + Tx Hash (concatenated)
	newPublicKey := hdw.PublicKeyFromIndex(arch.Wallet, arch.AccountIndex+1)
	pubKeyConcatTxHash := append(crypto.FromECDSAPub(newPublicKey)[1:], []byte(arweaveTxId)...)
	hash := crypto.Keccak256Hash(pubKeyConcatTxHash)
	assetIdSig, err := crypto.Sign(hash.Bytes(), arch.PrivateKey)
	if err != nil {
		arch.fileUploadError("Couldnt sign the arweave tx: "+err.Error(), "There was an error with the file.", http.StatusBadRequest, w)
		return
	}

	R, S, V := utility.SigRSV(assetIdSig)

	w.Header().Set("Content-Type", "application/json")
	response := ResponseToEmbalmer{
		NewPublicKey:    crypto.FromECDSAPub(newPublicKey)[1:],
		AssetDoubleHash: assetDoubleHash,
		AssetId:         arweaveTxId,
		V:               V,
		R:               R,
		S:               S,
	}

	log.Printf("Sending Response to Embalmer: %v", response)

	json.NewEncoder(w).Encode(response)
	arch.fileHandlerCheck()
}

// ListenForFile .
func (arch *Archaeologist) ListenForFile() {
	if !arch.IsServerRunning() {
		arch.InitServer()
		arch.StartServer()
	}
}

// IsServerRunning .
func (arch *Archaeologist) IsServerRunning() bool {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", arch.FilePort), timeout)
	return err == nil
}

// InitAndTestServer .
func (arch *Archaeologist) InitAndTestServer() {
	arch.InitServer()
	log.Printf("Testing Server...")
	go func() {
		if err := arch.Server.ListenAndServe(); err != nil {
			log.Fatalf("Could not start server. Please check that the port in CONFIG is set correctly and is open. Error: %v", err)
		}
	}()
	log.Printf("Server test sucessful, shutting down.")
	arch.Server.Shutdown(context.Background())
}

// InitServer .
func (arch *Archaeologist) InitServer() {
	sm := http.NewServeMux()
	sm.Handle("/ping", http.HandlerFunc(arch.pingHandler))
	sm.Handle("/file", http.HandlerFunc(arch.fileUploadHandler))
	arch.Server = &http.Server{Addr: "localhost:" + arch.FilePort, Handler: utility.LimitMiddleware(sm)}
}

// Start Server .
func (arch *Archaeologist) StartServer() {
	go func() {
		log.Printf("Server starting on %s:", arch.Server.Addr)
		if err := arch.Server.ListenAndServe(); err != nil {
			log.Println("Server shutting down:", err)
		}
	}()

	// Stop the server if the program exits
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop
	arch.ShutdownServer()
}

// ShutdownServer .
func (arch *Archaeologist) ShutdownServer() {
	if arch.IsServerRunning() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := arch.Server.Shutdown(ctx); err != nil {
			log.Println("Error shutting down http server (server may already be shutdown):", err)
		} else {
			log.Println("Server has been shutdown")
		}
	}
}

// IsArchSarcophagus returns true if the sarcophagus is in state
func (arch *Archaeologist) IsArchSarcophagus(doubleHash [32]byte) bool {
	_, ok := arch.Sarcophaguses[doubleHash]
	return ok
}

// RemoveArchSarcophagus deletes the sarcophagus from state if it exists
// also removes the file handler
func (arch *Archaeologist) RemoveArchSarcophagus(doubleHash [32]byte) {
	if arch.IsArchSarcophagus(doubleHash) {
		delete(arch.Sarcophaguses, doubleHash)

		_, ok := arch.FileHandlers[doubleHash]
		if ok {
			arch.fileHandlerCheck()
			delete(arch.FileHandlers, doubleHash)
		}
	}
}
