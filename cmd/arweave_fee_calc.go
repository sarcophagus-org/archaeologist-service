// Utility program -- not used by archaeologist service
// Used to calculate arweave fees
// Will output estimated cost for creating the transaction
// As well as estimated fee per byte
package main

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"log"
	"math/big"
	"math/rand"
	"strconv"
)

func main(){
	arClient := ar.InitArweaveClient("https://arweave.net:443")
	feeOneByte, _ := arClient.GetTransactionPrice(genBytes(1), nil)
	feeOneMB, _ := arClient.GetTransactionPrice(genBytes(1000000), nil)

	one, _ := strconv.Atoi(strconv.FormatInt(feeOneByte, 10))
	mb, _ := strconv.Atoi(strconv.FormatInt(feeOneMB, 10))

	perByte := (mb - one) / 1000000
	transCost := one - perByte

	log.Printf("transaction creation cost: %v", utility.ToDecimal(big.NewInt(int64(transCost)), 12))
	log.Printf("Per Byte Cost: %v", utility.ToDecimal(big.NewInt(int64(perByte)), 12))
}

func genBytes(size int) []byte {
	randBytes := make([]byte, size)

	rand.Seed(200)
	rand.Read(randBytes)

	return randBytes
}