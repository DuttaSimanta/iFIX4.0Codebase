 package entities

import (
	"encoding/json"
	"io"
)

type MapldapgrouproleEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Logintypeid   		int64  `json:"logintypeid"`
	Roleid              int64  `json:"roleid"`
	Rolename            string `json:"rolename"`
	Groupid             int64  `json:"groupid"`
	Groupname           string `json:"groupname"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Logintype          string `json:"logintype"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`

}

type MapldapgrouproleEntities struct {
	Total  int64            `json:"total"`
	Values []MapldapgrouproleEntity `json:"values"`
}

type MapldapgrouproleResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details MapldapgrouproleEntities `json:"details"`
}

type MapldapgrouproleResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MapldapgrouproleEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
