// Various utility functions used for creating and interating with ethereum client

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

// InitEthClient connects to and returns eth client supplied in the config file
func InitEthClient(ethNode string) (*ethclient.Client, error) {
	cli, err := ethclient.Dial(ethNode)
	if err != nil {
		return nil, fmt.Errorf("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	return cli, nil
}

// SarcoAddress validates and returns the address of the Sarcophagus address
// using the contract address supplied in the config file
func SarcoAddress(contractAddress string, client *ethclient.Client) (common.Address, error) {
	address := common.HexToAddress(contractAddress)
	if isContract := utility.IsContract(address, client); !isContract {
		return address, fmt.Errorf("Config value CONTRACT_ADDRESS is not a valid contract. Please check the value is correct")
	}

	return address, nil
}

// WaitMined waits for an ethereum transaction to be confirmed (or fail)
// Returns an error if there was an error confirming the transaction
func WaitMined(client *ethclient.Client, txHash common.Hash, label string) error {
	var miningErr error = nil
	retries := 4
	time.Sleep(3 * time.Second)
	log.Printf("Waiting for %s transaction to be mined", label)

	for {
		_, pending, err := client.TransactionByHash(context.Background(), txHash)

		if err != nil {
			if retries > 0 {
				log.Printf("%v transaction not found, retrying", label)
				time.Sleep(3 * time.Second)
				retries -= 1
			} else {
				miningErr = err
				break
			}
		}

		fmt.Print(".")
		time.Sleep(1 * time.Second)

		if !pending {
			fmt.Println("")
			break
		}
	}

	time.Sleep(6 * time.Second)

	return miningErr
}