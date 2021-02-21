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

	if _, ok := arch.Sarcophaguses[event.Identifier]; ok {
		// Lookup the account index for this sarcophagus
		if sarcoAccountIndex, ok := arch.SarcophagusesAccountIndex[event.Identifier]; ok {

			// grab the private key for this account index
			privateKey := hdw.PrivateKeyFromIndex(arch.Wallet, sarcoAccountIndex)

			// Update resurrection time for Sarcophagus in state
			// If there are any other unwrapping scheduled for this sarcophagus,
			// those scheduleUnwrap timers will still run, but they will see the resurrection time
			// has changed, and silently return without creating an unwrap tx
			arch.Sarcophaguses[event.Identifier] = event.ResurrectionTime
			scheduleUnwrap(&arch.SarcoSession, arch.ArweaveTransactor.Client.(*api.Client), event.ResurrectionTime, arch, event.Identifier, privateKey, event.AssetId)
		} else {
			log.Printf("We dont have an account index for this sarcophagus to use for the rewrapping: %v",  event.Identifier)
		}
	} else {
		log.Printf("We dont have a sarcophagus to update for the rewrapping: %v",  event.Identifier)
	}
}