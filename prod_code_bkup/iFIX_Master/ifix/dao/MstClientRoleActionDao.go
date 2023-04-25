package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var roleactioninsert = "INSERT INTO mstclientroleaction(clientid,mstorgnhirarchyid,roleid,actionid) VALUES (?,?,?,?)"
var roleactiondelete = "UPDATE mstclientroleaction SET deleteflg=1 WHERE id=?"
var roleactionduplicate = "SELECT count(id) total FROM mstclientroleaction WHERE clientid=? AND mstorgnhirarchyid=? AND roleid=? AND actionid=? AND deleteflg=0 AND activeflg=1"

//var roleactioncountForClient = "SELECT count(id) total FROM mstclientroleaction WHERE clientid=? AND mstorgnhirarchyid=? AND deleteflg=0 AND activeflg=1"
var roleactionupdate = "UPDATE mstclientroleaction SET clientid=?,mstorgnhirarchyid=?,roleid=?,actionid=? WHERE id=?"

//var roleactiongetForClient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID FROM mstclientroleaction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"

var roleactioncount = "SELECT count(id) total FROM mstclientroleaction WHERE deleteflg=0 AND activeflg=1"
var roleactionget = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID FROM mstclientroleaction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
var mstactionget = "SELECT id as ID,actionname as Actionname FROM mstaction"
var rolewiseaction = "SELECT a.id as ID,a.actionname as Actionname FROM mstaction a,mstclientroleaction b where a.id=b.actionid and b.clientid=? and b.mstorgnhirarchyid=? and b.roleid=? and b.activeflg=1 and b.deleteflg=0"
var deleterolewise = "DELETE from mstclientroleaction where clientid=? and mstorgnhirarchyid=? and roleid=?"

//CheckDuplicateRoleAction check duplicate record
func (mdao DbConn) CheckDuplicateRoleAction(tz *entities.MstClientRoleActionEntity) (entities.MstClientRoleActionEntities, error) {
	logger.Log.Println("Query -->", roleactionduplicate)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID)
	value := entities.MstClientRoleActionEntities{}
	err := mdao.DB.QueryRow(roleactionduplicate, tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("CheckDuplicateRoleAction Get Statement Prepare Error", err)
		return value, err
	}
}

// Delete Action Role wise
func (mdao DbConn) DeleteActionRoleWise(tz *entities.MstClientRoleActionEntity) error {
	logger.Log.Println("Query -->", deleterolewise)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID)
	stmt, err := mdao.DB.Prepare(deleterolewise)
	defer stmt.Close()
	if err != nil {
		log.Print("DeleteActionRoleWise Prepare Statement  Error", err)
		logger.Log.Print("DeleteActionRoleWise Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID)
	if err != nil {
		log.Print("DeleteActionRoleWise Execute Statement  Error", err)
		logger.Log.Print("DeleteActionRoleWise Execute Statement  Error", err)
		return err
	}
	return nil
}

//InsertRoleActionData data insertd in mstclientuser table
func (mdao DbConn) InsertRoleActionData(tz *entities.MstClientRoleActionEntity) (int64, error) {
	logger.Log.Println("Query -->", roleactioninsert)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID)
	stmt, err := mdao.DB.Prepare(roleactioninsert)
	defer stmt.Close()
	if err != nil {
		log.Print("InsertMapRoleUserData Prepare Statement  Error", err)
		logger.Log.Print("InsertMapRoleUserData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID)
	if err != nil {
		log.Print("InsertRoleActionData Execute Statement  Error", err)
		logger.Log.Print("InsertRoleActionData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//UpdatRoleActionData update mstclientuser table
func (mdao DbConn) UpdatRoleActionData(data *entities.MstClientRoleActionEntity) error {
	logger.Log.Println("Query -->", roleactionupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.RoleID, data.ActionID, data.ID)
	stmt, err := mdao.DB.Prepare(roleactionupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("UpdatRoleActionData Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.ClientID, data.MstorgnhirarchyID, data.RoleID, data.ActionID, data.ID)
	if err != nil {
		logger.Log.Print("UpdatRoleActionData Execute Statement  Error", err)
		return err
	}
	return nil
}

//DeleteRoleActionData update mstclientuser table
func (mdao DbConn) DeleteRoleActionData(tz *entities.MstClientRoleActionEntity) error {
	logger.Log.Println("Query -->", roleactiondelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := mdao.DB.Prepare(roleactiondelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("DeleteRoleActionData Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("DeleteRoleActionData Execute Statement  Error", err)
		return err
	}
	return nil
}

//GetRoleActionCountForClient get user count with condition
func (mdao DbConn) GetRoleActionCountForClient(tz *entities.MstClientRoleActionEntity, OrgnTypeID int64) (entities.MstClientRoleActionEntities, error) {
	value := entities.MstClientRoleActionEntities{}
	var roleactioncountForClient string
	var params []interface{}
	if OrgnTypeID == 1 {
		roleactioncountForClient = "SELECT count(id) total FROM mstclientroleaction WHERE deleteflg=0 AND activeflg=1"
	} else if OrgnTypeID == 2 {
		roleactioncountForClient = "SELECT count(id) total FROM mstclientroleaction WHERE clientid=? AND deleteflg=0 AND activeflg=1"
		params = append(params, tz.ClientID)
	} else {
		roleactioncountForClient = "SELECT count(id) total FROM mstclientroleaction WHERE clientid=? AND mstorgnhirarchyid=? AND deleteflg=0 AND activeflg=1"
		params = append(params, tz.ClientID)
		params = append(params, tz.MstorgnhirarchyID)
	}
	err := mdao.DB.QueryRow(roleactioncountForClient, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("GetRoleActionCount Statement Prepare Error", err)
		return value, err
	}
}

//GetAllRoleActionForClient get user count with condition
func (mdao DbConn) GetAllRoleActionForClient(page *entities.MstClientRoleActionEntity, OrgnType int64) ([]entities.MstClientRoleActionEntity, error) {
	values := []entities.MstClientRoleActionEntity{}
	var roleactiongetForClient string
	var params []interface{}
	if OrgnType == 1 {
		roleactiongetForClient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID FROM mstclientroleaction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e WHERE a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		roleactiongetForClient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID FROM mstclientroleaction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e WHERE a.clientid =? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		roleactiongetForClient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID FROM mstclientroleaction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.MstorgnhirarchyID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := mdao.DB.Query(roleactiongetForClient, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllRoleAction Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientRoleActionEntity{}
		rows.Scan(&value.ID, &value.Clientname, &value.Mstorgnhirarchyname, &value.Rolename, &value.Actionname, &value.ClientID, &value.MstorgnhirarchyID, &value.RoleID, &value.ActionID)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}

//GetRoleActionCount get user count with condition
func (mdao DbConn) GetRoleActionCount() (entities.MstClientRoleActionEntities, error) {
	logger.Log.Println("Query -->", roleactioncount)
	//logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID)
	value := entities.MstClientRoleActionEntities{}
	err := mdao.DB.QueryRow(roleactioncount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("GetRoleActionCount Statement Prepare Error", err)
		return value, err
	}
}

//GetAllRoleAction get user count with condition
func (mdao DbConn) GetAllRoleAction(page *entities.MstClientRoleActionEntity) ([]entities.MstClientRoleActionEntity, error) {
	logger.Log.Println("Query -->", roleactionget)
	logger.Log.Println("parameters -->", page.Offset, page.Limit)
	values := []entities.MstClientRoleActionEntity{}
	rows, err := mdao.DB.Query(roleactionget, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllRoleAction Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientRoleActionEntity{}
		rows.Scan(&value.ID, &value.Clientname, &value.Mstorgnhirarchyname, &value.Rolename, &value.Actionname, &value.ClientID, &value.MstorgnhirarchyID, &value.RoleID, &value.ActionID)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}

//
////GetMstAction get Action data
func (mdao DbConn) GetMstAction() ([]entities.MstActionEntity, error) {
	logger.Log.Println("Query -->", mstactionget)
	values := []entities.MstActionEntity{}
	rows, err := mdao.DB.Query(mstactionget)
	defer rows.Close()
	if err != nil {
		log.Println("GetMstAction Statement Prepare Error", err)
		logger.Log.Print("GetMstAction Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstActionEntity{}
		rows.Scan(&value.ID, &value.Actionname)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}

// Select role wise action
func (mdao DbConn) GetRoleWiseAction(tz *entities.MstClientRoleActionEntity) ([]entities.MstActionEntity, error) {
	log.Println("In side dao")
	values := []entities.MstActionEntity{}
	rows, err := mdao.DB.Query(rolewiseaction, tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllModules Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstActionEntity{}
		rows.Scan(&value.ID, &value.Actionname)
		values = append(values, value)
	}
	return values, nil
}
