package entities

import (
	"encoding/json"
	"io"
)

type MstclientslaEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Slaname             string `json:"slaname"`
	Slatimereset        int64  `json:"slatimereset"`
	Slaupgradereset     int64  `json:"slaupgradereset"`
	Sladowngradereset   int64  `json:"sladowngradereset"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}

type Mstslaname struct {
	Id      int64  `json:"id"`
	Slaname string `json:"slaname"`
}

type MstclientslaEntities struct {
	Total  int64                `json:"total"`
	Values []MstclientslaEntity `json:"values"`
}

type MstslanameResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details []Mstslaname `json:"details"`
}

type MstclientslaResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details MstclientslaEntities `json:"details"`
}

type MstclientslaResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstclientslaEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
