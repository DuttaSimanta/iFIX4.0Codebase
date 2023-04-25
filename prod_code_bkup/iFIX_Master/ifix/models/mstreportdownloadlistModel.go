package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func ReportDownloadList(tz *entities.ReportDownloadEntity) (int64, bool, error, string) {
	logger.Log.Println("In side ReportDownloadList")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}

	dataAccess := dao.DbConn{DB: db}

	// count, err := dataAccess.CheckDuplicateMstAttribute(tz)
	// if err != nil {

	// 	return 0, false, err, "Something Went Wrong"
	// }
	// if count.Total == 0 {

	id, err := dataAccess.ReportDownloadList(tz)
	if err != nil {

		return 0, false, err, "Something Went Wrong"
	}
	return id, true, err, ""
	// }
	// return 0, false, err, "Data Already Exist"

}
