package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstslafcrecorddiffAllResponse(successMessage string, responseData entities.MstslafcrecorddiffEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstslafcrecorddiffResponse{}
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

func ThrowMstslafcrecorddiffIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstslafcrecorddiffResponseInt{}
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

func ThrowSLAmetertypeAllResponse(successMessage string, responseData []entities.MstslametertypeEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstslametertypeResponse{}
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

func ThrowSLAmetertermnamesAllResponse(successMessage string, responseData []entities.MstslaindicatortermEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstslaindicatortermResponse{}
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

func InsertMstslafcrecorddiff(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslafcrecorddiffEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstslafcrecorddiff(&data)
		ThrowMstslafcrecorddiffIntResponse(msg, data, w, success)
	}
}

func GetAllMstslafcrecorddiff(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslafcrecorddiffEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstslafcrecorddiff(&data)
		ThrowMstslafcrecorddiffAllResponse(msg, data, w, success)
	}
}

func DeleteMstslafcrecorddiff(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslafcrecorddiffEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstslafcrecorddiff(&data)
		ThrowMstslafcrecorddiffIntResponse(msg, 0, w, success)
	}
}

func UpdateMstslafcrecorddiff(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslafcrecorddiffEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstslafcrecorddiff(&data)
		ThrowMstslafcrecorddiffIntResponse(msg, 0, w, success)
	}
}

// 19.04.2021
func GetSLAmetertype(w http.ResponseWriter, req *http.Request) {
	data, success, _, msg := models.GetSLAmetertype()
	ThrowSLAmetertypeAllResponse(msg, data, w, success)

}

func GetSLAtermnames(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstslafcrecorddiffEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetSLAmetermnames(&data)
		ThrowSLAmetertermnamesAllResponse(msg, data, w, success)
	}
}
