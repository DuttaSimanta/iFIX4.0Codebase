package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowProcessMonitorResponse(successMessage string, responseData []entities.ProcessMonitorEntity, w http.ResponseWriter, success bool) {
	var response = entities.ProcessMonitorResponse{}
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
func ThrowProcessMonitorServerResponse(successMessage string, responseData []entities.ProcessMonitorServerEntity, w http.ResponseWriter, success bool) {
	var response = entities.ProcessMonitorServerResponse{}
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

func Getallserver(w http.ResponseWriter, req *http.Request) {
	var data = entities.ProcessMonitorEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getallserver()
		ThrowProcessMonitorServerResponse(msg, data, w, success)
	}
}
func Getprocessbyserver(w http.ResponseWriter, req *http.Request) {
	var data = entities.ProcessMonitorEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getprocessbyserver(&data)
		ThrowProcessMonitorResponse(msg, data, w, success)
	}
}
func Getstatusbyprocess(w http.ResponseWriter, req *http.Request) {
	var data = entities.ProcessMonitorEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getstatusbyprocess(&data)
		ThrowProcessMonitorResponse(msg, data, w, success)
	}
}