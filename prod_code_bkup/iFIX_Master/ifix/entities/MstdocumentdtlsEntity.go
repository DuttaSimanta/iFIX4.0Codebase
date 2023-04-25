package entities

import (
	"encoding/json"
	"io"
)

type MstdocumentdtlsEntity struct {
	Id                            int64   `json:"id"`
	Clientid                      int64   `json:"clientid"`
	Mstorgnhirarchyid             int64   `json:"mstorgnhirarchyid"`
	Recorddifftypeid              int64   `json:"recorddifftypeid"`
	Recorddiffid                  int64   `json:"recorddiffid"`
	Groupid                       []int64 `json:"groupid"`
	Documentname                  string  `json:"documentname"`
	Orginaldocumentname           string  `json:"orginaldocumentname"`
	Documentpath                  string  `json:"documentpath"`
	Credentialid                  int64   `json:"credentialid"`
	Activeflg                     int64   `json:"activeflg"`
	Offset                        int64   `json:"offset"`
	Limit                         int64   `json:"limit"`
	Clientname                    string  `json:"clientname"`
	Mstorgnhirarchyname           string  `json:"mstorgnhirarchyname"`
	Recorddifferentiationtypename string  `json:"recorddifferentiationtypename"`
	Recorddifferentiationname     string  `json:"recorddifferentiationname"`
	Supportgroupname              string  `json:"supportgroupname"`
	Supportgroupid                string  `json:"supportgroupid"`
	Usagecount                    int64   `json:"usagecount"`
}

type MstdocumentdtlsEntities struct {
	Total  int64                   `json:"total"`
	Values []MstdocumentdtlsEntity `json:"values"`
}

type MstdocumentdtlsResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details MstdocumentdtlsEntities `json:"details"`
}

type MstdocumentdtlsResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstdocumentdtlsEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
