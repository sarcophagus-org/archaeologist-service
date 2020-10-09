package ethereum

import (
	"context"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
	"math/big"
)

func registerArchaeologist(session *contracts.SarcophagusSession, config *models.Config) {
	log.Println("***REGISTERING ARCHAEOLOGIST***")
	tx, err := session.RegisterArchaeologist(
		archPublicKeyBytes,
		config.ENDPOINT,
		archAddress,
		big.NewInt(config.FEE_PER_BYTE),
		big.NewInt(config.MIN_BOUNTY),
		big.NewInt(config.MIN_DIGGING_FEE),
		big.NewInt(config.MAX_RESURRECTION_TIME),
		big.NewInt(freeBond),
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error registering Archaeologist: %v Config values ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND have been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Register Archaeologist Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func updateArchaeologist(session *contracts.SarcophagusSession, config *models.Config) {
	log.Println("***UPDATING ARCHAEOLOGIST***")
	tx, err := session.UpdateArchaeologist(
		config.ENDPOINT,
		archAddress,
		big.NewInt(config.FEE_PER_BYTE),
		big.NewInt(config.MIN_BOUNTY),
		big.NewInt(config.MIN_DIGGING_FEE),
		big.NewInt(config.MAX_RESURRECTION_TIME),
		big.NewInt(freeBond),
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error updating Archaeologist: %v Config values ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND have been reset to 0. You will need to reset these.", err)
	}

	log.Printf("Update Archaeologist Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func handleFreeBondTransactions(session *contracts.TokenSession, archExists bool) {
	if freeBond > 0 {
		archSarcoBalance := ArchSarcoBalance()

		// Check if archSarcoBalance < freeBond
		if archSarcoBalance.Cmp(big.NewInt(freeBond)) == -1 {
			log.Fatalf("Your balance is too low to cover the free bond transfer. \n Balance Needed: %v \n Current Balance: %v", freeBond, archSarcoBalance)
		}

		tx, err := session.Approve(
			sarcoAddress,
			big.NewInt(freeBond),
		)

		if err != nil {
			log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config values ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
		}

		log.Printf("Approval Transaction for %v Sarco Tokens successful. Transaction ID: %v", freeBond, tx.Hash().Hex())
		log.Printf("Gas Used for Approval: %v", tx.Gas())
	} else if freeBond < 0 && archExists {

	}
}

func RegisterOrUpdateArchaeologist(config *models.Config) {
	sarcoSession := NewSarcophagusSession(context.Background())

	archaeologist, err := sarcoSession.Archaeologists(archAddress)
	if err != nil {
		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed: %v", err)
	}

	archExists := archaeologist.Exists
	tokenSession := NewSarcophagusTokenSession(context.Background())
	handleFreeBondTransactions(&tokenSession, archExists)

	if archExists {
		if freeBond > 0 ||
			archaeologist.Endpoint != config.ENDPOINT ||
			archaeologist.PaymentAddress != archAddress ||
			archaeologist.FeePerByte.Cmp(big.NewInt(config.FEE_PER_BYTE)) != 0 ||
			archaeologist.MinimumBounty.Cmp(big.NewInt(config.MIN_BOUNTY)) != 0 ||
			archaeologist.MinimumDiggingFee.Cmp(big.NewInt(config.MIN_DIGGING_FEE)) != 0 ||
			archaeologist.MaximumResurrectionTime.Cmp(big.NewInt(config.MAX_RESURRECTION_TIME)) != 0 {
			updateArchaeologist(&sarcoSession, config)
		} else {
			log.Printf("Archaeologist did not need to get updated, no config values have changed.")
		}
	} else {
		registerArchaeologist(&sarcoSession, config)
	}

	archaeologistUpdated, _ := sarcoSession.Archaeologists(archAddress)
	log.Printf("Current Free Bond: %v", archaeologistUpdated.FreeBond)
	log.Printf("Current Cursed Bond: %v", archaeologistUpdated.CursedBond)
}
