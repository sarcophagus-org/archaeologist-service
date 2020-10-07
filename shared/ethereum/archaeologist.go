package ethereum

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"
)

func RegisterOrUpdateArchaeologist(config *models.Config) {
	archaeologist, err := sarcophagusContract.Archaeologists(&bind.CallOpts{}, archAddress)
	if err != nil {
		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed: %v", err)
	}

	nonce := GetNonce()
	gasPrice := GetSuggestedGasPrice()
	auth := bind.NewKeyedTransactor(archPrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	if archaeologist.Exists {
		// TODO: Update arch
		log.Println("Arch exists!")
	} else {
		tx, err := sarcophagusContract.RegisterArchaeologist(
			auth,
			archPublicKeyBytes,
			config.FILE_PORT,
			archAddress,
			big.NewInt(config.FEE_PER_BYTE),
			big.NewInt(config.MIN_BOUNTY),
			big.NewInt(config.MIN_DIGGING_FEE),
			big.NewInt(config.MAX_RESURRECTION_TIME),
			freeBond,
		)

		if err != nil {
			log.Fatalf("Transaction reverted. Error registering Archaeologist: %v", err)
		}

		log.Println("tx sent: %s", tx.Hash().Hex())
	}
}