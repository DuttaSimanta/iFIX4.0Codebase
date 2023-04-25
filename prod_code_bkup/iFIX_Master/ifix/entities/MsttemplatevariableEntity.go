package entities

import (
	"encoding/json"
	"io"
)

type MsttemplatevariableEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Templatename        string `json:"templatename"`
	Tableid             int64  `json:"tableid"`
	Fieldid             int64  `json:"fieldid"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Tablename           string `json:"tablename"`
	Fieldname           string `json:"fieldname"`
}

type MsttemplatevariableEntities struct {
	Total  int64                       `json:"total"`
	Values []MsttemplatevariableEntity `json:"values"`
}

type MsttemplatevariableResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details MsttemplatevariableEntities `json:"details"`
}

type MsttemplatevariableResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MsttemplatevariableEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
