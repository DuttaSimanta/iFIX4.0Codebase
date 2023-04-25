package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var userroleactioninsert = "INSERT INTO mstclientroleuseraction(clientid,mstorgnhirarchyid,roleid,actionid,userid) VALUES (?,?,?,?,?)"
var userroleactiondelete = "UPDATE mstclientroleuseraction SET deleteflg=1 WHERE id=?"
var userroleactionduplicate = "SELECT count(id) total FROM mstclientroleuseraction WHERE clientid=? AND mstorgnhirarchyid=? AND roleid=? AND actionid=? AND userid=? AND deleteflg=0 AND activeflg=1"

//var userroleactioncountForClient = "SELECT count(a.id) total FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1"
var userroleactionupdate = "UPDATE mstclientroleuseraction SET clientid=?,mstorgnhirarchyid=?,roleid=?,actionid=?,userid=? WHERE id=?"

//var userroleactiongetForClient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,f.name as Username,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID,a.userid as RefuserID FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"

var userroleactioncount = "SELECT count(a.id) total FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1"
var userroleactionget = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,f.name as Username,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID,a.userid as RefuserID FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
var roleuserwiseaction = "SELECT a.id as ID,a.actionname as Actionname FROM mstaction a,mstclientroleuseraction b where a.id=b.actionid and b.clientid=? and b.mstorgnhirarchyid=? and b.roleid=? and b.activeflg=1 and b.deleteflg=0"
var deleteroleuserwise = "DELETE from mstclientroleuseraction where clientid=? and mstorgnhirarchyid=? and roleid=? and userid"

// Delete Action Role User wise
func (mdao DbConn) DeleteActionRoleUserWise(tz *entities.MstUserRoleActionEntity) error {
	logger.Log.Println("Query -->", deleterolewise)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID)
	stmt, err := mdao.DB.Prepare(deleteroleuserwise)
	defer stmt.Close()
	if err != nil {
		log.Print("deleteroleuserwise Prepare Statement  Error", err)
		logger.Log.Print("deleteroleuserwise Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID)
	if err != nil {
		log.Print("deleteroleuserwise Execute Statement  Error", err)
		logger.Log.Print("deleteroleuserwise Execute Statement  Error", err)
		return err
	}
	return nil
}

//CheckDuplicateUserRoleAction check duplicate record
func (mdao DbConn) CheckDuplicateUserRoleAction(tz *entities.MstUserRoleActionEntity) (entities.MstUserRoleActionEntities, error) {
	logger.Log.Println("Query -->", userroleactionduplicate)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID, tz.RefuserID)
	value := entities.MstUserRoleActionEntities{}
	err := mdao.DB.QueryRow(userroleactionduplicate, tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID, tz.RefuserID).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("CheckDuplicateUserRoleAction Get Statement Prepare Error", err)
		return value, err
	}
}

//InsertUserRoleActionData data insertd in mstclientuser table
func (mdao DbConn) InsertUserRoleActionData(tz *entities.MstUserRoleActionEntity) (int64, error) {
	logger.Log.Println("Query -->", userroleactioninsert)
	logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID, tz.RefuserID)
	stmt, err := mdao.DB.Prepare(userroleactioninsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertUserRoleActionData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID, tz.ActionID, tz.RefuserID)
	if err != nil {
		logger.Log.Print("InsertUserRoleActionData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//UpdatUserRoleActionData update mstclientuser table
func (mdao DbConn) UpdatUserRoleActionData(data *entities.MstUserRoleActionEntity) error {
	logger.Log.Println("Query -->", userroleactionupdate)
	logger.Log.Println("parameters -->", data.ClientID, data.MstorgnhirarchyID, data.RoleID, data.ActionID, data.RefuserID, data.ID)
	stmt, err := mdao.DB.Prepare(userroleactionupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("UpdatUserRoleActionData Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.ClientID, data.MstorgnhirarchyID, data.RoleID, data.ActionID, data.RefuserID, data.ID)
	if err != nil {
		logger.Log.Print("UpdatUserRoleActionData Execute Statement  Error", err)
		return err
	}
	return nil
}

//DeleteUserRoleActionData update mstclientuser table
func (mdao DbConn) DeleteUserRoleActionData(tz *entities.MstUserRoleActionEntity) error {
	logger.Log.Println("Query -->", userroleactiondelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := mdao.DB.Prepare(userroleactiondelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("DeleteUserRoleActionData Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("DeleteUserRoleActionData Execute Statement  Error", err)
		return err
	}
	return nil
}

//GetUserRoleActionCountForClient get user count with condition
func (mdao DbConn) GetUserRoleActionCountForClient(tz *entities.MstUserRoleActionEntity, OrgnTypeID int64) (entities.MstUserRoleActionEntities, error) {
	value := entities.MstUserRoleActionEntities{}
	var userroleactioncountForClient string
	var params []interface{}
	if OrgnTypeID == 1 {
		userroleactioncountForClient = "SELECT count(a.id) total FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1"
	} else if OrgnTypeID == 2 {
		userroleactioncountForClient = "SELECT count(a.id) total FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE a.clientid =? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1"
		params = append(params, tz.ClientID)
	} else {
		userroleactioncountForClient = "SELECT count(a.id) total FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1"
		params = append(params, tz.ClientID)
		params = append(params, tz.MstorgnhirarchyID)
	}
	err := mdao.DB.QueryRow(userroleactioncountForClient, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("GetUserRoleActionCountForClient Statement Prepare Error", err)
		return value, err
	}
}

//GetAllUserRoleActionForClient get user count with condition
func (mdao DbConn) GetAllUserRoleActionForClient(page *entities.MstUserRoleActionEntity, OrgnType int64) ([]entities.MstUserRoleActionEntity, error) {
	values := []entities.MstUserRoleActionEntity{}
	var userroleactiongetForClient string
	var params []interface{}
	if OrgnType == 1 {
		userroleactiongetForClient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,f.name as Username,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID,a.userid as RefuserID FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		userroleactiongetForClient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,f.name as Username,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID,a.userid as RefuserID FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE a.clientid =? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		userroleactiongetForClient = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, c.rolename as Rolename, e.actionname as Actionname,f.name as Username,a.clientid as ClientID,a.mstorgnhirarchyid as MstorgnhirarchyID,a.roleid as RoleID,a.actionid as ActionID,a.userid as RefuserID FROM mstclientroleuseraction a, mstclient b, mstclientuserrole c,mstorgnhierarchy d, mstaction e, mstclientuser f WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.roleid = c.id AND a.actionid = e.id AND a.userid = f.id AND a.deleteflg = 0 AND a.activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.MstorgnhirarchyID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	logger.Log.Println("userroleactiongetForClient >>>>>>>>>>>>>", userroleactiongetForClient)
	rows, err := mdao.DB.Query(userroleactiongetForClient, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllUserRoleActionForClient Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserRoleActionEntity{}
		rows.Scan(&value.ID, &value.Clientname, &value.Mstorgnhirarchyname, &value.Rolename, &value.Actionname, &value.Username, &value.ClientID, &value.MstorgnhirarchyID, &value.RoleID, &value.ActionID, &value.RefuserID)
		//logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	//logger.Log.Println("values -->", values)
	return values, nil
}

//GetUserRoleActionCount get user count with condition
func (mdao DbConn) GetUserRoleActionCount() (entities.MstUserRoleActionEntities, error) {
	logger.Log.Println("Query -->", userroleactioncount)
	//logger.Log.Println("parameters -->", tz.ClientID, tz.MstorgnhirarchyID)
	value := entities.MstUserRoleActionEntities{}
	err := mdao.DB.QueryRow(userroleactioncount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("GetUserRoleActionCount Statement Prepare Error", err)
		return value, err
	}
}

//GetAllUserRoleAction get user count with condition
func (mdao DbConn) GetAllUserRoleAction(page *entities.MstUserRoleActionEntity) ([]entities.MstUserRoleActionEntity, error) {
	logger.Log.Println("Query -->", userroleactionget)
	logger.Log.Println("parameters -->", page.Offset, page.Limit)
	values := []entities.MstUserRoleActionEntity{}
	rows, err := mdao.DB.Query(userroleactionget, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Print("GetAllUserRoleAction Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstUserRoleActionEntity{}
		rows.Scan(&value.ID, &value.Clientname, &value.Mstorgnhirarchyname, &value.Rolename, &value.Actionname, &value.Username, &value.ClientID, &value.MstorgnhirarchyID, &value.RoleID, &value.ActionID, &value.RefuserID)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}

// Select role user wise action
func (mdao DbConn) GetRoleUserWiseAction(tz *entities.MstUserRoleActionEntity) ([]entities.MstActionEntity, error) {
	log.Println("In side dao")
	values := []entities.MstActionEntity{}
	rows, err := mdao.DB.Query(roleuserwiseaction, tz.ClientID, tz.MstorgnhirarchyID, tz.RoleID)
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
