package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstclientsla = "INSERT INTO mstclientsla (clientid, mstorgnhirarchyid, slaname, slatimereset, slaupgradereset, sladowngradereset) VALUES (?,?,?,?,?,?)"
var duplicateMstclientsla = "SELECT count(id) total FROM  mstclientsla WHERE clientid = ? AND mstorgnhirarchyid = ? AND slaname = ? AND slatimereset = ? AND slaupgradereset = ? AND sladowngradereset = ? AND deleteflg = 0 AND activeflg=1"

// var getMstclientsla = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.slaname as Slaname, a.slatimereset as Slatimereset, a.slaupgradereset as Slaupgradereset, a.sladowngradereset as Sladowngradereset, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname FROM mstclientsla a WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
// var getMstclientslacount = "SELECT count(a.id) total FROM  mstclientsla a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid =c.id"
var updateMstclientsla = "UPDATE mstclientsla SET mstorgnhirarchyid = ?, slaname = ?, slatimereset = ?, slaupgradereset = ?, sladowngradereset = ? WHERE id = ? "
var deleteMstclientsla = "UPDATE mstclientsla SET deleteflg = '1' WHERE id = ? "
var getslanames = "select id,slaname from mstclientsla where clientid = ? AND mstorgnhirarchyid = ? AND deleteflg = 0 AND activeflg=1 "

func (dbc DbConn) CheckDuplicateMstclientsla(tz *entities.MstclientslaEntity) (entities.MstclientslaEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstclientsla")
	value := entities.MstclientslaEntities{}
	err := dbc.DB.QueryRow(duplicateMstclientsla, tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaname, tz.Slatimereset, tz.Slaupgradereset, tz.Sladowngradereset).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstclientsla Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstclientsla(tz *entities.MstclientslaEntity) (int64, error) {
	logger.Log.Println("In side InsertMstclientsla")
	logger.Log.Println("Query -->", insertMstclientsla)
	stmt, err := dbc.DB.Prepare(insertMstclientsla)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstclientsla Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaname, tz.Slatimereset, tz.Slaupgradereset, tz.Sladowngradereset)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Slaname, tz.Slatimereset, tz.Slaupgradereset, tz.Sladowngradereset)
	if err != nil {
		logger.Log.Println("InsertMstclientsla Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstclientsla(page *entities.MstclientslaEntity, OrgnType int64) ([]entities.MstclientslaEntity, error) {
	logger.Log.Println("In side GelAllMstclientsla")
	values := []entities.MstclientslaEntity{}
	var getMstclientsla string
	var params []interface{}
	if OrgnType == 1 {
		getMstclientsla = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.slaname as Slaname, a.slatimereset as Slatimereset, a.slaupgradereset as Slaupgradereset, a.sladowngradereset as Sladowngradereset, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname FROM mstclientsla a WHERE a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstclientsla = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.slaname as Slaname, a.slatimereset as Slatimereset, a.slaupgradereset as Slaupgradereset, a.sladowngradereset as Sladowngradereset, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname FROM mstclientsla a WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstclientsla = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.slaname as Slaname, a.slatimereset as Slatimereset, a.slaupgradereset as Slaupgradereset, a.sladowngradereset as Sladowngradereset, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname FROM mstclientsla a WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstclientsla, params...)

	// rows, err := dbc.DB.Query(getMstclientsla, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstclientsla Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstclientslaEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Slaname, &value.Slatimereset, &value.Slaupgradereset, &value.Sladowngradereset, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstclientsla(tz *entities.MstclientslaEntity) error {
	logger.Log.Println("In side UpdateMstclientsla")
	stmt, err := dbc.DB.Prepare(updateMstclientsla)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstclientsla Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Slaname, tz.Slatimereset, tz.Slaupgradereset, tz.Sladowngradereset, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstclientsla Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstclientsla(tz *entities.MstclientslaEntity) error {
	logger.Log.Println("In side DeleteMstclientsla")
	stmt, err := dbc.DB.Prepare(deleteMstclientsla)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstclientsla Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstclientsla Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstclientslaCount(tz *entities.MstclientslaEntity, OrgnTypeID int64) (entities.MstclientslaEntities, error) {
	logger.Log.Println("In side GetMstclientslaCount")
	value := entities.MstclientslaEntities{}

	var getMstclientslacount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstclientslacount = "SELECT count(a.id) total FROM  mstclientsla a,mstclient b,mstorgnhierarchy c WHERE   a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid =c.id"
	} else if OrgnTypeID == 2 {
		getMstclientslacount = "SELECT count(a.id) total FROM  mstclientsla a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND   a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid =c.id"
		params = append(params, tz.Clientid)
	} else {
		getMstclientslacount = "SELECT count(a.id) total FROM  mstclientsla a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid =c.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstclientslacount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstclientslacount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstclientslaCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetSlanames(page *entities.MstclientslaEntity) ([]entities.Mstslaname, error) {
	logger.Log.Println("In side GelAllMstclientsla")
	values := []entities.Mstslaname{}
	rows, err := dbc.DB.Query(getslanames, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetSlanames Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Mstslaname{}
		rows.Scan(&value.Id, &value.Slaname)
		values = append(values, value)
	}
	return values, nil
}
