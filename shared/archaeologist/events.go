package archaeologist

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func EventsSubscribe(arch *models.Archaeologist) {
	sarcoEvents, err := contracts.NewEvents(arch.SarcoAddress, arch.Client)
	if err != nil {
		log.Printf("Error creating events contract")
	}

	/* Add arch address to slice for filtering purposes */
	var archAddress = []common.Address{arch.ArchAddress}

	/* Create Sarcophagus Subscription */
	csSink := make(chan *contracts.EventsCreateSarcophagus)
	csSub, err := sarcoEvents.WatchCreateSarcophagus(&bind.WatchOpts{}, csSink, nil, archAddress)
	if err != nil {
		log.Fatalf("Error subscribing to CreateSarcophagus event: %v", err)
	}

	/* Update Sarcophagus Subscription */
	/* TODO: Pass in slice of created sarcophaguses as query param to only trigger event for ones we care about */
	usSink := make(chan *contracts.EventsUpdateSarcophagus)
	usSub, err := sarcoEvents.WatchUpdateSarcophagus(&bind.WatchOpts{}, usSink, nil)
	if err != nil {
		log.Fatalf("Error subscribing to UpdateSarcophagus event: %v", err)
	}

	log.Println("Listening For Events...")

	for {
		select {
		case err := <-csSub.Err():
			if err != nil {
				log.Println("Error with Create Sarcophagus Subscription:", err)
			}
		case err := <-usSub.Err():
			if err != nil {
				log.Println("Error with Update Sarcophagus Subscription:", err)
			}
		case event := <-csSink:
			go handleCreateSarcophagus(event, arch)
		case event := <-usSink:
			handleUpdateSarcophagus(event, arch)
		}
	}
}
