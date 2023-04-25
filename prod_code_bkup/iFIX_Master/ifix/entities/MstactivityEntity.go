package entities

import (
	"encoding/json"
	"io"
)

type MstactivityEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Actiontypeid        int64  `json:"actiontypeid"`
	Processid           int64  `json:"processid"`
	Actionname          string `json:"actionname"`
	Description         string `json:"description"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Processname         string `json:"processname"`
	Actiontypename      string `json:"actiontypename"`
}
type MstactivitySingleEntity struct {
	Id                  int64  `json:"id"`
	Actiontypeid        int64  `json:"actiontypeid"`
	Actionname          string `json:"actionname"`
}

type MstactiontypeEntity struct {
	Id             int64  `json:"id"`
	Actiontypename string `json:"actiontypename"`
}

type MstactivityEntities struct {
	Total  int64               `json:"total"`
	Values []MstactivityEntity `json:"values"`
}

type MstactivityResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Details MstactivityEntities `json:"details"`
}

type MstactiontypeResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details []MstactiontypeEntity `json:"details"`
}
type MstactionResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details []MstactivitySingleEntity `json:"details"`
}

type MstactivityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstactivityEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
