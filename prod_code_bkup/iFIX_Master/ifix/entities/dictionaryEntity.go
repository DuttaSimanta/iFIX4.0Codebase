package entities

import (
	"encoding/json"
	"io"
)

type DictionaryEntity struct {
	Id                  	int64    `json:"id"`
	Clientid            	int64    `json:"clientid"`
	Mstorgnhirarchyid   	int64    `json:"mstorgnhirarchyid"`
	Databasename        	string   `json:"databasename"`
	Databaseuser        	string   `json:"databaseuser"`
	Databasepassword    	string   `json:"databasepassword"`
	Tablename           	string   `json:"tablename"`
	Mstdatadictionarydbid   int64    `json:"mstdatadictionarydbid"`
	Tableid   				int64    `json:"tableid"`
	Columnname    			string   `json:"columnname"`
	Columntype           	string   `json:"columntype"`
	Isnull   				int64    `json:"isnull"`
	Audittransactionid   	int64    `json:"audittransactionid"`
}

func (w *DictionaryEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
