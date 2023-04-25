package entities
import (
	"encoding/json"
	"io"
)


type MstExcelTemplateEntity struct {

     Id                       int64  `json:"id"`
     HeaderName               string `json:"headername"`
     Clientid                 int64  `json:"clientid"`
     Clientname               string `json:"clienname"`
     Mstorgnhirarchyid        int64  `json:"mstorgnhirarchyid"`
     Mstorgnhirarchyname      string `json:"mstorgnhirarchyname"`
     RecordDiffTypeid         int64  `json:"recorddifftypeid`
     RecordDiffTypeName       string `json:"recorddifftypename"`
     RecordDiffid             int64  `json:"recorddiffid"`
     RecordDiffName           string  `json:"recorddiffname"`
     TemplateTypeid           int64   `json:"templatetypeid"`
     TemplateTypeName         string  `json:"templatetypename"`
     SeqNo                    int64   `json:"seqno"`
     /*
     Typedifftypeid             int64 `json:"typedifftypeid"`
     Typedifftypename          string `json:"typedifftypename"`
     Typediffid                int64 `json:"typediffid"`
     Typediffname              string `json:"typediffname"`
     Difftypeid                   int64 `json:"difftypeid"`
     Difftypename             string `json:"difftypename"`
     Diffid                   int64 `json:"diffid"`
     Diffname                 string `json:"diffname"`
     Priority                 int64 `json:"priority"`*/
     Activeflg                int64  `json:"activeglg"`
     Offset                   int64  `json:"offset"`
	 Limit                    int64  `json:"limit"`
}
type MstExcelTemplateEntities struct {
	Total  int64            `json:"total"`
	Values []MstExcelTemplateEntity `json:"values"`
}

type MstExcelTemplateResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details MstExcelTemplateEntities  `json:"details"`
}

type MstExcelTemplateResponseInt struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details int64  `json:"details"`
}

func (w *MstExcelTemplateEntity) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(w)
}
