package entities


import (
  "encoding/json"
  "io"
  )

type ErrormesgEntity struct{
    Id           int64          `json:"id"`
    Errorcode           string          `json:"errorcode"`
    Errormsg           string          `json:"errormsg"`
    Offset              int64  `json:"offset"`
    Limit               int64  `json:"limit"`
}

type ErrormesgEntities struct{
    Total int64     `json:"total"`
    Values []ErrormesgEntity      `json:"values"`
}


type ErrormesgResponse struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         ErrormesgEntities     `json:"details"`
}


type ErrormesgResponseInt struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         int64      `json:"details"`
}


func (w *ErrormesgEntity) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(w)
}


