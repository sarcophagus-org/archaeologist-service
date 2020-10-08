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
		log.Fatalf("Transaction reverted. Error registering Archaeologist: %v Config values ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND have been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Register Archaeologist tx sent: %s", tx.Hash().Hex())
}

func updateArchaeologist(session *contracts.SarcophagusSession, config *models.Config) {
	log.Println("UPDATING ARCHAEOLOGIST")

	tx, err := session.UpdateArchaeologist(
		config.FILE_PORT,
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

	log.Printf("Register Archaeologist tx sent: %s", tx.Hash().Hex())
}

func handleFreeBondTransactions(session *contracts.TokenSession, archExists bool) {
	if freeBond > 0 {
		gasPrice := session.TransactOpts.GasPrice
		gasLimit := session.TransactOpts.GasLimit

		// TODO: Consider how to get a true estimate for balanceNeededForApproval
		// We need the gas estimate for the approval transaction *and* the transaction the approval is for
		// The gas costs may be different for each of these.
		// EstimateGasLimit() needs to be implemented

		freeBondBigInt := big.NewInt(freeBond)
		balanceNeededForApproval := BalanceNeededForApproval(gasPrice, gasLimit, freeBondBigInt)
		currentBalance := ArchBalance()

		// Check if currentBalance < balanceNeededToApprove
		if currentBalance.Cmp(balanceNeededForApproval) == -1 {
			log.Fatalf("Your balance is too low to cover the free bond transfer. \n Balance Needed: %v \n Current Balance: %v", balanceNeededForApproval, currentBalance)
		}

		tx, err := session.Approve(
			sarcoAddress,
			big.NewInt(freeBond),
		)

		if err != nil {
			log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
		}

		log.Printf("Approval Transaction successful. \n Approved for amount: %v \n Transaction ID: %v", freeBond, tx.Hash().Hex())
		log.Printf("Gas Price: %v", tx.GasPrice())
		log.Printf("Gas Used: %v", tx.Gas())
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
		if freeBond > 0 ||
			archaeologist.PaymentAddress != config.PAYMENT_ADDRESS ||
			archaeologist.FeePerByte != config.FEE_PER_BYTE ||
			archaeologist.MinimumBounty != config.MIN_BOUNTY ||
			archaeologist.MinimumDiggingFee != config.MIN_DIGGING_FEE ||
			archaeologist.MaximumResurrectionTime != config.MAX_RESURRECTION_TIME {
			updateArchaeologist(&sarcoSession, config)
		}
	} else {
		registerArchaeologist(&sarcoSession, config)
	}
}
