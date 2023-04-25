package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	//"fmt"
)

//a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
var insertMstsupportgrptermmap = "INSERT INTO mstsupportgrptermmap (clientid, mstorgnhirarchyid,termid,grpid,readpermission,writepermission) VALUES (?,?,?,?,?,?)"
var duplicateMstsupportgrptermmap = "SELECT count(id) total FROM  mstsupportgrptermmap WHERE clientid = ? AND mstorgnhirarchyid = ? AND termid = ? AND grpid=? AND readpermission=? AND writepermission=?  AND deleteflg = 0 AND activeflg=1"

//var getMstsupportgrptermmap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.termid as termid, a.grpid as Grpid, a.activeflg as Activeflg,b.name as Clientname, c.name  as Mstorgnhirarchyname, d.termname as Termname, e.supportgroupname  as Grpname  FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstclientsupportgroup e WHERE a.activeflg=1 and a.deleteflg =0 and a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id ORDER BY a.id DESC LIMIT ?,?"

// var getMstsupportgrptermmap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.termid as termid, a.grpid as Grpid, a.activeflg as Activeflg,b.name as Clientname, c.name  as Mstorgnhirarchyname, d.termname as Termname, e.name  as Grpname,a.readpermission,a.writepermission  FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstsupportgrp e WHERE a.activeflg=1 and a.deleteflg =0 and a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id ORDER BY a.id DESC LIMIT ?,?"
// var getMstsupportgrptermmapcount = "SELECT count(a.id) total FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstsupportgrp e WHERE a.activeflg=1 and a.deleteflg =0 and a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id "
var updateMstsupportgrptermmap = "UPDATE mstsupportgrptermmap SET mstorgnhirarchyid = ?, clientid = ?,termid=?,grpid=?,readpermission=?,writepermission=? WHERE id = ? "
var deleteMstsupportgrptermmap = "UPDATE mstsupportgrptermmap SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstsupportgrptermmap(tz *entities.MstsupportgrptermmapEntity, i int) (entities.MstsupportgrptermmapEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstsupportgrptermmap")
	value := entities.MstsupportgrptermmapEntities{}

	err := dbc.DB.QueryRow(duplicateMstsupportgrptermmap, tz.Clientid, tz.Mstorgnhirarchyid, tz.Termid[i], tz.Grpid, tz.Readpermission, tz.Writepermission).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstsupportgrptermmap Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstsupportgrptermmap(tz *entities.MstsupportgrptermmapEntity, i int) (int64, error) {
	logger.Log.Println("In side InsertMstsupportgrptermmap")
	logger.Log.Println("Query -->", insertMstsupportgrptermmap)
	stmt, err := dbc.DB.Prepare(insertMstsupportgrptermmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstsupportgrptermmap Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Termid[i], tz.Grpid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Termid[i], tz.Grpid, tz.Readpermission, tz.Writepermission)
	if err != nil {
		logger.Log.Println("InsertMstsupportgrptermmap Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstsupportgrptermmap(tz *entities.MstsupportgrptermmapEntity, OrgnType int64) ([]entities.MstsupportgrptermmapEntity, error) {
	logger.Log.Println("In side GelAllMstsupportgrptermmap")
	values := []entities.MstsupportgrptermmapEntity{}
	var getMstsupportgrptermmap string
	var params []interface{}
	if OrgnType == 1 {
		getMstsupportgrptermmap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.termid as termid, a.grpid as Grpid, a.activeflg as Activeflg,b.name as Clientname, c.name  as Mstorgnhirarchyname, d.termname as Termname, e.name  as Grpname,a.readpermission,a.writepermission  FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstsupportgrp e WHERE a.activeflg=1 and a.deleteflg =0 and  a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstsupportgrptermmap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.termid as termid, a.grpid as Grpid, a.activeflg as Activeflg,b.name as Clientname, c.name  as Mstorgnhirarchyname, d.termname as Termname, e.name  as Grpname,a.readpermission,a.writepermission  FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstsupportgrp e WHERE a.activeflg=1 and a.deleteflg =0 and a.clientid=? and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstsupportgrptermmap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.termid as termid, a.grpid as Grpid, a.activeflg as Activeflg,b.name as Clientname, c.name  as Mstorgnhirarchyname, d.termname as Termname, e.name  as Grpname,a.readpermission,a.writepermission  FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstsupportgrp e WHERE a.activeflg=1 and a.deleteflg =0 and a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getMstsupportgrptermmap, params...)
	// rows, err := dbc.DB.Query(getMstsupportgrptermmap, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstsupportgrptermmap Get Statement Prepare Error", err)
		return values, err
	}
	var termid int64
	for rows.Next() {
		value := entities.MstsupportgrptermmapEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &termid, &value.Grpid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Termname, &value.Grpname, &value.Readpermission, &value.Writepermission)
		value.Termid = append(value.Termid, termid)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstsupportgrptermmap(tz *entities.MstsupportgrptermmapEntity) error {
	logger.Log.Println("In side UpdateMstsupportgrptermmap")
	stmt, err := dbc.DB.Prepare(updateMstsupportgrptermmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstsupportgrptermmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Clientid, tz.Termid[0], tz.Grpid, tz.Readpermission, tz.Writepermission, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstsupportgrptermmap Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc DbConn) DeleteMstsupportgrptermmap(tz *entities.MstsupportgrptermmapEntity) error {
	logger.Log.Println("In side DeleteMstsupportgrptermmap")
	stmt, err := dbc.DB.Prepare(deleteMstsupportgrptermmap)

	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstsupportgrptermmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstsupportgrptermmap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstsupportgrptermmapCount(tz *entities.MstsupportgrptermmapEntity, OrgnTypeID int64) (entities.MstsupportgrptermmapEntities, error) {
	logger.Log.Println("In side GetMstsupportgrptermmapCount")
	value := entities.MstsupportgrptermmapEntities{}
	var getMstsupportgrptermmapcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstsupportgrptermmapcount = "SELECT count(a.id) total FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstsupportgrp e WHERE a.activeflg=1 and a.deleteflg =0 and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id "
	} else if OrgnTypeID == 2 {
		getMstsupportgrptermmapcount = "SELECT count(a.id) total FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstsupportgrp e WHERE a.activeflg=1 and a.deleteflg =0 and a.clientid=? and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id "
		params = append(params, tz.Clientid)
	} else {
		getMstsupportgrptermmapcount = "SELECT count(a.id) total FROM mstsupportgrptermmap a , mstclient b, mstorgnhierarchy c ,mstrecordterms d,mstsupportgrp e WHERE a.activeflg=1 and a.deleteflg =0 and a.clientid=? and a.mstorgnhirarchyid=? and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.termid=d.id and a.grpid=e.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstsupportgrptermmapcount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getMstsupportgrptermmapcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstsupportgrptermmapCount Get Statement Prepare Error", err)
		return value, err
	}
}
