package entities

import (
	"encoding/json"
	"io"
)

//id, clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, prefix, year, month, day, configurezero, isclient
type RecordConfigIncrementEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	ClientName          string `json:"clientname"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:mstorgnhirarchyname`
	Recorddifftypeid    int64  `json:"recorddifftypeid"`
	RecorddifftypeName  string `json:"recorddifftypename"`
	Recorddiffid        int64  `json:"recorddiffid"`
	RecorddiffName      string `json:"recorddiffname"`
	Prefix              string `json:"prefix"`
	Year                string `json:"year"`
	Month               string `json:"month"`
	Day                 string `json:"day"`
	Configurezero       string `json:"configurezero"`
	IsClient            int64  `json:"isclient"`
	Number              int64  `json:"number"`

	// Groupid             []int64 `json:"groupid"`
	// Groupname           string `json:groupname`
	// Message             string `json:"message"`
	// ActualStarttime     string `json:actualstarttime`
	// ActualEndtime       string `json:actualendtime`
	// Starttime           int64  `json:starttime`
	// Endtime             int64  `json:endtime`
	// Sequence            int64  `json:sequence`
	// Color               string `json:"color"`
	// Size                int64   `json:"size"`
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type RecordConfigIncrementEntities struct {
	Total  int64                         `json:"total"`
	Values []RecordConfigIncrementEntity `json:"values"`
}

type RecordConfigIncrementResponse struct {
	Success bool                          `json:"success"`
	Message string                        `json:"message"`
	Details RecordConfigIncrementEntities `json:"details"`
}

type RecordConfigIncrementResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *RecordConfigIncrementEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
