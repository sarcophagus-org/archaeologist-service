package ar

import (
	"context"
	"fmt"
	"github.com/Dev43/arweave-go/api"
	"github.com/Dev43/arweave-go/transactor"
	"github.com/Dev43/arweave-go/wallet"
	"log"
)

func ArweaveBalance(client *api.Client, arWallet *wallet.Wallet) string {
	balance, err := client.GetBalance(context.Background(), arWallet.Address())
	if err != nil {
		log.Fatal("couldnt get arweave balance %s", balance)
	}

	return balance
}

func InitArweaveTransactor(arweaveNode string) (*transactor.Transactor, error) {
	ar, err := transactor.NewTransactor(arweaveNode)

	if err != nil {
		return nil, fmt.Errorf("Could not connect to arweave node. Error: %v\n", err)
	}

	return ar, nil
}

func InitArweaveWallet(arweaveKeyFileName string) (*wallet.Wallet, error) {
	wallet_ := wallet.NewWallet()

	if err := wallet_.LoadKeyFromFile(fmt.Sprintf(arweaveKeyFileName)); err != nil {
		return nil, fmt.Errorf("Could not load config value ARWEAVE_KEY_FILE. Please check the config.yml file Error: %v", err)
	}

	return wallet_, nil
}
