// Sarco-- holds sarcophagus information to be stored

package models

import "math/big"

type Sarco struct {
	ResurrectionTime *big.Int
	AccountIndex     int
	Updated          bool
	UnwrapAttempts	 int
}
