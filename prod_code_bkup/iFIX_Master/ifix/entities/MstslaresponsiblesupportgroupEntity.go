package entities

import (
	"encoding/json"
	"io"
)

type MstslaresponsiblesupportgroupEntity struct {
	Id                           int64  `json:"id"`
	Clientid                     int64  `json:"clientid"`
	Mstorgnhirarchyid            int64  `json:"mstorgnhirarchyid"`
	Mstslafullfillmentcriteriaid int64  `json:"mstslafullfillmentcriteriaid"`
	Mstclientsupportgroupid      int64  `json:"mstclientsupportgroupid"`
	Mstslaid                     int64  `json:"mstslaid"`
	Activeflg                    int64  `json:"activeflg"`
	Offset                       int64  `json:"offset"`
	Limit                        int64  `json:"limit"`
	Clientname                   string `json:"clientname"`
	Mstorgnhirarchyname          string `json:"mstorgnhirarchyname"`
	Slaname                      string `json:"slaname"`
	Grpname                      string `json:"grpname"`
}
type MstslaresponsiblesupportgroupEntities struct {
	Total  int64                                 `json:"total"`
	Values []MstslaresponsiblesupportgroupEntity `json:"values"`
}

type Mstslanames struct {
	Id      int64  `json:"id"`
	Slaname string `json:"slaname"`
}

type MstslanamesResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Details []Mstslanames `json:"details"`
}

type MstslaresponsiblesupportgroupResponse struct {
	Success bool                                  `json:"success"`
	Message string                                `json:"message"`
	Details MstslaresponsiblesupportgroupEntities `json:"details"`
}

type MstslaresponsiblesupportgroupResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstslaresponsiblesupportgroupEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
