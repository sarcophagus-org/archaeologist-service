package main

import (
	"flag"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/spf13/viper"
	"log"
	"os"
)

const fileToDoubleHash = "/tmp/test.txt"

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
	file, err := os.Open(fileToDoubleHash)
	if err != nil {
		log.Fatalf("Couldnt open file to double hash")
	}

	fileBytes, _ := utility.FileToBytes(file)

	// We may not need this flag. Setting up in case we need more control on what to call.
	sarcoFlag := flag.String("type", "create", "Create or Update a Sarcophagus")

	flag.Parse()

	if *sarcoFlag == "create" {
		emb.CreateSarcophagus(config.RECIPIENT_PRIVATE_KEY, fileBytes)
		log.Println("Embalmer Sarco Balance:", emb.EmbalmerSarcoBalance())
	}
}
