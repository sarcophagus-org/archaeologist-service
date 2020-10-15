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
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
)

const fileToDecrypt = "/tmp/encrypted.txt"
const decryptedOutputFilePath = "/tmp/decryptedtwo.txt"

func archPrivateKeyString() string {
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

	if utility.IsHex(privateKey) {
		privateKey = privateKey[2:]
	}

	return privateKey
}

func archPrivateKeyECDSA(privateKey string) *ecdsa.PrivateKey {
	privateKeyECDSA, _ := crypto.HexToECDSA(privateKey)

	return privateKeyECDSA
}

func main() {
	file, err := os.Open(fileToDecrypt)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Fatalf("Issue copying file to buffer")
	}

	fileBytes := buf.Bytes()

	pkBytes := crypto.FromECDSA(archPrivateKeyECDSA(archPrivateKeyString()))

	// note that we already have corresponding pubKey
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	// Try decrypting and verify if it's the same message.
	plaintext, err := btcec.Decrypt(privKey, fileBytes)
	if err != nil {
		log.Fatalf("Could not decrypt file: %v", err)
	}

	bigErr := ioutil.WriteFile(decryptedOutputFilePath, plaintext, 0755)
	if bigErr != nil {
		log.Fatalf("Error writing file: %v", bigErr)
	}

	file.Close()
}