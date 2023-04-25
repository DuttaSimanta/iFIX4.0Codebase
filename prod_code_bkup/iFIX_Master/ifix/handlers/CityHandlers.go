package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

func ThrowCityAllResponse(successMessage string, responseData entities.CityEntities, w http.ResponseWriter, success bool) {
	var response = entities.CityResponse{}
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

// func ThrowCityIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
//     var response = entities.CityResponseInt{}
//     response.Success = success
//     response.Message = successMessage
//     response.Details = responseData
//     jsonResponse, jsonError := json.Marshal(response)
//     if jsonError != nil {
//         logger.Log.Fatal("Internel Server Error")
//     }
//     w.Header().Set("Content-Type", "application/json")
//     w.Write(jsonResponse)
// }

// func InsertCity(w http.ResponseWriter, req *http.Request) {
//     var data = entities.CityEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         data, success, _, msg := models.InsertCity(&data)
//         ThrowCityIntResponse(msg, data, w, success)
//     }
// }

func GetAllCity(w http.ResponseWriter, req *http.Request) {
	// var data = entities.CityEntity{}
	// jsonError := data.FromJSON(req.Body)
	// if jsonError != nil {
	// 	log.Print(jsonError)
	// 	logger.Log.Println(jsonError)
	// 	entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	// } else {
	data, success, _, msg := models.GetAllCity()
	ThrowCityAllResponse(msg, data, w, success)
	//}
}

// func DeleteCity(w http.ResponseWriter, req *http.Request) {
//     var data = entities.CityEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         success, _, msg := models.DeleteCity(&data)
//         ThrowCityIntResponse(msg, 0, w, success)
//     }
// }

// func UpdateCity(w http.ResponseWriter, req *http.Request) {
//     var data = entities.CityEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         success, _, msg := models.UpdateCity(&data)
//         ThrowCityIntResponse(msg, 0, w, success)
//     }
// }
