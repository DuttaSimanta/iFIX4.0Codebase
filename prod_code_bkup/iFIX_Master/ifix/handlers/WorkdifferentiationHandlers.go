package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowWorkdifferentiationvalueResponse(successMessage string, responseData []entities.WorkdifferentiationsingleEntity, w http.ResponseWriter, success bool) {
	var response = entities.WorkdifferentiationsingleResponse{}
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
func ThrowWorkdifferentiationAllResponse(successMessage string, responseData entities.WorkdifferentiationEntities, w http.ResponseWriter, success bool) {
	var response = entities.WorkdifferentiationResponse{}
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

func ThrowWorkdifferentiationIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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
func ThrowWorkinglabelnameAllResponse(successMessage string, responseData entities.WorkinglabelnameEntities, w http.ResponseWriter, success bool) {
	var response = entities.WorkinglabelnameResponse{}
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

func GetWorkinglabelname(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkdifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllWorkinglabelname(&data)
		ThrowWorkinglabelnameAllResponse(msg, data, w, success)
	}
}
func InsertWorkdifferentiation(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkdifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertWorkdifferentiation(&data)
		ThrowWorkdifferentiationIntResponse(msg, data, w, success)
	}
}

func GetAllWorkdifferentiation(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkdifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllWorkdifferentiation(&data)
		ThrowWorkdifferentiationAllResponse(msg, data, w, success)
	}
}
func Getworkdifferentiationvalue(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkdifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getworkdifferentiationvalue(&data)
		ThrowWorkdifferentiationvalueResponse(msg, data, w, success)
	}
}

func DeleteWorkdifferentiation(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkdifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteWorkdifferentiation(&data)
		ThrowWorkdifferentiationIntResponse(msg, 0, w, success)
	}
}

func UpdateWorkdifferentiation(w http.ResponseWriter, req *http.Request) {
	var data = entities.WorkdifferentiationEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateWorkdifferentiation(&data)
		ThrowWorkdifferentiationIntResponse(msg, 0, w, success)
	}
}
