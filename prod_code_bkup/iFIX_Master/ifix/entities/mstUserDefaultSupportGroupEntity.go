package entities

import (
	"encoding/json"
	"io"
)

type MstUserDefaultSupportGroupEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Refuserid           int64  `json:"refuserid"`
	Groupid             int64  `json:"groupid"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Groupname           string `json:"groupname"`
	Refusername            string `json:"refusername"`
}

type MstUserDefaultSupportGroupEntities struct {
	Total  int64                              `json:"total"`
	Values []MstUserDefaultSupportGroupEntity `json:"values"`
}

type MstUserDefaultSupportGroupResponse struct {
	Success bool                               `json:"success"`
	Message string                             `json:"message"`
	Details MstUserDefaultSupportGroupEntities `json:"details"`
}

type MstUserDefaultSupportGroupResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstUserDefaultSupportGroupEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
