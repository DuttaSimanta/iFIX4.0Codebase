package entities


import (
  "encoding/json"
  "io"
  )

type MststatetermEntity struct{
    Id           		int64	`json:"id"`
    Clientid           		int64	`json:"clientid"`
    Mstorgnhirarchyid           int64	`json:"mstorgnhirarchyid"`
    Clientname          	string 	`json:"clientname"`
    Mstorgnhirarchyname 	string 	`json:"mstorgnhirarchyname"`
    Recorddifftypeid           	int64	`json:"recorddifftypeid"`
    Mstrecorddifferentiationtypename string `json:"mstrecorddifferentiationtypename"`
    Recorddiffid           	int64	`json:"recorddiffid"`
    Mstrecorddifferentiationname string `json:"mstrecorddifferentiationname"`
    Recordtermid           	int64	`json:"recordtermid"`
    Termname           		string  `json:"termname"`
    Recordtermvalue           	string	`json:"recordtermvalue"`
    Iscompulsory           	int64	`json:"iscompulsory"`
    Activeflg           	int64	`json:"activeflg"`
    Audittransactionid          int64	`json:"audittransactionid"`
    Offset              	int64	`json:"offset"`
    Limit               	int64	`json:"limit"`
}

type MststatetermEntities struct{
    Total int64     `json:"total"`
    Values []MststatetermEntity      `json:"values"`
}


type MststatetermResponse struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         MststatetermEntities     `json:"details"`
}


type MststatetermResponseInt struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         int64      `json:"details"`
}


func (w *MststatetermEntity) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(w)
}


