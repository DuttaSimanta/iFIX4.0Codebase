package handlers

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"

	//"log"
	"encoding/json"
	"net/http"
)

// ThrowFuncMasterResponse function is used to throw success response of All data in JSON format
func ThrowFuncMasterResponse(successMessage string, responseData []entities.FuncmasterEntity, w http.ResponseWriter, success bool) {
	var response = entities.FuncmasterRespEntity{}
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

// ThrowFuncIntResponse function is used to throw success response of integer data in JSON format
func ThrowFuncIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.ApiResponseInt{}
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

// ThrowFuncSingleResponse function is used to throw success response of All data in JSON format
func ThrowFuncSingleResponse(successMessage string, responseData []entities.FuncmappingsingleRespEntity, w http.ResponseWriter, success bool) {
	var response = entities.FuncmappingsingleResponese{}
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

// ThrowFuncResponse function is used to throw success response of All data in JSON format
func ThrowFuncResponse(successMessage string, responseData entities.FuncmappingEntitities, w http.ResponseWriter, success bool) {
	var response = entities.FuncmasterResponse{}
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
func Getfunctionality(w http.ResponseWriter, req *http.Request) {
	data, success, _, msg := models.Getfunctionality()
	ThrowFuncMasterResponse(msg, data, w, success)
}
func Insertfuncmapping(w http.ResponseWriter, req *http.Request) {
	var data = entities.FuncmappingEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.Insertfuncmapping(&data)
		ThrowFuncIntResponse(msg, data, w, success)
	}
}

func Getfuncmappingbytype(w http.ResponseWriter, req *http.Request) {

	var data = entities.FuncmappingEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getfuncmappingbytype(&data)
		ThrowFuncSingleResponse(msg, data, w, success)
	}
}
func Getfuncmappingbycatalogtype(w http.ResponseWriter, req *http.Request) {

	var data = entities.FuncmappingEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getfuncmappingbycatalogtype(&data)
		ThrowFuncSingleResponse(msg, data, w, success)
	}
}
func Getfuncmappingbytypeforquery(w http.ResponseWriter, req *http.Request) {

	var data = entities.FuncmappingEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getfuncmappingbytypeforquery(&data)
		ThrowFuncSingleResponse(msg, data, w, success)
	}
}
func Getfuncmappingdetails(w http.ResponseWriter, req *http.Request) {
	var data = entities.FuncmappingEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print("data", jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Getfuncmappingdetails(&data)
		ThrowFuncResponse(msg, data, w, success)
	}
}
func Deletefunctionmapping(w http.ResponseWriter, req *http.Request) {
	var data = entities.FuncmappingEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		success, _, msg := models.Deletefunctionmapping(&data)
		ThrowFuncIntResponse(msg, 0, w, success)
	}
}
