package ethereum

import (
	"context"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func initAuth () *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(archPrivateKey)
	auth.Nonce = nil // uses nonce of pending state
	auth.Value = big.NewInt(0)
	auth.GasLimit = 0 // estimate gas limit
	auth.GasPrice = GetSuggestedGasPrice() // estimates gas price

	return auth
}

func GetSuggestedGasPrice() *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("couldn't get the suggested gas price: %v", err)
	}
	return gasPrice
}

func NewSarcophagusSession(ctx context.Context) (session contracts.SarcophagusSession) {
	auth := initAuth()

	return contracts.SarcophagusSession{
		Contract: sarcophagusContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

func NewSarcophagusTokenSession(ctx context.Context) (session contracts.TokenSession) {
	auth := initAuth()

	return contracts.TokenSession{
		Contract: sarcophagusTokenContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}