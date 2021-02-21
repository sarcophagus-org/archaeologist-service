// Not used by the service, for testing purposes only
// Used in conjunction with the embalmer package to test embalmer actions

package main

import (
	"flag"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"log"
	"math/big"
	"time"
)

func main(){
	config := new(embalmer.EmbalmerConfig)
	config.LoadEmbalmerConfig("embalmer_config", "../embalmer")
	emb := new(embalmer.Embalmer)
	defaultResTime := int64(120)
	typeFlag := flag.String("type", "create", "Create or Update a Sarcophagus")
	resurrectionFlag := flag.Int64("res",  defaultResTime, "Resurrection Time")
	seedFlag := flag.Int64("seed", 1, "Seed to generate random bytes")

	embalmer.InitEmbalmer(emb, config, *resurrectionFlag)

	flag.Parse()
	fileBytes, assetDoubleHashBytes := embalmer.DoubleHashBytesFromSeed(*seedFlag, 20)

	log.Printf("Asset double hash bytes: %v", assetDoubleHashBytes)

	if *typeFlag == "create" {
		emb.CreateSarcophagus(config.RECIPIENT_PRIVATE_KEY, assetDoubleHashBytes, "Test Sarco")
		log.Println("Embalmer Sarco Balance:", emb.EmbalmerSarcoBalance())
	}

	if *typeFlag == "rewrap" {
		emb.RewrapSarcophagus(assetDoubleHashBytes, big.NewInt(time.Now().Unix() + defaultResTime * 2))
	}

	if *typeFlag == "clean" {
		emb.CleanupSarcophagus(assetDoubleHashBytes)
	}

	if *typeFlag == "bury" {
		emb.BurySarcophagus(assetDoubleHashBytes)
	}

	if *typeFlag == "cancel" {
		emb.CancelSarcophagus(assetDoubleHashBytes)
	}

	if *typeFlag == "update" {
		emb.UpdateSarcophagus(assetDoubleHashBytes, fileBytes)
	}
}
