package handlers

import (
	"encoding/json"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func getContextPath() (string, error) {

	wd, err := os.Getwd()
	if err != nil {
		return "", errors.New("ERROR: Unable to get WD")
	}
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	return contextPath, nil
}

// LocationPriority/src/handlers/UploadLocationPriorityHandlers.go

func UserWithGroupAndCategoryUpload(w http.ResponseWriter, req *http.Request) {

	logger.Log.Println("UserWithGroupAndCategoryUpload ====>")

	var successResponse entities.BulkUserWithGroupAndCategoryUploadResponse
	var errResponse entities.BulkUserWithGroupAndCategoryUploadErrorResponse

	data := entities.UserWithGroupAndCategoryForBulkUploadEntity{}

	// var result map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Log.Println(err)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
		entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
		return
	}
	jsonErr := json.Unmarshal(body, &data)
	logger.Log.Println("Payload====>", data)
	if jsonErr != nil {
		logger.Log.Println(jsonErr)
		errResponse.Status = false
		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
		entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
		return
	}

	logger.Log.Println("Payload====>", data)

	// clientID := int64(result["clientid"].(float64))
	// orgID := int64(result["mstorgnhirarchyid"].(float64))
	// recordDiffTypeID := int64(result["recorddifftypeid"].(float64))
	// recordDiffID := int64(result["recorddiffid"].(float64))
	// originalFileName := result["originalfilename"].(string)
	// uploadedFileName := result["uploadedfilename"].(string)
	// uploadErr := models.LocationPriorityUpload(clientID, orgID, recordDiffTypeID, recordDiffID, originalFileName, uploadedFileName)

	uploadErr := models.BulkUserWithGroupAndCategoryUpload(&data)
	if uploadErr != nil {
		log.Println(uploadErr)
		errResponse.Status = false
		errResponse.Message = uploadErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Bulk User With Support Group Uploaded Successfully"
		entities.ThrowBulkUserWithGroupAndCategoryUploadJSONResponse(successResponse, w)
		return
	}

}
