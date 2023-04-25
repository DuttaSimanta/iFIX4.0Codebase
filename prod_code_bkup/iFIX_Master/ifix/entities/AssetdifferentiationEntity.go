package entities

import (
	"encoding/json"
	"io"
)

type AssetdifferentiationEntity struct {
	Id                       int64  `json:"id"`
	Clientid                 int64  `json:"clientid"`
	Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
	Trnassetid               int64  `json:"trnassetid"`
	Mstdifferentiationtypeid int64  `json:"mstdifferentiationtypeid"`
	Mstdifferentiationid     int64  `json:"mstdifferentiationid"`
	Value           	 string `json:"value"`
	Deleteflg                int64  `json:"deleteflg"`
	Activeflg                int64  `json:"activeflg"`
	Offset                   int64  `json:"offset"`
	Limit                    int64  `json:"limit"`
	Clientname               string `json:"clientname"`
	Mstorgnhirarchyname      string `json:"mstorgnhirarchyname"`
	Recorddifftypename       string `json:"recorddifftypename"`
	Recorddiffname           string `json:"recorddiffname"`
	Assetid                  string `json:"assetid"`
}

type AssetdifferentiationEntities struct {
	Total  int64                        `json:"total"`
	Values []AssetdifferentiationEntity `json:"values"`
}

type AssetdifferentiationResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Details AssetdifferentiationEntities `json:"details"`
}

type AssetdifferentiationResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *AssetdifferentiationEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
