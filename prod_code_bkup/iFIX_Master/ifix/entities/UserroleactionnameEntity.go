package entities

import (
	"encoding/json"
	"io"
)

type UserroleactionnameEntity struct {
	ID                int64 `json:"id"`
	Clientid          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
	UserID            int64 `json:"userid"`
}

type UserroleactionnameEntityResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Values  []int64 `json:"values"`
}

func (w *UserroleactionnameEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
