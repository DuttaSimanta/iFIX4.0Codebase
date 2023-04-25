package entities

import (
	"encoding/json"
	"io"
)

//MstActionEntity contains all required data fields
type MstActionEntity struct {
	ID         int64  `json:"id"`
	Actionname string `json:"actionname"`
}

//FromJSON is used for convert data into JSON format
func (p *MstActionEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MstActionEntityResponse is a response with all details
type MstActionEntityResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Details []MstActionEntity `json:"details"`
}
