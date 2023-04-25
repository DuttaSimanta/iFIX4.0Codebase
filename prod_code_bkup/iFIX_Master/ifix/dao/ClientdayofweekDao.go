package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertClientdayofweek = "INSERT INTO mstclientdayofweek (clientid, mstorgnhirarchyid, dayofweekid, starttimeinteger, starttime, endtimeinteger, endtime, nextdayforward,originallogoname,uploadedlogoname) VALUES (?,?,?,?,?,?,?,?,?,?)"
var duplicateClientdayofweek = "SELECT count(id) total FROM  mstclientdayofweek WHERE clientid = ? AND mstorgnhirarchyid = ? AND dayofweekid = ? AND deleteflg = 0"

//var getClientdayofweek = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstclientdayofweek a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
//var getClientdayofweekcount = "SELECT count(id) total FROM  mstclientdayofweek WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateClientdayofweek = "UPDATE mstclientdayofweek SET mstorgnhirarchyid = ?, dayofweekid = ?, starttimeinteger = ?, starttime = ?, endtimeinteger = ?, endtime = ?, nextdayforward = ? WHERE id = ? "
var deleteClientdayofweek = "UPDATE mstclientdayofweek SET deleteflg = '1' WHERE id = ? "
var clientorganizationname = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,b.name as Clientname,c.name as Mstorgnhirarchyname from mstclientdayofweek a,mstclient b,mstorgnhierarchy c where a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstdifferentiationtypeid = c.id"
var clientwisedayofweek = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg FROM mstclientdayofweek a WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1"

func (dbc DbConn) CheckDuplicateClientdayofweek(tz *entities.ClientdayofweekEntity, tz1 *entities.ClientdayofweekdetailsEntity) (entities.ClientdayofweekEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientdayofweek")
	value := entities.ClientdayofweekEntities{}
	err := dbc.DB.QueryRow(duplicateClientdayofweek, tz.Clientid, tz.Mstorgnhirarchyid, tz1.Dayofweekid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientdayofweek Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) InsertClientdayofweek(tz *entities.ClientdayofweekEntity, tz1 *entities.ClientdayofweekdetailsEntity) (int64, error) {
	logger.Log.Println("In side InsertClientdayofweek")
	logger.Log.Println("Query -->", insertClientdayofweek)
	stmt, err := dbc.TX.Prepare(insertClientdayofweek)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertClientdayofweek Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz1.Dayofweekid, tz1.Starttimeinteger, tz1.Starttime, tz1.Endtimeinteger, tz1.Endtime, tz1.Nextdayforward, tz.Originallogoname, tz.Uploadedfilename)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz1.Dayofweekid, tz1.Starttimeinteger, tz1.Starttime, tz1.Endtimeinteger, tz1.Endtime, tz1.Nextdayforward, tz.Originallogoname, tz.Uploadedfilename)
	if err != nil {
		logger.Log.Println("InsertClientdayofweek Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllClientdayofweek(page *entities.ClientdayofweekEntity, OrgnType int64) ([]entities.ClientdayofweekresponseEntity, error) {
	logger.Log.Println("In side GelAllClientdayofweek")
	values := []entities.ClientdayofweekresponseEntity{}
	var getClientdayofweek string
	var params []interface{}
	if OrgnType == 1 {
		getClientdayofweek = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname, a.originallogoname as Originallogoname, a.uploadedlogoname as Uploadedfilename FROM mstclientdayofweek a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getClientdayofweek = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname, a.originallogoname as Originallogoname, a.uploadedlogoname as Uploadedfilename FROM mstclientdayofweek a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getClientdayofweek = "SELECT a.id as ID,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.dayofweekid as Dayofweekid, a.starttimeinteger as Starttimeinteger, a.starttime as Starttime, a.endtimeinteger as Endtimeinteger, a.endtime as Endtime, a.nextdayforward as Nextdayforward, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname, a.originallogoname as Originallogoname, a.uploadedlogoname as Uploadedfilename FROM mstclientdayofweek a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getClientdayofweek, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientdayofweek Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientdayofweekresponseEntity{}
		rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Dayofweekid, &value.Starttimeinteger, &value.Starttime, &value.Endtimeinteger, &value.Endtime, &value.Nextdayforward, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Originallogoname, &value.Uploadedfilename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetClientdayofweekCount(tz *entities.ClientdayofweekEntity, OrgnTypeID int64) (entities.ClientdayofweekEntities, error) {
	logger.Log.Println("In side GetClientdayofweekCount")
	value := entities.ClientdayofweekEntities{}
	var getClientdayofweekcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getClientdayofweekcount = "SELECT count(id) total FROM  mstclientdayofweek WHERE  deleteflg =0 and activeflg=1"
	} else if OrgnTypeID == 2 {
		getClientdayofweekcount = "SELECT count(id) total FROM  mstclientdayofweek WHERE clientid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getClientdayofweekcount = "SELECT count(id) total FROM  mstclientdayofweek WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}

	err := dbc.DB.QueryRow(getClientdayofweekcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetClientdayofweekCount Get Statement Prepare Error", err)
		return value, err
	}
}

// func (dbc DbConn) UpdateClientdayofweek(tz *entities.ClientdayofweekEntity) error {
// 	logger.Log.Println("In side UpdateClientdayofweek")
// 	stmt, err := dbc.DB.Prepare(updateClientdayofweek)
// 	defer stmt.Close()
// 	if err != nil {
// 		logger.Log.Println("UpdateClientdayofweek Prepare Statement  Error", err)
// 		return err
// 	}
// 	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Dayofweekid, tz.Starttimeinteger, tz.Starttime, tz.Endtimeinteger, tz.Endtime, tz.Nextdayforward, tz.Id)
// 	if err != nil {
// 		logger.Log.Println("UpdateClientdayofweek Execute Statement  Error", err)
// 		return err
// 	}
// 	return nil
// }

func (dbc DbConn) DeleteClientdayofweek(tz *entities.ClientdayofweekEntity) error {
	logger.Log.Println("In side DeleteClientdayofweek")
	stmt, err := dbc.DB.Prepare(deleteClientdayofweek)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteClientdayofweek Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println("DeleteClientdayofweek Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetClientwisedayofweek(page *entities.ClientdayofweekEntity) ([]entities.ClientdayofweekresponseEntity, error) {
	logger.Log.Println("In side GelAllClientdayofweek")
	values := []entities.ClientdayofweekresponseEntity{}
	rows, err := dbc.DB.Query(clientwisedayofweek, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientdayofweek Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientdayofweekresponseEntity{}
		rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Dayofweekid, &value.Starttimeinteger, &value.Starttime, &value.Endtimeinteger, &value.Endtime, &value.Nextdayforward, &value.Activeflg)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetOrgnType(ClientID int64, OrgnID int64) (int64, error) {
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
