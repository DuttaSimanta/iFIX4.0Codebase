package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstsupportgrptermmapAllResponse(successMessage string, responseData entities.MstsupportgrptermmapEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstsupportgrptermmapResponse{}
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


func ThrowMstsupportgrptermmapIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstsupportgrptermmapResponseInt{}
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


func InsertMstsupportgrptermmap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstsupportgrptermmapEntity{}
    jsonError := data.FromJSON(req.Body)
     
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstsupportgrptermmap(&data)
        ThrowMstsupportgrptermmapIntResponse(msg, data, w, success)
    }
}


func GetAllMstsupportgrptermmap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstsupportgrptermmapEntity{}
    
      jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstsupportgrptermmap(&data)
        ThrowMstsupportgrptermmapAllResponse(msg, data, w, success)
    }
}


func DeleteMstsupportgrptermmap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstsupportgrptermmapEntity{}
    jsonError := data.FromJSON(req.Body)
     if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstsupportgrptermmap(&data)
        ThrowMstsupportgrptermmapIntResponse(msg, 0, w, success)
    }
}


func UpdateMstsupportgrptermmap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstsupportgrptermmapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstsupportgrptermmap(&data)
        ThrowMstsupportgrptermmapIntResponse(msg, 0, w, success)
    }
}
