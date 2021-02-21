// Utility functions for interacting with the hdwallet library
// The hd wallet is used to generate new key pairs for each sarcophagus

package hdw

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
)

const derivationPathPrefix = "m/44'/60'/0'/0/"

// AccountFromIndex .
func AccountFromIndex(wallet *hdwallet.Wallet, index int) accounts.Account {
	account, err := wallet.Derive(DerivationPathFromIndex(index), false)
	if err != nil {
		log.Fatalf("There was an error creating an account: %v", err)
	}

	return account
}

// DerivationPathFromIndex .
func DerivationPathFromIndex(index int) accounts.DerivationPath {
	derivationPath, err := hdwallet.ParseDerivationPath(fmt.Sprintf("%s%d", derivationPathPrefix, index))
	if err != nil {
		log.Fatalf("There was an error parsing derivation path and creating account for hdwallet: %v", err)
	}

	return derivationPath
}

// PublicKeyFromIndex -- given an index of an account on the hd wallet, return the corresponding public key
func PublicKeyFromIndex(wallet *hdwallet.Wallet, index int) *ecdsa.PublicKey {
	account := AccountFromIndex(wallet, index)
	pubKey, err := wallet.PublicKey(account)
	if err != nil {
		log.Fatalf("There was an error getting public key from account index: %d, %v", index, err)
	}

	return pubKey
}

// PublicKeyBytesFromIndex .
func PublicKeyBytesFromIndex(wallet *hdwallet.Wallet, index int) []byte {
	pubKey := PublicKeyFromIndex(wallet, index)
	return crypto.FromECDSAPub(pubKey)[1:]
}

// PrivateKeyFromIndex -- given an index of an account on the hd wallet, return the corresponding private key
func PrivateKeyFromIndex(wallet *hdwallet.Wallet, index int) *ecdsa.PrivateKey {
	account := AccountFromIndex(wallet, index)
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		log.Fatalf("There was an error getting private key from account index: %d, %v", index, err)
	}

	return privateKey
}