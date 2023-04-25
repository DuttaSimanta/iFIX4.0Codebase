package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstSupportGroupWorkingHoursAllResponse(successMessage string, responseData entities.MstSupportGroupWorkingHoursEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstSupportGroupWorkingHoursResponse{}
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

func ThrowMstSupportGroupWorkingHoursIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstSupportGroupWorkingHoursResponseInt{}
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

func ThrowSupportGroupWiseWorkingHoursResponse(successMessage string, responseData []entities.MstSupportGroupWorkingHoursresponseEntity, w http.ResponseWriter, success bool) {
	var response = entities.SupportGroupWiseWorkingHoursResponse{}
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

func InsertMstSupportGroupWorkingHours(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstSupportGroupWorkingHoursEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstSupportGroupWorkingHours(&data)
		ThrowMstSupportGroupWorkingHoursIntResponse(msg, data, w, success)
	}
}

func GetAllMstSupportGroupWorkingHoursk(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstSupportGroupWorkingHoursEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstSupportGroupWorkingHours(&data)
		ThrowMstSupportGroupWorkingHoursAllResponse(msg, data, w, success)
	}
}

func DeleteMstSupportGroupWorkingHours(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstSupportGroupWorkingHoursEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstSupportGroupWorkingHours(&data)
		ThrowMstSupportGroupWorkingHoursIntResponse(msg, 0, w, success)
	}
}

func UpdateMstSupportGroupWorkingHours(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstSupportGroupWorkingHoursUpdateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstSupportGroupWorkingHours(&data)
		ThrowMstSupportGroupWorkingHoursIntResponse(msg, 0, w, success)
	}
}

func GetSupportGroupWiseWorkingHours(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstSupportGroupWorkingHoursEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetSupportGroupWiseWorkingHours(&data)
		ThrowSupportGroupWiseWorkingHoursResponse(msg, data, w, success)
	}
}
