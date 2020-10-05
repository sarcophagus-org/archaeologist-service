package ethereum

import (
	"context"
	"log"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"

	sarcophagus "github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
)

var client *ethclient.Client
var ethPrivateKey *ecdsa.PrivateKey
var ethPublicKey *ecdsa.PublicKey
var ethAddress common.Address
var sarcophagusContract *sarcophagus.Sarcophagus

func EthBalance() *big.Int {
	balance, _ := client.BalanceAt(context.Background(), ethAddress, nil)

	return balance
}

func ArchaeologistCount() *big.Int {
	archCount, err := sarcophagusContract.ArchaeologistCount(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve archaeologist count: %v", err)
	}

	return archCount
}

func InitSarcophagusContract(contractAddress string){
	sarcoContract, err := sarcophagus.NewSarcophagus(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate Sarcophagus contract: %v", err)
	}

	sarcophagusContract = sarcoContract
}

func GetSuggestedGasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("couldn't get the suggested gas price", err)
	}
	return gasPrice, err
}

func InitEthClient(ethNode string) {
	cli, err := ethclient.Dial(ethNode)
	if err != nil {
		log.Fatal("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	client = cli
}

func InitEthKeysAndAddress(privateKey string) {
	ethPrivKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	pub := ethPrivKey.Public()
	publicKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	ethPrivateKey = ethPrivKey
	ethPublicKey = publicKey
	ethAddress = common.HexToAddress(crypto.PubkeyToAddress(*ethPublicKey).Hex())
}
