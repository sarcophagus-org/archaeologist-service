package archaeologist

import (
	"github.com/Dev43/arweave-go/api"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

// handleRewrapSarcophagus schedules a new unwrap for the sarcopgagus at the specified resurrection time
func handleRewrapSarcophagus(event *contracts.EventsRewrapSarcophagus, arch *models.Archaeologist) {
	log.Println("Rewrap Sarcophagus Event Sent:", event.Identifier)

	if sarcophagus, ok := arch.Sarcophaguses[event.Identifier]; ok {
		if sarcophagus.ResurrectionTime.Cmp(event.ResurrectionTime) != 0 {
			// grab the private key for this account index
			privateKey := hdw.PrivateKeyFromIndex(arch.Wallet, sarcophagus.AccountIndex)

			// Update resurrection time for Sarcophagus in state
			sarcophagus.ResurrectionTime = event.ResurrectionTime
			scheduleUnwrap(&arch.SarcoSession, arch.ArweaveTransactor.Client.(*api.Client), event.ResurrectionTime, arch, event.Identifier, privateKey, event.AssetId)
		} else {
			log.Printf("Unwrapping already scheduled for: %v, skipping rewrap",  event.Identifier)
		}
	} else {
		log.Printf("We dont have a sarcophagus to update for the rewrapping: %v",  event.Identifier)
	}
}