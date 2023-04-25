package entities

import (
	"encoding/json"
	"io"
)

type CityEntity struct {
	Id       int64  `json:"id"`
	Cityname string `json:"cityname"`
}

type CityEntities struct {
	Total  int64        `json:"total"`
	Values []CityEntity `json:"values"`
}

type CityResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Details CityEntities `json:"details"`
}

type CityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *CityEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
