package entities

import (
	"encoding/json"
	"io"
)

type AssetvalidateEntity struct {
	Id                            int64  `json:"id"`
	Clientid                      int64  `json:"clientid"`
	Mstorgnhirarchyid             int64  `json:"mstorgnhirarchyid"`
	Mstdifferentiationtypeid      int64  `json:"mstdifferentiationtypeid"`
	Mstdifferentiationid          int64  `json:"mstdifferentiationid"`
	Validationrule                string `json:"validationrule"`
	Activeflg                     int64  `json:"activeflg"`
	Audittransactionid            int64  `json:"audittransactionid"`
	Offset                        int64  `json:"offset"`
	Limit                         int64  `json:"limit"`
	Clientname                    string `json:"clientname"`
	Mstorgnhirarchyname           string `json:"mstorgnhirarchyname"`
	Recorddifferentiationname     string `json:"recorddifferentiationname"`
	Recorddifferentiationtypename string `json:"recorddifferentiationtypename"`
}

type AssetvalidateEntities struct {
	Total  int64                 `json:"total"`
	Values []AssetvalidateEntity `json:"values"`
}

type AssetvalidateResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details AssetvalidateEntities `json:"details"`
}

type AssetvalidateResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *AssetvalidateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
