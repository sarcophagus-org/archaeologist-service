package embalmer

import (
	"crypto/ecdsa"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/contracts"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/ethereum"
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"time"
)

func InitEmbalmer(embalmer *Embalmer, config *EmbalmerConfig, resurrectionTime int64) {
	embalmer.Client = ethereum.InitEthClient(config.ETH_NODE)
	embalmer.EmbalmerPrivateKey, _ = utility.PrivateKeyHexToECDSA(config.EMBALMER_PRIVATE_KEY)
	archPrivateKey, _ := utility.PrivateKeyHexToECDSA(config.ARCH_PRIVATE_KEY)
	embalmer.ArchAddress = utility.PrivateKeyToAddress(archPrivateKey)
	embalmerPubKey := embalmer.EmbalmerPrivateKey.Public().(*ecdsa.PublicKey)
	embalmer.EmbalmerAddress = crypto.PubkeyToAddress(*embalmerPubKey)
	embalmer.SarcoAddress = ethereum.SarcoAddress(config.CONTRACT_ADDRESS, embalmer.Client)
	embalmer.SarcophagusContract, _ = contracts.NewSarcophagus(embalmer.SarcoAddress, embalmer.Client)
	embalmer.SarcophagusTokenContract, _ = contracts.NewToken(common.HexToAddress(config.TOKEN_ADDRESS), embalmer.Client )
	embalmer.ResurrectionTime = time.Now().Unix() + resurrectionTime
	embalmer.StorageFee = config.STORAGE_FEE
	embalmer.DiggingFee = config.DIGGING_FEE
	embalmer.Bounty = config.BOUNTY
}