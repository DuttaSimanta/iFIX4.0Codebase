package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowMsttemplateAllResponse(successMessage string, responseData entities.MsttemplateEntities, w http.ResponseWriter, success bool) {
	var response = entities.MsttemplateResponse{}
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
func ThrowNotificationTemplateAllResponse(successMessage string, responseData entities.MstNotificationTemplateEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstNotificationTemplateResponse{}
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

func ThrowNotificationVariableResponse(successMessage string, responseData []entities.MstNotificationVariable, w http.ResponseWriter, success bool) {
	var response = entities.MstNotificationVariableResponse{}
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

func ThrowMsttemplateIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MsttemplateResponseInt{}
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
func ThrowMsttemplateIntArrResponse(successMessage string, responseData []int64, w http.ResponseWriter, success bool) {
	var response = entities.MsttemplateResponseIntArr{}
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
func ThrowNotificationEventResponse(successMessage string, responseData []entities.MstNotificationEvent, w http.ResponseWriter, success bool) {
	var response = entities.MstNotificationEventResponse{}
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
func GetNotificationEvents(w http.ResponseWriter, req *http.Request) {
	data, success, _, msg := models.GetNotificationEvents()
	ThrowNotificationEventResponse(msg, data, w, success)
}
func InsertNotificationTemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstNotificationTemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertNotificationTemplate(&data)
		ThrowMsttemplateIntArrResponse(msg, data, w, success)
	}
}
func GetAllNotificationTemplates(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstNotificationTemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllNotificationTemplates(&data)
		ThrowNotificationTemplateAllResponse(msg, data, w, success)
	}
}

func GetAllNotificationVariables(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstNotificationTemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllNotificationVariables(&data)
		ThrowNotificationVariableResponse(msg, data, w, success)
	}
}
func InsertMsttemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MsttemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMsttemplate(&data)
		ThrowMsttemplateIntResponse(msg, data, w, success)
	}
}

func GetAllMsttemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MsttemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMsttemplate(&data)
		ThrowMsttemplateAllResponse(msg, data, w, success)
	}
}

func DeleteMsttemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MsttemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMsttemplate(&data)
		ThrowMsttemplateIntResponse(msg, 0, w, success)
	}
}

func DeleteNotificationTemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstNotificationTemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteNotificationTemplate(&data)
		ThrowMsttemplateIntResponse(msg, 0, w, success)
	}
}

func UpdateMsttemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MsttemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMsttemplate(&data)
		ThrowMsttemplateIntResponse(msg, 0, w, success)
	}
}

func UpdateNotificationTemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstNotificationTemplateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateNotificationTemplate(&data)
		ThrowMsttemplateIntResponse(msg, 0, w, success)
	}
}
