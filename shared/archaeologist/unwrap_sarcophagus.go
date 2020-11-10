package archaeologist

import (
	"crypto/ecdsa"
	"github.com/Dev43/arweave-go/api"
	"github.com/Dev43/arweave-go/utils"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"time"
	"context"
)

// TODO: handle rewrapped sarc
func scheduleUnwrap(session *contracts.SarcophagusSession, client *api.Client, resurrectionTime *big.Int, assetDoubleHash [32]byte, privateKey *ecdsa.PrivateKey, assetId string) {
	timeToUnwrap := time.Until(time.Unix(resurrectionTime.Int64(), 0))

	var privateKeyBytes [32]byte
	copy(privateKeyBytes[:], crypto.FromECDSA(privateKey))
	log.Printf("Current Private Key BYTES: %v", privateKeyBytes)

	time.AfterFunc(timeToUnwrap, func() {
		singleHash, err := generateSingleHash(client, assetId, privateKey)
		if err != nil {
			log.Printf("Error generating single hash during unwrapping process: %v", err)
		}

		tx, err := session.UnwrapSarcophagus(assetDoubleHash, singleHash, privateKeyBytes)
		if err != nil {
			log.Fatalf("Transaction reverted. There was an error unwrapping the sarcophagus: %v", err)
		}

		log.Printf("Unwrap Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
		log.Printf("Gas Used: %v", tx.Gas())
		log.Printf("AssetDoubleHash: %v", assetDoubleHash)
	})

	log.Println("Unwrap scheduled in:", timeToUnwrap)
}

func generateSingleHash(client *api.Client, assetId string, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	dataString, err := client.GetData(context.Background(), assetId)
	if err != nil {
		return nil, err
	}

	dataBytes, err := utils.DecodeString(dataString)
	if err != nil {
		return nil, err
	}

	decryptedDataBytes, err := utility.DecryptFile(dataBytes, privateKey)
	if err != nil {
		return nil, err
	}

	return crypto.Keccak256(decryptedDataBytes), nil
}