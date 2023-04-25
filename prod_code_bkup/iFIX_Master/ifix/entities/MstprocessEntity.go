package entities

import (
	"encoding/json"
	"io"
)

type MstprocessEntity struct {
	Id                         int64  `json:"id"`
	Clientid                   int64  `json:"clientid"`
	Mstorgnhirarchyid          int64  `json:"mstorgnhirarchyid"`
	Processname                string `json:"processname"`
	Activeflg                  int64  `json:"activeflg"`
	Offset                     int64  `json:"offset"`
	Limit                      int64  `json:"limit"`
	Recorddifftypeid           int64  `json:"recorddifftypeid"`
	Recorddiffid               int64  `json:"recorddiffid"`
	Mstdatadictionaryfieldid   int64  `json:"mstdatadictionaryfieldid"`
	Mstprocesstoentityid       int64  `json:"mstprocesstoentityid"`
	Mstprocessrecordmapid      int64  `json:"mstprocessrecordmapid"`
	Clientname                 string `json:"clientname"`
	Mstorgnhirarchyname        string `json:"mstorgnhirarchyname"`
	Recorddifftypname          string `json:"recorddifftypname"`
	Recorddiffname             string `json:"recorddiffname"`
	Mstdatadictionaryfieldname string `json:"mstdatadictionaryfieldname"`
	Mstdatadictionarytablename string `json:"mstdatadictionarytablename"`
	Forrecorddifftypeid        int64  `json:"forrecorddifftypeid"`
	Forrecorddiffid            int64  `json:"forrecorddiffid"`
	Tableid                    int64  `json:"tableid"`
	Mstdatadictionarydbid      int64  `json:"mstdatadictionarydbid"`
	Catname                    string
	Parentcatname              string
}
type MstprocessEntities struct {
	Total  int64              `json:"total"`
	Values []MstprocessEntity `json:"values"`
}

type MstprocessResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details MstprocessEntities `json:"details"`
}

type MstprocessResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstprocessEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
