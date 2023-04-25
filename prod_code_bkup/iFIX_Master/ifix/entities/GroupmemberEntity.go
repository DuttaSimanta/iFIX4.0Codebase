package entities

import (
	"encoding/json"
	"io"
)

type GroupmemberEntity struct {
	Id                  int64   `json:"id"`
	Clientid            int64   `json:"clientid"`
	Mstorgnhirarchyid   int64   `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyids  string   `json:"mstorgnhirarchyids"`
	Groupid             int64   `json:"groupid"`
	Refuserid           int64   `json:"refuserid"`
	Activeflg           int64   `json:"activeflg"`
	Offset              int64   `json:"offset"`
	Limit               int64   `json:"limit"`
	Clientname          string  `json:"clientname"`
	Loginname           string  `json:"loginname"`
	Email           	string  `json:"email"`
	Mstorgnhirarchyname string  `json:"mstorgnhirarchyname"`
	Supportgroupname    string  `json:"supportgroupname"`
	Username            string  `json:"username"`
	Type                string  `json:"type"`
	Userids             []int64 `json:"userids"`
	ToMstorgnhirarchyid []int64 `json:"tomstorgnhirarchyid"`
	ToClientid          int64   `json:"toclientid"`
}

type GroupmemberEntities struct {
	Total  int64               `json:"total"`
	Values []GroupmemberEntity `json:"values"`
}

type GroupmemberResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Details GroupmemberEntities `json:"details"`
}

type GroupmemberResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *GroupmemberEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
