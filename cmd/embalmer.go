package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/embalmer"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"
)

func createTmpFile(encryptedFileBytes []byte) *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	// Remember to clean up the file afterwards
	fmt.Println("Created File: " + tmpFile.Name())

	// Example writing to the file
	if _, err = tmpFile.Write(encryptedFileBytes); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	return tmpFile
}

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
		sarcoSession := emb.NewSarcophagusSession(context.Background())
		contractArch, err := sarcoSession.Archaeologists(emb.ArchAddress)
		if err != nil {
			log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed. Please check CONTRACT_ADDRESS is correct in the config file: %v", err)
		}

		currentPublicKeyBytes := append([]byte{4}, contractArch.CurrentPublicKey...)
		pubKeyEcdsa, err := crypto.UnmarshalPubkey(currentPublicKeyBytes)
		if err != nil {
			log.Fatalf("Error unmarshaling public key during embalmer update:", err)
		}

		pubkeyBytes := crypto.FromECDSAPub(pubKeyEcdsa)
		pubKey, err := btcec.ParsePubKey(pubkeyBytes, btcec.S256())
		if err != nil {
			log.Fatalf("error casting public key to btcec: %v", err)
		}
		encryptedBytes, err := btcec.Encrypt(pubKey, fileBytes)
		if err != nil {
			log.Fatalf("Error encrypting file: %v", err)
		}
		log.Printf("ENCRYPTED BYTES: %v", encryptedBytes)
		tmpFile := createTmpFile(encryptedBytes)
		defer os.Remove(tmpFile.Name())

		emb.UpdateSarcophagus(assetDoubleHashBytes, tmpFile.Name())

		// Close the file
		if err := tmpFile.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
