package entitiess

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"src/config"
	Dao "src/dao"
	"src/entities"
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
func ExcelTemplateAndValueCheck(db *sql.DB, excelFile *Excel.File, clientID int64, mstOrgnHirarchyId int64) error {
	//	var headerName [] string
	for _, sheet := range excelFile.Sheets {
		sheetId, sheetNameErr := Dao.CheckSheetNamePresentInDBAndGetId(db, clientID, mstOrgnHirarchyId, sheet.Name)
		if sheetId == 0 {
			Logger.Log.Println(sheetNameErr)
			return errors.New("ERROR: DiffType Fetch Error ")
		}
		if sheetNameErr != nil {
			Logger.Log.Println(sheetNameErr)
			return sheetNameErr
		} else {
			headerIds, headerName, headerError := Dao.GetHeaderNameAndHeaderIds(db, clientID, mstOrgnHirarchyId, sheet.Name)
			if headerError != nil {
				Logger.Log.Println(headerError)
				return sheetNameErr
			}
			Logger.Log.Println("HeaderIDS=>", headerIds)
			Logger.Log.Println("headerName=>", headerName)
			var rowCount int64 = 0
			for _, row := range sheet.Rows {
				var coloumnCount int64 = 0
				for _, cell := range row.Cells {
					if rowCount == 0 {
						// if len(row.Cells) != len(headerName) {
						// 	Logger.Log.Println("Header Length Not Matched", len(row.Cells))
						// 	return errors.New("Header Length Not Matched")
						// }
						text := cell.String()
						text = strings.Trim(text, " ")
						Logger.Log.Println(text)
						Logger.Log.Println(headerName[coloumnCount])
						if !strings.EqualFold(text, headerName[coloumnCount]) {
							Logger.Log.Println("Header Template not matched")
							return errors.New("Header Template not matched")
						}

					} else {

					}

					coloumnCount++
				}

				rowCount++
			}
		}
	}

	return nil
}
func AssetUpload(clientID int64, mstOrgnHirarchyId int64, originalFileName string, uploadedFileName string) error {

	Logger.Log.Println("In assetUpload Service")
	//Logger.Log.Println("clientid===>",clientID)
	//Logger.Log.Println("MSt===>",mstOrgnHirarchyId)
	//Logger.Log.Println("Filename==>",filename)
	//Logger.Log.Println("url===>",url)
	contextPath, contextPatherr := getContextPath()
	if contextPatherr != nil {
		Logger.Log.Println(contextPatherr)
		return contextPatherr
	}
	filePath := contextPath + "/resource/downloads/" + originalFileName

	fileDownloadErr := FileUtils.DownloadFileFromUrl(clientID, mstOrgnHirarchyId, originalFileName, uploadedFileName, filePath)
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
		//getRecord
		db, dBerr := config.GetDB()
		if dBerr != nil {
			Logger.Log.Println(dBerr)
			return errors.New("ERROR: Unable to connect DB")
		}

		excelTemplateAndValueCheckError := ExcelTemplateAndValueCheck(db, excelFile, clientID, mstOrgnHirarchyId)

		if excelTemplateAndValueCheckError != nil {
			Logger.Log.Println(excelTemplateAndValueCheckError)
			return errors.New("ERROR: Invalid Excel")
		} else {

			for _, sheet := range excelFile.Sheets {
				diffTypeID, fetchDiffTypeError := Dao.CheckSheetNamePresentInDBAndGetId(db, clientID, mstOrgnHirarchyId, sheet.Name)
				if fetchDiffTypeError != nil {
					Logger.Log.Println(fetchDiffTypeError)
					return errors.New("ERROR: DiffType Fetch Error ")
				}
				headerIds, headerName, headerError := Dao.GetHeaderNameAndHeaderIds(db, clientID, mstOrgnHirarchyId, sheet.Name)
				if headerError != nil {
					Logger.Log.Println(headerError)
					return headerError
				}
				Logger.Log.Println("HeaderName==>", headerName)
				var rowCount int64 = 1
				for _, row := range sheet.Rows[1:] {
					var coloumnCount int64 = 0
					var trnAsset entities.TrnAsset
					//tx,txError := db.Begin()
					tx, txError := db.Begin()
					if txError != nil {
						Logger.Log.Println(txError)
						return errors.New("ERROR: Unable to start transaction")
					}
					assetID, err := Dao.GetLastAssetId(db, tx, clientID, mstOrgnHirarchyId)
					if err != nil {
						tx.Rollback()
						Logger.Log.Println(fetchDiffTypeError)
						return errors.New("ERROR: DiffType Fetch Error ")
					}
					Logger.Log.Println("Asset ID===>", assetID)
					//Logger.Log.Println("code===>",assetID,err)
					trnAsset.ClientId = clientID
					trnAsset.MstOrgnHirarchyId = mstOrgnHirarchyId
					trnAsset.MstDifftypeid = diffTypeID
					trnAsset.AssetId = assetID
					trnAsset.AdditionalAttr = ""
					trnAsset.ActiveFlag = 1
					trnAsset.DeleteFlag = 0
					//trnAsset.AuditTransactionId = 1

					lastInsertedTrnAssetId, insertTrnAssetError := Dao.InsertTrnAsset(db, tx, &trnAsset)
					if insertTrnAssetError != nil {
						tx.Rollback()
						Logger.Log.Println(fetchDiffTypeError)
						return errors.New("ERROR: TrnAssetInsert Error ")
					}
					if lastInsertedTrnAssetId == 0 {
						tx.Rollback()
						Logger.Log.Println(fetchDiffTypeError)
						return errors.New("ERROR: LastTrnAssetInsertedId Error ")
					}
					var coloumn []string
					for _, cell := range row.Cells {
						var mapAssetDiff entities.MapAssetDifferentiation

						text := cell.String()
						text = strings.Trim(text, " ")
						coloumn = append(coloumn, text)
						//log.Printf("%s\n", text)
						mapAssetDiff.Clientid = clientID
						mapAssetDiff.Mstorgnhirarchyid = mstOrgnHirarchyId
						mapAssetDiff.Mstdifferentiationtypeid = diffTypeID
						mapAssetDiff.Mstdifferentiationid = headerIds[coloumnCount]
						mapAssetDiff.Trnassetid = lastInsertedTrnAssetId
						mapAssetDiff.Value = text
						mapAssetDiff.Deleteflg = 0
						mapAssetDiff.Activeflg = 1
						//	mapAssetDiff.AuditTransactionId = 1
						mapAssetDiffError := Dao.InsertMapAssetDiff(db, tx, &mapAssetDiff)
						if mapAssetDiffError != nil {
							Logger.Log.Println(mapAssetDiffError)
						}

						coloumnCount++
					}
					Logger.Log.Println("Row==>", rowCount)
					Logger.Log.Println("Row==>", rowCount)
					Logger.Log.Println("Coloumn==>", coloumn)
					Logger.Log.Println("Coloumn==>", coloumn)

					commitErr := tx.Commit()
					if commitErr != nil {
						Logger.Log.Println(commitErr)
						return errors.New("ERROR: Unable to commit  Asset")

					}
					rowCount++
				}
			}
		}
	}

	return nil
}
