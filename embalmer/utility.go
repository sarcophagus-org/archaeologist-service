package embalmer

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func FileBytesToDoubleHashBytes(fileBytes []byte) [32]byte {
	assetSingleHash := crypto.Keccak256(fileBytes)
	assetDoubleHash := crypto.Keccak256(assetSingleHash)

	/* Convert Double Hash to 32 byte array */
	var assetDoubleHashBytes [32]byte
	copy(assetDoubleHashBytes[:], assetDoubleHash)

	return assetDoubleHashBytes
}