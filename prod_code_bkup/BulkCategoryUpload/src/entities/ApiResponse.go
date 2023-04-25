package entities

import (
	"encoding/json"
	//"io"
	"log"
	"net/http"
	Logger "src/logger"
	//"Errors"
)

type APIResponse struct {
	Status   bool   `json:"success"`
	Message  string `json:"message"`
	Response string `json:"response"`
}
type APIResponseDownload struct {
	Status           bool   `json:"success"`
	Message          string `json:"message"`
	UploadedFileName string `json:"uploadedfilename"`
	OriginalFileName string `json:"originalfilename"`
}

// ErrorResponse Structure used to handle error  response using json
type ErrorResponse struct {
	Status  bool   `json:"success"`
	Message string `json:"message"`
	//Response []string `json:"response"`
}
type FileForUpload struct {
	ClientID     int64  `json:"clientid"`
	OrgID        int64  `json:"mstorgnhirarchyid"`
	FileToUpload []byte `json:"myFile"`
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

//......................................
type FileuploadEntities struct {
	Values []FileuploadEntity `json:"values"`
}

type FileuploadResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Details FileuploadEntity `json:"details"`
}

func BlankPathCheckResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "404 not found."
	log.Println("Blank request called")
	Logger.Log.Println("ERROR: Blank request called")
	return response
}
func ThrowJSONDownloadResponse(response APIResponseDownload, w http.ResponseWriter) {
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
