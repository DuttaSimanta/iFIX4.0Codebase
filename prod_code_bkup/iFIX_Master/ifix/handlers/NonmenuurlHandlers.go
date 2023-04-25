package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowNonmenuurlResponse(successMessage string, responseData []entities.NonmenuurlEntity, w http.ResponseWriter, success bool) {
	var response = entities.NonmenuurlsingleResponse{}
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
func ThrowNonmenuurlAllResponse(successMessage string, responseData entities.NonmenuurlEntities, w http.ResponseWriter, success bool) {
	var response = entities.NonmenuurlResponse{}
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

func ThrowNonmenuurlIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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

func ThrowMstUrlKeyAllResponse(successMessage string, responseData entities.MsturlkeyEntities, w http.ResponseWriter, success bool) {
	var response = entities.MsturlkeyResponse{}
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

func InsertNonmenuurl(w http.ResponseWriter, req *http.Request) {
	var data = entities.NonmenuurlEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertNonmenuurl(&data)
		ThrowNonmenuurlIntResponse(msg, data, w, success)
	}
}

func GetAllNonmenuurl(w http.ResponseWriter, req *http.Request) {
	var data = entities.NonmenuurlEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllNonmenuurl(&data)
		ThrowNonmenuurlAllResponse(msg, data, w, success)
	}
}
func Geturlbykey(w http.ResponseWriter, req *http.Request) {
	var data = entities.NonmenuurlEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Geturlbykey(&data)
		ThrowNonmenuurlResponse(msg, data, w, success)
	}
}

func DeleteNonmenuurl(w http.ResponseWriter, req *http.Request) {
	var data = entities.NonmenuurlEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteNonmenuurl(&data)
		ThrowNonmenuurlIntResponse(msg, 0, w, success)
	}
}

func UpdateNonmenuurl(w http.ResponseWriter, req *http.Request) {
	var data = entities.NonmenuurlEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateNonmenuurl(&data)
		ThrowNonmenuurlIntResponse(msg, 0, w, success)
	}
}

func GetAllUrlkey(w http.ResponseWriter, req *http.Request) {
	var data = entities.MsturlkeyInputEntity{}
	jsonError := data.FromInputJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllUrlkey(&data)
		ThrowMstUrlKeyAllResponse(msg, data, w, success)
	}
}
