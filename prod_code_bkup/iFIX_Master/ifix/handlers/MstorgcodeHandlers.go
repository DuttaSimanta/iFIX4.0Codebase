package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstorgcodewithtoolAllResponse(successMessage string, responseData entities.MstorgcodeEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstorgcodeResponse{}
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

func ThrowMstorgcodewithtoolIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstorgcodeResponseInt{}
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

func ThrowGettoolscodeResponse(successMessage string, responseData []entities.Gettoolscode, w http.ResponseWriter, success bool) {
	var response = entities.GettoolsResponse{}
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

func ThrowGetorgcodeResponse(successMessage string, responseData []entities.Getorgcode, w http.ResponseWriter, success bool) {
	var response = entities.GetorgResponse{}
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

func InsertMstorgcode(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgcodeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstorgcode(&data)
		ThrowMstorgcodewithtoolIntResponse(msg, data, w, success)
	}
}

func GetAllMstorgcode(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgcodeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstorgcode(&data)
		ThrowMstorgcodewithtoolAllResponse(msg, data, w, success)
	}
}

func DeleteMstorgcode(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgcodeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstorgcode(&data)
		ThrowMstorgcodewithtoolIntResponse(msg, 0, w, success)
	}
}

func UpdateMstorgcode(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgcodeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstorgcode(&data)
		ThrowMstorgcodewithtoolIntResponse(msg, 0, w, success)
	}
}

func GetAlltoolvalue(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgcodeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAlltoolvalue(&data)
		ThrowGettoolscodeResponse(msg, data, w, success)
	}
}

func GetAllorgvalue(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstorgcodeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllorgvalue(&data)
		ThrowGetorgcodeResponse(msg, data, w, success)
	}
}
