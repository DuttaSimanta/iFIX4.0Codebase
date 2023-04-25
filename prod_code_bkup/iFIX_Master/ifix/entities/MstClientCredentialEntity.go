package entities

import (
	"encoding/json"
	"io"
)

type MstClientCredentialEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
    Clientname          string `json:"clientname"`
    Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
    Credentialtypeid    int64 `json:"credentialtypeid"`
    Credentialtypename  string `json:"credentialtypename"`
    CredentialAccount   string `json:"credentialaccount"`
    CredentialPassword  string `json:"credentialpassword"`
    CredentialKey       string `json:"credentialkey"`
    CredentialEndPoint  string `json:"credentialendpoint"`
    DefaultConfig       int64  `json:"defaultconfig"`
    DefaultConfigName   string `json:"defaultconfigname"`
	 
 	Activeflg           int64  `json:"activeflg"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
 }

type MstClientCredentialEntities struct {
	Total  int64            `json:"total"`
	Values []MstClientCredentialEntity `json:"values"`
}

type MstClientCredentialResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details MstClientCredentialEntities `json:"details"`
}

type MstClientCredentialResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstClientCredentialEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
