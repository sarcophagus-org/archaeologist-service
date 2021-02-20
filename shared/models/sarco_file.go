// SarcoFile -- the body of the data sent from the embalmer to the file upload endpoint
// is unmarshalled into this struct for use when uploading file to arweave

package models

type SarcoFile struct {
	FileType  string `json:"fileType"`
	FileBytes string `json:"fileBytes"`
}