package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowClientsupportgroupnewAllResponse(successMessage string, responseData entities.ClientsupportgroupnewEntities, w http.ResponseWriter, success bool) {
	var response = entities.ClientsupportgroupnewResponse{}
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
func ThrowGetsupportgroupbyorgAllResponse(successMessage string, responseData []entities.GetsupportgroupbyorgEntity, w http.ResponseWriter, success bool) {
	var response = entities.GetsupportgroupResponse{}
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

// ThrowGroupSingleResponse function is used to throw success response of All data in JSON format
func ThrowGroupnewSingleResponse(successMessage string, responseData []entities.ClientsupportgroupnewsingleEntity, w http.ResponseWriter, success bool) {
	var response = entities.ClientsupportgroupnewsingleResponse{}
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
func ThrowClientsupportgroupnewIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.ClientsupportgroupnewResponseInt{}
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

func InsertClientsupportgroupnew(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupnewEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertClientsupportgroupnew(&data)
		ThrowClientsupportgroupnewIntResponse(msg, data, w, success)
	}
}

func GetAllClientsupportgroupnew(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupnewEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllClientsupportgroupnew(&data)
		ThrowClientsupportgroupnewAllResponse(msg, data, w, success)
	}
}

func DeleteClientsupportgroupnew(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupnewEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteClientsupportgroupnew(&data)
		ThrowClientsupportgroupnewIntResponse(msg, 0, w, success)
	}
}

func UpdateClientsupportgroupnew(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupnewEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateClientsupportgroupnew(&data)
		ThrowClientsupportgroupnewIntResponse(msg, 0, w, success)
	}
}

func InsertClientsupportgroupfromto(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupnewEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertClientsupportgroupfromto(&data)
		ThrowClientsupportgroupnewIntResponse(msg, data, w, success)
	}
}

func GetAllClientsupportgroupbyclient(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupnewEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllClientsupportgroupbyclient(&data)
		ThrowClientsupportgroupnewAllResponse(msg, data, w, success)
	}
}
func Getsupportgroupbyorg(w http.ResponseWriter, req *http.Request) {
	var data = entities.ClientsupportgroupnewEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.Getsupportgroupbyorg(&data)
		ThrowGetsupportgroupbyorgAllResponse(msg, data, w, success)
	}
}
