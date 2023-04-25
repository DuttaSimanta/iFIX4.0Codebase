package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertClientsupportgroup = "INSERT INTO mstclientsupportgroup (clientid, mstorgnhirarchyid, supportgroupname, supportgrouplevelid, mstclienttimezoneid, reporttimezoneid, email,hascatalog) VALUES (?,?,?,?,?,?,?,?)"
var duplicateClientsupportgroup = "SELECT count(id) total FROM mstclientsupportgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND supportgroupname = ? AND supportgrouplevelid = ? AND mstclienttimezoneid=? AND reporttimezoneid=? AND email=? AND deleteflg = 0 AND activeflg=1"
var duplicateClientsupportgroupupdate = "SELECT count(id) total FROM mstclientsupportgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND supportgroupname = ? AND supportgrouplevelid = ? AND mstclienttimezoneid=? AND reporttimezoneid=? AND email=? AND hascatalog=? AND deleteflg = 0 AND activeflg=1"
var getClientsupportgroup = "SELECT a.id as Id,a.hascatalog, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.supportgroupname as Supportgroupname, a.supportgrouplevelid as Supportgrouplevelid, a.mstclienttimezoneid as Mstclienttimezoneid, a.reporttimezoneid as Reporttimezoneid, a.email as Email, a.activeflg as Activeflg,a.isworkflow as Isworkflow,b.name as Clientname,c.name as Mstorgnhirarchyname,d.name as Supportgrplevelname,e.zone_name as Timezonename,f.zone_name as Reporttimezonename  FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id ORDER BY a.id DESC LIMIT ?,?"
var getClientsupportgroupcount = "SELECT count(a.id) as total FROM mstclientsupportgroup a,mstclient b,mstorgnhierarchy c,supportgrouplevel d,zone e,zone f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.supportgrouplevelid=d.id and a.mstclienttimezoneid=e.zone_id and a.reporttimezoneid=f.zone_id"
var updateClientsupportgroup = "UPDATE mstclientsupportgroup SET mstorgnhirarchyid = ?, supportgroupname = ?, supportgrouplevelid = ?, mstclienttimezoneid = ?, reporttimezoneid = ?, email = ?,hascatalog=? WHERE id = ? "
var deleteClientsupportgroup = "UPDATE mstclientsupportgroup SET deleteflg = '1' WHERE id = ? "

//var getgroupbyorgid=" SELECT id,supportgroupname,supportgrouplevelid levelid from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and activeflg=1 and deleteflg = 0"
var getgroupbyorgid = "SELECT distinct a.grpid,b.name,a.supportgrouplevelid levelid from mstclientsupportgroup a,mstsupportgrp b where a.clientid=? and a.mstorgnhirarchyid=? and a.activeflg=1 and a.deleteflg = 0 and a.grpid = b.id and b.activeflg=1 and b.deleteflg=0 order by b.name"
var getprocessgroupbyorgid = "SELECT distinct a.grpid,b.name,a.supportgrouplevelid levelid from mstclientsupportgroup a,mstsupportgrp b where a.clientid=? and a.mstorgnhirarchyid=? and a.isworkflow='Y' and a.grpid = b.id and a.activeflg=1 and a.deleteflg = 0  and b.activeflg=1 and b.deleteflg=0"
var getgroupid="SELECT id from mstsupportgrp where clientid=? and name=? and activeflg=1 and deleteflg=0"
func (dbc DbConn) Searchgroupbyname(clientid int64,name string) ([]entities.ClientsupportgroupsingleEntity, error) {
	//log.Println("In side dao")
	values := []entities.ClientsupportgroupsingleEntity{}
	rows, err := dbc.DB.Query(getgroupid, clientid, name)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Searchgroupbyname Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupsingleEntity{}
		rows.Scan(&value.Id)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getgroupbyorgid(page *entities.ClientsupportgroupEntity) ([]entities.ClientsupportgroupsingleEntity, error) {
	logger.Log.Println("In side Getgroupbyorgid")
	logger.Log.Println(getgroupbyorgid)
	values := []entities.ClientsupportgroupsingleEntity{}
	rows, err := dbc.DB.Query(getgroupbyorgid, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getgroupbyorgid Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupsingleEntity{}
		rows.Scan(&value.Id, &value.Supportgroupname, &value.Levelid)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getprocessgroupbyorgid(page *entities.ClientsupportgroupEntity) ([]entities.ClientsupportgroupsingleEntity, error) {
	logger.Log.Println("In side Getgroupbyorgid")
	logger.Log.Println(getgroupbyorgid)
	values := []entities.ClientsupportgroupsingleEntity{}
	rows, err := dbc.DB.Query(getprocessgroupbyorgid, page.Clientid, page.Mstorgnhirarchyid)

	if err != nil {
		logger.Log.Println("Getprocessgroupbyorgid Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.ClientsupportgroupsingleEntity{}
		rows.Scan(&value.Id, &value.Supportgroupname, &value.Levelid)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getprocessgroupbyorgids(page *entities.ClientsupportgroupEntity) ([]entities.ClientsupportgroupsingleEntity, error) {
	values := []entities.ClientsupportgroupsingleEntity{}
	var getprocessgroupbyorgids = "SELECT distinct a.grpid,b.name,a.supportgrouplevelid levelid from mstclientsupportgroup a,mstsupportgrp b where a.clientid=? and a.mstorgnhirarchyid in ("+page.Mstorgnhirarchyids+") and a.isworkflow='Y' and a.grpid = b.id and a.activeflg=1 and a.deleteflg = 0  and b.activeflg=1 and b.deleteflg=0"
	rows, err := dbc.DB.Query(getprocessgroupbyorgids, page.Clientid)

	if err != nil {
		logger.Log.Println("Getprocessgroupbyorgids Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.ClientsupportgroupsingleEntity{}
		rows.Scan(&value.Id, &value.Groupname, &value.Levelid)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) CheckDuplicateClientsupportgroup(tz *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroup")
	value := entities.ClientsupportgroupEntities{}
	err := dbc.DB.QueryRow(duplicateClientsupportgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email).Scan(&value.Total)
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
}
func (dbc DbConn) CheckDuplicateClientsupportgroupupdate(tz *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroupupdate")
	value := entities.ClientsupportgroupEntities{}
	err := dbc.DB.QueryRow(duplicateClientsupportgroupupdate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Hascatalog).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateClientsupportgroupupdate Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertClientsupportgroup(tz *entities.ClientsupportgroupEntity) (int64, error) {
	logger.Log.Println("In side InsertClientsupportgroup")
	logger.Log.Println("Query -->", insertClientsupportgroup)
	stmt, err := dbc.DB.Prepare(insertClientsupportgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertClientsupportgroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Hascatalog)
	if err != nil {
		logger.Log.Println("InsertClientsupportgroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllClientsupportgroup(page *entities.ClientsupportgroupEntity) ([]entities.ClientsupportgroupEntity, error) {
	logger.Log.Println("In side GelAllClientsupportgroup")
	logger.Log.Println(getClientsupportgroup)
	values := []entities.ClientsupportgroupEntity{}
	rows, err := dbc.DB.Query(getClientsupportgroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientsupportgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupEntity{}
		rows.Scan(&value.Id, &value.Hascatalog, &value.Clientid, &value.Mstorgnhirarchyid, &value.Supportgroupname, &value.Supportgrouplevelid, &value.Mstclienttimezoneid, &value.Reporttimezoneid, &value.Email, &value.Activeflg, &value.Isworkflow, &value.Clientname, &value.Mstorgnhirarchyname, &value.Supportgrplevelname, &value.Timezonename, &value.Reporttimezonename)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateClientsupportgroup(tz *entities.ClientsupportgroupEntity) error {
	logger.Log.Println("In side UpdateClientsupportgroup")
	stmt, err := dbc.DB.Prepare(updateClientsupportgroup)
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroup Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Hascatalog, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteClientsupportgroup(tz *entities.ClientsupportgroupEntity) error {
	logger.Log.Println("In side DeleteClientsupportgroup")
	stmt, err := dbc.DB.Prepare(deleteClientsupportgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetClientsupportgroupCount(tz *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, error) {
	logger.Log.Println("In side GetClientsupportgroupCount")
	value := entities.ClientsupportgroupEntities{}
	err := dbc.DB.QueryRow(getClientsupportgroupcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetClientsupportgroupCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) CheckDuplicateClientsupportgroupforupdate(tz *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroup")
	value := entities.ClientsupportgroupEntities{}
	err := dbc.DB.QueryRow(duplicateClientsupportgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid).Scan(&value.Total)
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
}

//All method definations with transactions.........

func CheckDuplicateClientsupportgroupwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroup")
	value := entities.ClientsupportgroupEntities{}
	err := tx.QueryRow(duplicateClientsupportgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid).Scan(&value.Total)
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
}

func InsertClientsupportgroupwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) (int64, error) {
	logger.Log.Println("In side InsertClientsupportgroup")
	logger.Log.Println("Query -->", insertClientsupportgroup)
	stmt, err := tx.Prepare(insertClientsupportgroup)

	if err != nil {
		logger.Log.Println("InsertClientsupportgroup Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Hascatalog)
	if err != nil {
		logger.Log.Println("InsertClientsupportgroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func GetAllClientsupportgroupwithtransaction(tx *sql.Tx, page *entities.ClientsupportgroupEntity) ([]entities.ClientsupportgroupEntity, error) {
	logger.Log.Println("In side GelAllClientsupportgroup")
	logger.Log.Println(getClientsupportgroup)
	values := []entities.ClientsupportgroupEntity{}
	rows, err := tx.Query(getClientsupportgroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientsupportgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.ClientsupportgroupEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Supportgroupname, &value.Supportgrouplevelid, &value.Mstclienttimezoneid, &value.Reporttimezoneid, &value.Email, &value.Activeflg, &value.Isworkflow, &value.Clientname, &value.Mstorgnhirarchyname, &value.Supportgrplevelname, &value.Timezonename, &value.Reporttimezonename)
		values = append(values, value)
	}
	return values, nil
}

func UpdateClientsupportgroupwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) error {
	logger.Log.Println("In side UpdateClientsupportgroup")
	stmt, err := tx.Prepare(updateClientsupportgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid, tz.Mstclienttimezoneid, tz.Reporttimezoneid, tz.Email, tz.Hascatalog, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateClientsupportgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteClientsupportgroupwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) error {
	logger.Log.Println("In side DeleteClientsupportgroup")
	stmt, err := tx.Prepare(deleteClientsupportgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteClientsupportgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func GetClientsupportgroupCountwithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, error) {
	logger.Log.Println("In side GetClientsupportgroupCount")
	value := entities.ClientsupportgroupEntities{}
	err := tx.QueryRow(getClientsupportgroupcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetClientsupportgroupCount Get Statement Prepare Error", err)
		return value, err
	}
}

func CheckDuplicateClientsupportgroupforupdatewithtransaction(tx *sql.Tx, tz *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateClientsupportgroup")
	value := entities.ClientsupportgroupEntities{}
	err := tx.QueryRow(duplicateClientsupportgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Supportgroupname, tz.Supportgrouplevelid).Scan(&value.Total)
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
}
