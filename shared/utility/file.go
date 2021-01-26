package utility

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"log"
)

func SignatureValid(signature string, signedMsg []byte, msgSenderAddress common.Address) bool {
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		log.Printf("Could not decode signature: %v", err)
		return false
	}

	// if using [32]byte, postfix the signedMsg with [:]
	hopefullySignerPubKey, err := crypto.SigToPub(signedMsg, signatureBytes)
	if err != nil {
		log.Printf("Could not derive public key from hash and signature: %v", err)
		return false
	}

	hopefullySignerAddress := crypto.PubkeyToAddress(*hopefullySignerPubKey)
	if hopefullySignerAddress != msgSenderAddress {
		log.Printf("Address derived from the provided signature does not match the address sender")
		return false
	}

	log.Printf("Signature is valid!")

	return true
}

func DecryptFile(fileBytes []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	log.Printf("priv key: %v", privateKey)
	log.Printf("fileBytes: %v", fileBytes)

	eciesPrivateKey := ecies.ImportECDSA(privateKey)
	return eciesPrivateKey.Decrypt(fileBytes, nil, nil)
}

func FileBytesToDoubleHashBytes(fileBytes []byte) [32]byte {
	assetSingleHash := crypto.Keccak256(fileBytes)
	assetDoubleHash := crypto.Keccak256(assetSingleHash)

	/* Convert Double Hash to 32 byte slice */
	var assetDoubleHashBytes [32]byte
	copy(assetDoubleHashBytes[:], assetDoubleHash)

	return assetDoubleHashBytes
}