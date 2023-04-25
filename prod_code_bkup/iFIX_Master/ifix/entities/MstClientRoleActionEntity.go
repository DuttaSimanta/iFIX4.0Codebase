package entities

import (
	"encoding/json"
	"io"
)

//MstClientRoleActionEntity contains all required data fields
type MstClientRoleActionEntity struct {
	ID                  int64   `json:"id"`
	ClientID            int64   `json:"clientid"`
	MstorgnhirarchyID   int64   `json:"mstorgnhirarchyid"`
	RoleID              int64   `json:"roleid"`
	ActionID            int64   `json:"actionid"`
	ActionIDs           []int64 `json:"actionids"`
	Deleteflag          int64   `json:"deleteflag"`
	Activeflag          int64   `json:"activeflag"`
	Offset              int64   `json:"offset"`
	Limit               int64   `json:"limit"`
	Clientname          string  `json:"clientname"`
	Mstorgnhirarchyname string  `json:"mstorgnhirarchyname"`
	Rolename            string  `json:"rolename"`
	Actionname          string  `json:"actionname"`
}
type RoleActionEntityResp struct {
	Clientname string `json:"clientname"`
	Orgname    int64  `json:"orgname"`
	Username   int64  `json:"userid"`
	Roleid     int64  `json:"Roleid"`
	Actionid   int64  `json:"actionid"`
}

//FromJSON is used for convert data into JSON format
func (p *MstClientRoleActionEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MstClientRoleActionEntities is a entity with two fields
type MstClientRoleActionEntities struct {
	Total  int64                       `json:"total"`
	Values []MstClientRoleActionEntity `json:"values"`
}

//MstClientRoleActionEntityResponse is a response with all details
type MstClientRoleActionEntityResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details MstClientRoleActionEntities `json:"details"`
}

//MstClientRoleActionEntityResponseInt is a response with int
type MstClientRoleActionEntityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
