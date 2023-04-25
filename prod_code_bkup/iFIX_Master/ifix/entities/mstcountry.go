package entities

import (
	"encoding/json"
	"io"
)


type MstCountry struct {
	Id      	int64	`json:"id"`
	CountryCode	string	`json:"countrycode"`	
	CountryName    	string	`json:"countryname"`
}

type MstCountryResponse struct {
	Status  	bool `json:"success"`
	Message 	string `json:"message"`
	Response 	MstCountry `json:"response"`
}

type MstCountryAllResponse struct {
	Status  	bool `json:"success"`
	Message 	string `json:"message"`
	Response 	[]MstCountry `json:"response"`
}

func (p *MstCountry) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
