package entities

import (
	"encoding/json"
	"io"
)

type RecorddifferentiationtypeEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Typename            string `json:"typename"`
	Seqno               int64  `json:"seqno"`
	Parentid            int64  `json:"parentid"`
	Activeflg           int64  `json:"activeflg"`
	Audittransactionid  int64  `json:"audittransactionid"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Parentname          string `json:"parentname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}

type RecorddifferentiationtypeEntities struct {
	Total  int64                             `json:"total"`
	Values []RecorddifferentiationtypeEntity `json:"values"`
}

type RecorddifferentiationtypeResponse struct {
	Success bool                              `json:"success"`
	Message string                            `json:"message"`
	Details RecorddifferentiationtypeEntities `json:"details"`
}

type RecorddifferentiationtypeResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *RecorddifferentiationtypeEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
