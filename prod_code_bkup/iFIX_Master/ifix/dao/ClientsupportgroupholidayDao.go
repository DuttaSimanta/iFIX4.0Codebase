package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertClientsupportgroupholiday = "INSERT INTO mstclientsupportgroupholiday (clientid, mstorgnhirarchyid, mstclientsupportgroupid, dateofholiday, plannedornot, dayofweekid, starttime, starttimeinteger, endtime, endtimeinteger) VALUES (?,?,?,round(UNIX_TIMESTAMP(?)),?,?,?,?,?,?)"
var duplicateClientsupportgroupholiday = "SELECT count(id) total FROM  mstclientsupportgroupholiday WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstclientsupportgroupid = ? AND dateofholiday = ? AND deleteflg = 0 AND activeflg=1"

//var getClientsupportgroupholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Mstclientsupportgroupid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.supportgroupname as Supportgroupname FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstclientsupportgroup e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id ORDER BY a.id DESC LIMIT ?,?"
// var getClientsupportgroupholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Mstclientsupportgroupid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstsupportgrp e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id ORDER BY a.id DESC LIMIT ?,?"
// var getClientsupportgroupholidaycount = "SELECT count(a.id) as total FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstsupportgrp e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id"
var updateClientsupportgroupholiday = "UPDATE mstclientsupportgroupholiday SET mstorgnhirarchyid = ?, mstclientsupportgroupid = ?, dateofholiday = round(UNIX_TIMESTAMP(?)), plannedornot = ?, dayofweekid = ?, starttime = ?, starttimeinteger = ?, endtime = ?, endtimeinteger = ? WHERE id = ? "
var deleteClientsupportgroupholiday = "UPDATE mstclientsupportgroupholiday SET deleteflg = '1' WHERE id = ? "
var getsupportgrpname = "SELECT id as ID,supportgroupname as Name FROM mstclientsupportgroup WHERE clientid=? AND mstorgnhirarchyid=? AND deleteflg = 0 AND activeflg=1"

func (dbc DbConn) CheckDuplicateClientsupportgroupholiday(tz *entities.ClientsupportgroupholidayEntity) (entities.ClientsupportgroupholidayEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroupholiday")
	value := entities.ClientsupportgroupholidayEntities{}
	err := dbc.DB.QueryRow(duplicateClientsupportgroupholiday, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstclientsupportgroupid, tz.Dateofholiday).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientsupportgroupholiday Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertClientsupportgroupholiday(tz *entities.ClientsupportgroupholidayEntity) (int64, error) {
	logger.Log.Println("In side InsertClientsupportgroupholiday")
	logger.Log.Println("Query -->", insertClientsupportgroupholiday)
	stmt, err := dbc.DB.Prepare(insertClientsupportgroupholiday)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertClientsupportgroupholiday Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstclientsupportgroupid, tz.Dateofholiday, tz.Plannedornot, tz.Dayofweekid, tz.Starttime, tz.Starttimeinteger, tz.Endtime, tz.Endtimeinteger)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstclientsupportgroupid, tz.Dateofholiday, tz.Plannedornot, tz.Dayofweekid, tz.Starttime, tz.Starttimeinteger, tz.Endtime, tz.Endtimeinteger)
	if err != nil {
		logger.Log.Println("InsertClientsupportgroupholiday Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllClientsupportgroupholiday(tz *entities.ClientsupportgroupholidayEntity, OrgnType int64) ([]entities.ClientsupportgroupholidayEntity, error) {
	logger.Log.Println("In side GelAllClientsupportgroupholiday")
	values := []entities.ClientsupportgroupholidayEntity{}
	var getClientsupportgroupholiday string
	var params []interface{}
	if OrgnType == 1 {
		getClientsupportgroupholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Mstclientsupportgroupid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstsupportgrp e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getClientsupportgroupholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Mstclientsupportgroupid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstsupportgrp e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getClientsupportgroupholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.mstclientsupportgroupid as Mstclientsupportgroupid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,e.name as Supportgroupname FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstsupportgrp e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getClientsupportgroupholiday, params...)
	// rows, err := dbc.DB.Query(getClientsupportgroupholiday, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientsupportgroupholiday Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupholidayEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstclientsupportgroupid, &value.Dateofholiday, &value.Plannedornot, &value.Dayofweekid, &value.Starttime, &value.Starttimeinteger, &value.Endtime, &value.Endtimeinteger, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Supportgroupname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateClientsupportgroupholiday(tz *entities.ClientsupportgroupholidayEntity) error {
	logger.Log.Println("In side UpdateClientsupportgroupholiday")
	stmt, err := dbc.DB.Prepare(updateClientsupportgroupholiday)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroupholiday Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstclientsupportgroupid, tz.Dateofholiday, tz.Plannedornot, tz.Dayofweekid, tz.Starttime, tz.Starttimeinteger, tz.Endtime, tz.Endtimeinteger, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroupholiday Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteClientsupportgroupholiday(tz *entities.ClientsupportgroupholidayEntity) error {
	logger.Log.Println("In side DeleteClientsupportgroupholiday")
	stmt, err := dbc.DB.Prepare(deleteClientsupportgroupholiday)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroupholiday Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroupholiday Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetClientsupportgroupholidayCount(tz *entities.ClientsupportgroupholidayEntity, OrgnTypeID int64) (entities.ClientsupportgroupholidayEntities, error) {
	logger.Log.Println("In side GetClientsupportgroupholidayCount")
	value := entities.ClientsupportgroupholidayEntities{}
	var getClientsupportgroupholidaycount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getClientsupportgroupholidaycount = "SELECT count(a.id) as total FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstsupportgrp e WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id"
	} else if OrgnTypeID == 2 {
		getClientsupportgroupholidaycount = "SELECT count(a.id) as total FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstsupportgrp e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id"
		params = append(params, tz.Clientid)
	} else {
		getClientsupportgroupholidaycount = "SELECT count(a.id) as total FROM mstclientsupportgroupholiday a,mstclient b,mstorgnhierarchy c,mstsupportgrp e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstclientsupportgroupid = e.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getClientsupportgroupholidaycount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getClientsupportgroupholidaycount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetClientsupportgroupholidayCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllSupportgrpname(page *entities.ClientsupportgroupholidayEntity) ([]entities.SupportgrpEntity, error) {
	logger.Log.Println("In side GelAllClientsupportgroupholiday")
	values := []entities.SupportgrpEntity{}
	rows, err := dbc.DB.Query(getsupportgrpname, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllSupportgrpname Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.SupportgrpEntity{}
		rows.Scan(&value.ID, &value.Name)
		values = append(values, value)
	}
	return values, nil
}
