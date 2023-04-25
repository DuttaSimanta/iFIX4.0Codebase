package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMapprocessstateAllResponse(successMessage string, responseData entities.MapprocessstateEntities, w http.ResponseWriter, success bool) {
	var response = entities.MapprocessstateResponse{}
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

func ThrowMapprocessstateIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MapprocessstateResponseInt{}
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

func InsertMapprocessstate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapprocessstateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMapprocessstate(&data)
		ThrowMapprocessstateIntResponse(msg, data, w, success)
	}
}

func GetAllMapprocessstate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapprocessstateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMapprocessstate(&data)
		ThrowMapprocessstateAllResponse(msg, data, w, success)
	}
}

func DeleteMapprocessstate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapprocessstateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMapprocessstate(&data)
		ThrowMapprocessstateIntResponse(msg, 0, w, success)
	}
}

func UpdateMapprocessstate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapprocessstateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMapprocessstate(&data)
		ThrowMapprocessstateIntResponse(msg, 0, w, success)
	}
}
