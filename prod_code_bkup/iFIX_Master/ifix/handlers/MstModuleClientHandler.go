package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)


// ThrowModuleByClientResponse function is used to throw success response of All data in JSON format
func ThrowModuleByClientResponse(successMessage string, responseData []entities.MstModuleByClientEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstModuleByClientEntityResponse{}
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

// ThrowModuleClientAllResponse function is used to throw success response of All data in JSON format
func ThrowModuleClientAllResponse(successMessage string, responseData entities.MstModuleClientEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstModuleClientEntityResponse{}
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

// ThrowModuleClientIntResponse function is used to throw success response of integer data in JSON format
func ThrowModuleClientIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstClientUserRoleResponseInt{}
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

//AddModuleClient method is used for insert client module data
func AddModuleClient(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstModuleClientEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.AddModuleClients(&data)
		ThrowModuleClientIntResponse(msg, data, w, success)
	}
}

//DeleteModuleClient method is used for delete client module data
func DeleteModuleClient(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstModuleClientEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteModuleClients(&data)
		ThrowModuleClientIntResponse(msg, 0, w, success)
	}
}

//UpdateModuleClient  method is used for update client module data
func UpdateModuleClient(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstModuleClientEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateModuleClients(&data)
		ThrowModuleClientIntResponse(msg, 0, w, success)
	}
}

//GetAllModuleClients  method is used for get clients data
func GetAllModuleClients(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstModuleClientEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllModuleClients(&data)
		ThrowModuleClientAllResponse(msg, data1, w, success)
	}
}
//GetModuleByOrgId  method is used for get clients data
func GetModuleByOrgId(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstModuleClientEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetModuleByOrgId(&data)
		ThrowModuleByClientResponse(msg, data1, w, success)
	}
}
