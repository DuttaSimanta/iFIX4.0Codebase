package dao

import (
	"database/sql"
	"errors"

	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertUserWithGroupAndCategoryForBulkUpload = "INSERT INTO mapuserwithgroupandcategory (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, categoryid, groupid, userid) VALUES (?,?,?,?,?,?,?)"
var checkduplicateUserWithGroupAndCategoryForBulkUpload = "SELECT count(id) total FROM  mapuserwithgroupandcategory WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND categoryid = ? AND userid = ? AND deleteflg = 0 and activeflg=1"
var updateUserWithGroupAndCategoryForBulkUpload = "UPDATE mapuserwithgroupandcategory SET clientid = ?, mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ?, categoryid = ?, groupid = ?, userid = ? WHERE id = ? "
var getUpdatedIDForBulkUpload = "SELECT id FROM  mapuserwithgroupandcategory WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND categoryid = ? AND userid = ? AND deleteflg = 0 and activeflg=1"

func BulkUserDetails(db *sql.DB, clientID int64, mstOrgnHirarchyId int64) ([]int64, []string, error) {
	var userIds []int64
	var userLoginNames []string
	var selectUserDetails string = "select id,loginname from mstclientuser where clientid=? and mstorgnhirarchyid=? and deleteflag=0 and activeflag=1"
	UserDetailsResultSet, err := db.Query(selectUserDetails, clientID, mstOrgnHirarchyId)
	if err != nil {
		logger.Log.Println("ERROR: UserDetailsResultSet Fetch Error")
		return userIds, userLoginNames, errors.New("ERROR: UserDetailsResultSet Fetch Error")
	}
	defer UserDetailsResultSet.Close()
	for UserDetailsResultSet.Next() {
		var userId int64
		var userLoginName string
		err = UserDetailsResultSet.Scan(&userId, &userLoginName)
		if err != nil {
			logger.Log.Println("ERROR: UserDetailsResultSet scan Error")
			return userIds, userLoginNames, errors.New("ERROR: UserDetailsResultSet scan Error")

		}
		userIds = append(userIds, userId)
		userLoginNames = append(userLoginNames, userLoginName)
	}
	return userIds, userLoginNames, nil

}
func BulkGroupDetails(db *sql.DB, clientID int64, mstOrgnHirarchyId int64) ([]int64, []string, error) {
	var userIds []int64
	var userLoginNames []string
	var selectSupportGroupDetails string = "select id,name from mstsupportgrp where clientid=? and mstorgnhirarchyid=? and deleteflg=0 and activeflg=1"
	SupportGroupDetailsResultSet, err := db.Query(selectSupportGroupDetails, clientID, mstOrgnHirarchyId)
	if err != nil {
		logger.Log.Println("ERROR: SupportGroupDetailsResultSet Fetch Error", err)
		return userIds, userLoginNames, errors.New("ERROR: SupportGroupDetailsResultSet Fetch Error")
	}
	defer SupportGroupDetailsResultSet.Close()
	for SupportGroupDetailsResultSet.Next() {
		var userId int64
		var userLoginName string
		err = SupportGroupDetailsResultSet.Scan(&userId, &userLoginName)
		if err != nil {
			logger.Log.Println("ERROR: SupportGroupDetailsResultSet scan Error")
			return userIds, userLoginNames, errors.New("ERROR: SupportGroupDetailsResultSet scan Error")

		}
		userIds = append(userIds, userId)
		userLoginNames = append(userLoginNames, userLoginName)
	}
	return userIds, userLoginNames, nil

}
func AddTXUserWithGroupAndCategory(db *sql.DB, tx *sql.Tx, tz *entities.UserWithGroupAndCategoryForBulkUploadEntity) (int64, error) {
	logger.Log.Println("In side AddTXUserWithGroupAndCategory")
	logger.Log.Println("Query -->", insertUserWithGroupAndCategoryForBulkUpload)
	stmt, err := db.Prepare(insertUserWithGroupAndCategoryForBulkUpload)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("Insert Bulk User With Group And Category  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Groupid, tz.Userid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Groupid, tz.Userid)
	if err != nil {
		logger.Log.Println("Insert Bulk User With Group And Category  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func UpdateTXUserWithGroupAndCategory(db *sql.DB, tx *sql.Tx, tz *entities.UserWithGroupAndCategoryForBulkUploadEntity, updatedID int64) (int64, error) {
	logger.Log.Println("In side UpdateTXUserWithGroupAndCategory")
	logger.Log.Println("Query -->", updateUserWithGroupAndCategoryForBulkUpload)
	stmt, err := db.Prepare(updateUserWithGroupAndCategoryForBulkUpload)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("Update Bulk User With Group And Category  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Groupid, tz.Userid, updatedID)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Groupid, tz.Userid, updatedID)
	if err != nil {
		logger.Log.Println("Update Bulk User With Group And Category  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func GetTemplateHeaderNamesForValidation(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, recordDiffId int64) ([]string, error) {
	var headerName []string
	logger.Log.Println("Client===>", clientID)
	logger.Log.Println("Record DiffId ===>", recordDiffId)
	var selectHeaderForCategoryQuery string = "select headername from mstexceltemplate where clientid=? and mstorgnhirarchyid=? and templatetypeid=5 and recorddiffid=? and  deleteflg=0 order by seqno asc"
	//fetching category header Details and storing into slice
	categoryHeadeResultSet, err := db.Query(selectHeaderForCategoryQuery, clientID, mstOrgnHirarchyId, recordDiffId)
	if err != nil {
		logger.Log.Println(err)

		return headerName, errors.New("ERROR: Unable to fetch User Login ID and Support Group Name Header")
	}
	defer categoryHeadeResultSet.Close()
	for categoryHeadeResultSet.Next() {
		var header string
		//	var  diffTypeId int64
		err = categoryHeadeResultSet.Scan(&header)
		if err != nil {
			logger.Log.Println(err)

			return headerName, errors.New("ERROR: Unable to scan User Login ID and Support Group Name Header")
		}
		headerName = append(headerName, header)
	}
	return headerName, nil
}

func CheckDuplicateBulkUser(tz *entities.UserWithGroupAndCategoryForBulkUploadEntity, db *sql.DB) (entities.UserWithGroupAndCategoryForBulkUploadEntities, error) {
	logger.Log.Println("In side CheckDuplicateBulkUser")
	value := entities.UserWithGroupAndCategoryForBulkUploadEntities{}
	err := db.QueryRow(checkduplicateUserWithGroupAndCategoryForBulkUpload, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Userid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateAsset Get Statement Prepare Error", err)
		return value, err
	}
}

func GetUpdatedID(tz *entities.UserWithGroupAndCategoryForBulkUploadEntity, db *sql.DB) (int64, error) {
	logger.Log.Println("In side GetUpdatedID")
	var value int64
	var values []int64
	err := db.QueryRow(getUpdatedIDForBulkUpload, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Categoryid, tz.Userid).Scan(&value)
	values = append(values, value)
	if err != nil {
		return 0, err
	}
	if len(values) == 0 {
		return 0, nil
	} else {
		return values[0], nil
	}
}
