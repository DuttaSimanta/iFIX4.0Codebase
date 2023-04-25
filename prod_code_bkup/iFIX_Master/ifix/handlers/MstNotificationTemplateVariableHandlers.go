package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMstTemplateVariableAllResponse(successMessage string, responseData entities.MstTemplateVariableEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstTemplateVariableResponse{}
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

func ThrowMstTemplateVariableListResponse(successMessage string, responseData []entities.MstTemplateVariableEntityList, w http.ResponseWriter, success bool) {
	var response = entities.MstTemplateVariableListResponse{}
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

func ThrowMstTemplateVariableIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstTemplateVariableResponseInt{}
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

func AddMstTemplateVariable(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstTemplateVariableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.AddMstTemplateVariable(&data)
		ThrowMstTemplateVariableIntResponse(msg, data, w, success)
	}
}
func AddMstTemplateVariablecopy(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstTemplateVariableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.AddMstTemplateVariablecopy(&data)
		ThrowMstTemplateVariableIntResponse(msg, data, w, success)
	}
}

func GetAllMstTemplateVariable(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstTemplateVariableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstTemplateVariable(&data)
		ThrowMstTemplateVariableAllResponse(msg, data, w, success)
	}
}

func DeleteMstTemplateVariable(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstTemplateVariableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstTemplateVariable(&data)
		ThrowMstTemplateVariableIntResponse(msg, 0, w, success)
	}
}

func UpdateMstTemplateVariable(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstTemplateVariableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstTemplateVariable(&data)
		ThrowMstTemplateVariableIntResponse(msg, 0, w, success)
	}
}

func GetAllMstTemplateVariableList(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstTemplateVariableEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstTemplateVariableList(&data)
		ThrowMstTemplateVariableListResponse(msg, data, w, success)
	}
}
