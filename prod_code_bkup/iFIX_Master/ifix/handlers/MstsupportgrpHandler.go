package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowmstsupportgrpAllResponse(successMessage string, responseData entities.MstsupportgrpEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstsupportgrpResponse{}
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

func ThrowmstsupportgrpIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstsupportgrpResponseInt{}
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

func ThrowmstsupportgrpbycopyableAllResponse(successMessage string, responseData []entities.MstsupportgrpbycopyableEntity, w http.ResponseWriter, success bool) {
	var response = entities.MstsupportgrpbycopyableResponse{}
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

func Insertmstsupportgrp(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstsupportgrpEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Insertmstsupportgrp(&data)
		ThrowmstsupportgrpIntResponse(msg, data, w, success)
	}
}

func GetAllmstsupportgrp(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstsupportgrpEntity{}
	jsonError := data.FromJSON(req.Body)
	//logger.Log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllmstsupportgrp(&data)
		ThrowmstsupportgrpAllResponse(msg, data, w, success)
	}
}

func Deletemstsupportgrp(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstsupportgrpEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.Deletemstsupportgrp(&data)
		ThrowmstsupportgrpIntResponse(msg, 0, w, success)
	}
}

func Updatemstsupportgrp(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstsupportgrpEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.Updatemstsupportgrp(&data)
		ThrowmstsupportgrpIntResponse(msg, 0, w, success)
	}
}

func GetAllmstsupportgrpbycopyable(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstsupportgrpEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllmstsupportgrpbycopyable(&data)
		ThrowmstsupportgrpbycopyableAllResponse(msg, data, w, success)
	}
}
func Getallcreatedsupportgrp(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstsupportgrpEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getallcreatedsupportgrp(&data)
		ThrowmstsupportgrpbycopyableAllResponse(msg, data, w, success)
	}
}
