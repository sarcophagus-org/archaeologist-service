package archaeologist

import (
	"github.com/Dev43/arweave-go/api"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

func handleUpdateSarcophagus(event *contracts.EventsUpdateSarcophagus, arch *models.Archaeologist) {
	log.Println("Update Sarcophagus Event Sent for asset ID:", event.AssetId)

	// Delete open file handler for the double hash
	delete(arch.FileHandlers, event.Identifier)

	if resurrectionTime, ok := arch.Sarcophaguses[event.Identifier]; ok {
		privateKey := hdw.PrivateKeyFromIndex(arch.Wallet, arch.AccountIndex)

		// Save the account index in state mapping to be used by rewrapping
		arch.SarcophagusesAccountIndex[event.Identifier] = arch.AccountIndex

		arweaveClient := arch.ArweaveTransactor.Client.(*api.Client)
		log.Printf("Scheduling Unwrap for: %v", resurrectionTime)
		scheduleUnwrap(&arch.SarcoSession, arweaveClient, resurrectionTime, arch, event.Identifier, privateKey, event.AssetId)

		arch.AccountIndex += 1
		arch.CurrentPublicKeyBytes = hdw.PublicKeyBytesFromIndex(arch.Wallet, arch.AccountIndex)
	} else {
		log.Printf("We dont have a sarcophagus to update for the double hash: %v",  event.Identifier)
	}
}