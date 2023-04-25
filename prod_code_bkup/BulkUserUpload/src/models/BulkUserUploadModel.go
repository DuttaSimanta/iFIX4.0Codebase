package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"src/config"
	"src/dao"

	FileUtils "src/fileutils"
	Logger "src/logger"
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

func ExcelTemplateAndValueCheck(db *sql.DB, excelFile *Excel.File, clientID int64, orgID int64) error {
	sheetName := "User Master"
	for _, sheet := range excelFile.Sheets {

		if !strings.EqualFold(sheetName, sheet.Name) {
			Logger.Log.Println("Expected Sheet Name", sheetName)
			Logger.Log.Println("Excel Sheet Name", sheet.Name)
			return errors.New("ERROR: Invalid Sheet Name")
		} else {
			Logger.Log.Println("Expected Sheet Name", sheetName)
			Logger.Log.Println("Excel Sheet Name", sheet.Name)
			headerName, headerError := dao.GetHeaderName(db, clientID, orgID)
			if headerError != nil {
				Logger.Log.Println(headerError)
				return headerError
			}
			//log.Println("HeaderIDS=>", headerIds)
			var rowCount int64 = 0
			for _, row := range sheet.Rows {
				var coloumnCount int64 = 0
				for _, cell := range row.Cells {
					if rowCount == 0 {
						if len(row.Cells) != len(headerName) {
							Logger.Log.Println("Header Length Not Matched", len(row.Cells))
							return errors.New("Header Length Not Matched")
						}
						text := cell.String()
						text = strings.Trim(text, " ")
						Logger.Log.Printf("%s\n", text)
						Logger.Log.Println(headerName[coloumnCount])
						if !strings.EqualFold(text, headerName[coloumnCount]) {
							log.Println("Header Template not matched")
							return errors.New("Header Template not matched")
						}

					}
					coloumnCount++
				}

				rowCount++
			}
		}
	}
	return nil
}

func BulkUserUploadUsingExcel(requestData map[string]interface{}) error {
	Logger.Log.Println("In BulkUserUploadUsingExcel")
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	originalFileName := requestData["originalfilename"].(string)
	uploadedFileName := requestData["uploadedfilename"].(string)
	grpID := int64(requestData["groupid"].(float64))
	roleID := int64(requestData["roleid"].(float64))
	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		Logger.Log.Println(contextPatherr)
		return contextPatherr
	}
	filePath := contextPath + "/resource/downloads/" + originalFileName

	fileDownloadErr := FileUtils.DownloadFileFromUrl(clientID, orgID, originalFileName, uploadedFileName, filePath)
	if fileDownloadErr != nil {
		Logger.Log.Println(fileDownloadErr)
		return fileDownloadErr
	}
	log.Println("File DownLoaded Successfull...Path==>", filePath)
	excelFile, excelFileOpenErr := Excel.OpenFile(filePath)

	if excelFileOpenErr != nil {
		Logger.Log.Println(excelFileOpenErr)

		return errors.New("ERROR: Unable to Open Excel File")
	} else {
		db, dBerr := config.GetDB()
		if dBerr != nil {
			Logger.Log.Println(dBerr)
			return errors.New("ERROR: Unable to connect DB")
		}

		templateValueCheckError := ExcelTemplateAndValueCheck(db, excelFile, clientID, orgID)
		if templateValueCheckError != nil {
			Logger.Log.Println(templateValueCheckError)

			return templateValueCheckError
		} else {
			var count int64
			var errorstring string
			for _, sheet := range excelFile.Sheets[:1] {

				var rowCount int64 = 1

				for _, row := range sheet.Rows[1:] {
					var coloumnCount int64 = 0
					var coloumn []string
					//for loop for coloumn reading and storing in slice
					for _, cell := range row.Cells {
						text := cell.String()
						//log.Printf("%s\n", text)
						text = strings.Trim(text, " ")
						coloumn = append(coloumn, text)
						coloumnCount++
					}
					Logger.Log.Println("Value of ColoumnClount===>", coloumnCount)
					//rowCount++
					Logger.Log.Println("Row  NO ==>", rowCount)
					userUploadError := dao.UserUploadWithRoleMapAndGroupMap(db, clientID, orgID, coloumn, grpID, roleID)
					if userUploadError != nil {
						count++
						Logger.Log.Println("ROW where error Found for user Insert==>", rowCount)
						Logger.Log.Println("Error ==>", userUploadError)
						errorstring = errorstring + "Row No.-" + fmt.Sprintf("%v", rowCount) + ". Due to " + fmt.Sprintf("%v", userUploadError) + "\n"
						//return errors.New(errorstring)

					}
					rowCount++

				}

			}
			Logger.Log.Println("Final Errorlist===>", errorstring)
		}

	}
	return nil
}

// func BulkUserDownload(clientID int64, mstOrgnHirarchyId int64, groupid []int64) (string, string, error) {
// 	contextPath, contextPatherr := getContextPath()
// 	if contextPatherr != nil {
// 		Logger.Log.Println(contextPatherr)
// 		return "", "", contextPatherr
// 	}
// 	Logger.Log.Println("BulkUserDownload M1")
// 	db, dBerr := config.GetDB()
// 	Logger.Log.Println("BulkUserDownload M1")
// 	if dBerr != nil {
// 		Logger.Log.Println(dBerr)
// 		fmt.Println(dBerr)
// 		return "", "", errors.New("ERROR: Unable to connect DB")
// 	}
// 	OrgName, OrgNameErr := dao.GetOrgName(db, clientID, mstOrgnHirarchyId)
// 	if OrgNameErr != nil {
// 		fmt.Println(OrgNameErr)
// 		Logger.Log.Println(OrgNameErr)
// 		return "", "", errors.New("ERROR: Dao error")
// 	}
// 	filePath := contextPath + "/resource/downloads/" + OrgName + "_" + "CTIS.xlsx"
// 	fmt.Println(clientID, mstOrgnHirarchyId)

// 	headerName, headerErr := dao.GetHeaderName(db, clientID, mstOrgnHirarchyId)
// 	if headerErr != nil {
// 		Logger.Log.Println(headerErr)
// 		return "", "", errors.New("ERROR: Dao error")
// 	}

// 	if len(headerName) == 0 {
// 		return "", "", errors.New("No Headers available for this organization")
// 	}
// 	firstnames, lastnames, loginnames, useremails, pmobilenos, smobilenos, divisions, brands, designations, citys, branchs, vipusers, usertypes, userErr := dao.GetUserDetails(db, clientID, mstOrgnHirarchyId, groupid)
// 	if userErr != nil {
// 		Logger.Log.Println(userErr)
// 		return "", "", errors.New("ERROR: Dao error")
// 	}

// 	if len(firstnames) == 0 {
// 		return "", "", errors.New("No Valuse available")
// 	}

// 	file := Excel.NewFile()
// 	sheet, sheetErr := file.AddSheet("User Master")
// 	if sheetErr != nil {
// 		Logger.Log.Print(sheetErr)
// 		return "", "", errors.New("ERROR: sheet adding error")
// 	}

// 	for i := 0; i <= len(firstnames); i++ {
// 		Logger.Log.Println("Rowcount ====>", i)
// 		row := sheet.AddRow()
// 		if i == 0 {
// 			for j := 0; j < len(headerName); j++ {
// 				cell := row.AddCell()
// 				cell.Value = headerName[j]
// 			}
// 		} else {
// 			Logger.Log.Println("First names", firstnames[i-1])
// 			for k := 0; k < len(headerName); k++ {
// 				if k == 0 {
// 					cell := row.AddCell()
// 					cell.Value = firstnames[i-1]
// 				} else if k == 1 {
// 					cell := row.AddCell()
// 					cell.Value = lastnames[i-1]
// 				} else if k == 2 {
// 					cell := row.AddCell()
// 					cell.Value = loginnames[i-1]
// 				} else if k == 3 {
// 					cell := row.AddCell()
// 					cell.Value = useremails[i-1]
// 				} else if k == 4 {
// 					cell := row.AddCell()
// 					cell.Value = pmobilenos[i-1]
// 				} else if k == 5 {
// 					cell := row.AddCell()
// 					cell.Value = smobilenos[i-1]
// 				} else if k == 6 {
// 					cell := row.AddCell()
// 					cell.Value = divisions[i-1]
// 				} else if k == 7 {
// 					cell := row.AddCell()
// 					cell.Value = brands[i-1]
// 				} else if k == 8 {
// 					cell := row.AddCell()
// 					cell.Value = designations[i-1]
// 				} else if k == 9 {
// 					cell := row.AddCell()
// 					cell.Value = citys[i-1]
// 				} else if k == 10 {
// 					cell := row.AddCell()
// 					cell.Value = branchs[i-1]
// 				} else if k == 11 {
// 					cell := row.AddCell()
// 					cell.Value = vipusers[i-1]
// 				} else if k == 12 {
// 					cell := row.AddCell()
// 					cell.Value = usertypes[i-1]
// 				}
// 			}
// 		}

// 	}

// 	saveErr := file.Save(filePath)
// 	if saveErr != nil {
// 		Logger.Log.Print(saveErr)

// 		//fmt.Printf(err.Error())
// 		return "", "", errors.New("ERROR: File saving error")
// 	}
// 	props, err := FileUtils.ReadPropertiesFile(contextPath + "/resource/application.properties")
// 	originalFileName, newFileName, err := FileUtils.FileUploadAPICall(clientID, mstOrgnHirarchyId, props["fileUploadUrl"], filePath)
// 	if err != nil {
// 		Logger.Log.Println("Error while downloading", "-", err)
// 	}

// 	return originalFileName, newFileName, nil
// }
