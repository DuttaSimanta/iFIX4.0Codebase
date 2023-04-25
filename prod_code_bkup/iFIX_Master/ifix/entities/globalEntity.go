package entities

import (
	"encoding/json"
	"io"
)

type ZoneEntity struct{
	Id        		int64    `json:"id"`
	Zonename       	   string   `json:"zonename"`
}

type ZoneEntityResp struct{
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	[]ZoneEntity `json:"details"`
}
type GlobalEntityResp struct{
	Success  	bool `json:"success"`
}
func (p *ZoneEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}