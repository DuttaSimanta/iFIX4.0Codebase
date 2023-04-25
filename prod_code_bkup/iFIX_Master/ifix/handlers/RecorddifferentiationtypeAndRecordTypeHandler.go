package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecorddifftypeAndRecordtypeIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.RecorddifftypeAndRecordTypeResponseInt{}
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

func InsertRecorddifftypeAndRecordtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifftypeAndRecordTypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertRecorddifftypeAndRecordtype(&data)
		ThrowRecorddifftypeAndRecordtypeIntResponse(msg, data, w, success)
	}
}
func DeleteRecorddifftypeAndRecordtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifftypeAndRecordTypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteRecorddifftypeAndRecordtype(&data)
		ThrowRecorddifftypeAndRecordtypeIntResponse(msg, 0, w, success)
	}
}

func UpdateRecorddifftypeAndRecordtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifftypeAndRecordTypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateRecorddifftypeAndRecordtype(&data)
		ThrowRecorddifftypeAndRecordtypeIntResponse(msg, 0, w, success)
	}
}
