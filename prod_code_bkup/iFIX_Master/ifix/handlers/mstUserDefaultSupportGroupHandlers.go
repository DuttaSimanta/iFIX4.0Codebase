package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstUserDefaultSupportGroupAllResponse(successMessage string, responseData entities.MstUserDefaultSupportGroupEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstUserDefaultSupportGroupResponse{}
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

func ThrowMstUserDefaultSupportGroupIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstUserDefaultSupportGroupResponseInt{}
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

func InsertMstUserDefaultSupportGroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserDefaultSupportGroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstUserDefaultSupportGroup(&data)
		ThrowMstUserDefaultSupportGroupIntResponse(msg, data, w, success)
	}
}

func GetAllMstUserDefaultSupportGroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserDefaultSupportGroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstUserDefaultSupportGroup(&data)
		ThrowMstUserDefaultSupportGroupAllResponse(msg, data, w, success)
	}
}

func DeleteMstUserDefaultSupportGroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserDefaultSupportGroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstUserDefaultSupportGroup(&data)
		ThrowMstUserDefaultSupportGroupIntResponse(msg, 0, w, success)
	}
}

func UpdateMstUserDefaultSupportGroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserDefaultSupportGroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstUserDefaultSupportGroup(&data)
		ThrowMstUserDefaultSupportGroupIntResponse(msg, 0, w, success)
	}
}




func MstUserSupportGroupChange(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstUserDefaultSupportGroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.MstUserSupportGroupChange(&data)
		ThrowMstUserDefaultSupportGroupIntResponse(msg, 0, w, success)
	}
}



