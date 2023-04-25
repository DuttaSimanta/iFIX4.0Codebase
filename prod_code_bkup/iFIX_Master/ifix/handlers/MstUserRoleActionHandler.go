package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)



// ThrowUserRoleActionAllResponse function is used to throw success response of All data in JSON format
func ThrowUserRoleActionAllResponse(successMessage string, responseData entities.MstUserRoleActionEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstUserRoleActionEntityResponse{}
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

// ThrowUserRoleActionIntResponse function is used to throw success response of integer data in JSON format
func ThrowUserRoleActionIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstUserRoleActionEntityResponseInt{}
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

//AddUserActionRole method is used for insert user data
func AddUserActionRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.AddUserRoleAction(&data)
		ThrowUserRoleActionIntResponse(msg, data, w, success)
	}
}

//DeleteUserRoleAction method is used for delete user data
func DeleteUserRoleAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteUserRoleAction(&data)
		ThrowUserRoleActionIntResponse(msg, 0, w, success)
	}
}

//UpdateUserRoleAction  method is used for update user data
func UpdateUserRoleAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateUserRoleAction(&data)
		ThrowUserRoleActionIntResponse(msg, 0, w, success)
	}
}

//GetAllUserRoleActionForClient  method is used for get client user data
func GetAllUserRoleActionForClient(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllUserRoleActionForClient(&data)
		ThrowUserRoleActionAllResponse(msg, data1, w, success)
	}
}

//GetAllUserRoleAction  method is used for get client user data
func GetAllUserRoleAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllUserRoleAction(&data)
		ThrowUserRoleActionAllResponse(msg, data1, w, success)
	}
}
//GetRoleWiseAction  method is used for get client user data
func GetRoleUserWiseAction(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserRoleActionEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetRoleUserWiseAction(&data)
		ThrowActionAllResponse(msg, data1, w, success)
	}
}
