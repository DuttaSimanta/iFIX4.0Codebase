package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstldapAllResponse(successMessage string, responseData entities.MstldapEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstldapResponse{}
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
func ThrowMstldapStringResponse(successMessage string, responseData []string, w http.ResponseWriter, success bool) {
    var response = entities.MstldapfieldResponse{}
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


func ThrowMstldapIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstldapResponseInt{}
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


func AddMstldap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstldapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstldap(&data)
        ThrowMstldapIntResponse(msg, data, w, success)
    }
}


func GetAllMstldap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstldapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstldap(&data)
        ThrowMstldapAllResponse(msg, data, w, success)
    }
}
func Gettabledetails(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstldapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.Gettabledetails(&data)
        ThrowMstldapStringResponse(msg, data, w, success)
    }
}


func DeleteMstldap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstldapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstldap(&data)
        ThrowMstldapIntResponse(msg, 0, w, success)
    }
}


func UpdateMstldap(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstldapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstldap(&data)
        ThrowMstldapIntResponse(msg, 0, w, success)
    }
}
func UpdateMstldapCertificate(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstldapEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstldapCertificate(&data)
        ThrowMstldapIntResponse(msg, 0, w, success)
    }
}
