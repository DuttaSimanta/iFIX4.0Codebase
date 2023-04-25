package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstRecorddiffpriority = "INSERT INTO mstrecorddifferentiationpriority ( clientid, mstorgnhirarchyid, typedifferentiationtypeid, typedifferentiationid, differentiationtypeid, differentiationid, priority) VALUES (?,?,?,?,?,?,?)"
var duplicateMstRecorddiffpriority = "SELECT count(id) total FROM  mstrecorddifferentiationpriority WHERE clientid = ? AND mstorgnhirarchyid = ?  AND typedifferentiationtypeid=? AND typedifferentiationid=? AND differentiationtypeid=? AND differentiationid=? AND priority=? AND activeflg =1 AND deleteflg = 0 "

//var getMstRecorddiffpriority= "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.typedifferentiationtypeid as Typedifftypeid, a.typedifferentiationid as Typediffid, a.differentiationtypeid as Difftypeid,a.differentiationid as Diffid,a.priority as Priority,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname ,d.typename AS Typedifftypename,f.name AS Difftypename,e.typename AS Typediffname,g.name AS Diffname FROM mstrecorddifferentiationpriority a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.typedifferentiationtypeid=d.id  AND a.typedifferentiationid = f.id AND a.differentiationtypeid = e.id AND a.differentiationid = g.id AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 AND f.activeflg=1 AND  f.deleteflg=0 AND g.activeflg=1 AND g.deleteflg=0   ORDER BY a.id DESC LIMIT ?,?"
//var getMstRecorddiffprioritycount = "SELECT count(a.id) as total FROM mstrecorddifferentiationpriority a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.typedifferentiationtypeid=d.id  AND a.typedifferentiationid = f.id AND a.differentiationtypeid = e.id AND a.differentiationid = g.id AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 AND f.activeflg=1 AND  f.deleteflg=0 AND g.activeflg=1 AND g.deleteflg=0   ORDER BY a.id DESC "
var updateMstRecorddiffpriority = "UPDATE mstrecorddifferentiationpriority SET clientid=?,mstorgnhirarchyid = ?, typedifferentiationtypeid = ?, typedifferentiationid = ?, differentiationtypeid = ?,differentiationid=?,priority=? WHERE id = ? "
var deleteMstRecorddiffpriority = "UPDATE mstrecorddifferentiationpriority SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstRecorddiffpriority(tz *entities.MstRecorddiffpriorityEntity) (entities.MstRecorddiffpriorityEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstRecorddiffpriority ")
	value := entities.MstRecorddiffpriorityEntities{}
	err := dbc.DB.QueryRow(duplicateMstRecorddiffpriority, tz.Clientid, tz.Mstorgnhirarchyid, tz.Typedifftypeid, tz.Typediffid, tz.Difftypeid, tz.Diffid, tz.Priority).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstRecorddiffpriority Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) AddMstRecorddiffpriority(tz *entities.MstRecorddiffpriorityEntity) (int64, error) {
	logger.Log.Println("In side AddMstRecorddiffpriority")
	logger.Log.Println("Query -->", insertMstRecorddiffpriority)
	stmt, err := dbc.DB.Prepare(insertMstRecorddiffpriority)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstRecorddiffpriority Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Typedifftypeid, tz.Typediffid, tz.Difftypeid, tz.Diffid, tz.Priority)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Typedifftypeid, tz.Typediffid, tz.Difftypeid, tz.Diffid, tz.Priority)
	if err != nil {
		logger.Log.Println("AddMstRecorddiffpriority Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstRecorddiffpriority(tz *entities.MstRecorddiffpriorityEntity, OrgnType int64) ([]entities.MstRecorddiffpriorityEntity, error) {
	logger.Log.Println("In side GetAllMstRecorddiffpriority")
	values := []entities.MstRecorddiffpriorityEntity{}
	var getMstRecorddiffpriority string
	var params []interface{}
	if OrgnType == 1 {
		getMstRecorddiffpriority = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.typedifferentiationtypeid as Typedifftypeid, a.typedifferentiationid as Typediffid, a.differentiationtypeid as Difftypeid,a.differentiationid as Diffid,a.priority as Priority,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname ,d.typename AS Typedifftypename,f.name AS Difftypename,e.typename AS Typediffname,g.name AS Diffname FROM mstrecorddifferentiationpriority a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.typedifferentiationtypeid=d.id  AND a.typedifferentiationid = f.id AND a.differentiationtypeid = e.id AND a.differentiationid = g.id AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 AND f.activeflg=1 AND  f.deleteflg=0 AND g.activeflg=1 AND g.deleteflg=0   ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstRecorddiffpriority = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.typedifferentiationtypeid as Typedifftypeid, a.typedifferentiationid as Typediffid, a.differentiationtypeid as Difftypeid,a.differentiationid as Diffid,a.priority as Priority,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname ,d.typename AS Typedifftypename,f.name AS Difftypename,e.typename AS Typediffname,g.name AS Diffname FROM mstrecorddifferentiationpriority a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.typedifferentiationtypeid=d.id  AND a.typedifferentiationid = f.id AND a.differentiationtypeid = e.id AND a.differentiationid = g.id AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 AND f.activeflg=1 AND  f.deleteflg=0 AND g.activeflg=1 AND g.deleteflg=0   ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstRecorddiffpriority = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.typedifferentiationtypeid as Typedifftypeid, a.typedifferentiationid as Typediffid, a.differentiationtypeid as Difftypeid,a.differentiationid as Diffid,a.priority as Priority,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname ,d.typename AS Typedifftypename,f.name AS Difftypename,e.typename AS Typediffname,g.name AS Diffname FROM mstrecorddifferentiationpriority a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.typedifferentiationtypeid=d.id  AND a.typedifferentiationid = f.id AND a.differentiationtypeid = e.id AND a.differentiationid = g.id AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 AND f.activeflg=1 AND  f.deleteflg=0 AND g.activeflg=1 AND g.deleteflg=0   ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstRecorddiffpriority, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstRecorddiffpriority Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstRecorddiffpriorityEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Typedifftypeid, &value.Typediffid, &value.Difftypeid, &value.Diffid, &value.Priority, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Typedifftypename, &value.Typediffname, &value.Difftypename, &value.Diffname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstRecorddiffpriority(tz *entities.MstRecorddiffpriorityEntity) error {
	logger.Log.Println("In side UpdateMstRecorddiffpriority")
	stmt, err := dbc.DB.Prepare(updateMstRecorddiffpriority)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstRecorddiffpriority Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Typedifftypeid, tz.Typediffid, tz.Difftypeid, tz.Diffid, tz.Priority, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstRecorddiffpriority Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstRecorddiffpriority(tz *entities.MstRecorddiffpriorityEntity) error {
	logger.Log.Println("In side DeleteMstRecorddiffpriority")
	stmt, err := dbc.DB.Prepare(deleteMstRecorddiffpriority)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstRecorddiffpriority Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstRecorddiffpriority Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstRecorddiffpriorityCount(tz *entities.MstRecorddiffpriorityEntity, OrgnTypeID int64) (entities.MstRecorddiffpriorityEntities, error) {
	logger.Log.Println("In side GetMstRecorddiffpriorityCount")
	value := entities.MstRecorddiffpriorityEntities{}
	var getMstRecorddiffprioritycount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstRecorddiffprioritycount = "SELECT count(a.id) as total FROM mstrecorddifferentiationpriority a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.typedifferentiationtypeid=d.id  AND a.typedifferentiationid = f.id AND a.differentiationtypeid = e.id AND a.differentiationid = g.id AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 AND f.activeflg=1 AND  f.deleteflg=0 AND g.activeflg=1 AND g.deleteflg=0   ORDER BY a.id DESC "
	} else if OrgnTypeID == 2 {
		getMstRecorddiffprioritycount = "SELECT count(a.id) as total FROM mstrecorddifferentiationpriority a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.typedifferentiationtypeid=d.id  AND a.typedifferentiationid = f.id AND a.differentiationtypeid = e.id AND a.differentiationid = g.id AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 AND f.activeflg=1 AND  f.deleteflg=0 AND g.activeflg=1 AND g.deleteflg=0   ORDER BY a.id DESC "
		params = append(params, tz.Clientid)
	} else {
		getMstRecorddiffprioritycount = "SELECT count(a.id) as total FROM mstrecorddifferentiationpriority a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.typedifferentiationtypeid=d.id  AND a.typedifferentiationid = f.id AND a.differentiationtypeid = e.id AND a.differentiationid = g.id AND d.activeflg=1 AND d.deleteflg=0 AND e.activeflg=1 AND e.deleteflg=0 AND f.activeflg=1 AND  f.deleteflg=0 AND g.activeflg=1 AND g.deleteflg=0   ORDER BY a.id DESC "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstRecorddiffprioritycount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstRecorddiffpriorityCount Get Statement Prepare Error", err)
		return value, err
	}
}
