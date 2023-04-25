package entities

import (
	"encoding/json"
	"io"
)

type ModuleEntity struct{
	Id        		   int64    `json:"id"`
	Modulename         string   `json:"modulename"`
	Moduledescription  string   `json:"moduledescription"`
}
type ModuleEntities struct{
	Total int64 `json:"total"`
	Values []ModuleEntity  `json:"values"`
}
type ModuleResponse struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	ModuleEntities `json:"details"`
}
type ModuleResponseInt struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	int64 `json:"details"`
}
func (w *ModuleEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}