package entities

import (
	"encoding/json"
	"io"
)

//MstClientEntity contains all required data fields
type MstClientEntity struct {
	ID              int64  `json:"id"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Keyperson       string `json:"keyperson"`
	Keyemail        string `json:"keyemail"`
	Keymobile       string `json:"keymobile"`
	Baseflag        string `json:"baseflag"`
	Spocname        string `json:"spocname"`
	Spocemail       string `json:"spocemail"`
	Spocnumber      string `json:"spocnumber"`
	Clientauditflag int64  `json:"clientauditflag"`
	Offset          int64  `json:"offset"`
	Limit           int64  `json:"limit"`
}

type AllMstClientEntity struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	OrgnID int64  `json:"orgnid"`
}


type AllMstClientEntityResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details []AllMstClientEntity `json:"details"`
}
//FromJSON is used for convert data into JSON format
func (p *MstClientEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MstClientEntities is a entity with two fields
type MstClientEntities struct {
	Total  int64             `json:"total"`
	Values []MstClientEntity `json:"values"`
}

//MstClientEntityResponse is a response with all details
type MstClientEntityResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Details MstClientEntities `json:"details"`
}

//MstClientEntityResponseInt is a response with int
type MstClientEntityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
