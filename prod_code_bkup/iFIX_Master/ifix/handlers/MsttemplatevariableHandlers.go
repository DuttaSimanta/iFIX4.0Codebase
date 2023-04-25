package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMsttemplatevariableAllResponse(successMessage string, responseData entities.MsttemplatevariableEntities, w http.ResponseWriter, success bool) {
    var response = entities.MsttemplatevariableResponse{}
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


func ThrowMsttemplatevariableIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MsttemplatevariableResponseInt{}
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


func InsertMsttemplatevariable(w http.ResponseWriter, req *http.Request) {
    var data = entities.MsttemplatevariableEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMsttemplatevariable(&data)
        ThrowMsttemplatevariableIntResponse(msg, data, w, success)
    }
}


func GetAllMsttemplatevariable(w http.ResponseWriter, req *http.Request) {
    var data = entities.MsttemplatevariableEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMsttemplatevariable(&data)
        ThrowMsttemplatevariableAllResponse(msg, data, w, success)
    }
}


func DeleteMsttemplatevariable(w http.ResponseWriter, req *http.Request) {
    var data = entities.MsttemplatevariableEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMsttemplatevariable(&data)
        ThrowMsttemplatevariableIntResponse(msg, 0, w, success)
    }
}


func UpdateMsttemplatevariable(w http.ResponseWriter, req *http.Request) {
    var data = entities.MsttemplatevariableEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMsttemplatevariable(&data)
        ThrowMsttemplatevariableIntResponse(msg, 0, w, success)
    }
}


