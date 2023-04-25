package entities

import (
	"encoding/json"
	"io"
)

type DashboardQueryEntity struct {
	Id int64 `json:"id"`
	//HeaderName string `json:"headername"`
	Clientid             int64  `json:"clientid"`
	Clientname           string `json:"clienname"`
	Mstorgnhirarchyid    int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname  string `json:"mstorgnhirarchyname"`
	RecordDiffid         int64  `json:"recorddiffid"`
	RecordDiffName       string `json:"recorddiffname"`
	Tilesid              int64  `json:"tilesid"`
	TilesName            string `json:"tilesname"`
	QueryType            int64  `json:"querytype"`
	QueryTypename        string `json:"querytypename"`
	Query                string `json:"query"`
	QueryParam           string `json:"queryparam"`
	JoinQuery            string `json:"joinquery"`
	Ismanagerialview     int64  `json:"ismanegerialview"`
	IsmanagerialviewName string `json:"ismanagerialviewname"`

	Activeflg int64 `json:"activeglg"`
	Offset    int64 `json:"offset"`
	Limit     int64 `json:"limit"`
}
type DashboardQueryEntities struct {
	Total  int64                  `json:"total"`
	Values []DashboardQueryEntity `json:"values"`
}

type DashboardQueryResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details DashboardQueryEntities `json:"details"`
}

type DashboardQueryResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *DashboardQueryEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
