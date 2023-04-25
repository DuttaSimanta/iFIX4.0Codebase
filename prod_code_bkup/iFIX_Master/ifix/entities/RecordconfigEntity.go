package entities

import (
	"encoding/json"
	"io"
)

type RecordconfigEntity struct {
	Id                            int64  `json:"id"`
	Clientid                      int64  `json:"clientid"`
	Mstorgnhirarchyid             int64  `json:"mstorgnhirarchyid"`
	Recorddifftypeid              int64  `json:"recorddifftypeid"`
	Recorddiffid                  int64  `json:"recorddiffid"`
	Recordtypeid                  int64  `json:"recordtypeid"`
	Prefix                        string `json:"prefix"`
	Year                          string `json:"year"`
	Month                         string `json:"month"`
	Day                           string `json:"day"`
	Configurezero                 string `json:configurezero`
	IsClient                      int64  `json:"isclient"`
	Activeflg                     int64  `json:"activeflg"`
	Offset                        int64  `json:"offset"`
	Limit                         int64  `json:"limit"`
	Clientname                    string `json:"clientname"`
	Mstorgnhirarchyname           string `json:"mstorgnhirarchyname"`
	Recorddifferentiationname     string `json:"recorddifferentiationname"`
	Recorddifferentiationtypename string `json:"recorddifferentiationtypename"`
}

type RecordconfigEntities struct {
	Total  int64                `json:"total"`
	Values []RecordconfigEntity `json:"values"`
}

type RecordconfigResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details RecordconfigEntities `json:"details"`
}

type RecordconfigResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *RecordconfigEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
