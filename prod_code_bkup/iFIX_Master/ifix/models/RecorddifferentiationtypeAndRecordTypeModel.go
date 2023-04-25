package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertRecorddifftypeAndRecordtype(tz *entities.RecorddifftypeAndRecordTypeEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Bannermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//j := 0
	var id int64
	id = 0
	var k int64
	k = 0
	/* Starting Transaction*/
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	//dataAccess := dao.DbConn{DB: db}
	dataAccess := dao.TxConn{TX: tx}
	count, err := dataAccess.CheckDuplicateRecorddifftype(tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err = dataAccess.InsertRecorddifftype(tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			return 0, false, err, "Something Went Wrong"
		}
		k++
	}
	tz.Torecorddifftypeid = id
	tz.Torecorddiffid = 0
	count1, err := dataAccess.CheckDuplicateRecrdType(tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	if count1.Total == 0 {
		id, err = dataAccess.InsertRecrdType(tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			return 0, false, err, "Something Went Wrong"
		}
		k++
	}
	if k == 2 {
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Banner  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Rollback()
		if err != nil {
			// log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Banner  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}
}

func UpdateRecorddifftypeAndRecordtype(tz *entities.RecorddifftypeAndRecordTypeEntity) (bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupnewmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	var k int64
	k = 0
	//dataAccess := dao.DbConn{DB: db}
	dataAccess := dao.TxConn{TX: tx}
	count, err := dataAccess.CheckDuplicateRecorddifftype(tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err = dataAccess.UpdateRecorddifftype(tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			return false, err, "Something Went Wrong"
		}
		k++
	}
	//tz.Torecorddifftypeid=id
	tz.Torecorddifftypeid = tz.Id
	count1, err := dataAccess.CheckDuplicateRecrdType(tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	if count1.Total == 0 {
		err = dataAccess.UpdateRecrdType(tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			return false, err, "Something Went Wrong"
		}
		k++
	}
	if k == 2 {
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Banner  Statement Commit error", err)
			return false, err, ""
		}
		return true, err, ""
	} else {
		err = tx.Rollback()
		if err != nil {
			// log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Banner  Statement Commit error", err)
			return false, err, ""
		}
		return false, nil, "Data Already Exist."
	}
}

// func DeleteRecorddifftypeAndRecordtype(tz *entities.RecorddifftypeAndRecordTypeEntity) (bool, error, string) {

// 	logger.Log.Println("In side Clientsupportgroupnewmodel")
// 	lock.Lock()
// 	defer lock.Unlock()
// 	db, err := config.ConnectMySqlDbSingleton()
// 	if err != nil {
// 		logger.Log.Println("database connection failure", err)
// 		return false, err, "Something Went Wrong"
// 	}
// 	tx, err := db.Begin()
// 	if err != nil {
// 		logger.Log.Println("database transaction connection failure", err)
// 		tx.Rollback()
// 		return false, err, "Something Went Wrong"
// 	}
// 	//dataAccess := dao.DbConn{DB: db}
// 	dataAccess := dao.TxConn{TX: tx}
// 	err1 := dataAccess.DeleteRecorddifftype(tz)
// 	if err1 != nil {
// 		logger.Log.Println(err1)
// 		tx.Rollback()
// 		//db.Close()
// 		return false, err, "Data deletion failure."
// 	}

// 	err2 := dataAccess.DeleteRecrdType(tz)
// 	if err2 != nil {
// 		logger.Log.Println(err2)
// 		tx.Rollback()
// 		//db.Close()
// 		return false, err, "Data deletion failure."
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		// log.Print("MoveWorkflow  Statement Commit error", err)
// 		logger.Log.Print("Banner  Statement Commit error", err)
// 		return false, err, ""
// 	}
// 	return true, err, ""
// }
func DeleteRecorddifftypeAndRecordtype(tz *entities.RecorddifftypeAndRecordTypeEntity) (bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupnewmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err2 := dataAccess.DeleteRecrdType(tz)
	if err2 != nil {
		logger.Log.Println(err2)
		// tx.Rollback()
		//db.Close()
		return false, err, "Data deletion failure."
	}
	return true, err, ""
}
