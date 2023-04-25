package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstbusinessmatrixAllResponse(successMessage string, responseData entities.MstbusinessmatrixEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstbusinessmatrixResponse{}
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

func ThrowMstbusinessmatrixIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstbusinessmatrixResponseInt{}
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

func ThrowLastlevelcategorynamesResponse(successMessage string, responseData []entities.MstlastlevelEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstlastlevelEntityResponse{}
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

func InsertMstbusinessmatrix(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstbusinessmatrixEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstbusinessmatrix(&data)
		ThrowMstbusinessmatrixIntResponse(msg, data, w, success)
	}
}

func GetAllMstbusinessmatrix(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstbusinessmatrixEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstbusinessmatrix(&data)
		ThrowMstbusinessmatrixAllResponse(msg, data, w, success)
	}
}

func DeleteMstbusinessmatrix(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstbusinessmatrixEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstbusinessmatrix(&data)
		ThrowMstbusinessmatrixIntResponse(msg, 0, w, success)
	}
}

func UpdateMstbusinessmatrix(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstbusinessmatrixEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstbusinessmatrix(&data)
		ThrowMstbusinessmatrixIntResponse(msg, 0, w, success)
	}
}

func Checkbusinessmatrixconfig(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstbusinessmatrixEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Checkbusinessmatrixconfig(&data)
		ThrowMstbusinessmatrixIntResponse(msg, data, w, success)
	}
}

func Getlastlevelcategoryname(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstbusinessmatrixEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getlastlevelcategoryname(&data)
		ThrowLastlevelcategorynamesResponse(msg, data, w, success)
	}
}
