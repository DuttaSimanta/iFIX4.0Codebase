package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstslastateAllResponse(successMessage string, responseData entities.MstslastateEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstslastateResponse{}
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


func ThrowMstslastateIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstslastateResponseInt{}
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


func InsertMstslastate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslastateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstslastate(&data)
        ThrowMstslastateIntResponse(msg, data, w, success)
    }
}


func GetAllMstslastate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslastateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstslastate(&data)
        ThrowMstslastateAllResponse(msg, data, w, success)
    }
}


func DeleteMstslastate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslastateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstslastate(&data)
        ThrowMstslastateIntResponse(msg, 0, w, success)
    }
}


func UpdateMstslastate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstslastateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstslastate(&data)
        ThrowMstslastateIntResponse(msg, 0, w, success)
    }
}


