package handlers

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
	"encoding/json"
)


// ThrowLoginAllResponse function is used to throw success response of All data in JSON format
func ThrowZoneResponse(successMessage string, responseData []entities.ZoneEntity, w http.ResponseWriter, success bool) {
	var response = entities.ZoneEntityResp{}
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
func Searchzone(w http.ResponseWriter, req *http.Request) {
	var data = entities.ZoneEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		data, success, _, msg := models.Searchzone(&data)
		ThrowZoneResponse(msg, data, w, success)
	}
}
func Heartbeat(w http.ResponseWriter, req *http.Request) {
	var response = entities.GlobalEntityResp{}
	response.Success = true
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}