package entities

import (
	"encoding/json"
	"io"
)

//MstUserRoleActionEntity contains all required data fields
type MstUserRoleActionEntity struct {
	ID                  int64  `json:"id"`
	ClientID            int64  `json:"clientid"`
	MstorgnhirarchyID   int64  `json:"mstorgnhirarchyid"`
	RoleID              int64  `json:"roleid"`
	ActionID            int64  `json:"actionid"`
	Actionids           []int64  `json:"actionids"`
	RefuserID           int64  `json:"refuserid"`
	Deleteflag          int64  `json:"deleteflag"`
	Activeflag          int64  `json:"activeflag"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Rolename            string `json:"rolename"`
	Actionname          string `json:"actionname"`
	Username            string `json:"username"`
}

//FromJSON is used for convert data into JSON format
func (p *MstUserRoleActionEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MstUserRoleActionEntities is a entity with two fields
type MstUserRoleActionEntities struct {
	Total  int64                     `json:"total"`
	Values []MstUserRoleActionEntity `json:"values"`
}

//MstUserRoleActionEntityResponse is a response with all details
type MstUserRoleActionEntityResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details MstUserRoleActionEntities `json:"details"`
}

//MstUserRoleActionEntityResponseInt is a response with int
type MstUserRoleActionEntityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
