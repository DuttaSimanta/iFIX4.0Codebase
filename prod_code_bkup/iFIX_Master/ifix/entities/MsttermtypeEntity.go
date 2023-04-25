package entities


import (
  "encoding/json"
  "io"
  )

type MsttermtypeEntity struct{
    Id           	int64	`json:"id"`
    Termtypename        string  `json:"termtypename"`
    Offset              int64  	`json:"offset"`
    Limit               int64  	`json:"limit"`
}

type MsttermtypeEntities struct{
    Total int64     	`json:"total"`
    Values []MsttermtypeEntity	`json:"values"`
}


type MsttermtypeResponse struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         MsttermtypeEntities     `json:"details"`
}


type MsttermtypeResponseInt struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         int64      `json:"details"`
}


func (w *MsttermtypeEntity) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(w)
}


