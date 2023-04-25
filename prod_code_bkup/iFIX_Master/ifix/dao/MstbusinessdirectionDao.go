package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstbusinessdirection = "INSERT INTO mstbusinessdirection (clientid, mstorgnhirarchyid, mstrecorddifferentiationtypeid, mstrecorddifferentiationid, direction, baseconfig) VALUES (?,?,?,?,?,?)"
var duplicateMstbusinessdirection = "SELECT count(id) total FROM  mstbusinessdirection WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstrecorddifferentiationtypeid = ? AND mstrecorddifferentiationid = ? AND direction = ? AND deleteflg = 0"
var duplicateCheckForDefaultPriority = "SELECT count(id) total FROM  mstbusinessdirection WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstrecorddifferentiationtypeid = ? AND mstrecorddifferentiationid = ? AND baseconfig <> 0 AND deleteflg = 0"
// var getMstbusinessdirection = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstrecorddifferentiationtypeid as Mstrecorddifferentiationtypeid, mstrecorddifferentiationid as Mstrecorddifferentiationid, direction as Direction, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=mstrecorddifferentiationtypeid and deleteflg =0 and activeflg=1) AS Mstrecorddifferentiationtypename,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationid and deleteflg =0 and activeflg=1) as Mstrecorddifferentiationname FROM mstbusinessdirection WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
// var getMstbusinessdirectioncount = "SELECT count(id) total FROM  mstbusinessdirection WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
var updateMstbusinessdirection = "UPDATE mstbusinessdirection SET mstorgnhirarchyid = ?, mstrecorddifferentiationtypeid = ?, mstrecorddifferentiationid = ?, direction = ?, baseconfig = ? WHERE id = ? "
var deleteMstbusinessdirection = "UPDATE mstbusinessdirection SET deleteflg = '1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstbusinessdirection(tz *entities.MstbusinessdirectionEntity) (entities.MstbusinessdirectionEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstbusinessdirection")
	value := entities.MstbusinessdirectionEntities{}
	err := dbc.DB.QueryRow(duplicateMstbusinessdirection, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtypeid, tz.Mstrecorddifferentiationid, tz.Direction).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstbusinessdirection Get Statement Prepare Error", err)
		return value, err
	}
}


func (dbc DbConn) DuplicateCheckingForDefaultPriority(tz *entities.MstbusinessdirectionEntity) (entities.MstbusinessdirectionEntities, error){
	logger.Log.Println("\n  ------>>>>>   In side DuplicateCheckingForDefaultPriority ")
	value := entities.MstbusinessdirectionEntities{}
	if tz.Baseconfig == 1 {
		err := dbc.DB.QueryRow(duplicateCheckForDefaultPriority, tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtypeid, tz.Mstrecorddifferentiationid).Scan(&value.Total)
		switch err {
		case sql.ErrNoRows:
			value.Total = 0
			return value, nil
		case nil:
			logger.Log.Println("\n value.Total ====    ", value.Total)
			return value, nil
		default:
			logger.Log.Println("DuplicateCheckingForDefaultPriority Get Statement Prepare Error", err)
			return value, err
		}
	} else {
		value.Total = 0
		return value, nil
	}
}


func (dbc DbConn) InsertMstbusinessdirection(tz *entities.MstbusinessdirectionEntity) (int64, error) {
	logger.Log.Println("In side InsertMstbusinessdirection")
	logger.Log.Println("Query -->", insertMstbusinessdirection)
	stmt, err := dbc.DB.Prepare(insertMstbusinessdirection)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstbusinessdirection Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtypeid, tz.Mstrecorddifferentiationid, tz.Direction, tz.Baseconfig)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtypeid, tz.Mstrecorddifferentiationid, tz.Direction, tz.Baseconfig)
	if err != nil {
		logger.Log.Println("InsertMstbusinessdirection Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstbusinessdirection(page *entities.MstbusinessdirectionEntity, OrgnType int64) ([]entities.MstbusinessdirectionEntity, error) {
	logger.Log.Println("In side GelAllMstbusinessdirection")
	values := []entities.MstbusinessdirectionEntity{}
	var getMstbusinessdirection string
	var params []interface{}
	if OrgnType == 1 {
		getMstbusinessdirection = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstrecorddifferentiationtypeid as Mstrecorddifferentiationtypeid, mstrecorddifferentiationid as Mstrecorddifferentiationid, direction as Direction, baseconfig as Baseconfig, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=mstrecorddifferentiationtypeid and deleteflg =0 and activeflg=1) AS Mstrecorddifferentiationtypename,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationid and deleteflg =0 and activeflg=1) as Mstrecorddifferentiationname FROM mstbusinessdirection WHERE  deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstbusinessdirection = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstrecorddifferentiationtypeid as Mstrecorddifferentiationtypeid, mstrecorddifferentiationid as Mstrecorddifferentiationid, direction as Direction, baseconfig as Baseconfig, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=mstrecorddifferentiationtypeid and deleteflg =0 and activeflg=1) AS Mstrecorddifferentiationtypename,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationid and deleteflg =0 and activeflg=1) as Mstrecorddifferentiationname FROM mstbusinessdirection WHERE clientid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstbusinessdirection = "SELECT id as Id, clientid as Clientid, mstorgnhirarchyid as Mstorgnhirarchyid, mstrecorddifferentiationtypeid as Mstrecorddifferentiationtypeid, mstrecorddifferentiationid as Mstrecorddifferentiationid, direction as Direction, baseconfig as Baseconfig, activeflg as Activeflg,(select name from mstclient where id=clientid) as Clientname,(select name from mstorgnhierarchy where id=mstorgnhirarchyid) as Mstorgnhirarchyname,(select typename from mstrecorddifferentiationtype where id=mstrecorddifferentiationtypeid and deleteflg =0 and activeflg=1) AS Mstrecorddifferentiationtypename,(select name from mstrecorddifferentiation where id=mstrecorddifferentiationid and deleteflg =0 and activeflg=1) as Mstrecorddifferentiationname FROM mstbusinessdirection WHERE clientid = ? AND mstorgnhirarchyid = ?  AND deleteflg =0 and activeflg=1 ORDER BY id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}

	rows, err := dbc.DB.Query(getMstbusinessdirection, params...)

	// rows, err := dbc.DB.Query(getMstbusinessdirection, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstbusinessdirection Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstbusinessdirectionEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstrecorddifferentiationtypeid, &value.Mstrecorddifferentiationid, &value.Direction, &value.Baseconfig, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Mstrecorddifferentiationtypename, &value.Mstrecorddifferentiationname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstbusinessdirection(tz *entities.MstbusinessdirectionEntity) error {
	logger.Log.Println("In side UpdateMstbusinessdirection")
	stmt, err := dbc.DB.Prepare(updateMstbusinessdirection)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstbusinessdirection Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Mstrecorddifferentiationtypeid, tz.Mstrecorddifferentiationid, tz.Direction, tz.Baseconfig, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstbusinessdirection Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstbusinessdirection(tz *entities.MstbusinessdirectionEntity) error {
	logger.Log.Println("In side DeleteMstbusinessdirection")
	stmt, err := dbc.DB.Prepare(deleteMstbusinessdirection)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstbusinessdirection Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstbusinessdirection Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstbusinessdirectionCount(tz *entities.MstbusinessdirectionEntity, OrgnTypeID int64) (entities.MstbusinessdirectionEntities, error) {
	logger.Log.Println("In side GetMstbusinessdirectionCount")
	value := entities.MstbusinessdirectionEntities{}
	var getMstbusinessdirectioncount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstbusinessdirectioncount = "SELECT count(id) total FROM  mstbusinessdirection WHERE   deleteflg =0 and activeflg=1"
	} else if OrgnTypeID == 2 {
		getMstbusinessdirectioncount = "SELECT count(id) total FROM  mstbusinessdirection WHERE clientid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getMstbusinessdirectioncount = "SELECT count(id) total FROM  mstbusinessdirection WHERE clientid = ? AND mstorgnhirarchyid = ? AND  deleteflg =0 and activeflg=1"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstbusinessdirectioncount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstbusinessdirectioncount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstbusinessdirectionCount Get Statement Prepare Error", err)
		return value, err
	}
}
