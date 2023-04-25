package entities

import (
	"encoding/json"
	"io"
)

type ClientdayofweekEntity struct {
	ID                int64                          `json:"id"`
	Clientid          int64                          `json:"clientid"`
	Mstorgnhirarchyid int64                          `json:"mstorgnhirarchyid"`
	Offset            int64                          `json:"offset"`
	Uploadedfilename  string                         `json:"uploadedfilename"`
	Originallogoname  string                         `json:"originalfilename"`
	Limit             int64                          `json:"limit"`
	Details           []ClientdayofweekdetailsEntity `json:"details"`
}

type ClientdayofweekdetailsEntity struct {
	Dayofweekid      int64  `json:"dayofweekid"`
	Starttimeinteger int64  `json:"starttimeinteger"`
	Starttime        string `json:"starttime"`
	Endtimeinteger   int64  `json:"endtimeinteger"`
	Endtime          string `json:"endtime"`
	Nextdayforward   int64  `json:"nextdayforward"`
	Activeflg        int64  `json:"activeflg"`
}

type ClientdayofweekresponseEntity struct {
	ID                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Dayofweekid         int64  `json:"dayofweekid"`
	Starttimeinteger    int64  `json:"starttimeinteger"`
	Starttime           string `json:"starttime"`
	Endtimeinteger      int64  `json:"endtimeinteger"`
	Endtime             string `json:"endtime"`
	Nextdayforward      int64  `json:"nextdayforward"`
	Activeflg           int64  `json:"activeflg"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Uploadedfilename    string `json:"uploadedfilename"`
	Originallogoname    string `json:"originalfilename"`
}

type ClientdayofweekEntities struct {
	Total  int64                           `json:"total"`
	Values []ClientdayofweekresponseEntity `json:"values"`
}

type ClientdayofweekResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details ClientdayofweekEntities `json:"details"`
}

type ClientwisedayofweekResponse struct {
	Success bool                            `json:"success"`
	Message string                          `json:"message"`
	Details []ClientdayofweekresponseEntity `json:"details"`
}

type ClientdayofweekResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *ClientdayofweekEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *ClientdayofweekresponseEntity) DetailsFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
