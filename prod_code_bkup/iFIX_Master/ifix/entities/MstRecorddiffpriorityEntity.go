package entities
import (
	"encoding/json"
	"io"
)


type MstRecorddiffpriorityEntity struct {
     
     Id                       int64  `json:"id"`
     Clientid                 int64  `json:"clientid"`
     Clientname               string `json:"clienname"`
     Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
     Mstorgnhirarchyname      string `json:"mstorgnhirarchyname"`
     Typedifftypeid             int64 `json:"typedifftypeid"`
     Typedifftypename          string `json:"typedifftypename"`
     Typediffid                int64 `json:"typediffid"`
     Typediffname              string `json:"typediffname"`
     Difftypeid                   int64 `json:"difftypeid"`
     Difftypename             string `json:"difftypename"`
     Diffid                   int64 `json:"diffid"`
     Diffname                 string `json:"diffname"`
     Priority                 int64 `json:"priority"`
     Activeflg                int64  `json:"activeglg"`
     Offset                   int64  `json:"offset"`
	 Limit                    int64  `json:"limit"`
}
type MstRecorddiffpriorityEntities struct {
	Total  int64            `json:"total"`
	Values []MstRecorddiffpriorityEntity `json:"values"`
}

type MstRecorddiffpriorityResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details MstRecorddiffpriorityEntities  `json:"details"`
}

type MstRecorddiffpriorityResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstRecorddiffpriorityEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
