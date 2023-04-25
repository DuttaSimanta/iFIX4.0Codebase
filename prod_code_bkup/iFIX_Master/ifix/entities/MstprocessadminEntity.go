package entities

import (
	"encoding/json"
	"io"
)

type MstprocessadminEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Processid           int64  `json:"processid"`
	Refuserid           int64  `json:"refuserid"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Processname         string `json:"processname"`
	Username            string `json:"username"`
}

type MstprocessadminEntities struct {
	Total  int64                   `json:"total"`
	Values []MstprocessadminEntity `json:"values"`
}

type MstprocessadminResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details MstprocessadminEntities `json:"details"`
}

type MstprocessadminResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstprocessadminEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
