package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMststateAllResponse(successMessage string, responseData entities.MststateEntities, w http.ResponseWriter, success bool) {
    var response = entities.MststateResponse{}
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


func ThrowMststateIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MststateResponseInt{}
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


func InsertMststate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMststate(&data)
        ThrowMststateIntResponse(msg, data, w, success)
    }
}


func GetAllMststate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMststate(&data)
        ThrowMststateAllResponse(msg, data, w, success)
    }
}


func DeleteMststate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMststate(&data)
        ThrowMststateIntResponse(msg, 0, w, success)
    }
}


func UpdateMststate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststateEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMststate(&data)
        ThrowMststateIntResponse(msg, 0, w, success)
    }
}


