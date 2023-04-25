package entities


import (
  "encoding/json"
  "io"
  )

type MstrecordfieldEntity struct{
    Id           		int64	`json:"id"`
    Clientid           		int64	`json:"clientid"`
    Mstorgnhirarchyid           int64	`json:"mstorgnhirarchyid"`
    Mstrecordfieldtype          string	`json:"mstrecordfieldtype"`
    Clientname          	string 	`json:"clientname"`
    Mstorgnhirarchyname 	string 	`json:"mstorgnhirarchyname"`
    Recordtermid           	int64	`json:"recordtermid"`
    Termname           		string  `json:"termname"`
    Termtypename           	string  `json:"termtypename"`
    Activeflg           	int64	`json:"activeflg"`
    Audittransactionid          int64	`json:"audittransactionid"`
    MstrecordfielddiffEntities  []MstrecordfielddiffEntity	`json:"mstrecordfielddiff"`
    Offset              	int64	`json:"offset"`
    Limit               	int64	`json:"limit"`
}

type MstrecordfielddiffEntity struct{
    Id           		int64	`json:"id"`
    Clientid           		int64	`json:"clientid"`
    Mstorgnhirarchyid           int64	`json:"mstorgnhirarchyid"`
    Mstrecordfieldid		int64	`json:"mstrecordfieldid"`
    Recorddifftypeid           	int64	`json:"recorddifftypeid"`
    Mstrecorddifferentiationtypename string `json:"mstrecorddifferentiationtypename"`
    Recorddiffid           	int64	`json:"recorddiffid"`
    Mstrecorddifferentiationname string `json:"mstrecorddifferentiationname"`
    Activeflg           	int64	`json:"activeflg"`
    Audittransactionid          int64	`json:"audittransactionid"`
    RecorddifftypeParentid      int64	`json:"recorddifftypeParentid"`
}

type MstrecordfieldEntities struct{
    Total int64     `json:"total"`
    Values []MstrecordfieldEntity      `json:"values"`
}


type MstrecordfieldResponse struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         MstrecordfieldEntities     `json:"details"`
}


type MstrecordfieldResponseInt struct {
    Success         bool     `json:"success"`
    Message         string     `json:"message"`
    Details         int64      `json:"details"`
}


func (w *MstrecordfieldEntity) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(w)
}


