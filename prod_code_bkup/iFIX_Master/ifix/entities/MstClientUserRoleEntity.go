package entities

import (
	"encoding/json"
	"io"
)

//MstClientUserRoleEntity contains all required data fields
type MstClientUserRoleEntity struct {
	ID                int64  `json:"id"`
	ClientID          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	Rolename          string `json:"rolename"`
	Roledesc          string `json:"roledesc"`
	Adminflag         int64  `json:"adminflag"`
	UserID            int64  `json:"userid"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
}

//MstClientRoleEntityResp contains id and role name
type MstClientRoleEntityResp struct {
	ID                int64  `json:"id"`
	Rolename          string `json:"rolename"`
	Roledesc          string `json:"roledesc"`
	Issuperadmin          int `json:"issuperadmin"`
}


//FromJSON is used for convert data into JSON format
func (p *MstClientUserRoleEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MstClientUserRoleEntities is a entity with two fields
type MstClientUserRoleEntities struct {
	Total  int64                     `json:"total"`
	Values []MstClientUserRoleEntity `json:"values"`
}

//MstClientUserRoleResponse is a response with all details
type MstClientUserRoleResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details MstClientUserRoleEntities `json:"details"`
}
//MstClientUserRoleResponse is a response with all details
type MstClientRoleResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details []MstClientRoleEntityResp `json:"details"`
}

//MstClientUserRoleResponseInt is a response with int
type MstClientUserRoleResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

//curl -v localhost:8082/createrole -d '{"clientid":1,"mstorgnhirarchyid":2,"rolename":"AAA","roledesc":"BBB","adminflag":1,"userid":1}'
//{"success":true,"message":"","details":3}
