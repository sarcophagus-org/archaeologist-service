package archaeologist

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

	freeBond = 0
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

	freeBond = 0
	log.Printf("Update Archaeologist Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}

func approveFreeBondTransfer(session *contracts.TokenSession) {
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
		log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Approval Transaction for %v Sarco Tokens successful. Transaction ID: %v", freeBond, tx.Hash().Hex())
	log.Printf("Gas Used for Approval: %v", tx.Gas())
}

func RegisterOrUpdateArchaeologist(arch *models.Archaeologist) {
	contractArch, err := arch.SarcoSession.Archaeologists(arch.ArchAddress)
	if err != nil {
		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed: %v", err)
	}

	if arch.FreeBond > 0 {
		approveFreeBondTransfer(&arch.TokenSession)
	}

	if contractArch.Exists {
		if arch.FreeBond < 0 {
			withdrawAmount := new(big.Int).Abs(big.NewInt(arch.FreeBond))
			arch.WithdrawBond(withdrawAmount)
			arch.FreeBond = 0
		}

		if arch.FreeBond > 0 ||
			contractArch.Endpoint != arch.EndPoint ||
			contractArch.PaymentAddress != arch.ArchAddress ||
			contractArch.FeePerByte.Cmp(big.NewInt(arch.FeePerByte)) != 0 ||
			contractArch.MinimumBounty.Cmp(big.NewInt(arch.MinBounty)) != 0 ||
			contractArch.MinimumDiggingFee.Cmp(big.NewInt(arch.MinDiggingFee)) != 0 ||
			contractArch.MaximumResurrectionTime.Cmp(big.NewInt(arch.MaxResurectionTime)) != 0 {
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
