package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstrecordfieldAllResponse(successMessage string, responseData entities.MstrecordfieldEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstrecordfieldResponse{}
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


func ThrowMstrecordfieldIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstrecordfieldResponseInt{}
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


func InsertMstrecordfield(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordfieldEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstrecordfield(&data)
        ThrowMstrecordfieldIntResponse(msg, data, w, success)
    }
}


func GetAllMstrecordfield(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordfieldEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstrecordfield(&data)
        ThrowMstrecordfieldAllResponse(msg, data, w, success)
    }
}


func DeleteMstrecordfield(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordfieldEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstrecordfield(&data)
        ThrowMstrecordfieldIntResponse(msg, 0, w, success)
    }
}


func UpdateMstrecordfield(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordfieldEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstrecordfield(&data)
        ThrowMstrecordfieldIntResponse(msg, 0, w, success)
    }
}


