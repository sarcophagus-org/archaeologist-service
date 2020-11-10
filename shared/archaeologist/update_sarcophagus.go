package archaeologist

import (
	"github.com/Dev43/arweave-go/api"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

func handleUpdateSarcophagus(event *contracts.EventsUpdateSarcophagus, arch *models.Archaeologist) {
	log.Println("Update Sarcophagus Event Sent:", event.AssetId)
	/* TODO: Clean up sarcophagus in file handlers */

	if updatedSarc, ok := arch.Sarcophaguses[event.AssetDoubleHash]; ok {
		resurrectionTime := updatedSarc.ResurrectionTime
		privateKey := hdw.PrivateKeyFromIndex(arch.Wallet, updatedSarc.AccountIndex)
		log.Printf("account Index:", updatedSarc.AccountIndex)
		arweaveClient := arch.ArweaveTransactor.Client.(*api.Client)

		scheduleUnwrap(&arch.SarcoSession, arweaveClient, resurrectionTime, event.AssetDoubleHash, privateKey, event.AssetId)
	} else {
		log.Printf("We dont have a sarcophagus to update for the double hash: %v",  event.AssetDoubleHash)
	}
}