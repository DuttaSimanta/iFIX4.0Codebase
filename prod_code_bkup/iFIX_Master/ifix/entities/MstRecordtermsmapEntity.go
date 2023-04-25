package entities

import (
	"encoding/json"
	"io"
)

type MstRecordtermsmapEntity struct {
	ID           int64   `json:"id"`
	FromclientID int64   `json:"fromclientid"`
	FromorgnID   int64   `json:"fromorgnid"`
	ToclinentID  int64   `json:"toclientid"`
	ToorgnID     []int64 `json:"toorgnid"`
	TermsSeq     []int64 `json:"termsseq"`
	MapID        int64   `json:"mapid"`
	Offset       int64   `json:"offset"`
	Limit        int64   `json:"limit"`
	Termname     string
	Termvalue    string
	Termtype     int64
}

type MstStatetermEntity struct {
	ClientID         int64
	OrgnID           int64
	RecorddifftypeID int64
	RecorddiffID     int64
	Recordtermvalue  string
	Iscompulsery     int64
}

//FromJSON is used for convert data into JSON format
func (p *MstRecordtermsmapEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
