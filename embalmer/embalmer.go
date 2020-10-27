package embalmer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

type Embalmer struct {
	Client *ethclient.Client
	ArchPublicKeyBytes []byte
	EmbalmerPrivateKey *ecdsa.PrivateKey
	EmbalmerAddress common.Address
	SarcoAddress common.Address
	SarcophagusContract *contracts.Sarcophagus
	SarcophagusTokenContract *contracts.Token
	ResurrectionTime int64
	StorageFee int64
	DiggingFee int64
	Bounty int64
}

func (embalmer *Embalmer)  initAuth() *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(embalmer.EmbalmerPrivateKey)
	auth.Nonce = nil // uses nonce of pending state
	auth.Value = big.NewInt(0)
	auth.GasLimit = 0   // 0 estimates gas limit
	auth.GasPrice = nil // nil suggests gas price

	return auth
}

func (embalmer *Embalmer) NewSarcophagusSession(ctx context.Context) (session contracts.SarcophagusSession) {
	auth := embalmer.initAuth()

	return contracts.SarcophagusSession{
		Contract:     embalmer.SarcophagusContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

func (embalmer *Embalmer) NewSarcophagusTokenSession(ctx context.Context) (session contracts.TokenSession) {
	auth := embalmer.initAuth()

	return contracts.TokenSession{
		Contract:     embalmer.SarcophagusTokenContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

func (embalmer *Embalmer) EmbalmerSarcoBalance() *big.Int {
	balance, err := embalmer.SarcophagusTokenContract.BalanceOf(&bind.CallOpts{}, embalmer.EmbalmerAddress)

	if err != nil {
		log.Fatalf("Could not get sarcophagus balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func isHex(str string) bool {
	return strings.HasPrefix(str, "0x")
}

func (embalmer *Embalmer) approveCreateSarcophagusTransfer(session *contracts.TokenSession, approvalAmount *big.Int) {
	tx, err := session.Approve(
		embalmer.SarcoAddress,
		approvalAmount,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Approval Transaction for %v Create Sarcophagus Tokens successful. Transaction ID: %v", approvalAmount, tx.Hash().Hex())
	log.Printf("Gas Used for Approval: %v", tx.Gas())
}

func (embalmer *Embalmer) CreateSarcophagus(recipientPrivateKey string) {
	/* Initialize recipient public key bytes */
	if utility.IsHex(recipientPrivateKey) {
		recipientPrivateKey = recipientPrivateKey[2:]
	}
	recipPrivKey, _ := crypto.HexToECDSA(recipientPrivateKey)
	pub := recipPrivKey.Public().(*ecdsa.PublicKey)
	recipientPublicKeyBytes := crypto.FromECDSAPub(pub)[1:]

	/* Setup Asset Double Hash */
	/* Arbitrary initial value for sake of testing */
	/* Note: We may need to use solsha3 library to make these hashes, undetermined yet */
	/* https://github.com/miguelmota/go-solidity-sha3 */

	assetSingleHash := crypto.Keccak256([]byte{1, 2, 3, 4})
	assetDoubleHash := crypto.Keccak256(assetSingleHash)

	/* Sign asset double hash */
	signedAssetDoubleHash, err := crypto.Sign(assetDoubleHash, embalmer.EmbalmerPrivateKey)
	if err != nil {
		log.Fatalf("error signing asset double hash: %v", err)
	}

	/* Convert Double Hash to 32 byte array */
	var assetDoubleHashBytes [32]byte
	copy(assetDoubleHashBytes[:], assetDoubleHash)

	/* String for purposes of passing in as form param to arch http server */
	signedAssetDoubleHashString := hex.EncodeToString(signedAssetDoubleHash)

	log.Println("***CREATING SARCOPHAGUS***")
	log.Println("Signed Asset Double Hash:", signedAssetDoubleHashString)

	sarcoTokenSession := embalmer.NewSarcophagusTokenSession(context.Background())
	approvalAmount := new(big.Int).Add(new(big.Int).Add(big.NewInt(embalmer.Bounty), big.NewInt(embalmer.DiggingFee)), big.NewInt(embalmer.StorageFee))
	embalmer.approveCreateSarcophagusTransfer(&sarcoTokenSession, approvalAmount)

	sarcoSession := embalmer.NewSarcophagusSession(context.Background())
	tx, err := sarcoSession.CreateSarcophagus(
		"My Sarcophagus",
		embalmer.ArchPublicKeyBytes,
		big.NewInt(embalmer.ResurrectionTime),
		big.NewInt(embalmer.StorageFee),
		big.NewInt(embalmer.DiggingFee),
		big.NewInt(embalmer.Bounty),
		assetDoubleHashBytes,
		recipientPublicKeyBytes,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error creating Sarcophagus: %v", err)
	}

	log.Printf("Create Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}
