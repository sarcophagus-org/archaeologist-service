package utility

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
	"strings"
	"context"
)

func IsHex(str string) bool {
	return strings.HasPrefix(str, "0x")
}

func IsValidAddress(ethAddress string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(ethAddress)
}

func IsContract(address common.Address, client *ethclient.Client) bool {
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Could not get bytecode from contract address: %v", err)
	}

	isContract := len(bytecode) > 0
	return isContract
}

func PrivateKeyHexToECDSA(privateKey string) (*ecdsa.PrivateKey, error) {
	if IsHex(privateKey) {
		privateKey = privateKey[2:]
	}

	ethPrivateKey, err := crypto.HexToECDSA(privateKey)
	return ethPrivateKey, err
}

func PrivateToPublicKeyECDSA(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	pKey := privateKey.Public()
	publicKey, ok := pKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return publicKey
}