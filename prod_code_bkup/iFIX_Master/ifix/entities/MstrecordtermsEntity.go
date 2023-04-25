package entities

import (
	"encoding/json"
	"io"
)

type MstrecordtermsEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Termname            string `json:"termname"`
	Termtypeid          int64  `json:"termtypeid"`
	Termtypename        string `json:"termtypename"`
	Termvalue           string `json:"termvalue"`
	Termseq             int64 `json:"termseq"`
	Activeflg           int64  `json:"activeflg"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Audittransactionid  int64  `json:"audittransactionid"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
}

type TermsEntity struct {
	Id           int64  `json:"id"`
	Termname     string `json:"termname"`
	Termtypeid   int64  `json:"termtypeid"`
	Termtypename string `json:"termtypename"`
	Termvalue    string `json:"termvalue"`
	Activeflg    int64  `json:"activeflg"`
	Termseq      int64 `json:"termseq"`
}

type MstrecordtermsEntities struct {
	Total  int64                  `json:"total"`
	Values []MstrecordtermsEntity `json:"values"`
}

type MstrecordtermsResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details MstrecordtermsEntities `json:"details"`
}

type TermsResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Details []TermsEntity `json:"details"`
}

type MstrecordtermsResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstrecordtermsEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
