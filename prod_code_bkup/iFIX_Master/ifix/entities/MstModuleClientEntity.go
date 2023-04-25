package entities

import (
	"encoding/json"
	"io"
)

//MstModuleClientEntity contains all required data fields
type MstModuleClientEntity struct {
	ID                  int64  `json:"id"`
	ClientID            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	ModuleID            int64  `json:"moduleid"`
	Deleteflag          int64  `json:"deleteflag"`
	Activeflag          int64  `json:"activeflag"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Modulename          string `json:"modulename"`
	Fromdate            string `json:"fromdate"`
}
type MstModuleClientEntityResp struct {
	ID                  int64  `json:"id"`
	Clientname          string `json:"clientname"`
	ModuleID            int64  `json:"moduleid"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Modulename          string `json:"modulename"`
	Fromdate            string `json:"fromdate"`
}
//MstModuleByClientEntity contains data for client wise
type MstModuleByClientEntity struct {
	ID                  int64  `json:"id"`
	Modulename          string `json:"modulename"`
}
//MstModuleByClientEntityResponse is a response with all details
type MstModuleByClientEntityResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details []MstModuleByClientEntity `json:"details"`
}


//FromJSON is used for convert data into JSON format
func (p *MstModuleClientEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MstModuleClientEntities is a entity with two fields
type MstModuleClientEntities struct {
	Total  int64                   `json:"total"`
	Values []MstModuleClientEntityResp `json:"values"`
}

//MstModuleClientEntityResponse is a response with all details
type MstModuleClientEntityResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details MstModuleClientEntities `json:"details"`
}

//MstModuleClientEntityResponseInt is a response with int
type MstModuleClientEntityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
