package main

import (
	"fmt"
	"github.com/Dev43/arweave-go/api"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/archaeologist"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"log"
	"strings"
)

func main(){
	config := new(models.Config)
	config.LoadConfig("config", "config", true)
	arch := new(models.Archaeologist)
	errStrings := archaeologist.InitializeArchaeologist(arch, config, "config")
	if len(errStrings) > 0 {
		fmt.Println(fmt.Errorf(strings.Join(errStrings, "\n")))
		log.Fatal("**Please fix these errors in your config file and restart the service.**")
	}

	if len(arch.FileHandlers) > 0 {
		go arch.ListenForFile()
	}

	archaeologist.RegisterOrUpdateArchaeologist(arch)

	log.Printf("Eth Balance: %v", utility.ToDecimal(arch.EthBalance(), 18))
	log.Printf("Sarco Token Balance: %v", utility.ToDecimal(arch.SarcoBalance(), 18))
	log.Println("Arweave Balance:", utility.ToDecimal(arweave.ArweaveBalance(arch.ArweaveTransactor.Client.(*api.Client), arch.ArweaveWallet), 18))
	archaeologist.EventsSubscribe(arch)
}
