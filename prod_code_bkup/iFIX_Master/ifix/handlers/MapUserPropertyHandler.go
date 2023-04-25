package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowGetUserPropertyNameResponse(successMessage string, responseData []entities.GetUserPropertyNameEntity, w http.ResponseWriter, success bool) {
	var response = entities.GetUserPropertyNameResponse{}
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

func ThrowUserRolePropertyAllResponse(successMessage string, responseData entities.MapUserRolePropertyEntities, w http.ResponseWriter, success bool) {
	var response = entities.MapUserRolePropertyResponse{}
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

func ThrowUserRolePropertyIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MapUserRolePropertyResponseInt{}
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

func InsertUserRoleProperty(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapUserRolePropertyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertUserRoleProperty(&data)
		ThrowUserRolePropertyIntResponse(msg, data, w, success)
	}
}

func GetAllUserRoleProperty(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapUserRolePropertyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllUserRoleProperty(&data)
		ThrowUserRolePropertyAllResponse(msg, data, w, success)
	}
}

func GetUserPropertyName(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapUserRolePropertyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetUserPropertyName(&data)
		ThrowGetUserPropertyNameResponse(msg, data, w, success)
	}
}

func UpdateUserPropertyName(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapUserRolePropertyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateUserPropertyName(&data)
		ThrowUserRolePropertyIntResponse(msg, 0, w, success)
	}
}

func DeleteUserPropertyName(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapUserRolePropertyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteUserPropertyName(&data)
		ThrowUserRolePropertyIntResponse(msg, 0, w, success)
	}
}
