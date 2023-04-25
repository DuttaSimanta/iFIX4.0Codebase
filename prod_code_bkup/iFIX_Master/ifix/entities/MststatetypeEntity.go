package entities

import (
	"encoding/json"
	"io"
)

type MststatetypeEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Statetypename       string `json:"statetypename"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}

type MststatetypeEntities struct {
	Total  int64                `json:"total"`
	Values []MststatetypeEntity `json:"values"`
}

type MststatetypeResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details MststatetypeEntities `json:"details"`
}

type MststatetypeResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MststatetypeEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
