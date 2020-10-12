package main

import (
	"flag"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"github.com/spf13/viper"
	"log"
)

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
	embalmer.InitEthClient(config.ETH_NODE)
	embalmer.InitKeys(config.ARCH_PRIVATE_KEY, config.EMBALMER_PRIVATE_KEY)
	embalmer.InitSarcophagusContract(config.CONTRACT_ADDRESS)
	embalmer.InitSarcophagusTokenContract(config.TOKEN_ADDRESS)

	// We may not need this flag. Setting up in case we need more control on what to call.
	sarcoFlag := flag.String("type", "create", "Create or Update a Sarcophagus")

	flag.Parse()

	if *sarcoFlag == "create" {
		embalmer.CreateSarcophagus(config.RECIPIENT_PRIVATE_KEY)
		log.Println("Embalmer Sarco Balance:", embalmer.EmbalmerSarcoBalance())
	}
}
