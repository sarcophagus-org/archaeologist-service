package archaeologist

import (
	"context"
	"crypto/ecdsa"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func InitializeArchaeologist(arch *models.Archaeologist, config *models.Config) {
	var err error

	arch.FreeBond = calculateFreeBond(config.ADD_TO_FREE_BOND, config.REMOVE_FROM_FREE_BOND)
	arch.Client = initEthClient(config.ETH_NODE)
	arch.PrivateKey, err = utility.PrivateKeyHexToECDSA(config.ETH_PRIVATE_KEY)
	if err != nil {
		log.Fatalf("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	arch.PublicKey = utility.PrivateToPublicKeyECDSA(arch.PrivateKey)
	arch.PublicKeyBytes = crypto.FromECDSAPub(arch.PublicKey)[1:]
	arch.ArchAddress = initArchAddress(config.PAYMENT_ADDRESS, arch.PublicKey, arch.Client)
	arch.SarcoAddress = initSarcoAddress(config.CONTRACT_ADDRESS, arch.Client)
	arch.SarcoSession = initSarcophagusSession(arch.SarcoAddress, arch.Client, arch.PrivateKey)
	arch.TokenSession = initTokenSession(config.TOKEN_ADDRESS, arch.Client, arch.PrivateKey)
	arch.FeePerByte = config.FEE_PER_BYTE
	arch.MinBounty = config.MIN_BOUNTY
	arch.MinDiggingFee = config.MIN_DIGGING_FEE
	arch.MaxResurectionTime = config.MAX_RESURRECTION_TIME
	arch.Endpoint = config.ENDPOINT
	arch.FilePort = config.FILE_PORT
}

func calculateFreeBond(addFreeBond int64, removeFreeBond int64) int64 {
	var archFreeBond int64 = 0

	if addFreeBond > 0 {
		if removeFreeBond > 0 {
			log.Fatal("ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND cannot both be > 0")
		}
		archFreeBond = addFreeBond
	} else if removeFreeBond > 0 {
		archFreeBond = int64(-1) * removeFreeBond
	}

	return archFreeBond
}

func initEthClient(ethNode string) *ethclient.Client {
	cli, err := ethclient.Dial(ethNode)
	if err != nil {
		log.Fatalf("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	return cli
}

func initArchAddress(paymentAddress string, publicKey *ecdsa.PublicKey, client *ethclient.Client) common.Address {
	// If the optional payment address is supplied, verify it is a valid address
	var archAddress common.Address

	if paymentAddress != "" {
		if utility.IsValidAddress(paymentAddress) && !utility.IsContract(common.HexToAddress(paymentAddress), client) {
			archAddress = common.HexToAddress(paymentAddress)
		} else {
			log.Fatal("Payment address supplied in config is invalid. Please check that address.")
		}
	} else {
		archAddress = crypto.PubkeyToAddress(*publicKey)
	}

	return archAddress
}

func initSarcoAddress(contractAddress string, client *ethclient.Client) common.Address {
	address := common.HexToAddress(contractAddress)
	if isContract := utility.IsContract(address, client); !isContract {
		log.Fatal("Config value CONTRACT_ADDRESS is not a valid contract. Please check the value is correct.")
	}

	return address
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