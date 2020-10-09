package arweave

import (
	"context"
	"log"
	"fmt"

	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/Dev43/arweave-go/wallet"
	"github.com/Dev43/arweave-go/api"
)

var client *api.Client
var arWallet *wallet.Wallet

func ArweaveBalance() string {
	balance, err := client.GetBalance(context.Background(), arWallet.Address())
	if err != nil {
		log.Fatal("couldnt get arweave balance %s", balance)
	}

	return balance
}

func InitArweaveVars(config *models.Config) {
	w := wallet.NewWallet()

	if err := w.LoadKeyFromFile(fmt.Sprintf("config/%s", config.ARWEAVE_KEY_FILE)); err != nil {
		log.Fatal("Could not load config value ARWEAVE_KEY_FILE. Please check the config.yml file Error:", err)
	}

	cli, err := api.Dial(config.ARWEAVE_NODE)
	if err != nil {
		log.Fatal("Could not connect to arweave node. Error: %v\n", err)
	}

	client = cli
	arWallet = w
}
