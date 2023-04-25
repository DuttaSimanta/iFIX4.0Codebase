package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowDashboardQueryCopyAllResponse(successMessage string, responseData entities.DashboardQueryCopyEntities, w http.ResponseWriter, success bool) {
	var response = entities.DashboardQueryCopyResponse{}
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

func ThrowDashboardQueryCopyIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.DashboardQueryCopyResponseInt{}
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

func AddDashboardQueryCopy(w http.ResponseWriter, req *http.Request) {
	var data = entities.DashboardQueryCopyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.AddDashboardQueryCopy(&data)
		ThrowDashboardQueryCopyIntResponse(msg, data, w, success)
	}
}

func GetAllDashboardQueryCopy(w http.ResponseWriter, req *http.Request) {
	var data = entities.DashboardQueryCopyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllDashboardQueryCopy(&data)
		ThrowDashboardQueryCopyAllResponse(msg, data, w, success)
	}
}

func DeleteDashboardQueryCopy(w http.ResponseWriter, req *http.Request) {
	var data = entities.DashboardQueryCopyEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteDashboardQueryCopy(&data)
		ThrowDashboardQueryCopyIntResponse(msg, 0, w, success)
	}
}

// func UpdateDashboardQueryCopy(w http.ResponseWriter, req *http.Request) {
//     var data = entities.DashboardQueryCopyEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         success, _, msg := models.UpdateDashboardQueryCopy(&data)
//         ThrowDashboardQueryCopyIntResponse(msg, 0, w, success)
//     }
// }
