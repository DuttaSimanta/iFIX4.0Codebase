package handlers

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

/*func ThrowMapprocessstateAllResponse(successMessage string, responseData entities.MapprocessstateEntities, w http.ResponseWriter, success bool) {
	var response = entities.MapprocessstateResponse{}
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

func ThrowMapprocessstateIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MapprocessstateResponseInt{}
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
*/
func InsertMapprocesstemplatestate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapprocessstateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMapprocesstemplatestate(&data)
		ThrowMapprocessstateIntResponse(msg, data, w, success)
	}
}

func GetAllMapprocesstemplatestate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapprocessstateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMapprocesstemplatestate(&data)
		ThrowMapprocessstateAllResponse(msg, data, w, success)
	}
}

func DeleteMapprocesstemplatestate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapprocessstateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMapprocesstemplatestate(&data)
		ThrowMapprocessstateIntResponse(msg, 0, w, success)
	}
}

func UpdateMapprocesstemplatestate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MapprocessstateEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMapprocesstemplatestate(&data)
		ThrowMapprocessstateIntResponse(msg, 0, w, success)
	}
}
