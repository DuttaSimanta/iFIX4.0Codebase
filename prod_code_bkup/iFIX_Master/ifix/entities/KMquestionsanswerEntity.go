package entities

import (
	"encoding/json"
	"io"
)

type KMquestionsanswerEntity struct {
	ID                  int64  `json:"id"`
	QuestiondifftypeID  int64  `json:"questiondifftypeid"`
	QuestiondiffID      int64  `json:"questiondiffid"`
	Questions           string `json:"questions"`
	Questionheader      string `json:"questionheader"`
	AnswerdifftypeID    int64  `json:"answerdifftypeid"`
	AnswerdiffID        int64  `json:"answerdiffid"`
	Answer              string `json:"answer"`
	Answerheader        string `json:"answerheader"`
	Activeflg           int64  `json:"activeflg"`
	Clientid            int64  `json:"clientid"`
	Mstorgnhirarchyid   int64  `json:"mstorgnhirarchyid"`
	Clientname          string `json:"clientname"`
	Mstorgnhirarchyname string `json:"mstorgnhirarchyname"`
	Offset              int64  `json:"offset"`
	Limit               int64  `json:"limit"`
}

type KMquestionsanswerEntities struct {
	Total  int64                     `json:"total"`
	Values []KMquestionsanswerEntity `json:"values"`
}

type KMquestionsanswerResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details KMquestionsanswerEntities `json:"details"`
}

type KMquestionsanswerResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *KMquestionsanswerEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
