package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertRecorddifferentiongroup = "INSERT INTO mstrecorddifferentiongroup (clientid, mstorgnhirarchyid, mstworkdifferentiationtypeid,mstworkdifferentiationid, mstgroupid, mstuserid) VALUES (?,?,?,?,?,?)"
var duplicateRecorddifferentiongroup = "SELECT count(id) total FROM  mstrecorddifferentiongroup WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstworkdifferentiationtypeid=? AND mstworkdifferentiationid = ? AND mstgroupid = ? AND deleteflg = 0 AND activeflg=1"

//var getRecorddifferentiongroup = "SELECT a.id AS Id,a.clientid AS Clientid,a.mstorgnhirarchyid AS Mstorgnhirarchyid,a.mstworkdifferentiationtypeid AS Mstworkdifferentiationtypeid,a.mstgroupid AS Mstgroupid,a.mstuserid AS Mstuserid,a.activeflg AS Activeflg,b.name AS Clientname,c.name AS Mstorgnhirarchyname,e.id AS Recorddifftypeid,e.typename AS Recorddifftypename,f.id AS Recorddiffid,f.name AS Recorddiffname,g.typename AS Mstworkdifferentiationtypename,h.supportgroupname AS Supportgroupname,a.mstworkdifferentiationid as Mstworkdifferentiationid,i.name as Mstworkdifferentiationname,i.parentcategorynames FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstclientsupportgroup h,mstrecorddifferentiation i WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id ORDER BY a.id DESC LIMIT ?,?"
//var getRecorddifferentiongroup = "SELECT a.id AS Id,a.clientid AS Clientid,a.mstorgnhirarchyid AS Mstorgnhirarchyid,a.mstworkdifferentiationtypeid AS Mstworkdifferentiationtypeid,a.mstgroupid AS Mstgroupid,a.mstuserid AS Mstuserid,a.activeflg AS Activeflg,b.name AS Clientname,c.name AS Mstorgnhirarchyname,e.id AS Recorddifftypeid,e.typename AS Recorddifftypename,f.id AS Recorddiffid,f.name AS Recorddiffname,g.typename AS Mstworkdifferentiationtypename,h.name AS Supportgroupname,a.mstworkdifferentiationid as Mstworkdifferentiationid,i.name as Mstworkdifferentiationname,i.parentcategorynames FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstsupportgrp h,mstrecorddifferentiation i WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id ORDER BY a.id DESC LIMIT ?,?"
var getRecorddifferentiongroupcount = "SELECT count(a.id) as total FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstclientsupportgroup h,mstrecorddifferentiation i WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id"
var updateRecorddifferentiongroup = "UPDATE mstrecorddifferentiongroup SET mstorgnhirarchyid = ?, mstworkdifferentiationtypeid=?,mstworkdifferentiationid = ?, mstgroupid = ?, mstuserid = ? WHERE id = ? "
var deleteRecorddifferentiongroup = "UPDATE mstrecorddifferentiongroup SET deleteflg = '1' WHERE id = ? "
var getworkinglevel = "select a.mainrecorddifftypeid as Mstworkdifferentiationid,b.typename as Levelname from mstworkdifferentiation a,mstrecorddifferentiationtype b WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.forrecorddifftypeid=? and a.forrecorddiffid=? and a.deleteflg=0 and a.activeflg=1 and a.mainrecorddifftypeid = b.id"

func (dbc DbConn) CheckDuplicateRecorddifferentiongroup(tz *entities.RecorddifferentiongroupEntity) (entities.RecorddifferentiongroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecorddifferentiongroup")
	value := entities.RecorddifferentiongroupEntities{}
	err := dbc.DB.QueryRow(duplicateRecorddifferentiongroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstworkdifferentiationtypeid, tz.Mstworkdifferentiationid, tz.Mstgroupid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecorddifferentiongroup Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) CheckDuplicateRecorddifferentiongroupWithTx(tz *entities.RecorddifferentiongroupEntity, i int) (entities.RecorddifferentiongroupEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecorddifferentiongroup")
	value := entities.RecorddifferentiongroupEntities{}
	err := dbc.DB.QueryRow(duplicateRecorddifferentiongroup, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstworkdifferentiationtypeids[i], tz.Mstworkdifferentiationids[i], tz.Mstgroupid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecorddifferentiongroup Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc TxConn) InsertRecorddifferentiongroup(tz *entities.RecorddifferentiongroupEntity, i int) (int64, error) {
	logger.Log.Println("In side InsertRecorddifferentiongroup")
	logger.Log.Println("Query -->", insertRecorddifferentiongroup)
	stmt, err := dbc.TX.Prepare(insertRecorddifferentiongroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecorddifferentiongroup Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstworkdifferentiationtypeids[i], tz.Mstworkdifferentiationids[i], tz.Mstgroupid, tz.Mstuserid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstworkdifferentiationtypeids[i], tz.Mstworkdifferentiationids[i], tz.Mstgroupid, tz.Mstuserid)
	if err != nil {
		logger.Log.Println("InsertRecorddifferentiongroup Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllRecorddifferentiongroup(tz *entities.RecorddifferentiongroupEntity, OrgnType int64) ([]entities.RecorddifferentiongroupEntity, error) {
	logger.Log.Println("In side GelAllRecorddifferentiongroup")
	values := []entities.RecorddifferentiongroupEntity{}

	var getRecorddifferentiongroup string
	var params []interface{}
	if OrgnType == 1 {
		getRecorddifferentiongroup = "SELECT a.id AS Id,a.clientid AS Clientid,a.mstorgnhirarchyid AS Mstorgnhirarchyid,a.mstworkdifferentiationtypeid AS Mstworkdifferentiationtypeid,a.mstgroupid AS Mstgroupid,a.mstuserid AS Mstuserid,a.activeflg AS Activeflg,b.name AS Clientname,c.name AS Mstorgnhirarchyname,e.id AS Recorddifftypeid,e.typename AS Recorddifftypename,f.id AS Recorddiffid,f.name AS Recorddiffname,g.typename AS Mstworkdifferentiationtypename,h.name AS Supportgroupname,a.mstworkdifferentiationid as Mstworkdifferentiationid,i.name as Mstworkdifferentiationname,i.parentcategorynames FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstsupportgrp h,mstrecorddifferentiation i WHERE  a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getRecorddifferentiongroup = "SELECT a.id AS Id,a.clientid AS Clientid,a.mstorgnhirarchyid AS Mstorgnhirarchyid,a.mstworkdifferentiationtypeid AS Mstworkdifferentiationtypeid,a.mstgroupid AS Mstgroupid,a.mstuserid AS Mstuserid,a.activeflg AS Activeflg,b.name AS Clientname,c.name AS Mstorgnhirarchyname,e.id AS Recorddifftypeid,e.typename AS Recorddifftypename,f.id AS Recorddiffid,f.name AS Recorddiffname,g.typename AS Mstworkdifferentiationtypename,h.name AS Supportgroupname,a.mstworkdifferentiationid as Mstworkdifferentiationid,i.name as Mstworkdifferentiationname,i.parentcategorynames FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstsupportgrp h,mstrecorddifferentiation i WHERE a.clientid = ? AND  a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getRecorddifferentiongroup = "SELECT a.id AS Id,a.clientid AS Clientid,a.mstorgnhirarchyid AS Mstorgnhirarchyid,a.mstworkdifferentiationtypeid AS Mstworkdifferentiationtypeid,a.mstgroupid AS Mstgroupid,a.mstuserid AS Mstuserid,a.activeflg AS Activeflg,b.name AS Clientname,c.name AS Mstorgnhirarchyname,e.id AS Recorddifftypeid,e.typename AS Recorddifftypename,f.id AS Recorddiffid,f.name AS Recorddiffname,g.typename AS Mstworkdifferentiationtypename,h.name AS Supportgroupname,a.mstworkdifferentiationid as Mstworkdifferentiationid,i.name as Mstworkdifferentiationname,i.parentcategorynames FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstsupportgrp h,mstrecorddifferentiation i WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getRecorddifferentiongroup, params...)

	//rows, err := dbc.DB.Query(getRecorddifferentiongroup, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiongroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecorddifferentiongroupEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstworkdifferentiationtypeid, &value.Mstgroupid, &value.Mstuserid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifftypeid, &value.Recorddifftypename, &value.Recorddiffid, &value.Recorddiffname, &value.Mstworkdifferentiationtypename, &value.Supportgroupname, &value.Mstworkdifferentiationid, &value.Name, &value.Parentcategorynames)
		if len(value.Parentcategorynames) > 0 {
			value.Mstworkdifferentiationname = value.Name + "(" + value.Parentcategorynames + ")"
		} else {
			value.Mstworkdifferentiationname = value.Name
		}
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateRecorddifferentiongroup(tz *entities.RecorddifferentiongroupEntity) error {
	logger.Log.Println("In side UpdateRecorddifferentiongroup")
	stmt, err := dbc.DB.Prepare(updateRecorddifferentiongroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecorddifferentiongroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstworkdifferentiationtypeid, tz.Mstworkdifferentiationid, tz.Mstgroupid, tz.Mstuserid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecorddifferentiongroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteRecorddifferentiongroup(tz *entities.RecorddifferentiongroupEntity) error {
	logger.Log.Println("In side DeleteRecorddifferentiongroup")
	stmt, err := dbc.DB.Prepare(deleteRecorddifferentiongroup)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecorddifferentiongroup Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecorddifferentiongroup Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetRecorddifferentiongroupCount(tz *entities.RecorddifferentiongroupEntity, OrgnTypeID int64) (entities.RecorddifferentiongroupEntities, error) {
	logger.Log.Println("In side GetRecorddifferentiongroupCount")
	value := entities.RecorddifferentiongroupEntities{}

	var getRecorddifferentiongroupcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getRecorddifferentiongroupcount = "SELECT count(a.id) as total FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstclientsupportgroup h,mstrecorddifferentiation i WHERE a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id"
	} else if OrgnTypeID == 2 {
		getRecorddifferentiongroupcount = "SELECT count(a.id) as total FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstclientsupportgroup h,mstrecorddifferentiation i WHERE a.clientid = ? AND a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id"
		params = append(params, tz.Clientid)
	} else {
		getRecorddifferentiongroupcount = "SELECT count(a.id) as total FROM mstrecorddifferentiongroup a,mstclient b,mstorgnhierarchy c,mstworkdifferentiation d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiationtype g,mstclientsupportgroup h,mstrecorddifferentiation i WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg = 0 AND a.activeflg = 1 AND d.deleteflg = 0 AND d.activeflg = 1 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.mstworkdifferentiationtypeid = d.mainrecorddifftypeid AND d.forrecorddifftypeid = e.id AND d.forrecorddiffid = f.id AND d.mainrecorddifftypeid = g.id AND a.mstgroupid = h.id AND  a.mstworkdifferentiationid =i.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getRecorddifferentiongroupcount, params...).Scan(&value.Total)

	//err := dbc.DB.QueryRow(getRecorddifferentiongroupcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetRecorddifferentiongroupCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetWorkinglevel(page *entities.RecorddifferentiongroupEntity) ([]entities.WorkinglevelEntity, error) {
	logger.Log.Println("In side GelAllRecorddifferentiongroup")
	logger.Log.Println("Query --->", getworkinglevel)
	values := []entities.WorkinglevelEntity{}
	rows, err := dbc.DB.Query(getworkinglevel, page.Clientid, page.Mstorgnhirarchyid, page.Recorddifftypeid, page.Recorddiffid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiongroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.WorkinglevelEntity{}
		err := rows.Scan(&value.Mstworkdifferentiationid, &value.Levelname)
		logger.Log.Println("Error --->", err)
		values = append(values, value)
	}
	logger.Log.Println("values --->", values)
	return values, nil
}
