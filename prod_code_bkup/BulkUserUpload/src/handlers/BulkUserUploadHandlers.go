package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"src/entities"
	Logger "src/logger"
	BulkUserUploadService "src/models"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
func ThrowBlankResponse(w http.ResponseWriter, req *http.Request) {
	entities.ThrowJSONResponse(entities.BlankPathCheckResponse(), w)
}

func BulkUserUpload(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.APIResponse
	var errResponse entities.ErrorResponse
	var payload map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &payload)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	Logger.Log.Println("Payload====>", payload)

	notificationErr := BulkUserUploadService.BulkUserUploadUsingExcel(payload)

	if notificationErr != nil {
		log.Println(notificationErr)
		errResponse.Status = false
		errResponse.Message = notificationErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Users Uploaded Successfully"
		entities.ThrowJSONResponse(successResponse, w)
		return
	}

}

func BulkUserDownload(w http.ResponseWriter, req *http.Request) {
	Logger.Log.Println("BulkUserUpload====>")
	var successResponse entities.APIResponseDownload
	var errResponse entities.ErrorResponse
	Logger.Log.Println("BulkUserUpload====>1")
	//var result map[string]interface{}
	var result = entities.MstDownloadUser{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	jsonErr := json.Unmarshal(body, &result)
	Logger.Log.Println("BulkUserUpload====>2")
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	Logger.Log.Println("Payload====>", result)

	// roleId := int64(result["roleid"].(float64))
	// usertype := string(result["usertype"].(string))

	Logger.Log.Println("Payload1====>", result)
	originalFileName, uploadedFileName, downloadErr := BulkUserUploadService.BulkUserDownload(result.ClientID, result.OrgID, result.Groupid)
	if downloadErr != nil {
		Logger.Log.Println(downloadErr)
		errResponse.Status = false
		errResponse.Message = downloadErr.Error()
		Logger.Log.Println(errResponse.Message)
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Bulk User Downloaded Successfully"
		successResponse.OriginalFileName = originalFileName
		successResponse.UploadedFileName = uploadedFileName
		entities.ThrowJSONDownloadResponse(successResponse, w)
		return
	}

}
