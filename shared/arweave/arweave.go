package arweave

import (
	"context"
	"github.com/Dev43/arweave-go/api"
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
