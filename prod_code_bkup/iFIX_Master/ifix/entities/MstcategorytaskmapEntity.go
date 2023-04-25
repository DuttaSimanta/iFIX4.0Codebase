package entities

import (
	"encoding/json"
	"io"
)

type MstcategorytaskmapEntity struct {
	Id                int64 `json:"id"`
	Clientid          int64 `json:"clientid"`
	Mstorgnhirarchyid int64 `json:"mstorgnhirarchyid"`

	Fromtickettypedifftypeid int64 `json:"fromtickettypedifftypeid"`
	Fromtickettypediffid     int64 `json:"fromtickettypediffid"`
	Fromcatdifftypeid        int64 `json:"fromcatdifftypeid"`
	Fromcatlabelid           int64 `json:"fromcatlabelid"`
	Fromcatdiffid            int64 `json:"fromcatdiffid"`
	Totickettypedifftypeid   int64 `json:"totickettypedifftypeid"`
	Totickettypediffid       int64 `json:"totickettypediffid"`
	Tocatdifftypeid          int64 `json:"tocatdifftypeid"`
	Tocatlabelid             int64 `json:"tocatlabelid"`
	Tocatdiffid              int64 `json:"tocatdiffid"`

	Activeflg                  int64  `json:"activeflg"`
	Offset                     int64  `json:"offset"`
	Limit                      int64  `json:"limit"`
	Clientname                 string `json:"clientname"`
	Mstorgnhirarchyname        string `json:"mstorgnhirarchyname"`
	Fromtickettypedifftypename string `json:"fromtickettypedifftypename"`
	Fromtickettypediffname     string `json:"fromtickettypediffname"`
	Fromcatdifftypename        string `json:"fromcatdifftypename"`
	Fromcatlabelname           string `json:"fromcatlabelname"`
	Fromcatdiffname            string `json:"fromcatdiffname"`
	Totickettypedifftypename   string `json:"totickettypedifftypename"`
	Totickettypediffname       string `json:"totickettypediffname"`
	Tocatdifftypename          string `json:"tocatdifftypename"`
	Tocatlabelname             string `json:"tocatlabelname"`
	Tocatdiffnam               string `json:"tocatdiffnam"`
}

type MstcategorytaskmapEntities struct {
	Total  int64                      `json:"total"`
	Values []MstcategorytaskmapEntity `json:"values"`
}

type MstcategorytaskmapResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details MstcategorytaskmapEntities `json:"details"`
}

type MstcategorytaskmapResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstcategorytaskmapEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
