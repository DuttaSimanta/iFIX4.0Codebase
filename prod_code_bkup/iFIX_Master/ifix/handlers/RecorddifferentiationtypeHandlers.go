package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecorddifferentiationtypeAllResponse(successMessage string, responseData entities.RecorddifferentiationtypeEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecorddifferentiationtypeResponse{}
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

func ThrowRecorddifferentiationtypeIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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

func InsertRecorddifferentiationtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertRecorddifferentiationtype(&data)
		ThrowRecorddifferentiationtypeIntResponse(msg, data, w, success)
	}
}

func GetAllRecorddifferentiationtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllRecorddifferentiationtype(&data)
		ThrowRecorddifferentiationtypeAllResponse(msg, data, w, success)
	}
}
func GetRecorddifferentiationtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetRecorddifferentiationtype(&data)
		ThrowRecorddifferentiationtypeAllResponse(msg, data, w, success)
	}
}

func DeleteRecorddifferentiationtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteRecorddifferentiationtype(&data)
		ThrowRecorddifferentiationtypeIntResponse(msg, 0, w, success)
	}
}

func UpdateRecorddifferentiationtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateRecorddifferentiationtype(&data)
		ThrowRecorddifferentiationtypeIntResponse(msg, 0, w, success)
	}
}
