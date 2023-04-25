package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowClientholidayAllResponse(successMessage string, responseData entities.ClientholidayEntities, w http.ResponseWriter, success bool) {
    var response = entities.ClientholidayResponse{}
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


func ThrowClientholidayIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.ClientholidayResponseInt{}
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


func InsertClientholiday(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientholidayEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertClientholiday(&data)
        ThrowClientholidayIntResponse(msg, data, w, success)
    }
}


func GetAllClientholiday(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientholidayEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllClientholiday(&data)
        ThrowClientholidayAllResponse(msg, data, w, success)
    }
}


func DeleteClientholiday(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientholidayEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteClientholiday(&data)
        ThrowClientholidayIntResponse(msg, 0, w, success)
    }
}


func UpdateClientholiday(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientholidayEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateClientholiday(&data)
        ThrowClientholidayIntResponse(msg, 0, w, success)
    }
}


