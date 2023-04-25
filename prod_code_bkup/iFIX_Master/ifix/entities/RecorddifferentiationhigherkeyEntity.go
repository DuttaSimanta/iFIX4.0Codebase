package entities

import (
	"encoding/json"
	"io"
)

//RecorddifferentiationhigherkeyEntity defines all data fields
type RecorddifferentiationhigherkeyEntity struct {
	ID                       int64  `json:"id"`
	Clientid                 int64  `json:"clientid"`
	Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
	Parentrecorddifftypeid   int64  `json:"parentrecorddifftypeid"` // ticket type
	Parentrecorddiffid       int64  `json:"parentrecorddiffid"`     // incident
	Childrecorddifftypeid    int64  `json:"childrecorddifftypeid"`  // category label
	Childrecorddiffid        int64  `json:"childrecorddiffid"`      // category value
	Parentid                 int64  `json:"parentid"`
	Name                     string `json:"name"`
	Seqno                    int64  `json:"seqno"`
	Activeflg                int64  `json:"activeflg"`
	Offset                   int64  `json:"offset"`
	Limit                    int64  `json:"limit"`
	Clientname               string `json:"clientname"`
	Mstorgnhirarchyname      string `json:"mstorgnhirarchyname"`
	Parentrecorddifftypename string `json:"parentrecorddifftypename"`
	Parentrecorddiffname     string `json:"parentrecorddiffname"`
	Childrecorddifftypename  string `json:"childrecorddifftypename"` // category label
	Childrecorddiffname      string `json:"childrecorddiffname"`     // category value
	Catename                 string
	Parentcatname            string
}

//RecorddifferentiationhigherkeyEntities is used for response
type RecorddifferentiationhigherkeyEntities struct {
	Total  int64                                  `json:"total"`
	Values []RecorddifferentiationhigherkeyEntity `json:"values"`
}

//RecorddifferentiationhigherkeyEntityResponse is used for all response
type RecorddifferentiationhigherkeyEntityResponse struct {
	Success bool                                   `json:"success"`
	Message string                                 `json:"message"`
	Details RecorddifferentiationhigherkeyEntities `json:"details"`
}

//RecorddifferentiationhigherkeyEntityResponseInt is used for int response
type RecorddifferentiationhigherkeyEntityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

//FromJSON is used to convert json
func (w *RecorddifferentiationhigherkeyEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
