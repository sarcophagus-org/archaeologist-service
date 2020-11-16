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

	/* Delete open file handler for the double hash */
	delete(arch.FileHandlers, event.AssetDoubleHash)

	/* Shutdown server if we dont have any open file handlers */
	/* Should we shut this down regardless? */
	if len(arch.FileHandlers) == 0 {
		arch.ShutdownServer()
	}

	if updatedSarc, ok := arch.Sarcophaguses[event.AssetDoubleHash]; ok {
		resurrectionTime := updatedSarc.ResurrectionTime
		privateKey := hdw.PrivateKeyFromIndex(arch.Wallet, arch.AccountIndex)
		arweaveClient := arch.ArweaveTransactor.Client.(*api.Client)
		scheduleUnwrap(&arch.SarcoSession, arweaveClient, resurrectionTime, event.AssetDoubleHash, privateKey, event.AssetId)

		arch.AccountIndex += 1
		arch.CurrentPublicKeyBytes = hdw.PublicKeyBytesFromIndex(arch.Wallet, arch.AccountIndex)
	} else {
		log.Printf("We dont have a sarcophagus to update for the double hash: %v",  event.AssetDoubleHash)
	}
}