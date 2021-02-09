package archaeologist

import (
	"bytes"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"log"
	"math/big"
)

func RegisterOrUpdateArchaeologist(arch *models.Archaeologist) {
	contractArch, err := arch.SarcoSession.Archaeologists(arch.ArchAddress)
	if err != nil {
		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed. Please check CONTRACT_ADDRESS is correct in the config file: %v", err)
	}

	if arch.FreeBond.Cmp(big.NewInt(0)) == 1 {
		arch.ApproveFreeBondTransfer()
	}

	if contractArch.Exists {
		if arch.FreeBond.Cmp(big.NewInt(0)) == -1 {
			withdrawAmount := new(big.Int).Abs(arch.FreeBond)
			arch.WithdrawBond(withdrawAmount)
			arch.FreeBond = big.NewInt(0)
		}

		if arch.FreeBond.Cmp(big.NewInt(0)) == 1 ||
			!bytes.Equal(contractArch.CurrentPublicKey, arch.CurrentPublicKeyBytes) ||
			contractArch.Endpoint != arch.Endpoint ||
			contractArch.PaymentAddress != arch.PaymentAddress ||
			contractArch.FeePerByte.Cmp(arch.FeePerByte) != 0 ||
			contractArch.MinimumBounty.Cmp(arch.MinBounty) != 0 ||
			contractArch.MinimumDiggingFee.Cmp(arch.MinDiggingFee) != 0 ||
			contractArch.MaximumResurrectionTime.Cmp(arch.MaxResurectionTime) != 0 {
			arch.UpdateArchaeologist()
		} else {
			log.Printf("Archaeologist did not need to get updated, no config values have changed.")
		}
	} else {
		arch.RegisterArchaeologist()
	}

	archaeologistUpdated, _ := arch.SarcoSession.Archaeologists(arch.ArchAddress)
	log.Printf("Current Free Bond: %v", utility.ToDecimal(archaeologistUpdated.FreeBond, 18))

	if archaeologistUpdated.FreeBond.Cmp(big.NewInt(0)) == 0 {
		log.Printf("CURRENT FREE BOND IS 0. YOU WILL BE UNABLE TO ACCEPT NEW JOBS.")
	}

	log.Printf("Current Cursed Bond: %v", utility.ToDecimal(archaeologistUpdated.CursedBond, 18))
}
