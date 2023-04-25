package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstslatimezoneAllResponse(successMessage string, responseData entities.MstslatimezoneEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstslatimezoneResponse{}
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


func ThrowMstslatimezoneIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstslatimezoneResponseInt{}
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


func InsertMstslatimezone(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslatimezoneEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstslatimezone(&data)
        ThrowMstslatimezoneIntResponse(msg, data, w, success)
    }
}


func GetAllMstslatimezone(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslatimezoneEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstslatimezone(&data)
        ThrowMstslatimezoneAllResponse(msg, data, w, success)
    }
}


func DeleteMstslatimezone(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslatimezoneEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstslatimezone(&data)
        ThrowMstslatimezoneIntResponse(msg, 0, w, success)
    }
}


func UpdateMstslatimezone(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslatimezoneEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstslatimezone(&data)
        ThrowMstslatimezoneIntResponse(msg, 0, w, success)
    }
}


