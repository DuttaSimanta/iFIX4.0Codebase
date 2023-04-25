package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

func ThrowPortgrouplevelAllResponse(successMessage string, responseData entities.SupportgrouplevelEntities, w http.ResponseWriter, success bool) {
	var response = entities.SupportgrouplevelResponse{}
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

// func ThrowPortgrouplevelIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
// 	var response = entities.SupportgrouplevelResponseInt{}
// 	response.Success = success
// 	response.Message = successMessage
// 	response.Details = responseData
// 	jsonResponse, jsonError := json.Marshal(response)
// 	if jsonError != nil {
// 		logger.Log.Fatal("Internel Server Error")
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonResponse)
// }

// func InsertPortgrouplevel(w http.ResponseWriter, req *http.Request) {
// 	var data = entities.SupportgrouplevelEntity{}
// 	jsonError := data.FromJSON(req.Body)
// 	if jsonError != nil {
// 		log.Print(jsonError)
// 		logger.Log.Println(jsonError)
// 		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
// 	} else {
// 		data, success, _, msg := models.InsertPortgrouplevel(&data)
// 		ThrowPortgrouplevelIntResponse(msg, data, w, success)
// 	}
// }

func GetAllPortgrouplevel(w http.ResponseWriter, req *http.Request) {
	// var data = entities.SupportgrouplevelEntity{}
	// jsonError := data.FromJSON(req.Body)
	// if jsonError != nil {
	// 	log.Print(jsonError)
	// 	logger.Log.Println(jsonError)
	// 	entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	// } else {
	data, success, _, msg := models.GetAllPortgrouplevel()
	ThrowPortgrouplevelAllResponse(msg, data, w, success)
	//}
}

// func DeletePortgrouplevel(w http.ResponseWriter, req *http.Request) {
// 	var data = entities.SupportgrouplevelEntity{}
// 	jsonError := data.FromJSON(req.Body)
// 	if jsonError != nil {
// 		log.Print(jsonError)
// 		logger.Log.Println(jsonError)
// 		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
// 	} else {
// 		success, _, msg := models.DeletePortgrouplevel(&data)
// 		ThrowPortgrouplevelIntResponse(msg, 0, w, success)
// 	}
// }

// func UpdatePortgrouplevel(w http.ResponseWriter, req *http.Request) {
// 	var data = entities.SupportgrouplevelEntity{}
// 	jsonError := data.FromJSON(req.Body)
// 	if jsonError != nil {
// 		log.Print(jsonError)
// 		logger.Log.Println(jsonError)
// 		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
// 	} else {
// 		success, _, msg := models.UpdatePortgrouplevel(&data)
// 		ThrowPortgrouplevelIntResponse(msg, 0, w, success)
// 	}
// }
