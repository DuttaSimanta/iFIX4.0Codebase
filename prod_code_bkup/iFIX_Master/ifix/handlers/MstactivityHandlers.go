package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstactivityAllResponse(successMessage string, responseData entities.MstactivityEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstactivityResponse{}
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
func ThrowMstactivityResponse(successMessage string, responseData []entities.MstactivitySingleEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstactionResponse{}
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

func ThrowMstactivityIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstactivityResponseInt{}
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

func ThrowMstactiontypenamesAllResponse(successMessage string, responseData []entities.MstactiontypeEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstactiontypeResponse{}
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

func InsertMstactivity(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstactivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstactivity(&data)
		ThrowMstactivityIntResponse(msg, data, w, success)
	}
}

func GetAllMstactivity(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstactivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstactivity(&data)
		ThrowMstactivityAllResponse(msg, data, w, success)
	}
}
func Getactivitywithtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstactivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getactivitywithtype(&data)
		ThrowMstactivityResponse(msg, data, w, success)
	}
}

func DeleteMstactivity(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstactivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstactivity(&data)
		ThrowMstactivityIntResponse(msg, 0, w, success)
	}
}

func UpdateMstactivity(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstactivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstactivity(&data)
		ThrowMstactivityIntResponse(msg, 0, w, success)
	}
}

func GetActiontypenames(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstactivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetActiontypenames(&data)
		ThrowMstactiontypenamesAllResponse(msg, data, w, success)
	}
}
