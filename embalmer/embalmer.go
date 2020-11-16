package embalmer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Embalmer struct {
	Client                   *ethclient.Client
	ArchAddress              common.Address
	EmbalmerPrivateKey       *ecdsa.PrivateKey
	EmbalmerAddress          common.Address
	SarcoAddress             common.Address
	SarcophagusContract      *contracts.Sarcophagus
	SarcophagusTokenContract *contracts.Token
	ResurrectionTime         int64
	StorageFee               int64
	DiggingFee               int64
	Bounty                   int64
}

func (embalmer *Embalmer) initAuth() *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(embalmer.EmbalmerPrivateKey)
	auth.Nonce = nil // uses nonce of pending state
	auth.Value = big.NewInt(0)
	auth.GasLimit = 0   // 0 estimates gas limit
	auth.GasPrice = nil // nil suggests gas price

	return auth
}

func (embalmer *Embalmer) NewSarcophagusSession(ctx context.Context) (session contracts.SarcophagusSession) {
	auth := embalmer.initAuth()

	return contracts.SarcophagusSession{
		Contract:     embalmer.SarcophagusContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

func (embalmer *Embalmer) NewSarcophagusTokenSession(ctx context.Context) (session contracts.TokenSession) {
	auth := embalmer.initAuth()

	return contracts.TokenSession{
		Contract:     embalmer.SarcophagusTokenContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

func (embalmer *Embalmer) EmbalmerSarcoBalance() *big.Int {
	balance, err := embalmer.SarcophagusTokenContract.BalanceOf(&bind.CallOpts{}, embalmer.EmbalmerAddress)

	if err != nil {
		log.Fatalf("Could not get sarcophagus balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func (embalmer *Embalmer) approveCreateSarcophagusTransfer(session *contracts.TokenSession, approvalAmount *big.Int) {
	tx, err := session.Approve(
		embalmer.SarcoAddress,
		approvalAmount,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Approval Transaction for %v Create Sarcophagus Tokens successful. Transaction ID: %v", approvalAmount, tx.Hash().Hex())
	log.Printf("Gas Used for Approval: %v", tx.Gas())
}

func (embalmer *Embalmer) CreateSarcophagus(recipientPrivateKey string, assetDoubleHashBytes [32]byte) {
	/* Initialize recipient public key bytes */
	if utility.IsHex(recipientPrivateKey) {
		recipientPrivateKey = recipientPrivateKey[2:]
	}
	recipPrivKey, _ := crypto.HexToECDSA(recipientPrivateKey)
	pub := recipPrivKey.Public().(*ecdsa.PublicKey)
	recipientPublicKeyBytes := crypto.FromECDSAPub(pub)[1:]

	log.Println("***CREATING SARCOPHAGUS***")

	sarcoTokenSession := embalmer.NewSarcophagusTokenSession(context.Background())
	approvalAmount := new(big.Int).Add(new(big.Int).Add(big.NewInt(embalmer.Bounty), big.NewInt(embalmer.DiggingFee)), big.NewInt(embalmer.StorageFee))
	embalmer.approveCreateSarcophagusTransfer(&sarcoTokenSession, approvalAmount)

	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	tx, err := sarcoSession.CreateSarcophagus(
		"My Sarcophagus",
		embalmer.ArchAddress,
		big.NewInt(embalmer.ResurrectionTime),
		big.NewInt(embalmer.StorageFee),
		big.NewInt(embalmer.DiggingFee),
		big.NewInt(embalmer.Bounty),
		assetDoubleHashBytes,
		recipientPublicKeyBytes,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error creating Sarcophagus: %v", err)
	}

	log.Printf("Create Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func (embalmer *Embalmer) UpdateSarcophagus(assetDoubleHash [32]byte, filename string) {
	log.Println("***UPDATING SARCOPHAGUS***")
	url := "http://127.0.0.1:8080/file"
	response, err := embalmer.SendFile(url, filename, "file")
	if err != nil {
		log.Fatalf("Error sending file:", err)
	}

	var responseToEmbalmer = new(models.ResponseToEmbalmer)
	err = json.Unmarshal(response, &responseToEmbalmer)
	if err != nil {
		log.Fatalf("Error: %v", string(response))
	}

	log.Printf("NewPublicKey:", responseToEmbalmer.NewPublicKey)
	log.Printf("AssetID:", responseToEmbalmer.AssetId)
	log.Printf("V", responseToEmbalmer.V)
	log.Printf("R", responseToEmbalmer.R)
	log.Printf("S", responseToEmbalmer.S)

	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	tx, err := sarcoSession.UpdateSarcophagus(
		responseToEmbalmer.NewPublicKey,
		assetDoubleHash,
		responseToEmbalmer.AssetId,
		responseToEmbalmer.V,
		responseToEmbalmer.R,
		responseToEmbalmer.S,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error updating Sarcophagus: %v", err)
	}

	log.Printf("Update Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func (embalmer *Embalmer) SendFile(url string, filename string, filetype string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(part, file)
	writer.Close()
	request, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return content, err
}
