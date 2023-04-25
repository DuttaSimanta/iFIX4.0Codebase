package dao

import (
	"database/sql"
	"errors"

	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func GetOrgName(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, rerecordDiffID int64) (string, string, error) {
	var orgName string
	var ticketTypeName string
	var OrgNameQuery string = "SELECT a.name,b.name FROM mstorgnhierarchy a,mstrecorddifferentiation b  where a.clientid = b.clientid and a.id = b.mstorgnhirarchyid and b.id=? and b.activeflg=1 and b.deleteflg=0"
	OrgNameScanErr := db.QueryRow(OrgNameQuery, rerecordDiffID).Scan(&orgName, &ticketTypeName)
	if OrgNameScanErr != nil {
		logger.Log.Println(OrgNameScanErr)
		return orgName, ticketTypeName, errors.New("ERROR: Scan Error For GetOrgName")
	}
	return orgName, ticketTypeName, nil
}

func GetUserWithGroupAndCategoryDetails(db *sql.DB, clientID int64, mstOrgnHirarchyId int64, recordDiffTypeID int64, recordDiffID int64, categoryID int64) ([]entities.UserWithGroupAndCategoryEntity, error) {
	// logger.Log.Println("Parameter -->", page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	values := []entities.UserWithGroupAndCategoryEntity{}
	var getUserDetailsForBulkDownload string
	// var params []interface{}
	//getAsset = "select a.location as location,b.name as torecorddiffname from mstlocationwisedifferentiationmap a,mstrecorddifferentiation b where a.clientid=? and a.mstorgnhirarchyid=? and a.fromrecorddiffid=? and a.torecorddiffid=b.id and a.deleteflg=0"
	//getUserDetailsForBulkDownload = "select groupid, userid FROM mapuserwithgroupandcategory where clientid=? and mstorgnhirarchyid=? and recorddifftypeid=? and recorddiffid=? and categoryid in (" + ids + ")"
	getUserDetailsForBulkDownload = "select b.name, c.loginname FROM mapuserwithgroupandcategory a, mstsupportgrp b, mstclientuser c where a.clientid=? and a.mstorgnhirarchyid=? and a.recorddifftypeid=? and a.recorddiffid=? and a.categoryid=? and a.userid=c.id and a.groupid=b.id and a.deleteflg=0"
	logger.Log.Println("In side getUserDetailsForBulkDownload==>", getUserDetailsForBulkDownload)
	rows, err := db.Query(getUserDetailsForBulkDownload, clientID, mstOrgnHirarchyId, recordDiffTypeID, recordDiffID, categoryID)

	// rows, err := dbc.DB.Query(getAsset, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllAsset Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.UserWithGroupAndCategoryEntity{}
		rows.Scan(&value.Groupname, &value.Username)
		values = append(values, value)
	}
	return values, nil
}
