package ethereum

import (
	"fmt"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"log"
	"time"
)

func InitEthClient(ethNode string) (*ethclient.Client, error) {
	cli, err := ethclient.Dial(ethNode)
	if err != nil {
		return nil, fmt.Errorf("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	return cli, nil
}

func SarcoAddress(contractAddress string, client *ethclient.Client) (common.Address, error) {
	address := common.HexToAddress(contractAddress)
	if isContract := utility.IsContract(address, client); !isContract {
		return address, fmt.Errorf("Config value CONTRACT_ADDRESS is not a valid contract. Please check the value is correct")
	}

	return address, nil
}

func WaitMined(client *ethclient.Client, txHash common.Hash, label string) error {
	var miningErr error = nil

	log.Printf("Waiting for %s transaction to be mined", label)

	for {
		_, pending, err := client.TransactionByHash(context.Background(), txHash)

		if err != nil {
			miningErr = err
			break
		}

		fmt.Print(".")
		time.Sleep(1 * time.Second)

		if !pending {
			fmt.Println("")
			break
		}
	}

	return miningErr
}