package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowClientmoduleurlroleuserAllResponse(successMessage string, responseData entities.ClientmoduleurlroleuserEntities, w http.ResponseWriter, success bool) {
    var response = entities.ClientmoduleurlroleuserResponse{}
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


func ThrowClientmoduleurlroleuserIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.ClientmoduleurlroleuserResponseInt{}
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


func InsertClientmoduleurlroleuser(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientmoduleurlroleuserEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertClientmoduleurlroleuser(&data)
        ThrowClientmoduleurlroleuserIntResponse(msg, data, w, success)
    }
}


func GetAllClientmoduleurlroleuser(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientmoduleurlroleuserEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllClientmoduleurlroleuser(&data)
        ThrowClientmoduleurlroleuserAllResponse(msg, data, w, success)
    }
}


func DeleteClientmoduleurlroleuser(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientmoduleurlroleuserEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteClientmoduleurlroleuser(&data)
        ThrowClientmoduleurlroleuserIntResponse(msg, 0, w, success)
    }
}


func UpdateClientmoduleurlroleuser(w http.ResponseWriter, req *http.Request) {
    var data = entities.ClientmoduleurlroleuserEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateClientmoduleurlroleuser(&data)
        ThrowClientmoduleurlroleuserIntResponse(msg, 0, w, success)
    }
}


