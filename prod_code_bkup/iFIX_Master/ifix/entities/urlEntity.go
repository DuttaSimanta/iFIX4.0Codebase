package entities

import (
	"encoding/json"
	"io"
)

type UrlEntity struct{
	Id        		   int64    `json:"id"`
	Clientid       	   int64    `json:"clientid"`
	Mstorgnhirarchyid  int64    `json:"mstorgnhirarchyid"`
	Moduleid           int64    `json:"moduleid"`
	OldUrl             int64    `json:"oldUrl"`
	Urlkey         	   string   `json:"urlkey"`
	Url  			   string   `json:"url"`
	Urldescription     string   `json:"urldescription"`
	Type     		   string   `json:"type"`
}
type UrlRespEntity struct{
	Id        		   int64    `json:"id"`
	Urlkey         	   string   `json:"urlkey"`
	Url  			   string   `json:"url"`
	Urldescription     string   `json:"urldescription"`
}
type ModuleUrlEntity struct{
	Id        		   int64    `json:"id"`
	Modulename         string   `json:"modulename"`
	Url  			   string   `json:"url"`
	Urlkey         	   string   `json:"urlkey"`
	Clientname         	   string   `json:"clientname"`
	Orgname         	   string   `json:"orgname"`

}

type UrlEntities struct{
	Total int64 `json:"total"`
	Values []UrlRespEntity  `json:"values"`
}
type ModuleUrlEntities struct{
	Total int64 `json:"total"`
	Values []ModuleUrlEntity  `json:"values"`
}

type UrlResponseOnly struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	[]UrlRespEntity `json:"details"`
}
type UrlResponse struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	UrlEntities `json:"details"`
}

type UrlResponseInt struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	int64 `json:"details"`
}
type ModuleUrlResponse struct {
	Success  	bool `json:"success"`
	Message 	string `json:"message"`
	Details 	ModuleUrlEntities `json:"details"`
}

func (w *UrlEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}

func (w *ModuleUrlEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}