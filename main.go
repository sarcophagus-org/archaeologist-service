package main

import (
	"fmt"

	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
)

func main(){
	balance := ethereum.ArchBalance()
	fmt.Printf("Balance: %d\n", balance)
}
