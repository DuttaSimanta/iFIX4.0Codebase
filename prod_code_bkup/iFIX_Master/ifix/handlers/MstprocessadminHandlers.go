package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstprocessadminAllResponse(successMessage string, responseData entities.MstprocessadminEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstprocessadminResponse{}
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


func ThrowMstprocessadminIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstprocessadminResponseInt{}
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


func InsertMstprocessadmin(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstprocessadminEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstprocessadmin(&data)
        ThrowMstprocessadminIntResponse(msg, data, w, success)
    }
}


func GetAllMstprocessadmin(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstprocessadminEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstprocessadmin(&data)
        ThrowMstprocessadminAllResponse(msg, data, w, success)
    }
}


func DeleteMstprocessadmin(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstprocessadminEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstprocessadmin(&data)
        ThrowMstprocessadminIntResponse(msg, 0, w, success)
    }
}


func UpdateMstprocessadmin(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstprocessadminEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstprocessadmin(&data)
        ThrowMstprocessadminIntResponse(msg, 0, w, success)
    }
}


