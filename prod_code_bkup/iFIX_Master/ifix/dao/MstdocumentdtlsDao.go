package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstdocumentdtls = "INSERT INTO mstdocumentdtls (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, groupid, documentname, documentpath, credentialid,orginaldocumentname) VALUES (?,?,?,?,?,?,?,?,?)"
var duplicateMstdocumentdtls = "SELECT count(id) total FROM  mstdocumentdtls WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND groupid = ? AND orginaldocumentname = ? AND deleteflg = 0 and activeflg=1"

//var getMstdocumentdtls = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.documentname as Documentname, a.documentpath as Documentpath, a.credentialid as Credentialid, a.activeflg as Activeflg,a.orginaldocumentname as Orginaldocumentname,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.supportgroupname as Supportgroupname,a.doccount as Usagecount FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstclientsupportgroup f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id ORDER BY a.id DESC LIMIT ?,?"
var getMstdocumentdtls = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.documentname as Documentname, a.documentpath as Documentpath, a.credentialid as Credentialid, a.activeflg as Activeflg,a.orginaldocumentname as Orginaldocumentname,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.name as Supportgroupname,a.doccount as Usagecount FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstsupportgrp f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id ORDER BY a.id DESC LIMIT ?,?"
var getMstdocumentdtlscount = "SELECT count(a.id) as total FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstsupportgrp f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id"
var updateMstdocumentdtls = "UPDATE mstdocumentdtls SET mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ?, groupid = ?, documentname = ?, documentpath = ?, credentialid = ?,orginaldocumentname=? WHERE id = ? "
var deleteMstdocumentdtls = "UPDATE mstdocumentdtls SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstdocumentdtls(tz *entities.MstdocumentdtlsEntity, grpid int64) (entities.MstdocumentdtlsEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstdocumentdtls")
	value := entities.MstdocumentdtlsEntities{}
	err := dbc.DB.QueryRow(duplicateMstdocumentdtls, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, grpid, tz.Orginaldocumentname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstdocumentdtls Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstdocumentdtls(tz *entities.MstdocumentdtlsEntity, grpid int64) (int64, error) {
	logger.Log.Println("In side InsertMstdocumentdtls")
	logger.Log.Println("Query -->", insertMstdocumentdtls)
	stmt, err := dbc.DB.Prepare(insertMstdocumentdtls)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstdocumentdtls Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, grpid, tz.Documentname, tz.Documentpath, tz.Credentialid, tz.Orginaldocumentname)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, grpid, tz.Documentname, tz.Documentpath, tz.Credentialid, tz.Orginaldocumentname)
	if err != nil {
		logger.Log.Println("InsertMstdocumentdtls Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstdocumentdtls(tz *entities.MstdocumentdtlsEntity, OrgnType int64) ([]entities.MstdocumentdtlsEntity, error) {
	logger.Log.Println("In side GelAllMstdocumentdtls")
	values := []entities.MstdocumentdtlsEntity{}
	var getMstdocumentdtls string
	var params []interface{}
	if OrgnType == 1 {
		getMstdocumentdtls = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.documentname as Documentname, a.documentpath as Documentpath, a.credentialid as Credentialid, a.activeflg as Activeflg,a.orginaldocumentname as Orginaldocumentname,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.name as Supportgroupname,a.doccount as Usagecount FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstsupportgrp f WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstdocumentdtls = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.documentname as Documentname, a.documentpath as Documentpath, a.credentialid as Credentialid, a.activeflg as Activeflg,a.orginaldocumentname as Orginaldocumentname,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.name as Supportgroupname,a.doccount as Usagecount FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstsupportgrp f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstdocumentdtls = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.groupid as Groupid, a.documentname as Documentname, a.documentpath as Documentpath, a.credentialid as Credentialid, a.activeflg as Activeflg,a.orginaldocumentname as Orginaldocumentname,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifferentiationtypename,e.name as Recorddifferentiationname,f.name as Supportgroupname,a.doccount as Usagecount FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstsupportgrp f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getMstdocumentdtls, params...)
	// rows, err := dbc.DB.Query(getMstdocumentdtls, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstdocumentdtls Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstdocumentdtlsEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Supportgroupid, &value.Documentname, &value.Documentpath, &value.Credentialid, &value.Activeflg, &value.Orginaldocumentname, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifferentiationtypename, &value.Recorddifferentiationname, &value.Supportgroupname, &value.Usagecount)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstdocumentdtls(tz *entities.MstdocumentdtlsEntity, grpid int64) error {
	logger.Log.Println("In side UpdateMstdocumentdtls")
	stmt, err := dbc.DB.Prepare(updateMstdocumentdtls)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstdocumentdtls Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, grpid, tz.Documentname, tz.Documentpath, tz.Credentialid, tz.Orginaldocumentname, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstdocumentdtls Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstdocumentdtls(tz *entities.MstdocumentdtlsEntity) error {
	logger.Log.Println("In side DeleteMstdocumentdtls")
	stmt, err := dbc.DB.Prepare(deleteMstdocumentdtls)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstdocumentdtls Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstdocumentdtls Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstdocumentdtlsCount(tz *entities.MstdocumentdtlsEntity, OrgnTypeID int64) (entities.MstdocumentdtlsEntities, error) {
	logger.Log.Println("In side GetMstdocumentdtlsCount")
	value := entities.MstdocumentdtlsEntities{}
	var getMstdocumentdtlscount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstdocumentdtlscount = "SELECT count(a.id) as total FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstsupportgrp f WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id"
	} else if OrgnTypeID == 2 {
		getMstdocumentdtlscount = "SELECT count(a.id) as total FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstsupportgrp f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id"
		params = append(params, tz.Clientid)
	} else {
		getMstdocumentdtlscount = "SELECT count(a.id) as total FROM mstdocumentdtls a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstsupportgrp f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.groupid=f.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstdocumentdtlscount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getMstdocumentdtlscount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstdocumentdtlsCount Get Statement Prepare Error", err)
		return value, err
	}
}
