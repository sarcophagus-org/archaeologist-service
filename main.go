package main

import (
	"log"
	"fmt"

	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/spf13/viper"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
)

func loadConfig() *models.Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var config models.Config

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Could not find config file. It should be setup under config/config.yml")
		} else {
			log.Fatal("Could not read config file. Please check it is configured correctly. Error: %s \n", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Could not load config file. Please check it is configured correctly. Error: %s \n", err)
	}

	return &config
}

func validateConfig(config *models.Config){
	ethereum.InitEthKeysAndAddress(config.ETH_PRIVATE_KEY[2:])
	ethereum.InitEthClient(config.ETH_NODE)
	ethereum.InitSarcophagusContract(config.CONTRACT_ADDRESS)
	ethereum.InitSarcophagusTokenContract(config.TOKEN_ADDRESS)
	arweave.InitArweaveVars(config)
}

func main(){
	config := loadConfig()
	validateConfig(config)

	ethBalance := ethereum.EthBalance()
	fmt.Printf("Eth Balance: %d\n", ethBalance)

	arweaveBalance := arweave.ArweaveBalance()
	fmt.Printf("Arweave Balance: %s\n", arweaveBalance)

	archCount := ethereum.ArchaeologistCount()
	log.Println("Archaeologist Count:", archCount)

	tokenName := ethereum.TokenName()
	log.Println("Token name:", tokenName)
}
