package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowWorkflowtemplateAllResponse(successMessage string, responseData []entities.WorkflowResponseEntity, w http.ResponseWriter, success bool) {
	var response = entities.WorkflowEntityResponse{}
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
func ThrowWorkflowTemplateStateResponse(successMessage string, responseData entities.WorkflowStateResponseEntity, w http.ResponseWriter, success bool) {
	var response = entities.WorkflowStateEntityResponse{}
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
func Getprocesstemplatedetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getprocesstemplatedetails(&data)
		ThrowWorkflowtemplateAllResponse(msg, data, w, success)
	}
}
func Createprocesstemplatetransition(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Createprocesstemplatetransition(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Insertprocesstemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Insertprocesstemplate(&data)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}

func Deletetemplatetransitionstate(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Deletetemplatetransitionstate(&data)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Upserttemplatetransitiondetails(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Upserttemplatetransitiondetails(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Gettemplatetransitionstatedetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Gettemplatetransitionstatedetails(&data)
		ThrowWorkflowTemplateStateResponse(msg, data, w, success)
	}
}
func Getprocesstemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getprocesstemplate(&data)
		ThrowWorkflowtemplateAllResponse(msg, data, w, success)
	}
}
