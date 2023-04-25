package entities

import (
	"encoding/json"
	"io"
)

type ClientsupportgroupEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyids   string  `json:"mstorgnhirarchyids"`
	Supportgroupname    string `json:"supportgroupname"`
	Supportgrouplevelid int64  `json:"supportgrouplevelid"`
	Mstclienttimezoneid int64  `json:"mstclienttimezoneid"`
	Reporttimezoneid    int64  `json:"reporttimezoneid"`
	Email               string `json:"email"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Supportgrplevelname string `json:"supportgrplevelname"`
	Timezonename        string `json:"timezonename"`
	Reporttimezonename  string `json:"reporttimezonename"`
	Isworkflow          string `json:"isworkflow"`
	Hascatalog          string `json:"hascatalog"`
	Externalgrpid       int64  `json:"externalgrpid"`

}
type ClientsupportgroupsingleEntity struct {
	Id                  int64  `json:"id"`
	Levelid             int64  `json:"levelid"`
	Supportgroupname    string `json:"supportgroupname"`
	Groupname    		string `json:"groupname"`
	Defaultgroup        int     `json:"defaultgroup"`
}

type ClientsupportgroupsingleResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details []ClientsupportgroupsingleEntity `json:"details"`
}

type ClientsupportgroupEntities struct {
	Total  int64                      `json:"total"`
	Values []ClientsupportgroupEntity `json:"values"`
}

type ClientsupportgroupResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details ClientsupportgroupEntities `json:"details"`
}

type ClientsupportgroupResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *ClientsupportgroupEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
