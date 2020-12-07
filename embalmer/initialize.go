package embalmer

import (
	"crypto/ecdsa"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"time"
)

func InitEmbalmer(embalmer *Embalmer, config *EmbalmerConfig, resurrectionTime int64) {
	embalmer.Client, _ = ethereum.InitEthClient(config.ETH_NODE)
	embalmer.EmbalmerPrivateKey, _ = utility.PrivateKeyHexToECDSA(config.EMBALMER_PRIVATE_KEY)
	archPrivateKey, _ := utility.PrivateKeyHexToECDSA(config.ARCH_PRIVATE_KEY)
	embalmer.ArchAddress = utility.PrivateKeyToAddress(archPrivateKey)
	embalmerPubKey := embalmer.EmbalmerPrivateKey.Public().(*ecdsa.PublicKey)
	embalmer.EmbalmerAddress = crypto.PubkeyToAddress(*embalmerPubKey)
	embalmer.SarcoAddress, _ = ethereum.SarcoAddress(config.CONTRACT_ADDRESS, embalmer.Client)
	embalmer.SarcophagusContract, _ = contracts.NewSarcophagus(embalmer.SarcoAddress, embalmer.Client)
	embalmer.SarcophagusTokenContract, _ = contracts.NewToken(common.HexToAddress(config.TOKEN_ADDRESS), embalmer.Client )
	embalmer.ResurrectionTime = big.NewInt(time.Now().Unix() + resurrectionTime)
	embalmer.StorageFee, _ = new(big.Int).SetString(config.STORAGE_FEE, 10)
	embalmer.DiggingFee, _ = new(big.Int).SetString(config.DIGGING_FEE, 10)
	embalmer.Bounty, _ = new(big.Int).SetString(config.BOUNTY, 10)
}