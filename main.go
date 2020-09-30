package main

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

func publicKeyECDSAFromPrivKey(privKey string) *ecdsa.PublicKey {
	privKeyECDSA, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return publicKeyECDSA
}

func addressFromPrivKey(privKey string) string {
	publicKeyECDSA := publicKeyECDSAFromPrivKey(privKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address
}

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	client, err := ethclient.Dial(os.Getenv("ETH_GATEWAY"))
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	archaeologistAddressHex := addressFromPrivKey(os.Getenv("ETH_PRIVATE_KEY")[2:])
	archaeologistAddress := common.HexToAddress(archaeologistAddressHex)
	balance, _ := client.BalanceAt(ctx, archaeologistAddress, nil)
	fmt.Printf("Balance: %d\n",balance)
}
