package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertmstsupportgrp = "INSERT INTO mstsupportgrp (clientid,mstorgnhirarchyid,name,copyable) VALUES (?,?,?,?)"
var duplicatemstsupportgrp = "SELECT count(id) total FROM  mstsupportgrp WHERE name = ? AND clientid=? AND mstorgnhirarchyid=? AND copyable=? AND deleteflg = 0 AND activeflg=1"

//var getmstsupportgrp = "SELECT a.id as Id,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, a.name as SupportgrpName,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,a.copyable as Copyable FROM mstsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id ORDER BY a.id DESC LIMIT ?,?"
//var getmstsupportgrpcount = "SELECT count(a.id) total FROM mstsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id "
var updatemstsupportgrp = "UPDATE mstsupportgrp SET clientid=?,mstorgnhirarchyid=?,name=?,copyable=? WHERE id = ? "
var deletemstsupportgrp = "UPDATE mstsupportgrp SET deleteflg = '1' WHERE id = ? "
var getmstsupportgrpbycopyable = "SELECT a.id as Id,a.name as SupportgrpName from mstsupportgrp a where clientid=? and mstorgnhirarchyid=? and a.copyable=1 and a.activeflg=1 and a.deleteflg=0"
var getallcreatedsupportgrp = "SELECT a.id as Id,a.name as SupportgrpName from mstsupportgrp a where clientid=? and mstorgnhirarchyid=?  and a.activeflg=1 and a.deleteflg=0"
func (dbc DbConn) CheckDuplicatemstsupportgrp(tz *entities.MstsupportgrpEntity) (entities.MstsupportgrpEntities, error) {
	logger.Log.Println("In side CheckDuplicatemstsupportgrp")
	value := entities.MstsupportgrpEntities{}
	err := dbc.DB.QueryRow(duplicatemstsupportgrp, tz.SupportgrpName, tz.Clientid, tz.Mstorgnhirarchyid, tz.Copyable).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicatemstsupportgrp Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) Insertmstsupportgrp(tz *entities.MstsupportgrpEntity) (int64, error) {
	logger.Log.Println("In side Insertmstsupportgrp")
	logger.Log.Println("Query -->", insertmstsupportgrp)
	stmt, err := dbc.DB.Prepare(insertmstsupportgrp)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("insertmstsupportgrp Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.SupportgrpName)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.SupportgrpName, tz.Copyable)
	if err != nil {
		logger.Log.Println("insertmstsupportgrp Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllmstsupportgrp(tz *entities.MstsupportgrpEntity, OrgnType int64) ([]entities.MstsupportgrpEntity, error) {
	logger.Log.Println("In side GetAllmstsupportgrp")
	values := []entities.MstsupportgrpEntity{}
	var getmstsupportgrp string
	var params []interface{}
	if OrgnType == 1 {
		getmstsupportgrp = "SELECT a.id as Id,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, a.name as SupportgrpName,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,a.copyable as Copyable FROM mstsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getmstsupportgrp = "SELECT a.id as Id,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, a.name as SupportgrpName,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,a.copyable as Copyable FROM mstsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.clientid=? and a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getmstsupportgrp = "SELECT a.id as Id,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, a.name as SupportgrpName,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,a.copyable as Copyable FROM mstsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getmstsupportgrp, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstsupportgrp Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstsupportgrpEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.SupportgrpName, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Copyable)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) Updatemstsupportgrp(tz *entities.MstsupportgrpEntity) error {
	logger.Log.Println("In side Updatemstsupportgrp")
	stmt, err := dbc.DB.Prepare(updatemstsupportgrp)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("Updatemstsupportgrp Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.SupportgrpName, tz.Copyable, tz.Id)
	if err != nil {
		logger.Log.Println("Updatemstsupportgrp Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) Deletemstsupportgrp(tz *entities.MstsupportgrpEntity) error {
	logger.Log.Println("In side Deletemstsupportgrp")
	stmt, err := dbc.DB.Prepare(deletemstsupportgrp)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("Deletemstsupportgrp Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("Deletemstsupportgrp Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetmstsupportgrpCount(tz *entities.MstsupportgrpEntity, OrgnTypeID int64) (entities.MstsupportgrpEntities, error) {
	logger.Log.Println("In side GetmstsupportgrpCount")
	value := entities.MstsupportgrpEntities{}
	var getmstsupportgrpcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getmstsupportgrpcount = "SELECT count(a.id) total FROM mstsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id "
	} else if OrgnTypeID == 2 {
		getmstsupportgrpcount = "SELECT count(a.id) total FROM mstsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.clientid=? and a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id "
		params = append(params, tz.Clientid)
	} else {
		getmstsupportgrpcount = "SELECT count(a.id) total FROM mstsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.deleteflg =0 and a.activeflg=1 and a.clientid=b.id and a.mstorgnhirarchyid=c.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getmstsupportgrpcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetmstsupportgrpCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllmstsupportgrpbycopyable(page *entities.MstsupportgrpEntity) ([]entities.MstsupportgrpbycopyableEntity, error) {
	logger.Log.Println("In side GetAllmstsupportgrp")
	values := []entities.MstsupportgrpbycopyableEntity{}
	rows, err := dbc.DB.Query(getmstsupportgrpbycopyable, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllmstsupportgrpbycopyable Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstsupportgrpbycopyableEntity{}
		rows.Scan(&value.Id, &value.SupportgrpName)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getallcreatedsupportgrp(page *entities.MstsupportgrpEntity) ([]entities.MstsupportgrpbycopyableEntity, error) {
	logger.Log.Println("In side GetAllmstsupportgrp")
	values := []entities.MstsupportgrpbycopyableEntity{}
	rows, err := dbc.DB.Query(getallcreatedsupportgrp, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("getallcreatedsupportgrp Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstsupportgrpbycopyableEntity{}
		rows.Scan(&value.Id, &value.SupportgrpName)
		values = append(values, value)
	}
	return values, nil
}
