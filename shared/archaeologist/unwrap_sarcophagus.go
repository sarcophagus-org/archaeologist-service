package archaeologist

import (
	"context"
	"crypto/ecdsa"
	"github.com/Dev43/arweave-go/api"
	"github.com/Dev43/arweave-go/utils"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
	"time"
)

// TODO: handle rewrapped sarc
func scheduleUnwrap(session *contracts.SarcophagusSession, arweaveClient *api.Client, resurrectionTime *big.Int, arch *models.Archaeologist, assetDoubleHash [32]byte, privateKey *ecdsa.PrivateKey, assetId string) {
	timeToUnwrap := time.Until(time.Unix(resurrectionTime.Int64(), 0))

	var privateKeyBytes [32]byte
	copy(privateKeyBytes[:], crypto.FromECDSA(privateKey))
	log.Printf("Current Private Key BYTES: %v", privateKeyBytes)

	time.AfterFunc(timeToUnwrap, func() {
		if resTime, ok := arch.Sarcophaguses[assetDoubleHash]; ok {
			if resTime.Cmp(resurrectionTime) == 0 {
				singleHash, err := generateSingleHash(arweaveClient, assetId, privateKey)
				if err != nil {
					log.Printf("Error generating single hash during unwrapping process. Unwrapping cancelled: %v", err)
				} else {

					/* TODO: Add retry on the estimate gas */
					/* If it reverts 3 times in a row, then... */
					/* Add a minute to the time */
					/*
						Estimate Gas is used to check if the unwrap will succeed
					*/
					err := estimateGasForUnwrap(arch, assetDoubleHash, singleHash, privateKeyBytes)

					if err != nil {
						log.Printf("unwrapping aborted, transaction will fail: %v", err)
					} else {
						tx, err := session.UnwrapSarcophagus(assetDoubleHash, singleHash, privateKeyBytes)
						if err != nil {
							log.Printf("Transaction reverted. There was an error unwrapping the sarcophagus: %v", err)
						} else {
							log.Printf("Unwrap Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
							log.Printf("Gas Used: %v", tx.Gas())
							log.Printf("AssetDoubleHash: %v", assetDoubleHash)

							/* Sarcophagus is unwrapped, remove from state */
							delete(arch.Sarcophaguses, assetDoubleHash)
						}
					}
				}
			} else {
				// Resurrection time is different in state for this sarc, meaning sarc has been rewrapped
				log.Printf("Sarco has been rewrapped, rescheduling!")
				scheduleUnwrap(session, arweaveClient, resTime, arch, assetDoubleHash, privateKey, assetId)
			}
		}
	})

	log.Println("Unwrap scheduled in:", timeToUnwrap)
}

func generateSingleHash(arweaveClient *api.Client, assetId string, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	dataString, err := arweaveClient.GetData(context.Background(), assetId)
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

func estimateGasForUnwrap(arch *models.Archaeologist, assetDoubleHash [32]byte, singleHash []byte, privateKeyBytes [32]byte) error {
	parsed, err := abi.JSON(strings.NewReader(contracts.SarcophagusABI))
	if err != nil {
		return err
	}

	input, err := parsed.Pack("unwrapSarcophagus", assetDoubleHash, singleHash, privateKeyBytes)
	if err != nil {
		return err
	}
	msg := ethereum.CallMsg{From: arch.SarcoSession.CallOpts.From, To: &arch.SarcoAddress, Data: input}
	gasLimit, err := arch.Client.EstimateGas(context.Background(), msg)
	log.Printf("gas limit for unwrapping: %v", gasLimit)
	if err != nil {
		return err
	}
	return nil
}