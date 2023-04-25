package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"strconv"
)

var insertMstrecordfield = "INSERT INTO mstrecordfield (clientid, mstorgnhirarchyid,mstrecordfieldtype, recordtermid,activeflg,deleteflg,audittransactionid) VALUES (?,?,?,?,1,0,1) "
var insetMstrecordfielddiff = "INSERT INTO mstrecordfielddiff(clientid, mstorgnhirarchyid, mstrecordfieldid, recorddifftypeid, recorddiffid, deleteflg, activeflg, audittransactionid) VALUES (?,?,?,?,?,0,1,1) "

//var duplicateMstrecordfield = "SELECT count(id) total FROM  mstrecordfield WHERE clientid = ? AND mstorgnhirarchyid = ? AND mstrecordfieldtype=? AND recordtermid = ? AND deleteflg = 0"

// var getMstrecordfield = "SELECT mstrecordfield.id, mstrecordfield.clientid, mstrecordfield.mstorgnhirarchyid, mstrecordfield.mstrecordfieldtype, mstrecordfield.recordtermid, mstrecordfield.activeflg, mstrecordterms.termname,msttermtype.termtypename,mstclient.name clientname,mstorgnhierarchy.name orgname FROM mstrecordfield JOIN mstclient ON mstrecordfield.clientid=mstclient.id JOIN mstorgnhierarchy ON mstrecordfield.clientid=mstorgnhierarchy.clientid AND mstrecordfield.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecordterms ON mstrecordfield.clientid=mstrecordterms.clientid AND mstrecordfield.mstorgnhirarchyid=mstrecordterms.mstorgnhirarchyid AND mstrecordfield.recordtermid = mstrecordterms.id AND mstrecordterms.activeflg=1 AND mstrecordterms.deleteflg=0 JOIN msttermtype ON mstrecordterms.termtypeid = msttermtype.id WHERE mstrecordfield.clientid=? AND mstrecordfield.mstorgnhirarchyid=? AND mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0 ORDER BY mstrecordfield.id DESC LIMIT ?,?"
var getMstrecordfielddiff = "SELECT mstrecordfielddiff.id,mstrecordfielddiff.clientid,mstrecordfielddiff.mstorgnhirarchyid,mstrecordfielddiff.mstrecordfieldid,mstrecordfielddiff.recorddifftypeid, mstrecordfielddiff.recorddiffid,mstrecordfielddiff.activeflg,mstrecorddifferentiationtype.typename,mstrecorddifferentiation.name,mstrecorddifferentiationtype.parentid FROM mstrecordfielddiff,mstrecorddifferentiationtype, mstrecorddifferentiation WHERE mstrecordfielddiff.activeflg = 1 AND mstrecordfielddiff.deleteflg = 0 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiation.deleteflg = 0 AND mstrecorddifferentiation.activeflg=1 AND mstrecordfielddiff.recorddifftypeid=mstrecorddifferentiationtype.id AND mstrecordfielddiff.recorddifftypeid=mstrecorddifferentiation.recorddifftypeid AND  mstrecordfielddiff.recorddiffid = mstrecorddifferentiation.id  AND mstrecordfielddiff.clientid = ? AND mstrecordfielddiff.mstorgnhirarchyid = ? AND mstrecordfielddiff.mstrecordfieldid = ? "

// var getMstrecordfieldcount = "SELECT count(mstrecordfield.id) total FROM mstrecordfield JOIN mstclient ON mstrecordfield.clientid=mstclient.id JOIN mstorgnhierarchy ON mstrecordfield.clientid=mstorgnhierarchy.clientid AND mstrecordfield.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecordterms ON mstrecordfield.clientid=mstrecordterms.clientid AND mstrecordfield.mstorgnhirarchyid=mstrecordterms.mstorgnhirarchyid AND mstrecordfield.recordtermid = mstrecordterms.id AND mstrecordterms.activeflg=1 AND mstrecordterms.deleteflg=0 JOIN msttermtype ON mstrecordterms.termtypeid = msttermtype.id WHERE mstrecordfield.clientid=? AND mstrecordfield.mstorgnhirarchyid=? AND mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0"

var updateMstrecordfield = "UPDATE mstrecordfield SET mstrecordfieldtype = ?, recordtermid = ?  WHERE id = ? "

var deleteMstrecordfield = "UPDATE mstrecordfield SET deleteflg = '1' WHERE id = ? "

var deleteMstrecordfielddiff = "UPDATE mstrecordfielddiff SET deleteflg = '1' WHERE mstrecordfieldid = ? "

func (dbc DbConn) CheckDuplicateMstrecordfield(tz *entities.MstrecordfieldEntity) (entities.MstrecordfieldEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstrecordfield")
	var diffchecksql = "SELECT COUNT(mstrecordfield.id) total FROM mstrecordfield WHERE mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0 AND mstrecordfield.clientid=" + strconv.Itoa(int(tz.Clientid)) + " AND mstrecordfield.mstorgnhirarchyid=" + strconv.Itoa(int(tz.Mstorgnhirarchyid)) + " AND mstrecordfield.mstrecordfieldtype='" + tz.Mstrecordfieldtype + "' AND mstrecordfield.recordtermid=" + strconv.Itoa(int(tz.Recordtermid)) + " AND mstrecordfield.id IN ( SELECT mstrecordfielddiff.mstrecordfieldid FROM mstrecordfielddiff WHERE mstrecordfielddiff.activeflg=1 AND mstrecordfielddiff.deleteflg=0 AND mstrecordfielddiff.clientid=" + strconv.Itoa(int(tz.Clientid)) + " AND mstrecordfielddiff.mstorgnhirarchyid=" + strconv.Itoa(int(tz.Mstorgnhirarchyid)) + " GROUP BY mstrecordfielddiff.mstrecordfieldid HAVING COUNT(mstrecordfielddiff.id)=" + strconv.Itoa(len(tz.MstrecordfielddiffEntities)) + " AND mstrecordfielddiff.mstrecordfieldid IN ( SELECT mstrecordfielddiff.mstrecordfieldid FROM mstrecordfielddiff WHERE mstrecordfielddiff.activeflg=1 AND mstrecordfielddiff.deleteflg=0 AND mstrecordfielddiff.clientid=" + strconv.Itoa(int(tz.Clientid)) + " AND mstrecordfielddiff.mstorgnhirarchyid=" + strconv.Itoa(int(tz.Mstorgnhirarchyid)) + " GROUP BY mstrecordfielddiff.mstrecordfieldid HAVING "
	var tmpsql = ""
	for _, v := range tz.MstrecordfielddiffEntities {
		if tmpsql != "" {
			tmpsql = tmpsql + " + "
		}
		tmpsql = tmpsql + " SUM(IF(mstrecordfielddiff.recorddifftypeid=" + strconv.Itoa(int(v.Recorddifftypeid)) + " AND mstrecordfielddiff.recorddiffid=" + strconv.Itoa(int(v.Recorddiffid)) + ",1,0)) "
	}
	diffchecksql = diffchecksql + tmpsql + " = " + strconv.Itoa(len(tz.MstrecordfielddiffEntities)) + " )) "

	value := entities.MstrecordfieldEntities{}
	err := dbc.DB.QueryRow(diffchecksql).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstrecordfield Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstrecordfield(tz *entities.MstrecordfieldEntity) (int64, error) {
	logger.Log.Println("In side InsertMstrecordfield")
	logger.Log.Println("Query -->", insertMstrecordfield)
	stmt, err := dbc.DB.Prepare(insertMstrecordfield)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstrecordfield Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecordfieldtype, tz.Recordtermid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Mstrecordfieldtype, tz.Recordtermid)
	if err != nil {
		logger.Log.Println("InsertMstrecordfield Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) InsertMstrecordfielddiff(tz *entities.MstrecordfieldEntity, tz1 *entities.MstrecordfielddiffEntity, filedid int64) (int64, error) {
	logger.Log.Println("In side InsertMstrecordfielddiff")
	logger.Log.Println("Query -->", insetMstrecordfielddiff)
	stmt, err := dbc.DB.Prepare(insetMstrecordfielddiff)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstrecordfielddiff Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, filedid, tz1.Recorddifftypeid, tz1.Recorddiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, filedid, tz1.Recorddifftypeid, tz1.Recorddiffid)
	if err != nil {
		logger.Log.Println("InsertMstrecordfield Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstrecordfield(tz *entities.MstrecordfieldEntity, OrgnType int64) ([]entities.MstrecordfieldEntity, error) {
	logger.Log.Println("In side GelAllMstrecordfield")
	values := []entities.MstrecordfieldEntity{}
	var getMstrecordfield string
	var params []interface{}
	if OrgnType == 1 {
		getMstrecordfield = "SELECT mstrecordfield.id, mstrecordfield.clientid, mstrecordfield.mstorgnhirarchyid, mstrecordfield.mstrecordfieldtype, mstrecordfield.recordtermid, mstrecordfield.activeflg, mstrecordterms.termname,msttermtype.termtypename,mstclient.name clientname,mstorgnhierarchy.name orgname FROM mstrecordfield JOIN mstclient ON mstrecordfield.clientid=mstclient.id JOIN mstorgnhierarchy ON mstrecordfield.clientid=mstorgnhierarchy.clientid AND mstrecordfield.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecordterms ON mstrecordfield.clientid=mstrecordterms.clientid AND mstrecordfield.mstorgnhirarchyid=mstrecordterms.mstorgnhirarchyid AND mstrecordfield.recordtermid = mstrecordterms.id AND mstrecordterms.activeflg=1 AND mstrecordterms.deleteflg=0 JOIN msttermtype ON mstrecordterms.termtypeid = msttermtype.id WHERE  mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0 ORDER BY mstrecordfield.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstrecordfield = "SELECT mstrecordfield.id, mstrecordfield.clientid, mstrecordfield.mstorgnhirarchyid, mstrecordfield.mstrecordfieldtype, mstrecordfield.recordtermid, mstrecordfield.activeflg, mstrecordterms.termname,msttermtype.termtypename,mstclient.name clientname,mstorgnhierarchy.name orgname FROM mstrecordfield JOIN mstclient ON mstrecordfield.clientid=mstclient.id JOIN mstorgnhierarchy ON mstrecordfield.clientid=mstorgnhierarchy.clientid AND mstrecordfield.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecordterms ON mstrecordfield.clientid=mstrecordterms.clientid AND mstrecordfield.mstorgnhirarchyid=mstrecordterms.mstorgnhirarchyid AND mstrecordfield.recordtermid = mstrecordterms.id AND mstrecordterms.activeflg=1 AND mstrecordterms.deleteflg=0 JOIN msttermtype ON mstrecordterms.termtypeid = msttermtype.id WHERE mstrecordfield.clientid=? AND mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0 ORDER BY mstrecordfield.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstrecordfield = "SELECT mstrecordfield.id, mstrecordfield.clientid, mstrecordfield.mstorgnhirarchyid, mstrecordfield.mstrecordfieldtype, mstrecordfield.recordtermid, mstrecordfield.activeflg, mstrecordterms.termname,msttermtype.termtypename,mstclient.name clientname,mstorgnhierarchy.name orgname FROM mstrecordfield JOIN mstclient ON mstrecordfield.clientid=mstclient.id JOIN mstorgnhierarchy ON mstrecordfield.clientid=mstorgnhierarchy.clientid AND mstrecordfield.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecordterms ON mstrecordfield.clientid=mstrecordterms.clientid AND mstrecordfield.mstorgnhirarchyid=mstrecordterms.mstorgnhirarchyid AND mstrecordfield.recordtermid = mstrecordterms.id AND mstrecordterms.activeflg=1 AND mstrecordterms.deleteflg=0 JOIN msttermtype ON mstrecordterms.termtypeid = msttermtype.id WHERE mstrecordfield.clientid=? AND mstrecordfield.mstorgnhirarchyid=? AND mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0 ORDER BY mstrecordfield.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getMstrecordfield, params...)
	// rows, err := dbc.DB.Query(getMstrecordfield, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstrecordfield Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstrecordfieldEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstrecordfieldtype, &value.Recordtermid, &value.Activeflg, &value.Termname, &value.Termtypename, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllMstrecordfielddiff(page *entities.MstrecordfieldEntity) ([]entities.MstrecordfielddiffEntity, error) {
	logger.Log.Println("In side GetAllMstrecordfielddiff")
	values := []entities.MstrecordfielddiffEntity{}
	rows, err := dbc.DB.Query(getMstrecordfielddiff, page.Clientid, page.Mstorgnhirarchyid, page.Id)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstrecordfielddiff Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstrecordfielddiffEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Mstrecordfieldid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Activeflg, &value.Mstrecorddifferentiationtypename, &value.Mstrecorddifferentiationname, &value.RecorddifftypeParentid)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstrecordfield(tz *entities.MstrecordfieldEntity) error {
	logger.Log.Println("In side UpdateMstrecordfield")
	stmt, err := dbc.DB.Prepare(updateMstrecordfield)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstrecordfield Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstrecordfieldtype, tz.Recordtermid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstrecordfield Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstrecordfield(tz *entities.MstrecordfieldEntity) error {
	logger.Log.Println("In side DeleteMstrecordfield")
	stmt, err := dbc.DB.Prepare(deleteMstrecordfield)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstrecordfield Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstrecordfield Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstrecordfielddiff(tz *entities.MstrecordfieldEntity) error {
	logger.Log.Println("In side DeleteMstrecordfielddiff")
	stmt, err := dbc.DB.Prepare(deleteMstrecordfielddiff)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstrecordfielddiff Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstrecordfielddiff Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstrecordfieldCount(tz *entities.MstrecordfieldEntity, OrgnTypeID int64) (entities.MstrecordfieldEntities, error) {
	logger.Log.Println("In side GetMstrecordfieldCount")
	value := entities.MstrecordfieldEntities{}
	var getMstrecordfieldcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstrecordfieldcount = "SELECT count(mstrecordfield.id) total FROM mstrecordfield JOIN mstclient ON mstrecordfield.clientid=mstclient.id JOIN mstorgnhierarchy ON mstrecordfield.clientid=mstorgnhierarchy.clientid AND mstrecordfield.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecordterms ON mstrecordfield.clientid=mstrecordterms.clientid AND mstrecordfield.mstorgnhirarchyid=mstrecordterms.mstorgnhirarchyid AND mstrecordfield.recordtermid = mstrecordterms.id AND mstrecordterms.activeflg=1 AND mstrecordterms.deleteflg=0 JOIN msttermtype ON mstrecordterms.termtypeid = msttermtype.id WHERE mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getMstrecordfieldcount = "SELECT count(mstrecordfield.id) total FROM mstrecordfield JOIN mstclient ON mstrecordfield.clientid=mstclient.id JOIN mstorgnhierarchy ON mstrecordfield.clientid=mstorgnhierarchy.clientid AND mstrecordfield.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecordterms ON mstrecordfield.clientid=mstrecordterms.clientid AND mstrecordfield.mstorgnhirarchyid=mstrecordterms.mstorgnhirarchyid AND mstrecordfield.recordtermid = mstrecordterms.id AND mstrecordterms.activeflg=1 AND mstrecordterms.deleteflg=0 JOIN msttermtype ON mstrecordterms.termtypeid = msttermtype.id WHERE mstrecordfield.clientid=? AND mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0"
		params = append(params, tz.Clientid)
	} else {
		getMstrecordfieldcount = "SELECT count(mstrecordfield.id) total FROM mstrecordfield JOIN mstclient ON mstrecordfield.clientid=mstclient.id JOIN mstorgnhierarchy ON mstrecordfield.clientid=mstorgnhierarchy.clientid AND mstrecordfield.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecordterms ON mstrecordfield.clientid=mstrecordterms.clientid AND mstrecordfield.mstorgnhirarchyid=mstrecordterms.mstorgnhirarchyid AND mstrecordfield.recordtermid = mstrecordterms.id AND mstrecordterms.activeflg=1 AND mstrecordterms.deleteflg=0 JOIN msttermtype ON mstrecordterms.termtypeid = msttermtype.id WHERE mstrecordfield.clientid=? AND mstrecordfield.mstorgnhirarchyid=? AND mstrecordfield.activeflg=1 AND mstrecordfield.deleteflg=0"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstrecordfieldcount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getMstrecordfieldcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstrecordfieldCount Get Statement Prepare Error", err)
		return value, err
	}
}
