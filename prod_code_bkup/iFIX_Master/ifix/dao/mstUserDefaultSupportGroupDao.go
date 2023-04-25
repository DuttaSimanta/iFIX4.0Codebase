package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstUserDefaultSupportGroup = "INSERT INTO mstdefaultsupportgrp (clientid, mstorgnhirarchyid, userid, groupid) VALUES (?,?,?,?)"
var duplicateMstUserDefaultSupportGroup = "SELECT count(id) total FROM  mstdefaultsupportgrp WHERE clientid = ? AND mstorgnhirarchyid = ? AND userid = ? AND deleteflg = 0"
var duplicateMstUserDefaultSupportGroupUpdate = "SELECT count(a.id) as total FROM mstdefaultsupportgrp a WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.userid = ? AND a.id <> ?"

//var getMstUserDefaultSupportGroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.userid as Refuserid, a.groupid as Groupid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Refusername,e.name as Groupname FROM mstdefaultsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"

//var getMstUserDefaultSupportGroupcount = "SELECT count(a.id) as total FROM mstdefaultsupportgrp a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id"

var updateMstUserDefaultSupportGroup = "UPDATE mstdefaultsupportgrp SET clientid = ?, mstorgnhirarchyid = ?, userid = ?, groupid = ? WHERE id = ? "
var deleteMstUserDefaultSupportGroup = "UPDATE mstdefaultsupportgrp SET deleteflg = '1' WHERE id = ? "


var duplicateMstUserSupportGroupChange = "SELECT id FROM mstdefaultsupportgrp WHERE clientid = ? AND mstorgnhirarchyid = ? AND userid = ? AND deleteflg = 0"


func (dbc DbConn) CheckDuplicateMstUserDefaultSupportGroup(tz *entities.MstUserDefaultSupportGroupEntity) (entities.MstUserDefaultSupportGroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstUserDefaultSupportGroup")
	value := entities.MstUserDefaultSupportGroupEntities{}
	err := dbc.DB.QueryRow(duplicateMstUserDefaultSupportGroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Refuserid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstUserDefaultSupportGroup Get Statement Prepare Error", err)
		return value, err
	}
}


func (dbc DbConn) CheckDuplicateMstUserDefaultSupportGroupUpdate(tz *entities.MstUserDefaultSupportGroupEntity) (entities.MstUserDefaultSupportGroupEntities, error) {
	logger.Log.Println("In side duplicateMstUserDefaultSupportGroupUpdate")
	value := entities.MstUserDefaultSupportGroupEntities{}
	err := dbc.DB.QueryRow(duplicateMstUserDefaultSupportGroupUpdate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Refuserid, tz.Id).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("duplicateMstUserDefaultSupportGroupUpdate Get Statement Prepare Error", err)
		return value, err
	}
}


func (dbc DbConn) InsertMstUserDefaultSupportGroup(tz *entities.MstUserDefaultSupportGroupEntity) (int64, error) {
	logger.Log.Println("In side InsertMstUserDefaultSupportGroup")
	logger.Log.Println("Query -->", insertMstUserDefaultSupportGroup)
	stmt, err := dbc.DB.Prepare(insertMstUserDefaultSupportGroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstUserDefaultSupportGroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Refuserid, tz.Groupid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Refuserid, tz.Groupid)
	if err != nil {
		logger.Log.Println("InsertMstUserDefaultSupportGroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) UpdateMstUserDefaultSupportGroup(tz *entities.MstUserDefaultSupportGroupEntity) error {
	logger.Log.Println("In side UpdateMstUserDefaultSupportGroup")
	stmt, err := dbc.DB.Prepare(updateMstUserDefaultSupportGroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstUserDefaultSupportGroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Refuserid, tz.Groupid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstUserDefaultSupportGroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstUserDefaultSupportGroup(tz *entities.MstUserDefaultSupportGroupEntity) error {
	logger.Log.Println("In side DeleteMstUserDefaultSupportGroup")
	stmt, err := dbc.DB.Prepare(deleteMstUserDefaultSupportGroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstUserDefaultSupportGroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstUserDefaultSupportGroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstUserDefaultSupportGroupCount(tz *entities.MstUserDefaultSupportGroupEntity, OrgnTypeID int64) (entities.MstUserDefaultSupportGroupEntities, error) {
	logger.Log.Println("In side GetMstUserDefaultSupportGroupCount")
	value := entities.MstUserDefaultSupportGroupEntities{}
	var getMstUserDefaultSupportGroupcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstUserDefaultSupportGroupcount = "SELECT count(a.id) as total FROM mstdefaultsupportgrp a,mstclient b,mstorgnhierarchy c,mstclientuser d,mstsupportgrp e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.userid = d.id AND a.groupid = e.id "
	} else if OrgnTypeID == 2 {
		getMstUserDefaultSupportGroupcount = "SELECT count(a.id) as total FROM mstdefaultsupportgrp a,mstclient b,mstorgnhierarchy c,mstclientuser d,mstsupportgrp e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.userid = d.id AND a.groupid = e.id "
		params = append(params, tz.Clientid)
	} else {
		getMstUserDefaultSupportGroupcount = "SELECT count(a.id) as total FROM mstdefaultsupportgrp a,mstclient b,mstorgnhierarchy c,mstclientuser d,mstsupportgrp e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.userid = d.id AND a.groupid = e.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstUserDefaultSupportGroupcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstUserDefaultSupportGroupCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllMstUserDefaultSupportGroup(tz *entities.MstUserDefaultSupportGroupEntity, OrgnType int64) ([]entities.MstUserDefaultSupportGroupEntity, error) {
	logger.Log.Println("In side GelAllMstUserDefaultSupportGroup")
	values := []entities.MstUserDefaultSupportGroupEntity{}
	var getMstUserDefaultSupportGroup string
	var params []interface{}
	if OrgnType == 1 {
		getMstUserDefaultSupportGroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.userid as Refuserid, a.groupid as Groupid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Refusername,e.name as Groupname FROM mstdefaultsupportgrp a,mstclient b,mstorgnhierarchy c,mstclientuser d,mstsupportgrp e WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.userid = d.id AND a.groupid = e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstUserDefaultSupportGroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.userid as Refuserid, a.groupid as Groupid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Refusername,e.name as Groupname FROM mstdefaultsupportgrp a,mstclient b,mstorgnhierarchy c,mstclientuser d,mstsupportgrp e WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.userid = d.id AND a.groupid = e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstUserDefaultSupportGroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.userid as Refuserid, a.groupid as Groupid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Refusername,e.name as Groupname FROM mstdefaultsupportgrp a,mstclient b,mstorgnhierarchy c,mstclientuser d,mstsupportgrp e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.userid = d.id AND a.groupid = e.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstUserDefaultSupportGroup, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstUserDefaultSupportGroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserDefaultSupportGroupEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Refuserid, &value.Groupid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname,&value.Refusername,&value.Groupname)
		values = append(values, value)
	}
	return values, nil
}







func (dbc DbConn) CheckDuplicateMstUserSupportGroupChange(tz *entities.MstUserDefaultSupportGroupEntity) ([]int, error) {
	logger.Log.Println("In side duplicateMstUserSupportGroupChange")
	var values []int
	rows, err := dbc.DB.Query(duplicateMstUserSupportGroupChange, tz.Clientid, tz.Mstorgnhirarchyid, tz.Refuserid)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllUserRoleAction Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		var value int
		rows.Scan(&value)
		// logger.Log.Println("value 11111  -->", value)
		values = append(values, value)
	}
	// logger.Log.Println("values 22222  -->", values)
	return values, nil
}

func (dbc DbConn) UpdateMstUserSupportGroupChange(tz *entities.MstUserDefaultSupportGroupEntity, ids []int) error {
	logger.Log.Println("In side UpdateMstUserDefaultSupportGroup")
	for i:=0;i<len(ids);i++{
		updatedID := ids[i]
		stmt, err := dbc.DB.Prepare(updateMstUserDefaultSupportGroup)
		defer stmt.Close()
		if err != nil {
			logger.Log.Println("UpdateMstUserDefaultSupportGroup Prepare Statement  Error", err)
			return err
		}
		_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Refuserid, tz.Groupid, updatedID)
		if err != nil {
			logger.Log.Println("UpdateMstUserDefaultSupportGroup Execute Statement  Error", err)
			return err
		}
	}
	return nil
}


