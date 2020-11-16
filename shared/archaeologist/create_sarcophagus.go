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

	if bytes.Compare(event.ArchaeologistPublicKey, arch.CurrentPublicKeyBytes) != 0 {
		log.Print("Public Key on Sarcophagus does not match current Public Key. Not listening for file.")
		return
	}

	_, ok := arch.Sarcophaguses[event.AssetDoubleHash]
	if ok {
		/* The sarco already exists! Oh no! */

	}

	arch.Sarcophaguses[event.AssetDoubleHash] = models.Sarcophagus{ResurrectionTime: event.ResurrectionTime}
	arch.FileHandlers[event.AssetDoubleHash] = event.StorageFee

	arch.InitServer(arch.FilePort)
	arch.StartServer()
}