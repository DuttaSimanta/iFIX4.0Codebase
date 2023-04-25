package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMsttemplatevariable = "INSERT INTO msttemplatevariable (clientid, mstorgnhirarchyid, templatename, tableid, fieldid) VALUES (?,?,?,?,?)"
var duplicateMsttemplatevariable = "SELECT count(id) total FROM  msttemplatevariable WHERE clientid = ? AND mstorgnhirarchyid = ? AND templatename = ? AND tableid = ? AND fieldid = ? AND deleteflg = 0"
var getMsttemplatevariable = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.templatename as Templatename, a.tableid as Tableid, a.fieldid as Fieldid, a.activeflg as Activeflg,(select name from mstclient where id = a.clientid ) as Clientname,(select name from mstorgnhierarchy where id = a.mstorgnhirarchyid ) as Mstorgnhirarchyname,(select columnname from mstdatadictionaryfield where id = a.fieldid ) as Fieldname,(select tablename from mstdatadictionarytable where id =a.tableid ) as Tablename FROM msttemplatevariable a WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
var getMsttemplatevariablecount = "SELECT count(a.id) as total FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c,mstdatadictionaryfield d,mstdatadictionarytable e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and b.id = a.clientid and c.id = a.mstorgnhirarchyid and d.id = a.fieldid and e.id =a.tableid "
var updateMsttemplatevariable = "UPDATE msttemplatevariable SET mstorgnhirarchyid = ?, templatename = ?, tableid = ?, fieldid = ? WHERE id = ? "
var deleteMsttemplatevariable = "UPDATE msttemplatevariable SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMsttemplatevariable(tz *entities.MsttemplatevariableEntity) (entities.MsttemplatevariableEntities, error) {
	logger.Log.Println("In side CheckDuplicateMsttemplatevariable")
	value := entities.MsttemplatevariableEntities{}
	err := dbc.DB.QueryRow(duplicateMsttemplatevariable, tz.Clientid, tz.Mstorgnhirarchyid, tz.Templatename, tz.Tableid, tz.Fieldid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMsttemplatevariable Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMsttemplatevariable(tz *entities.MsttemplatevariableEntity) (int64, error) {
	logger.Log.Println("In side InsertMsttemplatevariable")
	logger.Log.Println("Query -->", insertMsttemplatevariable)
	stmt, err := dbc.DB.Prepare(insertMsttemplatevariable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMsttemplatevariable Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Templatename, tz.Tableid, tz.Fieldid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Templatename, tz.Tableid, tz.Fieldid)
	if err != nil {
		logger.Log.Println("InsertMsttemplatevariable Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMsttemplatevariable(page *entities.MsttemplatevariableEntity) ([]entities.MsttemplatevariableEntity, error) {
	logger.Log.Println("In side GelAllMsttemplatevariable")
	values := []entities.MsttemplatevariableEntity{}
	rows, err := dbc.DB.Query(getMsttemplatevariable, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMsttemplatevariable Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MsttemplatevariableEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Templatename, &value.Tableid, &value.Fieldid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Fieldname, &value.Tablename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMsttemplatevariable(tz *entities.MsttemplatevariableEntity) error {
	logger.Log.Println("In side UpdateMsttemplatevariable")
	stmt, err := dbc.DB.Prepare(updateMsttemplatevariable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMsttemplatevariable Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Templatename, tz.Tableid, tz.Fieldid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMsttemplatevariable Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMsttemplatevariable(tz *entities.MsttemplatevariableEntity) error {
	logger.Log.Println("In side DeleteMsttemplatevariable")
	stmt, err := dbc.DB.Prepare(deleteMsttemplatevariable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMsttemplatevariable Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMsttemplatevariable Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMsttemplatevariableCount(tz *entities.MsttemplatevariableEntity) (entities.MsttemplatevariableEntities, error) {
	logger.Log.Println("In side GetMsttemplatevariableCount")
	value := entities.MsttemplatevariableEntities{}
	err := dbc.DB.QueryRow(getMsttemplatevariablecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMsttemplatevariableCount Get Statement Prepare Error", err)
		return value, err
	}
}
