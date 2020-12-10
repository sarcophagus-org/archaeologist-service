package embalmer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
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

func CreateTmpFile(encryptedFileBytes []byte) *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	// Remember to clean up the file afterwards
	fmt.Println("Created File: " + tmpFile.Name())

	// Example writing to the file
	if _, err = tmpFile.Write(encryptedFileBytes); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	return tmpFile
}