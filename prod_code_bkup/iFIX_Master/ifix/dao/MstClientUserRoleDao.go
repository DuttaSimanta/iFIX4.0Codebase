package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var roleinsert = "INSERT INTO mstclientuserrole(clientid,mstorgnhirarchyid,rolename,roledesc,issuperadmin) VALUES(?,?,?,?,?)"
var roledelete = "UPDATE mstclientuserrole SET deleteflg=1 WHERE id=?"
var updaterole = "UPDATE mstclientuserrole SET clientid=?,mstorgnhirarchyid=?,rolename=?,roledesc=?, issuperadmin=? WHERE id=?"
var duplicateRole = "SELECT count(id) total from mstclientuserrole where clientid=? AND mstorgnhirarchyid=? AND rolename=? AND deleteflg =0 AND id<>?"
var duplicateRoleforupdate = "SELECT count(id) total from mstclientuserrole where clientid=? AND mstorgnhirarchyid=? AND rolename=? AND issuperadmin =? AND id=?"

//var rolegetcount = "SELECT count(a.id) total FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0"
var rolegetallcount = "SELECT count(a.id) total FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0"

var roleget = "SELECT distinct a.id as ID,a.clientid,a.mstorgnhirarchyid,b.name as Clientname, d.name as Mstorgnhirarchyname, a.rolename as Rolename,a.roledesc as Roledesc,a.issuperadmin as Adminflag FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
var rolebyorgid = "SELECT id as ID,rolename as Rolename,roledesc,issuperadmin FROM mstclientuserrole WHERE clientid =? AND mstorgnhirarchyid = ? AND deleteflg = 0"

//CheckDuplicateCientRoles check duplicate record
var rolegetall = "SELECT a.id as ID,b.name as Clientname, d.name as Mstorgnhirarchyname, a.rolename as Rolename,a.roledesc as Roledesc,a.issuperadmin as Adminflag FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE  a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"

func (mdao DbConn) CheckDuplicateCientRoles(tz *entities.MstClientUserRoleEntity) (entities.MstClientUserRoleEntities, error) {
	logger.Log.Println("Update Query -->", duplicateRole)
	logger.Log.Println("parameters -->", tz.ClientID, tz.Mstorgnhirarchyid, tz.Rolename)
	value := entities.MstClientUserRoleEntities{}
	err := mdao.DB.QueryRow(duplicateRole, tz.ClientID, tz.Mstorgnhirarchyid, tz.Rolename, tz.ID).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("CheckDuplicateCientRoles Get Statement Prepare Error", err)
		return value, err
	}
}

//InsertRoleData data insertd in mstclientuserrole table
func (mdao DbConn) InsertRoleData(data *entities.MstClientUserRoleEntity) (int64, error) {
	logger.Log.Println("Insert query -->", roleinsert)
	logger.Log.Println("parameters -->", data.ClientID, data.Mstorgnhirarchyid, data.Rolename, data.Roledesc, data.Adminflag, data.UserID)
	stmt, err := mdao.DB.Prepare(roleinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertModule Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(data.ClientID, data.Mstorgnhirarchyid, data.Rolename, data.Roledesc, data.Adminflag)
	if err != nil {
		logger.Log.Print("InsertModule Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//UpdateRoleData update record number table
func (mdao DbConn) UpdateRoleData(data *entities.MstClientUserRoleEntity) error {
	logger.Log.Println("Update Query -->", updaterole)
	logger.Log.Println("parameters -->", data.ClientID, data.Mstorgnhirarchyid, data.Rolename, data.Roledesc, data.Adminflag, data.ID)
	stmt, err := mdao.DB.Prepare(updaterole)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Client Role Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.ClientID, data.Mstorgnhirarchyid, data.Rolename, data.Roledesc, data.Adminflag, data.ID)
	if err != nil {
		logger.Log.Print("Update Client Role Execute Statement  Error", err)
		return err
	}
	return nil
}

//DeleteRoleData update record number table
func (mdao DbConn) DeleteRoleData(tz *entities.MstClientUserRoleEntity) error {
	logger.Log.Println("Delete Query -->", roledelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := mdao.DB.Prepare(roledelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Delete Role Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("Delete Role Execute Statement  Error", err)
		return err
	}
	return nil
}

//GetRoleCount get role count with condition
// func (mdao DbConn) GetRoleCount(tz *entities.MstClientUserRoleEntity) (entities.MstClientUserRoleEntities, error) {
// 	logger.Log.Println("Count Query -->", rolegetcount)
// 	logger.Log.Println("parameters -->", tz.ClientID, tz.Mstorgnhirarchyid)
// 	value := entities.MstClientUserRoleEntities{}
// 	if tz.ClientID > 0 && tz.Mstorgnhirarchyid > 0 {
// 		err := mdao.DB.QueryRow(rolegetcount, tz.ClientID, tz.Mstorgnhirarchyid).Scan(&value.Total)
// 		switch err {
// 		case sql.ErrNoRows:
// 			value.Total = 0
// 			return value, nil
// 		case nil:
// 			return value, nil
// 		default:
// 			log.Print("GetRoleCount Get Statement Prepare Error", err)
// 			return value, err
// 		}
// 	} else {
// 		err := mdao.DB.QueryRow(rolegetallcount).Scan(&value.Total)
// 		switch err {
// 		case sql.ErrNoRows:
// 			value.Total = 0
// 			return value, nil
// 		case nil:
// 			return value, nil
// 		default:
// 			log.Print("GetRoleCount Get Statement Prepare Error", err)
// 			return value, err
// 		}
// 	}
// }

func (mdao DbConn) GetRoleCount(tz *entities.MstClientUserRoleEntity, OrgnTypeID int64) (entities.MstClientUserRoleEntities, error) {
	value := entities.MstClientUserRoleEntities{}
	var params []interface{}
	var rolegetcount string
	if OrgnTypeID == 1 {
		rolegetcount = "SELECT count(a.id) total FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0"
	} else if OrgnTypeID == 2 {
		rolegetcount = "SELECT count(a.id) total FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE a.clientid =? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0"
		params = append(params, tz.ClientID)
	} else {
		rolegetcount = "SELECT count(a.id) total FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0"
		params = append(params, tz.ClientID)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := mdao.DB.QueryRow(rolegetcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetRoleCount Get Statement Prepare Error", err)
		return value, err
	}

}

func (mdao DbConn) GetAllRoleData(page *entities.MstClientUserRoleEntity, OrgnType int64) ([]entities.MstClientUserRoleEntity, error) {
	values := []entities.MstClientUserRoleEntity{}
	var params []interface{}
	var roleget string
	if OrgnType == 1 {
		roleget = "SELECT distinct a.id as ID,a.clientid,a.mstorgnhirarchyid,b.name as Clientname, d.name as Mstorgnhirarchyname, a.rolename as Rolename,a.roledesc as Roledesc,a.issuperadmin as Adminflag FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		roleget = "SELECT distinct a.id as ID,a.clientid,a.mstorgnhirarchyid,b.name as Clientname, d.name as Mstorgnhirarchyname, a.rolename as Rolename,a.roledesc as Roledesc,a.issuperadmin as Adminflag FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE a.clientid =? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		roleget = "SELECT distinct a.id as ID,a.clientid,a.mstorgnhirarchyid,b.name as Clientname, d.name as Mstorgnhirarchyname, a.rolename as Rolename,a.roledesc as Roledesc,a.issuperadmin as Adminflag FROM mstclientuserrole a, mstclient b,mstorgnhierarchy d WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.clientid = b.id AND a.mstorgnhirarchyid = d.id AND a.deleteflg = 0 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := mdao.DB.Query(roleget, params...)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllRoleData Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientUserRoleEntity{}
		rows.Scan(&value.ID, &value.ClientID, &value.Mstorgnhirarchyid, &value.Clientname, &value.Mstorgnhirarchyname, &value.Rolename, &value.Roledesc, &value.Adminflag)
		values = append(values, value)
	}
	return values, nil

}

//Getrolebyorgid get role by orgid and clientid
func (mdao DbConn) Getrolebyorgid(page *entities.MstClientUserRoleEntity) ([]entities.MstClientRoleEntityResp, error) {
	logger.Log.Println("Getrolebyorgid Query -->", roleget)
	logger.Log.Println("parameters -->", page.ClientID, page.Mstorgnhirarchyid)
	values := []entities.MstClientRoleEntityResp{}
	rows, err := mdao.DB.Query(rolebyorgid, page.ClientID, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		log.Print("Getrolebyorgid Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientRoleEntityResp{}
		rows.Scan(&value.ID, &value.Rolename, &value.Roledesc, &value.Issuperadmin)
		values = append(values, value)
	}
	return values, nil
}
func (mdao DbConn) CheckDuplicateCientRolesforupdate(tz *entities.MstClientUserRoleEntity) (entities.MstClientUserRoleEntities, error) {
	logger.Log.Println("Update Query -->", duplicateRole)
	logger.Log.Println("parameters -->", tz.ClientID, tz.Mstorgnhirarchyid, tz.Rolename)
	value := entities.MstClientUserRoleEntities{}
	err := mdao.DB.QueryRow(duplicateRoleforupdate, tz.ClientID, tz.Mstorgnhirarchyid, tz.Rolename, tz.Adminflag, tz.ID).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("CheckDuplicateCientRoles Get Statement Prepare Error", err)
		return value, err
	}
}
