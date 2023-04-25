package entities

import (
	"encoding/json"
	"io"
)

type MstrecorddifftypeEntity struct {
	ID               int64  `json:"id"`
	Typename         string `json:"typename"`
	Parentname       string `json:"parentname"`
	Seqno            int64  `json:"seqno"`
	Istextfield      int64  `json:"istextfield"`
	Recorddifftypeid int64  `json:"recorddifftypeid"`
}
type RecordDiffEntity struct {
	ID                int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	Parentid          int64  `json:"parentid"`
	Recorddifftypeid  int64  `json:"recorddifftypeid"`
	Recorddiffid      int64  `json:"recorddiffid"`
	Name              string `json:"name"`
	Seqno             int64  `json:"seqno"`
	Offset            int64  `json:"offset"`
	Limit             int64  `json:"limit"`
}
type RecordDiffEntityResp struct {
	ID                int64  `json:"id"`
	Clientid          int64  `json:"clientid"`
	Mstorgnhirarchyid int64  `json:"mstorgnhirarchyid"`
	Recorddifftypeid  int64  `json:"recorddifftypeid"`
	Name              string `json:"name"`
	Seqno             int64  `json:"seqno"`
	Clientname        string `json:"clientname"`
	Orgname           string `json:"orgname"`
	Type              string `json:"Type"`
	Activeflg         int64  `json:"activeflg"`
}
type RecordDiffEntities struct {
	Total  int64                  `json:"total"`
	Values []RecordDiffEntityResp `json:"values"`
}
type RecordDiffTypeResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details []MstrecorddifftypeEntity `json:"details"`
}
type RecordDiffResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Details RecordDiffEntities `json:"details"`
}

func (w *RecordDiffEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
