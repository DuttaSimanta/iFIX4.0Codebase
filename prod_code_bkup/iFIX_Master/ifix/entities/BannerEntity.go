package entities

import (
	"encoding/json"
	"io"
)

type BannerEntity struct {
	Id                  int64  `json:"id"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Mstorgnhirarchyname string `json:mstorgnhirarchyname`
	Groupid             []int64 `json:"groupid"`
	Groupname           string `json:groupname`
	Message             string `json:"message"`
	ActualStarttime     string `json:actualstarttime`
	ActualEndtime       string `json:actualendtime`
	Starttime           int64  `json:starttime`
	Endtime             int64  `json:endtime`
	Sequence            int64  `json:sequence`
	Color               string `json:"color"`
	Size                int64   `json:"size"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
}
type BannerMessageEntity struct{
	Message  string   `json:"message"`
	Color               string   `json:"color"`
	Size                int64   `json:"size"`
}

type BannerEntities struct {
	Total  int64            `json:"total"`
	Values []BannerEntity `json:"values"`
}

type BannerResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details BannerEntities `json:"details"`
}

type BannerResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

type BannerResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details []BannerMessageEntity  `json:"details"`
}

func (w *BannerEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
