package entities

import (
	"encoding/json"
	"io"
)

type TicketCustomerEntity struct {
	//Id int64 `json:"id"`
	// Modulename         string   `json:"modulename"`
	// Moduledescription  string   `json:"moduledescription"`
	Clientid            int64  `json:"clientid"`
	Refuserid              int64  `json:"refuserid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}
type TicketCustomerEntities struct {
	Total  int64                  `json:"total"`
	Values []TicketCustomerEntity `json:"values"`
}
type TicketCustomerResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details TicketCustomerEntities `json:"details"`
}

/*type ModuleResponseInt struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	int64 `json:"details"`
}*/
func (w *TicketCustomerEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
