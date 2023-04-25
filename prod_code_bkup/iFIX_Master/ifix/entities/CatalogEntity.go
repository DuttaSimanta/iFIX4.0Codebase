package entities

import (
	"encoding/json"
	"io"
)


type CatalogRecordResponse struct {
	Success bool                    `json:"success"`
	Message string                  `json:"message"`
	Details  RecordEntity       `json:"details"`
}
type CatalogTicketTypeEntity struct {
	TypeId 	int64    `json:"typeid"`
	Id 		int64          `json:"id"`
}
type RecordEntity struct {
	Catagories  []ParentCategoryEntity  `json:"catagories"`
	Ttype       CatalogTicketTypeEntity  `json:"ttype"`
	Catalog     ParentCategoryEntity    `json:"catalog"`
}
type ParentCategoryEntity struct {
	ID   string  `json:"id"`
	NAME string  `json:"name"`
	Torecorddiffid int64  `json:"torecorddiffid"`
}
type CatalogEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Fromrecorddifftypeid   int64  `json:"fromrecorddifftypeid"`
	Fromrecorddiffid   	int64  `json:"fromrecorddiffid"`
	Catalogname         string `json:"catalogname"`
	Activeflg           int64  `json:"activeflg"`
	Audittransactionid  int64  `json:"audittransactionid"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
}
type CatalogEntitySingle struct {
	Id                  int64  `json:"id"`
	Catalogname         string `json:"catalogname"`

}

type CatalogEntities struct {
	Total  int64           `json:"total"`
	Values []CatalogEntity `json:"values"`
}

type CatalogResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Details CatalogEntities `json:"details"`
}
type CatalogSingleResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Details []CatalogEntitySingle `json:"details"`
}

type CatalogResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *CatalogEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
