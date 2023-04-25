package entities

import (
	"encoding/json"
	"io"
)

type RecordTermAdditionalMapEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Recordtermid        int64  `json:"recordtermid"`
	RecordtermName      string `json:"recordtermname"`
	RecordfieldtypeName string `json:"recordfieldtypename"`
	DisplaySeq          int64  `json:"displayseq"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
}
type AdditionalTabEntity struct {
	Id      int64  `json:"id"`
	TabName string `json:"tabname"`
}

type RecordTermAdditionalMapEntities struct {
	Total  int64                           `json:"total"`
	Values []RecordTermAdditionalMapEntity `json:"values"`
}
type AdditionalTabResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details []AdditionalTabEntity `json:"details"`
}
type RecordTermAdditionalMapResponse struct {
	Success bool                            `json:"success"`
	Message string                          `json:"message"`
	Details RecordTermAdditionalMapEntities `json:"details"`
}

type RecordTermAdditionalMapResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *RecordTermAdditionalMapEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
