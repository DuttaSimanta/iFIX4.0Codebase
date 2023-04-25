package dao

import (
	"database/sql"
	"errors"
	"log"
	model "src/entities"
	Logger "src/logger"
)

func GetCategoryLevelNameAndId(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, recordDiffId int64) ([]string, []int64, error) {
	//var seqNo int64=0
	log.Println("GetCategoryLevelNameAndId")
	var categoryLevelNames []string
	var categoryLevelIds []int64

	var selectCategoryLevelNameAndIdQuery string = "SELECT id,typename FROM mstrecorddifferentiationtype where parentid=1 and  id in (SELECT torecorddifftypeid FROM mstrecordtype where clientid=? and mstorgnhirarchyid=? and torecorddiffid=0 and fromrecorddiffid=?) and deleteflg=0 and activeflg=1"
	categoryLevelNameAndIdResultSet, err := db.Query(selectCategoryLevelNameAndIdQuery, clientID, mstOrgnHirarchyId, recordDiffId)
	if err != nil {
		Logger.Log.Println("ERROR: categoryLevelNameAndIdResultSet Fetch Error")
		return categoryLevelNames, categoryLevelIds, errors.New("ERROR: categoryLevelNameAndIdResultSet Fetch Error")
	}

	defer categoryLevelNameAndIdResultSet.Close()
	for categoryLevelNameAndIdResultSet.Next() {
		var categoryLevelName string
		var categoryLevelId int64
		err = categoryLevelNameAndIdResultSet.Scan(&categoryLevelId, &categoryLevelName)
		if err != nil {
			Logger.Log.Println("ERROR: categoryLevelNameAndIdResultSet scan Error")
			return categoryLevelNames, categoryLevelIds, errors.New("ERROR: categoryLevelNameAndIdResultSet scan Error")
		}
		categoryLevelNames = append(categoryLevelNames, categoryLevelName)
		categoryLevelIds = append(categoryLevelIds, categoryLevelId)

		log.Printf("categoryLevelName==>%s, categoryLevelId==>%d ", categoryLevelName, categoryLevelId)
	}
	return categoryLevelNames, categoryLevelIds, nil
}

func GetImactUrgencyPriorityDetails(db *sql.DB, clientID int64, mstOrgnHirarchyId int64) ([]string, []int64, []string, []int64, []string, []int64, error) {
	var impactNames []string
	var impactIds []int64
	var urgencyNames []string
	var urgencyIds []int64
	var priorityNames []string
	var priorityIds []int64

	//Fetching impact details and storing into slice
	var selectImpactForCategory string = "select id,name from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=6 and deleteflg=0)"
	impactForCategoryResultSet, err := db.Query(selectImpactForCategory, clientID, mstOrgnHirarchyId)
	if err != nil {
		Logger.Log.Println("ERROR: impactForCategoryResultSet Fetch Error")
		return impactNames, impactIds, urgencyNames, urgencyIds, priorityNames, priorityIds, errors.New("ERROR: impactForCategoryResultSet Fetch Error")
	}
	defer impactForCategoryResultSet.Close()
	for impactForCategoryResultSet.Next() {
		var impactName string
		var impactId int64
		err = impactForCategoryResultSet.Scan(&impactId, &impactName)
		if err != nil {
			Logger.Log.Println("ERROR: impactForCategoryResultSet scan Error")
			return impactNames, impactIds, urgencyNames, urgencyIds, priorityNames, priorityIds, errors.New("ERROR: impactForCategoryResultSet scan Error")
		}
		// print out the attribute

		impactIds = append(impactIds, impactId)
		impactNames = append(impactNames, impactName)
	}

	//Fetching Urgency details and storing into slice
	var selectUrgencyForCategory string = "select id,name from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=7 and deleteflg=0)"
	urgencyForCategoryResultSet, err := db.Query(selectUrgencyForCategory, clientID, mstOrgnHirarchyId)
	if err != nil {
		Logger.Log.Println("ERROR: urgencyForCategoryResultSet Fetch Error")
		return impactNames, impactIds, urgencyNames, urgencyIds, priorityNames, priorityIds, errors.New("ERROR: urgencyForCategoryResultSet Fetch Error")
	}
	defer urgencyForCategoryResultSet.Close()
	for urgencyForCategoryResultSet.Next() {
		var urgencyName string
		var urgencyId int64
		err = urgencyForCategoryResultSet.Scan(&urgencyId, &urgencyName)
		if err != nil {
			Logger.Log.Println("ERROR: urgencyForCategoryResultSet scan Error")
			return impactNames, impactIds, urgencyNames, urgencyIds, priorityNames, priorityIds, errors.New("ERROR: urgencyForCategoryResultSet scan Error")

		}
		// print out the attribute

		urgencyIds = append(urgencyIds, urgencyId)
		urgencyNames = append(urgencyNames, urgencyName)
	}
	var selectPriorityForCategory string = "select id,name from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=4 and deleteflg=0)"
	priorityForCategoryResultSet, err := db.Query(selectPriorityForCategory, clientID, mstOrgnHirarchyId)
	if err != nil {
		Logger.Log.Println("ERROR: priorityForCategoryResultSet Fetch Error")
		return impactNames, impactIds, urgencyNames, urgencyIds, priorityNames, priorityIds, errors.New("ERROR: priorityForCategoryResultSet Fetch Error")
	}
	defer priorityForCategoryResultSet.Close()
	for priorityForCategoryResultSet.Next() {
		var priorityname string
		var priorityId int64
		err = priorityForCategoryResultSet.Scan(&priorityId, &priorityname)
		if err != nil {
			Logger.Log.Println("ERROR: priorityForCategoryResultSet scan Error")
			return impactNames, impactIds, urgencyNames, urgencyIds, priorityNames, priorityIds, errors.New("ERROR: priorityForCategoryResultSet scan Error")

		}
		// print out the attribute

		priorityIds = append(priorityIds, priorityId)
		priorityNames = append(priorityNames, priorityname)
	}
	return impactNames, impactIds, urgencyNames, urgencyIds, priorityNames, priorityIds, nil

}

func GetTemplateHeaderNamesForValidation(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, recordDiffId int64) ([]string, error) {
	var headerName []string
	Logger.Log.Println("Client===>", clientID)
	Logger.Log.Println("Record DiffId ===>", recordDiffId)
	var selectHeaderForCategoryQuery string = "select headername from mstexceltemplate where clientid=? and mstorgnhirarchyid=? and templatetypeid=3 and recorddiffid=? and  deleteflg=0 order by seqno asc"
	//fetching category header Details and storing into slice
	categoryHeadeResultSet, err := db.Query(selectHeaderForCategoryQuery, clientID, mstOrgnHirarchyId, recordDiffId)
	if err != nil {
		Logger.Log.Println(err)

		return headerName, errors.New("ERROR: Unable to fetch categoryHeadeResultSet")
	}
	defer categoryHeadeResultSet.Close()
	for categoryHeadeResultSet.Next() {
		var header string
		//	var  diffTypeId int64
		err = categoryHeadeResultSet.Scan(&header)
		if err != nil {
			Logger.Log.Println(err)

			return headerName, errors.New("ERROR: Unable to scan categoryHeadeResultSet")
		}
		headerName = append(headerName, header)
	}
	return headerName, nil
}
func GetCategoryExistCount(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, recordDiffParentid int64,
	coloumnValue string, categoryLevelId int64) (int64, error) {
	var categoryExistcount int64 = 0
	//var presentID int64 = 0
	var categoryExistQuery string = "Select count(name) as count from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and parentid=? and name =? and recorddifftypeid=?"
	err := db.QueryRow(categoryExistQuery, clientID, mstOrgnHirarchyId, recordDiffParentid, coloumnValue, categoryLevelId).Scan(&categoryExistcount)
	if err != nil {
		Logger.Log.Println(err)

		return categoryExistcount, errors.New("ERROR: Unable to scan categoryExistcount")
	}
	// if categoryExistcount > 0 {

	// 	var categoryExistQuery string = "Select count(id) as count from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and parentid=? and name =?"
	// 	err := db.QueryRow(categoryExistQuery, clientID, mstOrgnHirarchyId, recordDiffParentid, coloumnValue).Scan(&presentID, &categoryExistcount)
	// 	if err != nil {
	// 		Logger.Log.Println(err)

	// 		return categoryExistcount, errors.New("ERROR: Unable to scan categoryExistcount")
	// 	}
	// }

	return categoryExistcount, nil
}

func GetImmediatePatentsId(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, recordDiffParentid int64, coloumnValue string, categoryLevelId int64) (int64, error) {
	var parentId int64
	var getParentIDQuery string = "Select id from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and parentid=? and name =? and recorddifftypeid=?"
	getParentIDScanErr := db.QueryRow(getParentIDQuery, clientID, mstOrgnHirarchyId, recordDiffParentid, coloumnValue, categoryLevelId).Scan(&parentId)
	if getParentIDScanErr != nil {
		Logger.Log.Println(getParentIDScanErr)
		return parentId, errors.New("ERROR: Scan Error For GetImmediateParentsId ")
	}
	return parentId, nil
}

func InsertMstDiffAndMstRecord(db *sql.DB, tx *sql.Tx, mstDiff *model.MstRecordDifferentiation, mstRecord *model.MstRecordType) (int64, error) {
	var insertmstDiffQuery string = "INSERT INTO mstrecorddifferentiation (`clientid`,`mstorgnhirarchyid`,`recorddifftypeid`,`parentid`,`name`,`seqno`,`deleteflg`,`activeflg`,`parentcategoryids`,`parentcategorynames`) VALUES(?,?,?,?,?,?,?,?,?,?)"
	stmtMstDiff, err := tx.Prepare(insertmstDiffQuery)
	defer stmtMstDiff.Close()
	if err != nil {
		Logger.Log.Println(err)
		return 0, errors.New("ERROR: SQL Prepare Error in insertMstDiffAndMstRecord")
	}
	mstDiffResultSet, err := stmtMstDiff.Exec(mstDiff.ClientId, mstDiff.MstOrgnHirarchyId, mstDiff.RecordDiffTypeId, mstDiff.ParentId, mstDiff.Name, mstDiff.SeqNo, mstDiff.DeleteFlag, mstDiff.ActiveFlag, mstDiff.ParentCategoryids, mstDiff.ParentCategoryNames)
	if err != nil {
		Logger.Log.Println(err)
		return 0, errors.New("ERROR: SQL Execution Error in mstDiffResultSet")
	}
	lastInsertedmstDiffId, err := mstDiffResultSet.LastInsertId()
	if err != nil {
		Logger.Log.Println(err)
		tx.Rollback()
		return 0, errors.New("ERROR: lastInsertedmstDiffId fetching error")
	}
	mstRecord.ToRecordDiffId = lastInsertedmstDiffId
	var insertmstRecordQuery string = "INSERT INTO mstrecordtype (`clientid`,`mstorgnhirarchyid`,`fromrecorddifftypeid`,`fromrecorddiffid`,`torecorddifftypeid`,`torecorddiffid`,`deleteflg`,`activeflg`) VALUES(?,?,?,?,?,?,?,?)"
	stmtMstRecord, err := tx.Prepare(insertmstRecordQuery)
	defer stmtMstRecord.Close()
	if err != nil {
		Logger.Log.Println(err)
		tx.Rollback()
		return 0, errors.New("ERROR: SQL Prepare Error in insertMstDiffAndMstRecord")
	}
	_, er := stmtMstRecord.Exec(mstRecord.ClientId, mstRecord.MstOrgnHirarchyId, mstRecord.FromRecordDiffTypeId, mstRecord.FromRecordDiffId, mstRecord.ToRecordDiffTypeId, mstRecord.ToRecordDiffId, mstRecord.DeleteFlag, mstRecord.ActiveFlag)
	if er != nil {
		Logger.Log.Println(er)
		tx.Rollback()
		return 0, errors.New("ERROR: SQL Execution Error in stmtMstRecord")
	}
	return lastInsertedmstDiffId, nil
}

func InsertMstBusinessMatrix(db *sql.DB, tx *sql.Tx, mstBusinessMatrix *model.MstBusinessMatrix) error {
	var insertmstBusinessMatrixQuery string = "INSERT INTO mstbusinessmatrix(`clientid`,`mstorgnhirarchyid`,`mstrecorddifferentiationtickettypeid`,`mstrecorddifferentiationcatid`,`mstrecorddifferentiationimpactid`,`mstrecorddifferentiationurgencyid`,`mstrecorddifferentiationpriorityid`,`activeflg`,`deleteflg`) VALUES(?,?,?,?,?,?,?,?,?)"
	stmtMstBusinessMatrix, err := tx.Prepare(insertmstBusinessMatrixQuery)
	if err != nil {
		Logger.Log.Println(err)
		tx.Rollback()
		return errors.New("ERROR: SQL Prepare Error in InsertMstBusinessMatrix")
	}
	_, er := stmtMstBusinessMatrix.Exec(mstBusinessMatrix.ClientId, mstBusinessMatrix.MstOrgnHirarchyId, mstBusinessMatrix.MstRecordDifferentiationTicketTypeId, mstBusinessMatrix.MstRecordDifferentiationCatId, mstBusinessMatrix.MstRecordDifferentiationImpactId, mstBusinessMatrix.MstRecordDifferentiationUrgencyId, mstBusinessMatrix.MstRecordDifferentiationPriorityId, mstBusinessMatrix.ActiveFlag, mstBusinessMatrix.DeleteFlag)
	if er != nil {
		Logger.Log.Println(err)
		tx.Rollback()
		return errors.New("ERROR: SQL Execution Error in InsertMstBusinessMatrix")
	}

	log.Println("BusinessMatrix===>", mstBusinessMatrix)
	return nil
}
func InsertMapCategoryWithEstimateTime(db *sql.DB, tx *sql.Tx, mapCategoryWithEstimateTime *model.MapCategoryWithEstimateTime) (int64, error) {
	log.Println("mapCategoryWithEstimateTime====>", mapCategoryWithEstimateTime)
	var insertMapCategoryWithEstimateTime string = "INSERT INTO mapcategorywithestimatetime(`clientid`,`mstorgnhirarchyid`,`recorddiffid`,`estimatedtime`,`efficiency`,`changetype`,`deleteflg`,`activeflg`) VALUES(?,?,?,?,?,?,?,?)"
	stmtMapCategoryWithEstimateTime, err := tx.Prepare(insertMapCategoryWithEstimateTime)
	if err != nil {
		Logger.Log.Println(err)
		tx.Rollback()
		return 0, errors.New("ERROR: Prepare Error for insertMapCategoryWithEstimateTime")

	}
	MapCategoryWithEstimateTime, err1 := stmtMapCategoryWithEstimateTime.Exec(mapCategoryWithEstimateTime.ClientId, mapCategoryWithEstimateTime.MstOrgnHirarchyId, mapCategoryWithEstimateTime.RecordDiffId, mapCategoryWithEstimateTime.EstimatedTime, mapCategoryWithEstimateTime.Efficiency, mapCategoryWithEstimateTime.ChangeType, mapCategoryWithEstimateTime.DeleteFlag, mapCategoryWithEstimateTime.ActiveFlag)

	if err1 != nil {
		Logger.Log.Println(err1)
		tx.Rollback()
		return 0, errors.New("ERROR: Exec Error for insertMapCategoryWithEstimateTime")
	}
	lastInsertedId, err2 := MapCategoryWithEstimateTime.LastInsertId()
	if err2 != nil {
		Logger.Log.Println(err2)
		tx.Rollback()
		return 0, errors.New("ERROR: LastinsertedId Fetch Error for insertMapCategoryWithEstimateTime")
	}
	return lastInsertedId, nil
}
