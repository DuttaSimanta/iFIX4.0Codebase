package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// ThrowClientUserRoleAllResponse function is used to throw success response of All data in JSON format
func ThrowClientUserRoleAllResponse(successMessage string, responseData entities.MapClientUserRoleUserEntities, w http.ResponseWriter, success bool) {
	var response = entities.MapClientUserRoleUserEntityResponse{}
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

// ThrowClientUserRoleIntResponse function is used to throw success response of integer data in JSON format
func ThrowClientUserRoleIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MapClientUserRoleUserEntityResponseInt{}
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

//AddClientUserRole method is used for insert user data
func AddClientUserRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapClientUserRoleUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.AddClientUserRole(&data)
		ThrowClientUserRoleIntResponse(msg, data, w, success)
	}
}

//DeleteClientUserRole method is used for delete user data
func DeleteClientUserRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapClientUserRoleUserEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteClientUserRole(&data)
		ThrowClientUserRoleIntResponse(msg, 0, w, success)
	}
}

//UpdateClientUserRole  method is used for update user data
func UpdateClientUserRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapClientUserRoleUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateClientUserRole(&data)
		ThrowClientUserRoleIntResponse(msg, 0, w, success)
	}
}

//GetAllClientUserRole  method is used for get client user data
func GetAllClientUserRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapClientUserRoleUserEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllClientUserRole(&data)
		ThrowClientUserRoleAllResponse(msg, data1, w, success)
	}
}
