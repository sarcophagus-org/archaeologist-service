package archaeologist

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

func handleUpdateSarcophagus(event *contracts.EventsUpdateSarcophagus, arch *models.Archaeologist) {
	log.Println("Update Sarcophagus Event Sent:", event.AssetId)
	/* TODO: Clean up sarcophagus in file handlers */

	/* TODO: See if we can return resurrection time in update sarc event and use below */
	/* scheduleUnwrap should return timer that we store pointer to in state. call stop on this timer if/when rewrap is called */

	resurrectionTime := int64(0)
	scheduleUnwrap(resurrectionTime, arch)
}