package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertMapfunctionalitywithgroup = "INSERT INTO mapfunctionalitywithgroup (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, mstfunctionailyid, diffid, groupid,recorddifftypestatusid,recorddiffstatusid) VALUES (?,?,?,?,?,?,?,?,?)"
var insertMapfunctionalitywithgroupwithuser = "INSERT INTO mapfunctionalitywithgroup (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, mstfunctionailyid, diffid, groupid,userid) VALUES (?,?,?,?,?,?,?,?)"
var duplicateMapfunctionalitywithgroup = "SELECT count(id) total FROM  mapfunctionalitywithgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND mstfunctionailyid = ? AND diffid = ? AND groupid = ? AND deleteflg = 0 and activeflg=1 and recorddifftypestatusid=? and recorddiffstatusid=?"
var duplicateMapfunctionalitywithgroupwithuser = "SELECT count(id) total FROM  mapfunctionalitywithgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND mstfunctionailyid = ? AND diffid = ? AND groupid = ? AND userid = ? AND deleteflg = 0 and activeflg=1"

//var getMapfunctionalitywithgroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mstfunctionailyid as Mstfunctionailyid, a.diffid as Diffid, a.groupid as Groupid, coalesce(a.userid,'0') as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifftypname,e.name as Recorddiffname,f.name as Mstfunctionailyname,g.description as Diffname,coalesce(h.name,'') as Refusername,k.name FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mapfunctionality g,mstsupportgrp k,mapfunctionalitywithgroup a left join mstclientuser h on h.id=a.userid WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.funcdescid AND g.clientid=? AND g.mstorgnhirarchyid=? AND d.deleteflg = 0 AND d.activeflg = 1 AND e.deleteflg = 0 AND e.activeflg = 1 AND g.deleteflg = 0 AND g.activeflg = 1 AND g.funcid = f.id AND k.id=a.groupid ORDER BY a.id DESC LIMIT ?,?"
var getMapfunctionalitywithgroup = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mstfunctionailyid as Mstfunctionailyid, a.diffid as Diffid, a.groupid as Groupid, coalesce(a.userid,'0') as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifftypname,e.name as Recorddiffname,f.name as Mstfunctionailyname,g.description as Diffname,coalesce(h.name,'') as Refusername,k.name,coalesce(a.recorddifftypestatusid,'') Recorddifftypestatusid,coalesce(a.recorddiffstatusid,'') Recorddiffstatusid ,coalesce(l.typename,'') as Recorddifftypestatusname,coalesce(m.name,'') as Recorddiffstatusname FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mapfunctionality g,mstsupportgrp k,mapfunctionalitywithgroup a left join mstclientuser h on h.id=a.userid left join mstrecorddifferentiationtype l on a.recorddifftypestatusid=l.id left join mstrecorddifferentiation m on a.recorddiffstatusid = m.id WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.funcdescid AND g.clientid=2 AND g.mstorgnhirarchyid=2 AND d.deleteflg = 0 AND d.activeflg = 1 AND e.deleteflg = 0 AND e.activeflg = 1 AND g.deleteflg = 0 AND g.activeflg = 1 AND g.funcid = f.id AND k.id=a.groupid ORDER BY a.id DESC LIMIT ?,?"
var getMapfunctionalitywithgroupwithterm = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mstfunctionailyid as Mstfunctionailyid, a.diffid as Diffid, a.groupid as Groupid, a.userid as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifftypname,e.name as Recorddiffname,f.name as Mstfunctionailyname,g.termname as Diffname,h.name as Refusername FROM mapfunctionalitywithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mstrecordterms g,mstclientuser h WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.mstfunctionailyid=?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.id and a.userid=h.id ORDER BY a.id DESC LIMIT ?,?"

//var getMapfunctionalitywithgroupcount = "SELECT count(id) total FROM  mapfunctionalitywithgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var getMapfunctionalitywithgroupcount = "SELECT count(a.id) as total FROM mapfunctionalitywithgroup a left join mstclientuser h on a.userid = h.id and h.activeflag=1 and h.deleteflag=0,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mapfunctionality g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.id "
var getMapfunctionalitywithgroupcountwithterm = "SELECT count(a.id) as total FROM mapfunctionalitywithgroup a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mstrecordterms g,mstclientuser h WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.mstfunctionailyid=?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.id and a.userid=h.id"
var updateMapfunctionalitywithgroup = "UPDATE mapfunctionalitywithgroup SET mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ?, mstfunctionailyid = ?, diffid = ?, groupid = ?,recorddifftypestatusid=?,recorddiffstatusid=? WHERE id = ? "
var updateMapfunctionalitywithgroupwithuser = "UPDATE mapfunctionalitywithgroup SET mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ?, mstfunctionailyid = ?, diffid = ?, groupid = ?,userid=? WHERE id = ? "
var deleteMapfunctionalitywithgroup = "UPDATE mapfunctionalitywithgroup SET deleteflg = '1' WHERE id = ? "
var getsupportgrporganization = "select distinct b.id,b.name,a.isworkflow from mstclientsupportgroup a,mstorgnhierarchy b where a.clientid=? and a.deleteflg=0 and a.activeflg=1 and a.mstorgnhirarchyid=b.id"

func (dbc DbConn) CheckDuplicateMapfunctionalitywithgroup(tz *entities.MapfunctionalitywithgroupEntity, diffid int64, grpid int64, recorddiffstatusid int64) (entities.MapfunctionalitywithgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapfunctionalitywithgroup")
	value := entities.MapfunctionalitywithgroupEntities{}
	err := dbc.DB.QueryRow(duplicateMapfunctionalitywithgroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, diffid, grpid, tz.Recorddifftypestatusid, recorddiffstatusid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapfunctionalitywithgroup Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) CheckDuplicateMapfunctionalitywithgroupwithoutStatus(tz *entities.MapfunctionalitywithgroupEntity, diffid int64, grpid int64) (entities.MapfunctionalitywithgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapfunctionalitywithgroup")
	value := entities.MapfunctionalitywithgroupEntities{}
	var query = "SELECT count(id) total FROM  mapfunctionalitywithgroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND mstfunctionailyid = ? AND diffid = ? AND groupid = ? AND deleteflg = 0 and activeflg=1"
	err := dbc.DB.QueryRow(query, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, diffid, grpid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapfunctionalitywithgroup Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) CheckDuplicateMapfunctionalitywithgroupwithuser(tz *entities.MapfunctionalitywithgroupEntity, diffid int64, userid int64) (entities.MapfunctionalitywithgroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateMapfunctionalitywithgroup")
	value := entities.MapfunctionalitywithgroupEntities{}
	err := dbc.DB.QueryRow(duplicateMapfunctionalitywithgroupwithuser, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, diffid, tz.Groupid, userid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMapfunctionalitywithgroup Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMapfunctionalitywithgroup(tz *entities.MapfunctionalitywithgroupEntity, diffid int64, grpid int64, recorddiffstatusid int64) (int64, error) {
	logger.Log.Println("In side InsertMapfunctionalitywithgroup")
	logger.Log.Println("Query -->", insertMapfunctionalitywithgroup)
	stmt, err := dbc.DB.Prepare(insertMapfunctionalitywithgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapfunctionalitywithgroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, tz.Diffid, tz.Groupid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, diffid, grpid, tz.Recorddifftypestatusid, recorddiffstatusid)
	if err != nil {
		logger.Log.Println("InsertMapfunctionalitywithgroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) InsertMapfunctionalitywithgroupwithoutstatus(tz *entities.MapfunctionalitywithgroupEntity, diffid int64, grpid int64) (int64, error) {
	logger.Log.Println("In side InsertMapfunctionalitywithgroup")
	logger.Log.Println("Query -->", insertMapfunctionalitywithgroup)
	var query = "INSERT INTO mapfunctionalitywithgroup (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, mstfunctionailyid, diffid, groupid) VALUES (?,?,?,?,?,?,?)"
	stmt, err := dbc.DB.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapfunctionalitywithgroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, tz.Diffid, tz.Groupid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, diffid, grpid)
	if err != nil {
		logger.Log.Println("InsertMapfunctionalitywithgroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) InsertMapfunctionalitywithgroupwithuser(tz *entities.MapfunctionalitywithgroupEntity, diffid int64, userid int64) (int64, error) {
	logger.Log.Println("In side InsertMapfunctionalitywithgroup")
	logger.Log.Println("Query -->", insertMapfunctionalitywithgroup)
	stmt, err := dbc.DB.Prepare(insertMapfunctionalitywithgroupwithuser)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMapfunctionalitywithgroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, tz.Diffid, tz.Groupid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, diffid, tz.Groupid, userid)
	if err != nil {
		logger.Log.Println("InsertMapfunctionalitywithgroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMapfunctionalitywithgroup(tz *entities.MapfunctionalitywithgroupEntity, OrgnType int64) ([]entities.MapfunctionalitywithgroupResponseEntity, error) {
	logger.Log.Println("In side GelAllMapfunctionalitywithgroup")
	values := []entities.MapfunctionalitywithgroupResponseEntity{}
	var query string
	// if page.Mstfunctionailyid == 2 {
	// 	query = getMapfunctionalitywithgroupwithterm
	// } else {
	// 	query = getMapfunctionalitywithgroup
	// }
	logger.Log.Println("Fetch query --->", query)
	log.Print("OrgnType:", OrgnType)
	var getMststate string
	var params []interface{}
	if OrgnType == 1 {
		getMststate = "SELECT distinct a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mstfunctionailyid as Mstfunctionailyid, a.diffid as Diffid, a.groupid as Groupid, coalesce(a.userid,'0') as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,COALESCE(d.typename,'') as Recorddifftypname,COALESCE(e.name ,'') as Recorddiffname,f.name as Mstfunctionailyname,g.description as Diffname,coalesce(h.name,'') as Refusername,k.name,coalesce(a.recorddifftypestatusid,'') Recorddifftypestatusid,coalesce(a.recorddiffstatusid,'') Recorddiffstatusid ,coalesce(l.typename,'') as Recorddifftypestatusname,coalesce(m.name,'') as Recorddiffstatusname FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mapfunctionality g,mstsupportgrp k,mapfunctionalitywithgroup a left join mstclientuser h on h.id=a.userid left join mstrecorddifferentiationtype l on a.recorddifftypestatusid=l.id left join mstrecorddifferentiation m on a.recorddiffstatusid = m.id WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.funcdescid AND  d.deleteflg = 0 AND d.activeflg = 1 AND e.deleteflg = 0 AND e.activeflg = 1 AND g.deleteflg = 0 AND g.activeflg = 1 AND g.funcid = f.id AND k.id=a.groupid ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMststate = "SELECT distinct a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mstfunctionailyid as Mstfunctionailyid, a.diffid as Diffid, a.groupid as Groupid, coalesce(a.userid,'0') as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,COALESCE(d.typename,'') as Recorddifftypname,COALESCE(e.name ,'') as Recorddiffname,f.name as Mstfunctionailyname,g.description as Diffname,coalesce(h.name,'') as Refusername,k.name,coalesce(a.recorddifftypestatusid,'') Recorddifftypestatusid,coalesce(a.recorddiffstatusid,'') Recorddiffstatusid ,coalesce(l.typename,'') as Recorddifftypestatusname,coalesce(m.name,'') as Recorddiffstatusname FROM mstclient b,mstorgnhierarchy c,mstfunctionality f,mapfunctionality g,mstsupportgrp k,mapfunctionalitywithgroup a left join mstclientuser h on h.id=a.userid left join mstrecorddifferentiationtype l on a.recorddifftypestatusid=l.id left join mstrecorddifferentiation m on a.recorddiffstatusid = m.id LEFT JOIN mstrecorddifferentiation e ON a.recorddiffid = e.id AND e.deleteflg = 0 AND e.activeflg = 1 LEFT JOIN mstrecorddifferentiationtype d ON a.recorddifftypeid = d.id AND d.deleteflg = 0 AND d.activeflg = 1 WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstfunctionailyid=f.id and a.diffid=g.funcdescid AND g.clientid=?  AND g.deleteflg = 0 AND g.activeflg = 1 AND g.funcid = f.id AND k.id=a.groupid ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMststate = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as Recorddifftypeid, a.recorddiffid as Recorddiffid, a.mstfunctionailyid as Mstfunctionailyid, a.diffid as Diffid, a.groupid as Groupid, coalesce(a.userid,'0') as Refuserid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,COALESCE(d.typename,'') as Recorddifftypname, COALESCE(e.name ,'') as Recorddiffname,f.name as Mstfunctionailyname,g.description as Diffname,coalesce(h.name,'') as Refusername,k.name,coalesce(a.recorddifftypestatusid,'') Recorddifftypestatusid,coalesce(a.recorddiffstatusid,'') Recorddiffstatusid ,coalesce(l.typename,'') as Recorddifftypestatusname,coalesce(m.name,'') as Recorddiffstatusname FROM mstclient b,mstorgnhierarchy c,mstfunctionality f,mapfunctionality g,mstsupportgrp k,mapfunctionalitywithgroup a left join mstclientuser h on h.id=a.userid left join mstrecorddifferentiationtype l on a.recorddifftypestatusid=l.id left join mstrecorddifferentiation m on a.recorddiffstatusid = m.id LEFT JOIN mstrecorddifferentiation e ON a.recorddiffid = e.id AND e.deleteflg = 0 AND e.activeflg = 1 LEFT JOIN mstrecorddifferentiationtype d ON a.recorddifftypeid = d.id AND d.deleteflg = 0 AND d.activeflg = 1 WHERE a.clientid = ? AND a.mstorgnhirarchyid=? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.mstfunctionailyid=f.id and a.diffid=g.funcdescid AND g.clientid=? AND g.mstorgnhirarchyid=?  AND g.deleteflg = 0 AND g.activeflg = 1 AND g.funcid = f.id AND k.id=a.groupid ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getMststate, params...)
	// rows, err := dbc.DB.Query(getMapfunctionalitywithgroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit) //page.Clientid, page.Mstorgnhirarchyid,
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMapfunctionalitywithgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MapfunctionalitywithgroupResponseEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Mstfunctionailyid, &value.Diffid, &value.Groupid, &value.Refuserid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifftypname, &value.Recorddiffname, &value.Mstfunctionailyname, &value.Diffname, &value.Refusername, &value.Grpname, &value.Recorddifftypestatusid, &value.Recorddiffstatusid, &value.Recorddifftypestatusname, &value.Recorddiffstatusname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMapfunctionalitywithgroup(tz *entities.MapfunctionalitywithgroupEntity, diffid int64, grpid int64, recorddiffstatusid int64) error {
	logger.Log.Println("In side UpdateMapfunctionalitywithgroup")
	stmt, err := dbc.DB.Prepare(updateMapfunctionalitywithgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMapfunctionalitywithgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, diffid, grpid, tz.Recorddifftypestatusid, recorddiffstatusid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMapfunctionalitywithgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) UpdateMapfunctionalitywithgroupwithuser(tz *entities.MapfunctionalitywithgroupEntity, diffid int64, userid int64) error {
	logger.Log.Println("In side UpdateMapfunctionalitywithgroup")
	stmt, err := dbc.DB.Prepare(updateMapfunctionalitywithgroupwithuser)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMapfunctionalitywithgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Mstfunctionailyid, diffid, tz.Groupid, userid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMapfunctionalitywithgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMapfunctionalitywithgroup(tz *entities.MapfunctionalitywithgroupEntity) error {
	logger.Log.Println("In side DeleteMapfunctionalitywithgroup")
	stmt, err := dbc.DB.Prepare(deleteMapfunctionalitywithgroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapfunctionalitywithgroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMapfunctionalitywithgroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMapfunctionalitywithgroupCount(tz *entities.MapfunctionalitywithgroupEntity, OrgnTypeID int64) (entities.MapfunctionalitywithgroupEntities, error) {
	logger.Log.Println("In side GetMapfunctionalitywithgroupCount")
	value := entities.MapfunctionalitywithgroupEntities{}
	var query string
	// if tz.Mstfunctionailyid == 2 {
	// 	query = getMapfunctionalitywithgroupcountwithterm
	// } else {
	// 	query = getMapfunctionalitywithgroupcount
	// }
	var getMststatecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMststatecount = "SELECT count(a.id) as total FROM mapfunctionalitywithgroup a left join mstclientuser h on a.userid = h.id and h.activeflag=1 and h.deleteflag=0,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mapfunctionality g WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.id "
	} else if OrgnTypeID == 2 {
		getMststatecount = "SELECT count(a.id) as total FROM mapfunctionalitywithgroup a left join mstclientuser h on a.userid = h.id and h.activeflag=1 and h.deleteflag=0,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mapfunctionality g WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.id "
		params = append(params, tz.Clientid)
	} else {
		getMststatecount = "SELECT count(a.id) as total FROM mapfunctionalitywithgroup a left join mstclientuser h on a.userid = h.id and h.activeflag=1 and h.deleteflag=0,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstfunctionality f,mapfunctionality g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.recorddifftypeid=d.id and a.recorddiffid=e.id and a.mstfunctionailyid=f.id and a.diffid=g.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMststatecount, params...).Scan(&value.Total)
	logger.Log.Println("Fetch count query ----->", query)
	// err := dbc.DB.QueryRow(getMapfunctionalitywithgroupcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMapfunctionalitywithgroupCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllOrganizationgrpnames(page *entities.MapfunctionalitywithgroupEntity) ([]entities.Organizationgrpname, error) {
	logger.Log.Println("In side GelAllMapfunctionalitywithgroup")
	values := []entities.Organizationgrpname{}
	rows, err := dbc.DB.Query(getsupportgrporganization, page.Clientid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMapfunctionalitywithgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Organizationgrpname{}
		rows.Scan(&value.Id, &value.Name, &value.Isworkflow)
		values = append(values, value)
	}
	return values, nil
}
