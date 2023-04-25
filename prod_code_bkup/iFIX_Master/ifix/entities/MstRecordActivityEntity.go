package entities

import (
	"encoding/json"
	"io"
)

type MstRecordActivityEntity struct {
	Id                   int64    `json:"id"`
	Clientid             int64    `json:"clientid"`
	ToClientid           int64    `json:"toclientid"`
	Clientname           string   `json:"clientname"`
	Mstorgnhirarchyid    int64    `json:"mstorgnhirarchyid"`
	ToMstorgnhirarchyids []int64  `json:"tomstorgnhirarchyids"`
	Mstorgnhirarchyname  string   `json:"mstorgnhirarchyname"`
	Activitydesc         string   `json:"activitydesc"`
	Activitydesces       []string `json:"activitydesces"`
	Sequence             int64    `json:"sequence"`
	Activeflg            int64    `json:"activeglg"`
	Offset               int64    `json:"offset"`
	Limit                int64    `json:"limit"`
}
type Activitydesces struct {
	Activitydesc string `json:"activitydesc"`
}
type ActivitydescesResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details []Activitydesces `json:"details"`
}
type MstRecordActivityEntities struct {
	Total  int64                     `json:"total"`
	Values []MstRecordActivityEntity `json:"values"`
}

type MstRecordActivityResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details MstRecordActivityEntities `json:"details"`
}

type MstRecordActivityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstRecordActivityEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
