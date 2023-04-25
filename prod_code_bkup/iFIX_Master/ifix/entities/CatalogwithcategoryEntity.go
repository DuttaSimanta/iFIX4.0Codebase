package entities

import (
	"encoding/json"
	"io"
)

type CatalogwithcategoryEntity struct {
	Id                     int64  `json:"id"`
	Clientid               int64  `json:"clientid"`
	Mstorgnhirarchyid      int64  `json:"mstorgnhirarchyid"`
	Fromrecorddifftypeid   int64  `json:"fromrecorddifftypeid"`
	Fromrecorddiffid       int64  `json:"fromrecorddiffid"`
	Torecorddifftypeid     int64  `json:"torecorddifftypeid"`
	Torecorddiffid         int64  `json:"torecorddiffid"`
	Forrecorddiffid        int64  `json:"forrecorddiffid"`
	Catalogid              int64  `json:"catalogid"`
	Activeflg              int64  `json:"activeflg"`
	Audittransactionid     int64  `json:"audittransactionid"`
	Offset                 int64  `json:"offset"`
	Limit                  int64  `json:"limit"`
	Clientname             string `json:"clientname"`
	Mstorgnhirarchyname    string `json:"mstorgnhirarchyname"`
	Fromrecorddifftypename string `json:"fromrecorddifftypename"`
	Fromrecorddiffname     string `json:"fromrecorddiffname"`
	Torecorddifftypename   string `json:"torecorddifftypename"`
	Torecorddiffname       string `json:"torecorddiffname"`
	Catalogname            string `json:"catalogname"`
	Parentname            string `json:"parentname"`
	Torecorddiffids      []int64  `json:"torecorddiffids"`
}

type CatalogwithsingleEntity struct{
	Torecorddifftypeid     int64  `json:"typeid"`
	Fromrecorddifftypeid   int64  `json:"fromrecorddifftypeid"`
	Fromrecorddiffid       int64  `json:"fromrecorddiffid"`
	Torecorddiffid         int64  `json:"id"`
	Torecorddiffname       string `json:"title"`
	Parentpath       string `json:"parentpath"`
	Seqno                  int64 `json:"seqno"`
}
type CatalogwithcategoryEntities struct {
	Total  int64                       `json:"total"`
	Values []CatalogwithcategoryEntity `json:"values"`
}

type CatalogwithcategoryResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details CatalogwithcategoryEntities `json:"details"`
}
type CatalogwithcategorysingleResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details []CatalogwithsingleEntity `json:"details"`
}

type CatalogwithcategoryResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *CatalogwithcategoryEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
