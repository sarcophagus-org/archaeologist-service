// Not used by the service, for testing purposes only
package embalmer

import (
	"github.com/decent-labs/airfoil-sarcophagus-archaeologist-service/shared/utility"
	"math/rand"
)

func DoubleHashBytesFromSeed(seed int64, size int) ([]byte, [32]byte) {
	/* Generate random bytes to use as payload for each sarco */
	fileBytes := make([]byte, size)

	rand.Seed(seed)
	rand.Read(fileBytes)
	return fileBytes, utility.FileBytesToDoubleHashBytes(fileBytes)
}