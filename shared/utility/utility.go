package utility

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
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

func ValidatePositiveNumber(num int64, numField string) int64 {
	if num <= 0 {
		log.Fatalf("%s must be greater than 0. Please check the value in the config file", numField)
	}

	return num
}

func ValidateIpAddress(ip string, ipField string) string {
	if net.ParseIP(ip) == nil {
		log.Fatalf("%s IP Address is invalid. Please check the value in the config file", ipField)
	}

	return ip
}

func ValidateTimeInFuture(unixTimestamp int64, timeField string) int64 {
	now := time.Now()
	unixNow := now.Unix()

	if unixTimestamp <= unixNow {
		log.Fatalf("%s must be in the future. Please check the value in the config file", timeField)
	}

	return unixTimestamp
}

func FileToBytes(file *os.File) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, file)

	return buf.Bytes(), err
}