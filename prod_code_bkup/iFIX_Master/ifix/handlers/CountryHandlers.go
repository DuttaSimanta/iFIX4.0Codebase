package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

func ThrowCountryAllResponse(successMessage string, responseData entities.CountryEntities, w http.ResponseWriter, success bool) {
	var response = entities.CountryResponse{}
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

// func ThrowCountryIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
//     var response = entities.CountryResponseInt{}
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

// func InsertCountry(w http.ResponseWriter, req *http.Request) {
//     var data = entities.CountryEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         data, success, _, msg := models.InsertCountry(&data)
//         ThrowCountryIntResponse(msg, data, w, success)
//     }
// }

func GetAllCountry(w http.ResponseWriter, req *http.Request) {
	// var data = entities.CountryEntity{}
	// jsonError := data.FromJSON(req.Body)
	// if jsonError != nil {
	//     log.Print(jsonError)
	//     logger.Log.Println(jsonError)
	//     entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	// } else {
	data, success, _, msg := models.GetAllCountry()
	ThrowCountryAllResponse(msg, data, w, success)
	//}
}

// func DeleteCountry(w http.ResponseWriter, req *http.Request) {
//     var data = entities.CountryEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         success, _, msg := models.DeleteCountry(&data)
//         ThrowCountryIntResponse(msg, 0, w, success)
//     }
// }

// func UpdateCountry(w http.ResponseWriter, req *http.Request) {
//     var data = entities.CountryEntity{}
//     jsonError := data.FromJSON(req.Body)
//     if jsonError != nil {
//         log.Print(jsonError)
//         logger.Log.Println(jsonError)
//         entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
//     } else {
//         success, _, msg := models.UpdateCountry(&data)
//         ThrowCountryIntResponse(msg, 0, w, success)
//     }
// }
