package archaeologist

import (
	"github.com/stretchr/testify/assert"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func TestCalculateFreeBond(t *testing.T) {
	addFreeBond, _ := new(big.Int).SetString("8000000000000000000", 10)
	removeFromFreeBond, _ := new(big.Int).SetString("0", 10)

	calculatedFreeBond, _ := calculateFreeBond(addFreeBond, removeFromFreeBond)

	if calculatedFreeBond.Cmp(addFreeBond) != 0 {
		t.Errorf("Free should equal %v, got %v", addFreeBond, calculatedFreeBond)
	}

	t.Log("CalculateFreeBond with positive addFreeBond passes")

	addFreeBond, _ = new(big.Int).SetString("0", 10)
	removeFromFreeBond, _ = new(big.Int).SetString("5000000000000000000", 10)

	calculatedFreeBond, _ = calculateFreeBond(addFreeBond, removeFromFreeBond)
	if calculatedFreeBond.Cmp(big.NewInt(0).Neg(removeFromFreeBond)) != 0 {
		t.Errorf("Free should equal -%v, got %v", removeFromFreeBond, calculatedFreeBond)
	}

	t.Log("CalculateFreeBond with positive removeFromFreeBond passes")

	addFreeBond, _ = new(big.Int).SetString("5000000000000000000", 10)
	removeFromFreeBond, _ = new(big.Int).SetString("5000000000000000000", 10)

	calculatedFreeBond, err := calculateFreeBond(addFreeBond, removeFromFreeBond)
	if err == nil {
		t.Errorf("Calculate free bond should throw an error when addFreeBond and removeFromFreeBond are positive, but is not.")
	}

	t.Log("CalculateFreeBond with positive and addFreeBond / removeFromFreeBond throws error passes")
}

func TestSetPaymentAddress(t *testing.T) {
	archAddress := common.HexToAddress("0x3968927caAbd8d5eF8ADe20B364E038ae785F855")
	paymentAddress := ""
	client := ethclient.Client{}

	address, err := setPaymentAddress(archAddress, paymentAddress, &client)
	assert.Nil(t, err)
	assert.Equal(t, archAddress, address)
}