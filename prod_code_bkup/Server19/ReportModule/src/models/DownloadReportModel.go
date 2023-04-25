package models

import (
	// FileUtils "src/fileutils"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"src/dao"
	"src/entities"
	"src/fileutils"
	"src/logger"
	"time"
	// Excel "github.com/tealeg/xlsx"
)

func JsonToExcelConverter(tz *entities.ResultSetRequestEntity) (bool, error, string) {
	///t := entities.APIResponseDownload{}
	contextPath, contextPatherr := os.Getwd()
	props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		logger.Log.Println(err)
		return false, err, "Unable to Get URL From utility.ReadPropertiesFile"
	}
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		return false, contextPatherr, "Contextpath ERROR"
	}

	sendData, err := json.Marshal(tz)
	if err != nil {
		logger.Log.Println(err)
		return false, err, "Unable to marshal data"
	}
	logger.Log.Println("Create Ticket Json=======>   ", string(sendData))

	resp, err := http.Post(props["Getqueryresulturl"], "application/json", bytes.NewBuffer(sendData))
	if err != nil {
		logger.Log.Println(err)
		return false, err, "Unable to Request For Data"
	}
	var result entities.JsonToExcelResponse
	respBody, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		logger.Log.Println(err1)
		return false, err, "Unable to read response data"
	}
	err2 := json.Unmarshal(respBody, &result)
	if err2 != nil {
		logger.Log.Println(err2)
		return false, err, "Unable to Unmarshal data"
	}
	if !result.Success {
		logger.Log.Println("Unsuccess To Get Record")
		return false, err, "Something Went Wrong"
	}
	if len(result.Details.RequestResultsetData) == 0 {
		logger.Log.Println("No Data Is There")
		return false, err, "No Records Found"
	}
	//logger.Log.Println(result)
	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return false, err, "Something Went Wrong"
		}
		db = dbcon
	}
	err, clientorg := dao.Getclientadnorgbyuser(tz.Userid, db)
	currentTime := time.Now()
	datetime := currentTime.Format("2006-01-02 15:04:05")
	datetime, err = dao.Getexacttime(clientorg[0].Clientid, clientorg[0].Mstorgnhirarchyid, datetime, db)
	logger.Log.Println("SEE ITT:", clientorg[0].Clientid, clientorg[0].Mstorgnhirarchyid, datetime, err)
	_, filePath, err := fileutils.JsonToExcelConverter(result, tz, datetime)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	OriginalFileName, UploadedFileName, err := fileutils.FileUploadAPICall(1, 1, props["fileUploadUrl"], filePath)
	if err != nil {
		logger.Log.Println("Error while downloading", "-", err)
	}
	downloadlist := entities.DownloadlistEntity{}
	downloadlist.OriginalFileName = OriginalFileName
	downloadlist.UploadedFileName = UploadedFileName
	downloadlist.Refuserid = tz.Userid

	sendDatabyte, err := json.Marshal(downloadlist)
	if err != nil {
		logger.Log.Println(err)
		return false, err, "Unable to marshal data"
	}
	logger.Log.Println("Create Ticket Json=======>   ", string(sendDatabyte))
	logger.Log.Println(props["insertdownloadlisturl"], "application/json", bytes.NewBuffer(sendDatabyte))
	resp1, err1 := http.Post(props["insertdownloadlisturl"], "application/json", bytes.NewBuffer(sendDatabyte))
	if err1 != nil {
		logger.Log.Println(err1)
		return false, err, "Unable to Request For Data"
	}
	var result1 entities.ReportDownloadResponseInt
	respBody1, err1 := ioutil.ReadAll(resp1.Body)
	if err1 != nil {
		logger.Log.Println(err1)
		return false, err, "Unable to read response data"
	}
	err3 := json.Unmarshal(respBody1, &result1)
	if err3 != nil {
		logger.Log.Println(err3)
		return false, err3, "Unable to Unmarshal data"
	}
	if !result1.Success {
		logger.Log.Println("Unsuccess To Get Record")
		return false, err3, "Something Went Wrong"
	}
	//fmt.Println(t.OriginalFileName, t.UploadedFileName)
	return true, nil, ""
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
