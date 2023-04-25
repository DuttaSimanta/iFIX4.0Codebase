package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMsttermtypeAllResponse(successMessage string, responseData entities.MsttermtypeEntities, w http.ResponseWriter, success bool) {
    var response = entities.MsttermtypeResponse{}
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


func ThrowMsttermtypeIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MsttermtypeResponseInt{}
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


func InsertMsttermtype(w http.ResponseWriter, req *http.Request) {
    var data = entities.MsttermtypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMsttermtype(&data)
        ThrowMsttermtypeIntResponse(msg, data, w, success)
    }
}


func GetAllMsttermtype(w http.ResponseWriter, req *http.Request) {
    var data = entities.MsttermtypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMsttermtype(&data)
        ThrowMsttermtypeAllResponse(msg, data, w, success)
    }
}


func DeleteMsttermtype(w http.ResponseWriter, req *http.Request) {
    var data = entities.MsttermtypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMsttermtype(&data)
        ThrowMsttermtypeIntResponse(msg, 0, w, success)
    }
}


func UpdateMsttermtype(w http.ResponseWriter, req *http.Request) {
    var data = entities.MsttermtypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMsttermtype(&data)
        ThrowMsttermtypeIntResponse(msg, 0, w, success)
    }
}


