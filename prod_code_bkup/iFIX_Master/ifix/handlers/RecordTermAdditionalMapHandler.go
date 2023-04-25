package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowGetAdditionalTabAllResponse(successMessage string, responseData []entities.AdditionalTabEntity, w http.ResponseWriter, success bool) {
	var response = entities.AdditionalTabResponse{}
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
func ThrowRecordTermAdditionalMapAllResponse(successMessage string, responseData entities.RecordTermAdditionalMapEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecordTermAdditionalMapResponse{}
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

func ThrowRecordTermAdditionalMapIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.RecordTermAdditionalMapResponseInt{}
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

func InsertRecordTermAdditionalMap(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordTermAdditionalMapEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertRecordTermAdditionalMap(&data)
		ThrowRecordTermAdditionalMapIntResponse(msg, data, w, success)
	}
}

func GetAllRecordTermAdditionalMap(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordTermAdditionalMapEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllRecordTermAdditionalMap(&data)
		ThrowRecordTermAdditionalMapAllResponse(msg, data, w, success)
	}
}

func DeleteRecordTermAdditionalMap(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordTermAdditionalMapEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteRecordTermAdditionalMap(&data)
		ThrowRecordTermAdditionalMapIntResponse(msg, 0, w, success)
	}
}

// func UpdateRecordTermAdditionalMap(w http.ResponseWriter, req *http.Request) {
// 	var data = entities.RecordTermAdditionalMapEntity{}
// 	jsonError := data.FromJSON(req.Body)
// 	if jsonError != nil {
// 		log.Print(jsonError)
// 		logger.Log.Println(jsonError)
// 		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
// 	} else {
// 		success, _, msg := models.UpdateRecordTermAdditionalMap(&data)
// 		ThrowRecordTermAdditionalMapIntResponse(msg, 0, w, success)
// 	}
// }
func GetAdditionalTab(w http.ResponseWriter, req *http.Request) {
	// var data = entities.RecordTermAdditionalMapEntity{}
	// jsonError := data.FromJSON(req.Body)
	// if jsonError != nil {
	// 	log.Print(jsonError)
	// 	logger.Log.Println(jsonError)
	// 	entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	// } else {
	data, success, _, msg := models.GetAdditionalTab()
	ThrowGetAdditionalTabAllResponse(msg, data, w, success)
	// }
}
