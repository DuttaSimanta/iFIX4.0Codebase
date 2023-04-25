package entities

import (
	"encoding/json"
	"io"
)

type RecorddifferentiongroupEntity struct {
	Id                             int64   `json:"id"`
	Clientid                       int64   `json:"clientid"`
	Mstorgnhirarchyid              int64   `json:"mstorgnhirarchyid"`
	Mstworkdifferentiationtypeid   int64   `json:"mstworkdifferentiationtypeid"`
	Mstworkdifferentiationid       int64   `json:"mstworkdifferentiationid"`
	Mstworkdifferentiationtypeids  []int64 `json:"mstworkdifferentiationtypeids"`
	Mstworkdifferentiationids      []int64 `json:"mstworkdifferentiationids"`
	Mstgroupid                     int64   `json:"mstgroupid"`
	Mstuserid                      int64   `json:"mstuserid"`
	Activeflg                      int64   `json:"activeflg"`
	Offset                         int64   `json:"offset"`
	Limit                          int64   `json:"limit"`
	Clientname                     string  `json:"clientname"`
	Mstorgnhirarchyname            string  `json:"mstorgnhirarchyname"`
	Recorddifftypeid               int64   `json:"recorddifftypeid"`
	Recorddifftypename             string  `json:"recorddifftypename"`
	Recorddiffid                   int64   `json:"recorddiffid"`
	Recorddiffname                 string  `json:"recorddiffname"`
	Mstworkdifferentiationtypename string  `json:"mstworkdifferentiationtypename"`
	Supportgroupname               string  `json:"Supportgroupname"`
	Mstworkdifferentiationname     string  `json:"mstworkdifferentiationname"`
	Name                           string
	Parentcategorynames            string
}

type WorkinglevelEntity struct {
	Mstworkdifferentiationid int64  `json:"mstworkdifferentiationid"`
	Levelname                string `json:"levelname"`
}

type RecorddifferentiongroupEntities struct {
	Total  int64                           `json:"total"`
	Values []RecorddifferentiongroupEntity `json:"values"`
}

type WorkinglevelEntities struct {
	Total  int64                `json:"total"`
	Values []WorkinglevelEntity `json:"values"`
}

type RecorddifferentiongroupResponse struct {
	Success bool                            `json:"success"`
	Message string                          `json:"message"`
	Details RecorddifferentiongroupEntities `json:"details"`
}

type WorkinglevelResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details WorkinglevelEntities `json:"details"`
}

type RecorddifferentiongroupResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *RecorddifferentiongroupEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
