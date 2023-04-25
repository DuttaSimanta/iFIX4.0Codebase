package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMapcommontileswithgroup = "INSERT INTO mapcommontileswithgroup (clientid, mstorgnhirarchyid, urlkey, recorddifftypeid, recorddiffid, groupid) VALUES (?,?,?,?,?,?)"
var duplicateMapcommontileswithgroup = "SELECT count(id) total FROM  mapcommontileswithgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND urlkey = ? AND recorddifftypeid = ? AND recorddiffid = ? AND groupid = ? AND deleteflg = 0"

// var getMapcommontileswithgroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.urlkey as Urlkey, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.name as Urlname,g.name as Supportgrpname FROM mapcommontileswithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,msturlkey f,mstsupportgrp g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.urlkey=f.id and a.groupid=g.id ORDER BY a.id DESC LIMIT ?,?"
// var getMapcommontileswithgroupcount = "SELECT count(a.id) as total FROM mapcommontileswithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,msturlkey f,mstsupportgrp g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.urlkey=f.id and a.groupid=g.id"
var updateMapcommontileswithgroup = "UPDATE mapcommontileswithgroup SET mstorgnhirarchyid = ?, urlkey = ?, recorddifftypeid = ?, recorddiffid = ?, groupid = ? WHERE id = ? "
var deleteMapcommontileswithgroup = "UPDATE mapcommontileswithgroup SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMapcommontileswithgroup(tz *entities.MapcommontileswithgroupEntity, grpid int64) (entities.MapcommontileswithgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapcommontileswithgroup")
	value := entities.MapcommontileswithgroupEntities{}
	err := dbc.DB.QueryRow(duplicateMapcommontileswithgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Urlkey, tz.Recorddifftypeid, tz.Recorddiffid, grpid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapcommontileswithgroup Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMapcommontileswithgroup(tz *entities.MapcommontileswithgroupEntity, grpid int64) (int64, error) {
	logger.Log.Println("In side InsertMapcommontileswithgroup")
	logger.Log.Println("Query -->", insertMapcommontileswithgroup)
	stmt, err := dbc.DB.Prepare(insertMapcommontileswithgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapcommontileswithgroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Urlkey, tz.Recorddifftypeid, tz.Recorddiffid, grpid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Urlkey, tz.Recorddifftypeid, tz.Recorddiffid, grpid)
	if err != nil {
		logger.Log.Println("InsertMapcommontileswithgroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMapcommontileswithgroup(page *entities.MapcommontileswithgroupEntity, OrgnType int64) ([]entities.MapcommontileswithgroupEntity, error) {
	logger.Log.Println("In side GelAllMapcommontileswithgroup")
	values := []entities.MapcommontileswithgroupEntity{}
	var getMapcommontileswithgroup string
	var params []interface{}
	if OrgnType == 1 {
		getMapcommontileswithgroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.urlkey as Urlkey, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.name as Urlname,g.name as Supportgrpname FROM mapcommontileswithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,msturlkey f,mstsupportgrp g WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.urlkey=f.id and a.groupid=g.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMapcommontileswithgroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.urlkey as Urlkey, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.name as Urlname,g.name as Supportgrpname FROM mapcommontileswithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,msturlkey f,mstsupportgrp g WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.urlkey=f.id and a.groupid=g.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMapcommontileswithgroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.urlkey as Urlkey, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.name as Urlname,g.name as Supportgrpname FROM mapcommontileswithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,msturlkey f,mstsupportgrp g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.urlkey=f.id and a.groupid=g.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	logger.Log.Println(getMapcommontileswithgroup)

	rows, err := dbc.DB.Query(getMapcommontileswithgroup, params...)

	// rows, err := dbc.DB.Query(getMapcommontileswithgroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMapcommontileswithgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MapcommontileswithgroupEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Urlkey, &value.Recorddifftypeid, &value.Recorddiffid, &value.Supportgrpid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifferentiationtypename, &value.Recorddifferentiationname, &value.Urlname, &value.Supportgrpname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMapcommontileswithgroup(tz *entities.MapcommontileswithgroupEntity, grpid int64) error {
	logger.Log.Println("In side UpdateMapcommontileswithgroup")
	stmt, err := dbc.DB.Prepare(updateMapcommontileswithgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMapcommontileswithgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Urlkey, tz.Recorddifftypeid, tz.Recorddiffid, grpid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMapcommontileswithgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMapcommontileswithgroup(tz *entities.MapcommontileswithgroupEntity) error {
	logger.Log.Println("In side DeleteMapcommontileswithgroup")
	stmt, err := dbc.DB.Prepare(deleteMapcommontileswithgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapcommontileswithgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMapcommontileswithgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMapcommontileswithgroupCount(tz *entities.MapcommontileswithgroupEntity, OrgnTypeID int64) (entities.MapcommontileswithgroupEntities, error) {
	logger.Log.Println("In side GetMapcommontileswithgroupCount")
	value := entities.MapcommontileswithgroupEntities{}
	var getMapcommontileswithgroupcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMapcommontileswithgroupcount = "SELECT count(a.id) as total FROM mapcommontileswithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,msturlkey f,mstsupportgrp g WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.urlkey=f.id and a.groupid=g.id"
	} else if OrgnTypeID == 2 {
		getMapcommontileswithgroupcount = "SELECT count(a.id) as total FROM mapcommontileswithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,msturlkey f,mstsupportgrp g WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.urlkey=f.id and a.groupid=g.id"
		params = append(params, tz.Clientid)
	} else {
		getMapcommontileswithgroupcount = "SELECT count(a.id) as total FROM mapcommontileswithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,msturlkey f,mstsupportgrp g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.urlkey=f.id and a.groupid=g.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMapcommontileswithgroupcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMapcommontileswithgroupcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMapcommontileswithgroupCount Get Statement Prepare Error", err)
		return value, err
	}
}
