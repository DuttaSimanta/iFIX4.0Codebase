package entities

import (
	"encoding/json"
	"io"
)

type MstorgcodeEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Toolcode            string `json:"toolcode"`
	Orgcode             string `json:"orgcode"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}

type Gettoolscode struct {
	Id       int64  `json:"id"`
	Toolcode string `json:"toolcode"`
}
type Getorgcode struct {
	Id      int64  `json:"id"`
	Orgcode string `json:"orgcode"`
}

type MstorgcodeEntities struct {
	Total  int64              `json:"total"`
	Values []MstorgcodeEntity `json:"values"`
}

type MstorgcodeResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details MstorgcodeEntities `json:"details"`
}

type MstorgcodeResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

type GettoolsResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Details []Gettoolscode `json:"details"`
}

type GetorgResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details []Getorgcode `json:"details"`
}

func (w *MstorgcodeEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
