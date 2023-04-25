//SearchUser for implements business logic
package models

import (
	"errors"
	"fmt"
	"src/config"
	"src/dao"
	"src/fileutils"
	"src/logger"

	"github.com/tealeg/xlsx"
)

func BulkLocationWisePriorityDownload(clientID int64, mstOrgnHirarchyId int64, recordDiffID int64) (string, string, error) {
	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		return "", "", contextPatherr
	}

	db, dBerr := config.GetDB()

	if dBerr != nil {
		logger.Log.Println(dBerr)
		fmt.Println(dBerr)
		return "", "", errors.New("ERROR: Unable to connect DB")
	}
	OrgName, ticketTypeName, OrgNameErr := dao.GetOrgName(db, clientID, mstOrgnHirarchyId, recordDiffID)
	if OrgNameErr != nil {
		fmt.Println(OrgNameErr)
		logger.Log.Println(OrgNameErr)
		return "", "", errors.New("ERROR: dao error")
	}
	filePath := contextPath + "/resource/categoryexcelsheet/" + OrgName + "_" + ticketTypeName + "_" + "CTIS.xlsx"
	fmt.Println(clientID, mstOrgnHirarchyId)
	//defer db.Close()
	headerNames, headerErr := dao.GetTemplateHeaderNamesForValidation(db, clientID, mstOrgnHirarchyId, recordDiffID)
	if headerErr != nil {
		fmt.Println(headerErr)
		logger.Log.Println(headerErr)
		return "", "", errors.New("ERROR: dao error")
	}
	//fmt.Println("Lastrocordidis :", lasRecorddifftypeid)
	values, parentCategoryerr := dao.GetLocatioWisePriorityDetails(db, clientID, mstOrgnHirarchyId, recordDiffID)
	if parentCategoryerr != nil {
		logger.Log.Println(parentCategoryerr)
		return "", "", parentCategoryerr
	}
	headerLength := len(headerNames)
	if headerLength == 0 {
		return "", "", errors.New("ERROR: Header Length Is Zero")
	}
	file := xlsx.NewFile()
	sheet, sheetErr := file.AddSheet("Sheet1")
	if sheetErr != nil {
		logger.Log.Print(sheetErr)

		//fmt.Printf(err.Error())
		return "", "", errors.New("ERROR: sheet adding error")
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
			cell.Value = values[i-1].Location
			cell = row.AddCell()
			cell.Value = values[i-1].ToReccorddiffName
			//}
		}
	}
	saveErr := file.Save(filePath)
	if saveErr != nil {
		logger.Log.Print(saveErr)
		//fmt.Printf(err.Error())
		return "", "", errors.New("ERROR: File saving error")
	}
	props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	originalFileName, newFileName, err := fileutils.FileUploadAPICall(clientID, mstOrgnHirarchyId, props["fileUploadUrl"], filePath)
	if err != nil {
		logger.Log.Println("Error while downloading", "-", err)
	}
	return originalFileName, newFileName, nil
}
