package models

import "math/big"

/* TODO: We probably dont need account index, in which case we dont need a struct and can just use a mapping */

type Sarcophagus struct {
	ResurrectionTime *big.Int
	AccountIndex int
}