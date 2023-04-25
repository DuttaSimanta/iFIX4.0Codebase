package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecordConfigIncrementAllResponse(successMessage string, responseData entities.RecordConfigIncrementEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecordConfigIncrementResponse{}
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

func ThrowRecordConfigIncrementIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.RecordConfigIncrementResponseInt{}
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

// func ThrowBannerMessageResponse(successMessage string, responseData []entities.BannerMessageEntity, w http.ResponseWriter, success bool) {
//     var response = entities.BannerResponseMessage{}
//     response.Success = success
//     response.Message = successMessage
//     response.Details = responseData
//     jsonResponse, jsonError := json.Marshal(response)
//     if jsonError != nil {
//         logger.Log.Fatal("Internel Server Error")
//     }
//     w.Header().Set("Content-Type", "application/json")
//     w.Write(jsonResponse)
// }

func InsertRecordConfigIncrement(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordConfigIncrementEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertRecordConfigIncrement(&data)
		ThrowRecordConfigIncrementIntResponse(msg, data, w, success)
	}
}

func GetAllRecordConfigIncrement(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordConfigIncrementEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllRecordConfigIncrement(&data)
		ThrowRecordConfigIncrementAllResponse(msg, data, w, success)
	}
}

func DeleteRecordConfigIncrement(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordConfigIncrementEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteRecordConfigIncrement(&data)
		ThrowRecordConfigIncrementIntResponse(msg, 0, w, success)
	}
}

func UpdateRecordConfigIncrement(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordConfigIncrementEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateRecordConfigIncrement(&data)
		ThrowRecordConfigIncrementIntResponse(msg, 0, w, success)
	}
}
