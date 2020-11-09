package hdw

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
)

const derivationPathPrefix = "m/44'/60'/0'/0/"

func AccountFromIndex(wallet *hdwallet.Wallet, index int) accounts.Account {
	account, err := wallet.Derive(DerivationPathFromIndex(index), false)
	if err != nil {
		log.Fatalf("There was an error creating an account: %v", err)
	}

	return account
}

func DerivationPathFromIndex(index int) accounts.DerivationPath {
	derivationPath, err := hdwallet.ParseDerivationPath(fmt.Sprintf("%s%d", derivationPathPrefix, index))
	if err != nil {
		log.Fatalf("There was an error parsing derivation path and creating account for hdwallet: %v", err)
	}

	return derivationPath
}

func PublicKeyFromIndex(wallet *hdwallet.Wallet, index int) *ecdsa.PublicKey {
	account := AccountFromIndex(wallet, index)
	pubKey, err := wallet.PublicKey(account)
	if err != nil {
		log.Fatalf("There was an error getting public key from account index: %d, %v", index, err)
	}

	return pubKey
}