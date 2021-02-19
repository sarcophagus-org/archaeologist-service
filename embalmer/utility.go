package embalmer

import (
	"github.com/ethereum/go-ethereum/crypto"
	"math/rand"
)

func FileBytesToDoubleHashBytes(fileBytes []byte) [32]byte {
	assetSingleHash := crypto.Keccak256(fileBytes)
	assetDoubleHash := crypto.Keccak256(assetSingleHash)

	/* Convert Double Hash to 32 byte slice */
	var assetDoubleHashBytes [32]byte
	copy(assetDoubleHashBytes[:], assetDoubleHash)

	return assetDoubleHashBytes
}

func DoubleHashBytesFromSeed(seed int64, size int) ([]byte, [32]byte) {
	/* Generate random bytes to use as payload for each sarco */
	fileBytes := make([]byte, size)

	rand.Seed(seed)
	rand.Read(fileBytes)
	return fileBytes, FileBytesToDoubleHashBytes(fileBytes)
}