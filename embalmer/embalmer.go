package embalmer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/miguelmota/go-solidity-sha3"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var client *ethclient.Client
var archPublicKeyBytes []byte
var embalmerPrivateKey *ecdsa.PrivateKey
var sarcophagusContract *contracts.Sarcophagus

func initAuth () *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(embalmerPrivateKey)
	auth.Nonce = nil // uses nonce of pending state
	auth.Value = big.NewInt(0)
	auth.GasLimit = 0 // 0 estimates gas limit
	auth.GasPrice = nil // nil suggests gas price

	return auth
}

func NewSarcophagusSession(ctx context.Context) (session contracts.SarcophagusSession) {
	auth := initAuth()

	return contracts.SarcophagusSession{
		Contract: sarcophagusContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

func InitSarcophagusContract(contractAddress string) {
	address := common.HexToAddress(contractAddress)
	sarcoContract, err := contracts.NewSarcophagus(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Sarcophagus contract: %v", err)
	}

	sarcophagusContract = sarcoContract
}

func InitEthClient(ethNode string) {
	cli, err := ethclient.Dial(ethNode)
	if err != nil {
		log.Fatalf("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	client = cli
}

func InitKeys(archPrivateKey string, embalmerPrivKey string) {
	ethPrivKey, err := crypto.HexToECDSA(embalmerPrivKey)
	ethPrivKey, err := crypto.HexToECDSA(archPrivateKey)
	if err != nil {
		log.Fatalf("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	pub := ethPrivKey.Public()
	publicKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	archPublicKeyBytes = crypto.FromECDSAPub(publicKey)[1:]
}

func CreateSarcophagus() {
	/* Initialize recipient public key bytes */
	privateKey := "3ceff89a53061b0098f23794ce2874dce8f6829f53d961cd34aa665728bb6fe8"
	ethPrivKey, _ := crypto.HexToECDSA(privateKey)
	pub := ethPrivKey.Public()
	pubKey, _ := pub.(*ecdsa.PublicKey)
	recipientPublicKeyBytes := crypto.FromECDSAPub(pubKey)[1:]

	session := ethereum.NewSarcophagusSession(context.Background())

	/* Sarc init values. */
	resurrectionTime := big.NewInt(2000)
	storageFee := big.NewInt(2000)
	diggingFee := big.NewInt(2000)
	bounty := big.NewInt(2000)

	/* Setup Asset Double Hash */
	assetSingleHash := solsha3.SoliditySHA3(
		// types
		[]string{"string"},

		// values
		[]interface{}{
			"my silly message",
		},
	)
	assetDoubleHash := solsha3.SoliditySHA3(
		// types
		[]string{"bytes32"},

		// values
		[]interface{}{
			assetSingleHash,
		},
	)

	assetDoubleHashString := hex.EncodeToString(assetDoubleHash)
	var assetDoubleHashBytes [32]byte
	copy(assetDoubleHashBytes[:], assetDoubleHashString)

	log.Println("***CREATING SARCOPHAGUS***")
	tx, err := session.CreateSarcophagus(
	"My Sarcophagus",
		archPublicKeyBytes,
		resurrectionTime,
		storageFee,
		diggingFee,
		bounty,
		assetDoubleHashBytes,
		recipientPublicKeyBytes,
	)
}