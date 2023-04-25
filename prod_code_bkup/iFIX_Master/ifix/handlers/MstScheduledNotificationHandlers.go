package handlers
 
 
import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )
func ThrowMstScheduledNotificationAllResponse(successMessage string, responseData entities.MstScheduledNotificationEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstScheduledNotificationResponse{}
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


func ThrowMstScheduledNotificationIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstScheduledNotificationResponseInt{}
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

func ThrowGetClientAndOrgWiseclientuserAllResponse(successMessage string, responseData []entities.GetClientAndOrgWiseclientuserEntity, w http.ResponseWriter, success bool) {
    var response = entities.GetClientAndOrgWiseclientuserResponse{}
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

func AddMstScheduledNotification(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstScheduledNotificationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.AddMstScheduledNotification(&data)
        ThrowMstScheduledNotificationIntResponse(msg, data, w, success)
    }
}


func GetMstScheduledNotification(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstScheduledNotificationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstScheduledNotification(&data)
        ThrowMstScheduledNotificationAllResponse(msg, data, w, success)
    }
}


func DeleteMstScheduledNotification(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstScheduledNotificationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstScheduledNotification(&data)
        ThrowMstScheduledNotificationIntResponse(msg, 0, w, success)
    }
}


func UpdateMstScheduledNotification(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstScheduledNotificationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstScheduledNotification(&data)
        ThrowMstScheduledNotificationIntResponse(msg, 0, w, success)
    }
}

func GetClientAndOrgWiseclientuser(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstScheduledNotificationEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetClientAndOrgWiseclientuser(&data)
        ThrowGetClientAndOrgWiseclientuserAllResponse(msg, data, w, success)
    }
}