package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertClientholiday = "INSERT INTO mstclientholiday (clientid, mstorgnhirarchyid, dateofholiday, plannedornot, dayofweekid, starttime, starttimeinteger, endtime, endtimeinteger) VALUES (?,?,round(UNIX_TIMESTAMP(?)),?,?,?,?,?,?)"
var duplicateClientholiday = "SELECT count(id) total FROM  mstclientholiday WHERE clientid = ? AND mstorgnhirarchyid = ? AND dateofholiday = ? AND deleteflg = 0"

// var getClientholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstclientholiday a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
// var getClientholidaycount = "SELECT count(a.id) total FROM mstclientholiday a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
var updateClientholiday = "UPDATE mstclientholiday SET mstorgnhirarchyid = ?, dateofholiday = round(UNIX_TIMESTAMP(?)), plannedornot = ?, dayofweekid = ?, starttime = ?, starttimeinteger = ?, endtime = ?, endtimeinteger = ? WHERE id = ? "
var deleteClientholiday = "UPDATE mstclientholiday SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateClientholiday(tz *entities.ClientholidayEntity) (entities.ClientholidayEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientholiday")
	value := entities.ClientholidayEntities{}
	err := dbc.DB.QueryRow(duplicateClientholiday, tz.Clientid, tz.Mstorgnhirarchyid, tz.Dateofholiday).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientholiday Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertClientholiday(tz *entities.ClientholidayEntity) (int64, error) {
	logger.Log.Println("In side InsertClientholiday")
	logger.Log.Println("Query -->", insertClientholiday)
	stmt, err := dbc.DB.Prepare(insertClientholiday)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertClientholiday Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Dateofholiday, tz.Plannedornot, tz.Dayofweekid, tz.Starttime, tz.Starttimeinteger, tz.Endtime, tz.Endtimeinteger)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Dateofholiday, tz.Plannedornot, tz.Dayofweekid, tz.Starttime, tz.Starttimeinteger, tz.Endtime, tz.Endtimeinteger)
	if err != nil {
		logger.Log.Println("InsertClientholiday Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllClientholiday(tz *entities.ClientholidayEntity, OrgnType int64) ([]entities.ClientholidayEntity, error) {
	logger.Log.Println("In side GelAllClientholiday")
	values := []entities.ClientholidayEntity{}
	var getClientholiday string
	var params []interface{}
	if OrgnType == 1 {
		getClientholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstclientholiday a,mstclient b,mstorgnhierarchy c WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getClientholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstclientholiday a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getClientholiday = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, FROM_UNIXTIME(a.dateofholiday) as Dateofholiday, a.plannedornot as Plannedornot, a.dayofweekid as Dayofweekid, a.starttime as Starttime, a.starttimeinteger as Starttimeinteger, a.endtime as Endtime, a.endtimeinteger as Endtimeinteger, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstclientholiday a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getClientholiday, params...)

	// rows, err := dbc.DB.Query(getClientholiday, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientholiday Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientholidayEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Dateofholiday, &value.Plannedornot, &value.Dayofweekid, &value.Starttime, &value.Starttimeinteger, &value.Endtime, &value.Endtimeinteger, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateClientholiday(tz *entities.ClientholidayEntity) error {
	logger.Log.Println("In side UpdateClientholiday")
	stmt, err := dbc.DB.Prepare(updateClientholiday)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateClientholiday Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Dateofholiday, tz.Plannedornot, tz.Dayofweekid, tz.Starttime, tz.Starttimeinteger, tz.Endtime, tz.Endtimeinteger, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateClientholiday Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteClientholiday(tz *entities.ClientholidayEntity) error {
	logger.Log.Println("In side DeleteClientholiday")
	stmt, err := dbc.DB.Prepare(deleteClientholiday)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteClientholiday Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteClientholiday Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetClientholidayCount(tz *entities.ClientholidayEntity, OrgnTypeID int64) (entities.ClientholidayEntities, error) {
	logger.Log.Println("In side GetClientholidayCount")
	value := entities.ClientholidayEntities{}
	var getClientholidaycount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getClientholidaycount = "SELECT count(a.id) total FROM mstclientholiday a,mstclient b,mstorgnhierarchy c WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
	} else if OrgnTypeID == 2 {
		getClientholidaycount = "SELECT count(a.id) total FROM mstclientholiday a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
		params = append(params, tz.Clientid)
	} else {
		getClientholidaycount = "SELECT count(a.id) total FROM mstclientholiday a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getClientholidaycount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getClientholidaycount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetClientholidayCount Get Statement Prepare Error", err)
		return value, err
	}
}
