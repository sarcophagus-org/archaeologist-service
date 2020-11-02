package models

type ResponseToEmbalmer struct {
	SignedArweaveTxId     string    `json:"ArweaveTxId"`
	AssetDoubleHash [32]byte `json:"AssetDoubleHash"`
	SigV            int    `json:"sigV"`
	SigR            string `json:"sigR"`
	SigS            string `json:"sigS"`
}
