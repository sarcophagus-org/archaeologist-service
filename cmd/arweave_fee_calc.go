// Used to calculate arweave fees
// Will output estimated cost for creating the transaction
// As well as estimated fee per byte
package main

import (
	"context"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/arweave"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"log"
	"math/big"
	"math/rand"
	"strconv"
)

func main(){
	tr, _ := ar.InitArweaveTransactor("https://arweave.net:443")
	feeOneByte, _ := tr.Client.GetReward(context.Background(), genBytes(1))
	feeOneMB, _ := tr.Client.GetReward(context.Background(), genBytes(1000000))

	one, _ := strconv.Atoi(feeOneByte)
	mb, _ := strconv.Atoi(feeOneMB)

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