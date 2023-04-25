package handlers

import (
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/models"
	"net/http"
)

// func BulkUserWithGroupAndCategoryDownload(w http.ResponseWriter, req *http.Request) {
// 	var successResponse entities.BulkUserWithGroupAndCategoryDownload
// 	var errResponse entities.BulkUserWithGroupAndCategoryUploadErrorResponse

// 	data := entities.UserWithGroupAndCategoryForBulkUploadEntity{}
// 	body, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		logger.Log.Println(err)
// 		errResponse.Status = false
// 		errResponse.Message = errors.New("ERROR: Not able to fetch Request Data").Error()
// 		entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
// 		return
// 	}
// 	jsonErr := json.Unmarshal(body, &data)
// 	if jsonErr != nil {
// 		logger.Log.Println(jsonErr)
// 		errResponse.Status = false
// 		errResponse.Message = errors.New("ERROR: Json Unmarshal error").Error()
// 		entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
// 		return
// 	}

// 	logger.Log.Println("Payload====>", data)

// 	// clientID := int64(result["clientid"].(float64))
// 	// orgID := int64(result["mstorgnhirarchyid"].(float64))
// 	// recordDiffID := int64(result["recorddiffid"].(float64))
// 	/*recordDiffTypeID := int64(result["recorddifftypeid"].(float64))
// 	recordDiffID := int64(result["recorddiffid"].(float64))
// 	var url string = result["url"].(string)*/

// 	filePath, tz, downloadErr := models.BulkUserWithGroupAndCategoryDownload(&data)
// 	if downloadErr != nil {
// 		log.Println(downloadErr)
// 		errResponse.Status = false
// 		errResponse.Message = downloadErr.Error()
// 		log.Println(errResponse.Message)
// 		entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
// 		return
// 	} else {
// 		// contextPath, contextPatherr := getContextPath()
// 		// if contextPatherr != nil {
// 		// 	logger.Log.Println(contextPatherr)
// 		// 	errResponse.Status = false
// 		// 	errResponse.Message = downloadErr.Error()
// 		// 	log.Println(errResponse.Message)
// 		// 	entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
// 		// 	return
// 		// }
// 		// props, propErr := utility.ReadPropertiesFile(contextPath + "/ifix/resource/application.properties")
// 		// if propErr != nil {
// 		// 	logger.Log.Println("Error while downloading", "-", propErr)
// 		// 	errResponse.Status = false
// 		// 	errResponse.Message = propErr.Error()
// 		// 	log.Println(errResponse.Message)
// 		// 	entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
// 		// 	return
// 		// }
// 		originalFileName, uploadedFileName, fileUploaderr := utility.FileUploadAPICall(tz.Clientid, tz.Mstorgnhirarchyid, config.FileUploadUrl, filePath)
// 		if fileUploaderr != nil {
// 			logger.Log.Println("Error while downloading", "-", fileUploaderr)
// 			errResponse.Status = false
// 			errResponse.Message = fileUploaderr.Error()
// 			log.Println(errResponse.Message)
// 			entities.ThrowBulkUserWithGroupAndCategoryUploadJSONErrorResponse(errResponse, w)
// 			return
// 		}
// 		successResponse.Status = true
// 		successResponse.Message = "Bulk User With Support Group Downloaded Successfully"
// 		successResponse.OriginalFileName = originalFileName
// 		successResponse.UploadedFileName = uploadedFileName
// 		entities.ThrowBulkUserWithGroupAndCategoryDownloadResponse(successResponse, w)
// 		return
// 	}

// }

func ThrowBulkUserWithGroupAndCategoryResponse(successMessage string, OriginalFileName string, UploadedFileName string, w http.ResponseWriter, success bool) {
	var response = entities.APIResponseDownload{}
	response.Status = success
	response.Message = successMessage
	response.OriginalFileName = OriginalFileName
	response.UploadedFileName = UploadedFileName
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func BulkUserWithGroupAndCategoryDownload(w http.ResponseWriter, req *http.Request) {
	var data = entities.UserWithGroupAndCategoryForBulkUploadEntity{}
	jsonError := data.FromJSON(req.Body)

	if jsonError != nil {
		logger.Log.Println(jsonError)
		entities.ThrowJSONResponse(entities.JSONParseErrorResponse(), w)
	} else {
		//responseError = validators.ValidateAddMstClient(data)
		//if(len(responseError)==0){
		OriginalFileName, UploadedFileName, success, _, msg := models.BulkUserWithGroupAndCategoryDownload(&data)
		ThrowBulkUserWithGroupAndCategoryResponse(msg, OriginalFileName, UploadedFileName, w, success)
	}
}
