package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowSlaTermEntryAllResponse(successMessage string, responseData entities.SlaTermEntryEntities, w http.ResponseWriter, success bool) {
	var response = entities.SlaTermEntryResponse{}
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

func ThrowSlaTermEntryIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.SlaTermEntryResponseInt{}
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

func AddSlaTermEntry(w http.ResponseWriter, req *http.Request) {
	var data = entities.SlaTermEntryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.AddSlaTermEntry(&data)
		ThrowSlaTermEntryIntResponse(msg, data, w, success)
	}
}

func GetAllSlaTermEntry(w http.ResponseWriter, req *http.Request) {
	var data = entities.SlaTermEntryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllSlaTermEntry(&data)
		ThrowSlaTermEntryAllResponse(msg, data, w, success)
	}
}

func DeleteSlaTermEntry(w http.ResponseWriter, req *http.Request) {
	var data = entities.SlaTermEntryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteSlaTermEntry(&data)
		ThrowSlaTermEntryIntResponse(msg, 0, w, success)
	}
}

// func UpdateSlaTermEntry(w http.ResponseWriter, req *http.Request) {
//     var data = entities.SlaTermEntryEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         success, _, msg := models.UpdateSlaTermEntry(&data)
//         ThrowSlaTermEntryIntResponse(msg, 0, w, success)
//     }
// }
