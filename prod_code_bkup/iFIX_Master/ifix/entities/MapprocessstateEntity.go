package entities

import (
	"encoding/json"
	"io"
)

type MapprocessstateEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Statetid            int64  `json:"statetid"`
	Statetypeid         int64  `json:"statetypeid"`
	Processid           int64  `json:"processid"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Seqno               int64  `json:"seqno"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Processname         string `json:"processname"`
	Statename           string `json:"statename"`
}

type MapprocessstateEntities struct {
	Total  int64                   `json:"total"`
	Values []MapprocessstateEntity `json:"values"`
}

type MapprocessstateResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details MapprocessstateEntities `json:"details"`
}

type MapprocessstateResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MapprocessstateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
