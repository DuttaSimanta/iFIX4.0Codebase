package models

import (
	"fmt"
	"src/config"
	"src/dao"
	"src/logger"

	"github.com/gofrs/uuid"
)

func UpdateIfixsysid() (bool, error, string) {
	logger.Log.Println("In side UpdateIfixsysid.....Updating Started")
	// lock.Lock()
	// defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Database connection failure"
	}
	// tx, err := db.Begin()
	// if err != nil {
	// 	logger.Log.Println("database transaction connection failure", err)
	// 	return false, err, "Database transaction connection failure"
	// }
	dataAccess := dao.DbConn{DB: db}
	// dataAccess1 := dao.TxConn{TX: tx}
	tables, err1 := dataAccess.GetTableToUpdateUuid()
	if err1 != nil {
		// tx.Rollback()
		// err = tx.Rollback()
		// if err != nil {
		// 	logger.Log.Print("UpdateIfixsysid  Statement Commit error", err)
		// 	return false, err, ""
		// }
		return false, err1, "Error while fetching tablelist"
	}
	for i := 0; i < len(tables); i++ {
		tx, err := db.Begin()
		if err != nil {
			logger.Log.Println("database transaction connection failure", err)
			return false, err, "Database transaction connection failure"
		}
		dataAccess1 := dao.TxConn{TX: tx}
		tablerowsid, err1 := dataAccess.GetTableRowsToUpdateUuid(tables[i])
		if err1 != nil {
			err = tx.Rollback()
			if err != nil {
				logger.Log.Print("UpdateIfixsysid  Statement Commit error", err)
				return false, err, ""
			}
			return false, err1, "Error while fetching records where ifix sysid is null"
		}
		for j := 0; j < len(tablerowsid); j++ {
			uid, err := uuid.NewV4()
			if err != nil {
				tx.Rollback()
				fmt.Println("Error is >>>>>.: ", err)
			}
			err = dataAccess1.UpdateIfixsysid(tablerowsid[j], tables[i], uid)
			if err != nil {
				err1 = tx.Rollback()
				if err != nil {
					logger.Log.Print("UpdateIfixsysid  Statement Commit error", err)
					return false, err1, ""
				}
				return false, err, "Error while updating the ifix sysid"
			}
		}
		err1 = tx.Commit()
		if err != nil {
			logger.Log.Print("UpdateIfixsysid  Statement Commit error", err)
			return false, err1, ""
		}
	}
	// tx.Commit()
	logger.Log.Println("UpdateIfixsysid End")
	return true, err, ""

}
