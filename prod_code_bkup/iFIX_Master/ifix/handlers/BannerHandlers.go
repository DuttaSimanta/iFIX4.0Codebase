package handlers


import (
  "iFIX/ifix/entities"
  "iFIX/ifix/models"
  "iFIX/ifix/logger"
  "log"
  "net/http"
  "encoding/json"
  )

func ThrowBannerAllResponse(successMessage string, responseData entities.BannerEntities, w http.ResponseWriter, success bool) {
    var response = entities.BannerResponse{}
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


func ThrowBannerIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
    var response = entities.BannerResponseInt{}
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

func ThrowBannerMessageResponse(successMessage string, responseData []entities.BannerMessageEntity, w http.ResponseWriter, success bool) {
    var response = entities.BannerResponseMessage{}
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


func InsertBanner(w http.ResponseWriter, req *http.Request) {
    var data = entities.BannerEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.InsertBanner(&data)
        ThrowBannerIntResponse(msg, data, w, success)
    }
}


func GetAllBanner(w http.ResponseWriter, req *http.Request) {
    var data = entities.BannerEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllBanner(&data)
        ThrowBannerAllResponse(msg, data, w, success)
    }
}


func DeleteBanner(w http.ResponseWriter, req *http.Request) {
    var data = entities.BannerEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.DeleteBanner(&data)
        ThrowBannerIntResponse(msg, 0, w, success)
    }
}


func UpdateBanner(w http.ResponseWriter, req *http.Request) {
    var data = entities.BannerEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateBanner(&data)
        ThrowMststateIntResponse(msg, 0, w, success)
    }
}


func GetAllMessage(w http.ResponseWriter, req *http.Request) {
    var data = entities.BannerEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        data, success, _, msg := models.GetAllMessage(&data)
        ThrowBannerMessageResponse(msg, data, w, success)
    }
}

func UpdateBannerSequence(w http.ResponseWriter, req *http.Request) {
    var data = entities.BannerEntity{}
    jsonError := data.FromJSON(req.Body)
    if jsonError != nil {
        log.Print(jsonError)
        logger.Log.Println(jsonError)
        entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
    } else {
        success, _, msg := models.UpdateBannerSequence(&data)
        ThrowMststateIntResponse(msg, 0, w, success)
    }
}