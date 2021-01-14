package main

import (
	"flag"
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
	var configFile = flag.String("config", "config", "Location of the config file.")
	configDir := "./"
	flag.Parse()

	config := new(models.Config)
	config.LoadConfig(*configFile, configDir, true)
	arch := new(models.Archaeologist)
	errStrings := archaeologist.InitializeArchaeologist(arch, config, configDir)
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
	log.Println("Arweave Balance:", utility.ToDecimal(arweave.ArweaveBalance(arch.ArweaveTransactor.Client.(*api.Client), arch.ArweaveWallet), 12))
	log.Printf("Arweave Address: %v", arch.ArweaveWallet.Address())
	archaeologist.EventsSubscribe(arch)
}
