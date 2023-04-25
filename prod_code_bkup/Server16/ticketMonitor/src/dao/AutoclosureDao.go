package dao

import (
	"database/sql"
	"errors"
	Logger "src/logger"
)

func AutoUserRecordCodeDelete(db *sql.DB) error {
	var Sql = "DELETE FROM mstuserrecord where (opendate+1800)<=UNIX_TIMESTAMP();"
	Logger.Log.Println("Query -->", Sql)
	stmt, err := db.Prepare(Sql)

	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}

	return nil
}
