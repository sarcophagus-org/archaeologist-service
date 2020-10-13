package ethereum

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func EventsSubscribe() {
	wsClient, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatalf("error with web socket address: %v", err)
	}

	sink := make(chan *contracts.EventsCreateSarcophagus)
	sarcoEvents, err := contracts.NewEvents(sarcoAddress, wsClient)
	sub, err := sarcoEvents.WatchCreateSarcophagus(&bind.WatchOpts{}, sink)

	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				log.Println(err)
			}
		case event := <-sink:
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
		}
	}
}
