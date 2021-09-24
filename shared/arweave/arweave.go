package ar

import (
	"fmt"
	"github.com/everFinance/goar"
	"log"
	"math/big"
)

func ArweaveBalance(client *goar.Client, arWallet *goar.Wallet) *big.Float {
	balance, err := client.GetWalletBalance(arWallet.Address)
	if err != nil {
		log.Fatal("couldnt get arweave balance %s", balance)
	}

	return balance
}

func InitArweaveClient(arweaveNode string) (*goar.Client) {
	return goar.NewClient(arweaveNode)
}

func InitArweaveWallet(arweaveKeyFileName string, arweaveNode string) (*goar.Wallet, error) {
	wallet_, err := goar.NewWalletFromPath(arweaveKeyFileName, arweaveNode)

	if err != nil {
		return nil, fmt.Errorf("Could not load config value ARWEAVE_KEY_FILE. Please check the config.yml file Error: %v", err)
	}

	return wallet_, nil
}
