// Not used by the service, for testing purposes only

package embalmer

import (
	"github.com/spf13/viper"
	"log"
)

type EmbalmerConfig struct {
	EMBALMER_PRIVATE_KEY string
	ARCH_PRIVATE_KEY string
	RECIPIENT_PRIVATE_KEY string
	ETH_NODE string
	CONTRACT_ADDRESS string
	TOKEN_ADDRESS string
	RESURRECTION_TIME int64
	STORAGE_FEE string
	DIGGING_FEE string
	BOUNTY string
	ARWEAVE_NODE string
}

func (config *EmbalmerConfig) LoadEmbalmerConfig(name string, path string) {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Could not find config file. It should be setup under config/embalmer_config.yml")
		} else {
			log.Fatalf("Could not read embalmer config file. Please check it is configured correctly. Error: %v \n", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Could not load config file. Please check it is configured correctly. Error: %v \n", err)
	}
}