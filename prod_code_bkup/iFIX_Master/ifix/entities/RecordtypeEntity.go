package entities

import (
	"encoding/json"
	"io"
)

type RecordtypeEntity struct {
	Id                   int64   `json:"id"`
	Clientid             int64   `json:"clientid"`
	Mstorgnhirarchyid    int64   `json:"mstorgnhirarchyid"`
	Fromrecorddifftypeid int64   `json:"fromrecorddifftypeid"`
	Fromrecorddiffid     int64   `json:"fromrecorddiffid"`
	Fromrecorddiffids    []int64 `json:"fromrecorddiffids"`

	Torecorddifftypeid     int64  `json:"torecorddifftypeid"`
	Torecorddiffid         int64  `json:"torecorddiffid"`
	Seqno                  int64  `json:"seqno"`
	Parentid               int64  `json:"Parentid"`
	Activeflg              int64  `json:"activeflg"`
	Audittransactionid     int64  `json:"audittransactionid"`
	Offset                 int64  `json:"offset"`
	Limit                  int64  `json:"limit"`
	Clientname             string `json:"clientname"`
	Mstorgnhirarchyname    string `json:"mstorgnhirarchyname"`
	Fromrecorddifftypename string `json:"fromrecorddifftypename"`
	Fromrecorddiffname     string `json:"fromrecorddiffname"`
	Torecorddifftypename   string `json:"torecorddifftypename"`
	Torecorddiffname       string `json:"torecorddiffname"`
	Title                  string `json:"title"`
	Description            string `json:"description"`
}
type Recordtypesingleentity struct {
	Id               int64  `json:"id"`
	Recorddifftypeid int64  `json:"recorddifftypeid"`
	Seqno            int64  `json:"seqno"`
	Typename         string `json:"typename"`
	Parentpath       string `json:"parentpath"`
}
type RecordtypeEntities struct {
	Total  int64              `json:"total"`
	Values []RecordtypeEntity `json:"values"`
}

type RecordtypesingleResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	Details []Recordtypesingleentity `json:"details"`
}
type RecordtypeResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details RecordtypeEntities `json:"details"`
}

type RecordtypeResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *RecordtypeEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
