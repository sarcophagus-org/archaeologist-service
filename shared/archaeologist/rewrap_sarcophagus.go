package archaeologist

import (
	"github.com/Dev43/arweave-go/api"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

func handleRewrapSarcophagus(event *contracts.EventsRewrapSarcophagus, arch *models.Archaeologist) {
	log.Println("Rewrap Sarcophagus Event Sent:", event.Identifier)

	/* Update resurrection time for Sarcophagus in state */
	if _, ok := arch.Sarcophaguses[event.Identifier]; ok {
		if sarcoAccountIndex, ok := arch.SarcophagusesAccountIndex[event.Identifier]; ok {
			privateKey := hdw.PrivateKeyFromIndex(arch.Wallet, sarcoAccountIndex)
			arch.Sarcophaguses[event.Identifier] = event.ResurrectionTime
			scheduleUnwrap(&arch.SarcoSession, arch.ArweaveTransactor.Client.(*api.Client), event.ResurrectionTime, arch, event.Identifier, privateKey, event.AssetId)
		} else {
			log.Printf("We dont have an account index for this sarcophagus to use for the rewrapping: %v",  event.Identifier)
		}
	} else {
		log.Printf("We dont have a sarcophagus to update for the rewrapping: %v",  event.Identifier)
	}
}