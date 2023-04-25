package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstslastartendcriteriaAllResponse(successMessage string, responseData entities.MstslastartendcriteriaEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstslastartendcriteriaResponse{}
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

func ThrowMstslastartendcriteriaIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstslastartendcriteriaResponseInt{}
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

func ThrowMstslanameAllResponse(successMessage string, responseData []entities.MstslanameagaistworkflowEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstslanameagaistworkflowEntityResponse{}
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

func InsertMstslastartendcriteria(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslastartendcriteriaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstslastartendcriteria(&data)
		ThrowMstslastartendcriteriaIntResponse(msg, data, w, success)
	}
}

func GetAllMstslastartendcriteria(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslastartendcriteriaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstslastartendcriteria(&data)
		ThrowMstslastartendcriteriaAllResponse(msg, data, w, success)
	}
}

func DeleteMstslastartendcriteria(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslastartendcriteriaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstslastartendcriteria(&data)
		ThrowMstslastartendcriteriaIntResponse(msg, 0, w, success)
	}
}

func UpdateMstslastartendcriteria(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslastartendcriteriaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstslastartendcriteria(&data)
		ThrowMstslastartendcriteriaIntResponse(msg, 0, w, success)
	}
}

func GetSlanameagainstworkflowid(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslastartendcriteriaEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetSlanameagainstworkflowid(&data)
		ThrowMstslanameAllResponse(msg, data, w, success)
	}
}
