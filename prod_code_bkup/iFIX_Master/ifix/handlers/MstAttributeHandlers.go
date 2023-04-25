package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowAttributeAllResponse(successMessage string, responseData []entities.Attributes, w http.ResponseWriter, success bool) {
	var response = entities.AttributesResponse{}
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

func ThrowMstAttributeAllResponse(successMessage string, responseData entities.MstAttributeEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstAttributeResponse{}
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

func ThrowMstAttributeIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstAttributeResponseInt{}
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

// func AddMstAttributeCopy(w http.ResponseWriter, req *http.Request) {
// 	var data = entities.MstAttributeEntity{}
// 	jsonError := data.FromJSON(req.Body)
// 	if jsonError != nil {
// 		log.Print(jsonError)
// 		logger.Log.Println(jsonError)
// 		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
// 	} else {
// 		data, success, _, msg := models.AddMstAttributeCopy(&data)
// 		ThrowMstAttributeIntResponse(msg, data, w, success)
// 	}
// }
func AddMstAttribute(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstAttributeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.AddMstAttribute(&data)
		ThrowMstAttributeIntResponse(msg, data, w, success)
	}
}

func GetAllMstAttribute(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstAttributeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstAttribute(&data)
		ThrowMstAttributeAllResponse(msg, data, w, success)
	}
}

func DeleteMstAttribute(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstAttributeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstAttribute(&data)
		ThrowMstAttributeIntResponse(msg, 0, w, success)
	}
}

func UpdateMstAttribute(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstAttributeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstAttribute(&data)
		ThrowMstAttributeIntResponse(msg, 0, w, success)
	}
}
func GetMstAttribute(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstAttributeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetMstAttribute(&data)
		ThrowAttributeAllResponse(msg, data, w, success)
	}
}
