package arweave

import (
	"context"
	"log"
	"fmt"

	"github.com/Dev43/arweave-go/wallet"
	"github.com/Dev43/arweave-go/api"
)

func init() {
	// create a new wallet instance
	w := wallet.NewWallet()
	// extract the key from the wallet instance
	err := w.LoadKeyFromFile("config/arweave.json")
	if err != nil {
		log.Fatal(err)
	}

	ipAddress := "http://localhost:8000/arweave"
	c, err := api.Dial(ipAddress)
	if err != nil {
		log.Fatal(err)
	}

	balance, err := c.GetBalance(context.Background(), w.Address())
	if err != nil {
		log.Fatal("couldnt get balance %s", balance)
	}
	fmt.Printf("Arweave Balance: %s\n", balance)
}
