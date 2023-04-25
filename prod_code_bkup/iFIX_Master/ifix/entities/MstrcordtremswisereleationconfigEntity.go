package entities

import (
	"encoding/json"
	"io"
)

type MstrcordtremswisereleationconfigEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Recorddifftypeid    int64  `json:"recorddifftypeid"`
	Recorddiffid        int64  `json:"recorddiffid"`
	Releationid         int64  `json:"releationid"`
	Termsid             int64  `json:"termsid"`
	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Releationname       string `json:"releationname"`
	Termname            string `json:"termname"`
	Recorddifftypename  string `json:"recorddifftypename"`
	Recorddiffname      string `json:"recorddiffname"`
}

type Recordreleationdetails struct {
	ID            int64  `json:"id"`
	Releationname string `json:"releationname"`
}

type Recordtermnames struct {
	ID    int64  `json:"id"`
	Names string `json:"releationname"`
}

type MstrcordtremswisereleationconfigEntities struct {
	Total  int64                                    `json:"total"`
	Values []MstrcordtremswisereleationconfigEntity `json:"values"`
}

type MstrcordtremswisereleationconfigResponse struct {
	Success bool                                     `json:"success"`
	Message string                                   `json:"message"`
	Details MstrcordtremswisereleationconfigEntities `json:"details"`
}

type MstrcordtremswisereleationconfigResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

type RecordreleationdetailsAllResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message"`
	Details []Recordreleationdetails `json:"details"`
}

type RecordtermnamesAllResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Details []Recordtermnames `json:"details"`
}

func (w *MstrcordtremswisereleationconfigEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
