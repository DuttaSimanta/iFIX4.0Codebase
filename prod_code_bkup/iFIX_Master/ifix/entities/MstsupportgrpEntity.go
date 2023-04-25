package entities

import (
	"encoding/json"
	"io"
)

type MstsupportgrpEntity struct {
	Id                  int64  `json:"id"`
	SupportgrpName      string `json:"supportgrpname"`
	Activeflg           int64  `json:"activeflg"`

	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Copyable            int64  `json:"copyable"`
	/*Roleid              int64  `json:"roleid"`
	Rolename            string `json:rolename`
	Groupid             int64  `json:"groupid"`
	Groupname           string `json:"groupname"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`*/
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
}
type MstsupportgrpbycopyableEntity struct{
	Id                  int64  `json:"id"`
	SupportgrpName      string `json:"supportgrpname"`
}
type MstsupportgrpEntities struct {
	Total  int64            `json:"total"`
	Values []MstsupportgrpEntity `json:"values"`
}

type MstsupportgrpbycopyableResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details []MstsupportgrpbycopyableEntity `json:"details"`
}

type MstsupportgrpResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details MstsupportgrpEntities `json:"details"`
}


type MstsupportgrpResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstsupportgrpEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
