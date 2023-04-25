package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// ThrowMstClientUserAllResponse function is used to throw success response of All data in JSON format
func ThrowMstClientUserAllResponse(successMessage string, responseData entities.MstClientUserRoleEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstClientUserRoleResponse{}
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

// ThrowMstClientRoleResponse function is used to throw success response of All data in JSON format
func ThrowMstClientRoleResponse(successMessage string, responseData []entities.MstClientRoleEntityResp, w http.ResponseWriter, success bool) {
	var response = entities.MstClientRoleResponse{}
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

// ThrowMstClientUserIntResponse function is used to throw success response of integer data in JSON format
func ThrowMstClientUserIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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

//AddRole method is used for insert user role data
func AddRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserRoleEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.AddRoleModel(&data)
		ThrowMstClientUserIntResponse(msg, data, w, success)
	}
}

//DeleteRole method is used for delete user role data
func DeleteRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserRoleEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteRoleModel(&data)
		ThrowMstClientUserIntResponse(msg, 0, w, success)
	}
}

//UpdateRole  method is used for update user role data
func UpdateRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserRoleEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateRoleModel(&data)
		ThrowMstClientUserIntResponse(msg, 0, w, success)
	}
}

//GetAllRole  method is used for get client user data
func GetAllRole(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserRoleEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllRoleModel(&data)
		ThrowMstClientUserAllResponse(msg, data1, w, success)
	}
}
//Getrolebyorgid  method is used for get client wise role data
func Getrolebyorgid(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientUserRoleEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.Getrolebyorgid(&data)
		ThrowMstClientRoleResponse(msg, data1, w, success)
	}
}
