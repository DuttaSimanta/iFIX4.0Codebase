package dao

import (
	"src/logger"

	"github.com/gofrs/uuid"
)

var gettable = "select msttablename from msttransporttable where deleteflg=0"

func (dbc DbConn) GetTableToUpdateUuid() ([]string, error) {
	logger.Log.Println("In side Gettype")
	values := []string{}

	rows, err := dbc.DB.Query(gettable)
	logger.Log.Println(gettable)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Gettype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		var value string
		rows.Scan(&value)
		values = append(values, value)
	}
	return values, nil
}
func (dbc DbConn) GetTableRowsToUpdateUuid(table string) ([]int64, error) {
	logger.Log.Println("In side Gettype")
	values := []int64{}
	gettablerows := "select id from " + table + " where ifixsysid is NULL or ifixsysid=''"
	logger.Log.Println(gettablerows)
	rows, err := dbc.DB.Query(gettablerows)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Gettype Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		var value int64
		rows.Scan(&value)
		values = append(values, value)
	}
	return values, nil
}
func (dbc TxConn) UpdateIfixsysid(id int64, table string, uid uuid.UUID) error {
	logger.Log.Println("In side UpdateTransporttable")
	updateifixsysid := "UPDATE " + table + " SET ifixsysid = ? WHERE id = ? "
	stmt, err := dbc.TX.Prepare(updateifixsysid)
	logger.Log.Println(updateifixsysid)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateTransporttable Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(uid, id)
	if err != nil {
		logger.Log.Println("UpdateTransporttable Execute Statement  Error", err)
		return err
	}
	return nil
}
