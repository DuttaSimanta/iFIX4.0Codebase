//SearchUser  method is used for search  user data
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"src/entities"
	"src/logger"
	"src/models"
)

func ThrowLocationtAllResponse(successMessage string, responseData entities.LocationEntities, w http.ResponseWriter, success bool) {
	var response = entities.LocationResponse{}
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
func ThrowSearchLocationIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.LocationResponseInt{}
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

func ThrowSearchLocationResponse(successMessage string, responseData []entities.LocationSearchEntity, w http.ResponseWriter, success bool) {
	var response = entities.LocationSearchEntityResponse{}
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

func ThrowSelectLocationResponse(successMessage string, responseData []entities.LocationSelectEntity, w http.ResponseWriter, success bool) {
	var response = entities.LocationSelectEntityResponse{}
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
func SearchLocation(w http.ResponseWriter, req *http.Request) {
	var data = entities.LocationPriorityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.SearchLocation(&data)
		ThrowSearchLocationResponse(msg, data1, w, success)
	}
}
func SelectLocation(w http.ResponseWriter, req *http.Request) {
	var data = entities.LocationPriorityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.SelectLocation(&data)
		ThrowSelectLocationResponse(msg, data1, w, success)
	}
}
func AddLocation(w http.ResponseWriter, req *http.Request) {
	var data = entities.LocationPriorityEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data1, success, _, msg := models.AddLocation(&data)
		ThrowSearchLocationIntResponse(msg, data1, w, success)
	}
}
func GetAllLocation(w http.ResponseWriter, req *http.Request) {
	var data = entities.LocationPriorityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllLocation(&data)
		ThrowLocationtAllResponse(msg, data, w, success)
	}
}
func DeleteLocation(w http.ResponseWriter, req *http.Request) {
	var data = entities.LocationPriorityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteLocation(&data)
		ThrowSearchLocationIntResponse(msg, 0, w, success)
	}
}

func UpdateLocation(w http.ResponseWriter, req *http.Request) {
	var data = entities.LocationPriorityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateLocation(&data)
		ThrowSearchLocationIntResponse(msg, 0, w, success)
	}
}
