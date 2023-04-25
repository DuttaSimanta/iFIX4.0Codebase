package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertRecorddifferentiation = "INSERT INTO mstrecorddifferentiation (clientid, mstorgnhirarchyid, recorddifftypeid, parentid, name, seqno) VALUES (?,?,?,?,?,?)"
var duplicateRecorddifferentiation = "SELECT count(id) total FROM  mstrecorddifferentiation WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND deleteflg = 0"
var duplicateParentcategory = "SELECT count(id) as total FROM  mstrecorddifferentiation WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND parentid=? AND name=? AND deleteflg = 0 AND activeflg=1"

var getRecorddifferentiation = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.recorddifftypeid as recorddifftypeid, a.parentid as Parentid, a.name as Name, a.seqno as seqno, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifftypname FROM mstrecorddifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND d.id=a.recorddifftypeid AND d.parentid in (select id from mstrecorddifferentiationtype where seqno=? and activeflg=1 and deleteflg=0 ) ORDER BY a.id DESC LIMIT ?,?"
var getRecorddifferentiationcount = "SELECT count(a.id) as total FROM mstrecorddifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? and a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND d.id=a.recorddifftypeid AND d.parentid in (select id from mstrecorddifferentiationtype where seqno=? and activeflg=1 and deleteflg=0 )"

//var updateRecorddifferentiation = "UPDATE mstrecorddifferentiation SET mstorgnhirarchyid = ?, recorddifftypeid = ?, parentid = ?, name = ?, seqno = ? WHERE id = ? "
var updateRecorddifferentiation = "UPDATE mstrecorddifferentiation SET mstorgnhirarchyid = ?, name = ? WHERE id = ? "
var deleteRecorddifferentiation = "UPDATE mstrecorddifferentiation SET deleteflg = '1' WHERE id = ? "
var getdifferentiationname = "select id as Id,name as Recorddiffname,COALESCE(parentcategorynames,'') AS Parentcatname from mstrecorddifferentiation where clientid =? and mstorgnhirarchyid=? and name like ? and recorddifftypeid=? and deleteflg=0 and activeflg=1"
var getdiffvaluebyparentid = "SELECT id,recorddifftypeid,parentid,name,parentcategorynames from mstrecorddifferentiation where id=? and activeflg=1 and deleteflg=0 "

//var categorysearch = "SELECT id,parentcategorynames from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid IN (select id from mstrecorddifferentiationtype where clientid=? and mstorgnhirarchyid=? and parentid=? and activeflg=1 and deleteflg=0) and parentcategorynames like ? and activeflg=1 and deleteflg=0 limit 15"
var categorysearch = "SELECT a.id,a.parentcategorynames from mstrecorddifferentiation a,mapcatalogwithcategory b where a.clientid=? and a.mstorgnhirarchyid=? and a.recorddifftypeid IN (select id from mstrecorddifferentiationtype where clientid=? and mstorgnhirarchyid=? and parentid=? and activeflg=1 and deleteflg=0) and a.parentcategorynames like ? and a.activeflg=1 and a.deleteflg=0 and a.id = b.torecorddiffid AND b.activeflg=1 and b.deleteflg=0 limit 15"
var getdiffdetails = "SELECT recorddifftypeid,id from mstrecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recorddifftypeid in (SELECT id from mstrecorddifferentiationtype where seqno=? and activeflg=1 and deleteflg=0 )  and seqno=? and  activeflg=1 and deleteflg=0"

func (mdao DbConn) Getdiffdetailsbyseq(tz *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentionSingle, error) {
	taskcatids := []entities.RecorddifferentionSingle{}
	stmt, error := mdao.DB.Prepare(getdiffdetails)
	if error != nil {
		logger.Log.Println("Exception in Getdiffdetailsbyseq Prepare Statement..")
		return nil, error
	}
	defer stmt.Close()
	rows, err := stmt.Query(tz.Clientid, tz.Mstorgnhirarchyid, tz.Typeseqno, tz.Seqno)
	if err != nil {
		logger.Log.Println("Exception in Getdiffdetailsbyseq Query Statement..")
		return nil, error
	}
	for rows.Next() {
		taskcatid := entities.RecorddifferentionSingle{}
		if err := rows.Scan(&taskcatid.Recorddifftypeid, &taskcatid.Id); err != nil {
			logger.Log.Println("Error in fetching data", err)
		}
		taskcatids = append(taskcatids, taskcatid)
	}
	return taskcatids, nil
}
func (dbc DbConn) Searchcategory(page *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentionSingle, error) {
	logger.Log.Println("In side Searchcategory")
	values := []entities.RecorddifferentionSingle{}
	rows, err := dbc.DB.Query(categorysearch, page.Clientid, page.Mstorgnhirarchyid, page.Clientid, page.Mstorgnhirarchyid, page.Recorddifftypeid, "%"+page.Name+"%")
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Searchcategory Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecorddifferentionSingle{}
		rows.Scan(&value.Id, &value.Parentcategorynames)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetRecorddifferentiationbyparent(page *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentionSingle, error) {
	logger.Log.Println("In side GelAllRecorddifferentiation")
	logger.Log.Println(getdiffvaluebyparentid)
	values := []entities.RecorddifferentionSingle{}
	rows, err := dbc.DB.Query(getdiffvaluebyparentid, page.Id)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetRecorddifferentiationbyparent Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecorddifferentionSingle{}
		rows.Scan(&value.Id, &value.Recorddifftypeid, &value.Parentid, &value.Name, &value.Parentcategorynames)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) CheckDuplicateRecorddifferentiation(tz *entities.RecorddifferentiationEntity) (entities.RecorddifferentiationEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecorddifferentiation")
	value := entities.RecorddifferentiationEntities{}
	err := dbc.DB.QueryRow(duplicateRecorddifferentiation, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecorddifferentiation Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) CheckDuplicateParentcategoriesforupdate(tz *entities.RecorddifferentiationEntity, newparentcategorynames string) (entities.RecorddifferentiationEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecorddifferentiation")
	value := entities.RecorddifferentiationEntities{}
	err := dbc.DB.QueryRow(duplicateParentcategory, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Parentcatagorytypeid, tz.Name).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecorddifferentiation Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertRecorddifferentiation(tz *entities.RecorddifferentiationEntity) (int64, error) {
	logger.Log.Println("In side InsertRecorddifferentiation")
	logger.Log.Println("Query -->", insertRecorddifferentiation)
	stmt, err := dbc.DB.Prepare(insertRecorddifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecorddifferentiation Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Parentid, tz.Name, tz.Seqno)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Parentid, tz.Name, tz.Seqno)
	if err != nil {
		logger.Log.Println("InsertRecorddifferentiation Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllRecorddifferentiation(page *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentiationEntity, error) {
	logger.Log.Println("In side GelAllRecorddifferentiation")
	logger.Log.Println(getRecorddifferentiation)
	values := []entities.RecorddifferentiationEntity{}
	rows, err := dbc.DB.Query(getRecorddifferentiation, page.Clientid, page.Mstorgnhirarchyid, page.Seqno, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecorddifferentiationEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Recorddifftypeid, &value.Parentid, &value.Name, &value.Seqno, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifftypname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateRecorddifferentiation(tz *entities.RecorddifferentiationEntity) error {
	logger.Log.Println("In side UpdateRecorddifferentiation")
	stmt, err := dbc.DB.Prepare(updateRecorddifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecorddifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Name, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateRecorddifferentiation Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteRecorddifferentiation(tz *entities.RecorddifferentiationEntity) error {
	logger.Log.Println("In side DeleteRecorddifferentiation")
	stmt, err := dbc.DB.Prepare(deleteRecorddifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecorddifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecorddifferentiation Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetRecorddifferentiationCount(tz *entities.RecorddifferentiationEntity) (entities.RecorddifferentiationEntities, error) {
	logger.Log.Println("In side GetRecorddifferentiationCount")
	value := entities.RecorddifferentiationEntities{}
	err := dbc.DB.QueryRow(getRecorddifferentiationcount, tz.Clientid, tz.Mstorgnhirarchyid, tz.Seqno).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetRecorddifferentiationCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetRecorddifferentiationname(page *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentiationnameEntity, error) {
	logger.Log.Println("In side GelAllRecorddifferentiation")
	logger.Log.Println(getdifferentiationname)
	values := []entities.RecorddifferentiationnameEntity{}
	rows, err := dbc.DB.Query(getdifferentiationname, page.Clientid, page.Mstorgnhirarchyid, "%"+page.Name+"%", page.Parentcatagorytypeid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecorddifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.RecorddifferentiationnameEntity{}
		rows.Scan(&value.Id, &value.Diffname, &value.Parentcatname)
		logger.Log.Println("Parent Category name --------->", &value.Parentcatname)
		logger.Log.Println("Parent Category length --------->", len(value.Parentcatname))
		if len(value.Parentcatname) > 0 {
			value.Recorddiffname = value.Diffname + " (" + value.Parentcatname + ")"
		} else {
			value.Recorddiffname = value.Diffname
		}
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetRecorddifferentiationoldname(ID int64) (string, string, error) {
	logger.Log.Println("In side GetRecorddifferentiationoldname")
	var Diffname string
	var OldParentcategorynames string
	var sql = "SELECT name,parentcategorynames FROM mstrecorddifferentiation WHERE id = ?"
	rows, err := dbc.DB.Query(sql, ID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetRecorddifferentiationoldname Get Statement Prepare Error", err)
		return Diffname, OldParentcategorynames, err
	}
	for rows.Next() {
		err = rows.Scan(&Diffname, &OldParentcategorynames)
		logger.Log.Println("GetRecorddifferentiationoldname rows.next() Error", err)
	}
	return Diffname, OldParentcategorynames, nil
}

func (dbc DbConn) Updateparentcategorynames(ClientID int64, OrgnID int64, ID int64, newparentcategoryname string, OldParentcategorynames string) error {
	logger.Log.Println("Updatestageid parameters -->", ClientID, OrgnID, ID, OldParentcategorynames, newparentcategoryname)
	var sql = "update mstrecorddifferentiation SET parentcategorynames = replace(parentcategorynames,?,?) where clientid=?  AND mstorgnhirarchyid=? AND deleteflg=0"
	stmt, err := dbc.DB.Prepare(sql)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}
	//'%->946->%'
	defer stmt.Close()
	_, err = stmt.Exec(OldParentcategorynames, newparentcategoryname, ClientID, OrgnID)
	// logger.Log.Println("%->"+strconv.FormatInt(ID, 10), "%->"+strconv.FormatInt(ID, 10)+"->%", strconv.FormatInt(ID, 10)+"->%")
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}
