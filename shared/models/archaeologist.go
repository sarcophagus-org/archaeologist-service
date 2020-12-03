package models

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/Dev43/arweave-go"
	"github.com/Dev43/arweave-go/transactor"
	"github.com/Dev43/arweave-go/tx"
	"github.com/Dev43/arweave-go/wallet"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"io/ioutil"
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
	ArweaveWallet         *wallet.Wallet
	ArweaveTransactor     *transactor.Transactor
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
	Server				  *http.Server
	Sarcophaguses         map[[32]byte]*big.Int
	FileHandlers		  map[[32]byte]*big.Int
}

const (
	MB = 1 << 20
)

func (arch *Archaeologist) SarcoBalance() *big.Int {
	balance, err := arch.TokenSession.BalanceOf(arch.ArchAddress)

	if err != nil {
		log.Fatalf("Could not get sarcophagus balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func (arch *Archaeologist) EthBalance() *big.Int {
	balance, err := arch.Client.BalanceAt(context.Background(), arch.ArchAddress, nil)

	if err != nil {
		log.Fatalf("Could not get eth balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func (arch *Archaeologist) WithdrawBond(bondToWithdraw *big.Int) {
	txn, err := arch.SarcoSession.WithdrawBond(bondToWithdraw)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Withdrawing Bond: %v \n Config value REMOVE_FROM_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Withdrawal of %v Sarco Tokens successful. Transaction ID: %v", bondToWithdraw, txn.Hash().Hex())
	log.Printf("Gas Used for Withdrawal: %v", txn.Gas())
}

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

	arch.FreeBond = big.NewInt(0)
	log.Printf("Register Archaeologist Successful. Transaction ID: %s", txn.Hash().Hex())
	log.Printf("Gas Used: %v", txn.Gas())
}

func (arch *Archaeologist) UpdateArchaeologist() {
	log.Println("***UPDATING ARCHAEOLOGIST***")
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

	arch.FreeBond = big.NewInt(0)
	log.Printf("Update Archaeologist Successful. Transaction ID: %s", txn.Hash().Hex())
	log.Printf("Gas Used: %v", txn.Gas())
}

func (arch *Archaeologist) ApproveFreeBondTransfer() {
	archSarcoBalance := arch.SarcoBalance()

	// Check if archSarcoBalance < freeBond
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

	log.Printf("Approval Transaction for %v Sarco Tokens successful. Transaction ID: %v", arch.FreeBond, txn.Hash().Hex())
	log.Printf("Gas Used for Approval: %v", txn.Gas())
}

func (arch *Archaeologist) CreateArweaveTransaction(ctx context.Context, w arweave.WalletSigner, amount string, data []byte, target string) (*tx.Transaction, error) {
	tr := arch.ArweaveTransactor
	lastTx, err := tr.Client.LastTransaction(ctx, w.Address())
	if err != nil {
		return nil, err
	}

	price, err := tr.Client.GetReward(ctx, []byte(data))
	if err != nil {
		return nil, err
	}

	// Non encoded transaction fields
	txn := tx.NewTransaction(
		lastTx,
		w.PubKeyModulus(),
		amount,
		target,
		data,
		price,
	)

	return txn, nil
}

func (arch *Archaeologist) UploadFileToArweave(fileBytes []byte) (*tx.Transaction, error) {
	// create a transaction
	ar := arch.ArweaveTransactor
	w := arch.ArweaveWallet

	/*
		Arweave Transaction:
		Amount and Target are blank, as we aren't sending arweave tokens to anyone
	*/
	log.Printf("uploading file bytes: %v", fileBytes)
	txBuilder, err := arch.CreateArweaveTransaction(context.TODO(), w, "0", fileBytes, "")
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

func (arch *Archaeologist) fileHandlerCheck() {
	fileHandlerLen := len(arch.FileHandlers)
	if fileHandlerLen <= 1 {
		arch.FileHandlers = map[[32]byte]*big.Int{}
		arch.ShutdownServer()
	}
}

func (arch *Archaeologist) fileUploadError(logMsg string, httpErrMsg string, httpErrType int, w http.ResponseWriter) {
	log.Printf(logMsg)
	http.Error(w, httpErrMsg, httpErrType)

	arch.fileHandlerCheck()
}

func (arch *Archaeologist) fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		arch.fileUploadError("File handler received non-post method, exiting.", "Method not allowed", http.StatusMethodNotAllowed, w)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		arch.fileUploadError("Couldnt receive file:" + err.Error(), "There was an error receiving your file:"+err.Error(), http.StatusBadRequest, w)
		return
	}

	defer file.Close()
	fileBytes, _ := ioutil.ReadAll(file)
	fileByteLen := len(fileBytes)

	/* Validate Size. */
	if fileByteLen > (3 * MB) {
		arch.fileUploadError("File was too large to receive. Size:" + string(fileByteLen), "The file sent is larger than the limit of 3MB.", http.StatusBadRequest, w)
		return
	}

	log.Print("Decrypting file...")
	/* Decrypt the outer layer of file */
	currentPrivateKey := hdw.PrivateKeyFromIndex(arch.Wallet, arch.AccountIndex)
	decryptedFileBytes, err := utility.DecryptFile(fileBytes, currentPrivateKey)
	assetDoubleHash := utility.FileBytesToDoubleHashBytes(decryptedFileBytes)
	if err != nil {
		arch.fileUploadError("Error decrypting file:" + err.Error(), "The file cannot be decrypted by archaeologist. Confirm it was encrypted with the correct public key.", http.StatusBadRequest, w)
		return
	}

	/* Validate the inner layer double hash matches a double hash */
	storageFee, ok := arch.FileHandlers[assetDoubleHash]
	if !ok{
		errMsg := "The double hash of the file does not match any open double hashes."
		arch.fileUploadError(errMsg, errMsg, http.StatusBadRequest, w)
		return
	}

	/* Validate Storage Fee is sufficient */
	storageExpectation := new(big.Int).Mul(big.NewInt(int64(fileByteLen)), arch.FeePerByte)
	if storageExpectation.Cmp(storageFee) == 1  {
		errMsg := fmt.Sprintf("The storage fee is not enough. Expected storage fee of at least: %v, storage fee was: %v", storageExpectation, storageFee)
		arch.fileUploadError(errMsg, errMsg, http.StatusBadRequest, w)
		return
	}

	log.Printf("File %s was validated successfully:", header.Filename)

	arweaveTx, err := arch.UploadFileToArweave(fileBytes)
	if err != nil {
		errMsg := fmt.Sprintf("There was an error with the file. Error: %v", err)
		arch.fileUploadError(errMsg, errMsg, http.StatusBadRequest, w)
		return
	}

	log.Printf("Transaction from arweave successful: %v", arweaveTx.Hash())

	/* Respond to embalmer with:
		New Public Key
	   	Arweave Tx Hash
		Signature of New Public Key + Tx Hash
	*/
	arweaveTxHash := arweaveTx.Hash()
	newPublicKey := hdw.PublicKeyFromIndex(arch.Wallet, arch.AccountIndex+1)
	pubKeyConcatTxHash := append(crypto.FromECDSAPub(newPublicKey)[1:], []byte(arweaveTxHash)...)
	hash := crypto.Keccak256Hash(pubKeyConcatTxHash)
	assetIdSig, err := crypto.Sign(hash.Bytes(), arch.PrivateKey)
	if err != nil {
		arch.fileUploadError("Couldnt sign the arweave tx: " + err.Error(), "There was an error with the file.", http.StatusBadRequest, w)
		return
	}

	R, S, V := utility.SigRSV(assetIdSig)

	w.Header().Set("Content-Type", "application/json")
	response := ResponseToEmbalmer{
		NewPublicKey:    crypto.FromECDSAPub(newPublicKey)[1:],
		AssetId:         arweaveTxHash,
		AssetDoubleHash: assetDoubleHash,
		V:               V,
		R:               R,
		S:               S,
	}

	json.NewEncoder(w).Encode(response)
	arch.ShutdownServer()
}

func (arch *Archaeologist) ListenForFile() {
	if !arch.IsServerRunning() {
		arch.InitServer()
		arch.StartServer()
	}
}

func (arch *Archaeologist) IsServerRunning() bool {
	timeout := 1 * time.Second
	_, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", arch.FilePort), timeout)
	if err != nil {
		return false
	}
	return true
}

func (arch *Archaeologist) InitServer() {
	sm := http.NewServeMux()
	sm.Handle("/file", http.HandlerFunc(arch.fileUploadHandler))
	arch.Server = &http.Server{Addr: ":" + arch.FilePort, Handler: utility.LimitMiddleware(sm)}
}

func (arch *Archaeologist) StartServer() {
	go func() {
		log.Printf("Listening for file on %s:", arch.Server.Addr)
		if err := arch.Server.ListenAndServe(); err != nil {
			log.Println("Error starting http server:", err)
		}
	}()

	// Stop the server if the program exits
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (pkill -2)
	<-stop
	arch.ShutdownServer()
}

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

func (arch *Archaeologist) IsArchSarcophagus(doubleHash [32]byte) bool {
	_, ok := arch.Sarcophaguses[doubleHash]
	return ok
}

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