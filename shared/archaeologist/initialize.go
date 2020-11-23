package archaeologist

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/Dev43/arweave-go/transactor"
	"github.com/Dev43/arweave-go/wallet"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/hdw"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
	"math/big"
)

func InitializeArchaeologist(arch *models.Archaeologist, config *models.Config) {
	// TODO: Program exits on first error. Update to track # of errors in config and output messages for all.
	// TODO: Validations for Port and IP Address -- consider testing opening/closing port

	var err error

	arch.FreeBond = calculateFreeBond(big.NewInt(config.ADD_TO_FREE_BOND), big.NewInt(config.REMOVE_FROM_FREE_BOND))
	arch.Client = ethereum.InitEthClient(config.ETH_NODE)
	arch.ArweaveTransactor = initArweaveTransactor(config.ARWEAVE_NODE)
	arch.ArweaveWallet = initArweaveWallet(config.ARWEAVE_KEY_FILE)
	arch.PrivateKey, err = utility.PrivateKeyHexToECDSA(config.ETH_PRIVATE_KEY)
	if err != nil {
		log.Fatalf("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}
	arch.ArchAddress = utility.PrivateKeyToAddress(arch.PrivateKey)
	arch.SarcoAddress = ethereum.SarcoAddress(config.CONTRACT_ADDRESS, arch.Client)
	arch.SarcoSession = initSarcophagusSession(arch.SarcoAddress, arch.Client, arch.PrivateKey)
	arch.TokenSession = initTokenSession(config.TOKEN_ADDRESS, arch.Client, arch.PrivateKey)
	arch.Wallet, err = hdwallet.NewFromMnemonic(config.MNEMONIC)
	if err != nil {
		log.Fatalf("could not setup HD wallet from mnemonic: %v", err)
	}

	arch.Sarcophaguses, arch.FileHandlers = buildSarcophagusesState(&arch.SarcoSession)

	arch.AccountIndex = 0
	arch.CurrentPrivateKey = hdw.PrivateKeyFromIndex(arch.Wallet, arch.AccountIndex)
	arch.CurrentPublicKeyBytes = hdw.PublicKeyBytesFromIndex(arch.Wallet, arch.AccountIndex)
	arch.PaymentAddress = validatePaymentAddress(config.PAYMENT_ADDRESS, arch.Client)
	arch.FeePerByte = utility.ValidatePositiveNumber(big.NewInt(config.FEE_PER_BYTE), "FEE_PER_BYTE")
	arch.MinBounty = utility.ValidatePositiveNumber(big.NewInt(config.MIN_BOUNTY), "MIN_BOUNTY")
	arch.MinDiggingFee = utility.ValidatePositiveNumber(big.NewInt(config.MIN_DIGGING_FEE), "MIN_DIGGING_FEE")
	arch.MaxResurectionTime = utility.ValidateTimeInFuture(big.NewInt(config.MAX_RESURRECTION_TIME), "MAX_RESURRECTION_TIME")
	arch.Endpoint = utility.ValidateIpAddress(config.ENDPOINT, "ENDPOINT")
	arch.FilePort = config.FILE_PORT
}

func buildSarcophagusesState (session *contracts.SarcophagusSession) (map[[32]byte]models.Sarcophagus, map[[32]byte]*big.Int) {
	sarcophaguses := map[[32]byte]models.Sarcophagus{}
	fileHandlers := map[[32]byte]*big.Int{}

	/* TODO:
		1. Build sarco states
			- Retrieve all sarcos on the contract
			- Grab the ones that use us as the arch
			- Determine the current public key based on # of sarcs
				- Set the account index (used by HD wallet)
				- If we can't do this, then we need to use current pub key on the contract to determine which index we are on for the HD wallet
					- Do this by iterating through pub keys on HD wallet to see which one matches
					- Put some upper bound on the # of public keys we search for
		2. Build open file handlers for any open 'created' sarcos that arent updated
			- Check block timestamp. If they are lapsed past a certain time --- what do we do?
		3. Schedule resurrections
	*/

	return sarcophaguses, fileHandlers
}

func calculateFreeBond(addFreeBond *big.Int, removeFreeBond *big.Int) *big.Int {
	var zero = big.NewInt(0)
	var archFreeBond = zero

	if addFreeBond.Cmp(zero) == 1 {
		if removeFreeBond.Cmp(zero) == 1 {
			log.Fatal("ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND cannot both be > 0")
		}
		archFreeBond = addFreeBond
	} else if removeFreeBond.Cmp(zero) == 1 {
		archFreeBond = archFreeBond.Neg(removeFreeBond)
	}

	return archFreeBond
}

func initArweaveTransactor(arweaveNode string) *transactor.Transactor {
	ar, err := transactor.NewTransactor(arweaveNode)

	if err != nil {
		log.Fatal("Could not connect to arweave node. Error: %v\n", err)
	}

	return ar
}

func initArweaveWallet(arweaveKeyFileName string) *wallet.Wallet {
	wallet := wallet.NewWallet()

	if err := wallet.LoadKeyFromFile(fmt.Sprintf("config/%s", arweaveKeyFileName)); err != nil {
		log.Fatal("Could not load config value ARWEAVE_KEY_FILE. Please check the config.yml file Error:", err)
	}

	return wallet
}

func validatePaymentAddress(paymentAddress string, client *ethclient.Client) common.Address {
	var archAddress common.Address

	if paymentAddress != "" {
		if utility.IsValidAddress(paymentAddress) && !utility.IsContract(common.HexToAddress(paymentAddress), client) {
			archAddress = common.HexToAddress(paymentAddress)
		} else {
			log.Fatal("Payment address supplied in config is invalid. Please check that address.")
		}
	}

	return archAddress
}

func initSarcophagusSession(contractAddress common.Address, client *ethclient.Client, privateKey *ecdsa.PrivateKey) contracts.SarcophagusSession {
	sarcoContract, err := contracts.NewSarcophagus(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Sarcophagus contract: %v", err)
	}

	session := NewSarcophagusSession(context.Background(), sarcoContract, privateKey)

	return session
}

func initTokenSession(tokenAddress string, client *ethclient.Client, privateKey *ecdsa.PrivateKey) contracts.TokenSession {
	address := common.HexToAddress(tokenAddress)
	tokenContract, err := contracts.NewToken(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Sarcophagus contract: %v", err)
	}

	session := NewTokenSession(context.Background(), tokenContract, privateKey)

	return session
}