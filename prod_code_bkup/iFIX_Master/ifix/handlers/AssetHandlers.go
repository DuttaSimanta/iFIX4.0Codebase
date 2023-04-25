package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"log"
	"net/http"
)

func ThrowAssetSearchResponse(successMessage string, responseData entities.AssetSearchResEntity, w http.ResponseWriter, success bool) {
	var response = entities.AssetSearchResponse{}
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

func ThrowAssetTypesResponse(successMessage string, responseData []entities.Assettype, w http.ResponseWriter, success bool) {
	var response = entities.AssettypeResponse{}
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

func ThrowAssetByTypeResponse(successMessage string, responseData entities.AssetEntitiesByType, w http.ResponseWriter, success bool) {
	var response = entities.AssetByTypeResponse{}
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

func ThrowAssetDiffValResponse(successMessage string, responseData entities.AssetEntityDiffVals, w http.ResponseWriter, success bool) {
	var response = entities.AssetEntityDiffValResponse{}
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

func ThrowMapAssetWithTypeResponse(successMessage string, responseData entities.AssetMapWithRecordTypes, w http.ResponseWriter, success bool) {
	var response = entities.AssetMapWithRecordTypeResponse{}
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

func ThrowAssetAllResponse(successMessage string, responseData entities.AssetEntities, w http.ResponseWriter, success bool) {
	var response = entities.AssetResponse{}
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

func ThrowAssetIntResponse(successMessage string, responseData int64, w http.ResponseWriter, success bool) {
	var response = entities.AssetResponseInt{}
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

func InsertAsset(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.InsertAsset(&data)
		ThrowAssetIntResponse(msg, data, w, success)
	}
}

func GetAllAsset(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAllAsset(&data)
		ThrowAssetAllResponse(msg, data, w, success)
	}
}

func DeleteAsset(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.DeleteAsset(&data)
		ThrowAssetIntResponse(msg, 0, w, success)
	}
}

func UpdateAsset(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateAsset(&data)
		ThrowAssetIntResponse(msg, 0, w, success)
	}
}

func GetAssetBYType(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data, success, _, msg := models.GetAssetBYType(&data)
		ThrowAssetByTypeResponse(msg, data, w, success)
	}
}

func GetAssetDiffVal(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data1, success, _, msg := models.GetAssetDiffVal(&data)
		ThrowAssetDiffValResponse(msg, data1, w, success)
	}
}

func UpdateAssetDiffVal(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntityDiffValUpdate{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		success, _, msg := models.UpdateAssetDiffVal(&data)
		ThrowAssetIntResponse(msg, 0, w, success)
	}
}

func GetClietWiseAsset(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data1, success, _, msg := models.GetClietWiseAsset(&data)
		ThrowMapAssetWithTypeResponse(msg, data1, w, success)
	}
}

func GetAssetTypes(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data1, success, _, msg := models.GetAssetTypes(&data)
		ThrowAssetTypesResponse(msg, data1, w, success)
	}
}

func GetAssetAttributes(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		data1, success, _, msg := models.GetAssetAttributes(&data)
		ThrowAssetTypesResponse(msg, data1, w, success)
	}
}

func GetAssetByTypeNAtrrValue(w http.ResponseWriter, req *http.Request) {
	var data = entities.AssetSearchEntity{}
	jsonError := data.FromJSON(req.Body)
	if jsonError != nil {
		log.Print(jsonError)
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	}
	data1, success, _, msg := models.GetAssetByTypeNAtrrValue(&data)
	ThrowAssetSearchResponse(msg, data1, w, success)

}
