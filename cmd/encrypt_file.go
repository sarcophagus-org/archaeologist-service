package main

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const fileToEncrypt = "/tmp/test.txt"
const encryptedOutputFilePath = "/tmp/encrypted.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isHex(str string) bool {
	return strings.HasPrefix(str, "0x")
}

func archPrivateKey() *ecdsa.PrivateKey {
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Could not find config file. It should be setup under config/config.yml")
		} else {
			log.Fatalf("Could not read config file. Please check it is configured correctly. Error: %v \n", err)
		}
	}

	privateKey := viper.GetString("eth_private_key")
	log.Println("Private Key", privateKey)

	if isHex(privateKey) {
		privateKey = privateKey[2:]
	}

	ethPrivKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	return ethPrivKey
}

func archPublicKey() *btcec.PublicKey {
	privateKey := archPrivateKey()
	pub := privateKey.Public()
	publicKeyECDSA, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	pubKey, err := btcec.ParsePubKey(publicKeyBytes, btcec.S256())
	if err != nil {
		log.Fatalf("error casting public key to btcec: %v", err)
	}

	return pubKey
}

func main() {
	file, err := os.Open(fileToEncrypt)
	check(err)

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatalf("Issue copying file to buffer")
	}

	fileBytes := buf.Bytes()

	encryptedBytes, err := btcec.Encrypt(archPublicKey(), fileBytes)

	if err != nil {
		log.Fatalf("Error encrypting file: %v", err)
	}

	bigErr := ioutil.WriteFile(encryptedOutputFilePath, encryptedBytes, 0755)
	if bigErr != nil {
		log.Fatalf("Error writing file: %v", bigErr)
	}

	file.Close()
}
