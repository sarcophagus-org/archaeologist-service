package ethereum

import (
	"context"
	"crypto/ecdsa"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"regexp"

	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
)

var client *ethclient.Client
var archPrivateKey *ecdsa.PrivateKey
var archPublicKey *ecdsa.PublicKey
var archPublicKeyBytes []byte
var archAddress common.Address
var sarcoAddress common.Address
var sarcophagusContract *contracts.Sarcophagus
var sarcoTokenAddress common.Address
var sarcophagusTokenContract *contracts.Token
var freeBond int64

func ArchEthBalance() *big.Int {
	balance, err := client.BalanceAt(context.Background(), archAddress, nil)

	if err != nil {
		log.Fatalf("Could not get eth balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func ArchSarcoBalance() *big.Int {
	balance, err := sarcophagusTokenContract.BalanceOf(&bind.CallOpts{}, archAddress)

	if err != nil {
		log.Fatalf("Could not get sarcophagus balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func SetFreeBond(addFreeBond int64, removeFreeBond int64) {
	var archFreeBond int64 = 0

	if addFreeBond > 0 {
		if removeFreeBond > 0 {
			log.Fatal("ADD_TO_FREE_BOND and REMOVE_FROM_FREE_BOND cannot both be > 0")
		}
		archFreeBond = addFreeBond
	} else if removeFreeBond > 0 {
		archFreeBond = int64(-1) * removeFreeBond
	}

	freeBond = archFreeBond
}

func IsValidAddress(ethAddress string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(ethAddress)
}

func IsContract(address common.Address) bool {
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Could not get bytecode from contract address: %v", err)
	}

	isContract := len(bytecode) > 0
	return isContract
}

func InitSarcophagusContract(contractAddress string) {
	address := common.HexToAddress(contractAddress)
	if isContract := IsContract(address); !isContract {
		log.Fatal("Contract for config value CONTRACT_ADDRESS is not valid. Please check the value is correct.")
	}

	sarcoAddress = address
	sarcoContract, err := contracts.NewSarcophagus(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Sarcophagus contract: %v", err)
	}

	sarcophagusContract = sarcoContract
}

func InitSarcophagusTokenContract(tokenAddress string) {
	address := common.HexToAddress(tokenAddress)
	if isContract := IsContract(address); !isContract {
		log.Fatal("config value TOKEN_ADDRESS is not a valid contract. Please check the value is correct.")
	}

	sarcoTokenAddress = address
	tokenContract, err := contracts.NewToken(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Sarcophagus Token contract: %v", err)
	}

	sarcophagusTokenContract = tokenContract
}

func InitEthClient(ethNode string) {
	cli, err := ethclient.Dial(ethNode)
	if err != nil {
		log.Fatalf("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	client = cli
}

func InitEthKeysAndAddress(privateKey string, paymentAddress string) {
	if utility.IsHex(privateKey) {
		privateKey = privateKey[2:]
	}

	ethPrivKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	pub := ethPrivKey.Public()
	publicKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	archPublicKeyBytes = crypto.FromECDSAPub(publicKey)[1:]
	archPrivateKey = ethPrivKey

	// If the optional payment address is supplied, verify it is a valid address
	if paymentAddress != "" {
		if IsValidAddress(paymentAddress) && !IsContract(common.HexToAddress(paymentAddress)) {
			archAddress = common.HexToAddress(paymentAddress)
		} else {
			log.Fatal("Payment address supplied in config is invalid. Please check that address.")
		}
	} else {
		archAddress = crypto.PubkeyToAddress(*publicKey)
	}
}
