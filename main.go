package main

import (
	"flag"
	"fmt"
	"github.com/Dev43/arweave-go/api"
	"github.com/btcsuite/btcd/btcec"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/archaeologist"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"log"
	"strings"
)

func main(){
	log.Printf("%v", loadArt())

	// Add curve used by pub/priv keys in hdwallet to the accepted curves for ecies
	// This is the curve used by the Web App to encrypt the file bytes
	ecies.AddParamsForCurve(btcec.S256(), ecies.ECIES_AES128_SHA256)

	// Load the config file values. Flag with config location is optional
	var configFile = flag.String("config", "config", "Location of the config file.")
	configDir := "./"
	flag.Parse()

	config := new(models.Config)
	config.LoadConfig(*configFile, configDir, true)

	// Initialize the Archaeologist with values from the config
	arch := new(models.Archaeologist)
	errStrings := archaeologist.InitializeArchaeologist(arch, config)

	// List any errors. If any errors, exit
	if len(errStrings) > 0 {
		fmt.Println(fmt.Errorf(strings.Join(errStrings, "\n")))
		log.Fatal("**Please fix these errors in your config file and restart the service.**")
	}

	// Start the server to listen for files
	go arch.ListenForFile()

	// Register/Update the Archaeologist on the contract
	archaeologist.RegisterOrUpdateArchaeologist(arch)

	log.Printf("Eth Balance: %v", utility.ToDecimal(arch.EthBalance(), 18))
	log.Printf("Sarco Token Balance: %v", utility.ToDecimal(arch.SarcoBalance(), 18))
	log.Println("Arweave Balance:", utility.ToDecimal(ar.ArweaveBalance(arch.ArweaveTransactor.Client.(*api.Client), arch.ArweaveWallet), 12))
	log.Printf("Arweave Address: %v", arch.ArweaveWallet.Address())

	// Listen for contract events
	archaeologist.EventsSubscribe(arch)
}

func loadArt() string{
	return "\n\n ██   █▄▄▄▄ ▄█▄     ▄  █ ██   ▄███▄   ████▄ █    ████▄   ▄▀  ▄█    ▄▄▄▄▄      ▄▄▄▄▀ \n█ █  █  ▄▀ █▀ ▀▄  █   █ █ █  █▀   ▀  █   █ █    █   █ ▄▀    ██   █     ▀▄ ▀▀▀ █    \n█▄▄█ █▀▀▌  █   ▀  ██▀▀█ █▄▄█ ██▄▄    █   █ █    █   █ █ ▀▄  ██ ▄  ▀▀▀▀▄       █    \n█  █ █  █  █▄  ▄▀ █   █ █  █ █▄   ▄▀ ▀████ ███▄ ▀████ █   █ ▐█  ▀▄▄▄▄▀       █     \n   █   █   ▀███▀     █     █ ▀███▀             ▀       ███   ▐              ▀      \n  █   ▀             ▀     █                                                        \n ▀                       ▀                                                         \n\n"
}
