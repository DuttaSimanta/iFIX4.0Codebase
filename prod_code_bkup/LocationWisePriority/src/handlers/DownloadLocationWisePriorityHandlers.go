//SearchUser  method is used for search  user data
package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"src/entities"
	"src/logger"
	"src/models"
)

func BulkLocationWisePriorityDownload(w http.ResponseWriter, req *http.Request) {
	var successResponse entities.APIResponseDownload
	var errResponse entities.ErrorResponse

	var result map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	}

	logger.Log.Println("Payload====>", result)

	clientID := int64(result["clientid"].(float64))
	orgID := int64(result["mstorgnhirarchyid"].(float64))
	recordDiffID := int64(result["recorddiffid"].(float64))
	/*recordDiffTypeID := int64(result["recorddifftypeid"].(float64))
	recordDiffID := int64(result["recorddiffid"].(float64))
	var url string = result["url"].(string)*/

	originalFileName, uploadedFileName, downloadErr := models.BulkLocationWisePriorityDownload(clientID, orgID, recordDiffID)
	if downloadErr != nil {
		log.Println(downloadErr)
		errResponse.Status = false
		errResponse.Message = downloadErr.Error()
		log.Println(errResponse.Message)
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Bulk Category Downloaded Successfully"
		successResponse.OriginalFileName = originalFileName
		successResponse.UploadedFileName = uploadedFileName
		entities.ThrowJSONDownloadResponse(successResponse, w)
		return
	}

}
