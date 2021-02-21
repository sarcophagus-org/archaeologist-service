package utility

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"log"
)

// DecryptFile uses ecies library to decrypt file bytes
// Used to validate correct public key was used on the sarcophagus payload when handling file upload to arweave
func DecryptFile(fileBytes []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	log.Printf("priv key: %v", privateKey)
	log.Printf("fileBytes: %v", fileBytes)

	eciesPrivateKey := ecies.ImportECDSA(privateKey)
	return eciesPrivateKey.Decrypt(fileBytes, nil, nil)
}

// FileBytesToDoubleHashBytes is used to generate a sarcophagus identifier
// File bytes passed as an argument here are expected to be encrypted by the recipient's public key
func FileBytesToDoubleHashBytes(fileBytes []byte) [32]byte {
	assetSingleHash := crypto.Keccak256(fileBytes)
	assetDoubleHash := crypto.Keccak256(assetSingleHash)

	/* Convert Double Hash to 32 byte slice */
	var assetDoubleHashBytes [32]byte
	copy(assetDoubleHashBytes[:], assetDoubleHash)

	return assetDoubleHashBytes
}