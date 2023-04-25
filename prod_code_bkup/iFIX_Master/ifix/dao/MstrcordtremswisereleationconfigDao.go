package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstrcordtremswisereleationconfig = "INSERT INTO mstrcordtremswisereleationconfig (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, releationid, termsid) VALUES (?,?,?,?,?,?)"
var duplicateMstrcordtremswisereleationconfig = "SELECT count(id) total FROM  mstrcordtremswisereleationconfig WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND releationid = ? AND termsid = ? AND deleteflg = 0 and activeflg=1"

//var getMstrcordtremswisereleationconfig = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, recorddifftypeid as Recorddifftypeid, recorddiffid as Recorddiffid, releationid as Releationid, termsid as Termsid, activeflg as Activeflg,(select name from mstclient where id = clientid ) as Clientname,(select name from mstorgnhierarchy where id = mstorgnhirarchyid ) as Mstorgnhirarchyname,(SELECT typename FROM mstrecorddifferentiationtype WHERE id = recorddifftypeid AND deleteflg=0 AND activeflg=1) as Recorddifftypename,(SELECT name FROM mstrecorddifferentiation WHERE id = recorddiffid AND deleteflg=0 AND activeflg=1) as Recorddiffname,(SELECT name FROM mstrecorddifferentiation WHERE id = releationid AND deleteflg=0 AND activeflg=1) as Releationname,(SELECT termname FROM mstrecordterms WHERE id = termsid AND deleteflg=0 AND activeflg=1) as Termname FROM mstrcordtremswisereleationconfig WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
//var getMstrcordtremswisereleationconfigcount = "SELECT count(id) total FROM  mstrcordtremswisereleationconfig WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMstrcordtremswisereleationconfig = "UPDATE mstrcordtremswisereleationconfig SET mstorgnhirarchyid = ?, recorddifftypeid = ?, recorddiffid = ?, releationid = ?, termsid = ? WHERE id = ? "
var deleteMstrcordtremswisereleationconfig = "UPDATE mstrcordtremswisereleationconfig SET deleteflg = '1' WHERE id = ? "
var getrecordreleationname = "SELECT a.id,a.name FROM mstrecorddifferentiation as a,mstrecordtype c WHERE c.clientid=? AND c.mstorgnhirarchyid=? AND c.fromrecorddifftypeid=? AND c.fromrecorddiffid=? AND c.torecorddifftypeid in (SELECT id FROM mstrecorddifferentiationtype WHERE seqno=10) AND c.torecorddiffid = a.id AND a.deleteflg=0 AND a.activeflg=1 AND c.deleteflg=0 AND c.activeflg=1"
var termnames = "SELECT a.id,a.termname FROM mstrecordterms a,mststateterm b WHERE b.clientid=? AND b.mstorgnhirarchyid=? AND b.recorddifftypeid=? AND b.recorddiffid=? and a.id = b.recordtermid AND a.activeflg=1 AND a.deleteflg=0 AND b.activeflg=1 AND b.deleteflg=0"

func (dbc DbConn) CheckDuplicateMstrcordtremswisereleationconfig(tz *entities.MstrcordtremswisereleationconfigEntity) (entities.MstrcordtremswisereleationconfigEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstrcordtremswisereleationconfig")
	value := entities.MstrcordtremswisereleationconfigEntities{}
	err := dbc.DB.QueryRow(duplicateMstrcordtremswisereleationconfig, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Releationid, tz.Termsid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstrcordtremswisereleationconfig Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstrcordtremswisereleationconfig(tz *entities.MstrcordtremswisereleationconfigEntity) (int64, error) {
	logger.Log.Println("In side InsertMstrcordtremswisereleationconfig")
	logger.Log.Println("Query -->", insertMstrcordtremswisereleationconfig)
	stmt, err := dbc.DB.Prepare(insertMstrcordtremswisereleationconfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstrcordtremswisereleationconfig Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Releationid, tz.Termsid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Releationid, tz.Termsid)
	if err != nil {
		logger.Log.Println("InsertMstrcordtremswisereleationconfig Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstrcordtremswisereleationconfig(tz *entities.MstrcordtremswisereleationconfigEntity, OrgnType int64) ([]entities.MstrcordtremswisereleationconfigEntity, error) {
	logger.Log.Println("In side GelAllMstrcordtremswisereleationconfig")
	values := []entities.MstrcordtremswisereleationconfigEntity{}
	var getMstrcordtremswisereleationconfig string
	var params []interface{}

	if OrgnType == 1 {
		getMstrcordtremswisereleationconfig = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, recorddifftypeid as Recorddifftypeid, recorddiffid as Recorddiffid, releationid as Releationid, termsid as Termsid, activeflg as Activeflg,(select name from mstclient where id = clientid ) as Clientname,(select name from mstorgnhierarchy where id = mstorgnhirarchyid ) as Mstorgnhirarchyname,(SELECT typename FROM mstrecorddifferentiationtype WHERE id = recorddifftypeid AND deleteflg=0 AND activeflg=1) as Recorddifftypename,(SELECT name FROM mstrecorddifferentiation WHERE id = recorddiffid AND deleteflg=0 AND activeflg=1) as Recorddiffname,(SELECT name FROM mstrecorddifferentiation WHERE id = releationid AND deleteflg=0 AND activeflg=1) as Releationname,(SELECT termname FROM mstrecordterms WHERE id = termsid AND deleteflg=0 AND activeflg=1) as Termname FROM mstrcordtremswisereleationconfig WHERE deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstrcordtremswisereleationconfig = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, recorddifftypeid as Recorddifftypeid, recorddiffid as Recorddiffid, releationid as Releationid, termsid as Termsid, activeflg as Activeflg,(select name from mstclient where id = clientid ) as Clientname,(select name from mstorgnhierarchy where id = mstorgnhirarchyid ) as Mstorgnhirarchyname,(SELECT typename FROM mstrecorddifferentiationtype WHERE id = recorddifftypeid AND deleteflg=0 AND activeflg=1) as Recorddifftypename,(SELECT name FROM mstrecorddifferentiation WHERE id = recorddiffid AND deleteflg=0 AND activeflg=1) as Recorddiffname,(SELECT name FROM mstrecorddifferentiation WHERE id = releationid AND deleteflg=0 AND activeflg=1) as Releationname,(SELECT termname FROM mstrecordterms WHERE id = termsid AND deleteflg=0 AND activeflg=1) as Termname FROM mstrcordtremswisereleationconfig WHERE clientid = ? AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstrcordtremswisereleationconfig = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, recorddifftypeid as Recorddifftypeid, recorddiffid as Recorddiffid, releationid as Releationid, termsid as Termsid, activeflg as Activeflg,(select name from mstclient where id = clientid ) as Clientname,(select name from mstorgnhierarchy where id = mstorgnhirarchyid ) as Mstorgnhirarchyname,(SELECT typename FROM mstrecorddifferentiationtype WHERE id = recorddifftypeid AND deleteflg=0 AND activeflg=1) as Recorddifftypename,(SELECT name FROM mstrecorddifferentiation WHERE id = recorddiffid AND deleteflg=0 AND activeflg=1) as Recorddiffname,(SELECT name FROM mstrecorddifferentiation WHERE id = releationid AND deleteflg=0 AND activeflg=1) as Releationname,(SELECT termname FROM mstrecordterms WHERE id = termsid AND deleteflg=0 AND activeflg=1) as Termname FROM mstrcordtremswisereleationconfig WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstrcordtremswisereleationconfig, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstrcordtremswisereleationconfig Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstrcordtremswisereleationconfigEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Releationid, &value.Termsid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifftypename, &value.Recorddiffname, &value.Releationname, &value.Termname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstrcordtremswisereleationconfig(tz *entities.MstrcordtremswisereleationconfigEntity) error {
	logger.Log.Println("In side UpdateMstrcordtremswisereleationconfig")
	stmt, err := dbc.DB.Prepare(updateMstrcordtremswisereleationconfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstrcordtremswisereleationconfig Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Releationid, tz.Termsid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstrcordtremswisereleationconfig Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstrcordtremswisereleationconfig(tz *entities.MstrcordtremswisereleationconfigEntity) error {
	logger.Log.Println("In side DeleteMstrcordtremswisereleationconfig")
	stmt, err := dbc.DB.Prepare(deleteMstrcordtremswisereleationconfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstrcordtremswisereleationconfig Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstrcordtremswisereleationconfig Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstrcordtremswisereleationconfigCount(tz *entities.MstrcordtremswisereleationconfigEntity, OrgnTypeID int64) (entities.MstrcordtremswisereleationconfigEntities, error) {
	logger.Log.Println("In side GetMstrcordtremswisereleationconfigCount")
	value := entities.MstrcordtremswisereleationconfigEntities{}
	var getMstrcordtremswisereleationconfigcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstrcordtremswisereleationconfigcount = "SELECT count(id) total FROM  mstrcordtremswisereleationconfig WHERE deleteflg =0 and activeflg=1"
	} else if OrgnTypeID == 2 {
		getMstrcordtremswisereleationconfigcount = "SELECT count(id) total FROM  mstrcordtremswisereleationconfig WHERE clientid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getMstrcordtremswisereleationconfigcount = "SELECT count(id) total FROM  mstrcordtremswisereleationconfig WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstrcordtremswisereleationconfigcount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstrcordtremswisereleationconfigCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetRecordreleationnames(page *entities.MstrcordtremswisereleationconfigEntity) ([]entities.Recordreleationdetails, error) {
	logger.Log.Println("In side GelAllMstrcordtremswisereleationconfig")
	values := []entities.Recordreleationdetails{}
	rows, err := dbc.DB.Query(getrecordreleationname, page.Clientid, page.Mstorgnhirarchyid, page.Recorddifftypeid, page.Recorddiffid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstrcordtremswisereleationconfig Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Recordreleationdetails{}
		rows.Scan(&value.ID, &value.Releationname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetRecordtermnames(page *entities.MstrcordtremswisereleationconfigEntity) ([]entities.Recordtermnames, error) {
	logger.Log.Println("In side GelAllMstrcordtremswisereleationconfig")
	values := []entities.Recordtermnames{}
	rows, err := dbc.DB.Query(termnames, page.Clientid, page.Mstorgnhirarchyid, page.Recorddifftypeid, page.Recorddiffid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstrcordtremswisereleationconfig Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Recordtermnames{}
		rows.Scan(&value.ID, &value.Names)
		values = append(values, value)
	}
	return values, nil
}
