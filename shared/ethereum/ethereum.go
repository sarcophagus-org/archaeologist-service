package ethereum

import (
	"context"
	"log"
	"fmt"
	"crypto/ecdsa"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

var client *ethclient.Client
var ethPrivateKey *ecdsa.PrivateKey
var ethPublicKey *ecdsa.PublicKey
var ethAddress common.Address

func init() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := ethclient.Dial(os.Getenv("ETH_GATEWAY"))
	if err != nil {
		log.Fatal("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	ethPrivateKey, err  := crypto.HexToECDSA(os.Getenv("ETH_PRIVATE_KEY")[2:])
	if err != nil {
		log.Fatal(err)
	}

	Pub := ethPrivateKey.Public()
	ethPublicKey, ok := Pub.(*ecdsa.Pub)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	ethAddress := crypto.PubkeyToAddress(*ethPublicKey).Hex()

	balance, _ := client.BalanceAt(ctx, ethAddress, nil)
	fmt.Printf("Balance: %d\n",balance)
}
