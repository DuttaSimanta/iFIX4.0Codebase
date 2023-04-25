package entities

import (
	"encoding/json"
	"io"
)

type MstprocessrecordmapEntity struct {
	Id                 int64 `json:"id"`
	Clientid           int64 `json:"clientid"`
	Mstorgnhirarchyid  int64 `json:"mstorgnhirarchyid"`
	Recorddifftypeid   int64 `json:"recorddifftypeid"`
	Recorddiffid       int64 `json:"recorddiffid"`
	Mstprocessid       int64 `json:"mstprocessid"`
	Activeflg          int64 `json:"activeflg"`
	Audittransactionid int64 `json:"audittransactionid"`
	Offset             int64 `json:"offset"`
	Limit              int64 `json:"limit"`
}

type MstprocessrecordmapEntities struct {
	Total  int64                       `json:"total"`
	Values []MstprocessrecordmapEntity `json:"values"`
}

type MstprocessrecordmapResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Details MstprocessrecordmapEntities `json:"details"`
}

type MstprocessrecordmapResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstprocessrecordmapEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
