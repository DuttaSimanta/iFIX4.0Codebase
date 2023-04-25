package entities

import (
	"encoding/json"
	"io"
)
type RecorddifftypeAndRecordTypeEntity struct{
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Typename            string `json:"typename"`
	Seqno               int64  `json:"seqno"`
	Parentid            int64  `json:"parentid"`
	Fromrecorddifftypeid   int64  `json:"fromrecorddifftypeid"`
	Fromrecorddiffid       int64  `json:"fromrecorddiffid"`
	Torecorddifftypeid     int64  `json:"torecorddifftypeid"`
	Torecorddiffid         int64  `json:"torecorddiffid"`
	/*Activeflg           int64  `json:"activeflg"`
	Audittransactionid  int64  `json:"audittransactionid"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`*/
}
type RecorddifftypeAndRecordTypeEntities struct {
	Total  int64              `json:"total"`
	Values []RecorddifftypeAndRecordTypeEntity `json:"values"`
}
type RecorddifftypeAndRecordTypeResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
func (w *RecorddifftypeAndRecordTypeEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
