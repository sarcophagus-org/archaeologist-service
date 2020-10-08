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
		log.Fatalf("Transaction reverted. Error registering Archaeologist: %v Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Register Archaeologist tx sent: %s", tx.Hash().Hex())
}

func handleFreeBondTransactions(session *contracts.TokenSession, archExists bool) {
	if freeBond > 0 {
		balanceNeededToApprove := big.NewInt(0).Add(session.TransactOpts.GasPrice, big.NewInt(freeBond))
		if balanceNeededToApprove.Cmp(ArchBalance()) == -1 {
			log.Fatalf("Your balance is too low to cover the free bond transfer. \n Balance Needed: %v \n Current Balance: %v", balanceNeededToApprove, ArchBalance)
		}

		tx, err := session.Approve(
			sarcoAddress,
			big.NewInt(freeBond),
		)

		if err != nil {
			log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
		}

		log.Printf("Approval Transaction successful. Approved for amount: %v \n Transaction ID: %v", freeBond, tx.Hash().Hex())
		log.Printf("Gas Estimate: %v", session.TransactOpts.GasPrice)
		log.Printf("Gas Used: %v", tx.GasPrice())
	} else if freeBond < 0 && archExists {
		
	}
}

func RegisterOrUpdateArchaeologist(config *models.Config) {
	archaeologist, err := sarcophagusContract.Archaeologists(&bind.CallOpts{}, archAddress)
	if err != nil {
		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed: %v", err)
	}

	archExists := archaeologist.Exists

	tokenSession := NewSarcophagusTokenSession(context.Background())
	handleFreeBondTransactions(&tokenSession, archExists)

	sarcoSession := NewSarcophagusSession(context.Background())

	if archExists {
		// TODO: Update arch
		log.Println("Arch exists!")
	} else {
		registerArchaeologist(&sarcoSession, config)
	}
}
