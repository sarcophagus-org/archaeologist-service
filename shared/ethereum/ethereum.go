package ethereum

import (
	"context"
	"log"
	"crypto/ecdsa"
	"math/big"

	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
)

var client *ethclient.Client
var ethPrivateKey *ecdsa.PrivateKey
var ethPublicKey *ecdsa.PublicKey
var ethAddress common.Address

func EthBalance() *big.Int {
	balance, _ := client.BalanceAt(context.Background(), ethAddress, nil)

	return balance
}

func ValidateContractAddress(contractAddress string){
	// Instantiate the contract and display its name
	token, err := NewToken(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	log.Println("Token name:", name)
}

func GetSuggestedGasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("couldn't get the suggested gas price", err)
	}
	return gasPrice, err
}

func InitEthVars(config *models.Config) {
	cli, err := ethclient.Dial(config.ETH_NODE)
	if err != nil {
		log.Fatal("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	ethPrivateKey, err = crypto.HexToECDSA(config.ETH_PRIVATE_KEY[2:])
	if err != nil {
		log.Fatal("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	pub := ethPrivateKey.Public()
	publicKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	client = cli
	ethPublicKey = publicKey
	ethAddress = common.HexToAddress(crypto.PubkeyToAddress(*ethPublicKey).Hex())
}
