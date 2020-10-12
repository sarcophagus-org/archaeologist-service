package embalmer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

var client *ethclient.Client
var archPublicKeyBytes []byte
var embalmerPrivateKey *ecdsa.PrivateKey
var embalmerAddress common.Address
var sarcoAddress common.Address
var sarcophagusContract *contracts.Sarcophagus
var sarcophagusTokenContract *contracts.Token

func initAuth () *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(embalmerPrivateKey)
	auth.Nonce = nil // uses nonce of pending state
	auth.Value = big.NewInt(0)
	auth.GasLimit = 0 // 0 estimates gas limit
	auth.GasPrice = nil // nil suggests gas price

	return auth
}

func NewSarcophagusSession(ctx context.Context) (session contracts.SarcophagusSession) {
	auth := initAuth()

	return contracts.SarcophagusSession{
		Contract: sarcophagusContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

func NewSarcophagusTokenSession(ctx context.Context) (session contracts.TokenSession) {
	auth := initAuth()

	return contracts.TokenSession{
		Contract: sarcophagusTokenContract,
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}
}

func InitSarcophagusContract(contractAddress string) {
	address := common.HexToAddress(contractAddress)
	sarcoContract, err := contracts.NewSarcophagus(address, client)
	sarcoAddress = address
	if err != nil {
		log.Fatalf("Failed to instantiate Sarcophagus contract: %v", err)
	}

	sarcophagusContract = sarcoContract
}

func InitSarcophagusTokenContract(tokenAddress string) {
	address := common.HexToAddress(tokenAddress)
	tokenContract, err := contracts.NewToken(address, client)
	if err != nil {
		log.Fatalf("Failed to instantiate Sarcophagus Token contract: %v", err)
	}

	sarcophagusTokenContract = tokenContract
}

func EmbalmerSarcoBalance() *big.Int {
	balance, err := sarcophagusTokenContract.BalanceOf(&bind.CallOpts{}, embalmerAddress)

	if err != nil {
		log.Fatalf("Could not get sarcophagus balance. Please check your config PRIVATE_KEY value is correct.")
	}

	return balance
}

func InitEthClient(ethNode string) {
	cli, err := ethclient.Dial(ethNode)
	if err != nil {
		log.Fatalf("could not connect to Ethereum node. Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}

	client = cli
}

func isHex(str string) bool {
	return strings.HasPrefix(str, "0x")
}

func InitKeys(archPrivateKey string, embalmerPrivKey string) {
	if isHex(archPrivateKey) {
		archPrivateKey = archPrivateKey[2:]
	}
	if isHex(embalmerPrivKey) {
		embalmerPrivKey = embalmerPrivKey[2:]
	}
	ethPrivKey, err := crypto.HexToECDSA(archPrivateKey)
	if err != nil {
		log.Fatalf("could not load eth private key.  Please check the ETH_NODE value in the config file. Error: %v\n", err)
	}
	pub := ethPrivKey.Public()
	publicKey, _ := pub.(*ecdsa.PublicKey)
	archPublicKeyBytes = crypto.FromECDSAPub(publicKey)[1:]

	embalmerPrivateKey, _ = crypto.HexToECDSA(embalmerPrivKey)
	embalmerPubKey := embalmerPrivateKey.Public().(*ecdsa.PublicKey)
	embalmerAddress = crypto.PubkeyToAddress(*embalmerPubKey)
}

func approveCreateSarcophagusTransfer(session *contracts.TokenSession, approvalAmount *big.Int) {
	tx, err := session.Approve(
		sarcoAddress,
		approvalAmount,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error Approving Transaction: %v \n Config value ADD_TO_FREE_BOND has been reset to 0. You will need to reset this.", err)
	}

	log.Printf("Approval Transaction for %v Create Sarcophagus Tokens successful. Transaction ID: %v", approvalAmount, tx.Hash().Hex())
	log.Printf("Gas Used for Approval: %v", tx.Gas())
}

func CreateSarcophagus(recipientPrivateKey string) {
	/* Initialize recipient public key bytes */
	if isHex(recipientPrivateKey) {
		recipientPrivateKey = recipientPrivateKey[2:]
	}
	recipPrivKey, _ := crypto.HexToECDSA(recipientPrivateKey)
	pub := recipPrivKey.Public().(*ecdsa.PublicKey)
	recipientPublicKeyBytes := crypto.FromECDSAPub(pub)[1:]

	/* Sarc init values. */
	/* For resurrection time, using: 1/11/2021 -- https://www.unixtimestamp.com/ */
	resurrectionTime := big.NewInt(1610323200)
	storageFee := big.NewInt(2)
	diggingFee := big.NewInt(6)
	bounty := big.NewInt(6)

	/* Setup Asset Double Hash */
	/* Note: We may need to use solsha3 library to make these hashes, undetermined yet */
	/* https://github.com/miguelmota/go-solidity-sha3 */

	assetSingleHash := crypto.Keccak256([]byte{1, 2, 3, 4})
	assetDoubleHash := crypto.Keccak256(assetSingleHash)

	assetDoubleHashString := hex.EncodeToString(assetDoubleHash)
	var assetDoubleHashBytes [32]byte
	copy(assetDoubleHashBytes[:], assetDoubleHashString)

	log.Println("***CREATING SARCOPHAGUS***")

	sarcoTokenSession := NewSarcophagusTokenSession(context.Background())
	approvalAmount := new(big.Int).Add(new(big.Int).Add(bounty, diggingFee), storageFee)
	approveCreateSarcophagusTransfer(&sarcoTokenSession, approvalAmount)

	sarcoSession := NewSarcophagusSession(context.Background())
	tx, err := sarcoSession.CreateSarcophagus(
	"My Sarcophagus",
		archPublicKeyBytes,
		resurrectionTime,
		storageFee,
		diggingFee,
		bounty,
		assetDoubleHashBytes,
		recipientPublicKeyBytes,
	)

	if err != nil {
		log.Fatalf("Transaction reverted. Error creating Sarcophagus: %v", err)
	}

	log.Printf("Create Sarcophagus Successful. Transaction ID: %s", tx.Hash().Hex())
	log.Printf("Gas Used: %v", tx.Gas())
}