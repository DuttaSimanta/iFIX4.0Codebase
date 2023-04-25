package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertClientsupportgroup(tz *entities.ClientsupportgroupEntity) (int64, bool, error, string) {

	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	if tz.Isworkflow == "Y" {
		tx, err := dbcon.Begin()
		if err != nil {
			// dbcon.Close()
			logger.Log.Println("Transaction creation error.", err)
			return 0, false, err, "Something Went Wrong"
		}
		count, err := dao.CheckDuplicateClientsupportgroupwithtransaction(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}

		if count.Total == 0 {
			lastinsertedID, err := dao.InsertClientsupportgroupwithtransaction(tx, tz)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				// dbcon.Close()
				return 0, false, err, "Data insertion failure."
			}
			count1, err := dao.CheckDuplicateMstgroupwithtransaction(tx, tz)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				// dbcon.Close()
				return 0, false, err, "Data insertion failure."
			}
			if count1.Total == 0 {
				_, err := dao.InsertMstgroupwithtransaction(tx, tz, lastinsertedID)
				if err != nil {
					logger.Log.Println(err)
					tx.Rollback()
					// dbcon.Close()
					return 0, false, err, "Data insertion failure."
				}
				tx.Commit()
				return lastinsertedID, true, err, ""
			} else {
				tx.Rollback()
				// dbcon.Close()
				return 0, false, err, "Already data exist.Please verify the data."
			}
		} else {
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Already data exist.Please verify the data."
		}
	} else {
		dataAccess := dao.DbConn{DB: dbcon}
		count, err := dataAccess.CheckDuplicateClientsupportgroup(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err := dataAccess.InsertClientsupportgroup(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			return id, true, err, ""
		} else {
			return 0, false, nil, "Data Already Exist."
		}
	} //else part end here....

}

func GetAllClientsupportgroup(page *entities.ClientsupportgroupEntity) (entities.ClientsupportgroupEntities, bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupmodel")
	t := entities.ClientsupportgroupEntities{}
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllClientsupportgroup(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetClientsupportgroupCount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteClientsupportgroup(tz *entities.ClientsupportgroupEntity) (bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupmodel")
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	if tz.Isworkflow == "Y" {
		tx, err := dbcon.Begin()
		if err != nil {
			// dbcon.Close()
			logger.Log.Println("Transaction creation error.", err)
			return false, err, "Something Went Wrong"
		}
		err1 := dao.DeleteMstgroupwithtransaction(tx, tz)
		if err1 != nil {
			logger.Log.Println(err1)
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Data deletion failure."
		}

		err2 := dao.DeleteClientsupportgroupwithtransaction(tx, tz)
		if err2 != nil {
			logger.Log.Println(err2)
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Data deletion failure."
		}
		tx.Commit()
		return true, nil, ""
	} else {
		dataAccess := dao.DbConn{DB: dbcon}
		err1 := dataAccess.DeleteClientsupportgroup(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		return true, nil, ""
	}

}

func UpdateClientsupportgroup(tz *entities.ClientsupportgroupEntity) (bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupmodel")
	// dbcon, err := config.ConnectMySqlDb()
	// defer dbcon.Close()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}

	if tz.Isworkflow == "Y" {
		tx, err := dbcon.Begin()
		if err != nil {
			// dbcon.Close()
			logger.Log.Println("Transaction creation error.", err)
			return false, err, "Something Went Wrong"
		}
		count, err := dao.CheckDuplicateMstgroupwithtransaction(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Data updation failure."
		}
		if count.Total == 0 {
			err := dao.UpdateMstgroupwithtransaction(tx, tz)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				// dbcon.Close()
				return false, err, "Data updation failure."
			}
		}
		count1, err1 := dao.CheckDuplicateClientsupportgroupwithtransaction(tx, tz)
		if err1 != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Data updation failure."
		}
		if count1.Total == 0 {
			err := dao.UpdateClientsupportgroupwithtransaction(tx, tz)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				// dbcon.Close()
				return false, err, "Data updation failure."
			}
			tx.Commit()
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}

	} else {
		dataAccess := dao.DbConn{DB: dbcon}
		count, err := dataAccess.CheckDuplicateClientsupportgroupupdate(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			err := dataAccess.UpdateClientsupportgroup(tz)
			if err != nil {
				return false, err, "Something Went Wrong"
			}
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}
	}
}
func Getgroupbyorgid(page *entities.ClientsupportgroupEntity) ([]entities.ClientsupportgroupsingleEntity, bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupmodel")
	t := []entities.ClientsupportgroupsingleEntity{}
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getgroupbyorgid(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getprocessgroupbyorgid(page *entities.ClientsupportgroupEntity) ([]entities.ClientsupportgroupsingleEntity, bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupmodel")
	t := []entities.ClientsupportgroupsingleEntity{}
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getprocessgroupbyorgid(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getprocessgroupbyorgids(page *entities.ClientsupportgroupEntity) ([]entities.ClientsupportgroupsingleEntity, bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupmodel")
	t := []entities.ClientsupportgroupsingleEntity{}
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getprocessgroupbyorgids(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
