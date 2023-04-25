package entities

import (
	"encoding/json"
	"io"
)

//id, clientid, mstorgnhirarchyid, name, metertypeid, seqno

type SlaTermEntryEntity struct {
	Id int64 `json:"id"`

	Clientid            int64    `json:"clientid"`
	ToClientid          int64    `json:"toclientid"`
	Clientname          string   `json:"clienname"`
	Mstorgnhirarchyid   int64    `json:"mstorgnhirarchyid"`
	ToMstorgnhirarchyid int64    `json:"tomstorgnhirarchyid"`
	Mstorgnhirarchyname string   `json:"mstorgnhirarchyname"`
	MeterNames          []string `json:"meternames"`
	MeterName           string   `json:"metername"`
	MetertTypeid        int64    `json:"metertypeid"`
	MetertypeName       string   `json:"metertypename"`
	Seqno               int64    `json:"seqno"`
	Activeflg           int64    `json:"activeglg"`
	Offset              int64    `json:"offset"`
	Limit               int64    `json:"limit"`
}
type SlaTermEntryEntities struct {
	Total  int64                `json:"total"`
	Values []SlaTermEntryEntity `json:"values"`
}

type SlaTermEntryResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details SlaTermEntryEntities `json:"details"`
}

type SlaTermEntryResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *SlaTermEntryEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
