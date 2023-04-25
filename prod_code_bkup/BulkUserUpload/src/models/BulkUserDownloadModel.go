package models

import (
	// "database/sql"
	"errors"
	"fmt"

	// "log"
	// "os"
	"src/config"
	"src/dao"

	FileUtils "src/fileutils"
	Logger "src/logger"

	// "strings"

	Excel "github.com/tealeg/xlsx"
)

func BulkUserDownload(clientID int64, mstOrgnHirarchyId int64, groupid []int64) (string, string, error) {
	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		Logger.Log.Println(contextPatherr)
		return "", "", contextPatherr
	}
	Logger.Log.Println("BulkUserDownload M1")
	db, dBerr := config.GetDB()
	Logger.Log.Println("BulkUserDownload M1")
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		fmt.Println(dBerr)
		return "", "", errors.New("ERROR: Unable to connect DB")
	}
	OrgName, OrgNameErr := dao.GetOrgName(db, clientID, mstOrgnHirarchyId)
	if OrgNameErr != nil {
		fmt.Println(OrgNameErr)
		Logger.Log.Println(OrgNameErr)
		return "", "", errors.New("ERROR: Dao error")
	}
	filePath := contextPath + "/resource/downloads/" + OrgName + "_" + "USER.xlsx"
	fmt.Println(clientID, mstOrgnHirarchyId)

	headerName, headerErr := dao.GetHeaderName(db, clientID, mstOrgnHirarchyId)
	if headerErr != nil {
		Logger.Log.Println(headerErr)
		return "", "", errors.New("ERROR: Dao error")
	}

	if len(headerName) == 0 {
		return "", "", errors.New("No Headers available for this organization")
	}
	firstnames, lastnames, loginnames, useremails, pmobilenos, smobilenos, divisions, brands, designations, citys, branchs, vipusers, usertypes, userErr := dao.GetUserDetails(db, clientID, mstOrgnHirarchyId, groupid)
	if userErr != nil {
		Logger.Log.Println(userErr)
		return "", "", errors.New("ERROR: Dao error")
	}

	if len(firstnames) == 0 {
		return "", "", errors.New("No Valuse available")
	}

	file := Excel.NewFile()
	sheet, sheetErr := file.AddSheet("User Master")
	if sheetErr != nil {
		Logger.Log.Print(sheetErr)
		return "", "", errors.New("ERROR: sheet adding error")
	}

	for i := 0; i <= len(firstnames); i++ {
		Logger.Log.Println("Rowcount ====>", i)
		row := sheet.AddRow()
		if i == 0 {
			for j := 0; j < len(headerName); j++ {
				cell := row.AddCell()
				cell.Value = headerName[j]
			}
		} else {
			Logger.Log.Println("First names", firstnames[i-1])
			for k := 0; k < len(headerName); k++ {
				if k == 0 {
					cell := row.AddCell()
					cell.Value = firstnames[i-1]
				} else if k == 1 {
					cell := row.AddCell()
					cell.Value = lastnames[i-1]
				} else if k == 2 {
					cell := row.AddCell()
					cell.Value = loginnames[i-1]
				} else if k == 3 {
					cell := row.AddCell()
					cell.Value = useremails[i-1]
				} else if k == 4 {
					cell := row.AddCell()
					cell.Value = pmobilenos[i-1]
				} else if k == 5 {
					cell := row.AddCell()
					cell.Value = smobilenos[i-1]
				} else if k == 6 {
					cell := row.AddCell()
					cell.Value = divisions[i-1]
				} else if k == 7 {
					cell := row.AddCell()
					cell.Value = brands[i-1]
				} else if k == 8 {
					cell := row.AddCell()
					cell.Value = designations[i-1]
				} else if k == 9 {
					cell := row.AddCell()
					cell.Value = citys[i-1]
				} else if k == 10 {
					cell := row.AddCell()
					cell.Value = branchs[i-1]
				} else if k == 11 {
					cell := row.AddCell()
					cell.Value = vipusers[i-1]
				} else if k == 12 {
					cell := row.AddCell()
					cell.Value = usertypes[i-1]
				}
			}
		}

	}

	saveErr := file.Save(filePath)
	if saveErr != nil {
		Logger.Log.Print(saveErr)

		//fmt.Printf(err.Error())
		return "", "", errors.New("ERROR: File saving error")
	}
	props, err := FileUtils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	originalFileName, newFileName, err := FileUtils.FileUploadAPICall(clientID, mstOrgnHirarchyId, props["fileUploadUrl"], filePath)
	if err != nil {
		Logger.Log.Println("Error while downloading", "-", err)
	}

	return originalFileName, newFileName, nil
}
