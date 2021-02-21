// Collection of functions for initializing sessions
// that are used for the archaeologist to interact with the contract

package archaeologist

import (
	"context"
	"crypto/ecdsa"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

// initAuth .
func initAuth (privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = nil // uses nonce of pending state
	auth.Value = big.NewInt(0)
	auth.GasLimit = 0 // 0 estimates gas limit
	auth.GasPrice = nil // nil suggests gas price

	return auth
}

// NewSarcophagusSession .
func NewSarcophagusSession(ctx context.Context, sarcophagusContract *contracts.Sarcophagus, privateKey *ecdsa.PrivateKey) contracts.SarcophagusSession {
	auth := initAuth(privateKey)

	return contracts.SarcophagusSession{
		Contract: sarcophagusContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

// NewTokenSession .
func NewTokenSession(ctx context.Context, tokenContract *contracts.Token, privateKey *ecdsa.PrivateKey)  contracts.TokenSession {
	auth := initAuth(privateKey)

	return contracts.TokenSession{
		Contract: tokenContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}