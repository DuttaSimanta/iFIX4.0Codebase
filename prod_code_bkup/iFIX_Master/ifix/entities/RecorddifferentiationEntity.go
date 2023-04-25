package entities

import (
	"encoding/json"
	"io"
)

type RecorddifferentiationEntity struct {
	Id                   int64  `json:"id"`
	Clientid             int64  `json:"clientid"`
	Mstorgnhirarchyid    int64  `json:"mstorgnhirarchyid"`
	Recorddifftypeid     int64  `json:"recorddifftypeid"`
	Parentid             int64  `json:"parentid"`
	Name                 string `json:"name"`
	Seqno                int64  `json:"seqno"`
	Typeseqno            int64  `json:"typeseqno"`
	Activeflg            int64  `json:"activeflg"`
	Audittransactionid   int64  `json:"audittransactionid"`
	Offset               int64  `json:"offset"`
	Limit                int64  `json:"limit"`
	Clientname           string `json:"clientname"`
	Mstorgnhirarchyname  string `json:"mstorgnhirarchyname"`
	Recorddifftypname    string `json:"recorddifftypname"`
	Parentcatagorytypeid int64  `json:"parentcatagorytypeid"`
}

type RecorddifferentiationEntities struct {
	Total  int64                         `json:"total"`
	Values []RecorddifferentiationEntity `json:"values"`
}
type RecorddifferentionSingle struct {
	Id                  int64  `json:"id"`
	Parentid            int64  `json:"parentid"`
	Name                string `json:"name"`
	Parentcategorynames string `json:"parentcategorynames"`
	Sortedcategorynames string `json:"sortedcategorynames"`
	Recorddifftypeid    int64  `json:"recorddifftypeid"`
}
type RecorddifferentionRec struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details []RecorddifferentionSingle `json:"details"`
}
type RecorddifferentiationResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Details RecorddifferentiationEntities `json:"details"`
}

type RecorddifferentiationResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

type RecorddifferentiationnameEntity struct {
	Id             int64  `json:"id"`
	Recorddiffname string `json:"recorddiffname"`
	Parentcatname  string
	Diffname       string
}
type RecorddifferentiationnameEntities struct {
	Values []RecorddifferentiationnameEntity `json:"values"`
}
type RecorddifferentiationnameResponse struct {
	Success bool                              `json:"success"`
	Message string                            `json:"message"`
	Details RecorddifferentiationnameEntities `json:"details"`
}

func (w *RecorddifferentiationEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *RecorddifferentiationnameEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
