package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowUserWithGroupAndCategoryAllResponse(successMessage string, responseData entities.UserWithGroupAndCategoryEntities, w http.ResponseWriter, success bool) {
	var response = entities.UserWithGroupAndCategoryResponse{}
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

func ThrowUserWithGroupAndCategoryIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.UserWithGroupAndCategoryResponseInt{}
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

func InsertUserWithGroupAndCategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.UserWithGroupAndCategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertUserWithGroupAndCategory(&data)
		ThrowUserWithGroupAndCategoryIntResponse(msg, data, w, success)
	}
}

func GetAllUserWithGroupAndCategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.UserWithGroupAndCategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllUserWithGroupAndCategory(&data)
		ThrowUserWithGroupAndCategoryAllResponse(msg, data, w, success)
	}
}

func DeleteUserWithGroupAndCategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.UserWithGroupAndCategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteUserWithGroupAndCategory(&data)
		ThrowUserWithGroupAndCategoryIntResponse(msg, 0, w, success)
	}
}

func UpdateUserWithGroupAndCategory(w http.ResponseWriter, req *http.Request) {
	var data = entities.UserWithGroupAndCategoryEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateUserWithGroupAndCategory(&data)
		ThrowUserWithGroupAndCategoryIntResponse(msg, 0, w, success)
	}
}
