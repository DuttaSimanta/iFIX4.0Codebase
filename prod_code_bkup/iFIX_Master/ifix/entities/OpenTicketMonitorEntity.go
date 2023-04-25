package entities

import (
	"encoding/json"
	"io"
)

type OpenTicketEntity struct {
	Id          int64  `json:"id"`
	Groupid     int64  `json:"groupid"`
	Groupname   string `json:groupname`
	Recordid    int64  `json:"recordid"`
	Reccordcode string `json:reccordcode`
	Userid      int64  `json:"userid"`
	Username    string `json:username`
	Opendate    string `json:opendate`
	Offset      int64  `json:"offset"`
	Limit       int64  `json:"limit"`
}

type OpenTicketEntities struct {
	Total  int64              `json:"total"`
	Values []OpenTicketEntity `json:"values"`
}

type OpenTicketResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details OpenTicketEntities `json:"details"`
}

type OpenTicketResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *OpenTicketEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
