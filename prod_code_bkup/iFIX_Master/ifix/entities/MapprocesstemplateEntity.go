package entities

import (
	"encoding/json"
	"io"
)

type MapprocesstemplateEntity struct{
	Id                  int64    `json:"id"`
	Clientid            int64    `json:"clientid"`
	Mstorgnhirarchyid   int64    `json:"mstorgnhirarchyid"`
	Loggedinmstorgnhirarchyid   int64    `json:"loggedinmstorgnhirarchyid"`
	Processid           int64    `json:"processid"`
	Recorddiffids       []RecorddiffEntity  `json:"recorddiffids"`
	Templatetransitionid        int64    `json:"templatetransitionid"`
}
type RecorddiffEntity struct{
	Type                 int64    `json:"type"`
	Id                  int64    `json:"id"`
}

type ProcessdetailsEntity struct{
	Source                 []string    `json:"source"`
	Node                   string     `json:"node"`
	Targets                []int64    `json:"targets"`
	Issave                 bool       `json:"isSave"`
	Instate                []int64    `json:"inState"`
	OutState                []int64    `json:"outState"`
}

func (w *MapprocesstemplateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
