package entities

import (
	"encoding/json"
	"io"
)

type MstldapEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Clientname          string `json:"clienname"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	ServerName          string `json:"servername"`
	ServerUrl           string `json:"serverurl"`
	Binddn              string `json:"binddn"`
	Basedn              string `json:"basedn"`
	Password            string `json:"password"`
	Filterdn            string `json:"filterdn"`
	Ori_Certificate     string `json:"ori_certificate"`
	Chn_Certificate     string `json:"chn_certificate"`
	Activeflg           int64  `json:"activeglg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Tablename           string `json:"tablename"`
}
type MstldapEntities struct {
	Total  int64           `json:"total"`
	Values []MstldapEntity `json:"values"`
}

type MstldapResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Details MstldapEntities `json:"details"`
}
type MstldapfieldResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

type MstldapResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstldapEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
