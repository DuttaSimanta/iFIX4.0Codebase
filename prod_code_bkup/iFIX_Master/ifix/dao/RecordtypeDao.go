package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertRecordtype = "INSERT INTO mstrecordtype (clientid, mstorgnhirarchyid, fromrecorddifftypeid, fromrecorddiffid, torecorddifftypeid, torecorddiffid) VALUES (?,?,?,?,?,?)"
var duplicateRecordtype = "SELECT count(id) total FROM  mstrecordtype WHERE clientid = ? AND mstorgnhirarchyid = ? AND fromrecorddifftypeid = ? AND fromrecorddiffid = ? AND torecorddifftypeid = ? AND torecorddiffid=? AND deleteflg = 0"
var getRecordtype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,COALESCE(g.name,'NA') AS Torecorddiffname FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id ORDER BY a.id DESC LIMIT ?,?"
var getRecordtypecount = "SELECT count(a.id) as total FROM mstrecordtype a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id"
var updateRecordtype = "UPDATE mstrecordtype SET mstorgnhirarchyid = ?, fromrecorddifftypeid = ?, fromrecorddiffid = ?, torecorddifftypeid = ?, torecorddiffid = ? WHERE id = ? "
var deleteRecordtype = "UPDATE mstrecordtype SET deleteflg = '1' WHERE id = ? "
var labelbydiffid = "SELECT distinct b.id , b.typename from mstrecordtype a,mstrecorddifferentiationtype b where a.torecorddifftypeid=b.id and a.clientid=? and a.mstorgnhirarchyid=? and a.fromrecorddifftypeid = ? AND a.fromrecorddiffid = ? AND a.activeflg=1 AND a.deleteflg =0 AND b.seqno >99 AND b.activeflg=1 AND b.deleteflg =0 order by b.seqno"
var labelbydiffseq = "SELECT distinct b.id , b.typename,b.seqno from mstrecordtype a,mstrecorddifferentiationtype b where a.torecorddifftypeid=b.id and a.clientid=? and a.mstorgnhirarchyid=? and a.fromrecorddifftypeid = ? AND a.fromrecorddiffid = ? AND a.activeflg=1 AND a.deleteflg =0 AND b.parentid in (select id from mstrecorddifferentiationtype where seqno=? and activeflg =1 and deleteflg=0) AND b.activeflg=1 AND b.deleteflg =0 order by b.typename"
var lablelmappingbydifftype = "SELECT distinct b.id , b.typename,b.seqno from mstrecordtype a,mstrecorddifferentiationtype b where a.torecorddifftypeid=b.id and a.clientid=? and a.mstorgnhirarchyid=? and a.fromrecorddifftypeid = ? AND a.fromrecorddiffid = ? AND a.torecorddifftypeid IN (SELECT id from mstrecorddifferentiationtype where parentid= ? and activeflg=1 AND deleteflg =0) AND a.activeflg=1 AND a.deleteflg =0  AND b.activeflg=1 AND b.deleteflg =0 "
var parentid = "select id from mstrecorddifferentiationtype where seqno=?"

//var getstateidbyfromdiff="select b.mststateid from mstrecordtype a,maprecordstatetodifferentiation b where a.fromrecorddiffid =? and a.torecorddiffid=b.recorddiffid and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0"
var getstateidbyfromdiff = "select d.mststateid from mstrecordtype a ,mstrecordtype b,mstrecordtype c,maprecordstatetodifferentiation d WHERE a.fromrecorddifftypeid = 2  AND a.fromrecorddiffid = ? AND a.torecorddifftypeid=3 AND a.torecorddiffid=? AND a.torecorddiffid = b.fromrecorddiffid AND a.activeflg=1 AND a.deleteflg=0 AND c.fromrecorddifftypeid = 2 AND c.fromrecorddiffid = ? AND c.torecorddiffid = b.torecorddiffid AND c.torecorddiffid = d.recorddiffid AND c.activeflg=1 AND c.deleteflg=0 AND b.activeflg=1 AND b.deleteflg=0 AND d.activeflg=1 AND d.deleteflg=0;"
var mappeddiff = "select a.torecorddifftypeid,a.torecorddiffid,b.name,b.seqno,coalesce(b.parentcategorynames,'') parentpath from mstrecordtype a,mstrecorddifferentiation b where a.clientid=? and a.mstorgnhirarchyid=? and a.fromrecorddifftypeid=? and a.fromrecorddiffid=? and a.torecorddifftypeid in (select id from mstrecorddifferentiationtype where seqno=? and activeflg=1 and deleteflg=0) and a.torecorddiffid = b.id and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0"

func (mdao DbConn) Getstateidbyfromdiff(diffid int64, parentmappingdiffid int64, childmappingdiffid int64) ([]entities.Recordtypesingleentity, error) {
	logger.Log.Print("Getstateidbyfromdiff : ", diffid, parentmappingdiffid, childmappingdiffid)
	stateids := []entities.Recordtypesingleentity{}
	stmt, err := mdao.DB.Prepare(getstateidbyfromdiff)
	if err != nil {
		logger.Log.Print("Getstateidbyfromdiff Statement Prepare Error", err)
		log.Print("Getstateidbyfromdiff Statement Prepare Error", err)
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(parentmappingdiffid, diffid, childmappingdiffid)
	if err != nil {
		logger.Log.Print("Getstateidbyfromdiff Statement Execution Error", err)
		log.Print("Getstateidbyfromdiff Statement Execution Error", err)
		return nil, err
	}
	for rows.Next() {
		stateid := entities.Recordtypesingleentity{}
		rows.Scan(&stateid.Id)
		stateids = append(stateids, stateid)
	}
	return stateids, nil
}
func (dbc DbConn) Getlablelmappingbydifftype(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, error) {
	logger.Log.Println("In side Getlablelmappingbydifftype")
	//logger.Log.Println(getRecordtype)
	values := []entities.Recordtypesingleentity{}
	rows, err := dbc.DB.Query(lablelmappingbydifftype, page.Clientid, page.Mstorgnhirarchyid, page.Fromrecorddifftypeid, page.Fromrecorddiffid, page.Parentid)

	if err != nil {
		logger.Log.Println("Getlablelmappingbydifftype Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Recordtypesingleentity{}
		rows.Scan(&value.Id, &value.Typename, &value.Seqno)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getparentid(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, error) {
	logger.Log.Println("In side Getparentid")
	//logger.Log.Println(getRecordtype)
	values := []entities.Recordtypesingleentity{}
	rows, err := dbc.DB.Query(parentid, page.Seqno)

	if err != nil {
		logger.Log.Println("Getparentid Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Recordtypesingleentity{}
		rows.Scan(&value.Id)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getlabelbydiffid(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, error) {
	logger.Log.Println("In side GelAllRecordtype")
	values := []entities.Recordtypesingleentity{}
	rows, err := dbc.DB.Query(labelbydiffid, page.Clientid, page.Mstorgnhirarchyid, page.Fromrecorddifftypeid, page.Fromrecorddiffid)

	if err != nil {
		logger.Log.Println("Getlabelbydiffid Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Recordtypesingleentity{}
		rows.Scan(&value.Id, &value.Typename)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getlabelbydiffseq(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, error) {
	logger.Log.Println("In side GelAllRecordtype")
	values := []entities.Recordtypesingleentity{}
	rows, err := dbc.DB.Query(labelbydiffseq, page.Clientid, page.Mstorgnhirarchyid, page.Fromrecorddifftypeid, page.Fromrecorddiffid, page.Seqno)

	if err != nil {
		logger.Log.Println("Getlabelbydiffid Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Recordtypesingleentity{}
		rows.Scan(&value.Id, &value.Typename, &value.Seqno)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) Getmappeddiffbyseq(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, error) {
	logger.Log.Println("In side GelAllRecordtype")
	//logger.Log.Println(getRecordtype)
	values := []entities.Recordtypesingleentity{}
	rows, err := dbc.DB.Query(mappeddiff, page.Clientid, page.Mstorgnhirarchyid, page.Fromrecorddifftypeid, page.Fromrecorddiffid, page.Seqno)

	if err != nil {
		logger.Log.Println("Getmappeddiffbyseq Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.Recordtypesingleentity{}
		rows.Scan(&value.Recorddifftypeid, &value.Id, &value.Typename,&value.Seqno, &value.Parentpath)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateRecordtype(tz *entities.RecordtypeEntity) (entities.RecordtypeEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecordtype")
	value := entities.RecordtypeEntities{}
	err := dbc.DB.QueryRow(duplicateRecordtype, tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecordtype Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertRecordtype(tz *entities.RecordtypeEntity) (int64, error) {
	logger.Log.Println("In side InsertRecordtype")
	//logger.Log.Println("Query -->", insertRecordtype)
	stmt, err := dbc.DB.Prepare(insertRecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordtype Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid)
	if err != nil {
		logger.Log.Println("InsertRecordtype Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func InsertRecordtypetran(tx *sql.Tx, tz *entities.RecordtypeEntity, i int) (int64, error) {
	logger.Log.Println("In side InsertRecordtype")
	//logger.Log.Println("Query -->", insertRecordtype)
	stmt, err := tx.Prepare(insertRecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordtype Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffids[i], tz.Torecorddifftypeid, tz.Torecorddiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffids[i], tz.Torecorddifftypeid, tz.Torecorddiffid)
	if err != nil {
		logger.Log.Println("InsertRecordtype Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc DbConn) GetAllRecordtype(page *entities.RecordtypeEntity, OrgnType int64) ([]entities.RecordtypeEntity, error) {
	values := []entities.RecordtypeEntity{}
	var getRecordtype string
	var params []interface{}
	if OrgnType == 1 {
		getRecordtype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,COALESCE(g.name,'NA') AS Torecorddiffname FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getRecordtype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,COALESCE(g.name,'NA') AS Torecorddiffname FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getRecordtype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,COALESCE(g.name,'NA') AS Torecorddiffname FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getRecordtype, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecordtype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordtypeEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Fromrecorddifftypeid, &value.Fromrecorddiffid, &value.Torecorddifftypeid, &value.Torecorddiffid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Fromrecorddifftypename, &value.Fromrecorddiffname, &value.Torecorddifftypename, &value.Torecorddiffname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateRecordtype(tz *entities.RecordtypeEntity) error {
	logger.Log.Println("In side UpdateRecordtype")
	stmt, err := dbc.DB.Prepare(updateRecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecordtype Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteRecordtype(tz *entities.RecordtypeEntity) error {
	logger.Log.Println("In side DeleteRecordtype")
	stmt, err := dbc.DB.Prepare(deleteRecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecordtype Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetRecordtypeCount(tz *entities.RecordtypeEntity, OrgnTypeID int64) (entities.RecordtypeEntities, error) {
	logger.Log.Println("In side GetRecordtypeCount")
	value := entities.RecordtypeEntities{}
	var getRecordtypecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getRecordtypecount = "SELECT count(a.id) as total FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id "
	} else if OrgnTypeID == 2 {
		getRecordtypecount = "SELECT count(a.id) as total FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id "
		params = append(params, tz.Clientid)
	} else {
		getRecordtypecount = "SELECT count(a.id) as total FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getRecordtypecount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetRecordtypeCount Get Statement Prepare Error", err)
		return value, err
	}
}

func InsertRecordtype(tx *sql.Tx, tz *entities.RecordtypeEntity) (int64, error) {
	logger.Log.Println("In side InsertRecordtype")
	//logger.Log.Println("Query -->", insertRecordtype)
	stmt, err := tx.Prepare(insertRecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordtype Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid)
	if err != nil {
		logger.Log.Println("InsertRecordtype Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func Insertextrafieldvalue(tx *sql.Tx, Clientid int64, Mstorgnhirarchyid int64, id int64, Title string, Description string) (int64, error) {
	var sql = "INSERT INTO msttaskpropertyvalue(clientid,mstorgnhirarchyid,recordtypeid,title,description) VALUES(?,?,?,?,?)"
	stmt, err := tx.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordtype Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(Clientid, Mstorgnhirarchyid, id, Title, Description)
	if err != nil {
		logger.Log.Println("InsertRecordtype Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllTaskmapvalues(page *entities.RecordtypeEntity, OrgnType int64) ([]entities.RecordtypeEntity, error) {
	values := []entities.RecordtypeEntity{}
	var getRecordtype string
	var params []interface{}
	if OrgnType == 1 {
		getRecordtype = " SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,COALESCE(g.name,'NA') AS Torecorddiffname,COALESCE(h.title,''),COALESCE(h.description,'') FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,msttaskpropertyvalue h,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id and g.deleteflg=0 where AND h.recordtypeid=a.id AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and d.deleteflg=0 and e.deleteflg=0 and f.deleteflg=0 and h.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getRecordtype = " SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,COALESCE(g.name,'NA') AS Torecorddiffname,COALESCE(h.title,''),COALESCE(h.description,'') FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,msttaskpropertyvalue h,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id and g.deleteflg=0 where h.clientid = ? AND h.recordtypeid=a.id AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and d.deleteflg=0 and e.deleteflg=0 and f.deleteflg=0 and h.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getRecordtype = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.fromrecorddifftypeid as Fromrecorddifftypeid, a.fromrecorddiffid as Fromrecorddiffid, a.torecorddifftypeid as Torecorddifftypeid, a.torecorddiffid as Torecorddiffid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Fromrecorddifftypename,e.name as Fromrecorddiffname,f.typename as Torecorddifftypename,COALESCE(g.name,'NA') AS Torecorddiffname,COALESCE(h.title,''),COALESCE(h.description,'') FROM mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,msttaskpropertyvalue h,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id and g.deleteflg=0 where h.clientid = ? AND h.mstorgnhirarchyid = ? AND h.recordtypeid=a.id AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and d.deleteflg=0 and e.deleteflg=0 and f.deleteflg=0 and h.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getRecordtype, params...)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecordtype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecordtypeEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Fromrecorddifftypeid, &value.Fromrecorddiffid, &value.Torecorddifftypeid, &value.Torecorddiffid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Fromrecorddifftypename, &value.Fromrecorddiffname, &value.Torecorddifftypename, &value.Torecorddiffname, &value.Title, &value.Description)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GettaskmapCount(tz *entities.RecordtypeEntity, OrgnTypeID int64) (entities.RecordtypeEntities, error) {
	logger.Log.Println("In side GetRecordtypeCount")
	value := entities.RecordtypeEntities{}
	var getRecordtypecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getRecordtypecount = "SELECT count(a.id) as total FROM  mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,msttaskpropertyvalue h,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id and g.deleteflg=0 where h.recordtypeid=a.id AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and d.deleteflg=0 and e.deleteflg=0 and f.deleteflg=0 and h.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getRecordtypecount = "SELECT count(a.id) as total FROM  mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,msttaskpropertyvalue h,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id and g.deleteflg=0 where h.clientid = ? AND h.recordtypeid=a.id AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and d.deleteflg=0 and e.deleteflg=0 and f.deleteflg=0 and h.deleteflg=0"

		params = append(params, tz.Clientid)
	} else {
		// getRecordtypecount = "SELECT count(a.id) as total FROM mstrecordtype a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,mstrecorddifferentiation g,msttaskpropertyvalue h WHERE h.clientid = ? AND h.mstorgnhirarchyid = ? AND h.recordtypeid=a.id AND h.deleteflg =0 and h.activeflg=1 AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and a.torecorddiffid=g.id"
		getRecordtypecount = "SELECT count(a.id) as total FROM  mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f,msttaskpropertyvalue h,mstrecordtype a LEFT JOIN mstrecorddifferentiation g ON a.torecorddiffid = g.id and g.deleteflg=0 where h.clientid = ? AND h.mstorgnhirarchyid = ? AND h.recordtypeid=a.id AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.fromrecorddifftypeid=d.id and a.fromrecorddiffid=e.id and a.torecorddifftypeid=f.id and d.deleteflg=0 and e.deleteflg=0 and f.deleteflg=0 and h.deleteflg=0"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getRecordtypecount, params...).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetRecordtypeCount Get Statement Prepare Error", err)
		return value, err
	}
}

func DeleteTaskmap(tx *sql.Tx, tz *entities.RecordtypeEntity) error {
	logger.Log.Println("In side DeleteTaskmap")
	var sql = "UPDATE mstrecordtype SET deleteflg = '1' WHERE id = ? "
	stmt, err := tx.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecordtype Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteTaskProperty(tx *sql.Tx, tz *entities.RecordtypeEntity) error {
	logger.Log.Println("In side DeleteTaskmap")
	var sql = "UPDATE msttaskpropertyvalue SET deleteflg = '1' WHERE recordtypeid = ? "
	stmt, err := tx.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecordtype Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) CheckDuplicatetaskproperty(tz *entities.RecordtypeEntity) (entities.RecordtypeEntities, error) {
	logger.Log.Println("In side CheckDuplicatetaskproperty")
	var query = "SELECT COUNT(id) as total FROM msttaskpropertyvalue WHERE clientid=? AND mstorgnhirarchyid=? AND recordtypeid=? AND title=? AND description=?"
	value := entities.RecordtypeEntities{}
	err := dbc.DB.QueryRow(query, tz.Clientid, tz.Mstorgnhirarchyid, tz.Id, tz.Title, tz.Description).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicatetaskproperty Get Statement Prepare Error", err)
		return value, err
	}
}

func UpdateRecordtype(tx *sql.Tx, tz *entities.RecordtypeEntity) error {
	logger.Log.Println("In side UpdateRecordtype")
	stmt, err := tx.Prepare(updateRecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Fromrecorddifftypeid, tz.Fromrecorddiffid, tz.Torecorddifftypeid, tz.Torecorddiffid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecordtype Execute Statement  Error", err)
		return err
	}
	return nil
}

func UpdateTaskFieldValue(tx *sql.Tx, tz *entities.RecordtypeEntity) error {
	logger.Log.Println("In side UpdateTaskFieldValue")
	var sql = "UPDATE msttaskpropertyvalue SET title=?,description=? WHERE clientid=? AND mstorgnhirarchyid=? AND recordtypeid=? "
	stmt, err := tx.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordtype Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Title, tz.Description, tz.Clientid, tz.Mstorgnhirarchyid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecordtype Execute Statement  Error", err)
		return err
	}
	return nil
}
