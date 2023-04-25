//SearchUser for implements business logic
package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"src/config"
	"src/dao"
	"src/entities"
	"src/fileutils"
	"src/logger"
	"strconv"
	"strings"

	Excel "github.com/tealeg/xlsx"
)

func getContextPath() (string, error) {

	wd, err := os.Getwd()
	if err != nil {
		return "", errors.New("ERROR: Unable to get WD")
	}
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	return contextPath, nil
}
func excelTemplateAndPriorityCheck(db *sql.DB, excelFile *Excel.File, clientID int64, mstOrgnHirarchyId int64, priorityIds []int64, priorityNames []string, recordDiffId int64) error {
	//var headerParentId int64 = 1
	fmt.Println("recorddiffid", recordDiffId)
	headerName, headerTemplateErr := dao.GetTemplateHeaderNamesForValidation(db, clientID, mstOrgnHirarchyId, recordDiffId)
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
			logger.Log.Printf("coloumnCount=> %d, Coloumns===> %s", coloumnCount, coloumn)
			logger.Log.Println("PriorityIds--->", priorityIds)
			var priorityId int64 = 0
			logger.Log.Println("Test")
			if rowCount > 1 {
				//checking for Priority match
				for j := 0; j < len(priorityIds); j++ {
					if strings.EqualFold(coloumn[coloumnCount-1], priorityNames[j]) {
						priorityId = priorityIds[j]
						logger.Log.Printf("Priority Id===>%d   Name===>%s", priorityIds[j], priorityNames[j])
						break
					}
				}
				if priorityId == 0 {
					logger.Log.Println("Excel Priority Not Matched with Database Priority")
					return errors.New("Excel Priority Not Matched with Database Priority at line No = " + strconv.FormatInt(rowCount, 10))
				}
				if coloumn[coloumnCount-2] == "" {
					logger.Log.Println("Excel Priority is Empty")
					return errors.New("Excel Priority Empty In Line No = " + strconv.FormatInt(rowCount, 10))
				}
				// if !strings.Contains(coloumn[coloumnCount-3], ":") {
				// 	return errors.New("Not a ProperFormat for " + headerName[coloumnCount-2] + " at line No = " + strconv.FormatInt(rowCount, 10))
				// }
				// if !strings.Contains(coloumn[coloumnCount-2], "%") {
				// 	return errors.New("Not a ProperFormat for " + headerName[coloumnCount-1] + " at line No = " + strconv.FormatInt(rowCount, 10))
				// }
			}
			logger.Log.Println("rowCnts==>", rowCount)
			rowCount++
		}
	}
	return nil
}

//data map[string]interface{}
func LocationPriorityUpload(clientID int64, mstOrgnHirarchyId int64, recordDiffTypeId int64, recordDiffId int64, originalFileName string, uploadedFileName string) error {
	//logger.Log.Println(url)
	logger.Log.Println("In BulkCategoryUpload Service")

	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		return contextPatherr
	}
	filePath := contextPath + "/resource/downloads/" + originalFileName
	fileDownloadErr := fileutils.DownloadFileFromUrl(clientID, mstOrgnHirarchyId, originalFileName, uploadedFileName, filePath)
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

		db, dBerr := config.GetDB()
		if dBerr != nil {
			logger.Log.Println(dBerr)
			return errors.New("ERROR: Unable to connect DB")
		}
		tx, txErr := db.Begin()
		if txErr != nil {
			logger.Log.Println(txErr)
			return txErr
		}

		torecorddiffnames, torecorddiffids, priorityErr := dao.Prioritydetails(db, clientID, mstOrgnHirarchyId)
		if priorityErr != nil {
			logger.Log.Println(priorityErr)
			return priorityErr
		}

		logger.Log.Println("priorityNames===>", torecorddiffnames)
		logger.Log.Println("priorityIds===>", torecorddiffids)
		templateValueCheckError := excelTemplateAndPriorityCheck(db, excelFile, clientID, mstOrgnHirarchyId, torecorddiffids, torecorddiffnames, recordDiffId)
		if templateValueCheckError != nil {
			logger.Log.Println(templateValueCheckError)
			return templateValueCheckError
		} else {
			for _, sheet := range excelFile.Sheets[:1] {

				var rowCount int64 = 0

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
					var priorityid int64
					for i := 0; i < len(torecorddiffids); i++ {
						if torecorddiffnames[i] == coloumn[coloumnCount-1] {
							priorityid = torecorddiffids[i]
						}
					}
					if priorityid == 0 {
						return errors.New("ERROR:NO MATCH WITH PRIORITY")
					}
					values := entities.LocationPriorityEntity{}
					values.ClientID = clientID
					values.MstorgnhirarchyID = mstOrgnHirarchyId
					values.Recorddifftypeid = recordDiffTypeId
					values.Location = coloumn[coloumnCount-2]
					values.Recorddiffid = recordDiffId
					values.ToRecorddifftypeid = 5
					values.ToRecorddiffid = priorityid
					// tx, txErr := db.Begin()
					// if txErr != nil {
					// 	logger.Log.Println(txErr)
					// 	return txErr
					// }
					count, err := dao.CheckDuplicateLocation(&values, db)
					if err != nil {
						tx.Rollback()
						return err
					}
					if count.Total == 0 {
						_, insertMstDiffAndMstRecordErr := dao.AddTXLocation(db, tx, &values)
						if insertMstDiffAndMstRecordErr != nil {
							logger.Log.Println(insertMstDiffAndMstRecordErr)
							tx.Rollback()
							return insertMstDiffAndMstRecordErr
						}
					}
				}
				rowCount++
			}
		}
		tx.Commit()
	}
	//}

	return nil
}

// func BulkLocationWisePriorityDownload(clientID int64, mstOrgnHirarchyId int64, recordDiffID int64) (string, string, error) {
// 	contextPath, contextPatherr := getContextPath()
// 	if contextPatherr != nil {
// 		logger.Log.Println(contextPatherr)
// 		return "", "", contextPatherr
// 	}

// 	db, dBerr := config.GetDB()

// 	if dBerr != nil {
// 		logger.Log.Println(dBerr)
// 		fmt.Println(dBerr)
// 		return "", "", errors.New("ERROR: Unable to connect DB")
// 	}
// 	OrgName, ticketTypeName, OrgNameErr := dao.GetOrgName(db, clientID, mstOrgnHirarchyId, recordDiffID)
// 	if OrgNameErr != nil {
// 		fmt.Println(OrgNameErr)
// 		logger.Log.Println(OrgNameErr)
// 		return "", "", errors.New("ERROR: dao error")
// 	}
// 	filePath := contextPath + "/resource/categoryexcelsheet/" + OrgName + "_" + ticketTypeName + "_" + "CTIS.xlsx"
// 	fmt.Println(clientID, mstOrgnHirarchyId)
// 	//defer db.Close()
// 	headerNames, headerErr := dao.GetTemplateHeaderNamesForValidation(db, clientID, mstOrgnHirarchyId, recordDiffID)
// 	if headerErr != nil {
// 		fmt.Println(headerErr)
// 		logger.Log.Println(headerErr)
// 		return "", "", errors.New("ERROR: dao error")
// 	}

// 	//fmt.Println("Lastrocordidis :", lasRecorddifftypeid)
// 	values, parentCategoryerr := dao.GetLocatioWisePriorityDetails(db, clientID, mstOrgnHirarchyId, recordDiffID)
// 	if parentCategoryerr != nil {
// 		logger.Log.Println(parentCategoryerr)
// 		return "", "", parentCategoryerr
// 	}

// 	headerLength := len(headerNames)

// 	file := xlsx.NewFile()
// 	sheet, sheetErr := file.AddSheet("Category Master")
// 	if sheetErr != nil {
// 		logger.Log.Print(sheetErr)

// 		//fmt.Printf(err.Error())
// 		return "", "", errors.New("ERROR: sheet adding error")
// 	}
// 	for i := 0; i <= len(values); i++ {
// 		logger.Log.Println("ROwCOunt---->", i)
// 		row := sheet.AddRow()
// 		if i == 0 {
// 			for j := 0; j < headerLength; j++ {
// 				cell := row.AddCell()
// 				cell.Value = headerNames[j]
// 			}
// 		} else {
// 			// logger.Log.Println("ParentCategorynames====>", parentCategoryNames[i-1])

// 			// splittedParentCatagories := strings.Split(parentCategoryNames[i-1], "->") //(i-1) because for i=0 headernames is added
// 			// logger.Log.Println("cat level len====>", headerLength-6)
// 			// logger.Log.Println("Splitted Length====>", len(splittedParentCatagories))
// 			//for j := 0; j < headerLength; j++ {
// 			cell := row.AddCell()
// 			cell.Value = values[i-1].Location
// 			cell = row.AddCell()
// 			cell.Value = values[i-1].ToReccorddiffName
// 			//}
// 		}
// 	}
// 	saveErr := file.Save(filePath)
// 	if saveErr != nil {
// 		logger.Log.Print(saveErr)
// 		//fmt.Printf(err.Error())
// 		return "", "", errors.New("ERROR: File saving error")
// 	}
// 	props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")
// 	originalFileName, newFileName, err := fileutils.FileUploadAPICall(clientID, mstOrgnHirarchyId, props["fileUploadUrl"], filePath)
// 	if err != nil {
// 		logger.Log.Println("Error while downloading", "-", err)
// 	}
// 	return originalFileName, newFileName, nil
// }
