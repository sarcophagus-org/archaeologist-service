// ResponseToEmbalmer model is used to store the data returned to the embalmer
// after successful completion of the arweave file uploading process

package models

type ResponseToEmbalmer struct {
	NewPublicKey    []byte   `json:NewPublicKey`
	AssetId         string   `json:"AssetId"`
	AssetDoubleHash [32]byte `json:"AssetDoubleHash"`
	V               uint8    `json:"V"`
	R               [32]byte `json:"R"`
	S               [32]byte `json:"S"`
}
