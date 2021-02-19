package models

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ETH_NODE              string
	ETH_PRIVATE_KEY       string
	ARWEAVE_KEY_FILE      string
	ARWEAVE_MULTIPLIER    string
	ARWEAVE_NODE          string
	FILE_PORT             string
	ENDPOINT              string
	FEE_PER_BYTE          string
	MIN_BOUNTY            string
	MIN_DIGGING_FEE       string
	MAX_RESURRECTION_TIME string
	CONTRACT_ADDRESS      string
	TOKEN_ADDRESS         string
	ADD_TO_FREE_BOND      string
	REMOVE_FROM_FREE_BOND string
	PAYMENT_ADDRESS       string
	GAS_PRICE_OVERRIDE    string
	MNEMONIC              string
}

func (config *Config) LoadConfig(name string, path string, writeConfig bool) {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	viper.AddConfigPath("$GOPATH/bin")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Could not find config file. Please make sure it is setup under %v/%v", path, name)
		} else {
			log.Fatalf("Could not read config file. Please check it is configured correctly. Error: %v \n", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Could not load config file. Please check it is configured correctly. Error: %v \n", err)
	}

	// Write Free Bond Values to 0 in the config file.
	// They have already been loaded into the config struct.
	if writeConfig {
		viper.Set("ADD_TO_FREE_BOND", 0)
		viper.Set("REMOVE_FROM_FREE_BOND", 0)
		// viper.WriteConfig()
	}
}
