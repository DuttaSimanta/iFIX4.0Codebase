package models

import (
	"bytes"
	"encoding/json"
	"iFIX/ifix/config"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/utility"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func JsonToExcelConverter(tz *entities.ResultSetRequestEntity) (string, string, bool, error, string) {
	logger.Log.Println("Inside JsonToExcelConverter Model")

	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		return "", "", false, err, "Something Went Wrong"
	}

	sendData, err := json.Marshal(tz)
	if err != nil {
		logger.Log.Println(err)
		return "", "", false, err, "Unable to marshal data"
	}
	logger.Log.Println("Create Ticket Json=======>   ", string(sendData))

	resp, err := http.Post(config.GetjsonUrl, "application/json", bytes.NewBuffer(sendData))
	if err != nil {
		logger.Log.Println(err)
		return "", "", false, err, "Unable to Request For Data"
	}
	var result entities.JsonToExcelResponse
	respBody, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		logger.Log.Println(err1)
		return "", "", false, err, "Unable to read response data"
	}
	err2 := json.Unmarshal(respBody, &result)
	if err2 != nil {
		logger.Log.Println(err2)
		return "", "", false, err, "Unable to Unmarshal data"
	}
	if !result.Success {
		logger.Log.Println("Unsuccess To Get Record")
		return "", "", false, err, "Something Went Wrong"
	}
	if len(result.Details.RequestResultsetData) == 0 {
		logger.Log.Println("No Data Is There")
		return "", "", false, err, "No Records Found"
	}
	//logger.Log.Println(result)

	_, filePath, err := utility.JsonToExcelConverter(result, tz)
	if err != nil {
		return "", "", false, err, "Something Went Wrong"
	}
	buf, _ := ioutil.ReadFile(filePath)
	typee := http.DetectContentType(buf)
	// OriginalFileName, UploadedFileName, err := utility.FileUploadAPICall(1, 1, config.FileUploadUrl, filePath)
	// if err != nil {
	// 	logger.Log.Println("Error while downloading", "-", err)
	// }
	//fmt.Println(t.OriginalFileName, t.UploadedFileName)
	var data = entities.FileuploadEntity{}
	data.Clientid = 1
	data.Mstorgnhirarchyid = 1
	// imgbytes := buf.Bytes()
	// OriginalFileName, UploadedFileName, err := utility.FileUploadAPICall(values[0].Clientid, values[0].Mstorgnhirarchyid, props["fileUploadUrl"], filePath)
	dataDetails, success, err, msg := UploadFileWithConn(&data, buf, filePath, typee, db) //utility.FileUploadAPICall(values[0].Clientid, values[0].Mstorgnhirarchyid, props["fileUploadUrl"], filePath)
	if err != nil {
		logger.Log.Println("Error while downloading", "-", err)
		return "", "", false, err, "Something Went Wrong"
	}
	logger.Log.Println("===========FileUploadMessage==============", msg)
	logger.Log.Println("===========FileUploadSuccss==============", success)

	OriginalFileName := filepath.Base(dataDetails.Originalfile)
	UploadedFileName := dataDetails.Filename

	logger.Log.Println("===========OriginalFileName==============", OriginalFileName)
	logger.Log.Println("===========UploadedFileName==============", UploadedFileName)
	// val.OriginalFileName = OriginalFileName
	// val.UploadedFileName = UploadedFileName
	return OriginalFileName, UploadedFileName, true, nil, ""
}
