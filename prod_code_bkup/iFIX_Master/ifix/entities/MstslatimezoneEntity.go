package entities

import (
	"encoding/json"
	"io"
)

type MstslatimezoneEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstslaid            int64  `json:"mstslaid"`
	Msttimezoneid       int64  `json:"msttimezoneid"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Zonename            string `json:"zonename"`
	Slaname             string `json:"slaname"`
}

type MstslatimezoneEntities struct {
	Total  int64                  `json:"total"`
	Values []MstslatimezoneEntity `json:"values"`
}

type MstslatimezoneResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details MstslatimezoneEntities `json:"details"`
}

type MstslatimezoneResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstslatimezoneEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
