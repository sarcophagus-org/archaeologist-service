package archaeologist

import (
	"math/big"
	"testing"
)

func TestCalculateFreeBond(t *testing.T) {
	addFreeBond, _ := new(big.Int).SetString("8000000000000000000", 10)
	removeFromFreeBond, _ := new(big.Int).SetString("0", 10)

	calculatedFreeBond := calculateFreeBond(addFreeBond, removeFromFreeBond)

	if calculatedFreeBond.Cmp(addFreeBond) != 0 {
		t.Errorf("Free should equal %v, got %v", addFreeBond, calculatedFreeBond)
	}

	t.Log("CalculateFreeBond with positive addFreeBond passes")

	addFreeBond, _ = new(big.Int).SetString("0", 10)
	removeFromFreeBond, _ = new(big.Int).SetString("5000000000000000000", 10)

	calculatedFreeBond = calculateFreeBond(addFreeBond, removeFromFreeBond)
	if calculatedFreeBond.Cmp(big.NewInt(0).Neg(removeFromFreeBond)) != 0 {
		t.Errorf("Free should equal -%v, got %v", removeFromFreeBond, calculatedFreeBond)
	}

	t.Log("CalculateFreeBond with positive removeFromFreeBond passes")
}