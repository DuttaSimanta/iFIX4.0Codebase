package entities

import (
	"encoding/json"
	"io"
)

type MstSupportGroupWorkingHoursEntity struct {
	ID                int64                                      `json:"id"`
	Clientid          int64                                      `json:"clientid"`
	Mstorgnhirarchyid int64                                      `json:"mstorgnhirarchyid"`
	Supportgroupid    int64                                      `json:"supportgroupid"`
	Offset            int64                                      `json:"offset"`
	Limit             int64                                      `json:"limit"`
	Details           []MstSupportGroupWorkingHoursdetailsEntity `json:"details"`
}

type MstSupportGroupWorkingHoursdetailsEntity struct {
	Dayofweekid      int64  `json:"dayofweekid"`
	Starttimeinteger int64  `json:"starttimeinteger"`
	Starttime        string `json:"starttime"`
	Endtimeinteger   int64  `json:"endtimeinteger"`
	Endtime          string `json:"endtime"`
	Nextdayforward   int64  `json:"nextdayforward"`
	Activeflg        int64  `json:"activeflg"`
}

type MstSupportGroupWorkingHoursUpdateEntity struct {
	ID                int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	Supportgroupid    int64  `json:"supportgroupid"`
	Dayofweekid       int64  `json:"dayofweekid"`
	Starttimeinteger  int64  `json:"starttimeinteger"`
	Starttime         string `json:"starttime"`
	Endtimeinteger    int64  `json:"endtimeinteger"`
	Endtime           string `json:"endtime"`
	Nextdayforward    int64  `json:"nextdayforward"`
}

type MstSupportGroupWorkingHoursresponseEntity struct {
	ID                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Supportgroupid      int64  `json:"supportgroupid"`
	Dayofweekid         int64  `json:"dayofweekid"`
	Starttimeinteger    int64  `json:"starttimeinteger"`
	Starttime           string `json:"starttime"`
	Endtimeinteger      int64  `json:"endtimeinteger"`
	Endtime             string `json:"endtime"`
	Nextdayforward      int64  `json:"nextdayforward"`
	Activeflg           int64  `json:"activeflg"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Supportgroupname    string `json:"supportgroupname"`
}

type MstSupportGroupWorkingHoursEntities struct {
	Total  int64                                       `json:"total"`
	Values []MstSupportGroupWorkingHoursresponseEntity `json:"values"`
}

type MstSupportGroupWorkingHoursResponse struct {
	Success bool                                `json:"success"`
	Message string                              `json:"message"`
	Details MstSupportGroupWorkingHoursEntities `json:"details"`
}

type SupportGroupWiseWorkingHoursResponse struct {
	Success bool                                        `json:"success"`
	Message string                                      `json:"message"`
	Details []MstSupportGroupWorkingHoursresponseEntity `json:"details"`
}

type MstSupportGroupWorkingHoursResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstSupportGroupWorkingHoursEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *MstSupportGroupWorkingHoursUpdateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *MstSupportGroupWorkingHoursresponseEntity) DetailsFromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
