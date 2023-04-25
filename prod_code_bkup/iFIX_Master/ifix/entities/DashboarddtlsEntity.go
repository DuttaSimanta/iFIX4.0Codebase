package entities

import (
	"encoding/json"
	"io"
)

type DashboarddtlsEntity struct {
	Id                         int64  `json:"id"`
	Clientid                   int64  `json:"clientid"`
	Mstorgnhirarchyid          int64  `json:"mstorgnhirarchyid"`
	Mstrecorddifferentiationid int64  `json:"mstrecorddifferentiationid"`
	Mapfunctionalityid         int64  `json:"mapfunctionalityid"`
	Querytype                  int64  `json:"querytype"`
	Query                      string `json:"query"`
	Queryparam                 string `json:"queryparam"`
	Activeflg                  int64  `json:"activeflg"`
	Clientname                 string `json:"clientname"`
	Mstorgnhirarchyname        string `json:"mstorgnhirarchyname"`
	Recorddifferentiationname  string `json:"recorddifferentiationname"`
	Mapfunctionalityname       string `json:"mapfunctionalityname"`
	Offset                     int64  `json:"offset"`
	Limit                      int64  `json:"limit"`
}

type DashboarddtlsEntities struct {
	Total  int64                 `json:"total"`
	Values []DashboarddtlsEntity `json:"values"`
}

type DashboarddtlsResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details DashboarddtlsEntities `json:"details"`
}

type DashboarddtlsResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *DashboarddtlsEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
