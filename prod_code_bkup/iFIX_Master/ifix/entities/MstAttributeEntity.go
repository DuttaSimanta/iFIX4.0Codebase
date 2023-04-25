package entities

import (
	"encoding/json"
	"io"
)

type MstAttributeEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Adfsattribute       string `json:"adfsattribute"`
	Activeflg           int64  `json:"activeglg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
}

type Attributes struct {
	Adfsattribute string `json:"key"`
}
type AttributesResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details []Attributes `json:"details"`
}
type MstAttributeEntities struct {
	Total  int64                `json:"total"`
	Values []MstAttributeEntity `json:"values"`
}

type MstAttributeResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details MstAttributeEntities `json:"details"`
}

type MstAttributeResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstAttributeEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
