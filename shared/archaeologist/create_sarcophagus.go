package archaeologist

import (
	"bytes"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

// handleCreateSarcophagus gets called when the createSarcophagus event is emitted on the sarcophagus contract
// it validates the public key, adds the new sarcophagus to state, and listens for the file to be sent from embalmer
func handleCreateSarcophagus(event *contracts.EventsCreateSarcophagus, arch *models.Archaeologist) {
	log.Println("Create Sarcophagus Event Triggered")
	log.Println("Event Name:", event.Name)
	log.Println("Asset Double Hash:", event.Identifier)
	log.Println("Archaeologist:", event.Archaeologist)
	log.Println("Embalmer:", event.Embalmer)
	log.Println("Resurrection Time:", event.ResurrectionTime)
	log.Println("Resurrection Window:", event.ResurrectionWindow)
	log.Println("Bounty:", event.Bounty)
	log.Println("Storage Fee:", event.StorageFee)
	log.Println("Digging Fee:", event.DiggingFee)
	log.Println("CursedBond:", event.CursedBond)
	log.Println("CurrentPublicKey:", event.ArchaeologistPublicKey)

	// The public key on the event must match the current public key on the archaeologist
	if !bytes.Equal(event.ArchaeologistPublicKey, arch.CurrentPublicKeyBytes) {
		log.Printf("Public Key on Sarcophagus does not match current Public Key : %v. Not listening for file.", arch.CurrentPublicKeyBytes)
		return
	}

	// check if sarcophagus already exists in state, in case of a replayed duplicate event
	if arch.IsArchSarcophagus(event.Identifier) {
		log.Printf("The sarcophagus for this double hash already exists: %v", event.Identifier)
		return
	}

	sarco, _ := arch.SarcoSession.Sarcophagus(event.Identifier)
	if sarco.State != 1 {
		log.Printf("Created Sarco state is not 1, not adding sarco to state")
		return
	}

	mutex.Lock()
	// add sarcophagus to state
	arch.Sarcophaguses[event.Identifier] = &models.Sarco{
		ResurrectionTime: event.ResurrectionTime,
		AccountIndex:     arch.AccountIndex,
		Updated:          false,
		UnwrapAttempts:   0,
	}

	// open the endpoint for the file to be received from the embalmer
	arch.FileHandlers[event.Identifier] = event.StorageFee
	mutex.Unlock()
}
