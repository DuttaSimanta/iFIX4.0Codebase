package models

import (
	"database/sql"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

func GetUserActionnamewithapi(tz *entities.UserroleactionnameEntity) ([]int64, bool, error, string) {
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	ids, success, err, msg := GetUserActionname(db, tz)
	return ids, success, err, msg
}
func GetUserActionname(db *sql.DB, tz *entities.UserroleactionnameEntity) ([]int64, bool, error, string) {
	logger.Log.Println("In side model")

	var ids []int64
	//defer db.Close()

	dataAccess := dao.DbConn{DB: db}
	ids, err1 := dataAccess.GetUserRolewiseActionname(tz)
	if err1 != nil {
		return ids, false, err1, "Something Went Wrong"
	}
	if len(ids) > 0 {
		return ids, true, nil, ""
	} else {
		ids, err1 := dataAccess.GetRolewiseActionname(tz)
		if err1 != nil {
			return ids, false, err1, "Something Went Wrong"
		}
		return ids, true, nil, ""
	}
}
