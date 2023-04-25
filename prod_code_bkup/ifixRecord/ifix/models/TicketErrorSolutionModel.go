package models

import (
	"ifixRecord/ifix/dao"
	"ifixRecord/ifix/entities"
	"ifixRecord/ifix/logger"
)

func UpdateRecordfulTicketIdDeleteFlag(tz *entities.RecordfulTicketIdDeleteFlagUpdateEntity) (bool, error, string) {
	logger.Log.Println("In side UpdateRecordfulTicketIdDeleteFlag Model fucntion")
	lock.Lock()
	defer lock.Unlock()
	db, err := ConnectMySqlDb()
	if err != nil {
		logger.Log.Println("Error in DBConnection", err)
		return false, err, "Something Went Wrong"
	}
	logger.Log.Println("DB is connected...")
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateRecordfulTicketIdDeleteFlag(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	logger.Log.Println("Data accessed...")
	return true, err, ""
}

func UpdateTrnRecordCodeDeleteFlgById(tz *entities.TrnRecordCodeDeleteFlgUpdateByIdEntity) (bool, error, string) {
	logger.Log.Println("Inside UpdateRecordCodeDeleteFlgById Model function")
	lock.Lock()
	defer lock.Unlock()
	db, err := ConnectMySqlDb()
	if err != nil {
		logger.Log.Println("Error in DBConnection", err)
		return false, err, "Something Went Wrong"
	}
	logger.Log.Println("DB is connected...")
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateTrnRecordCodeDeleteFlgById(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	logger.Log.Println("Data accessed...")
	return true, err, ""

}
