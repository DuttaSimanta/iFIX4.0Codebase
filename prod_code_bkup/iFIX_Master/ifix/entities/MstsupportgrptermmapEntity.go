package entities

import (
	"encoding/json"
	"io"
)

type MstsupportgrptermmapEntity struct {
	Id                int64 `json:"id"`
	Clientid          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
	//Statename           string `json:"statename"`
	Termid              []int64 `json:termid`
	Grpid               int64   `json:grpid`
	Audittransactionid  int64   `json:audittransactionid`
	Activeflg           int64   `json:"activeflg"`
	Offset              int64   `json:"offset"`
	Limit               int64   `json:"limit"`
	Clientname          string  `json:"clientname"`
	Mstorgnhirarchyname string  `json:"mstorgnhirarchyname"`
	Termname            string  `json:termname`
	Grpname             string  `json:grpname`
	Readpermission      string  `json:"readpermission"`
	Writepermission     string  `json:"writepermission"`
}

type MstsupportgrptermmapEntities struct {
	Total  int64                        `json:"total"`
	Values []MstsupportgrptermmapEntity `json:"values"`
}

type MstsupportgrptermmapResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Details MstsupportgrptermmapEntities `json:"details"`
}

type MstsupportgrptermmapResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstsupportgrptermmapEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
