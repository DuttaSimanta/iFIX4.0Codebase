package dao

import (
	//"log"
	Logger "src/logger"
	//"src/config"
	"database/sql"
	"errors"
	"fmt"
	"src/entities"
)

func CheckSheetNamePresentInDBAndGetId(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, sheetName string) (int64, error) {
	Logger.Log.Println("CheckSheetNamePresentInDB", db)
	/* db, dBerr := config.GetDB()
	if dBerr != nil{
		Logger.Log.Println(dBerr)
		return  0,errors.New("ERROR: Unable to connect DB")
	}
	defer db.Close() */
	var sheetNameId int64 = 0
	var sheetNameExistQuery string = "Select id from mstrecorddifferentiationtype where clientid=? and mstorgnhirarchyid=? and typename =? and deleteflg=0 and activeflg=1 and parentid in(select id from mstrecorddifferentiationtype where seqno=5)"
	scanErr := db.QueryRow(sheetNameExistQuery, clientID, mstOrgnHirarchyId, sheetName).Scan(&sheetNameId)
	if scanErr != nil {
		Logger.Log.Println(scanErr)
		return 0, errors.New("ERROR: ROW Scan error")
	}
	if sheetNameId < 1 {
		return 0, errors.New("ERROR: SheetName Not Present In DB")
	}
	Logger.Log.Println("sheetname id==>", sheetNameId)
	return sheetNameId, nil
}
func GetHeaderNameAndHeaderIds(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, sheetName string) ([]int64, []string, error) {
	var headerNames []string
	var headerIds []int64
	Logger.Log.Println("GetHeaderNameAndHeaderIds", db)
	/* db, dBerr := config.GetDB()
	if dBerr != nil{
		Logger.Log.Println(dBerr)
		return  headerIds,headerNames, errors.New("ERROR: Unable to connect DB")
	} */
	//defer db.Close()
	var selectHeaderQuery string = "select id,name from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and  recorddifftypeid in (select id from mstrecorddifferentiationtype where typename=? and deleteflg=0 and activeflg=1 and parentid in(select id from mstrecorddifferentiationtype where seqno=5)) and deleteflg=0 and activeflg=1 order by seqno asc"
	//fetching category header Details and storing into slice
	HeadeResultSet, err := db.Query(selectHeaderQuery, clientID, mstOrgnHirarchyId, sheetName)
	if err != nil {
		Logger.Log.Println(err)
		return headerIds, headerNames, errors.New("ERROR: Unable to Fetch Header Details from DB")
	}
	defer HeadeResultSet.Close()
	for HeadeResultSet.Next() {
		var headerId int64
		var header string
		err = HeadeResultSet.Scan(&headerId, &header)
		if err != nil {
			Logger.Log.Println(err)
			return headerIds, headerNames, errors.New("ERROR: Unable to Scan Header Details")
		}
		headerNames = append(headerNames, header)
		headerIds = append(headerIds, headerId)
	}

	return headerIds, headerNames, nil
}
func GetLastAssetId(db *sql.DB, tx *sql.Tx, clientID int64, mstOrgnHirarchyId int64) (string, error) {
	/*db, dBerr := config.GetDB()
	 if dBerr != nil{
		Logger.Log.Println(dBerr)
		return  "", errors.New("ERROR: Unable to connect DB")
	}
	defer db.Close() */
	//var code string
	var lastID int64
	var code string
	var getLastAssetIdQuery string = "select code,uid from uidgen where clientid=? and mstorgnhirarchyid=? and difftypeid=6"
	err := db.QueryRow(getLastAssetIdQuery, clientID, mstOrgnHirarchyId).Scan(&code, &lastID)
	if err != nil {
		Logger.Log.Println(err)
		return "", errors.New("ERROR: Unable to Scan Asset Id Details")
	}
	//log.Println("lastid %0.7d",lastID)
	lastID++
	assetID := fmt.Sprintf("%0.7d", lastID)
	assetID = code + assetID
	/* tx,txError := db.Begin()
	if txError != nil{
		Logger.Log.Println(txError)
		return  "",errors.New("ERROR: Unable to start transaction")
	}  */
	var updateAssetLastIdQuery string = "update uidgen set uid=? where clientid=? and mstorgnhirarchyid=? and difftypeid=6"
	stmtAssetLastId, stmtErr := tx.Prepare(updateAssetLastIdQuery)
	if stmtErr != nil {
		Logger.Log.Println(stmtErr)
		return "", errors.New("ERROR: Unable to Prepare Statement")
	}
	defer stmtAssetLastId.Close()
	//var scanLastID int64=0
	res, err := stmtAssetLastId.Exec(lastID, clientID, mstOrgnHirarchyId)
	if err != nil {
		Logger.Log.Println(err)
		return "", errors.New("ERROR: ResultSet Fetching Error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		Logger.Log.Println(err)
		return "", errors.New("ERROR: ResultSet Fetching Error")
	}
	if count != 1 {
		Logger.Log.Println("Row not Updated Properly", count)
		return "", errors.New("ERROR: Row not Updated Properly")
	}
	/* commitErr := tx.Commit()
	if commitErr != nil{
		Logger.Log.Println(stmtErr)
		return  "",errors.New("ERROR: Unable to commit  lastId")

	} */
	Logger.Log.Println("RowsAffeCted==> ", count)
	return assetID, nil
}

func InsertTrnAsset(db *sql.DB, tx *sql.Tx, trnAsset *entities.TrnAsset) (int64, error) {
	//var lastInsertedTrnAssetId int64
	/* db, dBerr := config.GetDB()
	if dBerr != nil{
		Logger.Log.Println(dBerr)
		return  0 , errors.New("ERROR: Unable to connect DB")
	}
	defer db.Close()
	tx,txError := db.Begin()
	if txError != nil{
		Logger.Log.Println(txError)
		return  0,errors.New("ERROR: Unable to create transaction")
	}  */
	var InsertTrnAssetQuery string = "INSERT INTO trnasset(`clientid`,`mstorgnhirarchyid`,`mstdifferentiationtypeid`,`assetid`,`additionalattr`,`deleteflg`,`activeflg`) VALUES(?,?,?,?,?,?,?)"
	stmtInsertTrnAsset, stmtErr := tx.Prepare(InsertTrnAssetQuery)
	if stmtErr != nil {
		tx.Rollback()
		Logger.Log.Println(stmtErr)
		return 0, errors.New("ERROR: Unable to Prepare Statement")
	}
	defer stmtInsertTrnAsset.Close()
	InsertTrnAssetResultSet, insertErr := stmtInsertTrnAsset.Exec(trnAsset.ClientId, trnAsset.MstOrgnHirarchyId, trnAsset.MstDifftypeid, trnAsset.AssetId, trnAsset.AdditionalAttr, trnAsset.DeleteFlag, trnAsset.ActiveFlag)
	if insertErr != nil {
		tx.Rollback()
		Logger.Log.Println(insertErr)
		return 0, errors.New("ERROR: InsertTrnAssetResultSet  Error")
	}
	lastInsertedTrnAssetId, lastIdError := InsertTrnAssetResultSet.LastInsertId()
	if lastIdError != nil {
		tx.Rollback()
		Logger.Log.Println(lastIdError)
		return 0, errors.New("ERROR: last Inserted Id Error in TrnAsset")
	}
	return lastInsertedTrnAssetId, nil
}

func InsertMapAssetDiff(db *sql.DB, tx *sql.Tx, mapAssetDiff *entities.MapAssetDifferentiation) error {

	var InsertMapAssetDiffQuery string = "INSERT INTO mapassetdifferentiation(`clientid`,`mstorgnhirarchyid`,`mstdifferentiationtypeid`,`mstdifferentiationid`,`trnassetid`,`value`,`deleteflg`,`activeflg`,`audittransactionid`) VALUES(?,?,?,?,?,?,?,?,?)"
	stmtInsertMapAssetDiff, stmtErr := tx.Prepare(InsertMapAssetDiffQuery)
	if stmtErr != nil {
		tx.Rollback()
		Logger.Log.Println(stmtErr)
		return errors.New("ERROR: Unable to Prepare Statement")
	}
	defer stmtInsertMapAssetDiff.Close()
	InsertMapAssetDiffResultSet, insertErr := stmtInsertMapAssetDiff.Exec(mapAssetDiff.Clientid, mapAssetDiff.Mstorgnhirarchyid, mapAssetDiff.Mstdifferentiationtypeid, mapAssetDiff.Mstdifferentiationid, mapAssetDiff.Trnassetid, mapAssetDiff.Value, mapAssetDiff.Deleteflg, mapAssetDiff.Activeflg, mapAssetDiff.AuditTransactionId)
	if insertErr != nil {
		tx.Rollback()
		Logger.Log.Println(insertErr)
		return errors.New("ERROR: InsertTrnAssetResultSet  Error")
	}
	lastInsertedMapAssetDIffId, lastIdError := InsertMapAssetDiffResultSet.LastInsertId()
	if lastIdError != nil {
		tx.Rollback()
		Logger.Log.Println(lastIdError)
		return errors.New("ERROR: Last Inserted MapAssetDiff id  fetch  Error")
	}
	if lastInsertedMapAssetDIffId == 0 {
		tx.Rollback()
		Logger.Log.Println("lastInsertedMapAssetDIffId is zero")
		return errors.New("ERROR: Last Inserted MapAssetDiff id  fetch  Error")
	}
	return nil
}
