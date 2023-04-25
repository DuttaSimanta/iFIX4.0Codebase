package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertClientsupportgroupnew = "INSERT INTO mstclientsupportgroup (clientid, mstorgnhirarchyid, grpid, supportgrouplevelid, mstclienttimezoneid, reporttimezoneid, email,isworkflow,hascatalog,ismanagement) VALUES (?,?,?,?,?,?,?,?,?,?)"
var duplicateClientsupportgroupnew = "SELECT count(id) total FROM mstclientsupportgroup WHERE mstorgnhirarchyid = ? AND grpid = ? AND deleteflg = 0 AND activeflg=1"
var duplicateClientsupportgroupnewupdate = "SELECT count(id) total FROM mstclientsupportgroup WHERE clientid=? AND mstorgnhirarchyid = ? AND grpid = ? AND id <> ?  AND deleteflg = 0 AND activeflg=1"
var getClientsupportgroupnew = "SELECT a.id as Id,a.hascatalog,a.isworkflow, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.grpid as Supportgroupid ,g.name as Supportgroupname, a.supportgrouplevelid as Supportgrouplevelid, a.mstclienttimezoneid as Mstclienttimezoneid, a.reporttimezoneid as Reporttimezoneid, a.email as Email, a.activeflg as Activeflg,a.isworkflow as Isworkflow,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgrplevelname,e.zone_name as Timezonename,f.zone_name as Reporttimezonename  FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id ORDER BY a.id DESC LIMIT ?,?"

//var getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id"
var updateClientsupportgroupnew = "UPDATE mstclientsupportgroup SET mstorgnhirarchyid = ?, grpid = ?, supportgrouplevelid = ?, mstclienttimezoneid = ?, reporttimezoneid = ?, email = ?,isworkflow=?,hascatalog=?,ismanagement=? WHERE id = ? "
var deleteClientsupportgroupnew = "UPDATE mstclientsupportgroup SET deleteflg = '1' WHERE id = ? "

var getgroupnewbyorgid = " SELECT a.id as Id,b.name as Supportgroupname ,a.supportgrouplevelid as Levelid from mstclientsupportgroup a,mstsupportgrp b where a.clientid=? and a.mstorgnhirarchyid=? and a.activeflg=1 and a.deleteflg =var  0 and a.grpid = b.id"

var getrowbygrpid = "SELECT a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.grpid as Supportgroupid, a.supportgrouplevelid as Supportgrouplevelid, a.mstclienttimezoneid as Mstclienttimezoneid, a.reporttimezoneid as Reporttimezoneid, a.email as Email, a.isworkflow as Isworkflow, a.hascatalog as Hascatalog,a.ismanagement from mstclientsupportgroup a where a.clientid=? and a.mstorgnhirarchyid=? and a.grpid=? and activeflg=1 and deleteflg=0"
var getClientsupportgroupbyclient = "SELECT a.id as Id,a.hascatalog, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.grpid as Supportgroupid ,g.name as Supportgroupname, a.supportgrouplevelid as Supportgrouplevelid, a.mstclienttimezoneid as Mstclienttimezoneid, a.reporttimezoneid as Reporttimezoneid, a.email as Email, a.activeflg as Activeflg,a.isworkflow as Isworkflow,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgrplevelname,e.zone_name as Timezonename,f.zone_name as Reporttimezonename  FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.clientid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id ORDER BY a.id DESC LIMIT ?,?"
var getClientsupportgroupbyclientcount = "SELECT count(a.id) as total FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.clientid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id"

func (dbc DbConn) Getgroupnewbyorgid(page *entities.ClientsupportgroupnewEntity) ([]entities.ClientsupportgroupnewsingleEntity, error) {
	logger.Log.Println("In side Getgroupnewbyorgid")
	logger.Log.Println(getgroupnewbyorgid)
	values := []entities.ClientsupportgroupnewsingleEntity{}
	rows, err := dbc.DB.Query(getgroupnewbyorgid, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getgroupnewbyorgid Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupnewsingleEntity{}
		rows.Scan(&value.Id, &value.Supportgroupname, &value.Levelid)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) CheckDuplicateClientsupportgroupnew(tz *entities.ClientsupportgroupnewEntity) (entities.ClientsupportgroupnewEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroupnew")
	value := entities.ClientsupportgroupnewEntities{}
	err := dbc.DB.QueryRow(duplicateClientsupportgroupnew, tz.Mstorgnhirarchyid, tz.Supportgroupid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientsupportgroupnew Get Statement Prepare Error", err)
		return value, err
	}
}

/*func (dbc DbConn) CheckDuplicateClientsupportgroupnewupdate(tz *entities.ClientsupportgroupnewEntity) (entities.ClientsupportgroupnewEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroupnew")
	value := entities.ClientsupportgroupnewEntities{}
	err := dbc.DB.QueryRow(duplicateClientsupportgroupnew, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email,tz.Isworkflow,tz.Hascatalog).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientsupportgroupnew Get Statement Prepare Error", err)
		return value, err
	}
}*/

func (dbc DbConn) InsertClientsupportgroupnew(tz *entities.ClientsupportgroupnewEntity) (int64, error) {
	logger.Log.Println("In side InsertClientsupportgroupnew")
	logger.Log.Println("Query -->", insertClientsupportgroupnew)
	stmt, err := dbc.DB.Prepare(insertClientsupportgroupnew)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertClientsupportgroupnew Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Isworkflow, tz.Hascatalog)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Isworkflow, tz.Hascatalog, tz.IsManagement)
	if err != nil {
		logger.Log.Println("InsertClientsupportgroupnew Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllClientsupportgroupnew(tz *entities.ClientsupportgroupnewEntity, OrgnType int64) ([]entities.ClientsupportgroupnewEntity, error) {
	values := []entities.ClientsupportgroupnewEntity{}
	var getClientsupportgroupnew string
	var params []interface{}
	if OrgnType == 1 {
		getClientsupportgroupnew = "SELECT a.id as Id,a.hascatalog,a.isworkflow, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.grpid as Supportgroupid ,g.name as Supportgroupname, a.supportgrouplevelid as Supportgrouplevelid, a.mstclienttimezoneid as Mstclienttimezoneid, a.reporttimezoneid as Reporttimezoneid, a.email as Email,a.ismanagement, a.activeflg as Activeflg,a.isworkflow as Isworkflow,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgrplevelname,e.zone_name as Timezonename,f.zone_name as Reporttimezonename  FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getClientsupportgroupnew = "SELECT a.id as Id,a.hascatalog,a.isworkflow, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.grpid as Supportgroupid ,g.name as Supportgroupname, a.supportgrouplevelid as Supportgrouplevelid, a.mstclienttimezoneid as Mstclienttimezoneid, a.reporttimezoneid as Reporttimezoneid, a.email as Email,a.ismanagement, a.activeflg as Activeflg,a.isworkflow as Isworkflow,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgrplevelname,e.zone_name as Timezonename,f.zone_name as Reporttimezonename  FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getClientsupportgroupnew = "SELECT a.id as Id,a.hascatalog,a.isworkflow, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.grpid as Supportgroupid ,g.name as Supportgroupname, a.supportgrouplevelid as Supportgrouplevelid, a.mstclienttimezoneid as Mstclienttimezoneid, a.reporttimezoneid as Reporttimezoneid, a.email as Email,a.ismanagement, a.activeflg as Activeflg,a.isworkflow as Isworkflow,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgrplevelname,e.zone_name as Timezonename,f.zone_name as Reporttimezonename  FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getClientsupportgroupnew, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientsupportgroupnew Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupnewEntity{}
		rows.Scan(&value.Id, &value.Hascatalog, &value.Isworkflow, &value.Clientid, &value.Mstorgnhirarchyid, &value.Supportgroupid, &value.Supportgroupname, &value.Supportgrouplevelid, &value.Mstclienttimezoneid, &value.Reporttimezoneid, &value.Email, &value.IsManagement, &value.Activeflg, &value.Isworkflow, &value.Clientname, &value.Mstorgnhirarchyname, &value.Supportgrplevelname, &value.Timezonename, &value.Reporttimezonename)

		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateClientsupportgroupnew(tz *entities.ClientsupportgroupnewEntity) error {
	logger.Log.Println("In side UpdateClientsupportgroupnew")
	stmt, err := dbc.DB.Prepare(updateClientsupportgroupnew)
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroupnew Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Hascatalog, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroupnew Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteClientsupportgroupnew(tz *entities.ClientsupportgroupnewEntity) error {
	logger.Log.Println("In side DeleteClientsupportgroupnew")
	stmt, err := dbc.DB.Prepare(deleteClientsupportgroupnew)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroupnew Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroupnew Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetClientsupportgroupnewCount(tz *entities.ClientsupportgroupnewEntity, OrgnTypeID int64) (entities.ClientsupportgroupnewEntities, error) {
	logger.Log.Println("In side GetClientsupportgroupnewCount")
	value := entities.ClientsupportgroupnewEntities{}
	var getClientsupportgroupnewcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE  a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id"
	} else if OrgnTypeID == 2 {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id"
		params = append(params, tz.Clientid)
	} else {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f,mstsupportgrp g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id and a.grpid=g.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getClientsupportgroupnewcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetClientsupportgroupnewCount Get Statement Prepare Error", err)
		return value, err
	}
}

/*func (dbc DbConn) CheckDuplicateClientsupportgroupforupdate(tz *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroup")
	value := entities.ClientsupportgroupEntities{}
	err := dbc.DB.QueryRow(duplicateClientsupportgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientsupportgroup Get Statement Prepare Error", err)
		return value, err
	}
}*/

//All method definations with transactions.........

func CheckDuplicateClientsupportgroupnewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) (entities.ClientsupportgroupnewEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroupnew")
	value := entities.ClientsupportgroupnewEntities{}
	err := tx.QueryRow(duplicateClientsupportgroupnew, tz.Mstorgnhirarchyid, tz.Supportgroupid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientsupportgroupnew Get Statement Prepare Error", err)
		return value, err
	}
}

func InsertClientsupportgroupnewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) (int64, error) {
	logger.Log.Println("In side InsertClientsupportgroupnew")
	logger.Log.Println("Query -->", insertClientsupportgroupnew)
	stmt, err := tx.Prepare(insertClientsupportgroupnew)

	if err != nil {
		logger.Log.Println("InsertClientsupportgroupnew Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Isworkflow, tz.Hascatalog, tz.IsManagement)
	if err != nil {
		logger.Log.Println("InsertClientsupportgroupnew Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func UpdateClientsupportgroupnewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) error {
	logger.Log.Println("In side UpdateClientsupportgroupnew")
	stmt, err := tx.Prepare(updateClientsupportgroupnew)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroupnew Prepare Statement  Error", err)
		return err
	}
	logger.Log.Println(tz.Isworkflow)
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Isworkflow, tz.Hascatalog, tz.IsManagement, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroupnew Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteClientsupportgroupnewwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) error {
	logger.Log.Println("In side DeleteClientsupportgroupnew")
	stmt, err := tx.Prepare(deleteClientsupportgroupnew)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroupnew Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroupnew Execute Statement  Error", err)
		return err
	}
	return nil
}

func CheckDuplicateClientsupportgroupnewforupdatewithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupnewEntity) (entities.ClientsupportgroupnewEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroupnew")
	value := entities.ClientsupportgroupnewEntities{}
	err := tx.QueryRow(duplicateClientsupportgroupnewupdate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupid, tz.Id).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientsupportgroupnew Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) Getrowbygrpid(page *entities.ClientsupportgroupnewEntity) (*entities.ClientsupportgroupnewEntity, error) {
	logger.Log.Println("In side Getgroupnewbyorgid")
	logger.Log.Println(getrowbygrpid)
	value := entities.ClientsupportgroupnewEntity{}
	rows, err := dbc.DB.Query(getrowbygrpid, page.Clientid, page.Mstorgnhirarchyid, page.Supportgroupid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getgroupnewbyorgid Get Statement Prepare Error", err)
		return &value, err
	}
	for rows.Next() {
		//value := entities.ClientsupportgroupnewEntity{}
		rows.Scan(&value.Clientid, &value.Mstorgnhirarchyid, &value.Supportgroupid, &value.Supportgrouplevelid, &value.Mstclienttimezoneid, &value.Reporttimezoneid, &value.Email, &value.Isworkflow, &value.Hascatalog, &value.IsManagement)
		//values = append(values, value)
		break

	}
	logger.Log.Println("In side Getgroupnewbyorgid", value)

	return &value, nil
}

func (dbc DbConn) GetAllClientsupportgroupbyclient(page *entities.ClientsupportgroupnewEntity) ([]entities.ClientsupportgroupnewEntity, error) {
	logger.Log.Println("In side GelAllClientsupportgroupnew")
	logger.Log.Println(getClientsupportgroupnew)
	values := []entities.ClientsupportgroupnewEntity{}
	rows, err := dbc.DB.Query(getClientsupportgroupbyclient, page.Clientid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientsupportgroupnew Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupnewEntity{}
		rows.Scan(&value.Id, &value.Hascatalog, &value.Clientid, &value.Mstorgnhirarchyid, &value.Supportgroupid, &value.Supportgroupname, &value.Supportgrouplevelid, &value.Mstclienttimezoneid, &value.Reporttimezoneid, &value.Email, &value.Activeflg, &value.Isworkflow, &value.Clientname, &value.Mstorgnhirarchyname, &value.Supportgrplevelname, &value.Timezonename, &value.Reporttimezonename)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetAllClientsupportgroupbyclientcount(tz *entities.ClientsupportgroupnewEntity) (entities.ClientsupportgroupnewEntities, error) {
	logger.Log.Println("In side GetClientsupportgroupnewCount")
	value := entities.ClientsupportgroupnewEntities{}
	err := dbc.DB.QueryRow(getClientsupportgroupbyclientcount, tz.Clientid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetClientsupportgroupnewCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) Getsupportgroupbyorg(page *entities.ClientsupportgroupnewEntity, ids string) ([]entities.GetsupportgroupbyorgEntity, error) {
	logger.Log.Println("In side GelAllClientsupportgroupnew")
	getsupportgroupbyorg := "SELECT   a.id as id, a.clientid as clientid ,a.mstorgnhirarchyid as mstorgnhirarchyid,a.grpid as groupid,b.name as clientname,c.name as mstorgnhirarchyname,d.name as Groupname FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,mstsupportgrp d where a.clientid=? and a.mstorgnhirarchyid in (" + ids + ") and a.clientid=b.id and a.mstorgnhirarchyid=c.id and a.grpid=d.id and a.activeflg=1 and a.deleteflg=0 order by d.name"
	logger.Log.Println(getClientsupportgroupnew)
	values := []entities.GetsupportgroupbyorgEntity{}
	rows, err := dbc.DB.Query(getsupportgroupbyorg, page.Clientid)
	logger.Log.Println("working")
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientsupportgroupnew Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.GetsupportgroupbyorgEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Groupid, &value.Clientname, &value.Mstorgnhirarchyname, &value.Groupname)
		values = append(values, value)
	}
	return values, nil
}
