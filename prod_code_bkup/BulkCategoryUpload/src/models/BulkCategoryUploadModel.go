package models

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"src/config"
	Dao "src/dao"
	model "src/entities"
	FileUtils "src/fileutils"
	Logger "src/logger"
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
func excelTemplateAndValueCheck(db *sql.DB, excelFile *Excel.File, clientID int64, mstOrgnHirarchyId int64, priorityIds []int64, priorityNames []string, impactIds []int64, impactNames []string, urgencyIds []int64, urgencyNames []string, recordDiffId int64) error {
	//var headerParentId int64 = 1

	headerName, headerTemplateErr := Dao.GetTemplateHeaderNamesForValidation(db, clientID, mstOrgnHirarchyId, recordDiffId)
	if headerTemplateErr != nil {
		Logger.Log.Println(headerTemplateErr)

		return errors.New("ERROR: Unable to Get Template Details")
	}

	for _, sheet := range excelFile.Sheets[:1] {
		log.Printf("Sheet Name: %s\n", sheet.Name)
		if !strings.EqualFold(sheet.Name, "Category Master") {
			return errors.New("Sheet Name not matched")
		}
		var rowCount int64 = 0
		for _, row := range sheet.Rows {
			var coloumnCount int64 = 0
			var coloumn []string
			for _, cell := range row.Cells {
				if rowCount == 0 {
					if len(row.Cells) != len(headerName) {
						Logger.Log.Println("Header Length Not Matched")
						return errors.New("ERROR: Header Length Not Matched")
					}

					text := cell.String()
					log.Printf("%s\n", text)
					Logger.Log.Println(headerName[coloumnCount])
					if !strings.EqualFold(text, headerName[coloumnCount]) {
						Logger.Log.Println("Header Template not matched")

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
			Logger.Log.Printf("coloumnCount=> %d, Coloumns===> %s", coloumnCount, coloumn)
			Logger.Log.Println("impact--->", impactIds)
			Logger.Log.Println("urgency--->", urgencyIds)
			Logger.Log.Println("PriorityIds--->", priorityIds)
			if rowCount >= 1 {
				var i int64 = 0
				for i < (coloumnCount - 6) {
					if i == (coloumnCount - 6) {
						var impactId int64 = 0
						var urgencyId int64 = 0
						var priorityId int64 = 0
						Logger.Log.Println("Test")
						for j := range impactIds {
							if strings.EqualFold(coloumn[coloumnCount-6], impactNames[j]) {
								impactId = impactIds[j]
								Logger.Log.Printf("Impact Id===>%d   Name===>%s", impactIds[j], impactNames[j])
								break
							}
						}
						//checking for urgency match
						for j := range urgencyIds {
							if strings.EqualFold(coloumn[coloumnCount-5], urgencyNames[j]) {
								urgencyId = urgencyIds[j]
								Logger.Log.Printf("Urgency Id===>%d   Name===>%s", urgencyIds[j], urgencyNames[j])
								break
							}
						}
						//checking for Priority match
						for j := range priorityIds {
							if strings.EqualFold(coloumn[coloumnCount-4], priorityNames[j]) {
								priorityId = priorityIds[j]
								Logger.Log.Printf("Priority Id===>%d   Name===>%s", priorityIds[j], priorityNames[j])
								break
							}
						}
						if impactId == 0 {
							Logger.Log.Println("Excel Impact Not Matched with Database Impact")
							//var errorout string = "Excel Impact Not Matched with Database Impact for line = "+strconv.FormatInt(rowCount, 10)
							return errors.New("Excel Impact Not Matched with Database Impact at line No = " + strconv.FormatInt(rowCount, 10))
						}
						if urgencyId == 0 {
							Logger.Log.Println("Excel Urgency Not Matched with Database Urgency on line No =")
							return errors.New("Excel Urgency Not Matched with Database Urgency at line No = " + strconv.FormatInt(rowCount, 10))
						}
						if priorityId == 0 {
							Logger.Log.Println("Excel Priority Not Matched with Database Priority")
							return errors.New("Excel Priority Not Matched with Database Priority at line No = " + strconv.FormatInt(rowCount, 10))
						}

						if !strings.Contains(coloumn[coloumnCount-3], ":") {
							return errors.New("Not a ProperFormat for " + headerName[coloumnCount-2] + " at line No = " + strconv.FormatInt(rowCount, 10))
						}
						if !strings.Contains(coloumn[coloumnCount-2], "%") {
							return errors.New("Not a ProperFormat for " + headerName[coloumnCount-1] + " at line No = " + strconv.FormatInt(rowCount, 10))
						}
					}
					i++
				}
			}
			Logger.Log.Println("rowCnts==>", rowCount)
			rowCount++
		}
	}

	return nil
}

//data map[string]interface{}
func BulkCategoryUpload(clientID int64, mstOrgnHirarchyId int64, recordDiffTypeId int64, recordDiffId int64, originalFileName string, uploadedFileName string) error {
	//Logger.Log.Println(url)
	Logger.Log.Println("In BulkCategoryUpload Service")
	//Logger.Log.Println("clientid===>",clientID)
	//Logger.Log.Println("MSt===>",mstOrgnHirarchyId)
	//Logger.Log.Println("Filename==>",filename)
	//Logger.Log.Println("url===>",url)
	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		Logger.Log.Println(contextPatherr)
		return contextPatherr
	}
	// originalFileName := "ICCM_SR_Masterdata.xlsx"
	// uploadedFileName := "20210825-70fa5dad-cc68-4e8f-a645-d37bd6cb5196.xlsx"
	filePath := contextPath + "/resource/downloads/" + originalFileName

	fileDownloadErr := FileUtils.DownloadFileFromUrl(clientID, mstOrgnHirarchyId, originalFileName, uploadedFileName, filePath)
	//fileDownloadErr := FileUtils.DownloadFileFromUrl(url, filePath)
	if fileDownloadErr != nil {
		Logger.Log.Println(fileDownloadErr)
		return fileDownloadErr
	}
	Logger.Log.Println("<==================File DownLoaded Successfully=================>...Path==>", filePath)
	excelFile, excelFileOpenErr := Excel.OpenFile(filePath)

	if excelFileOpenErr != nil {
		Logger.Log.Println(excelFileOpenErr)

		return errors.New("ERROR: Unable to Open Excel File")
	} else {
		/* var headerName []string
		var mstDiffTypeId []int64
		var priorityNames []string
		var priorityIds []int64
		var impactNames []string
		var impactIds []int64
		var urgencyNames []string
		var urgencyIds []int64 */
		db, dBerr := config.GetDB()
		if dBerr != nil {
			Logger.Log.Println(dBerr)
			return errors.New("ERROR: Unable to connect DB")
		}
		//defer db.Close()
		categoryLevelNames, categoryLevelIds, categoryLevelNameAndIdserr := Dao.GetCategoryLevelNameAndId(db, clientID, mstOrgnHirarchyId, recordDiffId)
		if categoryLevelNameAndIdserr != nil {
			Logger.Log.Println(categoryLevelNameAndIdserr)
			return categoryLevelNameAndIdserr
		}

		Logger.Log.Println("excelFile==>", excelFile)
		Logger.Log.Println("categoryLevelNames===>", categoryLevelNames)
		Logger.Log.Println("categoryLevelIds===>", categoryLevelIds)
		impactNames, impactIds, urgencyNames, urgencyIds, priorityNames, priorityIds, impactUrgencyPriorityError := Dao.GetImactUrgencyPriorityDetails(db, clientID, mstOrgnHirarchyId)
		if impactUrgencyPriorityError != nil {
			Logger.Log.Println(impactUrgencyPriorityError)
			return impactUrgencyPriorityError
		}
		Logger.Log.Println("impactNames==>", impactNames)
		Logger.Log.Println("impactIds===>", impactIds)
		Logger.Log.Println("urgencyNames===>", urgencyNames)
		Logger.Log.Println("urgencyIds==>", urgencyIds)
		Logger.Log.Println("priorityNames===>", priorityNames)
		Logger.Log.Println("priorityIds===>", priorityIds)
		templateValueCheckError := excelTemplateAndValueCheck(db, excelFile, clientID, mstOrgnHirarchyId, priorityIds, priorityNames, impactIds, impactNames, urgencyIds, urgencyNames, recordDiffId)
		if templateValueCheckError != nil {
			Logger.Log.Println(templateValueCheckError)
			return templateValueCheckError
		} else {
			for _, sheet := range excelFile.Sheets[:1] {

				var rowCount int64 = 0

				for _, row := range sheet.Rows[1:] {
					var coloumnCount int64 = 0
					var coloumn []string
					var recordDiffParentid int64 = 0
					var parentCategoryIds string
					var parentCategoryNames string
					//for loop for coloumn reading and storing in slice
					for _, cell := range row.Cells {
						text := cell.String()
						//log.Printf("%s\n", text)
						text = strings.Trim(text, " ")
						coloumn = append(coloumn, text)
						coloumnCount++
					}
					Logger.Log.Println("Value of ColoumnClount===>", coloumnCount)
					var i int64 = 0
					if strings.Contains(coloumn[coloumnCount-2], "\\") {
						coloumn[coloumnCount-2] = strings.ReplaceAll(coloumn[coloumnCount-2], "\\", "")
						//log.Fatal("Error")
					}
					//for loop for category upload logic
					for i < (coloumnCount - 6) {
						var mstDiff model.MstRecordDifferentiation
						var mstRecord model.MstRecordType
						var categoryExistcount int64 = 0

						categoryExistcount, categoryExistCountErr := Dao.GetCategoryExistCount(db, clientID, mstOrgnHirarchyId, recordDiffParentid, coloumn[i], categoryLevelIds[i])
						if categoryExistCountErr != nil {
							Logger.Log.Println(categoryExistCountErr)
							return categoryExistCountErr
						}
						Logger.Log.Println("categoryExistcount===>", categoryExistcount)

						//if cateegory already exist
						if categoryExistcount >= 1 {
							parentId, getParentIDErr := Dao.GetImmediatePatentsId(db, clientID, mstOrgnHirarchyId, recordDiffParentid, coloumn[i], categoryLevelIds[i])
							if getParentIDErr != nil {
								Logger.Log.Println(getParentIDErr)
								return getParentIDErr
							}
							if i < (coloumnCount - 6) {
								parentCategoryNames = parentCategoryNames + "->" + coloumn[i]
							}
							if i == 0 {
								//parentCategoryIds = strconv.FormatInt(recordDiffParentid, 10)
							} else {
								parentCategoryIds = parentCategoryIds + "->" + strconv.FormatInt(recordDiffParentid, 10)
							}
							recordDiffParentid = parentId
							i++
							continue
						} else {
							if i < (coloumnCount - 6) {
								parentCategoryNames = parentCategoryNames + "->" + coloumn[i]
							}
							if i == 0 {
								//parentCategoryIds = strconv.FormatInt(recordDiffParentid, 10)
							} else {
								parentCategoryIds = parentCategoryIds + "->" + strconv.FormatInt(recordDiffParentid, 10)
							}

							//if new category found
							//mstDiff table object
							mstDiff.ActiveFlag = 1
							mstDiff.ClientId = clientID
							mstDiff.MstOrgnHirarchyId = mstOrgnHirarchyId
							mstDiff.DeleteFlag = 0
							mstDiff.Name = coloumn[i]
							mstDiff.ParentId = recordDiffParentid
							mstDiff.RecordDiffTypeId = categoryLevelIds[i]
							mstDiff.SeqNo = i + 1
							//mstDiff.AuditTransactionId =
							mstDiff.ParentCategoryids = strings.Trim(parentCategoryIds, "->")
							mstDiff.ParentCategoryNames = strings.Trim(parentCategoryNames, "->")
							//mst record table object
							mstRecord.ActiveFlag = 1
							mstRecord.AuditTransactionId = 1
							mstRecord.ClientId = clientID
							mstRecord.MstOrgnHirarchyId = mstOrgnHirarchyId
							mstRecord.FromRecordDiffTypeId = recordDiffTypeId
							mstRecord.FromRecordDiffId = recordDiffId
							mstRecord.ToRecordDiffTypeId = categoryLevelIds[i]

							tx, txErr := db.Begin()
							if txErr != nil {
								Logger.Log.Println(txErr)
								return txErr

							}
							lastInsertedmstDiffId, insertMstDiffAndMstRecordErr := Dao.InsertMstDiffAndMstRecord(db, tx, &mstDiff, &mstRecord)
							if insertMstDiffAndMstRecordErr != nil {

								Logger.Log.Println(insertMstDiffAndMstRecordErr)
								tx.Rollback()
								return insertMstDiffAndMstRecordErr
							}

							if i == (coloumnCount - 7) { //value 6 is for 5 more coloumn after working category.
								var mstBusinessMatrix model.MstBusinessMatrix
								//var priorityid int64
								//checking Impact match
								for j := range impactIds {
									if strings.EqualFold(coloumn[coloumnCount-6], impactNames[j]) {
										mstBusinessMatrix.MstRecordDifferentiationImpactId = impactIds[j]
										log.Printf("Impact Id===>%d   Name===>%s", impactIds[j], impactNames[j])
										break
									}
								}
								//checking for urgency match
								for j := range priorityIds {
									if strings.EqualFold(coloumn[coloumnCount-5], urgencyNames[j]) {
										mstBusinessMatrix.MstRecordDifferentiationUrgencyId = urgencyIds[j]
										log.Printf("Urgency Id===>%d   Name===>%s", urgencyIds[j], urgencyNames[j])
										break
									}
								}
								//checking for Priority match
								for j := range priorityIds {
									if strings.EqualFold(coloumn[coloumnCount-4], priorityNames[j]) {
										mstBusinessMatrix.MstRecordDifferentiationPriorityId = priorityIds[j]
										log.Printf("Priority Id===>%d   Name===>%s", priorityIds[j], priorityNames[j])
										break
									}
								}
								if mstBusinessMatrix.MstRecordDifferentiationImpactId == 0 {
									Logger.Log.Println("Excel Impact Not Matched with Database Impact")
								}
								if mstBusinessMatrix.MstRecordDifferentiationUrgencyId == 0 {
									Logger.Log.Println("Excel Urgency Not Matched with Database Urgency")
								}
								if mstBusinessMatrix.MstRecordDifferentiationPriorityId == 0 {
									Logger.Log.Println("Excel Priority Not Matched with Database Priority")
								}
								Logger.Log.Println("Working Category")
								mstBusinessMatrix.ClientId = clientID
								mstBusinessMatrix.MstOrgnHirarchyId = mstOrgnHirarchyId
								mstBusinessMatrix.MstRecordDifferentiationTicketTypeId = recordDiffId
								mstBusinessMatrix.MstRecordDifferentiationCatId = lastInsertedmstDiffId
								//mstBusinessMatrix.MstRecordDifferentiationImpactId  = 0
								//mstBusinessMatrix.MstRecordDifferentiationUrgencyId  = 0
								mstBusinessMatrix.DeleteFlag = 0
								mstBusinessMatrix.ActiveFlag = 1
								//mstBusinessMatrix.AuditTransactionId =1

								insertmstBusinessMatrixErr := Dao.InsertMstBusinessMatrix(db, tx, &mstBusinessMatrix)
								if insertmstBusinessMatrixErr != nil {
									Logger.Log.Println(insertmstBusinessMatrixErr)
									tx.Rollback()
									return insertmstBusinessMatrixErr
								}

								var mapCategoryWithEstimateTime model.MapCategoryWithEstimateTime
								mapCategoryWithEstimateTime.ClientId = clientID
								mapCategoryWithEstimateTime.MstOrgnHirarchyId = mstOrgnHirarchyId
								mapCategoryWithEstimateTime.RecordDiffId = mstRecord.ToRecordDiffId
								mapCategoryWithEstimateTime.EstimatedTime = coloumn[coloumnCount-3] //2nd last coloumn
								mapCategoryWithEstimateTime.Efficiency = coloumn[coloumnCount-2]    // last coloumn
								mapCategoryWithEstimateTime.ChangeType = coloumn[coloumnCount-1]
								mapCategoryWithEstimateTime.ActiveFlag = 1
								mapCategoryWithEstimateTime.DeleteFlag = 0
								//mapCategoryWithEstimateTime.AuditTransactionId = 1

								lastId, err := Dao.InsertMapCategoryWithEstimateTime(db, tx, &mapCategoryWithEstimateTime)
								if err != nil || lastId == 0 {
									Logger.Log.Println(err)
									tx.Rollback()
									return err
								}

							}
							err := tx.Commit()
							if err != nil {

								Logger.Log.Println(err)
								tx.Rollback()
								return err
							}
							recordDiffParentid = lastInsertedmstDiffId
						}

						i++
					}

					Logger.Log.Println("ROW No==>", rowCount)
					rowCount++
				}
			}

		}

	}
	return nil
}
