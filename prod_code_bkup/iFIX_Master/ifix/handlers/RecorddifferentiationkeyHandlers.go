package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecorddifferentiationhigherkeyAllResponse(successMessage string, responseData entities.RecorddifferentiationhigherkeyEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecorddifferentiationhigherkeyEntityResponse{}
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

func ThrowRecorddifferentiationhigherkeyIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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

func InsertRecorddifferentiationhigherkeyEntity(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationhigherkeyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Inserthighkeyvalue(&data)
		ThrowRecorddifferentiationhigherkeyIntResponse(msg, data, w, success)
	}
}

func GetAllRecorddifferentiationhigherkeyEntity(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationhigherkeyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllRecorddifferentiationHighkey(&data)
		ThrowRecorddifferentiationhigherkeyAllResponse(msg, data, w, success)
	}
}

func DeleteRecorddifferentiationhigherkeyEntity(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationhigherkeyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.Deletehigherkey(&data)
		ThrowRecorddifferentiationhigherkeyIntResponse(msg, 0, w, success)
	}
}

func UpdateRecorddifferentiationhigherkeyEntity(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationhigherkeyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.Updatehigherkey(&data)
		ThrowRecorddifferentiationhigherkeyIntResponse(msg, 0, w, success)
	}
}
