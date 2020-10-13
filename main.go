package main

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/spf13/viper"
	"log"
)

func loadConfig() *models.Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
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

	// Write Free Bond Values to 0 in the config file.
	// They have already been loaded into the config struct.
	viper.Set("ADD_TO_FREE_BOND", 0)
	viper.Set("REMOVE_FROM_FREE_BOND", 0)
	viper.WriteConfig()

	return &config
}

func validateConfig(config *models.Config){
	ethereum.SetFreeBond(config.ADD_TO_FREE_BOND, config.REMOVE_FROM_FREE_BOND)
	ethereum.InitEthClient(config.ETH_NODE)
	ethereum.InitEthKeysAndAddress(config.ETH_PRIVATE_KEY, config.PAYMENT_ADDRESS)
	ethereum.InitSarcophagusContract(config.CONTRACT_ADDRESS)
	ethereum.InitSarcophagusTokenContract(config.TOKEN_ADDRESS)
	arweave.InitArweaveVars(config)
}

func main(){
	config := loadConfig()
	validateConfig(config)
	ethereum.RegisterOrUpdateArchaeologist(config)

	log.Printf("Eth Balance: %v", ethereum.ArchEthBalance())
	log.Printf("Sarco Token Balance: %v", ethereum.ArchSarcoBalance())
	log.Println("Arweave Balance:", arweave.ArweaveBalance())

	ethereum.EventsSubscribe()
}
