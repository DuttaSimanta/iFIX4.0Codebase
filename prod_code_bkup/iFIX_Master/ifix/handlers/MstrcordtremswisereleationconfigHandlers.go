package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstrcordtremswisereleationconfigAllResponse(successMessage string, responseData entities.MstrcordtremswisereleationconfigEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstrcordtremswisereleationconfigResponse{}
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

func ThrowMstrcordtremswisereleationconfigIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstrcordtremswisereleationconfigResponseInt{}
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

func ThrowrecordreleationAllResponse(successMessage string, responseData []entities.Recordreleationdetails, w http.ResponseWriter, success bool) {
	var response = entities.RecordreleationdetailsAllResponse{}
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

func ThrowrecordtermnamesAllResponse(successMessage string, responseData []entities.Recordtermnames, w http.ResponseWriter, success bool) {
	var response = entities.RecordtermnamesAllResponse{}
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

func InsertMstrcordtremswisereleationconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstrcordtremswisereleationconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstrcordtremswisereleationconfig(&data)
		ThrowMstrcordtremswisereleationconfigIntResponse(msg, data, w, success)
	}
}

func GetAllMstrcordtremswisereleationconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstrcordtremswisereleationconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstrcordtremswisereleationconfig(&data)
		ThrowMstrcordtremswisereleationconfigAllResponse(msg, data, w, success)
	}
}

func DeleteMstrcordtremswisereleationconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstrcordtremswisereleationconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstrcordtremswisereleationconfig(&data)
		ThrowMstrcordtremswisereleationconfigIntResponse(msg, 0, w, success)
	}
}

func UpdateMstrcordtremswisereleationconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstrcordtremswisereleationconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstrcordtremswisereleationconfig(&data)
		ThrowMstrcordtremswisereleationconfigIntResponse(msg, 0, w, success)
	}
}

func GetRecordreleationnames(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstrcordtremswisereleationconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetRecordreleationnames(&data)
		ThrowrecordreleationAllResponse(msg, data, w, success)
	}
}

func GetRecordtermnames(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstrcordtremswisereleationconfigEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetRecordtermnames(&data)
		ThrowrecordtermnamesAllResponse(msg, data, w, success)
	}
}
