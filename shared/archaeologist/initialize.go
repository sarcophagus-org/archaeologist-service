package archaeologist

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Dev43/arweave-go/api"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
)

func InitializeArchaeologist(arch *models.Archaeologist, config *models.Config) []string {
	var err error
	var errStrings []string

	/*
		Sets archaeologist struct fields.
		Keeps a running list of errors. If any exist, outputs them to the console log and exits the service.
		UX may be better without the running list, as some of these settings piggy back on each other.
	 */

	arch.FreeBond, err = calculateFreeBond(utility.ToWei(config.ADD_TO_FREE_BOND, 18), utility.ToWei(config.REMOVE_FROM_FREE_BOND, 18))
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.Client, err = ethereum.InitEthClient(config.ETH_NODE)
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.ArweaveTransactor, err = ar.InitArweaveTransactor(config.ARWEAVE_NODE)
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.ArweaveWallet, err = ar.InitArweaveWallet(config.ARWEAVE_KEY_FILE)
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.PrivateKey, err = utility.PrivateKeyHexToECDSA(config.ETH_PRIVATE_KEY)
	if err != nil {
		errStrings = append(errStrings, fmt.Sprintf("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err))
	}

	arch.ArchAddress = utility.PrivateKeyToAddress(arch.PrivateKey)
	arch.SarcoAddress, err = ethereum.SarcoAddress(config.CONTRACT_ADDRESS, arch.Client)
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.SarcoSession, err = initSarcophagusSession(arch.SarcoAddress, arch.Client, arch.PrivateKey)
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.TokenSession, err = initTokenSession(config.TOKEN_ADDRESS, arch.Client, arch.PrivateKey)
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.Wallet, err = hdwallet.NewFromMnemonic(config.MNEMONIC)
	if err != nil {
		errStrings = append(errStrings, fmt.Sprintf("could not setup HD wallet from mnemonic: %v", err))
	}

	arch.PaymentAddress, err = setPaymentAddress(arch.ArchAddress, config.PAYMENT_ADDRESS, arch.Client)
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.FeePerByte, err = utility.ValidatePositiveNumber(utility.ToWei(config.FEE_PER_BYTE, 18), "FEE_PER_BYTE")
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.ArweaveMultiplier, err = decimal.NewFromString(config.ARWEAVE_MULTIPLIER)
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	if !arch.ArweaveMultiplier.IsPositive() {
		errStrings = append(errStrings, "Arweave Multiplier must be positive")
	}

	arch.MinBounty, err = utility.ValidatePositiveNumber(utility.ToWei(config.MIN_BOUNTY, 18), "MIN_BOUNTY")
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.MinDiggingFee, err = utility.ValidatePositiveNumber(utility.ToWei(config.MIN_DIGGING_FEE, 18), "MIN_DIGGING_FEE")
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.MaxResurectionTime = stringToBigInt(config.MAX_RESURRECTION_TIME)

	arch.Endpoint = config.ENDPOINT
	if err != nil {
		errStrings = append(errStrings, err.Error())
	}

	arch.FilePort = config.FILE_PORT

	arch.Sarcophaguses, arch.SarcophagusesAccountIndex, arch.FileHandlers, arch.AccountIndex = buildSarcophagusesState(arch)

	arch.CurrentPrivateKey = hdw.PrivateKeyFromIndex(arch.Wallet, arch.AccountIndex)
	arch.CurrentPublicKeyBytes = hdw.PublicKeyBytesFromIndex(arch.Wallet, arch.AccountIndex)
	arch.UnwrapAttempts	= map[[32]byte]int{}

	return errStrings
}

func buildSarcophagusesState (arch *models.Archaeologist) (map[[32]byte]*big.Int, map[[32]byte]int, map[[32]byte]*big.Int, int) {
	var sarcophaguses = map[[32]byte]*big.Int{}
	var sarcophagusesAccountIndex = map[[32]byte]int{}
	var fileHandlers = map[[32]byte]*big.Int{}

	// Create a slice of sarcos (double hashes) indexed by public keys
	var pubKeyMap = map[[64]byte][][32]byte{}

	var accountIndex = 0

	sarcoCount, err := arch.SarcoSession.ArchaeologistSarcophagusCount(arch.ArchAddress)
	if err != nil {
		log.Fatalf("Call to ArchaeologistSarcophagusCount in Contract failed. Please check CONTRACT_ADDRESS is correct in the config file: %v", err)
	}

	/*
		Iterate through all sarcos
		For any sarcos where we are the arch, determine state of sarco and build service state
		Schedule rewraps if Sarco is updated and resurrection time + window is in future

		// TODO: We only cleanup our own sarcos. We could be checking if *any* sarcos need cleanup.
	*/

	for i := big.NewInt(0); i.Cmp(sarcoCount) == -1; i = big.NewInt(0).Add(i, big.NewInt(1)) {
		doubleHash, _ := arch.SarcoSession.ArchaeologistSarcophagusIdentifier(arch.ArchAddress, i)
		sarco, _ := arch.SarcoSession.Sarcophagus(doubleHash)

		/*
			Sarco States:
			0 - Does not Exist
			1 - Exists
			2 - Done
		*/

		switch state := sarco.State; state {
		// Sarco Exists
		case 1:
			if utility.TimeWithWindowInFuture(sarco.ResurrectionTime, sarco.ResurrectionWindow) {

				// Track if the archaeologist public key on the sarcophagus matches
				// our current account index public key
				// If it does, no other sarcophagus has used this public key yet
				currentPublicKey := hdw.PublicKeyBytesFromIndex(arch.Wallet, accountIndex)
				pubKeyMatches := bytes.Equal(sarco.ArchaeologistPublicKey, currentPublicKey)

				// Update the public key mapping
				var currentPublicKeyIndex [64]byte
				copy(currentPublicKeyIndex[:], currentPublicKey)
				doubleHashList := append(pubKeyMap[currentPublicKeyIndex], doubleHash)
				pubKeyMap[currentPublicKeyIndex] = doubleHashList

				if sarco.AssetId == "" {
					// This is a created sarc that is not updated
					// If our current pub key matches the one on the sarcophagus,
					// this means no sarc has used our current public key yet
					// and we need to create a file handler for this sarc
					// as a file could potentially be sent for this sarcophagus
					if pubKeyMatches {
						fileHandlers[doubleHash] = sarco.StorageFee
					}
				} else {
					// We have an updated sarc that is updated but not unwrapped
					// Schedule an unwrap using the current account index private key
					// and increment the account index as this key pair has been used.
					privateKey := hdw.PrivateKeyFromIndex(arch.Wallet, accountIndex)
					sarcophagusesAccountIndex[doubleHash] = accountIndex
					scheduleUnwrap(&arch.SarcoSession, arch.ArweaveTransactor.Client.(*api.Client), sarco.ResurrectionTime, arch, doubleHash, privateKey, sarco.AssetId)
					fileHandlers = map[[32]byte]*big.Int{}
					accountIndex += 1

					// Remove any other sarcos from state that used this public key
					for i := range pubKeyMap[currentPublicKeyIndex] {
						if !bytes.Equal(pubKeyMap[currentPublicKeyIndex][i][:], doubleHash[:]) {
							delete(sarcophaguses, pubKeyMap[currentPublicKeyIndex][i])
						}
					}
				}

				if pubKeyMatches {
					sarcophaguses[doubleHash] = sarco.ResurrectionTime
				}
			} else {
				// Sarc's unwrap time + resurrection window is in the past
				log.Printf("Sarcophagus did not get unwrapped in time: %v", doubleHash)
				if sarco.AssetId != "" {
					// Sarco has been updated, increment account index as this sarco uses one of our key pairs.
					// Clear file handlers b/c we only want file handlers for our current account index
					fileHandlers = map[[32]byte]*big.Int{}
					accountIndex += 1
				}

				// Lets get some money by cleaning it up
				tx, err := arch.SarcoSession.CleanUpSarcophagus(doubleHash, arch.PaymentAddress)
				if err != nil {
					log.Printf("Cleanup Sarcophagus error: %v", err)
				}
				log.Printf("Cleanup Sarcophagus Tx Submitted. Transaction ID: %s", tx.Hash().Hex())
				log.Printf("Gas Used: %v", tx.Gas())
			}
		// Sarco is Done
		case 2:
			if sarco.AssetId != "" {
				// Sarco has been updated, increment account index as this sarco uses one of our key pairs.
				// Clear file handlers b/c we only want file handlers for our current account index
				fileHandlers = map[[32]byte]*big.Int{}
				accountIndex += 1
			}
		}
	}

	log.Printf("Sarcophaguses not yet complete: %v", sarcophaguses)
	log.Printf("Sarcophaguses waiting for a file: %v", fileHandlers)
	log.Printf("Current Account Index: %v", accountIndex)

	return sarcophaguses, sarcophagusesAccountIndex, fileHandlers, accountIndex
}

func calculateFreeBond(addFreeBond *big.Int, removeFreeBond *big.Int) (*big.Int, error) {
	var zero = big.NewInt(0)
	var archFreeBond = zero

	if addFreeBond.Cmp(zero) == 1 {
		if removeFreeBond.Cmp(zero) == 1 {
			return big.NewInt(0), fmt.Errorf("ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND cannot both be > 0")
		}
		archFreeBond = addFreeBond
	} else if removeFreeBond.Cmp(zero) == 1 {
		archFreeBond = archFreeBond.Neg(removeFreeBond)
	}

	return archFreeBond, nil
}

func setPaymentAddress(archAddress common.Address, paymentAddress string, client *ethclient.Client) (common.Address, error) {
	var addy common.Address

	if paymentAddress != "" {
		if utility.IsValidAddress(paymentAddress) && !utility.IsContract(common.HexToAddress(paymentAddress), client) {
			addy = common.HexToAddress(paymentAddress)
		} else {
			return common.Address{}, fmt.Errorf("payment address supplied in config is invalid, please check that address")
		}
	} else {
		addy = archAddress
	}

	return addy, nil
}

func initSarcophagusSession(contractAddress common.Address, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (contracts.SarcophagusSession, error) {
	sarcoContract, err := contracts.NewSarcophagus(contractAddress, client)
	if err != nil {
		return contracts.SarcophagusSession{}, fmt.Errorf("failed to instantiate Sarcophagus contract: %v", err)
	}

	session := NewSarcophagusSession(context.Background(), sarcoContract, privateKey)

	return session, nil
}

func initTokenSession(tokenAddress string, client *ethclient.Client, privateKey *ecdsa.PrivateKey) (contracts.TokenSession, error) {
	address := common.HexToAddress(tokenAddress)
	tokenContract, err := contracts.NewToken(address, client)
	if err != nil {
		return contracts.TokenSession{}, fmt.Errorf("failed to instantiate Sarcophagus contract: %v", err)
	}

	session := NewTokenSession(context.Background(), tokenContract, privateKey)

	return session, nil
}

func stringToBigInt(val string) *big.Int {
	intVal, ok := new(big.Int).SetString(val, 10)

	if !ok {
		log.Fatalf("Error casting string to big int: %v", val)
	}

	return intVal
}