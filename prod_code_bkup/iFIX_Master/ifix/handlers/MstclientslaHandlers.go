package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstclientslaAllResponse(successMessage string, responseData entities.MstclientslaEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstclientslaResponse{}
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

func ThrowMstclientslaIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstclientslaResponseInt{}
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

func ThrowMstclientslanamesResponse(successMessage string, responseData []entities.Mstslaname, w http.ResponseWriter, success bool) {
	var response = entities.MstslanameResponse{}
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


func InsertMstclientsla(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstclientslaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstclientsla(&data)
		ThrowMstclientslaIntResponse(msg, data, w, success)
	}
}

func GetAllMstclientsla(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstclientslaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstclientsla(&data)
		ThrowMstclientslaAllResponse(msg, data, w, success)
	}
}

func DeleteMstclientsla(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstclientslaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstclientsla(&data)
		ThrowMstclientslaIntResponse(msg, 0, w, success)
	}
}

func UpdateMstclientsla(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstclientslaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstclientsla(&data)
		ThrowMstclientslaIntResponse(msg, 0, w, success)
	}
}

func GetSlanames(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstclientslaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetSlanames(&data)
		ThrowMstclientslanamesResponse(msg, data, w, success)
	}
}
