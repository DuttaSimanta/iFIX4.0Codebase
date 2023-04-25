package entities

import (
	"encoding/json"
	"io"
)

type DashboardtilesinputEntity struct {
	Clientid                int64 `json:"clientid"`
	Mstorgnhirarchyid       int64 `json:"mstorgnhirarchyid"`
	UserID                  int64 `json:"userid"`
	GroupID                 int64 `json:"groupid"`
	Recorddifftypeid        int64 `json:"recorddifftypeid"`
	Recorddiffid            int64 `json:"recorddiffid"`
	Recorddifftypeseq       int64 `json:"recorddifftypeseq"`
	Recorddifftypestatusseq int64 `json:"recorddifftypestatusseq"`
	Ismanagerialview        int64 `json:"ismanagerialview"`
	Iscatalog               int64 `json:"iscatalog"`
}

type DashboardtilesresponseEntity struct {
	ID                int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	Mstfunctionailyid int64  `json:"mstfunctionailyid"`
	Diffid            int64  `json:"diffid"`
	Description       string `json:"description"`
	Seqno             int64  `json:"seqno"`
	Colorcode         string `json:"colorcode"`
	Image             string `json:"image"`
	Readpermission    int64  `json:"readpermission"`
	Writepermission   int64  `json:"writepermission"`
}
type RecordnamesEntity struct {
	Clientid          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`
	Linkrecordid      int64 `json:"linkrecordid"`
}

func (w *DashboardtilesinputEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
func (w *RecordnamesEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

type DashboardtilesAllResponse struct {
	Success bool                           `json:"success"`
	Message string                         `json:"message"`
	Details []DashboardtilesresponseEntity `json:"details"`
}
type RecordnamesResponse struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	Recordnames string `json:"recordnames"`
}
type Dashboardbuttontab struct {
	Tabs    []DashboardtilesresponseEntity
	Buttons []DashboardtilesresponseEntity
	Count   []DashboardtilesresponseEntity
}

type DashboardbuttontabAllResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details Dashboardbuttontab `json:"details"`
}
