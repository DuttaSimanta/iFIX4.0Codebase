package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"log"
)

/**
Get zone details
*/
func Searchzone(tz *entities.ZoneEntity) ([]entities.ZoneEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.ZoneEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Searchzone(tz)
	if err1 != nil {
		log.Println("database connection failure", err)
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, nil, ""
}
