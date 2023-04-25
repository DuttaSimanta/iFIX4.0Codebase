package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertWorkdifferentiation = "INSERT INTO mstworkdifferentiation (clientid, mstorgnhirarchyid, forrecorddifftypeid, forrecorddiffid, mainrecorddifftypeid) VALUES (?,?,?,?,?)"
var duplicateWorkdifferentiation = "SELECT count(id) total FROM  mstworkdifferentiation WHERE clientid = ? AND mstorgnhirarchyid = ? AND forrecorddifftypeid = ? AND forrecorddiffid = ? AND deleteflg = 0"

// var getWorkdifferentiation = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.forrecorddifftypeid as Forrecorddifftypeid, a.forrecorddiffid as Forrecorddiffid, a.mainrecorddifftypeid as Mainrecorddifftypeid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifftypname,e.name as Recorddiffname,f.typename as Recorddifftyplabel FROM mstworkdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.forrecorddifftypeid =d.id and a.forrecorddiffid=e.id and a.mainrecorddifftypeid =f.id ORDER BY a.id DESC LIMIT ?,?"
// var getWorkdifferentiationcount = "SELECT count(a.id) as total FROM mstworkdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.forrecorddifftypeid =d.id and a.forrecorddiffid=e.id and a.mainrecorddifftypeid =f.id"
var updateWorkdifferentiation = "UPDATE mstworkdifferentiation SET mstorgnhirarchyid = ?, forrecorddifftypeid = ?, forrecorddiffid = ?, mainrecorddifftypeid = ? WHERE id = ? "
var deleteWorkdifferentiation = "UPDATE mstworkdifferentiation SET deleteflg = '1' WHERE id = ? "
var workingcategory = "SELECT b.id,b.name as Recorddiffname FROM mstworkdifferentiation a,mstrecorddifferentiation b WHERE a.mainrecorddifftypeid=b.recorddifftypeid AND a.clientid=? and a.mstorgnhirarchyid=? and a.forrecorddifftypeid=? and a.forrecorddiffid=? and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0"
var getworkinglabel = "select distinct b.id,b.name,COALESCE(b.parentcategorynames,'') as parentcategorynames ,b.recorddifftypeid,(select typename from mstrecorddifferentiationtype where id=b.recorddifftypeid) as recorddifftypename,a.forrecorddifftypeid,a.forrecorddiffid from mstworkdifferentiation a,mstrecorddifferentiation b where a.clientid=? and a.mstorgnhirarchyid=? and a.forrecorddifftypeid=? and a.forrecorddiffid=? and a.mainrecorddifftypeid=b.recorddifftypeid and a.activeflg=1 and a.deleteflg=0 and b.activeflg=1 and b.deleteflg=0 order by b.name"

func (dbc DbConn) CheckDuplicateWorkdifferentiation(tz *entities.WorkdifferentiationEntity) (entities.WorkdifferentiationEntities, error) {
	logger.Log.Println("In side CheckDuplicateWorkdifferentiation")
	value := entities.WorkdifferentiationEntities{}
	err := dbc.DB.QueryRow(duplicateWorkdifferentiation, tz.Clientid, tz.Mstorgnhirarchyid, tz.Forrecorddifftypeid, tz.Forrecorddiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateWorkdifferentiation Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) InsertWorkdifferentiation(tz *entities.WorkdifferentiationEntity) (int64, error) {
	logger.Log.Println("In side InsertWorkdifferentiation")
	logger.Log.Println("Query -->", insertWorkdifferentiation)
	stmt, err := dbc.DB.Prepare(insertWorkdifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertWorkdifferentiation Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Forrecorddifftypeid, tz.Forrecorddiffid, tz.Mainrecorddifftypeid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Forrecorddifftypeid, tz.Forrecorddiffid, tz.Mainrecorddifftypeid)
	if err != nil {
		logger.Log.Println("InsertWorkdifferentiation Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) Getworkdifferentiationvalue(page *entities.WorkdifferentiationEntity) ([]entities.WorkdifferentiationsingleEntity, error) {
	logger.Log.Println("In side Getworkdifferentiationvalue")
	logger.Log.Println(workingcategory)
	values := []entities.WorkdifferentiationsingleEntity{}
	rows, err := dbc.DB.Query(workingcategory, page.Clientid, page.Mstorgnhirarchyid, page.Forrecorddifftypeid, page.Forrecorddiffid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllWorkdifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.WorkdifferentiationsingleEntity{}
		rows.Scan(&value.Id, &value.Recorddiffname)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetWokinglabelname(page *entities.WorkdifferentiationEntity) ([]entities.Workinglabelname, error) {
	logger.Log.Println("In side GelAllWorkdifferentiation")
	logger.Log.Println(getworkinglabel)
	values := []entities.Workinglabelname{}
	rows, err := dbc.DB.Query(getworkinglabel, page.Clientid, page.Mstorgnhirarchyid, page.Forrecorddifftypeid, page.Forrecorddiffid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllWorkdifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Workinglabelname{}
		rows.Scan(&value.ID, &value.Workingcatename, &value.Parentcatenames, &value.Recorddifftypid, &value.Recorddifftypname, &value.Forrecorddifftypeid, &value.Forrecorddiffid)
		if len(value.Parentcatenames) > 0 {
			value.Name = value.Workingcatename + " (" + value.Parentcatenames + ")"
		} else {
			value.Name = value.Workingcatename
		}
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllWorkdifferentiation(tz *entities.WorkdifferentiationEntity, OrgnType int64) ([]entities.WorkdifferentiationEntity, error) {
	logger.Log.Println("In side GelAllWorkdifferentiation")

	values := []entities.WorkdifferentiationEntity{}
	var getWorkdifferentiation string
	var params []interface{}
	if OrgnType == 1 {
		getWorkdifferentiation = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.forrecorddifftypeid as Forrecorddifftypeid, a.forrecorddiffid as Forrecorddiffid, a.mainrecorddifftypeid as Mainrecorddifftypeid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifftypname,e.name as Recorddiffname,f.typename as Recorddifftyplabel FROM mstworkdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f WHERE a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.forrecorddifftypeid =d.id and a.forrecorddiffid=e.id and a.mainrecorddifftypeid =f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getWorkdifferentiation = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.forrecorddifftypeid as Forrecorddifftypeid, a.forrecorddiffid as Forrecorddiffid, a.mainrecorddifftypeid as Mainrecorddifftypeid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifftypname,e.name as Recorddiffname,f.typename as Recorddifftyplabel FROM mstworkdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.forrecorddifftypeid =d.id and a.forrecorddiffid=e.id and a.mainrecorddifftypeid =f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getWorkdifferentiation = "SELECT a.id as Id,a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.forrecorddifftypeid as Forrecorddifftypeid, a.forrecorddiffid as Forrecorddiffid, a.mainrecorddifftypeid as Mainrecorddifftypeid, a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.typename as Recorddifftypname,e.name as Recorddiffname,f.typename as Recorddifftyplabel FROM mstworkdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.forrecorddifftypeid =d.id and a.forrecorddiffid=e.id and a.mainrecorddifftypeid =f.id ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	logger.Log.Println(getWorkdifferentiation)
	rows, err := dbc.DB.Query(getWorkdifferentiation, params...)

	// rows, err := dbc.DB.Query(getWorkdifferentiation, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllWorkdifferentiation Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.WorkdifferentiationEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Forrecorddifftypeid, &value.Forrecorddiffid, &value.Mainrecorddifftypeid, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.Recorddifftypname, &value.Recorddiffname, &value.Recorddifftyplabel)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateWorkdifferentiation(tz *entities.WorkdifferentiationEntity) error {
	logger.Log.Println("In side UpdateWorkdifferentiation")
	stmt, err := dbc.DB.Prepare(updateWorkdifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateWorkdifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Mstorgnhirarchyid, tz.Forrecorddifftypeid, tz.Forrecorddiffid, tz.Mainrecorddifftypeid, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateWorkdifferentiation Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteWorkdifferentiation(tz *entities.WorkdifferentiationEntity) error {
	logger.Log.Println("In side DeleteWorkdifferentiation")
	stmt, err := dbc.DB.Prepare(deleteWorkdifferentiation)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteWorkdifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteWorkdifferentiation Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetWorkdifferentiationCount(tz *entities.WorkdifferentiationEntity, OrgnTypeID int64) (entities.WorkdifferentiationEntities, error) {
	logger.Log.Println("In side GetWorkdifferentiationCount")
	value := entities.WorkdifferentiationEntities{}
	var getWorkdifferentiationcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getWorkdifferentiationcount = "SELECT count(a.id) as total FROM mstworkdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f WHERE  a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.forrecorddifftypeid =d.id and a.forrecorddiffid=e.id and a.mainrecorddifftypeid =f.id"
	} else if OrgnTypeID == 2 {
		getWorkdifferentiationcount = "SELECT count(a.id) as total FROM mstworkdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.forrecorddifftypeid =d.id and a.forrecorddiffid=e.id and a.mainrecorddifftypeid =f.id"
		params = append(params, tz.Clientid)
	} else {
		getWorkdifferentiationcount = "SELECT count(a.id) as total FROM mstworkdifferentiation a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstrecorddifferentiationtype f WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg =0 and a.activeflg=1 and a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.forrecorddifftypeid =d.id and a.forrecorddiffid=e.id and a.mainrecorddifftypeid =f.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getWorkdifferentiationcount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getWorkdifferentiationcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		log.Print("GetWorkdifferentiationCount Get Statement Prepare Error", err)
		return value, err
	}
}
