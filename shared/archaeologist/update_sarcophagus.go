package archaeologist

import (
	"github.com/Dev43/arweave-go/api"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

// handleUpdateSarcophagus - updates a created sarcophagus
// called when an "updateSarcophagus" event is emitted from the contract
func handleUpdateSarcophagus(event *contracts.EventsUpdateSarcophagus, arch *models.Archaeologist) {
	log.Println("Update Sarcophagus Event Sent for asset ID:", event.AssetId)

	// Delete open file handler for the double hash
	delete(arch.FileHandlers, event.Identifier)

	if sarcophagus, ok := arch.Sarcophaguses[event.Identifier]; ok {
		// Only schedule unwrap if sarcophagus has not been updated yet (in case of replayed events)
		if !sarcophagus.Updated {
			privateKey := hdw.PrivateKeyFromIndex(arch.Wallet, arch.AccountIndex)
			resurrectionTime := sarcophagus.ResurrectionTime

			arweaveClient := arch.ArweaveTransactor.Client.(*api.Client)
			log.Printf("Scheduling Unwrap for: %v", resurrectionTime)
			scheduleUnwrap(&arch.SarcoSession, arweaveClient, resurrectionTime, arch, event.Identifier, privateKey, event.AssetId)

			// key pair has been used for this sarcophagus, increment the account index and update the current public key
			arch.AccountIndex += 1
			arch.CurrentPublicKeyBytes = hdw.PublicKeyBytesFromIndex(arch.Wallet, arch.AccountIndex)
			sarcophagus.Updated = true
		}
	} else {
		log.Printf("We dont have a sarcophagus to update for the double hash: %v",  event.Identifier)
	}
}