package entities

import (
	"encoding/json"
	"io"
)

type MaprecordstatetodifferentiationEntity struct {
	Id                            int64  `json:"id"`
	Clientid                      int64  `json:"clientid"`
	Mstorgnhirarchyid             int64  `json:"mstorgnhirarchyid"`
	Recorddifftypeid              int64  `json:"recorddifftypeid"`
	Recorddiffid                  int64  `json:"recorddiffid"`
	Mststatetypeid                int64  `json:"mststatetypeid"`
	Mststateid                    int64  `json:"mststateid"`
	Activeflg                     int64  `json:"activeflg"`
	Offset                        int64  `json:"offset"`
	Limit                         int64  `json:"limit"`
	Clientname                    string `json:"clientname"`
	Mstorgnhirarchyname           string `json:"mstorgnhirarchyname"`
	Recorddifferentiationtypename string `json:"recorddifferentiationtypename"`
	Recorddifferentiationname     string `json:"recorddifferentiationname"`
	Statetypename                 string `json:"statetypename"`
	Statename                     string `json:"statename"`
}
type MaprecordstatetodifferentiationEntities struct {
	Total  int64                                   `json:"total"`
	Values []MaprecordstatetodifferentiationEntity `json:"values"`
}

type MaprecordstatetodifferentiationResponse struct {
	Success bool                                    `json:"success"`
	Message string                                  `json:"message"`
	Details MaprecordstatetodifferentiationEntities `json:"details"`
}

type MaprecordstatetodifferentiationResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MaprecordstatetodifferentiationEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
