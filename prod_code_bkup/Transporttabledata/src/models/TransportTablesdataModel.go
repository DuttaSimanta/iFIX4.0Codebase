package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"src/config"
	"src/dao"
	"src/entities"
	"src/fileutils"
	"src/logger"

	"io/ioutil"
	"os"
	"time"
)

// var lock = &sync.Mutex{}

func DownloadTablesdata(page *entities.Transporttable) (string, string, bool, error, string) {
	logger.Log.Println("In side DownloadTablesdata Model......Export started")
	t := []entities.ResultEntity{}
	contextPath, contextPatherr := os.Getwd() //getContextPath()
	logger.Log.Print("contextpath->", contextPath)
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		return "", "", false, contextPatherr, "Contextpath error"
	}
	//t := entities.RecordtypeEntities{}
	props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		logger.Log.Println(err)
		return "", "", false, err, "Unable to Get URL From utility.ReadPropertiesFile"
	}
	// lock.Lock()
	// defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return "", "", false, err, "Database connection failure"
	}
	dataAccess := dao.DbConn{DB: db}
	clientname, orgname, err1 := dataAccess.GetClientOrg(page)
	if err1 != nil {
		logger.Log.Println("fail to get client/org name:", err1)
		return "", "", false, err1, "Sql prepare statement error"
	}
	for i := 0; i < len(page.Table); i++ {
		values, err1 := dataAccess.GetSchema(page, i)
		if err1 != nil {
			logger.Log.Println("Fail to get schema:", err1)
			return "", "", false, err1, "Sql prepare statement error"
		}
		logger.Log.Println(values)
		result, err1 := dataAccess.GetTableData(values, page, i)
		if err1 != nil {
			if err1 == errors.New("Not Having Table Type") {
				return "", "", false, err1, "This tabletype is not configured"
			}
			logger.Log.Println("fail to get table data", err1)
			return "", "", false, err1, "Table data select statement error"
		}
		//logger.Log.Println(result)
		temp := entities.ResultEntity{}

		temp.Values = result
		temp.Tablename = page.Table[i].Tablename
		temp.Tabletype = page.Table[i].Tabletype
		t = append(t, temp)
	}
	var response = entities.ResultofalltableEntity{}
	response.Status = true
	response.Message = ""
	response.Response = t
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		logger.Log.Fatal("Internel Server Error")
	}
	currentTime := time.Now()
	datetime := currentTime.Format("2006.01.02 15:04:05")
	filePath := contextPath + "/resource/downloads/" + clientname + "_" + orgname + "_" + datetime + ".json"
	err = ioutil.WriteFile(filePath, jsonResponse, 0644)
	if err != nil {
		logger.Log.Fatal("Json file writing error")
	}
	OriginalFileName, UploadedFileName, err := fileutils.FileUploadAPICall(1, 1, props["fileUploadUrl"], filePath)
	if err != nil {
		logger.Log.Println("Error while downloading", "-", err)
	}
	logger.Log.Println("File Name: " + clientname + "_" + orgname + "_" + datetime + ".json" + " is uploaded in blob")
	logger.Log.Println("Export End")

	return OriginalFileName, UploadedFileName, true, err, ""

}

func UploadTablesdata(tz *entities.Uploadentity) (bool, error, string) {
	logger.Log.Println("In side UploadTablesdata Model......Import Started")
	//t := entities.RecordtypeEntities{}
	// t := []entities.ResultEntity{}
	contextPath, contextPatherr := os.Getwd()
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		return false, contextPatherr, "conetxtpath error"
	}
	filePath := contextPath + "/resource/downloads/" + tz.OriginalFileName
	fileDownloadErr := fileutils.DownloadFileFromUrl(tz.Clientid, tz.Mstorgnhirarchyid, tz.OriginalFileName, tz.UploadedFileName, filePath)
	if fileDownloadErr != nil {
		fmt.Println("Dowlloaderror")
		logger.Log.Println(fileDownloadErr)
		return false, fileDownloadErr, "File download error"
	}
	logger.Log.Println(filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Log.Println(err)
	}
	data := entities.ResultofalltableEntity{}
	err = json.Unmarshal(content, &data)
	page := data.Response
	// lock.Lock()
	// defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Database connection failure"
	}

	logger.Log.Println(len(page))
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	dataAccess1 := dao.TxConn{TX: tx}

	for i := 0; i < len(page); i++ {
		getschemafortable := entities.Transporttable{}
		tableentity := entities.Tablesandtype{page[i].Tablename, 0}
		getschemafortable.Table = append(getschemafortable.Table, tableentity)
		logger.Log.Println("hii")
		column, err1 := dataAccess.GetSchema(&getschemafortable, 0)
		if err1 != nil {
			tx.Rollback()
			return false, err1, "Sql prepare statement error"
		}
		logger.Log.Println(column)
		if page[i].Tabletype == 2 {
			err := dataAccess1.DeleteTabledata(page[i].Tablename, page[i].Values[0])
			if err != nil {
				tx.Rollback()
				logger.Log.Println("Insert fail", err)
				return false, err, "Insert statement error"
			}
		}
		for j := 0; j < len(page[i].Values); j++ {
			var params []interface{}
			var insertno []string
			var updatecolumn []string

			for k := 0; k < len(column); k++ {
				params = append(params, (page[i].Values[j][column[k]]))
				insertno = append(insertno, "?")
				updatecolumn = append(updatecolumn, column[k]+"= ?")
			}
			logger.Log.Println(params)
			count, err1 := dataAccess.CheckDuplicateData(page[i].Values[j]["id"], page[i].Tablename)
			if err1 != nil {
				tx.Rollback()
				logger.Log.Println("Duplicate checking error", err1)
				return false, err1, "Checkduplicate data error"
			}
			if count == 0 || page[i].Tabletype == 2 {
				err := dataAccess1.InsertTabledata(column, params, insertno, page[i].Tablename)
				if err != nil {
					tx.Rollback()
					logger.Log.Println("Insert fail", err)
					return false, err, "Insert statement error"
				}
			} else if tz.Isupdate == 1 {
				params = append(params, page[i].Values[j]["id"])
				err := dataAccess1.UpdateTabledata(updatecolumn, params, insertno, page[i].Tablename)
				if err != nil {
					tx.Rollback()
					logger.Log.Println("update fail ", err)
					return false, err, "Update statement error"
				}
			}
		}
	}
	tx.Commit()
	logger.Log.Println(tz.OriginalFileName + " File is imported")
	logger.Log.Println("Import Ended")

	return true, err, ""
}
