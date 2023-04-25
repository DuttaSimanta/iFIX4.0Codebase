package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMststatetermAllResponse(successMessage string, responseData entities.MststatetermEntities, w http.ResponseWriter, success bool) {
    var response = entities.MststatetermResponse{}
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


func ThrowMststatetermIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MststatetermResponseInt{}
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


func InsertMststateterm(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststatetermEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMststateterm(&data)
        ThrowMststatetermIntResponse(msg, data, w, success)
    }
}


func GetAllMststateterm(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststatetermEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMststateterm(&data)
        ThrowMststatetermAllResponse(msg, data, w, success)
    }
}


func DeleteMststateterm(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststatetermEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMststateterm(&data)
        ThrowMststatetermIntResponse(msg, 0, w, success)
    }
}


func UpdateMststateterm(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststatetermEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMststateterm(&data)
        ThrowMststatetermIntResponse(msg, 0, w, success)
    }
}


