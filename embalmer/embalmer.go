package embalmer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

type Embalmer struct {
	Client                   *ethclient.Client
	ArchAddress              common.Address
	EmbalmerPrivateKey       *ecdsa.PrivateKey
	EmbalmerAddress          common.Address
	SarcoAddress             common.Address
	SarcophagusContract      *contracts.Sarcophagus
	SarcophagusTokenContract *contracts.Token
	ResurrectionTime         *big.Int
	StorageFee               *big.Int
	DiggingFee               *big.Int
	Bounty                   *big.Int
	ArweaveNode              string
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

func (embalmer *Embalmer) approveEmbalmerTransfer(session *contracts.TokenSession, approvalAmount *big.Int) {
	tx, err := session.Approve(
		embalmer.SarcoAddress,
		approvalAmount,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Approval Transaction for %v Sarcophagus Tokens successful. Transaction ID: %v", utility.ToDecimal(approvalAmount, 18), tx.Hash().Hex())
	log.Printf("Gas Used for Approval: %v", tx.Gas())

	err = ethereum.WaitMined(embalmer.Client, tx.Hash(), "Approval of Embalmer Transfer")

	if err != nil {
		log.Fatalf("There was an error mining the approval of embalmer transfer: %v", err)
	}
}

func (embalmer *Embalmer) CreateSarcophagus(recipientPrivateKey string, assetDoubleHashBytes [32]byte, sarcoName string) {
	/* Initialize recipient public key bytes */
	if utility.IsHex(recipientPrivateKey) {
		recipientPrivateKey = recipientPrivateKey[2:]
	}
	recipPrivKey, _ := crypto.HexToECDSA(recipientPrivateKey)
	pub := recipPrivKey.Public().(*ecdsa.PublicKey)
	recipientPublicKeyBytes := crypto.FromECDSAPub(pub)[1:]

	log.Println("***CREATING SARCOPHAGUS***")

	sarcoTokenSession := embalmer.NewSarcophagusTokenSession(context.Background())
	bountyPlusDiggingFee := big.NewInt(0).Add(embalmer.Bounty, embalmer.DiggingFee)
	log.Printf("bountyPlusDiggingFee: %v", utility.ToDecimal(bountyPlusDiggingFee, 18))

	approvalAmount := big.NewInt(0).Add(big.NewInt(0).Add(embalmer.Bounty, embalmer.DiggingFee), embalmer.StorageFee)
	log.Printf("Approval Amount::: %v", utility.ToDecimal(approvalAmount, 18))
	embalmer.approveEmbalmerTransfer(&sarcoTokenSession, approvalAmount)

	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	tx, err := sarcoSession.CreateSarcophagus(
		sarcoName,
		embalmer.ArchAddress,
		embalmer.ResurrectionTime,
		embalmer.StorageFee,
		embalmer.DiggingFee,
		embalmer.Bounty,
		assetDoubleHashBytes,
		recipientPublicKeyBytes,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error creating Sarcophagus: %v", err)
	}

	log.Printf("Create Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())

	err = ethereum.WaitMined(embalmer.Client, tx.Hash(), "Approval of Embalmer Create Sarcophagus")

	if err != nil {
		log.Fatalf("There was an error mining the approval of embalmer create sarcophagus: %v", err)
	}
}

func (embalmer *Embalmer) EncryptFileBytes(fileBytes []byte) []byte {
	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	contractArch, err := sarcoSession.Archaeologists(embalmer.ArchAddress)
	if err != nil {
		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed. Please check CONTRACT_ADDRESS is correct in the config file: %v", err)
	}

	currentPublicKeyBytes := append([]byte{4}, contractArch.CurrentPublicKey...)
	pubKeyEcdsa, err := crypto.UnmarshalPubkey(currentPublicKeyBytes)
	if err != nil {
		log.Fatalf("Error unmarshaling public key during embalmer update: %v", err)
	}

	pubKey := ecies.ImportECDSAPublic(pubKeyEcdsa)
	encryptedBytes, err := ecies.Encrypt(rand.Reader, pubKey, fileBytes, nil, nil)
	if err != nil {
		log.Fatalf("Error encrypting file: %v", err)
	}
	log.Printf("ENCRYPTED BYTES: %v", encryptedBytes)
	return encryptedBytes
}

func (embalmer *Embalmer) UpdateSarcophagus(assetDoubleHash [32]byte, fileBytes []byte) {
	log.Println("***UPDATING SARCOPHAGUS***")
	encryptedBytes := embalmer.EncryptFileBytes(fileBytes)

	body := &models.SarcoFile{
		FileType:  "txt",
		FileBytes: base64.StdEncoding.EncodeToString(encryptedBytes),
	}

	url := "http://127.0.0.1:8080/file"
	response, err := embalmer.SendFile(url, body)
	if err != nil {
		log.Fatalf("Error sending file: %v", err)
	}

	var responseToEmbalmer = new(models.ResponseToEmbalmer)
	err = json.Unmarshal(response, &responseToEmbalmer)
	if err != nil {
		log.Fatalf("Error: %v", string(response))
	}

	log.Printf("NewPublicKey: %v", responseToEmbalmer.NewPublicKey)
	log.Printf("AssetID: %v", responseToEmbalmer.AssetId)
	log.Printf("V: %v", responseToEmbalmer.V)
	log.Printf("R: %v", responseToEmbalmer.R)
	log.Printf("S: %v", responseToEmbalmer.S)

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

	err = ethereum.WaitMined(embalmer.Client, tx.Hash(), "Embalmer Update Sarcophagus")

	if err != nil {
		log.Fatalf("There was an error mining the approval of embalmer update sarcophagus: %v", err)
	}
}

func (embalmer *Embalmer) RewrapSarcophagus(assetDoubleHash [32]byte, resurrectionTime *big.Int) {
	log.Println("***REWRAPPING SARCOPHAGUS***")

	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	sarcoTokenSession := embalmer.NewSarcophagusTokenSession(context.Background())
	embalmer.approveEmbalmerTransfer(&sarcoTokenSession, embalmer.DiggingFee)

	tx, err := sarcoSession.RewrapSarcophagus(
		assetDoubleHash,
		resurrectionTime,
		embalmer.DiggingFee,
		embalmer.Bounty,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error rewrapping Sarcophagus: %v", err)
	}

	log.Printf("Rewrap Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())

	err = ethereum.WaitMined(embalmer.Client, tx.Hash(), "Embalmer Rewrap Sarcophagus")

	if err != nil {
		log.Fatalf("There was an error mining the approval of embalmer rewrap sarcophagus: %v", err)
	}
}

func (embalmer *Embalmer) CleanupSarcophagus(assetDoubleHash [32]byte) {
	log.Println("***CLEANING UP SARCOPHAGUS***")

	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	tx, err := sarcoSession.CleanUpSarcophagus(
		assetDoubleHash,
		embalmer.EmbalmerAddress,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error cleaning Sarcophagus: %v", err)
	}

	log.Printf("Clean Up Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func (embalmer *Embalmer) BurySarcophagus(assetDoubleHash [32]byte) {
	log.Println("***BURYING SARCOPHAGUS***")

	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	tx, err := sarcoSession.BurySarcophagus(
		assetDoubleHash,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error burying Sarcophagus: %v", err)
	}

	log.Printf("Bury Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func (embalmer *Embalmer) CancelSarcophagus(assetDoubleHash [32]byte) {
	log.Println("***CANCELLING SARCOPHAGUS***")

	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	tx, err := sarcoSession.CancelSarcophagus(
		assetDoubleHash,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error cancelling Sarcophagus: %v", err)
	}

	log.Printf("Cancel Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func (embalmer *Embalmer) SendFile(url string, body *models.SarcoFile) ([]byte, error) {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	request, err := http.NewRequest("POST", url, payloadBuf)

	if err != nil {
		log.Fatal(err)
	}

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
