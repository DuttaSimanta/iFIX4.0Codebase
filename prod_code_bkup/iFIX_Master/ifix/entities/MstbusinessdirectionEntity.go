package entities

import (
	"encoding/json"
	"io"
)

type MstbusinessdirectionEntity struct {
	Id                               int64  `json:"id"`
	Clientid                         int64  `json:"clientid"`
	Mstorgnhirarchyid                int64  `json:"mstorgnhirarchyid"`
	Mstrecorddifferentiationtypeid   int64  `json:"mstrecorddifferentiationtypeid"`
	Mstrecorddifferentiationid       int64  `json:"mstrecorddifferentiationid"`
	Direction                        int64  `json:"direction"`
	Activeflg                        int64  `json:"activeflg"`
	Offset                           int64  `json:"offset"`
	Limit                            int64  `json:"limit"`
	Clientname                       string `json:"clientname"`
	Mstorgnhirarchyname              string `json:"mstorgnhirarchyname"`
	Mstrecorddifferentiationtypename string `json:"mstrecorddifferentiationtypename"`
	Mstrecorddifferentiationname     string `json:"mstrecorddifferentiationname"`
	Baseconfig     					 int64 `json:"baseconfig"`
}

type MstbusinessdirectionEntities struct {
	Total  int64                        `json:"total"`
	Values []MstbusinessdirectionEntity `json:"values"`
}

type MstbusinessdirectionResponse struct {
	Success bool                         `json:"success"`
	Message string                       `json:"message"`
	Details MstbusinessdirectionEntities `json:"details"`
}

type MstbusinessdirectionResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstbusinessdirectionEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
