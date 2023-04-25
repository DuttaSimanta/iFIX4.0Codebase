package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstslaresponsiblesupportgroupAllResponse(successMessage string, responseData entities.MstslaresponsiblesupportgroupEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstslaresponsiblesupportgroupResponse{}
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

func ThrowMstslanamesAllResponse(successMessage string, responseData []entities.Mstslanames, w http.ResponseWriter, success bool) {
	var response = entities.MstslanamesResponse{}
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

func ThrowMstslaresponsiblesupportgroupIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstslaresponsiblesupportgroupResponseInt{}
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

func InsertMstslaresponsiblesupportgroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslaresponsiblesupportgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstslaresponsiblesupportgroup(&data)
		ThrowMstslaresponsiblesupportgroupIntResponse(msg, data, w, success)
	}
}

func GetAllMstslaresponsiblesupportgroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslaresponsiblesupportgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstslaresponsiblesupportgroup(&data)
		ThrowMstslaresponsiblesupportgroupAllResponse(msg, data, w, success)
	}
}

func DeleteMstslaresponsiblesupportgroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslaresponsiblesupportgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstslaresponsiblesupportgroup(&data)
		ThrowMstslaresponsiblesupportgroupIntResponse(msg, 0, w, success)
	}
}

func UpdateMstslaresponsiblesupportgroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslaresponsiblesupportgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstslaresponsiblesupportgroup(&data)
		ThrowMstslaresponsiblesupportgroupIntResponse(msg, 0, w, success)
	}
}

func GetAllSlanames(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslaresponsiblesupportgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllSlanames(&data)
		ThrowMstslanamesAllResponse(msg, data, w, success)
	}
}

func GetFullfillmentcriteriaid(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslaresponsiblesupportgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetFullfillmentcriteriaid(&data)
		ThrowMstslaresponsiblesupportgroupIntResponse(msg, data, w, success)
	}
}
