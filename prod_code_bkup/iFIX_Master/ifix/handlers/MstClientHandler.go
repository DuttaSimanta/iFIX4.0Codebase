package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// ThrowClientAllResponse function is used to throw success response of All data in JSON format
func ThrowClientAllResponse(successMessage string, responseData entities.MstClientEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstClientEntityResponse{}
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

// ThrowClientIntResponse function is used to throw success response of integer data in JSON format
func ThrowClientIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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

func GetAllClientsnames(w http.ResponseWriter, req *http.Request) {
	data1, success, _, msg := models.GetAllClientsnames()
	ThrowClientAllNamesResponse(msg, data1, w, success)

}

func ThrowClientAllNamesResponse(successMessage string, responseData []entities.AllMstClientEntity, w http.ResponseWriter, success bool) {
	var response = entities.AllMstClientEntityResponse{}
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
//AddClient method is used for insert client data
func AddClient(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.AddClients(&data)
		ThrowClientIntResponse(msg, data, w, success)
	}
}

//DeleteClient method is used for delete client data
func DeleteClient(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteClients(&data)
		ThrowClientUsersIntResponse(msg, 0, w, success)
	}
}

//UpdateClient  method is used for update client data
func UpdateClient(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateClients(&data)
		ThrowClientUsersIntResponse(msg, 0, w, success)
	}
}

//GetAllClients  method is used for get clients data
func GetAllClients(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstClientEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.GetAllClients(&data)
		ThrowClientAllResponse(msg, data1, w, success)
	}
}
