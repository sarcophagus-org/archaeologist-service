package main

import (
	"fmt"

	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	_ "github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
)

func main(){
	balance := ethereum.ArchBalance()
	fmt.Printf("Balance: %d\n", balance)
}
