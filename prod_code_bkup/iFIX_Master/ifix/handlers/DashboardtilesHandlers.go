package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// ThrowTilesAllResponse function is used to throw success response of All data in JSON format
func ThrowTilesAllResponse(successMessage string, responseData []entities.DashboardtilesresponseEntity, w http.ResponseWriter, success bool) {
	var response = entities.DashboardtilesAllResponse{}
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

func ThrowTabButtonAllResponse(successMessage string, responseData entities.Dashboardbuttontab, w http.ResponseWriter, success bool) {
	var response = entities.DashboardbuttontabAllResponse{}
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
func ThrowRecordnamesResponse(successMessage string, responseData string, w http.ResponseWriter, success bool) {
	var response = entities.RecordnamesResponse{}
	response.Success = success
	response.Message = successMessage
	response.Recordnames = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetTilesnames(w http.ResponseWriter, req *http.Request) {
	var data = entities.DashboardtilesinputEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.GetTilesnames(&data)
		ThrowTilesAllResponse(msg, data, w, success)
	}
}

func GetTabsButtonnames(w http.ResponseWriter, req *http.Request) {
	var data = entities.DashboardtilesinputEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.GetTabsButtonnames(&data)
		ThrowTabButtonAllResponse(msg, data, w, success)
	}
}

func Getrecordname(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordnamesEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)

		//if(len(responseError)==0){
		data, success, _, msg := models.Getrecordname(&data)
		ThrowRecordnamesResponse(msg, data, w, success)
	}
}
