package entities

import (
	"encoding/json"
	"io"
)

type ClientholidayEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Dateofholiday       string `json:"dateofholiday"`
	Plannedornot        int64  `json:"plannedornot"`
	Dayofweekid         int64  `json:"dayofweekid"`
	Starttime           string `json:"starttime"`
	Starttimeinteger    int64  `json:"starttimeinteger"`
	Endtime             string `json:"endtime"`
	Endtimeinteger      int64  `json:"endtimeinteger"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}

type ClientholidayEntities struct {
	Total  int64                 `json:"total"`
	Values []ClientholidayEntity `json:"values"`
}

type ClientholidayResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details ClientholidayEntities `json:"details"`
}

type ClientholidayResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *ClientholidayEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
