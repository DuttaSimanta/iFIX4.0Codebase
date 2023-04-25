package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowClientsupportgroupholidayAllResponse(successMessage string, responseData entities.ClientsupportgroupholidayEntities, w http.ResponseWriter, success bool) {
	var response = entities.ClientsupportgroupholidayResponse{}
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

func ThrowClientsupportgroupholidayIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.ClientsupportgroupholidayResponseInt{}
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

func ThrowSupportgrpnameAllResponse(successMessage string, responseData []entities.SupportgrpEntity, w http.ResponseWriter, success bool) {
	var response = entities.SupportgrpResponse{}
	response.Success = success
	response.Message = successMessage
	response.Values = responseData
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func InsertClientsupportgroupholiday(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupholidayEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertClientsupportgroupholiday(&data)
		ThrowClientsupportgroupholidayIntResponse(msg, data, w, success)
	}
}

func GetAllClientsupportgroupholiday(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupholidayEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllClientsupportgroupholiday(&data)
		ThrowClientsupportgroupholidayAllResponse(msg, data, w, success)
	}
}

func DeleteClientsupportgroupholiday(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupholidayEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteClientsupportgroupholiday(&data)
		ThrowClientsupportgroupholidayIntResponse(msg, 0, w, success)
	}
}

func UpdateClientsupportgroupholiday(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupholidayEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateClientsupportgroupholiday(&data)
		ThrowClientsupportgroupholidayIntResponse(msg, 0, w, success)
	}
}

func GetAllSupportgrpname(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupholidayEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllSupportgrpname(&data)
		ThrowSupportgrpnameAllResponse(msg, data, w, success)
	}
}
