package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMststatetypeAllResponse(successMessage string, responseData entities.MststatetypeEntities, w http.ResponseWriter, success bool) {
    var response = entities.MststatetypeResponse{}
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


func ThrowMststatetypeIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MststatetypeResponseInt{}
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


func InsertMststatetype(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststatetypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMststatetype(&data)
        ThrowMststatetypeIntResponse(msg, data, w, success)
    }
}


func GetAllMststatetype(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststatetypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMststatetype(&data)
        ThrowMststatetypeAllResponse(msg, data, w, success)
    }
}


func DeleteMststatetype(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststatetypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMststatetype(&data)
        ThrowMststatetypeIntResponse(msg, 0, w, success)
    }
}


func UpdateMststatetype(w http.ResponseWriter, req *http.Request) {
    var data = entities.MststatetypeEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMststatetype(&data)
        ThrowMststatetypeIntResponse(msg, 0, w, success)
    }
}


