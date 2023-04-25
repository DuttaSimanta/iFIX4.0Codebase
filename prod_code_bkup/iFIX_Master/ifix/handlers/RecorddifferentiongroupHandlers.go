package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecorddifferentiongroupAllResponse(successMessage string, responseData entities.RecorddifferentiongroupEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecorddifferentiongroupResponse{}
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

func ThrowWoringlevelAllResponse(successMessage string, responseData entities.WorkinglevelEntities, w http.ResponseWriter, success bool) {
	var response = entities.WorkinglevelResponse{}
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

func ThrowRecorddifferentiongroupIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.RecorddifferentiongroupResponseInt{}
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

func InsertRecorddifferentiongroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiongroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertRecorddifferentiongroup(&data)
		ThrowRecorddifferentiongroupIntResponse(msg, data, w, success)
	}
}

func GetAllRecorddifferentiongroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiongroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllRecorddifferentiongroup(&data)
		ThrowRecorddifferentiongroupAllResponse(msg, data, w, success)
	}
}

func DeleteRecorddifferentiongroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiongroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteRecorddifferentiongroup(&data)
		ThrowRecorddifferentiongroupIntResponse(msg, 0, w, success)
	}
}

func UpdateRecorddifferentiongroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiongroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateRecorddifferentiongroup(&data)
		ThrowRecorddifferentiongroupIntResponse(msg, 0, w, success)
	}
}

func GetWorkinglevel(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiongroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetWorkinglevel(&data)
		ThrowWoringlevelAllResponse(msg, data, w, success)
	}
}
