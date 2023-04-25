package entities

import (
	"encoding/json"
	"io"
)

type MapcommontileswithgroupEntity struct {
	Id                            int64   `json:"id"`
	Clientid                      int64   `json:"clientid"`
	Mstorgnhirarchyid             int64   `json:"mstorgnhirarchyid"`
	Urlkey                        int64   `json:"urlkey"`
	Recorddifftypeid              int64   `json:"recorddifftypeid"`
	Recorddiffid                  int64   `json:"recorddiffid"`
	Groupid                       []int64 `json:"groupid"`
	Activeflg                     int64   `json:"activeflg"`
	Offset                        int64   `json:"offset"`
	Limit                         int64   `json:"limit"`
	Clientname                    string  `json:"clientname"`
	Mstorgnhirarchyname           string  `json:"mstorgnhirarchyname"`
	Recorddifferentiationtypename string  `json:"recorddifferentiationtypename"`
	Recorddifferentiationname     string  `json:"recorddifferentiationname"`
	Urlname                       string  `json:"urlname"`
	Supportgrpname                string  `json:"supportgrpname"`
	Supportgrpid                  string  `json:"supportgrpid"`
}

type MapcommontileswithgroupEntities struct {
	Total  int64                           `json:"total"`
	Values []MapcommontileswithgroupEntity `json:"values"`
}

type MapcommontileswithgroupResponse struct {
	Success bool                            `json:"success"`
	Message string                          `json:"message"`
	Details MapcommontileswithgroupEntities `json:"details"`
}

type MapcommontileswithgroupResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MapcommontileswithgroupEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
