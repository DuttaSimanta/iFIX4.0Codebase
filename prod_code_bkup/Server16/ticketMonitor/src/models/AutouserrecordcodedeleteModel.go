package models

import (
	"src/dao"
	"src/logger"
)

func AutoUserRecordCodeDelete() {
	// var t map[string]interface{}
	// var response map[string]interface{}
	// if mutexutility.MutexLocked(lock) == false {
	// 	lock.Lock()
	// 	defer lock.Unlock()
	// }
	// db, err := dbconfig.ConnectMySqlDb()
	// if err != nil {
	// 	logger.Log.Println("database connection failure", err)
	// 	return t, false, err, "Something Went Wrong"
	// }
	//defer db.Close()
	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
		}
		db = dbcon
	}
	// now := time.Now().Unix()
	err := dao.AutoUserRecordCodeDelete(db)
	if err != nil {
		logger.Log.Println("database connection failure", err)
	}
}
