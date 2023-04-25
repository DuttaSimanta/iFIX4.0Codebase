package models

import (
	"src/dao"
	"src/entities"
	"src/logger"
)

func GetDownloadList(page *entities.ReportDownloadListEntity) (entities.ReportDownloadListEntities, bool, error, string) {
	logger.Log.Println("In side GetDownloadList")
	t := entities.ReportDownloadListEntities{}
	// lock.Lock()
	// defer lock.Unlock()
	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return t, false, err, "Something Went Wrong"
		}
		db = dbcon
	}
	dataAccess := dao.DbConn{DB: db}
	values, err := dataAccess.GetDownloadList(page)
	if err != nil {
		return t, false, err, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err := dataAccess.GetDownloadListCount(page)

		if err != nil {
			return t, false, err, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
