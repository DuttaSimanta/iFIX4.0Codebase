package entities

import (
	"encoding/json"
	"io"
)

type MsturlkeyEntity struct {
	Id         int64  `json:"id"`
	Urlkeyname string `json:"Urlkeyname"`
}

type MsturlkeyInputEntity struct {
	Clientid          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
}

type MsturlkeyEntities struct {
	Values []MsturlkeyEntity `json:"values"`
}

type MsturlkeyResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Details MsturlkeyEntities `json:"details"`
}

func (w *MsturlkeyEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *MsturlkeyInputEntity) FromInputJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
