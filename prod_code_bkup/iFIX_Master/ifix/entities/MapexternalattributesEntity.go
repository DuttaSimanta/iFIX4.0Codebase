package entities

import (
	"encoding/json"
	"io"
)

type MapexternalattributesEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
 	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Systemid            int64 `json:"systemid"`
	SystemName          string `json:"systemname"`
	Map                 []Attr `json:"map"`
	Extattr             string `json:"extattr"`
	Sysattr             string `json:"sysattr"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
}
type Attr struct {
	Extattr   string `json:"extattr"`
	Sysattr   string `json:"sysattr"`
}
type MappedattributesResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details []Attr `json:"details"`
}


type MapexternalattributesEntities struct {
	Total  int64            `json:"total"`
	Values []MapexternalattributesEntity `json:"values"`
}

type MapexternalattributesResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details MapexternalattributesEntities `json:"details"`
}

type MapexternalattributesResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MapexternalattributesEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
