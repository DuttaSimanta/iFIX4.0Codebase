package entities

import (
	"encoding/json"
	//"io"
	"log"
	"net/http"
	//"Errors"
)

type APIResponse struct {
	Status   bool   `json:"success"`
	Message  string `json:"message"`
	Response string `json:"response"`
}

// ErrorResponse Structure used to handle error  response using json
type ErrorResponse struct {
	Status  bool   `json:"success"`
	Message string `json:"message"`
	//Response []string `json:"response"`
}
type Attachment struct {
	OriginalFilenames string `json:"originalfilenames"`
	UploadedFileName  string `json:"uploadedfilename"`
}

type InstantMessagingEntity struct {
	ClientID     int64        `json:"clientid"`
	OrgID        int64        `json:"mstorgnhirarchyid"`
	RecordID     int64        `json:"recordid"`
	EmailTo      string       `json:"emailto"`
	EmailCc      string       `json:"emailcc"`
	EmailSub     string       `json:"emailsub"`
	EmailBody    string       `json:"emailbody"`
	CreatedID    int64        `json:"userid"`
	CreatedGrpID int64        `json:"createdgrpid"`
	Attachments  []Attachment `json:"attachments"`
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

func BlankPathCheckResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "404 not found."
	log.Println("Blank request called")
	return response
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
