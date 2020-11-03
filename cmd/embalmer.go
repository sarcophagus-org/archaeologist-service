package main

import (
	"flag"
	"github.com/btcsuite/btcd/btcec"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
)

const fileDefault = "/tmp/test.txt"
const encryptedOutputFilePath = "/tmp/encrypted.txt"

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
		pubKey, err := btcec.ParsePubKey(emb.ArchPublicKeyBytes, btcec.S256())
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
