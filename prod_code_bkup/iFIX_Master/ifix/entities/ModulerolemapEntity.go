package entities

import (
	"encoding/json"
	"io"
)

type ModulerolemapEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Moduleid            int64  `json:"moduleid"`
	Roleid              int64  `json:"roleid"`
	Menuid              int64  `json:"menuid"`
	Activeflg           int64  `json:"activeflg"`
	Audittransactionid  int64  `json:"audittransactionid"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Rolename            string `json:"rolename"`
	Modulename          string `json:"modulename"`
	Menuname            string `json:"menuname"`
}

type ModulerolemapEntities struct {
	Total  int64                 `json:"total"`
	Values []ModulerolemapEntity `json:"values"`
}

type ModulerolemapResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Details ModulerolemapEntities `json:"details"`
}

type ModulerolemapResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *ModulerolemapEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
