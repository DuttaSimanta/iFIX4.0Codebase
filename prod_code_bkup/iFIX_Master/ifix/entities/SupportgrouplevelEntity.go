package entities

import (
	"encoding/json"
	"io"
)

type SupportgrouplevelEntity struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type SupportgrouplevelEntities struct {
	Total  int64                     `json:"total"`
	Values []SupportgrouplevelEntity `json:"values"`
}

type SupportgrouplevelResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details SupportgrouplevelEntities `json:"details"`
}

type SupportgrouplevelResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *SupportgrouplevelEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
