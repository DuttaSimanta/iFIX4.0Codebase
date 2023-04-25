package entities

import (
	"encoding/json"
	"io"
)

type CatconitionEntity struct {
	Seq   interface{} `json:"seq"`
	Value interface{} `json:"val"`
}
type ConditionEntity struct {
	Field   string      `json:"field"`
	Operand string      `json:"op"`
	Value   interface{} `json:"val"`
}

type ResultSetRequestEntity struct {
	Where          []ConditionEntity   `json:"where"`
	Cat            []CatconitionEntity `json:"cat"`
	Headers        []string            `json:"headers"`
	HeadersDisplay []string            `json:"headersdisplay"`
	Userid         int64               `json:"userid"`
}

type ResultSetJsonForExcelEntity struct {
	RequestResultsetData []map[string]interface{} `json:"result"`
	Total                int64                    `json:"total"`
}

type JsonToExcelResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details ResultSetJsonForExcelEntity `json:"details"`
}
type APIResponseDownload struct {
	Status           bool   `json:"success"`
	Message          string `json:"message"`
	UploadedFileName string `json:"uploadedfilename"`
	OriginalFileName string `json:"originalfilename"`
}

/*type JsonToExcelResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Details ResultSetEntity `json:"details"`
}
*/
func (w *ResultSetRequestEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

type FileuploadEntity struct {
	Id                 int64  `json:"id"`
	Clientid           int64  `json:"clientid"`
	Mstorgnhirarchyid  int64  `json:"mstorgnhirarchyid"`
	Credentialtype     string `json:"credentialtype"`
	Credentialaccount  string `json:"credentialaccount"`
	Credentialpassword string `json:"credentialpassword"`
	Credentialkey      string `json:"credentialkey"`
	Activeflg          int64  `json:"activeflg"`
	Originalfile       string `json:"originalfile"`
	Filename           string `json:"filename"`
	Path               string `json:"path"`
}
type DownloadlistEntity struct {
	Refuserid        int64  `json:"refuserid"`
	OriginalFileName string `json:"originalfilename"`
	UploadedFileName string `json:"uploadedfilename"`
}
type ReportDownloadResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
type UtilityEntity struct {
	Id                int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	Date              int64  `json:"date"`
	Timediff          int64  `json:"timediff"`
	Reporttimediff    int64  `json:"reporttimediff"`
	Reporttimeformat  string `json:"reporttimeformat"`
	Timeformat        string `json:"timeformat"`
}
type ClientOrgEntity struct {
	Clientid          interface{} `json:"clientid"`
	Mstorgnhirarchyid interface{} `json:"mstorgnhirarchyid"`
}
