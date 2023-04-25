package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertRecordConfigIncrement(tz *entities.RecordConfigIncrementEntity) (int64, bool, error, string) {
	logger.Log.Println("In side RecordConfigIncrement model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//	tx, err := db.Begin()
	dataAccess := dao.DbConn{DB: db}
	//dataAccess1 := dao.TxConn{TX: tx}
	if tz.IsClient == 0 {
		tx, err := db.Begin()
		dataAccess1 := dao.TxConn{TX: tx}
		count, err := dataAccess.CheckDuplicateRecordConfigIncrement(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err := dataAccess1.AddRecordConfig(tz)
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}
			_, err = dataAccess1.AddRecordIncrement(tz)
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}
			err = tx.Commit()
			if err != nil {
				//log.Print("MoveWorkflow  Statement Commit error", err)
				logger.Log.Print("RecordConfigIncrement  Statement Commit error", err)
				return 0, false, err, ""
			}
			return id, true, err, ""
		} else {
			return 0, false, nil, "Data Already Exist."
		}
	}
	if tz.IsClient == 1 {
		count, err := dataAccess.CheckDuplicateRecordConfig(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err := dataAccess.AddRecordConfigWithoutTx(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			return id, true, err, ""
		} else {
			return 0, false, nil, "Data Already Exist."
		}

	}
	return 0, false, err, "Something Went Wrong"
}

func GetAllRecordConfigIncrement(page *entities.RecordConfigIncrementEntity) (entities.RecordConfigIncrementEntities, bool, error, string) {
	logger.Log.Println("In side RecordConfigIncrement model")
	t := entities.RecordConfigIncrementEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(page.Clientid, page.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllRecordConfigIncrement(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetRecordConfigIncrementCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteRecordConfigIncrement(tz *entities.RecordConfigIncrementEntity) (bool, error, string) {
	logger.Log.Println("In side RecordConfigIncrement model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}

	dataAccess := dao.DbConn{DB: db}

	if tz.IsClient == 0 {
		tx, err := db.Begin()
		dataAccess1 := dao.TxConn{TX: tx}
		err1 := dataAccess1.DeleteRecordConfig(tz)
		if err1 != nil {
			tx.Rollback()
			return false, err1, "Something Went Wrong"
		}

		err1 = dataAccess1.DeleteRecordIncrement(tz)
		if err1 != nil {
			tx.Rollback()
			return false, err1, "Something Went Wrong"
		}
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("RecordConfigIncrement  Statement Commit error", err)
			return false, err, ""
		}
		return true, nil, ""
	}
	if tz.IsClient == 1 {
		err1 := dataAccess.DeleteRecordConfigWithoutTx(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		return true, nil, ""
	}
	return false, nil, "Something Went Wrong"
}

func UpdateRecordConfigIncrement(tz *entities.RecordConfigIncrementEntity) (bool, error, string) {
	logger.Log.Println("In side RecordConfigIncrement model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	dataAccess := dao.DbConn{DB: db}
	dataAccess1 := dao.TxConn{TX: tx}
	if tz.IsClient == 0 {
		count, err := dataAccess.CheckDuplicateRecordConfigIncrement(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 1 {
			err := dataAccess1.UpdateRecordConfig(tz)
			if err != nil {
				tx.Rollback()
				return false, err, "Something Went Wrong"
			}
			err = dataAccess1.UpdateRecordIncrement(tz)
			if err != nil {
				tx.Rollback()
				return false, err, "Something Went Wrong"
			}
			err = tx.Commit()
			if err != nil {
				//log.Print("MoveWorkflow  Statement Commit error", err)
				logger.Log.Print("RecordConfigIncrement  Statement Commit error", err)
				return false, err, ""
			}
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}
	}
	if tz.IsClient == 1 {
		count, err := dataAccess.CheckDuplicateRecordConfig(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			err := dataAccess.UpdateRecordConfigWithoutTx(tz)
			if err != nil {
				return false, err, "Something Went Wrong"
			}
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}

	}
	return false, err, "Something Went Wrong"

}
