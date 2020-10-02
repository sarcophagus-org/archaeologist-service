package main

import (
	"fmt"

	c "github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/spf13/viper"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	_ "github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
)

var config c.Config

func loadConfig(){
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Could not find config file. It should be setup under config/config.yml"))
		} else {
			panic(fmt.Errorf("Could not read config file. Please check it is configured correctly. Error: %s \n", err))
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Could not read config file. Please check it is configured correctly. Error: %s \n", err))
	}
}

func main(){
	loadConfig()
	fmt.Println("Eth Node is\t", config.ETH_NODE)
	fmt.Println("Eth Private Key is\t\t", config.ETH_PRIVATE_KEY)
	balance := ethereum.ArchBalance()
	fmt.Printf("Balance: %d\n", balance)
}
