package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstcategorytaskmapAllResponse(successMessage string, responseData entities.MstcategorytaskmapEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstcategorytaskmapResponse{}
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


func ThrowMstcategorytaskmapIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstcategorytaskmapResponseInt{}
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


func InsertMstcategorytaskmap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstcategorytaskmapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstcategorytaskmap(&data)
        ThrowMstcategorytaskmapIntResponse(msg, data, w, success)
    }
}


func GetAllMstcategorytaskmap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstcategorytaskmapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstcategorytaskmap(&data)
        ThrowMstcategorytaskmapAllResponse(msg, data, w, success)
    }
}


func DeleteMstcategorytaskmap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstcategorytaskmapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstcategorytaskmap(&data)
        ThrowMstcategorytaskmapIntResponse(msg, 0, w, success)
    }
}


func UpdateMstcategorytaskmap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstcategorytaskmapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstcategorytaskmap(&data)
        ThrowMstcategorytaskmapIntResponse(msg, 0, w, success)
    }
}


