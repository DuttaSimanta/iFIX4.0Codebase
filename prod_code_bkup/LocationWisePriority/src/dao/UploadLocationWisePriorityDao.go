package dao

import (
	"database/sql"
	"errors"
	"src/entities"
	"src/logger"
)

func Prioritydetails(db *sql.DB, clientID int64, mstOrgnHirarchyId int64) ([]string, []int64, error) {
	var priorityNames []string
	var priorityIds []int64
	var selectPriorityForCategory string = "select id,name from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=4 and deleteflg=0)"
	priorityForCategoryResultSet, err := db.Query(selectPriorityForCategory, clientID, mstOrgnHirarchyId)
	if err != nil {
		logger.Log.Println("ERROR: priorityForCategoryResultSet Fetch Error")
		return priorityNames, priorityIds, errors.New("ERROR: priorityForCategoryResultSet Fetch Error")
	}
	defer priorityForCategoryResultSet.Close()
	for priorityForCategoryResultSet.Next() {
		var priorityname string
		var priorityId int64
		err = priorityForCategoryResultSet.Scan(&priorityId, &priorityname)
		if err != nil {
			logger.Log.Println("ERROR: priorityForCategoryResultSet scan Error")
			return priorityNames, priorityIds, errors.New("ERROR: priorityForCategoryResultSet scan Error")

		}
		// print out the attribute

		priorityIds = append(priorityIds, priorityId)
		priorityNames = append(priorityNames, priorityname)
	}
	return priorityNames, priorityIds, nil

}
func AddTXLocation(db *sql.DB, tx *sql.Tx, tz *entities.LocationPriorityEntity) (int64, error) {
	logger.Log.Println("In side InsertAsset")
	logger.Log.Println("Query -->", insertLocation)
	stmt, err := db.Prepare(insertLocation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertAsset Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.ClientID, tz.MstorgnhirarchyID, tz.Location, tz.Recorddifftypeid, tz.Recorddiffid, tz.ToRecorddifftypeid, tz.ToRecorddiffid)
	res, err := stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.Location, tz.Recorddifftypeid, tz.Recorddiffid, tz.ToRecorddifftypeid, tz.ToRecorddiffid)
	if err != nil {
		logger.Log.Println("InsertAsset Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func GetTemplateHeaderNamesForValidation(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, recordDiffId int64) ([]string, error) {
	var headerName []string
	logger.Log.Println("Client===>", clientID)
	logger.Log.Println("Record DiffId ===>", recordDiffId)
	var selectHeaderForCategoryQuery string = "select headername from mstexceltemplate where clientid=? and mstorgnhirarchyid=? and templatetypeid=4 and recorddiffid=? and  deleteflg=0 order by seqno asc"
	//fetching category header Details and storing into slice
	categoryHeadeResultSet, err := db.Query(selectHeaderForCategoryQuery, clientID, mstOrgnHirarchyId, recordDiffId)
	if err != nil {
		logger.Log.Println(err)

		return headerName, errors.New("ERROR: Unable to fetch Priority and location Header")
	}
	defer categoryHeadeResultSet.Close()
	for categoryHeadeResultSet.Next() {
		var header string
		//	var  diffTypeId int64
		err = categoryHeadeResultSet.Scan(&header)
		if err != nil {
			logger.Log.Println(err)

			return headerName, errors.New("ERROR: Unable to scan Priority And Location Header")
		}
		headerName = append(headerName, header)
	}
	return headerName, nil
}

// func GetOrgName(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, rerecordDiffID int64) (string, string, error) {
// 	var orgName string
// 	var ticketTypeName string
// 	var OrgNameQuery string = "SELECT a.name,b.name FROM mstorgnhierarchy a,mstrecorddifferentiation b  where a.clientid = b.clientid and a.id = b.mstorgnhirarchyid and b.id=? and b.activeflg=1 and b.deleteflg=0"
// 	OrgNameScanErr := db.QueryRow(OrgNameQuery, rerecordDiffID).Scan(&orgName, &ticketTypeName)
// 	if OrgNameScanErr != nil {
// 		logger.Log.Println(OrgNameScanErr)
// 		return orgName, ticketTypeName, errors.New("ERROR: Scan Error For GetOrgName")
// 	}
// 	return orgName, ticketTypeName, nil
// }
// func GetLocatioWisePriorityDetails(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, recordDiffID int64) ([]entities.LocationPriorityEntity, error) {
// 	// logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
// 	values := []entities.LocationPriorityEntity{}
// 	var getAsset string
// 	// var params []interface{}
// 	getAsset = "select a.location as location,b.name as torecorddiffname from mstlocationwisedifferentiationmap a,mstrecorddifferentiation b where a.clientid=? and a.mstorgnhirarchyid=? and a.fromrecorddiffid=? and a.torecorddiffid=b.id"
// 	logger.Log.Println("In side GelAllAsset==>", getAsset)
// 	rows, err := db.Query(getAsset, clientID, mstOrgnHirarchyId, recordDiffID)

// 	// rows, err := dbc.DB.Query(getAsset, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
// 	defer rows.Close()
// 	if err != nil {
// 		logger.Log.Println("GetAllAsset Get Statement Prepare Error", err)
// 		return values, err
// 	}
// 	for rows.Next() {
// 		value := entities.LocationPriorityEntity{}
// 		rows.Scan(&value.Location, &value.ToReccorddiffName)
// 		values = append(values, value)
// 	}
// 	return values, nil
// }
