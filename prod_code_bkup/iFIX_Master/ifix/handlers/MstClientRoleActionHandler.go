package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// ThrowClientRoleActionAllResponse function is used to throw success response of All data in JSON format
func ThrowClientRoleActionAllResponse(successMessage string, responseData entities.MstClientRoleActionEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstClientRoleActionEntityResponse{}
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

// ThrowClientRoleActionIntResponse function is used to throw success response of integer data in JSON format
func ThrowClientRoleActionIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstClientRoleActionEntityResponseInt{}
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

// ThrowActionAllResponse function is used to throw success response of All data in JSON format
func ThrowActionAllResponse(successMessage string, responseData []entities.MstActionEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstActionEntityResponse{}
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

//AddActionRole method is used for insert user data
func AddActionRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.AddRoleAction(&data)
		ThrowClientRoleActionIntResponse(msg, data, w, success)
	}
}

//DeleteRoleAction method is used for delete user data
func DeleteRoleAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteRoleAction(&data)
		ThrowClientRoleActionIntResponse(msg, 0, w, success)
	}
}

//UpdateRoleAction  method is used for update user data
func UpdateRoleAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateRoleAction(&data)
		ThrowClientRoleActionIntResponse(msg, 0, w, success)
	}
}

//GetAllRoleActionForClient  method is used for get client user data
func GetAllRoleActionForClient(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllRoleActionForClient(&data)
		ThrowClientRoleActionAllResponse(msg, data1, w, success)
	}
}

//GetAllRoleAction  method is used for get client user data
func GetAllRoleAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllRoleAction(&data)
		ThrowClientRoleActionAllResponse(msg, data1, w, success)
	}
}

//GetAllAction  method is used for get client user data
func GetAllAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllAction()
		ThrowActionAllResponse(msg, data1, w, success)
	}
}
//GetRoleWiseAction  method is used for get client user data
func GetRoleWiseAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetRoleWiseAction(&data)
		ThrowActionAllResponse(msg, data1, w, success)
	}
}
