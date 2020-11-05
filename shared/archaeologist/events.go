package archaeologist

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func handleUpdateSarcophagus(event *contracts.EventsUpdateSarcophagus, arch *models.Archaeologist) {
	log.Println("Update Sarcophagus Event Sent:", event.AssetId)
	/* TODO: Clean up sarcophagus in file handlers */
	/* Add resurrection listener */
}

func handleCreateSarcophagus(event *contracts.EventsCreateSarcophagus, arch *models.Archaeologist) {
	log.Println("Name:", event.Name)
	log.Println("Asset Double Hash:", event.AssetDoubleHash)
	log.Println("Archaeologist:", event.Archaeologist)
	log.Println("Embalmer:", event.Embalmer)
	log.Println("Resurrection Time:", event.ResurrectionTime)
	log.Println("Resurrection Window:", event.ResurrectionWindow)
	log.Println("Bounty:", event.Bounty)
	log.Println("Storage Fee:", event.StorageFee)
	log.Println("Digging Fee:", event.DiggingFee)
	log.Println("CursedBond:", event.CursedBond)

	/* TODO: Update to handle multiple files (when 'create sarcophagus' is called multiple times) */
	/* Consider pushing file handlers to slice */
	/* Make server separate from file handler */

	fileHandler := &models.FileHandler{
		event.AssetDoubleHash,
		event.Embalmer,
		arch.PrivateKey,
		event.StorageFee,
		big.NewInt(arch.FeePerByte),
		arch.FilePort,
		arch.ArweaveTransactor,
		arch.ArweaveWallet,
	}

	/* Todo -- detect if we are already listening */
	fileHandler.HandleFileUpload()
}

func EventsSubscribe(arch *models.Archaeologist) {
	sarcoEvents, err := contracts.NewEvents(arch.SarcoAddress, arch.Client)
	if err != nil {
		log.Printf("Error creating events contract")
	}

	/* Create Sarcophagus Subscription */
	csSink := make(chan *contracts.EventsCreateSarcophagus)
	csSub, err := sarcoEvents.WatchCreateSarcophagus(&bind.WatchOpts{}, csSink)
	if err != nil {
		log.Fatalf("Error subscribing to CreateSarcophagus event: %v", err)
	}

	/* Update Sarcophagus Subscription */
	usSink := make(chan *contracts.EventsUpdateSarcophagus)
	usSub, err := sarcoEvents.WatchUpdateSarcophagus(&bind.WatchOpts{}, usSink)
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
