package entities

import (
	"encoding/json"
	"io"
)

type FuncmasterEntity struct{
	ID             int64    `json:"id"`
	Name	   string   `json:"name"`
}
type FuncmappingEntity struct{
	ID             		 int64    `json:"id"`
	Clientid             int64    `json:"clientid"`
	Mstorgnhirarchyid    int64    `json:"mstorgnhirarchyid"`
	Funcid            	 int64    `json:"funcid"`
	Funcdescid     		 int64    `json:"funcdescid"`
	Iscatalog     		 int64    `json:"iscatalog"`
	Description          string    `json:"description"`
	Seqno             	 int64    `json:"seqno"`
	Colorcode          	 string    `json:"colorcode"`
	Image          		 string    `json:"image"`
	Ismanegerialview     int64    `json:"ismanegerialview"`
	Offset 				 int64     `json:"offset"`
	Limit 				 int64     `json:"limit"`
}
type FuncmappingRespEntity struct{
	ID             		 int64    `json:"id"`
	Clientname             string    `json:"Clientname"`
	Orgname    			string    `json:"orgname"`
	Funcname            string    `json:"Funcname"`
	Description          string    `json:"description"`
	Seqno             	 int64    `json:"seqno"`
	Colorcode          	 string    `json:"colorcode"`
	Image          		 string    `json:"image"`
	Ismanegerialview     int64    `json:"ismanegerialview"`
	Activeflg 			 int64     `json:"activeflg"`
	Iscatalog     		 int64    `json:"iscatalog"`
}
type FuncmappingsingleRespEntity struct{
	Description          string    `json:"description"`
	Seqno             	 int64    `json:"seqno"`
	Funcdescid     		 int64    `json:"funcdescid"`
	Colorcode          	 string    `json:"colorcode"`
	Image          		 string    `json:"image"`
	Ismanegerialview     int64    `json:"ismanegerialview"`
	Iscatalog    		int64    `json:"iscatalog"`
	Activeflg 			 int64     `json:"activeflg"`
}

type FuncmappingsingleResponese struct{
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	[]FuncmappingsingleRespEntity `json:"details"`
}
type FuncmappingEntitities struct{
	Total int64 `json:"total"`
	Values []FuncmappingRespEntity  `json:"values"`
}

type FuncmasterRespEntity struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	[]FuncmasterEntity `json:"details"`
}
type FuncmasterResponse struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	FuncmappingEntitities `json:"details"`
}
func (w *FuncmappingEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}