package entities

import (
	"encoding/json"
	"io"
)

type RecordfulTicketIdDeleteFlagUpdateEntity struct {
	Recordid  int64  `json:"recordid"`
	Ticketid  string `json:"ticketid"`
	Deleteflg int64  `json:"deleteflg"`
}

type TrnRecordCodeDeleteFlgUpdateByIdEntity struct {
	Id        int64  `json:"id"`
	Code      string `json:"code"`
	Deleteflg int64  `json:"deleteflg"`
}

func (w *RecordfulTicketIdDeleteFlagUpdateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *TrnRecordCodeDeleteFlgUpdateByIdEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
