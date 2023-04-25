package entities

import (
	"encoding/json"
	"io"
)

type FileuploadEntity struct {
	Id                 int64  `json:"id"`
	Clientid           int64  `json:"clientid"`
	Mstorgnhirarchyid  int64  `json:"mstorgnhirarchyid"`
	Credentialtype     string `json:"credentialtype"`
	Credentialaccount  string `json:"credentialaccount"`
	Credentialpassword string `json:"credentialpassword"`
	Credentialkey      string `json:"credentialkey"`
	Activeflg          int64  `json:"activeflg"`
	Originalfile       string `json:"originalfile"`
	Filename           string `json:"filename"`
	Path               string `json:"path"`
}

type FileuploadEntities struct {
	Values []FileuploadEntity `json:"values"`
}

type FileuploadResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details FileuploadEntity `json:"details"`
}

func (w *FileuploadEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
