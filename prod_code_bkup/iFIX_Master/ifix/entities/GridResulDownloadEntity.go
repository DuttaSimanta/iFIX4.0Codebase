package entities

import (
	"encoding/json"
	"io"
)

type ConditionForGridEntity struct {
	Field   string      `json:"field"`
	Operand string      `json:"op"`
	Value   interface{} `json:"val"`
}
type OrderSeq struct {
	Field string `json:"field"`
	Dir   string `json:"dir"`
}
type ResultGridRequestEntity struct {
	Clientid                int64                    `json:"clientid"`
	Mstorgnhirarchyid       int64                    `json:"mstorgnhirarchyid"`
	Searchmstorgnhirarchyid string                   `json:"searchmstorgnhirarchyid"`
	RecordDiffid            int64                    `json:"recorddiffid"`
	RecordDiffSeq           int64                    `json:"recorddiffidseq"`
	Menuid                  int64                    `json:"menuid"`
	QueryType               int64                    `json:"querytype"`
	SupportGrpid            string                   `json:"supportgrpid"`
	Where                   []ConditionForGridEntity `json:"where"`
	Order                   []OrderSeq               `json:"order"`
	Headers                 []string                 `json:"headers"`
	HeadersDisplay          []string                 `json:"headersdisplay"`
	Iscatalog               int64                    `json:"iscatalog"`
	//Offset            int64                    `json:"offset"`
	//Limit             int64                    `json:"limit"`
	Userid int64 `json:"userid"`
}

type ResultGridJsonForExcelEntity struct {
	Menuid                int64                    `json:"menuid"`
	QueryType             int64                    `json:"querytype"`
	RequestResultGridData []map[string]interface{} `json:"result"`
	Total                 int64                    `json:"total"`
}

type JsonToExcelGridResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Details ResultGridJsonForExcelEntity `json:"details"`
}
type APIResponseGridDownload struct {
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
func (w *ResultGridRequestEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
