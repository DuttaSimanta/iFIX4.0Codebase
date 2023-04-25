package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstprocessadmin = "INSERT INTO mstprocessadmin (clientid, mstorgnhirarchyid, processid, userid) VALUES (?,?,?,?)"
var duplicateMstprocessadmin = "SELECT count(id) total FROM  mstprocessadmin WHERE clientid = ? AND mstorgnhirarchyid = ? AND processid = ? AND userid = ? AND deleteflg = 0 AND activeflg=1"

// var getMstprocessadmin = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.processid as Processid, a.userid as Userid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.name as Username FROM mstprocessadmin a,mstclient b,mstorgnhierarchy c,mstprocess d,mstuser e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND a.userid=e.id ORDER BY a.id DESC LIMIT ?,?"
// var getMstprocessadmincount = "SELECT count(a.id) as total FROM mstprocessadmin a,mstclient b,mstorgnhierarchy c,mstprocess d,mstuser e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND a.userid=e.id"
var updateMstprocessadmin = "UPDATE mstprocessadmin SET mstorgnhirarchyid = ?, processid = ?, userid = ? WHERE id = ? "
var deleteMstprocessadmin = "UPDATE mstprocessadmin SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstprocessadmin(tz *entities.MstprocessadminEntity) (entities.MstprocessadminEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstprocessadmin")
	value := entities.MstprocessadminEntities{}
	err := dbc.DB.QueryRow(duplicateMstprocessadmin, tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid, tz.Refuserid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstprocessadmin Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstprocessadmin(tz *entities.MstprocessadminEntity) (int64, error) {
	logger.Log.Println("In side InsertMstprocessadmin")
	logger.Log.Println("Query -->", insertMstprocessadmin)
	stmt, err := dbc.DB.Prepare(insertMstprocessadmin)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstprocessadmin Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid, tz.Refuserid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processid, tz.Refuserid)
	if err != nil {
		logger.Log.Println("InsertMstprocessadmin Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstprocessadmin(tz *entities.MstprocessadminEntity, OrgnType int64) ([]entities.MstprocessadminEntity, error) {
	logger.Log.Println("In side GelAllMstprocessadmin")
	values := []entities.MstprocessadminEntity{}

	var getMstprocessadmin string
	var params []interface{}
	if OrgnType == 1 {
		getMstprocessadmin = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.processid as Processid, a.userid as Userid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.name as Username FROM mstprocessadmin a,mstclient b,mstorgnhierarchy c,mstprocess d,mstuser e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND a.userid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstprocessadmin = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.processid as Processid, a.userid as Userid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.name as Username FROM mstprocessadmin a,mstclient b,mstorgnhierarchy c,mstprocess d,mstuser e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND a.userid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstprocessadmin = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.processid as Processid, a.userid as Userid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.processname as Processname,e.name as Username FROM mstprocessadmin a,mstclient b,mstorgnhierarchy c,mstprocess d,mstuser e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND a.userid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getMstprocessadmin, params...)

	// rows, err := dbc.DB.Query(getMstprocessadmin, tz.Clientid, tz.Mstorgnhirarchyid, tz.Offset, tz.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstprocessadmin Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstprocessadminEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Processid, &value.Refuserid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Processname, &value.Username)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstprocessadmin(tz *entities.MstprocessadminEntity) error {
	logger.Log.Println("In side UpdateMstprocessadmin")
	stmt, err := dbc.DB.Prepare(updateMstprocessadmin)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstprocessadmin Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Processid, tz.Refuserid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstprocessadmin Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstprocessadmin(tz *entities.MstprocessadminEntity) error {
	logger.Log.Println("In side DeleteMstprocessadmin")
	stmt, err := dbc.DB.Prepare(deleteMstprocessadmin)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstprocessadmin Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstprocessadmin Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstprocessadminCount(tz *entities.MstprocessadminEntity, OrgnTypeID int64) (entities.MstprocessadminEntities, error) {
	logger.Log.Println("In side GetMstprocessadminCount")
	value := entities.MstprocessadminEntities{}
	var getMstprocessadmincount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstprocessadmincount = "SELECT count(a.id) as total FROM mstprocessadmin a,mstclient b,mstorgnhierarchy c,mstprocess d,mstuser e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND a.userid=e.id"
	} else if OrgnTypeID == 2 {
		getMstprocessadmincount = "SELECT count(a.id) as total FROM mstprocessadmin a,mstclient b,mstorgnhierarchy c,mstprocess d,mstuser e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND a.userid=e.id"
		params = append(params, tz.Clientid)
	} else {
		getMstprocessadmincount = "SELECT count(a.id) as total FROM mstprocessadmin a,mstclient b,mstorgnhierarchy c,mstprocess d,mstuser e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.processid=d.id AND a.userid=e.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstprocessadmincount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstprocessadmincount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstprocessadminCount Get Statement Prepare Error", err)
		return value, err
	}
}
