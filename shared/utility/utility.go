package utility

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"io"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func IsHex(str string) bool {
	return strings.HasPrefix(str, "0x")
}

func IsValidAddress(ethAddress string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(ethAddress)
}

func IsContract(address common.Address, client *ethclient.Client) bool {
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Could not get bytecode from contract address: %v", err)
	}

	isContract := len(bytecode) > 0
	return isContract
}

func PrivateKeyHexToECDSA(privateKey string) (*ecdsa.PrivateKey, error) {
	if IsHex(privateKey) {
		privateKey = privateKey[2:]
	}

	ethPrivateKey, err := crypto.HexToECDSA(privateKey)
	return ethPrivateKey, err
}

func PrivateToPublicKeyECDSA(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	pKey := privateKey.Public()
	publicKey, ok := pKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return publicKey
}

func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) common.Address {
	pubKey := PrivateToPublicKeyECDSA(privateKey)
	return crypto.PubkeyToAddress(*pubKey)
}

func ValidatePositiveNumber(num *big.Int, numField string) (*big.Int, error) {
	if num.Cmp(big.NewInt(0)) != 1 {
		return nil, fmt.Errorf("%s must be greater than 0. Please check the value in the config file", numField)
	}

	return num, nil
}

func TimeInFuture(unixTimestamp *big.Int) bool {
	timestamp := unixTimestamp.Int64()
	now := time.Now()
	unixNow := now.Unix()
	return timestamp >= unixNow
}

func TimeWithWindowInFuture(unixTimestamp *big.Int, window *big.Int) bool {
	timePlusWindow := big.NewInt(0).Add(unixTimestamp, window)
	return TimeInFuture(timePlusWindow)
}

func FileToBytes(file *os.File) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, file)

	return buf.Bytes(), err
}

func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}
