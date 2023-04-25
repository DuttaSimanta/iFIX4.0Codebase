package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstprocess = "INSERT INTO mstprocess (clientid, mstorgnhirarchyid, processname) VALUES (?,?,?)"
var duplicateMstprocess = "SELECT count(id) total FROM  mstprocess WHERE clientid = ? AND mstorgnhirarchyid = ? AND processname = ? and activeflg=1 and deleteflg=0"

// var getMstprocess = "select a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, b.recorddifftypeid as Recorddifftypeid, b.recorddiffid as Recorddiffid, b.id as Mstprocessrecordmapid, c.id as Mstprocesstoentityid, c.mstdatadictionaryfieldid as Mstdatadictionaryfieldid, a.activeflg as Activeflg, a.processname as Processname ,  (select name from mstorgnhierarchy where id = a.mstorgnhirarchyid) as Mstorgnhirarchyname , (select typename from mstrecorddifferentiationtype where id = b.recorddifftypeid) as Recorddifftypname , (select name from mstrecorddifferentiation where id = b.recorddiffid) as Recorddiffname , (select columnname from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Mstdatadictionaryfieldname , (select tablename from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarytablename,(select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Tableid,(select mstdatadictionarydbid from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarydbid,(select parentcategorynames from mstrecorddifferentiation where id = b.recorddiffid) as Parentcatnames,(select forrecorddifftypeid from mstworkdifferentiation where mainrecorddifftypeid=b.recorddifftypeid AND forrecorddifftypeid=b.recorddifftypeid AND forrecorddiffid=b.recorddiffid and activeflg = '1' AND deleteflg = 0) AS Forrecorddifftypeid,(select forrecorddiffid from mstworkdifferentiation where mainrecorddifftypeid=b.recorddifftypeid AND forrecorddifftypeid=b.recorddifftypeid AND forrecorddiffid=b.recorddiffid and activeflg = '1' AND deleteflg = 0) AS Forrecorddiffid from mstprocess a, mstprocessrecordmap b, mapprocesstoentity c where a.clientid = ? and a.mstorgnhirarchyid = ? and a.activeflg = '1' and a.deleteflg=0 and a.id = b.mstprocessid and a.id = c.mstprocessid and b.activeflg = '1' and b.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
// var getMstprocesscount = "select count(distinct a.id) as total from mstprocess a, mstprocessrecordmap b, mapprocesstoentity c where a.id = b.mstprocessid and a.id = c.mstprocessid and a.clientid = ? and a.mstorgnhirarchyid = ? and a.activeflg = 1 and a.deleteflg=0 and b.activeflg = '1' and b.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0"
var updateMstprocess = "UPDATE mstprocess SET mstorgnhirarchyid = ?, processname = ? WHERE id = ? "
var deleteMstprocess = "UPDATE mstprocess SET deleteflg = '1' WHERE id=?"

func (dbc DbConn) CheckDuplicateMstprocess(tz *entities.MstprocessEntity) (entities.MstprocessEntities, error) {
	value := entities.MstprocessEntities{}
	err := dbc.DB.QueryRow(duplicateMstprocess, tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstprocess Get Statement Prepare Error", err)
		return value, err
	}
}

func CheckDuplicateMstprocesswithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) (entities.MstprocessEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstprocess")
	value := entities.MstprocessEntities{}
	err := tx.QueryRow(duplicateMstprocess, tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstprocess Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertMstprocess(tz *entities.MstprocessEntity) (int64, error) {
	logger.Log.Println("In side InsertMstprocess")
	logger.Log.Println("Query -->", insertMstprocess)
	stmt, err := dbc.DB.Prepare(insertMstprocess)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstprocess Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname)
	if err != nil {
		logger.Log.Println("InsertMstprocess Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func InsertMstprocesswithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) (int64, error) {
	logger.Log.Println("Query -->", insertMstprocess)
	stmt, err := tx.Prepare(insertMstprocess)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertMstprocess Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Processname)
	if err != nil {
		logger.Log.Println("InsertMstprocess Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstprocess(page *entities.MstprocessEntity, OrgnType int64) ([]entities.MstprocessEntity, error) {
	logger.Log.Println("In side GelAllMstprocess")
	values := []entities.MstprocessEntity{}
	var getMstprocess string
	var params []interface{}
	if OrgnType == 1 {
		getMstprocess = "select a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, b.recorddifftypeid as Recorddifftypeid, b.recorddiffid as Recorddiffid, b.id as Mstprocessrecordmapid, c.id as Mstprocesstoentityid, c.mstdatadictionaryfieldid as Mstdatadictionaryfieldid, a.activeflg as Activeflg, a.processname as Processname ,  (select name from mstorgnhierarchy where id = a.mstorgnhirarchyid) as Mstorgnhirarchyname , (select typename from mstrecorddifferentiationtype where id = b.recorddifftypeid) as Recorddifftypname , (select name from mstrecorddifferentiation where id = b.recorddiffid) as Recorddiffname , (select columnname from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Mstdatadictionaryfieldname , (select tablename from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarytablename,(select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Tableid,(select mstdatadictionarydbid from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarydbid,(select parentcategorynames from mstrecorddifferentiation where id = b.recorddiffid) as Parentcatnames,(select forrecorddifftypeid from mstworkdifferentiation where mainrecorddifftypeid=b.recorddifftypeid AND forrecorddifftypeid=b.recorddifftypeid AND forrecorddiffid=b.recorddiffid and activeflg = '1' AND deleteflg = 0) AS Forrecorddifftypeid,(select forrecorddiffid from mstworkdifferentiation where mainrecorddifftypeid=b.recorddifftypeid AND forrecorddifftypeid=b.recorddifftypeid AND forrecorddiffid=b.recorddiffid and activeflg = '1' AND deleteflg = 0) AS Forrecorddiffid from mstprocess a, mstprocessrecordmap b, mapprocesstoentity c where a.activeflg = '1' and a.deleteflg=0 and a.id = b.mstprocessid and a.id = c.mstprocessid and b.activeflg = '1' and b.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstprocess = "select a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, b.recorddifftypeid as Recorddifftypeid, b.recorddiffid as Recorddiffid, b.id as Mstprocessrecordmapid, c.id as Mstprocesstoentityid, c.mstdatadictionaryfieldid as Mstdatadictionaryfieldid, a.activeflg as Activeflg, a.processname as Processname ,  (select name from mstorgnhierarchy where id = a.mstorgnhirarchyid) as Mstorgnhirarchyname , (select typename from mstrecorddifferentiationtype where id = b.recorddifftypeid) as Recorddifftypname , (select name from mstrecorddifferentiation where id = b.recorddiffid) as Recorddiffname , (select columnname from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Mstdatadictionaryfieldname , (select tablename from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarytablename,(select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Tableid,(select mstdatadictionarydbid from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarydbid,(select parentcategorynames from mstrecorddifferentiation where id = b.recorddiffid) as Parentcatnames,(select forrecorddifftypeid from mstworkdifferentiation where mainrecorddifftypeid=b.recorddifftypeid AND forrecorddifftypeid=b.recorddifftypeid AND forrecorddiffid=b.recorddiffid and activeflg = '1' AND deleteflg = 0) AS Forrecorddifftypeid,(select forrecorddiffid from mstworkdifferentiation where mainrecorddifftypeid=b.recorddifftypeid AND forrecorddifftypeid=b.recorddifftypeid AND forrecorddiffid=b.recorddiffid and activeflg = '1' AND deleteflg = 0) AS Forrecorddiffid from mstprocess a, mstprocessrecordmap b, mapprocesstoentity c where a.clientid = ? and a.activeflg = '1' and a.deleteflg=0 and a.id = b.mstprocessid and a.id = c.mstprocessid and b.activeflg = '1' and b.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstprocess = "select a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, b.recorddifftypeid as Recorddifftypeid, b.recorddiffid as Recorddiffid, b.id as Mstprocessrecordmapid, c.id as Mstprocesstoentityid, c.mstdatadictionaryfieldid as Mstdatadictionaryfieldid, a.activeflg as Activeflg, a.processname as Processname ,  (select name from mstorgnhierarchy where id = a.mstorgnhirarchyid) as Mstorgnhirarchyname , (select typename from mstrecorddifferentiationtype where id = b.recorddifftypeid) as Recorddifftypname , (select name from mstrecorddifferentiation where id = b.recorddiffid) as Recorddiffname , (select columnname from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Mstdatadictionaryfieldname , (select tablename from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarytablename,(select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid) as Tableid,(select mstdatadictionarydbid from mstdatadictionarytable where id = (select tableid from mstdatadictionaryfield where id = c.mstdatadictionaryfieldid)) as Mstdatadictionarydbid,(select parentcategorynames from mstrecorddifferentiation where id = b.recorddiffid) as Parentcatnames,(select forrecorddifftypeid from mstworkdifferentiation where mainrecorddifftypeid=b.recorddifftypeid AND forrecorddifftypeid=b.recorddifftypeid AND forrecorddiffid=b.recorddiffid and activeflg = '1' AND deleteflg = 0) AS Forrecorddifftypeid,(select forrecorddiffid from mstworkdifferentiation where mainrecorddifftypeid=b.recorddifftypeid AND forrecorddifftypeid=b.recorddifftypeid AND forrecorddiffid=b.recorddiffid and activeflg = '1' AND deleteflg = 0) AS Forrecorddiffid from mstprocess a, mstprocessrecordmap b, mapprocesstoentity c where a.clientid = ? and a.mstorgnhirarchyid = ? and a.activeflg = '1' and a.deleteflg=0 and a.id = b.mstprocessid and a.id = c.mstprocessid and b.activeflg = '1' and b.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstprocess, params...)

	// rows, err := dbc.DB.Query(getMstprocess, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstprocess Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstprocessEntity{}
		err = rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Mstprocessrecordmapid, &value.Mstprocesstoentityid, &value.Mstdatadictionaryfieldid, &value.Activeflg, &value.Processname, &value.Mstorgnhirarchyname, &value.Recorddifftypname, &value.Catname, &value.Mstdatadictionaryfieldname, &value.Mstdatadictionarytablename, &value.Tableid, &value.Mstdatadictionarydbid, &value.Parentcatname, &value.Forrecorddifftypeid, &value.Forrecorddiffid)
		//logger.Log.Println("In side GelAllMstprocess------>", err)
		if len(value.Parentcatname) > 0 {
			value.Recorddiffname = value.Catname + "(" + value.Parentcatname + ")"
		} else {
			value.Recorddiffname = value.Catname
		}
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) UpdateMstprocess(tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side UpdateMstprocess")
	stmt, err := dbc.DB.Prepare(updateMstprocess)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstprocess Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Processname, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstprocess Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateMstprocesswithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side UpdateMstprocess")
	stmt, err := tx.Prepare(updateMstprocess)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstprocess Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Processname, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstprocess Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstprocess(tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side DeleteMstprocess")
	stmt, err := dbc.DB.Prepare(deleteMstprocess)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstprocess Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstprocess Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteMstprocesswithtransaction(tx *sql.Tx, tz *entities.MstprocessEntity) error {
	logger.Log.Println("In side DeleteMstprocess")
	stmt, err := tx.Prepare(deleteMstprocess)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstprocess Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstprocess Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstprocessCount(tz *entities.MstprocessEntity, OrgnTypeID int64) (entities.MstprocessEntities, error) {
	logger.Log.Println("In side GetMstprocessCount")
	value := entities.MstprocessEntities{}
	var getMstprocesscount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstprocesscount = "select count(distinct a.id) as total from mstprocess a, mstprocessrecordmap b, mapprocesstoentity c where a.id = b.mstprocessid and a.id = c.mstprocessid and  a.activeflg = 1 and a.deleteflg=0 and b.activeflg = '1' and b.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getMstprocesscount = "select count(distinct a.id) as total from mstprocess a, mstprocessrecordmap b, mapprocesstoentity c where a.id = b.mstprocessid and a.id = c.mstprocessid and a.clientid = ? and a.activeflg = 1 and a.deleteflg=0 and b.activeflg = '1' and b.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0"
		params = append(params, tz.Clientid)
	} else {
		getMstprocesscount = "select count(distinct a.id) as total from mstprocess a, mstprocessrecordmap b, mapprocesstoentity c where a.id = b.mstprocessid and a.id = c.mstprocessid and a.clientid = ? and a.mstorgnhirarchyid = ? and a.activeflg = 1 and a.deleteflg=0 and b.activeflg = '1' and b.deleteflg=0 and c.activeflg = '1' and c.deleteflg=0"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstprocesscount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstprocesscount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstprocessCount Get Statement Prepare Error", err)
		return value, err
	}
}
