package main

import (
	"flag"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/spf13/viper"
	"log"
)

func loadEmbalmerConfig() *models.Config {
	viper.SetConfigName("embalmer_config")
	viper.AddConfigPath("../config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var config models.Config

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Could not find config file. It should be setup under config/config.yml")
		} else {
			log.Fatalf("Could not read config file. Please check it is configured correctly. Error: %v \n", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Could not load config file. Please check it is configured correctly. Error: %v \n", err)
	}

	return &config
}

func main(){
	config := loadEmbalmerConfig()
	embalmer.InitEthClient(config.ETH_NODE)
	embalmer.InitKeys(config.ARCH_PRIVATE_KEY, config.EMBALMER_PRIVATE_KEY)
	embalmer.InitSarcophagusContract(config.CONTRACT_ADDRESS)

	// We may not need this flag. Setting up in case we need more control on what to call.
	sarcoFlag := flag.String("type", "create", "Create or Update a Sarcophagus")

	flag.Parse()

	if *sarcoFlag == "create" {
		embalmer.CreateSarcophagus()
	}
}
