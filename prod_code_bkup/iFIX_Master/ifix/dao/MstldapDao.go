package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertMstldap = "INSERT INTO mstldap (clientid, mstorgnhirarchyid, servername, serverurl, binddn, basedn, password, filterdn, ori_certificate, chn_certificate) VALUES (?,?,?,?,?,?,?,?,?,?)"
var duplicateMstldap = "SELECT count(id) total FROM  mstldap WHERE clientid = ? AND mstorgnhirarchyid = ?  AND activeflg =0 AND deleteflg = 0 "

var getallMstldap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.servername as ServerName, a.serverurl as ServerUrl, a.binddn as Binddn,a.basedn as Basedn,a.password as Password,a.filterdn as Filterdn,coalesce(a.ori_certificate,'') as Ori_Certificate,coalesce(a.chn_certificate,'' ) as Chn_Certificate,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstldap a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"

// var getMstldapcount = "SELECT count(a.id) as total FROM mstldap a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
var updateMstldap = "UPDATE mstldap SET clientid=?,mstorgnhirarchyid = ?, servername = ?, serverurl = ?, binddn = ?,basedn=?,password=?,filterdn=?,ori_certificate=?, chn_certificate=? WHERE id = ? "
var deleteMstldap = "UPDATE mstldap SET deleteflg = '1' WHERE id = ? "
var updateMstldapcertificate = "UPDATE mstldap SET ori_certificate=?,chn_certificate=? where id=?"
var tabledesc = "SELECT COLUMN_NAME FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE  `TABLE_NAME`=?"

func (mdao DbConn) Gettabledetails(tz *entities.MstldapEntity) ([]string, error) {
	var values []string

	stmt, err := mdao.DB.Prepare(tabledesc)
	if err != nil {
		logger.Log.Print("Gettabledetails Statement Prepare Error", err)
		log.Print("Gettabledetails Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Tablename)
	if err != nil {
		logger.Log.Print("Gettabledetails Statement Execution Error", err)
		log.Print("Gettabledetails Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		var value string
		rows.Scan(&value)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateMstldap(tz *entities.MstldapEntity) (entities.MstldapEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstldap")
	value := entities.MstldapEntities{}
	err := dbc.DB.QueryRow(duplicateMstldap, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstldap Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstldap(tz *entities.MstldapEntity) (int64, error) {
	logger.Log.Println("In side InsertMstldap")
	logger.Log.Println("Query -->", insertMstldap)
	stmt, err := dbc.DB.Prepare(insertMstldap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstldap Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.ServerName, tz.ServerUrl, tz.Binddn, tz.Basedn, tz.Filterdn)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.ServerName, tz.ServerUrl, tz.Binddn, tz.Basedn, tz.Password, tz.Filterdn, tz.Ori_Certificate, tz.Chn_Certificate)
	if err != nil {
		logger.Log.Println("InsertMstldap Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetMstldap(tz *entities.MstldapEntity, OrgnType int64) ([]entities.MstldapEntity, error) {
	logger.Log.Println("In side GelAllMstldap")
	values := []entities.MstldapEntity{}
	var getMstldap string
	var params []interface{}
	if OrgnType == 1 {
		getMstldap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.servername as ServerName, a.serverurl as ServerUrl, a.binddn as Binddn,a.basedn as Basedn,a.password as Password,a.filterdn as Filterdn,coalesce(a.ori_certificate,'') as Ori_Certificate,coalesce(a.chn_certificate,'' ) as Chn_Certificate,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstldap a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstldap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.servername as ServerName, a.serverurl as ServerUrl, a.binddn as Binddn,a.basedn as Basedn,a.password as Password,a.filterdn as Filterdn,coalesce(a.ori_certificate,'') as Ori_Certificate,coalesce(a.chn_certificate,'' ) as Chn_Certificate,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstldap a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstldap = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.servername as ServerName, a.serverurl as ServerUrl, a.binddn as Binddn,a.basedn as Basedn,a.password as Password,a.filterdn as Filterdn,coalesce(a.ori_certificate,'') as Ori_Certificate,coalesce(a.chn_certificate,'' ) as Chn_Certificate,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstldap a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getMstldap, params...)

	// rows, err := dbc.DB.Query(getMstldap, tz.Clientid, tz.Mstorgnhirarchyid, tz.Offset, tz.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstldap Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstldapEntity{}
		err = rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.ServerName, &value.ServerUrl, &value.Binddn, &value.Basedn, &value.Password, &value.Filterdn, &value.Ori_Certificate, &value.Chn_Certificate, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		if err != nil {
			logger.Log.Println("GetAllMstldap Scan Error", err)
		}
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetAllMstldap(page *entities.MstldapEntity) ([]entities.MstldapEntity, error) {
	logger.Log.Println("In side GelAllMstldap")
	values := []entities.MstldapEntity{}
	rows, err := dbc.DB.Query(getallMstldap, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstldap Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstldapEntity{}
		err = rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.ServerName, &value.ServerUrl, &value.Binddn, &value.Basedn, &value.Password, &value.Filterdn, &value.Ori_Certificate, &value.Chn_Certificate, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		if err != nil {
			logger.Log.Println("GetAllMstldap Scan Error", err)
		}
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstldap(tz *entities.MstldapEntity) error {
	logger.Log.Println("In side UpdateMstldap")
	stmt, err := dbc.DB.Prepare(updateMstldap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstldap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.ServerName, tz.ServerUrl, tz.Binddn, tz.Basedn, tz.Password, tz.Filterdn, tz.Ori_Certificate, tz.Chn_Certificate, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstldap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstldap(tz *entities.MstldapEntity) error {
	logger.Log.Println("In side DeleteMstldap")
	stmt, err := dbc.DB.Prepare(deleteMstldap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstldap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstldap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstldapCount(tz *entities.MstldapEntity, OrgnTypeID int64) (entities.MstldapEntities, error) {
	logger.Log.Println("In side GetMstldapCount")
	value := entities.MstldapEntities{}
	var getMstldapcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		logger.Log.Println("In side pareamt")
		getMstldapcount = "SELECT count(a.id) as total FROM mstldap a,mstclient b,mstorgnhierarchy c WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
	} else if OrgnTypeID == 2 {
		getMstldapcount = "SELECT count(a.id) as total FROM mstldap a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
		params = append(params, tz.Clientid)
	} else {
		getMstldapcount = "SELECT count(a.id) as total FROM mstldap a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstldapcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstldapcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstldapCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) UpdateMstldapCertificate(tz *entities.MstldapEntity) error {
	logger.Log.Println("In side UpdateMstldap")
	stmt, err := dbc.DB.Prepare(updateMstldapcertificate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstldap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Ori_Certificate, tz.Chn_Certificate, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstldap Execute Statement  Error", err)
		return err
	}
	return nil
}
