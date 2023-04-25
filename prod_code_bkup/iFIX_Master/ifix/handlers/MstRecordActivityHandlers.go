package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowActivitydescAllResponse(successMessage string, responseData []entities.Activitydesces, w http.ResponseWriter, success bool) {
	var response = entities.ActivitydescesResponse{}
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

func ThrowMstRecordActivityAllResponse(successMessage string, responseData entities.MstRecordActivityEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstRecordActivityResponse{}
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

func ThrowMstRecordActivityIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstRecordActivityResponseInt{}
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

func AddMstRecordActivityCopy(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstRecordActivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.AddMstRecordActivityCopy(&data)
		ThrowMstRecordActivityIntResponse(msg, data, w, success)
	}
}
func AddMstRecordActivity(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstRecordActivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.AddMstRecordActivity(&data)
		ThrowMstRecordActivityIntResponse(msg, data, w, success)
	}
}

func GetAllMstRecordActivity(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstRecordActivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstRecordActivity(&data)
		ThrowMstRecordActivityAllResponse(msg, data, w, success)
	}
}

func DeleteMstRecordActivity(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstRecordActivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstRecordActivity(&data)
		ThrowMstRecordActivityIntResponse(msg, 0, w, success)
	}
}

func UpdateMstRecordActivity(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstRecordActivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstRecordActivity(&data)
		ThrowMstRecordActivityIntResponse(msg, 0, w, success)
	}
}
func GetOrgWiseActivitydesc(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstRecordActivityEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetOrgWiseActivitydesc(&data)
		ThrowActivitydescAllResponse(msg, data, w, success)
	}
}
