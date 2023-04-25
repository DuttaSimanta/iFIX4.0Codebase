package entities

import (
	"encoding/json"
	"io"
)

type ConditionEntity struct {
	Field   string      `json:"field"`
	Operand string      `json:"op"`
	Value   interface{} `json:"val"`
}

type ResultSetRequestEntity struct {
	Where          []ConditionEntity `json:"where"`
	Headers        []string          `json:"headers"`
	HeadersDisplay []string          `json:"headersdisplay"`
	Userid         int64             `json:"userid"`
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
