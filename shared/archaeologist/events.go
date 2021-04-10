// Responsible for subscribing to all relevant events that get emitted from the contract
// and acting on those events

package archaeologist

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"log"
)

// watchCreateSarcophagus .
func watchCreateSarcophagus(sarcoEvents *contracts.Events, archAddress []common.Address) (chan *contracts.EventsCreateSarcophagus, event.Subscription) {
	sink := make(chan *contracts.EventsCreateSarcophagus)
	sub, err := sarcoEvents.WatchCreateSarcophagus(&bind.WatchOpts{}, sink, nil, archAddress)
	if err != nil {
		log.Fatalf("Error subscribing to CreateSarcophagus event: %v", err)
	}

	return sink, sub
}

// watchUpdateSarcophagus .
func watchUpdateSarcophagus(sarcoEvents *contracts.Events) (chan *contracts.EventsUpdateSarcophagus, event.Subscription) {
	sink := make(chan *contracts.EventsUpdateSarcophagus)
	sub, err := sarcoEvents.WatchUpdateSarcophagus(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		log.Fatalf("Error subscribing to UpdateSarcophagus event: %v", err)
	}

	return sink, sub
}

// watchRewrapSarcophagus .
func watchRewrapSarcophagus(sarcoEvents *contracts.Events) (chan *contracts.EventsRewrapSarcophagus, event.Subscription) {
	sink := make(chan *contracts.EventsRewrapSarcophagus)
	sub, err := sarcoEvents.WatchRewrapSarcophagus(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		log.Fatalf("Error subscribing to RewrapSarcophagus event: %v", err)
	}

	return sink, sub
}

// watchCleanUpSarcophagus .
func watchCleanUpSarcophagus(sarcoEvents *contracts.Events) (chan *contracts.EventsCleanUpSarcophagus, event.Subscription) {
	sink := make(chan *contracts.EventsCleanUpSarcophagus)
	sub, err := sarcoEvents.WatchCleanUpSarcophagus(&bind.WatchOpts{}, sink, nil, nil)
	if err != nil {
		log.Fatalf("Error subscribing to CleanUp Sarcophagus event: %v", err)
	}

	return sink, sub
}

// watchBurySarcophagus .
func watchBurySarcophagus(sarcoEvents *contracts.Events) (chan *contracts.EventsBurySarcophagus, event.Subscription) {
	sink := make(chan *contracts.EventsBurySarcophagus)
	sub, err := sarcoEvents.WatchBurySarcophagus(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		log.Fatalf("Error subscribing to Bury Sarcophagus event: %v", err)
	}

	return sink, sub
}

// watchCancelSarcophagus .
func watchCancelSarcophagus(sarcoEvents *contracts.Events) (chan *contracts.EventsCancelSarcophagus, event.Subscription) {
	sink := make(chan *contracts.EventsCancelSarcophagus)
	sub, err := sarcoEvents.WatchCancelSarcophagus(&bind.WatchOpts{}, sink, nil)
	if err != nil {
		log.Fatalf("Error subscribing to Cancel Sarcophagus event: %v", err)
	}

	return sink, sub
}

// watchAccuseArchaeologist .
func watchAccuseArchaeologist(sarcoEvents *contracts.Events) (chan *contracts.EventsAccuseArchaeologist, event.Subscription) {
	sink := make(chan *contracts.EventsAccuseArchaeologist)
	sub, err := sarcoEvents.WatchAccuseArchaeologist(&bind.WatchOpts{}, sink, nil, nil)
	if err != nil {
		log.Fatalf("Error subscribing to Accuse Archaeologist event: %v", err)
	}

	return sink, sub
}

// EventsSubscribe subscribes to the events and handles them
func EventsSubscribe(arch *models.Archaeologist) {
	sarcoEvents, err := contracts.NewEvents(arch.SarcoAddress, arch.Client)
	if err != nil {
		log.Printf("Error instantiating Events in EventsSubscribe: %v", err)
	}

	/* Add arch address to slice for filtering purposes */
	var archAddress = []common.Address{arch.ArchAddress}

	createSink, createSub := watchCreateSarcophagus(sarcoEvents, archAddress)
	updateSink, updateSub := watchUpdateSarcophagus(sarcoEvents)
	rewrapSink, rewrapSub := watchRewrapSarcophagus(sarcoEvents)
	cleanSink, cleanSub := watchCleanUpSarcophagus(sarcoEvents)
	burySink, burySub := watchBurySarcophagus(sarcoEvents)
	cancelSink, cancelSub := watchCancelSarcophagus(sarcoEvents)
	accuseSink, accuseSub := watchAccuseArchaeologist(sarcoEvents)

	log.Println("Listening For Events...")

	for {
		select {
		case err := <-createSub.Err():
			if err != nil {
				log.Fatalln("Error with Create Sarcophagus Subscription:", err)
			}
		case err := <-updateSub.Err():
			if err != nil {
				log.Fatalln("Error with Update Sarcophagus Subscription:", err)
			}
		case err := <-rewrapSub.Err():
			if err != nil {
				log.Fatalln("Error with Rewrap Sarcophagus Subscription:", err)
			}
		case err := <-cleanSub.Err():
			if err != nil {
				log.Fatalln("Error with Clean Sarcophagus Subscription:", err)
			}
		case err := <-burySub.Err():
			if err != nil {
				log.Fatalln("Error with Bury Sarcophagus Subscription:", err)
			}
		case err := <-cancelSub.Err():
			if err != nil {
				log.Fatalln("Error with Cancel Sarcophagus Subscription:", err)
			}
		case err := <-accuseSub.Err():
			if err != nil {
				log.Fatalln("Error with Accuse Archaeologist Subscription:", err)
			}
		case sarcoEvent := <-createSink:
			// the createSink only returns sarcophagi related to this archaeologist
			// updateSink and rewrapSink do not, so those have manual filtering before handling the event
			go handleCreateSarcophagus(sarcoEvent, arch)
		case sarcoEvent := <-updateSink:
			if arch.IsArchSarcophagus(sarcoEvent.Identifier) {
				go handleUpdateSarcophagus(sarcoEvent, arch)
			}
		case sarcoEvent := <-rewrapSink:
			if arch.IsArchSarcophagus(sarcoEvent.Identifier) {
				go handleRewrapSarcophagus(sarcoEvent, arch)
			}
		// these events indicate a sarcophagus is 'done'
		// and can be removed from state
		case sarcoEvent := <-cleanSink:
			arch.RemoveArchSarcophagus(sarcoEvent.Identifier)
		case sarcoEvent := <-burySink:
			arch.RemoveArchSarcophagus(sarcoEvent.Identifier)
		case sarcoEvent := <-cancelSink:
			arch.RemoveArchSarcophagus(sarcoEvent.Identifier)
		case sarcoEvent := <-accuseSink:
			arch.RemoveArchSarcophagus(sarcoEvent.Identifier)
		case <-arch.ReconnectChan:
			log.Printf("reconnecting...")
			return
		}
	}
}
