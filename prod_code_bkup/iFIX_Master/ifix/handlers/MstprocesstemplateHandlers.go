package handlers

import (
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

/*func ThrowMstprocessAllResponse(successMessage string, responseData entities.MstprocessEntities, w http.ResponseWriter, success bool) {
	var response = entities.MstprocessResponse{}
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

func ThrowMstprocessIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.MstprocessResponseInt{}
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
func InsertMstprocesstemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstprocessEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertMstprocesstemplatewithtransaction(&data)
		ThrowMstprocessIntResponse(msg, data, w, success)
	}
}

func GetAllMstprocesstemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstprocessEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllMstprocesstemplate(&data)
		ThrowMstprocessAllResponse(msg, data, w, success)
	}
}

func DeleteMstprocesstemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstprocessEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteMstprocesstemplatewithtransaction(&data)
		ThrowMstprocessIntResponse(msg, 0, w, success)
	}
}

func UpdateMstprocesstemplate(w http.ResponseWriter, req *http.Request) {
	var data = entities.MstprocessEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateMstprocesstemplatewithtransaction(&data)
		ThrowMstprocessIntResponse(msg, 0, w, success)
	}
}
