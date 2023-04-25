package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func GetAllPortgrouplevel() (entities.SupportgrouplevelEntities, bool, error, string) {
	logger.Log.Println("In side Portgrouplevelmodel")
	t := entities.SupportgrouplevelEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllPortgrouplevel()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t.Values = values
	return t, true, err, ""
}
