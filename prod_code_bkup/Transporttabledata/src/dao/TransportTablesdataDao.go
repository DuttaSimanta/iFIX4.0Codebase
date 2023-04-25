package dao

import (
	"database/sql"
	"errors"
	"src/entities"
	"src/logger"
	"strconv"
	"strings"
)

func (dbc DbConn) GetClientOrg(page *entities.Transporttable) (string, string, error) {
	logger.Log.Println("In side GetClientOrg")
	// table := "'" + page.Tablenames[i] + "'"
	getschema := "select a.name,b.name from mstclient a,mstorgnhierarchy b where a.id=? and b.id=?"
	// logger.Log.Println(getschema)
	var clientname string
	var orgname string

	rows, err := dbc.DB.Query(getschema, page.Clientid, page.Mstorgnhirarchyid)

	if err != nil {
		logger.Log.Println("GetClientOrg Get Statement Prepare Error", err)
		return "", "", err
	}
	defer rows.Close()
	for rows.Next() {

		rows.Scan(&clientname, &orgname)
	}
	return clientname, orgname, nil
}
func (dbc DbConn) GetSchema(page *entities.Transporttable, i int) ([]string, error) {
	logger.Log.Println("In side GetSchema")
	table := "'" + page.Table[i].Tablename + "'"
	getschema := "select COLUMN_NAME from INFORMATION_SCHEMA.COLUMNS where TABLE_Name=" + table + " order by ORDINAL_POSITION"
	logger.Log.Println(getschema)
	var values []string
	rows, err := dbc.DB.Query(getschema)

	if err != nil {
		logger.Log.Println("GetSchema Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		var value string
		rows.Scan(&value)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetTableData(column []string, page *entities.Transporttable, i int) ([]map[string]interface{}, error) {
	logger.Log.Println("In side GetTableData")
	str := strings.Join(column, ",")
	// logger.Log.Println()
	table := page.Table[i].Tablename
	var getschema string
	var params []interface{}
	values := []map[string]interface{}{}
	if page.Table[i].Tabletype == 1 {
		getschema = "select " + str + " from " + table + " where clientid=? and mstorgnhirarchyid=?"
		// params = append(params, 0)
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
	} else if page.Table[i].Tabletype == 0 {
		getschema = "select " + str + " from " + table
		// params = append(params, 0)
	} else if page.Table[i].Tabletype == 2 {
		getschema = "select " + str + " from " + table + " where clientid=? and mstorgnhirarchyid=?"
		// params = append(params, 0)
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
	} else {
		logger.Log.Println("This tabletype" + strconv.FormatInt(page.Table[i].Tabletype, 10) + " is not configured in backend code")

		return values, errors.New("Not Having Table Type")
	}
	logger.Log.Println("Query>>>>>>>", getschema)
	rows, err := dbc.DB.Query(getschema, params...)
	if err != nil {
		logger.Log.Println("GetTableData Get Statement Prepare Error", err)
		return values, err
	}
	//logger.Log.Println(rows)
	// defer rows.Close()
	// for rows.Next() {
	// 	var value []interface{}
	// 	rows.Scan(value...)
	// 	logger.Log.Println(value)
	// 	values = append(values, value)
	// }
	// return values, nil
	for rows.Next() {
		fb := entities.NewDbFieldBind()
		err = fb.PutFields(rows)
		if err != nil {
			return values, err
		}
		rows.Scan(fb.GetFieldPtrArr()...)
		//logger.Log.Println("values", fb.GetFieldArr())
		values = append(values, fb.GetFieldArr())
	}
	// logger.Log.Println(len(values))
	return values, nil
}
func (dbc DbConn) CheckDuplicateData(id interface{}, table string) (int64, error) {
	logger.Log.Println("In side CheckDuplicateData")
	var value int64
	value = 0
	getcount := "select count(id)  from " + table + " where  id=?"
	logger.Log.Println("query->>>>", getcount)
	err := dbc.DB.QueryRow(getcount, id).Scan(&value)
	switch err {
	case sql.ErrNoRows:
		value = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateData Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc TxConn) InsertTabledata(column []string, params []interface{}, insertno []string, table string) error {
	logger.Log.Println("In side InsertTabledata")
	columnn := strings.Join(column, ",")
	insertnoo := strings.Join(insertno, ",")
	var insertUidGen = "INSERT INTO " + table + "(" + columnn + ") VALUES (" + insertnoo + ")"
	logger.Log.Println("Query -->", insertUidGen)
	stmt, err := dbc.TX.Prepare(insertUidGen)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertTabledata Prepare Statement  Error", err)
		return err
	}
	logger.Log.Println("Parameter -->", params)
	res, err := stmt.Exec(params...)
	if err != nil {
		logger.Log.Println("InsertTabledata Execute Statement  Error", err)
		return err
	}
	_, err = res.LastInsertId()
	return nil
}
func (dbc TxConn) UpdateTabledata(column []string, params []interface{}, insertno []string, table string) error {
	logger.Log.Println("In side UpdateTabledata")
	columnnn := strings.Join(column, ",")

	var updatetable = "UPDATE " + table + " SET " + columnnn + " WHERE id = ? "
	logger.Log.Println("query for update->", updatetable)
	logger.Log.Println("parameter->", params)
	stmt, err := dbc.TX.Prepare(updatetable)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateTabledata Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(params...)
	if err != nil {
		logger.Log.Println("UpdateTabledata Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc TxConn) DeleteTabledata(tablename string, values map[string]interface{}) error {
	logger.Log.Println("In side DeleteMapCategoryWithKeyword")
	query := "DELETE FROM " + tablename + " WHERE clientid=? and mstorgnhirarchyid=?"
	stmt, err := dbc.TX.Prepare(query)
	logger.Log.Println(query)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMapCategoryWithKeyword Prepare Statement  Error", err)
		return err
	}
	deleteresult, err := stmt.Exec(values["clientid"], values["mstorgnhirarchyid"])
	logger.Log.Println(deleteresult)
	if err != nil {
		logger.Log.Println("DeleteMapCategoryWithKeyword Execute Statement  Error", err)
		return err
	}
	return nil
}
