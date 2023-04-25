package entities


import (
  "encoding/json"
  "io"
  )

type MstgroupEntity struct{
    Id           int64          `json:"id"`
    Clientid           int64          `json:"clientid"`
    Mstorgnhirarchyid           int64          `json:"mstorgnhirarchyid"`
    Processid           int64          `json:"processid"`
    Groupname           string          `json:"groupname"`
    Activeflg           int64          `json:"activeflg"`
    Audittransactionid           int64          `json:"audittransactionid"`
    Offset              int64  `json:"offset"`
    Limit               int64  `json:"limit"`
}

type MstgroupEntities struct{
    Total int64     `json:"total"`
    Values []MstgroupEntity      `json:"values"`
}


type MstgroupResponse struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         MstgroupEntities     `json:"details"`
}


type MstgroupResponseInt struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         int64      `json:"details"`
}


func (w *MstgroupEntity) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(w)
}


