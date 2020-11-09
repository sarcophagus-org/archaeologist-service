package archaeologist

import (
	"bytes"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
	"math/big"
)

func RegisterOrUpdateArchaeologist(arch *models.Archaeologist) {
	contractArch, err := arch.SarcoSession.Archaeologists(arch.ArchAddress)
	if err != nil {
		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed. Please check CONTRACT_ADDRESS is correct in the config file: %v", err)
	}

	if arch.FreeBond > 0 {
		arch.ApproveFreeBondTransfer()
	}

	if contractArch.Exists {
		if arch.FreeBond < 0 {
			withdrawAmount := new(big.Int).Abs(big.NewInt(arch.FreeBond))
			arch.WithdrawBond(withdrawAmount)
			arch.FreeBond = 0
		}

		if arch.FreeBond > 0 ||
			bytes.Compare(contractArch.CurrentPublicKey, arch.CurrentPublicKeyBytes) != 0 ||
			contractArch.Endpoint != arch.Endpoint ||
			contractArch.PaymentAddress != arch.ArchAddress ||
			contractArch.FeePerByte.Cmp(big.NewInt(arch.FeePerByte)) != 0 ||
			contractArch.MinimumBounty.Cmp(big.NewInt(arch.MinBounty)) != 0 ||
			contractArch.MinimumDiggingFee.Cmp(big.NewInt(arch.MinDiggingFee)) != 0 ||
			contractArch.MaximumResurrectionTime.Cmp(big.NewInt(arch.MaxResurectionTime)) != 0 {
			arch.UpdateArchaeologist()
		} else {
			log.Printf("Archaeologist did not need to get updated, no config values have changed.")
		}
	} else {
		arch.RegisterArchaeologist()
	}

	archaeologistUpdated, _ := arch.SarcoSession.Archaeologists(arch.ArchAddress)
	log.Printf("Current Free Bond: %v", archaeologistUpdated.FreeBond)
	log.Printf("Current Cursed Bond: %v", archaeologistUpdated.CursedBond)
}
