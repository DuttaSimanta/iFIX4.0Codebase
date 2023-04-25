package entities

import (
	"encoding/json"
	"io"
)

type StateEntity struct{
	Id                  int64    `json:"id"`
	Clientid            int64    `json:"clientid"`
	Mstorgnhirarchyid   int64    `json:"mstorgnhirarchyid"`
	Statetypeid   		int64    `json:"statetypeid"`
	Statename   		string    `json:"statename"`
	Description   		string    `json:"description"`
	Audittransactionid   int64    `json:"audittransactionid"`
}
func (w *StateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}