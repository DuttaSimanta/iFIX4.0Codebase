package entities

import (
	"encoding/json"
	"io"
)

type MststateEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Statetypeid         int64  `json:"statetypeid"`
	Statename           string `json:"statename"`
	Description         string `json:"description"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Seqno               int64  `json:"seqno"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Statetypename       string `json:"statetypename"`
}

type MststateEntities struct {
	Total  int64            `json:"total"`
	Values []MststateEntity `json:"values"`
}

type MststateResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details MststateEntities `json:"details"`
}

type MststateResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MststateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
