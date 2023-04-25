package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var maproleinsert = "INSERT INTO mapclientuserroleuser(clientid,mstorgnhirarchyid,roleid,userid) VALUES (?,?,?,?)"
var maproledelete = "UPDATE mapclientuserroleuser SET deleteflg=1 WHERE id=?"
var maproleduplicate = "SELECT count(id) total FROM mapclientuserroleuser WHERE clientid=? AND mstorgnhirarchyid=? AND roleid=? AND userid=? AND deleteflg=0 AND activeflg=1"
var maprolecount = "SELECT count(a.id) total FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND a.deleteflg = 0 AND c.deleteflg=0 AND e.deleteflag=0"
var maproleupdate = "UPDATE mapclientuserroleuser SET roleid=?,userid=? WHERE id=?"
var getmaproleuserall = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename,c.id as RoleID,e.name as Username,e.id as Refuserid FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND a.deleteflg = 0 AND c.deleteflg=0 AND e.deleteflag=0  ORDER BY a.id DESC LIMIT ?,?"

//var getmaproleuserforclient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename,a.roleid,a.userid as Refuserid, e.name as Username FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND e.deleteflag=0 and c.deleteflg=0 AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
//var maprolecountforclient = "SELECT count(a.id) total FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND e.deleteflag=0 and c.deleteflg=0 AND a.deleteflg = 0"

//CheckDuplicateMapRoleUser check duplicate record
func (mdao DbConn) CheckDuplicateMapRoleUser(tz *entities.MapClientUserRoleUserEntity) (entities.MapClientUserRoleUserEntities, error) {
	logger.Log.Println("maproleduplicate Query -->", maproleduplicate)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.Refuserid)
	value := entities.MapClientUserRoleUserEntities{}
	err := mdao.DB.QueryRow(maproleduplicate, tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.Refuserid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("maproleduplicate Get Statement Prepare Error", err)
		return value, err
	}
}

//InsertMapRoleUserData data insertd in mstclientuser table
func (mdao DbConn) InsertMapRoleUserData(tz *entities.MapClientUserRoleUserEntity) (int64, error) {
	logger.Log.Println("maproleinsert query -->", maproleinsert)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID)
	stmt, err := mdao.DB.Prepare(maproleinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertMapRoleUserData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.Refuserid)
	if err != nil {
		logger.Log.Print("InsertMapRoleUserData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}
func (mdao DbConn) InsertMapRoleUserDataTransaction(tz *entities.MapClientUserRoleUserEntity, tx *sql.Tx) (int64, error) {
	logger.Log.Println("maproleinsert query -->", maproleinsert)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID)
	stmt, err := tx.Prepare(maproleinsert)

	if err != nil {
		logger.Log.Print("InsertMapRoleUserDataTransaction Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.Refuserid)
	if err != nil {
		logger.Log.Print("InsertMapRoleUserDataTransaction Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//UpdateMapRoleUserData update mstclientuser table
func (mdao DbConn) UpdateMapRoleUserData(data *entities.MapClientUserRoleUserEntity) error {
	logger.Log.Println("UpdateMapRoleUserData Query -->", maproleupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.RoleID, data.ID)
	stmt, err := mdao.DB.Prepare(maproleupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Map Role User Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.RoleID, data.Refuserid, data.ID)
	if err != nil {
		logger.Log.Print("Update Map Role User Execute Statement  Error", err)
		return err
	}
	return nil
}

//DeleteMapRoleUserData update mstclientuser table
func (mdao DbConn) DeleteMapRoleUserData(tz *entities.MapClientUserRoleUserEntity) error {
	logger.Log.Println("maproledelete Query -->", maproledelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := mdao.DB.Prepare(maproledelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Delete Map Role User Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("Delete Map Role User Execute Statement  Error", err)
		return err
	}
	return nil
}

//GetMapRoleUserCount get user count with condition
func (mdao DbConn) GetMapRoleUserCount(tz *entities.MapClientUserRoleUserEntity, OrgnTypeID int64) (entities.MapClientUserRoleUserEntities, error) {
	logger.Log.Println("maprolecount query -->", maprolecount)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID)
	value := entities.MapClientUserRoleUserEntities{}
	var maprolecountforclient string
	var params []interface{}
	if OrgnTypeID == 1 {
		maprolecountforclient = "SELECT count(a.id) total FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND e.deleteflag=0 and c.deleteflg=0 AND a.deleteflg = 0"
	} else if OrgnTypeID == 2 {
		maprolecountforclient = "SELECT count(a.id) total FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE a.clientid =? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND e.deleteflag=0 and c.deleteflg=0 AND a.deleteflg = 0"
		params = append(params, tz.ClientID)
	} else {
		maprolecountforclient = "SELECT count(a.id) total FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND e.deleteflag=0 and c.deleteflg=0 AND a.deleteflg = 0"
		params = append(params, tz.ClientID)
		params = append(params, tz.MstorgnhirarchyID)
	}
	err := mdao.DB.QueryRow(maprolecountforclient, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("Get Map Role User Count Get Statement Prepare Error", err)
		logger.Log.Print("Get Map Role User Count Get Statement Prepare Error", err)
		return value, err
	}
}

//GetAllMapRoleUser get user count with condition
func (mdao DbConn) GetAllMapRoleUser(page *entities.MapClientUserRoleUserEntity, OrgnType int64) ([]entities.MapUserRoleEntityResp, error) {
	values := []entities.MapUserRoleEntityResp{}
	var getmaproleuserforclient string
	var params []interface{}
	if OrgnType == 1 {
		getmaproleuserforclient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename,a.roleid,a.userid as Refuserid, e.name as Username FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND e.deleteflag=0 and c.deleteflg=0 AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getmaproleuserforclient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename,a.roleid,a.userid as Refuserid, e.name as Username FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE a.clientid =? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND e.deleteflag=0 and c.deleteflg=0 AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getmaproleuserforclient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename,a.roleid,a.userid as Refuserid, e.name as Username FROM mapclientuserroleuser a, mstclient b, mstclientuserrole c,mstorgnhierarchy d,mstclientuser e WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.userid=e.id AND e.deleteflag=0 and c.deleteflg=0 AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.MstorgnhirarchyID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := mdao.DB.Query(getmaproleuserforclient, params...)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllMapRoleUser Get Statement Prepare Error", err)
		logger.Log.Print("GetAllMapRoleUser Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MapUserRoleEntityResp{}
		rows.Scan(&value.ID, &value.Clientname, &value.Mstorgnhirarchyname, &value.Rolename, &value.RoleID, &value.Refuserid, &value.Username)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}
