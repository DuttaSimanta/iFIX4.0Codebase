package entities

import (
	"encoding/json"
	"io"
)

type CountryEntity struct {
	Id          int64  `json:"id"`
	Countryname string `json:"countryname"`
}

type CountryEntities struct {
	Total  int64           `json:"total"`
	Values []CountryEntity `json:"values"`
}

type CountryResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Details CountryEntities `json:"details"`
}

type CountryResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *CountryEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
