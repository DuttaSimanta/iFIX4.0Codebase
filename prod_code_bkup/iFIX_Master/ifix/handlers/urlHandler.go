package handlers

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
	"encoding/json"
)


// ThrowUrlResponse function is used to throw success response of All data in JSON format
func ThrowUrlResponse(successMessage string, responseData []entities.UrlRespEntity, w http.ResponseWriter, success bool) {
	var response = entities.UrlResponseOnly{}
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
// ThrowUrlAllResponse function is used to throw success response of All data in JSON format
func ThrowUrlAllResponse(successMessage string, responseData entities.UrlEntities, w http.ResponseWriter, success bool) {
	var response = entities.UrlResponse{}
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
// ThrowModuleUrlAllResponse function is used to throw success response of All data in JSON format
func ThrowModuleUrlAllResponse(successMessage string, responseData entities.ModuleUrlEntities, w http.ResponseWriter, success bool) {
	var response = entities.ModuleUrlResponse{}
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
// ThrowUrlIntResponse function is used to throw success response of All data in JSON format
func ThrowUrlIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.UrlResponseInt{}
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

func InsertUrl(w http.ResponseWriter, req *http.Request) {
	var data = entities.UrlEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		log.Print("--------insert handler-----------")
		//if(len(responseError)==0){
		data, success, _, msg := models.InsertUrl(&data)
		ThrowUrlIntResponse(msg, data, w, success)
	}
}

func GetAllUrls(w http.ResponseWriter, req *http.Request) {
	var data = entities.PaginationEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.GetAllUrls(&data)
		ThrowUrlAllResponse(msg, data, w, success)
	}
}

func GetAllModuleUrls(w http.ResponseWriter, req *http.Request) {
	var data = entities.PaginationEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.GetAllModuleUrls(&data)
		ThrowModuleUrlAllResponse(msg, data, w, success)
	}
}
func DeleteUrl(w http.ResponseWriter, req *http.Request) {
	var data = entities.UrlEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteUrl(&data)
		ThrowUrlIntResponse(msg, 0, w, success)
	}
}
func DeleteModuleUrl(w http.ResponseWriter, req *http.Request) {
	var data = entities.ModuleUrlEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.DeleteModUrl(&data)
		ThrowUrlIntResponse(msg, 0, w, success)
	}
}
func UpdateUrl(w http.ResponseWriter, req *http.Request) {
	var data = entities.UrlEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.UpdateUrl(&data)
		ThrowUrlIntResponse(msg, 0, w, success)
	}
}
func GetDistinctUrl(w http.ResponseWriter, req *http.Request) {
	var data = entities.UrlEntity{}
	jsonError := data.FromJSON(req.Body)
	log.Println("inside controller")
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.GetDistinctUrl(&data)
		ThrowUrlResponse(msg, data, w, success)
	}
}
func GetRemainingUrl(w http.ResponseWriter, req *http.Request) {
	var data = entities.UrlEntity{}
	jsonError := data.FromJSON(req.Body)
	log.Println("inside controller")
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.GetRemainingUrl(&data)
		ThrowUrlResponse(msg, data, w, success)
	}
}
