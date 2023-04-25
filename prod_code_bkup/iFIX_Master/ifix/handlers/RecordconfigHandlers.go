package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecordconfigAllResponse(successMessage string, responseData entities.RecordconfigEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecordconfigResponse{}
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

func ThrowRecordconfigIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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

func InsertRecordconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertRecordconfig(&data)
		ThrowRecordconfigIntResponse(msg, data, w, success)
	}
}

func GetAllRecordconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllRecordconfig(&data)
		ThrowRecordconfigAllResponse(msg, data, w, success)
	}
}

func DeleteRecordconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteRecordconfig(&data)
		ThrowRecordconfigIntResponse(msg, 0, w, success)
	}
}

func UpdateRecordconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateRecordconfig(&data)
		ThrowRecordconfigIntResponse(msg, 0, w, success)
	}
}
