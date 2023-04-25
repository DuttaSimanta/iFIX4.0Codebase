package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowMstrecordtermsAllResponse(successMessage string, responseData entities.MstrecordtermsEntities, w http.ResponseWriter, success bool) {
    var response = entities.MstrecordtermsResponse{}
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

func ThrowMstrecordtermsListResponse(successMessage string, responseData []entities.TermsEntity, w http.ResponseWriter, success bool) {
    var response = entities.TermsResponse{}
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


func ThrowMstrecordtermsIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.MstrecordtermsResponseInt{}
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


func InsertMstrecordterms(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordtermsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertMstrecordterms(&data)
        ThrowMstrecordtermsIntResponse(msg, data, w, success)
    }
}


func GetAllMstrecordterms(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordtermsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMstrecordterms(&data)
        ThrowMstrecordtermsAllResponse(msg, data, w, success)
    }
}

func GetListMstrecordterms(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordtermsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetListMstrecordterms(&data)
        ThrowMstrecordtermsListResponse(msg, data, w, success)
    }
}


func DeleteMstrecordterms(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordtermsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteMstrecordterms(&data)
        ThrowMstrecordtermsIntResponse(msg, 0, w, success)
    }
}


func UpdateMstrecordterms(w http.ResponseWriter, req *http.Request) {
    var data = entities.MstrecordtermsEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateMstrecordterms(&data)
        ThrowMstrecordtermsIntResponse(msg, 0, w, success)
    }
}


