package models

import (
	"fmt"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"io/ioutil"
	"net/http"
	"path/filepath"

	// "strconv"

	"iFIX/ifix/entities"
	"iFIX/ifix/logger"

	// "iFIX/ifix/utility"

	"github.com/tealeg/xlsx"
)

func BulkUserWithGroupAndCategoryDownload(tz *entities.UserWithGroupAndCategoryForBulkUploadEntity) (string, string, bool, error, string) {
	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		return "", "", false, contextPatherr, "Something Went Wrong"
	}

	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return "", "", false, err, "Something Went Wrong"
	}
	_, txErr := db.Begin()
	if txErr != nil {
		logger.Log.Println(txErr)
		return "", "", false, err, "Something Went Wrong"
	}

	OrgName, ticketTypeName, OrgNameErr := dao.GetOrgName(db, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddiffid)
	if OrgNameErr != nil {
		fmt.Println(OrgNameErr)
		logger.Log.Println(OrgNameErr)
		return "", "", false, err, "Something Went Wrong"
	}
	filePath := contextPath + "/ifix/resource/downloads/" + OrgName + "_" + ticketTypeName + "_" + ".xlsx"
	fmt.Println(tz.Clientid, tz.Mstorgnhirarchyid)
	//defer db.Close()
	headerNames, headerErr := dao.GetTemplateHeaderNamesForValidation(db, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddiffid)
	if headerErr != nil {
		fmt.Println(headerErr)
		logger.Log.Println(headerErr)
		return "", "", false, err, "Something Went Wrong"
	}
	//fmt.Println("Lastrocordidis :", lasRecorddifftypeid)

	// var ids string
	// for j, i := range tz.WorkingCategories {
	// 	if j > 0 {
	// 		ids += ","
	// 	}
	// 	ids += strconv.Itoa(int(i))
	// }

	values, parentCategoryerr := dao.GetUserWithGroupAndCategoryDetails(db, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid)
	if parentCategoryerr != nil {
		logger.Log.Println(parentCategoryerr)
		return "", "", false, err, "Something Went Wrong"
	}
	headerLength := len(headerNames)
	if headerLength == 0 {
		return "", "", false, err, "Something Went Wrong"
	}
	file := xlsx.NewFile()
	sheet, sheetErr := file.AddSheet("Sheet1")
	if sheetErr != nil {
		logger.Log.Print(sheetErr)

		//fmt.Printf(err.Error())
		return "", "", false, err, "Something Went Wrong"
	}
	for i := 0; i <= len(values); i++ {
		logger.Log.Println("ROwCOunt---->", i)
		row := sheet.AddRow()
		if i == 0 {
			for j := 0; j < headerLength; j++ {
				cell := row.AddCell()
				cell.Value = headerNames[j]
			}
		} else {
			// logger.Log.Println("ParentCategorynames====>", parentCategoryNames[i-1])

			// splittedParentCatagories := strings.Split(parentCategoryNames[i-1], "->") //(i-1) because for i=0 headernames is added
			// logger.Log.Println("cat level len====>", headerLength-6)
			// logger.Log.Println("Splitted Length====>", len(splittedParentCatagories))
			//for j := 0; j < headerLength; j++ {
			cell := row.AddCell()
			cell.Value = values[i-1].Username
			cell = row.AddCell()
			cell.Value = values[i-1].Groupname
			//}
		}
	}
	saveErr := file.Save(filePath)
	if saveErr != nil {
		logger.Log.Print(saveErr)
		//fmt.Printf(err.Error())
		return "", "", false, err, "Something Went Wrong"
	}
	// props, err := utility.ReadPropertiesFile(contextPath + "/ifix/resource/application.properties")
	// originalFileName, newFileName, err := utility.FileUploadAPICall(tz.Clientid, tz.Mstorgnhirarchyid, props["fileUploadUrl"], filePath)
	// if err != nil {
	// 	logger.Log.Println("Error while downloading", "-", err)
	// }

	buf, _ := ioutil.ReadFile(filePath)
	typee := http.DetectContentType(buf)
	// OriginalFileName, UploadedFileName, err := utility.FileUploadAPICall(1, 1, config.FileUploadUrl, filePath)
	// if err != nil {
	//  logger.Log.Println("Error while downloading", "-", err)
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

	return OriginalFileName, UploadedFileName, true, nil, ""
}
