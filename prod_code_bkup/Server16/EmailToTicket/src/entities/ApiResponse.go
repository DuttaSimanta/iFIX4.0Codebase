package entities

import (
	"encoding/json"
	"io"

	//"io"
	"log"
	"net/http"
	//"Errors"
)

type UpdateRecordStatus struct {
	ClientID       int64 `json:"clientid"`
	OrgID          int64 `json:"mstorgnhirarchyid"`
	RecordID       int64 `json:"recordid"`
	CurrentStateID int64 `json:"reordstatusid"`
}

type APIResponse struct {
	Status   bool     `json:"success"`
	Message  string   `json:"message"`
	Response Response `json:"response"`
}

type FileForUpload struct {
	ClientID     int64  `json:"clientid"`
	OrgID        int64  `json:"mstorgnhirarchyid"`
	FileToUpload []byte `json:"myFile"`
}

type MoveWorkFlow struct {
	ClientID             int64 `json:"clientid"`
	OrgID                int64 `json:"mstorgnhirarchyid"`
	Recorddifftypeid     int64 `json:"recorddifftypeid"`
	RecordDiffID         int64 `json:"recorddiffid"`
	Previousstateid      int64 `json:"previousstateid"`
	Currentstateid       int64 `json:"currentstateid"`
	Manualstateselection int64 `json:"manualstateselection"`
	Transactionid        int64 `json:"transactionid"`
	Createdgroupid       int64 `json:"createdgroupid"`
	Mstgroupid           int64 `json:"mstgroupid"`
	Mstuserid            int64 `json:"mstuserid"`
	UserID               int64 `json:"userid"`
}

type StateSeq struct {
	ClientID  int64 `json:"clientid"`
	OrgID     int64 `json:"mstorgnhirarchyid"`
	TypeSeqNo int64 `json:"typeseqno"`
	SeqNo     int64 `json:"seqno"`
	UserID    int64 `json:"userid"`
}
type StateSeqRespEntity struct {
	Recorddifftypeid int64 `json:"recorddifftypeid"`
	RecordDiffID     int64 `json:"recorddiffid"`
	Mststateid       int64 `json:"'mststateid"`
}

// type StateSeqRespEntities struct {
// 	Values []StateSeqRespEntity `json:"'Values"`
// }
type StateSeqResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Details []StateSeqRespEntity `json:"details"`
}

func (p *StateSeqResponse) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

type Response struct {
	Ticketid string `json:"ticketid"`
	//Status string `json:"status"`
}
type RecordcommonEntity struct {
	ClientID          int64   `json:"clientid"`
	Mstorgnhirarchyid int64   `json:"mstorgnhirarchyid"`
	RecordID          int64   `json:"recordid"`
	RecordstageID     int64   `json:"recordstageid"`
	TermID            int64   `json:"termid"`
	Termvalue         string  `json:"termvalue"`
	ForuserID         int64   `json:"foruserid"`
	Recorddifftypeid  int64   `json:"recorddifftypeid"`
	Recorddiffid      int64   `json:"recorddiffid"`
	Termdescription   string  `json:"termdescription"`
	UserID            int64   `json:"userid"`
	Usergroupid       int64   `json:"usergroupid"`
	Termseq           int64   `json:"termseq"`
	ID                int64   `json:"id"`
	Sequance          []int64 `json:"sequance"`
}

type AdditionalField struct {
	Mstdifferentiationtypeid int64 `json:"mstdifferentiationtypeid"`
	Mstdifferentiationid     int64 `json:"mstdifferentiationid"`
}

type AdditionalFieldsList struct {
	ClientID              int64             `json:"clientid"`
	Mstorgnhirarchyid     int64             `json:"mstorgnhirarchyid"`
	Mstdifferentiationset []AdditionalField `json:"mstdifferentiationset"`
	UserID                int64             `json:"userid"`
}
type AdditionalFieldResponse struct {
	Fieldid       int64  `json:"fieldid"`
	Termsid       int64  `json:"message"`
	TermsName     string `json:"termsname"`
	TermsValue    string `json:"termsvalue"`
	Termstypeid   int64  `json:"termstypeid"`
	Termstypename string `json:"termstypename"`
	Ismandatory   int64  `json:"ismandatory"`
}

type AdditionalFieldListResponse struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Details []AdditionalFieldResponse `json:"response"`
}
type FileuploadEntity struct {
	Id                 int64  `json:"id"`
	Clientid           int64  `json:"clientid"`
	Mstorgnhirarchyid  int64  `json:"mstorgnhirarchyid"`
	Credentialtype     string `json:"credentialtype"`
	Credentialaccount  string `json:"credentialaccount"`
	Credentialpassword string `json:"credentialpassword"`
	Credentialkey      string `json:"credentialkey"`
	Activeflg          int64  `json:"activeflg"`
	Originalfile       string `json:"originalfile"`
	Filename           string `json:"filename"`
	Path               string `json:"path"`
}

type FileuploadEntities struct {
	Values []FileuploadEntity `json:"values"`
}

type FileuploadResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details FileuploadEntity `json:"details"`
}
type StatusResponse struct {
	Status       bool   `json:"success"`
	Message      string `json:"message"`
	TicketStatus string `json:"ticketstatus"`
}

// ErrorResponse Structure used to handle error  response using json
type ErrorResponse struct {
	Status  bool   `json:"success"`
	Message string `json:"message"`
	//Response []string `json:"response"`
}

func BlankPathCheckResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "404 not found."
	log.Println("Blank request called")
	return response
}
func ThrowJSONStatusResponse(response StatusResponse, w http.ResponseWriter) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowJSONResponse(response APIResponse, w http.ResponseWriter) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowJSONErrorResponse(response ErrorResponse, w http.ResponseWriter) {
	//log.Println(response)
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// NotPostMethodResponse function is used to return not post method response
func NotPostMethodResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "405 method not allowed."
	return response
}
