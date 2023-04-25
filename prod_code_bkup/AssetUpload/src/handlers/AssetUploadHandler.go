package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"src/entities"
	Logger "src/logger"
	Asset "src/models"
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

func AssetUpload(w http.ResponseWriter, req *http.Request) {
	Logger.Log.Println("Asset Upload Controller")
	var successResponse entities.APIResponse
	var errResponse entities.ErrorResponse

	var payload map[string]interface{}
	body, err := ioutil.ReadAll(req.Body)
	Logger.Log.Println("Payload====>", payload)
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

	var clientID int64 = int64(payload["clientid"].(float64))
	var orgID int64 = int64(payload["mstorgnhirarchyid"].(float64))
	originalFileName := payload["originalfilename"].(string)
	uploadedFileName := payload["uploadedfilename"].(string)

	// clientID, err := strconv.ParseInt(clientIDReq, 10, 64)
	// if err != nil {
	// 	Logger.Log.Println("ERROR: Please Provide Valid ClientId")
	// 	errResponse.Status = false
	// 	errResponse.Message = errors.New("ERROR: Please Provide Valid ClientId").Error()
	// 	entities.ThrowJSONErrorResponse(errResponse, w)
	// 	return
	// }
	// mstOrgnHirarchyId, err := strconv.ParseInt(mstOrgnHirarchyIdReq, 10, 64)
	// if err != nil {
	// 	Logger.Log.Println("ERROR: Please Provide Valid MstOrgnHirarchyId")
	// 	errResponse.Status = false
	// 	errResponse.Message = errors.New("ERROR: Please Provide Valid MstOrgnHirarchyId").Error()
	// 	entities.ThrowJSONErrorResponse(errResponse, w)
	// 	return
	// }
	// //filename := req.FormValue("filename")
	// //url := req.FormValue("url")

	uploadErr := Asset.AssetUpload(clientID, orgID, originalFileName, uploadedFileName)
	if uploadErr != nil {
		log.Println(uploadErr)
		errResponse.Status = false
		errResponse.Message = uploadErr.Error()
		//log.Println(errResponse.Message )
		entities.ThrowJSONErrorResponse(errResponse, w)
		return
	} else {
		successResponse.Status = true
		successResponse.Message = "Asset Upload Successful"
		entities.ThrowJSONResponse(successResponse, w)
		return
	}

}
