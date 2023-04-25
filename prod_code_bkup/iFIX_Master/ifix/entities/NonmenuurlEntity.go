package entities

import (
	"encoding/json"
	"io"
)

type NonmenuurlEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	UrlId               int64  `json:"urlid"`
	Url                 string `json:"url"`
	Deleteflg           int64  `json:"deleteflg"`
	Activeflg           int64  `json:"activeflg"`
	Audittransactionid  int64  `json:"audittransactionid"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Urlname             string `json:"Urlname"`
	Type                string `json:"type"`
	Code                string `json:"code"`
}

type NonmenuurlEntities struct {
	Total  int64              `json:"total"`
	Values []NonmenuurlEntity `json:"values"`
}

type NonmenuurlResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details NonmenuurlEntities `json:"details"`
}
type NonmenuurlsingleResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details []NonmenuurlEntity `json:"details"`
}

type NonmenuurlResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *NonmenuurlEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
