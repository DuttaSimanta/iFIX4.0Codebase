package entities

import (
	"encoding/json"
	"io"
)

type MstslastartendcriteriaEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Workflowid          int64  `json:"workflowid"`
	Stateid             int64  `json:"stateid"`
	Statetypeid         int64  `json:"statetypeid"`
	Slaid               int64  `json:"slaid"`
	Startorend          int64  `json:"startorend"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Workflowname        string `json:"workflowname"`
	Slaname             string `json:"slaname"`
	Statename           string `json:"statename"`
}

type MstslanameagaistworkflowEntity struct {
	Id      int64  `json:"id"`
	Slaname string `json:"slaname"`
}

type MstslastartendcriteriaEntities struct {
	Total  int64                          `json:"total"`
	Values []MstslastartendcriteriaEntity `json:"values"`
}

type MstslastartendcriteriaResponse struct {
	Success bool                           `json:"success"`
	Message string                         `json:"message"`
	Details MstslastartendcriteriaEntities `json:"details"`
}

type MstslanameagaistworkflowEntityResponse struct {
	Success bool                             `json:"success"`
	Message string                           `json:"message"`
	Details []MstslanameagaistworkflowEntity `json:"details"`
}

type MstslastartendcriteriaResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstslastartendcriteriaEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
