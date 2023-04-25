package entities


import (
  "encoding/json"
  "io"
  )

type MstgroupnewEntity struct{
    Id           int64          `json:"id"`
    Clientid           int64          `json:"clientid"`
    Mstorgnhirarchyid           int64          `json:"mstorgnhirarchyid"`
    Processid           int64          `json:"processid"`
    Groupid             int64       `json:"groupid"`
    Groupname           string          `json:"groupname"`
    Activeflg           int64          `json:"activeflg"`
    Audittransactionid           int64          `json:"audittransactionid"`
    Offset              int64  `json:"offset"`
    Limit               int64  `json:"limit"`
}

type MstgroupnewEntities struct{
    Total int64     `json:"total"`
    Values []MstgroupEntity      `json:"values"`
}


type MstgroupnewResponse struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         MstgroupEntities     `json:"details"`
}


type MstgroupnewResponseInt struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         int64      `json:"details"`
}


func (w *MstgroupnewEntity) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(w)
}

