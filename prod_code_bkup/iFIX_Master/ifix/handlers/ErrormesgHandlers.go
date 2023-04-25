package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowErrormesgAllResponse(successMessage string, responseData entities.ErrormesgEntities, w http.ResponseWriter, success bool) {
    var response = entities.ErrormesgResponse{}
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


func ThrowErrormesgIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.ErrormesgResponseInt{}
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


func InsertErrormesg(w http.ResponseWriter, req *http.Request) {
    var data = entities.ErrormesgEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertErrormesg(&data)
        ThrowErrormesgIntResponse(msg, data, w, success)
    }
}


func GetAllErrormesg(w http.ResponseWriter, req *http.Request) {
    var data = entities.ErrormesgEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllErrormesg(&data)
        ThrowErrormesgAllResponse(msg, data, w, success)
    }
}


func DeleteErrormesg(w http.ResponseWriter, req *http.Request) {
    var data = entities.ErrormesgEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteErrormesg(&data)
        ThrowErrormesgIntResponse(msg, 0, w, success)
    }
}


func UpdateErrormesg(w http.ResponseWriter, req *http.Request) {
    var data = entities.ErrormesgEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateErrormesg(&data)
        ThrowErrormesgIntResponse(msg, 0, w, success)
    }
}


