package entities

import (
	"encoding/json"
	"io"
)

type MstslafcrecorddiffEntity struct {
	Id                       int64  `json:"id"`
	Clientid                 int64  `json:"clientid"`
	Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
	Mstslaid                 int64  `json:"mstslaid"`
	Recorddifftypeidtype     int64  `json:"recorddifftypeidtype"`
	Recorddiffidtype         int64  `json:"recorddiffidtype"`
	Recorddifftypeidstatus   int64  `json:"recorddifftypeidstatus"`
	Recorddiffidstatus       int64  `json:"recorddiffidstatus"`
	Startstopindicator       int64  `json:"startstopindicator"`
	Activeflg                int64  `json:"activeflg"`
	Offset                   int64  `json:"offset"`
	Limit                    int64  `json:"limit"`
	Clientname               string `json:"clientname"`
	Mstorgnhirarchyname      string `json:"mstorgnhirarchyname"`
	Recorddifftypeidtypenm   string `json:"recorddifftypeidtypenm"`
	Recorddiffidtypenm       string `json:"recorddiffidtypenm"`
	Slaname                  string `json:"slaname"`
	Recorddifftypeidstatusnm string `json:"recorddifftypeidstatusnm"`
	Recorddiffidstatusnm     string `json:"recorddiffidstatusnm"`

	Recorddifftypetypeparent     int64 `json:"recorddifftypetypeparent"`
	Recorddifftypeidstatusparent int64 `json:"recorddifftypeidstatusparent"`

	SLAmetertypeID         int64  `json:"slametertypeid"`
	SLAmetertypename       string `json:"slametertypename"`
	Startstopindicatorname string `json:"startstopindicatorname"`
}

type MstslafcrecorddiffEntities struct {
	Total  int64                      `json:"total"`
	Values []MstslafcrecorddiffEntity `json:"values"`
}

type MstslafcrecorddiffResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details MstslafcrecorddiffEntities `json:"details"`
}

type MstslafcrecorddiffResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

type MstslametertypeEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MstslametertypeResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details []MstslametertypeEntity `json:"details"`
}

type MstslaindicatortermEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MstslaindicatortermResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details []MstslaindicatortermEntity `json:"details"`
}

func (w *MstslafcrecorddiffEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
