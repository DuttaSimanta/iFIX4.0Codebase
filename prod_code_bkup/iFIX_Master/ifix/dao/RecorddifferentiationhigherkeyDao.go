package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var duplicaterecorddifferentiationdata = "SELECT count(id) as total FROM  mstrecorddifferentiation WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND parentid=? AND name=? AND deleteflg = 0 AND activeflg=1"
var duplicatemstrecordtypedata = "SELECT count(id) as recordtypetotal FROM  mstrecordtype WHERE clientid = ? AND mstorgnhirarchyid = ? AND fromrecorddifftypeid = ? AND fromrecorddiffid=? AND torecorddifftypeid=? AND torecorddiffid=? AND deleteflg = 0 AND activeflg=1"
var insertrecorddifferentiationdata = "INSERT INTO mstrecorddifferentiation (clientid, mstorgnhirarchyid, recorddifftypeid, parentid, name, seqno) VALUES (?,?,?,?,?,?)"
var insertmstecordtypedata = "INSERT INTO mstrecordtype (clientid, mstorgnhirarchyid, fromrecorddifftypeid, fromrecorddiffid, torecorddifftypeid, torecorddiffid) VALUES (?,?,?,?,?,?)"
var deleterecorddifferentiation = "UPDATE mstrecorddifferentiation SET deleteflg=1 WHERE id=?"
var deletemstrecordtype = "UPDATE mstrecordtype SET deleteflg=1 WHERE torecorddiffid=?"

// var selecthigherkey = "SELECT f.id as ID,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, e.id as Parentrecorddifftypeid,d.id as Parentrecorddiffid,g.id as Childrecorddifftypeid,f.id as Childrecorddiffid,b.name as Clientname,c.name as Mstorgnhirarchyname,e.typename as Parentrecorddifftypename,d.name as Parentrecorddiffname,g.typename as Childrecorddifftypename,f.name as Childrecorddiffname,f.seqno as Seqno,f.parentid as Parentid,f.activeflg as Activeflg,f.parentcategorynames as Parentcatnames FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and  a.fromrecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=1) and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0' ORDER BY f.id DESC LIMIT ?,?"

//var selecthigherkey = "SELECT f.id as ID,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, e.id as Parentrecorddifftypeid,d.id as Parentrecorddiffid,g.id as Childrecorddifftypeid,f.id as Childrecorddiffid,b.name as Clientname,c.name as Mstorgnhirarchyname,e.typename as Parentrecorddifftypename,d.name as Parentrecorddiffname,g.typename as Childrecorddifftypename,f.name as Childrecorddiffname,f.seqno as Seqno,f.parentid as Parentid,f.activeflg as Activeflg,f.parentcategorynames as Parentcatnames FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0' ORDER BY f.id DESC LIMIT ?,?"
//var counthigherkey = "SELECT count(f.id) as total FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid  in (SELECT b.id FROM mstrecorddifferentiationtype a,  mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0'"
// var counthigherkey = "SELECT count(f.id) as total FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and  a.fromrecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=1) and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0'"
var updaterecorddifferentiation = "Update mstrecorddifferentiation set mstorgnhirarchyid=?,recorddifftypeid=?,parentid=?,name=?,seqno=? where id=?"
var updatemstrecordtype = "UPDATE mstrecordtype set mstorgnhirarchyid=?,fromrecorddifftypeid=?,fromrecorddiffid=?,torecorddifftypeid=?,torecorddiffid=? where torecorddiffid=?"

//CheckDuplicateRecorddifferentiationkey is used for check duplicate data
func CheckDuplicateRecorddifferentiationkey(tx *sql.Tx, tz *entities.RecorddifferentiationhigherkeyEntity) (int64, error) {
	logger.Log.Println("In side CheckDuplicateRecorddifferentiationkey")
	var total int64
	stmt, error := tx.Prepare(duplicaterecorddifferentiationdata)
	if error != nil {
		logger.Log.Println("Exception in CheckDuplicateRecorddifferentiationkey Prepare Statement..")
		return total, error
	}

	rows, err := stmt.Query(tz.Clientid, tz.Mstorgnhirarchyid, tz.Childrecorddifftypeid, tz.Parentid, tz.Name)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateRecorddifferentiationkey Query Statement..")
		return total, error
	}

	for rows.Next() {
		if err := rows.Scan(&total); err != nil {
			logger.Log.Println("Error in fetching data")
		}

	}
	fmt.Println("recorddifferentiation count value is :", total)
	return total, nil
}

//CheckDuplicateMstrecordtypedata is used for check duplicate data
func CheckDuplicateMstrecordtypedata(tx *sql.Tx, tz *entities.RecorddifferentiationhigherkeyEntity, Childrecorddiffid int64) (int64, error) {
	logger.Log.Println("In side CheckDuplicateMstrecordtypedata")
	var recordtypetotal int64
	stmt, error := tx.Prepare(duplicatemstrecordtypedata)
	if error != nil {
		logger.Log.Println("Exception in CheckDuplicateMstrecordtypedata Prepare Statement..")
		return recordtypetotal, error
	}

	rows, err := stmt.Query(tz.Clientid, tz.Mstorgnhirarchyid, tz.Parentrecorddifftypeid, tz.Parentrecorddiffid, tz.Childrecorddifftypeid, Childrecorddiffid)
	if err != nil {
		logger.Log.Println("Exception in CheckDuplicateMstrecordtypedata Query Statement..")
		return recordtypetotal, error
	}

	for rows.Next() {
		if err := rows.Scan(&recordtypetotal); err != nil {
			logger.Log.Println("Error in fetching data")
		}

	}
	fmt.Println("mstrecordtype total count value is :", recordtypetotal)
	return recordtypetotal, nil
}

//InsertRecorddifferentiation is used for insert data into differentiation table.
func InsertRecorddifferentiation(tx *sql.Tx, tz *entities.RecorddifferentiationhigherkeyEntity) (int64, error) {
	stmt, err := tx.Prepare(insertrecorddifferentiationdata)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Childrecorddifftypeid, tz.Parentid, tz.Name, tz.Seqno)
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Execution Error")
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}

	return lastInsertedID, nil
}

//InsertMstRecordtype is used for insert data into recordtype table.
func InsertMstRecordtype(tx *sql.Tx, tz *entities.RecorddifferentiationhigherkeyEntity, Childrecorddiffid int64) (int64, error) {
	stmt, err := tx.Prepare(insertmstecordtypedata)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Parentrecorddifftypeid, tz.Parentrecorddiffid, tz.Childrecorddifftypeid, Childrecorddiffid)
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Execution Error")
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}

	return lastInsertedID, nil
}

//DeleteRecorddifferentiation is used for delete data into differentiation table
func DeleteRecorddifferentiation(tx *sql.Tx, tz *entities.RecorddifferentiationhigherkeyEntity) error {
	logger.Log.Println("In side DeleteRecorddifferentiation")
	stmt, err := tx.Prepare(deleterecorddifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}

	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}

//DeleteMstrecordtype is used for delete data into mstrecordtype table
func DeleteMstrecordtype(tx *sql.Tx, tz *entities.RecorddifferentiationhigherkeyEntity) error {
	logger.Log.Println("In side DeleteMstrecordtype")
	stmt, err := tx.Prepare(deletemstrecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}

	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}

func (dbc DbConn) GetAllRecorddifferentiationHighkey(tz *entities.RecorddifferentiationhigherkeyEntity, OrgnType int64) ([]entities.RecorddifferentiationhigherkeyEntity, error) {
	logger.Log.Println("In side GetAllRecorddifferentiationHighkey")
	values := []entities.RecorddifferentiationhigherkeyEntity{}
	var selecthigherkey string
	var params []interface{}
	if OrgnType == 1 {
		selecthigherkey = "SELECT f.id as ID,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, e.id as Parentrecorddifftypeid,d.id as Parentrecorddiffid,g.id as Childrecorddifftypeid,f.id as Childrecorddiffid,b.name as Clientname,c.name as Mstorgnhirarchyname,e.typename as Parentrecorddifftypename,d.name as Parentrecorddiffname,g.typename as Childrecorddifftypename,f.name as Childrecorddiffname,f.seqno as Seqno,f.parentid as Parentid,f.activeflg as Activeflg,f.parentcategorynames as Parentcatnames FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where  a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and  a.fromrecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=1) and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0' ORDER BY f.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		selecthigherkey = "SELECT f.id as ID,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, e.id as Parentrecorddifftypeid,d.id as Parentrecorddiffid,g.id as Childrecorddifftypeid,f.id as Childrecorddiffid,b.name as Clientname,c.name as Mstorgnhirarchyname,e.typename as Parentrecorddifftypename,d.name as Parentrecorddiffname,g.typename as Childrecorddifftypename,f.name as Childrecorddiffname,f.seqno as Seqno,f.parentid as Parentid,f.activeflg as Activeflg,f.parentcategorynames as Parentcatnames FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and  a.fromrecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=1) and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0' ORDER BY f.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		selecthigherkey = "SELECT f.id as ID,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid, e.id as Parentrecorddifftypeid,d.id as Parentrecorddiffid,g.id as Childrecorddifftypeid,f.id as Childrecorddiffid,b.name as Clientname,c.name as Mstorgnhirarchyname,e.typename as Parentrecorddifftypename,d.name as Parentrecorddiffname,g.typename as Childrecorddifftypename,f.name as Childrecorddiffname,f.seqno as Seqno,f.parentid as Parentid,f.activeflg as Activeflg,f.parentcategorynames as Parentcatnames FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and  a.fromrecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=1) and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0' ORDER BY f.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	logger.Log.Println(selecthigherkey)

	rows, err := dbc.DB.Query(selecthigherkey, params...)
	// rows, err := dbc.DB.Query(selecthigherkey, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecorddifferentiationhigherkeyEntity{}
		err = rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Parentrecorddifftypeid, &value.Parentrecorddiffid, &value.Childrecorddifftypeid, &value.Childrecorddiffid, &value.Clientname, &value.Mstorgnhirarchyname, &value.Parentrecorddifftypename, &value.Parentrecorddiffname, &value.Childrecorddifftypename, &value.Catename, &value.Seqno, &value.Parentid, &value.Activeflg, &value.Parentcatname)
		//logger.Log.Println(err)
		if len(value.Parentcatname) > 0 {
			value.Childrecorddiffname = value.Catename + "(" + value.Parentcatname + ")"
		} else {
			value.Childrecorddiffname = value.Catename
		}

		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetRecorddifferentiationhighkeyCount(tz *entities.RecorddifferentiationhigherkeyEntity, OrgnTypeID int64) (entities.RecorddifferentiationhigherkeyEntities, error) {
	logger.Log.Println("In side GetRecorddifferentiationhighkeyCount")
	value := entities.RecorddifferentiationhigherkeyEntities{}
	var counthigherkey string
	var params []interface{}
	if OrgnTypeID == 1 {
		counthigherkey = "SELECT count(f.id) as total FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and  a.fromrecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=1) and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0'"
	} else if OrgnTypeID == 2 {
		counthigherkey = "SELECT count(f.id) as total FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and  a.fromrecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=1) and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0'"
		params = append(params, tz.Clientid)
	} else {
		counthigherkey = "SELECT count(f.id) as total FROM mstrecordtype a, mstclient b, mstorgnhierarchy c, mstrecorddifferentiation d, mstrecorddifferentiationtype e, mstrecorddifferentiation f, mstrecorddifferentiationtype g where a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid = b.id and a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid = e.id and  a.fromrecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=1) and a.fromrecorddiffid = d.id and a.torecorddifftypeid = g.id and a.torecorddiffid = f.id and a.torecorddifftypeid in (SELECT b.id FROM mstrecorddifferentiationtype a, mstrecorddifferentiationtype b where a.seqno = 0 and a.activeflg = '1' and a.deleteflg = '0' and a.id = b.parentid and b.activeflg = '1' and b.deleteflg = '0') and d.activeflg = '1' and d.deleteflg = '0' and e.activeflg = '1' and e.deleteflg = '0' and f.activeflg = '1' and f.deleteflg = '0' and g.activeflg = '1' and g.deleteflg = '0'"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(counthigherkey, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(counthigherkey, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetRecorddifferentiationhighkeyCount Get Statement Prepare Error", err)
		return value, err
	}
}

func UpdateRecorddifferentiationData(tx *sql.Tx, tz *entities.RecorddifferentiationhigherkeyEntity) error {
	logger.Log.Println("In side UpdateRecorddifferentiationData")
	stmt, err := tx.Prepare(updaterecorddifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}

	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Childrecorddifftypeid, tz.Parentid, tz.Name, tz.Seqno, tz.ID)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}

func UpdateMstrecordtypeData(tx *sql.Tx, tz *entities.RecorddifferentiationhigherkeyEntity) error {
	logger.Log.Println("In side UpdateRecorddifferentiationData")
	stmt, err := tx.Prepare(updatemstrecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}

	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Parentrecorddifftypeid, tz.Parentrecorddiffid, tz.Childrecorddifftypeid, tz.Childrecorddiffid, tz.Childrecorddiffid)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}

func (dbc DbConn) UpdateParentPath(parentID string, parentName string, lastinsertedID int64) error {
	logger.Log.Println("In side UpdateRecorddifferentiationData")
	var sql = "UPDATE mstrecorddifferentiation SET parentcategoryids=?,parentcategorynames=? WHERE id=?"
	stmt, err := dbc.DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}

	_, err = stmt.Exec(parentID, parentName, lastinsertedID)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}

func (mdao DbConn) GetParentdetails(tz *entities.RecorddifferentiationhigherkeyEntity) (string, string, error) {
	var parentcategoryids string
	var parentcategorynames string
	var sql = "SELECT parentcategoryids,parentcategorynames FROM mstrecorddifferentiation WHERE id=? AND clientid=? AND mstorgnhirarchyid=?"
	rows, err := mdao.DB.Query(sql, tz.Parentid, tz.Clientid, tz.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAdditional Get Statement Prepare Error", err)
		return parentcategoryids, parentcategorynames, err
	}
	for rows.Next() {
		err = rows.Scan(&parentcategoryids, &parentcategorynames)
		logger.Log.Println("GetParentdetails rows.next() Error", err)
	}
	return parentcategoryids, parentcategorynames, nil
}

func (mdao DbConn) GetParentnamesbyID(ID int64) (string, error) {
	var parentcategorynames string
	var sql = "SELECT name FROM mstrecorddifferentiation WHERE id=? " //in (SELECT parentid FROM mstrecorddifferentiation WHERE id=?)
	rows, err := mdao.DB.Query(sql, ID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAdditional Get Statement Prepare Error", err)
		return parentcategorynames, err
	}
	for rows.Next() {
		err = rows.Scan(&parentcategorynames)
		logger.Log.Println("GetParentdetails rows.next() Error", err)
	}
	return parentcategorynames, nil
}
