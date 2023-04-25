package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstslaresponsiblesupportgroup = "INSERT INTO mstslaresponsiblesupportgroup (clientid, mstorgnhirarchyid, mstslafullfillmentcriteriaid, mstclientsupportgroupid, mstslaid) VALUES (?,?,?,?,?)"
var duplicateMstslaresponsiblesupportgroup = "SELECT count(id) total FROM  mstslaresponsiblesupportgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstslafullfillmentcriteriaid = ? AND mstclientsupportgroupid = ? AND mstslaid = ? AND deleteflg = 0 and activeflg=1"

//var getMstslaresponsiblesupportgroup = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstslafullfillmentcriteriaid as Mstslafullfillmentcriteriaid, mstclientsupportgroupid as Mstclientsupportgroupid, mstslaid as Mstslaid, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select slaname from mstclientsla where id=mstslaid and deleteflg =0 and activeflg=1) as Slaname,(select supportgroupname from mstclientsupportgroup where id=mstclientsupportgroupid and deleteflg =0 and activeflg=1 ) as Grpname FROM mstslaresponsiblesupportgroup WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
var getMstslaresponsiblesupportgroup = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstslafullfillmentcriteriaid as Mstslafullfillmentcriteriaid, mstclientsupportgroupid as Mstclientsupportgroupid, mstslaid as Mstslaid, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select slaname from mstclientsla where id=mstslaid and deleteflg =0 and activeflg=1) as Slaname,(select name from mstsupportgrp where id=mstclientsupportgroupid and deleteflg =0 and activeflg=1 ) as Grpname FROM mstslaresponsiblesupportgroup WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
var getMstslaresponsiblesupportgroupcount = "SELECT count(a.id) total FROM  mstslaresponsiblesupportgroup a,mstclient b,mstorgnhierarchy c,mstclientsla d,mstsupportgrp e WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND  a.deleteflg =0 AND a.activeflg=1 AND b.id=a.clientid AND c.id=a.mstorgnhirarchyid AND d.id=a.mstslaid and d.deleteflg =0 and d.activeflg=1 and e.id=a.mstclientsupportgroupid and e.deleteflg =0 and e.activeflg=1"
var updateMstslaresponsiblesupportgroup = "UPDATE mstslaresponsiblesupportgroup SET mstorgnhirarchyid = ?, mstslafullfillmentcriteriaid = ?, mstclientsupportgroupid = ?, mstslaid = ? WHERE id = ? "
var deleteMstslaresponsiblesupportgroup = "UPDATE mstslaresponsiblesupportgroup SET deleteflg = '1' WHERE id = ? "
var slanames = "select a.id as Id,a.slaname as Slaname from mstclientsla a,mstslafullfillmentcriteria b where a.id=b.slaid and b.supportgroupspecific=1 and a.clientid=? and a.mstorgnhirarchyid=? and a.deleteflg=0 and a.activeflg=1 and b.clientid=? and b.mstorgnhirarchyid=? and b.deleteflg=0 and b.activeflg=1"
var criteriaid = "select id from mstslafullfillmentcriteria where clientid=? and mstorgnhirarchyid=? and slaid=? and deleteflg=0 and activeflg=1"

func (dbc DbConn) CheckDuplicateMstslaresponsiblesupportgroup(tz *entities.MstslaresponsiblesupportgroupEntity) (entities.MstslaresponsiblesupportgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstslaresponsiblesupportgroup")
	value := entities.MstslaresponsiblesupportgroupEntities{}
	err := dbc.DB.QueryRow(duplicateMstslaresponsiblesupportgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslafullfillmentcriteriaid, tz.Mstclientsupportgroupid, tz.Mstslaid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstslaresponsiblesupportgroup Get Statement Prepare Error", err)
		return value, err
	}

}

func (dbc DbConn) InsertMstslaresponsiblesupportgroup(tz *entities.MstslaresponsiblesupportgroupEntity) (int64, error) {
	logger.Log.Println("In side InsertMstslaresponsiblesupportgroup")
	logger.Log.Println("Query -->", insertMstslaresponsiblesupportgroup)
	stmt, err := dbc.DB.Prepare(insertMstslaresponsiblesupportgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstslaresponsiblesupportgroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslafullfillmentcriteriaid, tz.Mstclientsupportgroupid, tz.Mstslaid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstslafullfillmentcriteriaid, tz.Mstclientsupportgroupid, tz.Mstslaid)
	if err != nil {
		logger.Log.Println("InsertMstslaresponsiblesupportgroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstslaresponsiblesupportgroup(page *entities.MstslaresponsiblesupportgroupEntity) ([]entities.MstslaresponsiblesupportgroupEntity, error) {
	logger.Log.Println("In side GelAllMstslaresponsiblesupportgroup")
	values := []entities.MstslaresponsiblesupportgroupEntity{}
	rows, err := dbc.DB.Query(getMstslaresponsiblesupportgroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstslaresponsiblesupportgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstslaresponsiblesupportgroupEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstslafullfillmentcriteriaid, &value.Mstclientsupportgroupid, &value.Mstslaid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Slaname, &value.Grpname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstslaresponsiblesupportgroup(tz *entities.MstslaresponsiblesupportgroupEntity) error {
	logger.Log.Println("In side UpdateMstslaresponsiblesupportgroup")
	stmt, err := dbc.DB.Prepare(updateMstslaresponsiblesupportgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstslaresponsiblesupportgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstslafullfillmentcriteriaid, tz.Mstclientsupportgroupid, tz.Mstslaid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstslaresponsiblesupportgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstslaresponsiblesupportgroup(tz *entities.MstslaresponsiblesupportgroupEntity) error {
	logger.Log.Println("In side DeleteMstslaresponsiblesupportgroup")
	stmt, err := dbc.DB.Prepare(deleteMstslaresponsiblesupportgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstslaresponsiblesupportgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstslaresponsiblesupportgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstslaresponsiblesupportgroupCount(tz *entities.MstslaresponsiblesupportgroupEntity) (entities.MstslaresponsiblesupportgroupEntities, error) {
	logger.Log.Println("In side GetMstslaresponsiblesupportgroupCount")
	value := entities.MstslaresponsiblesupportgroupEntities{}
	err := dbc.DB.QueryRow(getMstslaresponsiblesupportgroupcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstslaresponsiblesupportgroupCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllSlanames(page *entities.MstslaresponsiblesupportgroupEntity) ([]entities.Mstslanames, error) {
	logger.Log.Println("In side GetAllSlanames")
	values := []entities.Mstslanames{}
	logger.Log.Println(slanames, page.Clientid, page.Mstorgnhirarchyid, page.Clientid, page.Mstorgnhirarchyid)
	rows, err := dbc.DB.Query(slanames, page.Clientid, page.Mstorgnhirarchyid, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()

	if err != nil {
		logger.Log.Println("GetAllSlanames Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Mstslanames{}
		rows.Scan(&value.Id, &value.Slaname)
		values = append(values, value)
	}
	//logger.Log.Println(values)
	return values, nil
}

func (dbc DbConn) GetFullfillmentcriteriaid(page *entities.MstslaresponsiblesupportgroupEntity) (int64, error) {
	logger.Log.Println("In side Checkmatrixconfig")
	stmt, err := dbc.DB.Prepare(criteriaid)
	defer stmt.Close()
	var criteriaid int64
	if err != nil {
		logger.Log.Println("Checkmatrixconfig Prepare Statement  Error", err)
		return 0, err
	}
	rows, err := stmt.Query(page.Clientid, page.Mstorgnhirarchyid, page.Mstslaid)
	if err != nil {
		logger.Log.Println("GetFullfillmentcriteriaid Execute Statement  Error", err)
		return 0, err
	}
	for rows.Next() {
		if err := rows.Scan(&criteriaid); err != nil {
			logger.Log.Println("Error in fetching data")
		}

	}

	return criteriaid, nil

}
