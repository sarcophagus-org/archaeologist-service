package ethereum

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
)

var client *ethclient.Client
var archPrivateKey *ecdsa.PrivateKey
var archPublicKey *ecdsa.PublicKey
var archPublicKeyBytes []byte
var archAddress common.Address
var sarcophagusContract *contracts.Sarcophagus
var sarcophagusTokenContract *contracts.Token
var freeBond int64

func EthBalance() *big.Int {
	balance, _ := client.BalanceAt(context.Background(), archAddress, nil)

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
		archFreeBond = removeFreeBond
	}

	freeBond = archFreeBond
}

func IsContract(address common.Address) bool {
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Could not get bytecode from contract address: %v", err)
	}

	isContract := len(bytecode) > 0
	return isContract
}

func ArchaeologistCount() *big.Int {
	archCount, err := sarcophagusContract.ArchaeologistCount(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve archaeologist count: %v", err)
	}

	return archCount
}

func TokenName() string {
	tokenName, err := sarcophagusTokenContract.Name(&bind.CallOpts{})

	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}

	return tokenName
}

func GetSuggestedGasPrice() *big.Int {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("couldn't get the suggested gas price: %v", err)
	}
	return gasPrice
}

func GetNonce() uint64 {
	nonce, err := client.PendingNonceAt(context.Background(), archAddress)
	if err != nil {
		log.Fatal(err)
	}

	return nonce
}

func InitSarcophagusContract(contractAddress string) {
	address := common.HexToAddress(contractAddress)
	if isContract := IsContract(address); !isContract {
		log.Fatal("Contract for config value CONTRACT_ADDRESS is not valid. Please check the value is correct.")
	}

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

func InitEthKeysAndAddress(privateKey string) {
	if privateKey[0:2] == "0x" {
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
	archAddress = crypto.PubkeyToAddress(*publicKey)
	archPrivateKey = ethPrivKey
}
