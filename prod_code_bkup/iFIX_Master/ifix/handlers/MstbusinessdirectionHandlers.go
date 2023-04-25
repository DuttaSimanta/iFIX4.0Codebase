package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstbusinessdirectionAllResponse(successMessage string, responseData entities.MstbusinessdirectionEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstbusinessdirectionResponse{}
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


func ThrowMstbusinessdirectionIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstbusinessdirectionResponseInt{}
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


func InsertMstbusinessdirection(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstbusinessdirectionEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstbusinessdirection(&data)
        ThrowMstbusinessdirectionIntResponse(msg, data, w, success)
    }
}


func GetAllMstbusinessdirection(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstbusinessdirectionEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstbusinessdirection(&data)
        ThrowMstbusinessdirectionAllResponse(msg, data, w, success)
    }
}


func DeleteMstbusinessdirection(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstbusinessdirectionEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstbusinessdirection(&data)
        ThrowMstbusinessdirectionIntResponse(msg, 0, w, success)
    }
}


func UpdateMstbusinessdirection(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstbusinessdirectionEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstbusinessdirection(&data)
        ThrowMstbusinessdirectionIntResponse(msg, 0, w, success)
    }
}


