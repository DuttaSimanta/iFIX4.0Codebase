package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	model "src/entities"
	Logger "src/logger"
	CategoryService "src/models"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func ThrowBlankResponse(w http.ResponseWriter, req *http.Request) {
	model.ThrowJSONResponse(model.BlankPathCheckResponse(), w)
}

func BulkCategoryUpload(w http.ResponseWriter, req *http.Request) {

	Logger.Log.Println("BulkCategoryUpload====>")

	var successResponse model.APIResponse
	var errResponse model.ErrorResponse

	var result map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &result)
	Logger.Log.Println("Payload====>", result)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	Logger.Log.Println("Payload====>", result)

	clientID := int64(result["clientid"].(float64))
	orgID := int64(result["mstorgnhirarchyid"].(float64))
	recordDiffTypeID := int64(result["recorddifftypeid"].(float64))
	recordDiffID := int64(result["recorddiffid"].(float64))
	originalFileName := result["originalfilename"].(string)
	uploadedFileName := result["uploadedfilename"].(string)

	uploadErr := CategoryService.BulkCategoryUpload(clientID, orgID, recordDiffTypeID, recordDiffID, originalFileName, uploadedFileName)
	if uploadErr != nil {
		log.Println(uploadErr)
		errResponse.Status = false
		errResponse.Message = uploadErr.Error()
		//log.Println(errResponse.Message )
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Bulk Category Upload Successful"
		model.ThrowJSONResponse(successResponse, w)
		return
	}

}

func BulkCategoryDownload(w http.ResponseWriter, req *http.Request) {
	log.Println("BulkCategoryUpload====>")
	var successResponse model.APIResponseDownload
	var errResponse model.ErrorResponse
	log.Println("BulkCategoryUpload====>1")
	var result map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		Logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	jsonErr := json.Unmarshal(body, &result)
	log.Println("BulkCategoryUpload====>2")
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	log.Println("Payload====>", result)

	clientID := int64(result["clientid"].(float64))
	orgID := int64(result["mstorgnhirarchyid"].(float64))
	recordDiffID := int64(result["recorddiffid"].(float64))
	/*recordDiffTypeID := int64(result["recorddifftypeid"].(float64))
	recordDiffID := int64(result["recorddiffid"].(float64))
	var url string = result["url"].(string)*/
	log.Println("Payload====>", result)
	originalFileName, uploadedFileName, downloadErr := CategoryService.BulkCategoryDownload(clientID, orgID, recordDiffID)
	if downloadErr != nil {
		log.Println(downloadErr)
		errResponse.Status = false
		errResponse.Message = downloadErr.Error()
		log.Println(errResponse.Message)
		model.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Bulk Category Downloaded Successfully"
		successResponse.OriginalFileName = originalFileName
		successResponse.UploadedFileName = uploadedFileName
		model.ThrowJSONDownloadResponse(successResponse, w)
		return
	}

}
