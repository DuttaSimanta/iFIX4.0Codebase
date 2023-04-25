package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMapfunctionalitywithgroupAllResponse(successMessage string, responseData entities.MapfunctionalitywithgroupEntities, w http.ResponseWriter, success bool) {
	var response = entities.MapfunctionalitywithgroupResponse{}
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

func ThrowGetOrganizationsAllResponse(successMessage string, responseData entities.OrganizationgrpnameEntities, w http.ResponseWriter, success bool) {
	var response = entities.OrganizationgrpnameResponse{}
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

func ThrowMapfunctionalitywithgroupIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MapfunctionalitywithgroupResponseInt{}
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

func InsertMapfunctionalitywithgroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapfunctionalitywithgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMapfunctionalitywithgroup(&data)
		ThrowMapfunctionalitywithgroupIntResponse(msg, data, w, success)
	}
}

func GetAllMapfunctionalitywithgroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapfunctionalitywithgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMapfunctionalitywithgroup(&data)
		ThrowMapfunctionalitywithgroupAllResponse(msg, data, w, success)
	}
}

func DeleteMapfunctionalitywithgroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapfunctionalitywithgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMapfunctionalitywithgroup(&data)
		ThrowMapfunctionalitywithgroupIntResponse(msg, 0, w, success)
	}
}

func UpdateMapfunctionalitywithgroup(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapfunctionalitywithgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMapfunctionalitywithgroup(&data)
		ThrowMapfunctionalitywithgroupIntResponse(msg, 0, w, success)
	}
}

func GetAllOrganizationgrpnames(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapfunctionalitywithgroupEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllOrganizationgrpnames(&data)
		ThrowGetOrganizationsAllResponse(msg, data, w, success)
	}
}
