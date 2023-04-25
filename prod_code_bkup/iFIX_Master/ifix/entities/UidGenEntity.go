package entities

import (
	"encoding/json"
	"io"
)

type UidGenEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Difftypeseq         int64  `json:"difftypeid"`
	Difftypename        string `json:"difftypename"`
	Code                string `json:"code"`
	Uid                 int64  `json:"uid"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}

type UidGenEntities struct {
	Total  int64          `json:"total"`
	Values []UidGenEntity `json:"values"`
}

type UidGenResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Details UidGenEntities `json:"details"`
}

type UidGenResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *UidGenEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

type MstorgnhierarchywithOrgtypeEntity struct {
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
}

//MstorgnhierarchyEntityResp contains name fields
type MstorgnhierarchywithOrgtypeEntityResp struct {
	ID                     int64  `json:"id"`
	Organizationname       string `json:"organizationname"`
	MstorgnhierarchytypeID int64  `json:"mstorgnhierarchytypeid"`
	Timeformat             string `json:"timeformat"`
}

type MstorgnhierarchyResponsewithorgtype struct {
	Success bool                                    `json:"success"`
	Message string                                  `json:"message"`
	Details []MstorgnhierarchywithOrgtypeEntityResp `json:"details"`
}

func (p *MstorgnhierarchywithOrgtypeEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
