package entities

import (
	"encoding/json"
	"io"
)

type MapUserRolePropertyEntity struct {
	Id                  int64   `json:"id"`
	Clientid            int64   `json:"clientid"`
	Mstorgnhirarchyid   int64   `json:"mstorgnhirarchyid"`
	Roleid              []int64 `json:"roleid"`
	Propertyid          int64   `json:"propertyid"`
	Activeflg           int64   `json:"activeflg"`
	Offset              int64   `json:"offset"`
	Limit               int64   `json:"limit"`
	Mstorgnhirarchyname string  `json:"mstorgnhirarchyname"`
	Rolename            string  `json:"rolename"`
	Propertyname        string  `json:"propertyname"`
}

type GetUserPropertyNameEntity struct {
	Id           int64  `json:"id"`
	Propertyname string `json:"propertyname"`
}

type MapUserRolePropertyEntities struct {
	Total  int64                       `json:"total"`
	Values []MapUserRolePropertyEntity `json:"values"`
}

type MapUserRolePropertyResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details MapUserRolePropertyEntities `json:"details"`
}

type MapUserRolePropertyResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

type GetUserPropertyNameResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details []GetUserPropertyNameEntity `json:"details"`
}

func (w *MapUserRolePropertyEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
