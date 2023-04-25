package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

// ThrowWorkflowUtilityResponse function is used to throw success response of All data in JSON format
func ThrowWorkflowUtilityResponse(successMessage string, responseData []entities.WorkflowSingleEntity, w http.ResponseWriter, success bool) {
	var response = entities.WorkflowUtilityResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func ThrowStateStatusResponse(successMessage string, responseData []entities.StateStatusEntity, w http.ResponseWriter, success bool) {
	var response = entities.StateStatusResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
// ThrowStateProcessResponse function is used to throw success response of All data in JSON format
func ThrowStateProcessResponse(successMessage string, responseData entities.StateCategory, w http.ResponseWriter, success bool) {
	var response = entities.StateProcessResponse{}
	response.Success = success
	response.Message = successMessage
	response.Details = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
func Insertuserticket(w http.ResponseWriter, req *http.Request) {
	var data = entities.TicketUserEntity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Insertuserticket(&data)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Deleteuserticket(w http.ResponseWriter, req *http.Request) {
	var data = entities.TicketUserEntity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Deleteuserticket(&data)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Getworklowutilitylist(w http.ResponseWriter, req *http.Request) {

	var data = entities.WorkflowUtilityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getworklowutilitylist(&data)
		ThrowWorkflowUtilityResponse(msg, data, w, success)
	}
}
func Getutilitydatabyfield(w http.ResponseWriter, req *http.Request) {

	var data = entities.WorkflowUtilityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getutilitydatabyfield(&data)
		ThrowWorkflowUtilityResponse(msg, data, w, success)
	}
}
func Getprocessbydiffid(w http.ResponseWriter, req *http.Request) {

	var data = entities.WorkflowUtilityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getprocessbydiffid(&data)
		ThrowWorkflowUtilityResponse(msg, data, w, success)
	}
}
func Getstatebyseq(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkflowUtilityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data",jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}else {
		data, success, _, msg := models.Getstatebyseq(&data)
		ThrowStateStatusResponse(msg, data, w, success)
	}
}

//SearchUserByOrgnId  method is used for search  user data
func Searchworkflowuser(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkflowUtilityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Searchworkflowuser(&data)
		ThrowSearchUsersResponse(msg, data1, w, success)
	}
}
func Getstatebyprocess(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkflowUtilityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Getstatebyprocess(&data)
		ThrowStateProcessResponse(msg, data1, w, success)
	}
}
func Getstatebyprocesstemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkflowUtilityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Getstatebyprocesstemplate(&data)
		ThrowStateProcessResponse(msg, data1, w, success)
	}
}
func Deleteprocessdetails(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.WorkflowUtilityEntity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Deleteprocessdetails(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Deleteprocesstemplatedetails(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.WorkflowUtilityEntity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Deleteprocesstemplatedetails(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
