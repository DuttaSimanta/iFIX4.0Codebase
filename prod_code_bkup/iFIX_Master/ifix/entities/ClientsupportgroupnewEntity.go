package entities

import (
	"encoding/json"
	"io"
)

type ClientsupportgroupnewEntity struct {
	Id                    int64   `json:"id"`
	Clientid              int64   `json:"clientid"`
	Mstorgnhirarchyid     int64   `json:"mstorgnhirarchyid"`
	Userid     int64   `json:"userid"`
	Mstorgnhirarchyids    []int64 `json:"mstorgnhirarchyids"`
	Supportgroupid        int64   `json:"supportgroupid"`
	Supportgroupname      string  `json:"supportgroupname"`
	Supportgrouplevelid   int64   `json:"supportgrouplevelid"`
	Mstclienttimezoneid   int64   `json:"mstclienttimezoneid"`
	Reporttimezoneid      int64   `json:"reporttimezoneid"`
	Email                 string  `json:"email"`
	Activeflg             int64   `json:"activeflg"`
	Offset                int64   `json:"offset"`
	Limit                 int64   `json:"limit"`
	Clientname            string  `json:"clientname"`
	Mstorgnhirarchyname   string  `json:"mstorgnhirarchyname"`
	Supportgrplevelname   string  `json:"supportgrplevelname"`
	Timezonename          string  `json:"timezonename"`
	Reporttimezonename    string  `json:"reporttimezonename"`
	Isworkflow            string  `json:"isworkflow"`
	Hascatalog            string  `json:"hascatalog"`
	Externalgrpid         int64   `json:"externalgrpid"`
	FromClientid          int64   `json:"fromclientid"`
	FromMstorgnhirarchyid int64   `json:"frommstorgnhirarchyid"`
	FromGroupids          []int64 `json:"fromgroupids"`
	ToClientid            int64   `json:"toclientid"`
	ToMstorgnhirarchyids  []int64 `json:"tomstorgnhirarchyids"`
	IsManagement          string  `json:"ismanagement"`
}
type GetsupportgroupbyorgEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Groupid             int64  `json:"groupid"`
	Groupname           string `json:"groupname"`
	Levelid           int64 `json:"levelid"`
	Hascatalog            string  `json:"hascatalog"`
}
type ClientsupportgroupnewsingleEntity struct {
	Id               int64  `json:"id"`
	Levelid          int64  `json:"levelid"`
	Supportgroupname string `json:"supportgroupname"`
}

type ClientsupportgroupnewsingleResponse struct {
	Success bool                                `json:"success"`
	Message string                              `json:"message"`
	Details []ClientsupportgroupnewsingleEntity `json:"details"`
}
type GetsupportgroupResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Details []GetsupportgroupbyorgEntity `json:"details"`
}

type ClientsupportgroupnewEntities struct {
	Total  int64                         `json:"total"`
	Values []ClientsupportgroupnewEntity `json:"values"`
}

type ClientsupportgroupnewResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Details ClientsupportgroupnewEntities `json:"details"`
}

type ClientsupportgroupnewResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *ClientsupportgroupnewEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
