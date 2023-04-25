package entities

import (
	"encoding/json"
	"io"
)

type MstslaentityEntity struct {
	Id                              int64  `json:"id"`
	Clientid                        int64  `json:"clientid"`
	Mstorgnhirarchyid               int64  `json:"mstorgnhirarchyid"`
	Slaid                           int64  `json:"slaid"`
	Associatedmstclienttableid      int64  `json:"associatedmstclienttableid"`
	Associatedmstclienttablefieldid int64  `json:"associatedmstclienttablefieldid"`
	Activeflg                       int64  `json:"activeflg"`
	Offset                          int64  `json:"offset"`
	Limit                           int64  `json:"limit"`
	Clientname                      string `json:"clientname"`
	Mstorgnhirarchyname             string `json:"mstorgnhirarchyname"`
	Tablename                       string `json:"tablename"`
	Fieldname                       string `json:"fieldname"`
	Dbid                            int64  `json:"dbid"`
	Slaname                         string `json:"Slaname"`
}

type MstslaentityEntities struct {
	Total  int64                `json:"total"`
	Values []MstslaentityEntity `json:"values"`
}

type MstslaentityResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details MstslaentityEntities `json:"details"`
}

type MstslaentityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstslaentityEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
