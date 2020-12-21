package archaeologist

import (
	"bytes"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
)

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
	log.Println("CurrentPublicKey:", event.ArchaeologistPublicKey)

	if !bytes.Equal(event.ArchaeologistPublicKey, arch.CurrentPublicKeyBytes) {
		log.Printf("Public Key on Sarcophagus does not match current Public Key : %v. Not listening for file.", arch.CurrentPublicKeyBytes)
		return
	}

	if arch.IsArchSarcophagus(event.AssetDoubleHash) {
		/* The sarco already exists in state but should not */
		log.Printf("The sarcophagus for this double hash already exists: %v", event.AssetDoubleHash)
		return
	}

	arch.Sarcophaguses[event.AssetDoubleHash] = event.ResurrectionTime
	arch.FileHandlers[event.AssetDoubleHash] = event.StorageFee
	go arch.ListenForFile()
}