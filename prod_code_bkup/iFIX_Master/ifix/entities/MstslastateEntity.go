package entities

import (
	"encoding/json"
	"io"
)

type MstslastateEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Statename           string `json:"statename"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}

type MstslastateEntities struct {
	Total  int64               `json:"total"`
	Values []MstslastateEntity `json:"values"`
}

type MstslastateResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Details MstslastateEntities `json:"details"`
}

type MstslastateResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstslastateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
