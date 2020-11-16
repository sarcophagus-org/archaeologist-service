package models

import "math/big"

/* TODO: Determine if we will be adding any other fields here. If not, dont need a struct */

type Sarcophagus struct {
	ResurrectionTime *big.Int
}