package ethereum
//
//import (
//	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"log"
//)
//
//func RegisterOrUpdateArchaeologist(config *models.Config) {
//	archaeologist, err := sarcophagusContract.Archaeologists(&bind.CallOpts{}, archAddress)
//	if err != nil {
//		log.Fatalf("Call to Archaeologists in Sarcophagus Contract failed: %v", err)
//	}
//
//	if (archaeologist.Exists) {
//		// Update arch
//	} else {
//		sarcophagusContract.RegisterArchaeologist(
//			&bind.TransactOpts{},
//			archPublicKey,
//			config.FILE_PORT,
//			archAddress,
//			big.NewInt(config.FEE_PER_BYTE),
//			config.MIN_BOUNTY,
//			config.MIN_DIGGING_FEE,
//			config.MAX_RESURRECTION_TIME,
//			freeBond,
//		)
//	}
//
//
//}