package ethereum

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func InitEthClient(ethNode string) *ethclient.Client {
	cli, err := ethclient.Dial(ethNode)
	if err != nil {
		log.Fatalf("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	return cli
}

func SarcoAddress(contractAddress string, client *ethclient.Client) common.Address {
	address := common.HexToAddress(contractAddress)
	if isContract := utility.IsContract(address, client); !isContract {
		log.Fatal("Config value CONTRACT_ADDRESS is not a valid contract. Please check the value is correct.")
	}

	return address
}