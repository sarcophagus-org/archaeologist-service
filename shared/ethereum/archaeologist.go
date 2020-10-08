package ethereum

import (
	"context"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func registerArchaeologist(session *contracts.SarcophagusSession, config *models.Config) {
	tx, err := session.RegisterArchaeologist(
		archPublicKeyBytes,
		config.FILE_PORT,
		archAddress,
		big.NewInt(config.FEE_PER_BYTE),
		big.NewInt(config.MIN_BOUNTY),
		big.NewInt(config.MIN_DIGGING_FEE),
		big.NewInt(config.MAX_RESURRECTION_TIME),
		big.NewInt(freeBond),
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error registering Archaeologist: %v", err)
	}

	log.Println("Register Archaeologist tx sent: %s", tx.Hash().Hex())
}

func handleFreeBondTransactions(session *contracts.TokenSession) {
	if freeBond > 0 {
		amountToApprove := big.NewInt(0).Add(session.TransactOpts.GasPrice, big.NewInt(freeBond))
		tx, err := session.Approve(
			sarcoAddress,
			amountToApprove,
		)

		if err != nil {
			log.Fatalf("Transaction reverted. Error Approving Transaction: %v", err)
		}

		log.Println("Allow transfer of to Free Bond tx sent:", amountToApprove, tx.Hash().Hex())
		log.Printf("Eth Balance: %v", EthBalance())
	}
}

func RegisterOrUpdateArchaeologist(config *models.Config) {
	archaeologist, err := sarcophagusContract.Archaeologists(&bind.CallOpts{}, archAddress)
	if err != nil {
		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed: %v", err)
	}

	tokenSession := NewSarcophagusTokenSession(context.Background())
	handleFreeBondTransactions(&tokenSession)

	sarcoSession := NewSarcophagusSession(context.Background())

	if archaeologist.Exists {
		// TODO: Update arch
		log.Println("Arch exists!")
	} else {
		registerArchaeologist(&sarcoSession, config)
	}
}
