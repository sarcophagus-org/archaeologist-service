package archaeologist

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
	"math/big"
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