package entities

import (
	"encoding/json"
	"io"
)

type TransporttableEntity struct {
	Id                   int64  `json:"id"`
	Msttablename         string `json:"msttablename"`
	Tabletype            int64  `json:"tabletype"`
	Tabletypedescription string `json:"tabletypedescription"`
	Activeflg            int64  `json:"activeflg"`
	Offset               int64  `json:"offset"`
	Limit                int64  `json:"limit"`
}
type TableEntity struct {
	Tablename string `json:"tablename"`
	Tabletype int64  `json:"tabletype"`
}
type GettableEntity struct {
	Tabletype            int64  `json:"tabletype"`
	Tabletypedescription string `json:"tabletypedescription"`
}
type TransporttableEntities struct {
	Total  int64                  `json:"total"`
	Values []TransporttableEntity `json:"values"`
}

type TransporttableResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details TransporttableEntities `json:"details"`
}
type TransporttabletypeResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details []GettableEntity `json:"details"`
}
type TableResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Details []TableEntity `json:"details"`
}
type TransporttableResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *TransporttableEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
