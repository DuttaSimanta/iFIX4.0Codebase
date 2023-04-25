package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowRecordtypeAllResponse(successMessage string, responseData entities.RecordtypeEntities, w http.ResponseWriter, success bool) {
	var response = entities.RecordtypeResponse{}
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

func ThrowRecordtypesingleResponse(successMessage string, responseData []entities.Recordtypesingleentity, w http.ResponseWriter, success bool) {
	var response = entities.RecordtypesingleResponse{}
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
func ThrowRecordtypeIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
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

func InsertRecordtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertRecordtype(&data)
		ThrowRecordtypeIntResponse(msg, data, w, success)
	}
}

func Getlabelbydiffid(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getlabelbydiffid(&data)
		ThrowRecordtypesingleResponse(msg, data, w, success)
	}
}
func Getlabelbydiffseq(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getlabelbydiffseq(&data)
		ThrowRecordtypesingleResponse(msg, data, w, success)
	}
}
func Getmappeddiffbyseq(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getmappeddiffbyseq(&data)
		ThrowRecordtypesingleResponse(msg, data, w, success)
	}
}
func Getlablelmappingbydifftype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getlablelmappingbydifftype(&data)
		ThrowRecordtypesingleResponse(msg, data, w, success)
	}
}
func GetAllRecordtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllRecordtype(&data)
		ThrowRecordtypeAllResponse(msg, data, w, success)
	}
}

func DeleteRecordtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteRecordtype(&data)
		ThrowRecordtypeIntResponse(msg, 0, w, success)
	}
}

func UpdateRecordtype(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateRecordtype(&data)
		ThrowRecordtypeIntResponse(msg, 0, w, success)
	}
}

func InsertMaptask(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertTaskmapping(&data)
		ThrowRecordtypeIntResponse(msg, data, w, success)
	}
}

func GetAllMaptask(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllTaskmap(&data)
		ThrowRecordtypeAllResponse(msg, data, w, success)
	}
}

func DeleteMaptask(w http.ResponseWriter, req *http.Request) {
	var data = entities.RecordtypeEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteTaskmap(&data)
		ThrowRecordtypeIntResponse(msg, 0, w, success)
	}
}
