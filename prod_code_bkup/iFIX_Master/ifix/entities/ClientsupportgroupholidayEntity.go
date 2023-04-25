package entities

import (
	"encoding/json"
	"io"
)

type ClientsupportgroupholidayEntity struct {
	Id                      int64  `json:"id"`
	Clientid                int64  `json:"clientid"`
	Mstorgnhirarchyid       int64  `json:"mstorgnhirarchyid"`
	Mstclientsupportgroupid int64  `json:"mstclientsupportgroupid"`
	Dateofholiday           string `json:"dateofholiday"`
	Plannedornot            int64  `json:"plannedornot"`
	Dayofweekid             int64  `json:"dayofweekid"`
	Starttime               string `json:"starttime"`
	Starttimeinteger        int64  `json:"starttimeinteger"`
	Endtime                 string `json:"endtime"`
	Endtimeinteger          int64  `json:"endtimeinteger"`
	Activeflg               int64  `json:"activeflg"`
	Audittransactionid      int64  `json:"audittransactionid"`
	Offset                  int64  `json:"offset"`
	Limit                   int64  `json:"limit"`
	Clientname              string `json:"clientname"`
	Mstorgnhirarchyname     string `json:"mstorgnhirarchyname"`
	Supportgroupname        string `json:"supportgroupname"`
}

type SupportgrpEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ClientsupportgroupholidayEntities struct {
	Total  int64                             `json:"total"`
	Values []ClientsupportgroupholidayEntity `json:"values"`
}

type ClientsupportgroupholidayResponse struct {
	Success bool                              `json:"success"`
	Message string                            `json:"message"`
	Details ClientsupportgroupholidayEntities `json:"details"`
}

type SupportgrpResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Values  []SupportgrpEntity `json:"values"`
}

type ClientsupportgroupholidayResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *ClientsupportgroupholidayEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
