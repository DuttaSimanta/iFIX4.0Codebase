package entities

import (
	"encoding/json"
	"io"
)

type WorkdifferentiationEntity struct {
	Id                   int64  `json:"id"`
	Clientid             int64  `json:"clientid"`
	Mstorgnhirarchyid    int64  `json:"mstorgnhirarchyid"`
	Forrecorddifftypeid  int64  `json:"forrecorddifftypeid"`
	Forrecorddiffid      int64  `json:"forrecorddiffid"`
	Mainrecorddifftypeid int64  `json:"mainrecorddifftypeid"`
	Activeflg            int64  `json:"activeflg"`
	Audittransactionid   int64  `json:"audittransactionid"`
	Offset               int64  `json:"offset"`
	Limit                int64  `json:"limit"`
	Clientname           string `json:"clientname"`
	Mstorgnhirarchyname  string `json:"mstorgnhirarchyname"`
	Recorddifftypname    string `json:"recorddifftypname"`
	Recorddiffname       string `json:"recorddiffname"`
	Recorddifftyplabel   string `json:"recorddifftyplabel"`
}
type WorkdifferentiationsingleEntity struct {
	Id             int64  `json:"id"`
	Recorddiffname string `json:"recorddiffname"`
}
type WorkdifferentiationsingleResponse struct {
	Success bool                              `json:"success"`
	Message string                            `json:"message"`
	Details []WorkdifferentiationsingleEntity `json:"details"`
}
type WorkdifferentiationEntities struct {
	Total  int64                       `json:"total"`
	Values []WorkdifferentiationEntity `json:"values"`
}

type WorkdifferentiationResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details WorkdifferentiationEntities `json:"details"`
}
type Workinglabelname struct {
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	Recorddifftypid     int64  `json:"recorddifftypid"`
	Forrecorddifftypeid int64  `json:"forrecorddifftypeid"`
	Forrecorddiffid     int64  `json:"forrecorddiffid"`
	Recorddifftypname   string `json:"recorddifftypname"`
	Workingcatename     string
	Parentcatenames     string
}
type WorkinglabelnameEntities struct {
	Values []Workinglabelname `json:"values"`
}

type WorkinglabelnameResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	Details WorkinglabelnameEntities `json:"details"`
}

type WorkdifferentiationResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *WorkdifferentiationEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
