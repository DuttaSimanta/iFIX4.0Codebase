package entities

import (
	"encoding/json"
	"io"
)

//MapClientUserRoleUserEntity contains all required data fields
type MapClientUserRoleUserEntity struct {
	ID                  int64  `json:"id"`
	ClientID            int64  `json:"clientid"`
	MstorgnhirarchyID   int64  `json:"mstorgnhirarchyid"`
	RoleID              int64  `json:"roleid"`
	Refuserid              int64  `json:"refuserid"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Rolename            string `json:"rolename"`
	Username            string `json:"Username"`
}

type MapUserRoleEntityResp struct {
	ID                  int64  `json:"id"`
	RoleID              int64  `json:"roleid"`
	Refuserid              int64  `json:"refuserid"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Rolename            string `json:"rolename"`
	Username            string `json:"Username"`
}

//FromJSON is used for convert data into JSON format
func (p *MapClientUserRoleUserEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MapClientUserRoleUserEntities is a entity with two fields
type MapClientUserRoleUserEntities struct {
	Total  int64                         `json:"total"`
	Values []MapUserRoleEntityResp `json:"values"`
}

//MapClientUserRoleUserEntityResponse is a response with all details
type MapClientUserRoleUserEntityResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Details MapClientUserRoleUserEntities `json:"details"`
}

//MapClientUserRoleUserEntityResponseInt is a response with int
type MapClientUserRoleUserEntityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
