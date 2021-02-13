package archaeologist

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

func handleRewrapSarcophagus(event *contracts.EventsRewrapSarcophagus, arch *models.Archaeologist) {
	log.Println("Rewrap Sarcophagus Event Sent:", event.Identifier)

	/* Update resurrection time for Sarcophagus in state */
	if _, ok := arch.Sarcophaguses[event.Identifier]; ok {
		arch.Sarcophaguses[event.Identifier] = event.ResurrectionTime
	} else {
		log.Printf("We dont have a sarcophagus to update for the rewrapping: %v",  event.Identifier)
	}
}