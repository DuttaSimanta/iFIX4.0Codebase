package models

import (
	"errors"
	"fmt"
	"log"
	"src/config"
	Dao "src/dao"
	"src/fileutils"
	Logger "src/logger"
	"strings"

	"github.com/tealeg/xlsx"
)

func BulkCategoryDownload(clientID int64, mstOrgnHirarchyId int64, recordDiffID int64) (string, string, error) {
	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		Logger.Log.Println(contextPatherr)
		return "", "", contextPatherr
	}
	log.Println("BulkCategoryDownload 1")
	db, dBerr := config.GetDB()
	log.Println("BulkCategoryDownload 1")
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		fmt.Println(dBerr)
		return "", "", errors.New("ERROR: Unable to connect DB")
	}
	OrgName, ticketTypeName, OrgNameErr := Dao.GetOrgName(db, clientID, mstOrgnHirarchyId, recordDiffID)
	if OrgNameErr != nil {
		fmt.Println(OrgNameErr)
		Logger.Log.Println(OrgNameErr)
		return "", "", errors.New("ERROR: Dao error")
	}
	filePath := contextPath + "/resource/categoryexcelsheet/" + OrgName + "_" + ticketTypeName + "_" + "CTIS.xlsx"
	fmt.Println(clientID, mstOrgnHirarchyId)
	//defer db.Close()
	headerNames, headerErr := Dao.Getheaderr(db, clientID, mstOrgnHirarchyId, recordDiffID)
	if headerErr != nil {
		fmt.Println(headerErr)
		Logger.Log.Println(headerErr)
		return "", "", errors.New("ERROR: Dao error")
	}

	//fmt.Println("Lastrocordidis :", lasRecorddifftypeid)
	parentCategoryNames, impactNames, urgencyNames, prirityNames, estimatedtimes, efficiencies, changeTypes, parentCategoryerr := Dao.GetParentcatagory(db, clientID, mstOrgnHirarchyId, recordDiffID)
	if parentCategoryerr != nil {
		Logger.Log.Println(parentCategoryerr)
		return "", "", parentCategoryerr
	}
	//impactnames, urgencyNames, prirityNames, impactUrgencyPriorityErr := Dao.GetImpactUrgencyPriorityNames(db, clientID, mstOrgnHirarchyId, lasRecorddifftypeid)
	//fmt.Println("all Catagory are:", len(parentCategoryNames), len(impactNames), impactNames)
	/*if impactUrgencyPriorityErr != nil {
		Logger.Log.Panicln(impactUrgencyPriorityErr)
		return impactUrgencyPriorityErr
	}
	estimatedtimes, efficiencies, estimatedtimeEfficiencyErr := Dao.GetEstimatedtimesEfficiencies(db, clientID, mstOrgnHirarchyId, lasRecorddifftypeid)
	//fmt.Println("all Catagory are:", parentcategorynames)
	if estimatedtimeEfficiencyErr != nil {
		Logger.Log.Panicln(estimatedtimeEfficiencyErr)
		return estimatedtimeEfficiencyErr
	}
	parentcategorynameslength := len(parentcategorynames)*/
	headerLength := len(headerNames)
	//splittedParentCatagories := make([]string, s)
	/*for i := 0; i < len(parentcategorynames); i++ {
		splittedParentCatagories[i] = strings.Split(parentcategorynames[i], "->")
	}*/
	file := xlsx.NewFile()
	sheet, sheetErr := file.AddSheet("Category Master")
	if sheetErr != nil {
		Logger.Log.Print(sheetErr)

		//fmt.Printf(err.Error())
		return "", "", errors.New("ERROR: sheet adding error")
	}
	//var elem string
	//fmt.Println("length is:", parentcategorynameslength)
	//fmt.Println(Parentcategorynames)
	for i := 0; i <= len(parentCategoryNames); i++ {
		Logger.Log.Println("ROwCOunt---->", i)
		row := sheet.AddRow()
		if i == 0 {
			for j := 0; j < headerLength; j++ {
				cell := row.AddCell()
				cell.Value = headerNames[j]
			}
		} else {
			Logger.Log.Println("ParentCategorynames====>", parentCategoryNames[i-1])

			splittedParentCatagories := strings.Split(parentCategoryNames[i-1], "->") //(i-1) because for i=0 headernames is added
			Logger.Log.Println("cat level len====>", headerLength-6)
			Logger.Log.Println("Splitted Length====>", len(splittedParentCatagories))
			for j := 0; j < headerLength; j++ {
				if len(splittedParentCatagories) == headerLength-6 {
					if j < headerLength-6 { //Adding all splitted catagory by "->"
						cell := row.AddCell()
						cell.Value = splittedParentCatagories[j]
					} else if j == headerLength-6 { //Adding impact Names
						cell := row.AddCell()
						cell.Value = impactNames[i-1]
					} else if j == headerLength-5 { //Adding urgency names
						cell := row.AddCell()
						cell.Value = urgencyNames[i-1]
					} else if j == headerLength-4 { //Adding priority
						cell := row.AddCell()
						cell.Value = prirityNames[i-1]
					} else if j == headerLength-3 { //Adding EstimatedTimes
						cell := row.AddCell()
						cell.Value = estimatedtimes[i-1]
					} else if j == headerLength-2 { //Adding Efficiencies
						cell := row.AddCell()
						cell.Value = efficiencies[i-1]
					} else if j == headerLength-1 { //Adding changetypes
						cell := row.AddCell()
						cell.Value = changeTypes[i-1]
					}
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
	props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	originalFileName, newFileName, err := fileutils.FileUploadAPICall(clientID, mstOrgnHirarchyId, props["fileUploadUrl"], filePath)
	if err != nil {
		Logger.Log.Println("Error while downloading", "-", err)
	}
	return originalFileName, newFileName, nil
}
