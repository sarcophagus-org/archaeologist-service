package ethereum

import (
	"context"
	"log"
	"crypto/ecdsa"
	"os"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

var client *ethclient.Client
var ethPrivateKey *ecdsa.PrivateKey
var ethPublicKey *ecdsa.PublicKey
var ethAddress common.Address

func ArchBalance() *big.Int {
	balance, _ := client.BalanceAt(context.Background(), ethAddress, nil)

	return balance
}

func GetSuggestedGasPrice() (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("couldn't get the suggested gas price", err)
	}
	return gasPrice, err
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err = ethclient.Dial(os.Getenv("ETH_GATEWAY"))
	if err != nil {
		log.Fatal("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	ethPrivateKey, err  = crypto.HexToECDSA(os.Getenv("ETH_PRIVATE_KEY")[2:])
	if err != nil {
		log.Fatal(err)
	}

	pub := ethPrivateKey.Public()
	ethPublicKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	ethAddress = common.HexToAddress(crypto.PubkeyToAddress(*ethPublicKey).Hex())
}
