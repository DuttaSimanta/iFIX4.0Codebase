package models

import (
	"database/sql"
	"errors"
	"fmt"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/utility"
	"log"
	"os"
	"strconv"
	"strings"

	Excel "github.com/tealeg/xlsx"
)

func getContextPath() (string, error) {
	logger.Log.Println("\n Inside Get Context Path.......")
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.New("ERROR: Unable to get WD")
	}
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	return contextPath, nil
}
func excelTemplateUserAndGroupCheck(db *sql.DB, excelFile *Excel.File, clientID int64, mstOrgnHirarchyId int64, userIds []int64, userNames []string, groupIds []int64, groupNames []string, recordDiffId int64) error {
	//var headerParentId int64 = 1
	fmt.Println("recorddiffid", recordDiffId)
	headerName, headerTemplateErr := dao.GetTemplateHeaderNamesForValidation(db, clientID, mstOrgnHirarchyId, recordDiffId)
	logger.Log.Println("\n headerName ===   ", headerName)
	if headerTemplateErr != nil {
		logger.Log.Println(headerTemplateErr)

		return errors.New("ERROR: Unable to Get Template Details")
	}
	// var headerName []string
	// headerName = append(headerName, "location")
	// headerName = append(headerName, "priority")
	for _, sheet := range excelFile.Sheets[:1] {
		log.Printf("Sheet Name: %s\n", sheet.Name)
		if !strings.EqualFold(sheet.Name, "Sheet1") {
			return errors.New("Sheet Name not matched")
		}
		var rowCount int64 = 0
		for _, row := range sheet.Rows {
			var coloumnCount int64 = 0
			var coloumn []string
			for _, cell := range row.Cells {
				fmt.Println("headerlength", len(row.Cells), len(headerName))
				if rowCount == 0 {
					if len(row.Cells) != len(headerName) {
						logger.Log.Println("Header Length Not Matched")
						return errors.New("ERROR: Header Length Not Matched")
					}
					text := cell.String()
					log.Printf("%s\n", text)
					logger.Log.Println(headerName[coloumnCount])
					if !strings.EqualFold(text, headerName[coloumnCount]) {
						logger.Log.Println("Header Template not matched")
						return errors.New("ERROR: Header Template not matched")
					}
				} else {
					text := cell.String()
					log.Printf("%s\n", text)
					text = strings.Trim(text, " ")
					coloumn = append(coloumn, text)
					log.Printf("%s\n", coloumn[coloumnCount])
				}
				coloumnCount++
			}
			var userId int64 = 0
			logger.Log.Println("\n Test =====   ", coloumnCount)
			if rowCount > 0 {
				//checking for userIds match
				// logger.Log.Println("\n userIds ====   ", userIds)
				for j := 0; j < len(userIds); j++ {
					logger.Log.Println("\n coloumn[coloumnCount-2], userNames[j] ===========    ", coloumn[coloumnCount-2], userNames[j])
					if strings.EqualFold(coloumn[coloumnCount-2], userNames[j]) {
						userId = userIds[j]
						logger.Log.Printf("userIds Id===>%d   Name===>%s", userIds[j], userNames[j])
						break
					}
				}
				if userId == 0 {
					logger.Log.Println("Excel User Login ID Not Matched with Database User")
					return errors.New("Excel User Login ID Not Matched with Database User at line No = " + strconv.FormatInt(rowCount, 10))
				}
			}
			var groupId int64 = 0
			logger.Log.Println("Test")
			if rowCount > 0 {
				//checking for groupIds match
				for j := 0; j < len(groupIds); j++ {
					if strings.EqualFold(coloumn[coloumnCount-1], groupNames[j]) {
						groupId = groupIds[j]
						logger.Log.Printf("groupIds Id===>%d   Name===>%s", groupIds[j], groupNames[j])
						break
					}
				}
				if groupId == 0 {
					logger.Log.Println("Excel Support Group Name Not Matched with Database Support Group")
					return errors.New("Excel Support Group Name Not Matched with Database Support Group at line No = " + strconv.FormatInt(rowCount, 10))
				}
			}
			logger.Log.Println("rowCnts==>", rowCount)
			rowCount++
		}
	}
	return nil
}

//data map[string]interface{}
//  func LocationPriorityUpload(clientID int64, mstOrgnHirarchyId int64, recordDiffTypeId int64, recordDiffId int64, originalFileName string, uploadedFileName string) error {

func BulkUserWithGroupAndCategoryUpload(tz *entities.UserWithGroupAndCategoryForBulkUploadEntity) error {
	//logger.Log.Println(url)
	logger.Log.Println("In BulkCategoryUpload Service")

	clientID := tz.Clientid
	mstOrgnHirarchyId := tz.Mstorgnhirarchyid
	originalFileName := tz.OriginalFileName
	uploadedFileName := tz.UploadedFileName
	recordDiffTypeId := tz.Recorddifftypeid
	recordDiffId := tz.Recorddiffid
	// workingCategories := tz.WorkingCategories
	categoryId := tz.Categoryid
	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		return contextPatherr
	}
	logger.Log.Println("\n contextPath =====    ", contextPath)
	filePath := contextPath + "/ifix/resource/downloads/" + originalFileName
	fileDownloadErr := utility.DownloadFileFromUrl(clientID, mstOrgnHirarchyId, originalFileName, uploadedFileName, filePath)
	//fileDownloadErr := FileUtils.DownloadFileFromUrl(url, filePath)
	if fileDownloadErr != nil {
		fmt.Println("Dowlloaderror")
		logger.Log.Println(fileDownloadErr)
		return fileDownloadErr
	}
	logger.Log.Println("<==================File DownLoaded Successfully=================>...Path==>", filePath)
	excelFile, excelFileOpenErr := Excel.OpenFile(filePath)
	fmt.Println(excelFile)
	if excelFileOpenErr != nil {
		logger.Log.Println(excelFileOpenErr)

		return errors.New("ERROR: Unable to Open Excel File")
	} else {

		lock.Lock()
		defer lock.Unlock()
		db, err := config.ConnectMySqlDbSingleton()
		if err != nil {
			logger.Log.Println("database connection failure", err)
			return errors.New("ERROR: Unable to connect DB")
		}
		tx, txErr := db.Begin()
		if txErr != nil {
			logger.Log.Println(txErr)
			return txErr
		}

		userIds, userLoginNames, userErr := dao.BulkUserDetails(db, clientID, mstOrgnHirarchyId)
		if userErr != nil {
			logger.Log.Println(userErr)
			return userErr
		}

		groupIds, groupNames, groupErr := dao.BulkGroupDetails(db, clientID, mstOrgnHirarchyId)
		if groupErr != nil {
			logger.Log.Println(groupErr)
			return groupErr
		}

		logger.Log.Println("userIds===>", userIds)
		logger.Log.Println("userLoginNames===>", userLoginNames)
		templateValueCheckError := excelTemplateUserAndGroupCheck(db, excelFile, clientID, mstOrgnHirarchyId, userIds, userLoginNames, groupIds, groupNames, recordDiffId)
		if templateValueCheckError != nil {
			logger.Log.Println(templateValueCheckError)
			return templateValueCheckError
		} else {
			for _, sheet := range excelFile.Sheets[:1] {

				var rowCount int64 = 0

				// for catID := 0; catID < len(tz.WorkingCategories); catID++ {

				for _, row := range sheet.Rows[1:] {
					var coloumnCount int64 = 0
					var coloumn []string

					for _, cell := range row.Cells {
						text := cell.String()
						//log.Printf("%s\n", text)
						text = strings.Trim(text, " ")
						coloumn = append(coloumn, text)
						coloumnCount++
					}
					logger.Log.Println("Value of ColoumnClount===>", coloumnCount)
					var userid int64
					for i := 0; i < len(userIds); i++ {
						if userLoginNames[i] == coloumn[coloumnCount-2] {
							userid = userIds[i]
						}
					}
					if userid == 0 {
						tx.Rollback()
						return errors.New("ERROR:NO MATCH WITH USER")
					}
					var groupid int64
					for i := 0; i < len(groupIds); i++ {
						if groupNames[i] == coloumn[coloumnCount-1] {
							groupid = groupIds[i]
						}
					}
					if groupid == 0 {
						tx.Rollback()
						return errors.New("ERROR:NO MATCH WITH SUPPORT GROUP")
					}
					values := entities.UserWithGroupAndCategoryForBulkUploadEntity{}
					values.Clientid = clientID
					values.Mstorgnhirarchyid = mstOrgnHirarchyId
					values.Recorddifftypeid = recordDiffTypeId
					values.Recorddiffid = recordDiffId
					// values.WorkingCategories = workingCategories
					// values.Categoryid = tz.WorkingCategories[catID]
					values.Categoryid = categoryId
					values.Userid = userid
					values.Groupid = groupid
					// tx, txErr := db.Begin()
					// if txErr != nil {
					// 	logger.Log.Println(txErr)
					// 	return txErr
					// }
					count, err := dao.CheckDuplicateBulkUser(&values, db)
					if err != nil {
						tx.Rollback()
						return err
					}
					if count.Total == 0 {
						logger.Log.Println("\n INSERT   .>>>>>>>>>>>>>   ")
						_, insertMstDiffAndMstRecordErr := dao.AddTXUserWithGroupAndCategory(db, tx, &values)
						if insertMstDiffAndMstRecordErr != nil {
							logger.Log.Println(insertMstDiffAndMstRecordErr)
							tx.Rollback()
							return insertMstDiffAndMstRecordErr
						}
					} else {
						logger.Log.Println("\n UPDATE ====>>>>>>>>>>    ")
						updatedID, err := dao.GetUpdatedID(&values, db)
						if err != nil {
							tx.Rollback()
							return err
						}
						_, updateMstDiffAndMstRecordErr := dao.UpdateTXUserWithGroupAndCategory(db, tx, &values, updatedID)
						if updateMstDiffAndMstRecordErr != nil {
							logger.Log.Println(updateMstDiffAndMstRecordErr)
							tx.Rollback()
							return updateMstDiffAndMstRecordErr
						}
					}
				}
				rowCount++

				// }
			}
		}
		tx.Commit()
	}
	//}

	return nil
}
