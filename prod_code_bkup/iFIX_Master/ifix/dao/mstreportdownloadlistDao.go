package dao

import (
	// "database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstDownloadList = "INSERT INTO mstreportgeneratedlog ( userid, originalfilename, uploadedfilename) VALUES (?,?,?) "

func (dbc DbConn) ReportDownloadList(tz *entities.ReportDownloadEntity) (int64, error) {
	logger.Log.Println("In side ReportDownloadList")
	logger.Log.Println("Query -->", insertMstDownloadList)
	stmt, err := dbc.DB.Prepare(insertMstDownloadList)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("ReportDownloadList Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Refuserid, tz.Originalfilename, tz.Uploadedfilename)
	res, err := stmt.Exec(tz.Refuserid, tz.Originalfilename, tz.Uploadedfilename)
	if err != nil {
		logger.Log.Println("ReportDownloadList Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
