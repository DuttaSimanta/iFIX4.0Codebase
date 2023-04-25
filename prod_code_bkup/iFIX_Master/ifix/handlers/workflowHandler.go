//***************************//
// Package handlers
// Date Of Creation: 12/01/2021
// Authour Name: Subham Chatterjee
// History: N/A
// Synopsis: This file is used to handle all the request and response related to workflow. It is used as Controller
// Functions: AddMstClient, GetMstClientById
// ThrowMstClientResponse() Parameter:  (<string>,<Structure entities.MstClient>,<http.ResponseWriter>)
// AddMstClient() Parameter:  (<http.ResponseWriter>, <*http.Request>)
// GetMstClientById() Parameter:  {<http.ResponseWriter>, <*http.Request>}
// Global Variable: N/A
// Version: 1.0.0
//***************************//

package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

// ThrowWorkflowStateResponse function is used to throw success response of All data in JSON format
func ThrowWorkflowTransitionResponse(successMessage string, responseData []entities.WorkflowTransitionEntity, w http.ResponseWriter, success bool) {
	var response = entities.WorkflowTransitionEntityResponse{}
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
// ThrowWorkflowStateResponse function is used to throw success response of All data in JSON format
func ThrowWorkflowStateResponse(successMessage string, responseData entities.WorkflowStateResponseEntity, w http.ResponseWriter, success bool) {
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
// ThrowWorkflowAllResponse function is used to throw success response of All data in JSON format
func ThrowWorkflowAllResponse(successMessage string, responseData []entities.WorkflowResponseEntity, w http.ResponseWriter, success bool) {
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

// ThrowWorkflowAllResponse function is used to throw success response of integer data in JSON format
func ThrowWorkflowIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.WorkflowResponseInt{}
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

// ThrowTransactionResponse function is used to throw success  response of All data in JSON format
func ThrowTransactionResponse(successMessage string, responseData []entities.TransactionRespEntity, w http.ResponseWriter, success bool) {
	var response = entities.TransactionEntityResponse{}
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

// InsertProcessDelegateUser function is used to Add Delegate user in a Process Transition
func InsertProcessDelegateUser(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.InsertProcessDelegateUser(&data)
		ThrowWorkflowIntResponse(msg, data, w, success)
		//if modelResponseError!=nil{
		//	entities.ThrowJSONResponse(entities.DbErrorResponse(),w)
		//} else {
		//	ThrowMstClientResponse(msg,mstClientResponseData,w)
		//}
		//} else {
		//	entities.ThrowErrorResponse(responseError,w)
		//}
	}
}

// MoveWorkflow function is used to Move record from one state to another
func MoveWorkflow(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.MoveWorkflowwithapi(&data)
		logger.Log.Print("\n\n MoveWorkflow : ")
		logger.Log.Print(success,msg)
		log.Print("\n\n MoveWorkflow : ")
		log.Print(success,msg)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Checkworkflowstate(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Checkworkflowstate(&data)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
// Checkworkflow function is used to Move record from one state to another
func Checkworkflow(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Checkworkflow(&data)
		logger.Log.Println(success,msg)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Insertprocess(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Insertprocess(&data)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Getprocessdetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getprocessdetails(&data)
		ThrowWorkflowAllResponse(msg, data, w, success)
	}
}

func Gettransitionstatedetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Gettransitionstatedetails(&data)
		ThrowWorkflowStateResponse(msg, data, w, success)
	}
}
func Gettransitiongroupdetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Gettransitiongroupdetails(&data)
		ThrowWorkflowAllResponse(msg, data, w, success)
	}
}
func Checkprocessdelete(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Checkprocessdelete(&data)
		ThrowWorkflowAllResponse(msg, data, w, success)
	}
}
func Createtransition(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	 if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Createtransition(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Upserttransitiondetails(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	 if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Upserttransitiondetails(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Deletetransitionstate(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	 if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Deletetransitionstate(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Changerecordgroup(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Changerecordgroupwithapi(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Getstatedetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//var token = req.Header.Get("Authorization")
		//success:=utility.CheckToken(token,data.Userid)
		//if !success{
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		data, success, _, msg := models.Getstatedetails(&data)
		ThrowTransactionResponse(msg, data, w, success)
	}
}
func Getnextstatedetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getnextstatedetails(&data)
		ThrowTransactionResponse(msg, data, w, success)
	}
}
func Gettransitionbyprocess(w http.ResponseWriter, req *http.Request) {
	var data = entities.Workflowentity{}
	//var responseError []string
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Gettransitionbyprocess(&data)
		ThrowWorkflowTransitionResponse(msg, data, w, success)
	}
}
func Updatechildstatus(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Updatechildstatuswithapi(&pdata)
		logger.Log.Print("Updatechildstatus resp ",success,msg)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Detachchildticket(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Detachchildticket(&pdata)
		logger.Log.Print("Detachchildticket resp ",success,msg)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Updatetaskstatus(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Updatetaskstatus(&pdata)
		logger.Log.Print("Updatetaskstatus resp ",success,msg)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Gethopcount(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Gethopcount(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
func Bulkapprovalfortickets(w http.ResponseWriter, req *http.Request) {
	var pdata = entities.Workflowentity{}
	pjsonError := pdata.FromJSON(req.Body)
	if pjsonError != nil {
		log.Print(pjsonError)
		logger.Log.Println(pjsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Bulkapprovalfortickets(&pdata)
		ThrowWorkflowIntResponse(msg, data, w, success)
	}
}
