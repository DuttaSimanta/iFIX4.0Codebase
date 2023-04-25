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

type EventParam struct {
	StatusSeq        int64 `json:"statusid"`
	PriorityID       int64 `json:"priorityid"`
	NoOfCount        int64 `json:"noofcount"`
	PriorityIDForSLA int64 `json:"processid"`
	ProcessComplete  int64 `json:"processcomplete"`
	NoOfDays         int64 `json:"noofdays"`
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

//NotPostMethodResponse function is used to return not post method response
func NotPostMethodResponse() APIResponse {
	var response = APIResponse{}
	response.Status = false
	response.Message = "405 method not allowed."
	return response
}
