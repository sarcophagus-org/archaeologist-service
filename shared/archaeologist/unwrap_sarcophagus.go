package archaeologist

import (
	"context"
	"crypto/ecdsa"
	"github.com/Dev43/arweave-go/api"
	"github.com/Dev43/arweave-go/utils"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	eth "github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const (
	UNWRAP_RETRY_LIMIT = 2 // How many attempts after first failure to retry the unwrapping if it fails
	UNWRAP_RETRY_INTERVAL_LB = 100 // lower bound of retry (will be multiplied by 10 milliseconds)
	UNWRAP_RETRY_INTERVAL_UB = 1000 // upper bound of retry (will be multiplied by 10 milliseconds)
)
var mutex = &sync.Mutex{}

// TODO: handle rewrapped sarc
func scheduleUnwrap(session *contracts.SarcophagusSession, arweaveClient *api.Client, resurrectionTime *big.Int, arch *models.Archaeologist, assetDoubleHash [32]byte, privateKey *ecdsa.PrivateKey, assetId string) {
	timeToUnwrap := time.Until(time.Unix(resurrectionTime.Int64(), 0))

	var privateKeyBytes [32]byte
	copy(privateKeyBytes[:], crypto.FromECDSA(privateKey))
	log.Printf("Current Private Key BYTES: %v", privateKeyBytes)

	time.AfterFunc(timeToUnwrap, func() {
		if resTime, ok := arch.Sarcophaguses[assetDoubleHash]; ok {
			if resTime.Cmp(resurrectionTime) == 0 {
				_, err := generateSingleHash(arweaveClient, assetId, privateKey)
				if err != nil {
					log.Printf("Error generating single hash during unwrapping process. Unwrapping cancelled: %v", err)
				} else {
					// TODO: Do we need to remove the sarco from state if the unwrap fails?
					mutex.Lock()
					attempts, ok := arch.UnwrapAttempts[assetDoubleHash]
					mutex.Unlock()

					if !ok {
						attempts = 1
					} else {
						attempts += 1
					}
					mutex.Lock()
					arch.UnwrapAttempts[assetDoubleHash] = attempts
					mutex.Unlock()

					/*
						Estimate Gas is used to check if the unwrap will succeed
					*/
					err := estimateGasForUnwrap(arch, assetDoubleHash, privateKeyBytes)

					if err != nil {
						log.Printf("Unwrapping aborted, transaction will fail: %v", err)
						if attempts <= UNWRAP_RETRY_LIMIT {
							time.Sleep(randomRetryInterval() * time.Millisecond)
							scheduleUnwrap(session, arweaveClient, resTime, arch, assetDoubleHash, privateKey, assetId)
						}
					} else {
						txn, err := session.UnwrapSarcophagus(assetDoubleHash, privateKeyBytes)
						if err != nil {
							log.Printf("Transaction reverted. There was an error unwrapping the sarcophagus: %v", err)
							if attempts <= UNWRAP_RETRY_LIMIT {
								time.Sleep(randomRetryInterval() * time.Millisecond)
								scheduleUnwrap(session, arweaveClient, resTime, arch, assetDoubleHash, privateKey, assetId)
							}
						} else {
							log.Printf("Unwrap Sarcophagus Transaction Submitted. Transaction ID: %s", txn.Hash().Hex())
							log.Printf("Gas Used: %v", txn.Gas())
							log.Printf("AssetDoubleHash: %v", assetDoubleHash)

							err = eth.WaitMined(arch.Client, txn.Hash(), "Unwrap Sarcophagus")

							if err != nil {
								if attempts <= UNWRAP_RETRY_LIMIT {
									log.Printf("There was an error mining the update archaeologist transaction: %v. Retrying...", err)
									scheduleUnwrap(session, arweaveClient, resTime, arch, assetDoubleHash, privateKey, assetId)
								}
							} else {
								log.Printf("Unwrap Sarcophagus Transaction Successful. Transaction ID: %s", txn.Hash().Hex())
							}

							/* Remove from state */
							arch.RemoveArchSarcophagus(assetDoubleHash)
						}
					}
				}
			} else {
				// Resurrection time is different in state for this sarc, meaning sarc has been rewrapped
				log.Printf("Sarco has been rewrapped, rescheduling!")
				scheduleUnwrap(session, arweaveClient, resTime, arch, assetDoubleHash, privateKey, assetId)
			}
		} else {
			// Sarcophagus does not exist in state
			// It has either been cleaned / buried / cancelled
			// Do nothing
			log.Printf("Unwrapping cancelled. Sarcophagus was cancelled, buried, or cleaned, or archaeologist was successfully accused.")
		}
	})

	log.Println("Unwrap scheduled in:", timeToUnwrap)
}

func randomRetryInterval() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration((rand.Intn(UNWRAP_RETRY_INTERVAL_UB - UNWRAP_RETRY_INTERVAL_LB) + UNWRAP_RETRY_INTERVAL_LB) * 10)
}

func generateSingleHash(arweaveClient *api.Client, assetId string, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	log.Printf("Getting arweave data for assetID: %v", assetId)
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

func estimateGasForUnwrap(arch *models.Archaeologist, assetDoubleHash [32]byte, privateKeyBytes [32]byte) error {
	parsed, err := abi.JSON(strings.NewReader(contracts.SarcophagusABI))
	if err != nil {
		return err
	}

	input, err := parsed.Pack("unwrapSarcophagus", assetDoubleHash, privateKeyBytes)
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