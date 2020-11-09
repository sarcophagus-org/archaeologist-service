package archaeologist

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/models"
	"log"
	"time"
)

// TODO: Pass in doublehash to unwrap
func scheduleUnwrap(resurrectionTime int64, arch *models.Archaeologist) {
	timeToUnwrap := time.Until(time.Unix(resurrectionTime, 0))
	log.Println("Unwrap scheduled in:", timeToUnwrap)

	time.AfterFunc(timeToUnwrap, func() {
		//	bytes32 assetDoubleHash,
		//	bytes memory singleHash,
		//	bytes32 privateKey
		// arch.SarcoSession.UnwrapSarcophagus()
	})
}