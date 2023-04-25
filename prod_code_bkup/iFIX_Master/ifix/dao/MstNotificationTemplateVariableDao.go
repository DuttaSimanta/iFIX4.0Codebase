package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstTemplateVariable = "INSERT INTO msttemplatevariable ( clientid, mstorgnhirarchyid, templatename, query, params, queryflag) VALUES (?,?,?,?,?,?)"
var duplicateMstTemplateVariable = "SELECT count(id) total FROM  msttemplatevariable WHERE clientid = ? AND mstorgnhirarchyid = ?  AND templatename=? AND query=? AND params=? AND queryflag=? AND activeflg =1 AND deleteflg = 0 "

// var getMstTemplateVariable= "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.templatename as TemplateName, a.query as Query, a.params as Params,a.queryflag as Queryflag,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname  FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"
// var getMstTemplateVariablecount = "SELECT count(a.id) as total FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
var updateMstTemplateVariable = "UPDATE msttemplatevariable SET clientid=?,mstorgnhirarchyid = ?, templatename = ?, query = ?, params = ?,queryflag=? WHERE id = ? "
var deleteMstTemplateVariable = "UPDATE msttemplatevariable SET deleteflg ='1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstTemplateVariable(tz *entities.MstTemplateVariableEntity) (entities.MstTemplateVariableEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstTemplateVariable ")
	value := entities.MstTemplateVariableEntities{}
	err := dbc.DB.QueryRow(duplicateMstTemplateVariable, tz.Clientid, tz.Mstorgnhirarchyid, tz.TemplateName, tz.Query, tz.Params, tz.Queryflag).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstTemplateVariable Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) AddMstTemplateVariable(tz *entities.MstTemplateVariableEntity) (int64, error) {
	logger.Log.Println("In side AddMstTemplateVariable")
	logger.Log.Println("Query -->", insertMstTemplateVariable)
	stmt, err := dbc.DB.Prepare(insertMstTemplateVariable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstTemplateVariable Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.TemplateName, tz.Query, tz.Params, tz.Queryflag)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.TemplateName, tz.Query, tz.Params, tz.Queryflag)
	if err != nil {
		logger.Log.Println("AddMstTemplateVariable Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc TxConn) AddMstTemplateVariablecopy(tz *entities.MstTemplateVariableEntity) (int64, error) {
	logger.Log.Println("In side AddMstTemplateVariable")
	logger.Log.Println("Query -->", insertMstTemplateVariable)
	stmt, err := dbc.TX.Prepare(insertMstTemplateVariable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstTemplateVariable Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.TemplateName, tz.Query, tz.Params, tz.Queryflag)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.TemplateName, tz.Query, tz.Params, tz.Queryflag)
	if err != nil {
		logger.Log.Println("AddMstTemplateVariable Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc DbConn) GetAllMstTemplateVariable(page *entities.MstTemplateVariableEntity, OrgnType int64) ([]entities.MstTemplateVariableEntity, error) {
	logger.Log.Println("In side GetAllMstTemplateVariable")
	values := []entities.MstTemplateVariableEntity{}

	var getMstTemplateVariable string
	var params []interface{}
	if OrgnType == 1 {
		getMstTemplateVariable = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.templatename as TemplateName, coalesce(a.query, '') as Query, coalesce(a.params, '') as Params,a.queryflag as Queryflag,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname  FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstTemplateVariable = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.templatename as TemplateName, coalesce(a.query, '') as Query, coalesce(a.params,'') as Params,a.queryflag as Queryflag,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname  FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstTemplateVariable = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.templatename as TemplateName, coalesce(a.query, '') as Query, coalesce(a.params, '') as Params,a.queryflag as Queryflag,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname  FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}
	rows, err := dbc.DB.Query(getMstTemplateVariable, params...)
	logger.Log.Println("Query====>", getMstTemplateVariable)

	// rows, err := dbc.DB.Query(getMstTemplateVariable, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstTemplateVariable Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstTemplateVariableEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.TemplateName, &value.Query, &value.Params, &value.Queryflag, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		logger.Log.Println(value.Clientname, value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllMstTemplateVariableList(tz *entities.MstTemplateVariableEntity) ([]entities.MstTemplateVariableEntityList, error) {
	logger.Log.Println("In side GetAllMstTemplateVariableList")
	values := []entities.MstTemplateVariableEntityList{}

	var getMstTemplateVariableList string

	getMstTemplateVariableList = "SELECT a.id as Id, a.templatename as TemplateName  FROM msttemplatevariable a WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =? AND a.mstorgnhirarchyid = ?"
	rows, err := dbc.DB.Query(getMstTemplateVariableList, tz.Clientid, tz.Mstorgnhirarchyid)

	// rows, err := dbc.DB.Query(getMstTemplateVariable, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstTemplateVariableList Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstTemplateVariableEntityList{}
		rows.Scan(&value.Id, &value.TemplateName)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstTemplateVariable(tz *entities.MstTemplateVariableEntity) error {
	logger.Log.Println("In side UpdateMstTemplateVariable")
	stmt, err := dbc.DB.Prepare(updateMstTemplateVariable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstTemplateVariable Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.TemplateName, tz.Query, tz.Params, tz.Queryflag, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstTemplateVariable Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstTemplateVariable(tz *entities.MstTemplateVariableEntity) error {
	logger.Log.Println("In side DeleteMstTemplateVariable", tz)
	stmt, err := dbc.DB.Prepare(deleteMstTemplateVariable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstTemplateVariable Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstTemplateVariable Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstTemplateVariableCount(tz *entities.MstTemplateVariableEntity, OrgnTypeID int64) (entities.MstTemplateVariableEntities, error) {
	logger.Log.Println("In side GetMstTemplateVariableCount")
	value := entities.MstTemplateVariableEntities{}
	var getMstTemplateVariablecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstTemplateVariablecount = "SELECT count(a.id) as total FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c WHERE  a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
	} else if OrgnTypeID == 2 {
		getMstTemplateVariablecount = "SELECT count(a.id) as total FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
		params = append(params, tz.Clientid)
	} else {
		getMstTemplateVariablecount = "SELECT count(a.id) as total FROM msttemplatevariable a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstTemplateVariablecount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMstTemplateVariablecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstTemplateVariableCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) GetTemplatevariableCopy(tz *entities.MstTemplateVariableEntity, i int) ([]entities.MstTemplateVariableEntity, error) {
	logger.Log.Println("In side GetAllMstTemplateVariableList")
	values := []entities.MstTemplateVariableEntity{}

	var getMstTemplateVariableList string

	getMstTemplateVariableList = "SELECT a.templatename, a.query as query, a.params as params,a.queryflag  FROM msttemplatevariable a WHERE a.clientid =? AND a.mstorgnhirarchyid = ? AND a.templatename=? AND  a.deleteflg =0 and a.activeflg=1 "
	rows, err := dbc.DB.Query(getMstTemplateVariableList, tz.Clientid, tz.Mstorgnhirarchyid, tz.TemplateNames[i])

	// rows, err := dbc.DB.Query(getMstTemplateVariable, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstTemplateVariableList Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstTemplateVariableEntity{}
		rows.Scan(&value.TemplateName, &value.Query, &value.Params, &value.Queryflag)
		values = append(values, value)
	}
	return values, nil
}
