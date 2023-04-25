package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstSupportGroupWorkingHours = "INSERT INTO mstclientsupportgroupdayofweek (clientid, mstorgnhirarchyid, mstclientsupportgroupid, dayofweekid, starttimeinteger, starttime, endtimeinteger, endtime, nextdayforward) VALUES (?,?,?,?,?,?,?,?,?)"
var duplicateMstSupportGroupWorkingHours = "SELECT count(id) total FROM  mstclientsupportgroupdayofweek WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstclientsupportgroupid = ? AND dayofweekid = ? AND deleteflg = 0"

//var getMstSupportGroupWorkingHours = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstclientsupportgroupdayofweek a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
//var getMstSupportGroupWorkingHourscount = "SELECT count(id) total FROM  mstclientsupportgroupdayofweek WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMstSupportGroupWorkingHours = "UPDATE mstclientsupportgroupdayofweek SET mstorgnhirarchyid = ?, mstclientsupportgroupid = ?, dayofweekid = ?, starttimeinteger = ?, starttime = ?, endtimeinteger = ?, endtime = ?, nextdayforward = ? WHERE id = ? "
var deleteMstSupportGroupWorkingHours = "UPDATE mstclientsupportgroupdayofweek SET deleteflg = '1' WHERE id = ? "

// var clientorganizationwisename = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Supportgroupid, b.name as Clientname,c.name as Mstorgnhirarchyname from mstclientsupportgroupdayofweek a,mstclient b,mstorgnhierarchy c where a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstdifferentiationtypeid = c.id"
var SupportGroupWiseWorkingHours = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Supportgroupid, a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg FROM mstclientsupportgroupdayofweek a WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1"

func (dbc DbConn) CheckDuplicateMstSupportGroupWorkingHours(tz *entities.MstSupportGroupWorkingHoursEntity, tz1 *entities.MstSupportGroupWorkingHoursdetailsEntity) (entities.MstSupportGroupWorkingHoursEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstSupportGroupWorkingHours")
	value := entities.MstSupportGroupWorkingHoursEntities{}
	err := dbc.DB.QueryRow(duplicateMstSupportGroupWorkingHours, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz1.Dayofweekid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstSupportGroupWorkingHours Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstSupportGroupWorkingHours(tz *entities.MstSupportGroupWorkingHoursEntity, tz1 *entities.MstSupportGroupWorkingHoursdetailsEntity) (int64, error) {
	logger.Log.Println("In side InsertMstSupportGroupWorkingHours")
	logger.Log.Println("Query -->", insertMstSupportGroupWorkingHours)
	stmt, err := dbc.DB.Prepare(insertMstSupportGroupWorkingHours)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstSupportGroupWorkingHours Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz1.Dayofweekid, tz1.Starttimeinteger, tz1.Starttime, tz1.Endtimeinteger, tz1.Endtime, tz1.Nextdayforward)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz1.Dayofweekid, tz1.Starttimeinteger, tz1.Starttime, tz1.Endtimeinteger, tz1.Endtime, tz1.Nextdayforward)
	if err != nil {
		logger.Log.Println("InsertMstSupportGroupWorkingHours Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstSupportGroupWorkingHours(page *entities.MstSupportGroupWorkingHoursEntity, OrgnType int64) ([]entities.MstSupportGroupWorkingHoursresponseEntity, error) {
	logger.Log.Println("In side GelAllMstSupportGroupWorkingHours")
	values := []entities.MstSupportGroupWorkingHoursresponseEntity{}
	var getMstSupportGroupWorkingHours string
	var params []interface{}
	if OrgnType == 1 {
		getMstSupportGroupWorkingHours = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Supportgroupid, a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgroupname FROM mstclientsupportgroupdayofweek a,mstclient b,mstorgnhierarchy c,mstsupportgrp d WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.mstclientsupportgroupid = d.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstSupportGroupWorkingHours = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Supportgroupid, a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgroupname FROM mstclientsupportgroupdayofweek a,mstclient b,mstorgnhierarchy c,mstsupportgrp d WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.mstclientsupportgroupid = d.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstSupportGroupWorkingHours = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Supportgroupid, a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgroupname FROM mstclientsupportgroupdayofweek a,mstclient b,mstorgnhierarchy c,mstsupportgrp d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.mstclientsupportgroupid = d.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstSupportGroupWorkingHours, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstSupportGroupWorkingHours Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstSupportGroupWorkingHoursresponseEntity{}
		rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Supportgroupid, &value.Dayofweekid, &value.Starttimeinteger, &value.Starttime, &value.Endtimeinteger, &value.Endtime, &value.Nextdayforward, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Supportgroupname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetMstSupportGroupWorkingHoursCount(tz *entities.MstSupportGroupWorkingHoursEntity, OrgnTypeID int64) (entities.MstSupportGroupWorkingHoursEntities, error) {
	logger.Log.Println("In side GetMstSupportGroupWorkingHoursCount")
	value := entities.MstSupportGroupWorkingHoursEntities{}
	var getMstSupportGroupWorkingHourscount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstSupportGroupWorkingHourscount = "SELECT count(id) total FROM  mstclientsupportgroupdayofweek WHERE  deleteflg =0 and activeflg=1"
	} else if OrgnTypeID == 2 {
		getMstSupportGroupWorkingHourscount = "SELECT count(id) total FROM  mstclientsupportgroupdayofweek WHERE clientid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getMstSupportGroupWorkingHourscount = "SELECT count(id) total FROM  mstclientsupportgroupdayofweek WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}

	err := dbc.DB.QueryRow(getMstSupportGroupWorkingHourscount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstSupportGroupWorkingHoursCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) UpdateMstSupportGroupWorkingHours(tz *entities.MstSupportGroupWorkingHoursUpdateEntity) error {
	logger.Log.Println("In side UpdateMstSupportGroupWorkingHours")
	stmt, err := dbc.DB.Prepare(updateMstSupportGroupWorkingHours)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstSupportGroupWorkingHours Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Dayofweekid, tz.Starttimeinteger, tz.Starttime, tz.Endtimeinteger, tz.Endtime, tz.Nextdayforward, tz.ID)
	if err != nil {
		logger.Log.Println("UpdateMstSupportGroupWorkingHours Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstSupportGroupWorkingHours(tz *entities.MstSupportGroupWorkingHoursEntity) error {
	logger.Log.Println("In side DeleteMstSupportGroupWorkingHours")
	stmt, err := dbc.DB.Prepare(deleteMstSupportGroupWorkingHours)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstSupportGroupWorkingHours Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println("DeleteMstSupportGroupWorkingHours Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetSupportGroupWiseWorkingHours(page *entities.MstSupportGroupWorkingHoursEntity) ([]entities.MstSupportGroupWorkingHoursresponseEntity, error) {
	logger.Log.Println("In side GelAllMstSupportGroupWorkingHours")
	values := []entities.MstSupportGroupWorkingHoursresponseEntity{}
	rows, err := dbc.DB.Query(SupportGroupWiseWorkingHours, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstSupportGroupWorkingHours Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstSupportGroupWorkingHoursresponseEntity{}
		rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Dayofweekid, &value.Starttimeinteger, &value.Starttime, &value.Endtimeinteger, &value.Endtime, &value.Nextdayforward, &value.Activeflg)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetOrganizationType(ClientID int64, OrgnID int64) (int64, error) {
	logger.Log.Println("In side GetOrgnType")
	var OrgnTypeID int64
	var sql = "SELECT mstorgnhierarchytypeid FROM mstorgnhierarchy WHERE clientid=? AND id=?"
	stmt, err := dbc.DB.Prepare(sql)
	if err != nil {
		logger.Log.Println(err)
		return OrgnTypeID, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(ClientID, OrgnID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetOrgnType Get Statement Prepare Error", err)
		return OrgnTypeID, err
	}
	for rows.Next() {
		err := rows.Scan(&OrgnTypeID)
		logger.Log.Println("Error is >>>>>>>", err)
	}
	return OrgnTypeID, nil
}
