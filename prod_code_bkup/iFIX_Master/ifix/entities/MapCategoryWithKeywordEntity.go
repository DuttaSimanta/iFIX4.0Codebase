package entities

import (
	"encoding/json"
	"io"
)

type MapCategoryWithKeywordEntity struct {
	Id                int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	Keyword           string `json:"keyword"`
	Categoryvalue     string `json:"categoryvalue"`
	// Templatename        string `json:"templatename"`
	// Tableid             int64  `json:"tableid"`
	// Fieldid             int64  `json:"fieldid"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	// Tablename           string `json:"tablename"`
	// Fieldname           string `json:"fieldname"`
}
type Getkeyword struct {
	Id            int64  `json:"id"`
	Keyword       string `json:"keyword"`
	Categoryvalue string `json:"categoryvalue"`
}
type MapCategoryWithKeywordEntities struct {
	Total  int64                          `json:"total"`
	Values []MapCategoryWithKeywordEntity `json:"values"`
}
type GetkeywordResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details []Getkeyword `json:"details"`
}

type MapCategoryWithKeywordResponse struct {
	Success bool                           `json:"success"`
	Message string                         `json:"message"`
	Details MapCategoryWithKeywordEntities `json:"details"`
}

type MapCategoryWithKeywordResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MapCategoryWithKeywordEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
