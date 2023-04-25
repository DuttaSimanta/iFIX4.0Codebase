package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMapCategoryWithKeywordAllResponse(successMessage string, responseData entities.MapCategoryWithKeywordEntities, w http.ResponseWriter, success bool) {
	var response = entities.MapCategoryWithKeywordResponse{}
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

func ThrowMapCategoryWithKeywordIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MapCategoryWithKeywordResponseInt{}
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
func ThrowGetkeywordResponse(successMessage string, responseData []entities.Getkeyword, w http.ResponseWriter, success bool) {
	var response = entities.GetkeywordResponse{}
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

func InsertMapCategoryWithKeyword(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapCategoryWithKeywordEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMapCategoryWithKeyword(&data)
		ThrowMapCategoryWithKeywordIntResponse(msg, data, w, success)
	}
}

func GetAllMapCategoryWithKeyword(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapCategoryWithKeywordEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMapCategoryWithKeyword(&data)
		ThrowMapCategoryWithKeywordAllResponse(msg, data, w, success)
	}
}

func DeleteMapCategoryWithKeyword(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapCategoryWithKeywordEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMapCategoryWithKeyword(&data)
		ThrowMapCategoryWithKeywordIntResponse(msg, 0, w, success)
	}
}

func UpdateMapCategoryWithKeyword(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapCategoryWithKeywordEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMapCategoryWithKeyword(&data)
		ThrowMapCategoryWithKeywordIntResponse(msg, 0, w, success)
	}
}

// func GetAllkeyword(w http.ResponseWriter, req *http.Request) {
// 	var data = entities.MapCategoryWithKeywordEntity{}
// 	jsonError := data.FromJSON(req.Body)
// 	if jsonError != nil {
// 		log.Print(jsonError)
// 		logger.Log.Println(jsonError)
// 		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
// 	} else {
// 		data, success, _, msg := models.GetAllkeyword(&data)
// 		ThrowGetkeywordResponse(msg, data, w, success)
// 	}
// }
func GetAllCategoryvalue(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapCategoryWithKeywordEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllCategoryvalue(&data)
		ThrowGetkeywordResponse(msg, data, w, success)
	}
}
