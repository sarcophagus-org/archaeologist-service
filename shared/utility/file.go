package utility

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"os"
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

func EncryptFile(file *os.File, publicKey *ecdsa.PublicKey) ([]byte, error) {
	fileBytes, _ := FileToBytes(file)

	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	pubKey, err := btcec.ParsePubKey(publicKeyBytes, btcec.S256())
	encryptedFileBytes, err := btcec.Encrypt(pubKey, fileBytes)

	return encryptedFileBytes, err
}

func DecryptFile(fileBytes []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privateKeyBytes)

	return btcec.Decrypt(privKey, fileBytes)
}