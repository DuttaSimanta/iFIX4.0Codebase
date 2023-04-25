package entities

import (
	"encoding/json"
	"io"
)


type MstClient struct {
	Id      	int64	`json:"id"`
	Code		string	`json:"code"`	
	Name    	string	`json:"name"`
	Description 	string  `json:"description"`
	OnboardDate   	string  `json:"onboarddate"`
	ClientAuditFlg  string  `json:"clientauditflg"`
}

type MstClientResponse struct {
	Status  	bool `json:"success"`
	Message 	string `json:"message"`
	Response 	MstClient `json:"response"`
}

type MstClientAllResponse struct {
	Status  	bool `json:"success"`
	Message 	string `json:"message"`
	Response 	[]MstClient `json:"response"`
}

func (p *MstClient) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
