package models

import (
	"context"
	"crypto/ecdsa"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type Archaeologist struct {
	Client             *ethclient.Client
	PrivateKey         *ecdsa.PrivateKey
	PublicKey          *ecdsa.PublicKey
	PublicKeyBytes     []byte
	ArchAddress        common.Address
	SarcoAddress       common.Address
	SarcoSession       contracts.SarcophagusSession
	SarcoTokenAddress  common.Address
	TokenSession       contracts.TokenSession
	FreeBond           int64
	FeePerByte         int64
	MinBounty          int64
	MinDiggingFee      int64
	MaxResurectionTime int64
}

func (arch *Archaeologist) SarcoBalance() *big.Int {
	balance, err := arch.TokenSession.BalanceOf(arch.ArchAddress)

	if err != nil {
		log.Fatalf("Could not get sarcophagus balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func (arch *Archaeologist) EthBalance() *big.Int {
	balance, err := arch.Client.BalanceAt(context.Background(), arch.ArchAddress, nil)

	if err != nil {
		log.Fatalf("Could not get eth balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func (arch *Archaeologist) WithdrawBond(bondToWithdraw *big.Int) {
	tx, err := arch.SarcoSession.WithdrawalBond(bondToWithdraw)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Withdrawing Bond: %v \n Config value REMOVE_FROM_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Withdrawal of %v Sarco Tokens successful. Transaction ID: %v", bondToWithdraw, tx.Hash().Hex())
	log.Printf("Gas Used for Withdrawal: %v", tx.Gas())
}
