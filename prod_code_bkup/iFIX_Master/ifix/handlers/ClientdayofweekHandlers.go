package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowClientdayofweekAllResponse(successMessage string, responseData entities.ClientdayofweekEntities, w http.ResponseWriter, success bool) {
	var response = entities.ClientdayofweekResponse{}
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

func ThrowClientdayofweekIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.ClientdayofweekResponseInt{}
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

func ThrowClientwisedayofweekResponse(successMessage string, responseData []entities.ClientdayofweekresponseEntity, w http.ResponseWriter, success bool) {
	var response = entities.ClientwisedayofweekResponse{}
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

func InsertClientdayofweek(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientdayofweekEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertClientdayofweek(&data)
		ThrowClientdayofweekIntResponse(msg, data, w, success)
	}
}

func GetAllClientdayofweek(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientdayofweekEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllClientdayofweek(&data)
		ThrowClientdayofweekAllResponse(msg, data, w, success)
	}
}

func DeleteClientdayofweek(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientdayofweekEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteClientdayofweek(&data)
		ThrowClientdayofweekIntResponse(msg, 0, w, success)
	}
}

// func UpdateClientdayofweek(w http.ResponseWriter, req *http.Request) {
//     var data = entities.ClientdayofweekEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         success, _, msg := models.UpdateClientdayofweek(&data)
//         ThrowClientdayofweekIntResponse(msg, 0, w, success)
//     }
// }

func GetClientwisedayofweek(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientdayofweekEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetClientwisedayofweek(&data)
		ThrowClientwisedayofweekResponse(msg, data, w, success)
	}
}
