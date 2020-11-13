package main

import (
	"flag"
	"github.com/btcsuite/btcd/btcec"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"context"
)

const fileDefault = "/tmp/test.txt"
const encryptedOutputFilePath = "/tmp/encrypted.txt"

/* TODO: Generate new file from random string */
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func loadEmbalmerConfig() *embalmer.EmbalmerConfig {
	viper.SetConfigName("embalmer_config")
	viper.AddConfigPath("../embalmer")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var embalmerConfig embalmer.EmbalmerConfig

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Could not find config file. It should be setup under config/embalmer_config.yml")
		} else {
			log.Fatalf("Could not read embalmer config file. Please check it is configured correctly. Error: %v \n", err)
		}
	}

	if err := viper.Unmarshal(&embalmerConfig); err != nil {
		log.Fatalf("Could not load config file. Please check it is configured correctly. Error: %v \n", err)
	}

	return &embalmerConfig
}

func main(){
	config := loadEmbalmerConfig()
	emb := new(embalmer.Embalmer)

	embalmer.InitEmbalmer(emb, config)

	assetPathFlag := flag.String("file", fileDefault, "File to use as payload for sarcophagus")
	typeFlag := flag.String("type", "create", "Create or Update a Sarcophagus")

	flag.Parse()

	file, err := os.Open(*assetPathFlag)
	if err != nil {
		log.Fatalf("Couldnt open file to double hash")
	}
	defer file.Close()

	fileBytes, _ := utility.FileToBytes(file)
	assetDoubleHashBytes := embalmer.FileBytesToDoubleHashBytes(fileBytes)

	if *typeFlag == "create" {
		emb.CreateSarcophagus(config.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytes)
		log.Println("Embalmer Sarco Balance:", emb.EmbalmerSarcoBalance())
	}

	if *typeFlag == "update" {
		sarcoSession := emb.NewSarcophagusSession(context.Background())
		contractArch, err := sarcoSession.Archaeologists(emb.ArchAddress)
		if err != nil {
			log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed. Please check CONTRACT_ADDRESS is correct in the config file: %v", err)
		}

		currentPublicKeyBytes := append([]byte{4}, contractArch.CurrentPublicKey...)
		pubKeyEcdsa, err := crypto.UnmarshalPubkey(currentPublicKeyBytes)
		if err != nil {
			log.Fatalf("Error unmarshaling public key during embalmer update:", err)
		}

		pubkeyBytes := crypto.FromECDSAPub(pubKeyEcdsa)
		pubKey, err := btcec.ParsePubKey(pubkeyBytes, btcec.S256())
		if err != nil {
			log.Fatalf("error casting public key to btcec: %v", err)
		}
		encryptedBytes, err := btcec.Encrypt(pubKey, fileBytes)
		if err != nil {
			log.Fatalf("Error encrypting file: %v", err)
		}

		bigErr := ioutil.WriteFile(encryptedOutputFilePath, encryptedBytes, 0755)
		if bigErr != nil {
			log.Fatalf("Error writing file: %v", bigErr)
		}
		emb.UpdateSarcophagus(assetDoubleHashBytes, encryptedOutputFilePath)
	}
}
