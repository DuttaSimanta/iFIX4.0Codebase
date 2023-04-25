package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstRecorddiffpriorityAllResponse(successMessage string, responseData entities.MstRecorddiffpriorityEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstRecorddiffpriorityResponse{}
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


func ThrowMstRecorddiffpriorityIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstRecorddiffpriorityResponseInt{}
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


func AddMstRecorddiffpriority(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstRecorddiffpriorityEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.AddMstRecorddiffpriority(&data)
        ThrowMstRecorddiffpriorityIntResponse(msg, data, w, success)
    }
}


func GetAllMstRecorddiffpriority(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstRecorddiffpriorityEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstRecorddiffpriority(&data)
        ThrowMstRecorddiffpriorityAllResponse(msg, data, w, success)
    }
}


func DeleteMstRecorddiffpriority(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstRecorddiffpriorityEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstRecorddiffpriority(&data)
        ThrowMstRecorddiffpriorityIntResponse(msg, 0, w, success)
    }
}


func UpdateMstRecorddiffpriority(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstRecorddiffpriorityEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstRecorddiffpriority(&data)
        ThrowMstRecorddiffpriorityIntResponse(msg, 0, w, success)
    }
}
