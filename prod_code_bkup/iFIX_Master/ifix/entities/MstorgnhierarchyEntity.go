package entities

import (
	"encoding/json"
	"io"
)

type MstorgnhierarchyFieldEntity struct {
	ID       int64 `json:"id"`
	ClientID int64 `json:"clientid"`
	Organizationname string `json:"organizationname"`
	Offset     int64  `json:"offset"`
	Limit      int64  `json:"limit"`
	Clientname string `json:"clientname"`
	Mfa              int64  `json:"mfa"`
	MfaName          string `json:"mfaname"`
	Notification     int64  `json:"notification"`
	NotificationName string `json:"notificationname"`
	Autoclose       int64  `json:"autoclose"`
	Autoclosefromdt string `json:"autoclosefromdt"`
	Autoclosetodt   string `json:"autoclosetodt"`
	AutocloseStatus string `json:"autoclosestatus"`
}
type MstorgnhierarchyFieldEntities struct {
	Total  int64                         `json:"total"`
	Values []MstorgnhierarchyFieldEntity `json:"values"`
}
type MstorgnhierarchyFieldEntityResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Details MstorgnhierarchyFieldEntities `json:"details"`
}

//MstorgnhierarchyEntity contains all required data fields
type MstorgnhierarchyEntity struct {
	ID                       int64  `json:"id"`
	ClientID                 int64  `json:"clientid"`
	ParentID                 int64  `json:"parentid"`
	MstorgnhierarchytypeID   int64  `json:"mstorgnhierarchytypeid"`
	Organizationname         string `json:"organizationname"`
	CityID                   int64  `json:"cityid"`
	CountryID                int64  `json:"countryid"`
	Code                     string `json:"code"`
	Location                 string `json:"location"`
	Pincode                  string `json:"pincode"`
	TimezoneID               int64  `json:"timezoneid"`
	ReporttimezoneID         int64  `json:"reporttimezoneid"`
	Offset                   int64  `json:"offset"`
	Limit                    int64  `json:"limit"`
	Clientname               string `json:"clientname"`
	Parentname               string `json:"parentname"`
	Cityname                 string `json:"cityname"`
	Countryname              string `json:"countryname"`
	Mstorgnhierarchytypename string `json:"mstorgnhierarchytypename"`
	Activationdate           string `json:"activationdate"`
	Timezonename             string `json:"timezonename"`
	Reporttimezonename       string `json:"reporttimezonename"`
	Timeformatid             int64  `json:"timeformatid"`
	Timeformat               string `json:"timeformat"`
	LogintypeID              int64  `json:"logintypeid"`
	LogintypeName            string `json:"logintypename"`
	Islocallogin             int64  `json:"islocallogin"`
	ReportTimeformatid       int64  `json:"reporttimeformatid"`
	ReportTimeformat         string `json:"reporttimeformat"`
	Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
	Mfa                      int64  `json:"mfa"`
	MfaName                  string `json:"mfaname"`
	Notification             int64  `json:"notification"`
	NotificationName         string `json:"notificationname"`
	Originalbgimage          string `json:"originalbgimage"`
	Uploadedbgimage          string `json:"uploadedbgimage"`
	Originallogoimage        string `json:"originallogoimage"`
	Uploadedlogoimage        string `json:"uploadedlogoimage"`
}

//MstorgnhierarchyEntityResp contains name fields
type MstorgnhierarchyEntityResp struct {
	ID               int64  `json:"id"`
	Organizationname string `json:"organizationname"`
	Timeformat       string `json:"timeformat"`
}
type LogintypeEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type MstorgnhierarchyResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Details []MstorgnhierarchyEntityResp `json:"details"`
}
type LogintypeEntityResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Details []LogintypeEntity `json:"details"`
}

//FromJSON is used for convert data into JSON format
func (p *MstorgnhierarchyEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
func (p *MstorgnhierarchyFieldEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

//MstorgnhierarchyEntities is a entity with two fields
type MstorgnhierarchyEntities struct {
	Total  int64                    `json:"total"`
	Values []MstorgnhierarchyEntity `json:"values"`
}

//MstorgnhierarchyEntityResponse is a response with all details
type MstorgnhierarchyEntityResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	Details MstorgnhierarchyEntities `json:"details"`
}

//MstorgnhierarchyEntityResponseInt is a response with int
type MstorgnhierarchyEntityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
