package entities

import (
	"encoding/json"
	"io"
)

type MapprocesstoentityEntity struct {
	Id                       int64 `json:"id"`
	Clientid                 int64 `json:"clientid"`
	Mstorgnhirarchyid        int64 `json:"mstorgnhirarchyid"`
	Mstprocessid             int64 `json:"mstprocessid"`
	Mstdatadictionaryfieldid int64 `json:"mstdatadictionaryfieldid"`
	Activeflg                int64 `json:"activeflg"`
	Offset                   int64 `json:"offset"`
	Limit                    int64 `json:"limit"`
}

type MapprocesstoentityEntities struct {
	Total  int64                      `json:"total"`
	Values []MapprocesstoentityEntity `json:"values"`
}

type MapprocesstoentityResponse struct {
	Success bool                       `json:"success"`
	Message string                     `json:"message"`
	Details MapprocesstoentityEntities `json:"details"`
}

type MapprocesstoentityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MapprocesstoentityEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
