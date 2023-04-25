package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstslaentityAllResponse(successMessage string, responseData entities.MstslaentityEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstslaentityResponse{}
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


func ThrowMstslaentityIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstslaentityResponseInt{}
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


func InsertMstslaentity(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslaentityEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstslaentity(&data)
        ThrowMstslaentityIntResponse(msg, data, w, success)
    }
}


func GetAllMstslaentity(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslaentityEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstslaentity(&data)
        ThrowMstslaentityAllResponse(msg, data, w, success)
    }
}


func DeleteMstslaentity(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslaentityEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstslaentity(&data)
        ThrowMstslaentityIntResponse(msg, 0, w, success)
    }
}


func UpdateMstslaentity(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslaentityEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstslaentity(&data)
        ThrowMstslaentityIntResponse(msg, 0, w, success)
    }
}


