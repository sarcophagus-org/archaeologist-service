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
	Endpoint           string
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

func (arch *Archaeologist) RegisterArchaeologist() {
	log.Println("***REGISTERING ARCHAEOLOGIST***")
	tx, err := arch.SarcoSession.RegisterArchaeologist(
		arch.PublicKeyBytes,
		arch.Endpoint,
		arch.ArchAddress,
		big.NewInt(arch.FeePerByte),
		big.NewInt(arch.MinBounty),
		big.NewInt(arch.MinDiggingFee),
		big.NewInt(arch.MaxResurectionTime),
		big.NewInt(arch.FreeBond),
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error registering Archaeologist: %v Config values ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND have been reset to 0. You will need to reset this.", err)
	}

	arch.FreeBond = 0
	log.Printf("Register Archaeologist Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func (arch *Archaeologist) UpdateArchaeologist() {
	log.Println("***UPDATING ARCHAEOLOGIST***")
	tx, err := arch.SarcoSession.UpdateArchaeologist(
		arch.Endpoint,
		arch.ArchAddress,
		big.NewInt(arch.FeePerByte),
		big.NewInt(arch.MinBounty),
		big.NewInt(arch.MinDiggingFee),
		big.NewInt(arch.MaxResurectionTime),
		big.NewInt(arch.FreeBond),
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error updating Archaeologist: %v Config values ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND have been reset to 0. You will need to reset these.", err)
	}

	arch.FreeBond = 0
	log.Printf("Update Archaeologist Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func (arch *Archaeologist) ApproveFreeBondTransfer() {
	archSarcoBalance := arch.SarcoBalance()

	// Check if archSarcoBalance < freeBond
	if archSarcoBalance.Cmp(big.NewInt(arch.FreeBond)) == -1 {
		log.Fatalf("Your balance is too low to cover the free bond transfer. \n Balance Needed: %v \n Current Balance: %v", arch.FreeBond, archSarcoBalance)
	}

	tx, err := arch.TokenSession.Approve(
		arch.SarcoAddress,
		big.NewInt(arch.FreeBond),
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Approval Transaction for %v Sarco Tokens successful. Transaction ID: %v", arch.FreeBond, tx.Hash().Hex())
	log.Printf("Gas Used for Approval: %v", tx.Gas())
}
