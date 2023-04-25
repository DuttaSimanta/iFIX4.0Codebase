package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecorddifferentiationRecResponse(successMessage string, responseData []entities.RecorddifferentionSingle, w http.ResponseWriter, success bool) {
	var response = entities.RecorddifferentionRec{}
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
func ThrowRecorddifferentiationAllResponse(successMessage string, responseData entities.RecorddifferentiationEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecorddifferentiationResponse{}
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


func ThrowRecorddifferentiationIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.ApiResponseInt{}
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

func InsertRecorddifferentiation(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertRecorddifferentiation(&data)
		ThrowRecorddifferentiationIntResponse(msg, data, w, success)
	}
}

func ThrowRecorddifferentiationnameAllResponse(successMessage string, responseData entities.RecorddifferentiationnameEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecorddifferentiationnameResponse{}
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

func GetAllRecorddifferentiation(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllRecorddifferentiation(&data)
		ThrowRecorddifferentiationAllResponse(msg, data, w, success)
	}
}

func DeleteRecorddifferentiation(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteRecorddifferentiation(&data)
		ThrowRecorddifferentiationIntResponse(msg, 0, w, success)
	}
}

func UpdateRecorddifferentiation(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateRecorddifferentiation(&data)
		ThrowRecorddifferentiationIntResponse(msg, 0, w, success)
	}
}

func GetRecorddifferentiationname(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetRecorddifferentiationname(&data)
		ThrowRecorddifferentiationnameAllResponse(msg, data, w, success)
	}
}
func GetRecorddifferentiationbyrecursive(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetRecorddifferentiationbyrecursive(&data)
		ThrowRecorddifferentiationRecResponse(msg, data, w, success)
	}
}
func Getdiffdetailsbyseq(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getdiffdetailsbyseq(&data)
		ThrowRecorddifferentiationRecResponse(msg, data, w, success)
	}
}
func GetRecorddifferentiationbyparent(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetRecorddifferentiationbyparent(&data)
		ThrowRecorddifferentiationRecResponse(msg, data, w, success)
	}
}
func Searchcategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecorddifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Searchcategory(&data)
		ThrowRecorddifferentiationRecResponse(msg, data, w, success)
	}
}
