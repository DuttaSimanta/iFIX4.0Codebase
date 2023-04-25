package entities

import (
	"encoding/json"
	"io"
)

//MstClientUserEntity contains all required data fields
type LocationPriorityEntity struct {
	ID                   int64  `json:"id"`
	ClientID             int64  `json:"clientid"`
	MstorgnhirarchyID    int64  `json:"mstorgnhirarchyid"`
	ClientName           string `json:"clientname"`
	MstorgnhirarchyName  string `json:"mstorgnhirarchyname"`
	Location             string `json:"location"`
	Recorddifftypeid     int64  `json:"recorddifftypeid"`
	RecorddifftypeName   string `json:"recorddifftypename"`
	Recorddiffid         int64  `json:"recorddiffid"`
	ReccorddiffName      string `json:"recorddiffname"`
	ToRecorddifftypeid   int64  `json:"torecorddifftypeid"`
	ToRecorddifftypeName string `json:"torecorddifftypename"`
	ToRecorddiffid       int64  `json:"torecorddiffid"`
	ToReccorddiffName    string `json:"torecorddiffname"`
	Activeflg            int64  `json:"activeflg"`
	Limit                int64  `json:"limit"`
	Offset               int64  `json:"offset"`
}
type LocationEntities struct {
	Total  int64                    `json:"total"`
	Values []LocationPriorityEntity `json:"values"`
}
type LocationResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}
type LocationResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details LocationEntities `json:"details"`
}

type LocationSearchEntity struct {
	ID       int64  `json:"id"`
	Location string `json:"location"`
}
type LocationSelectEntity struct {
	ID                 int64  `json:"id"`
	Location           string `json:"location"`
	Priorityid         int64  `json:"priorityid"`
	Priority           string `json:"priority"`
	Recorddifftypeid   int64  `json:"recorddifftypeid"`
	RecorddifftypeName string `json:"recorddifftypename"`
}
type LocationSearchEntityResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details []LocationSearchEntity `json:"details"`
}
type LocationSelectEntityResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details []LocationSelectEntity `json:"details"`
}

func (p *LocationPriorityEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
