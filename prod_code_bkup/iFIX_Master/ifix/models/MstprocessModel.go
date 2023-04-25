package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

/*func InsertMstprocess(tz *entities.MstprocessEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	dbcon, err := config.ConnectMySqlDb()
	defer dbcon.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: dbcon}
	count, err := dataAccess.CheckDuplicateMstprocess(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertMstprocess(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}*/

func InsertMstprocesswithtransaction(tz *entities.MstprocessEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	dataAccess := dao.DbConn{DB: dbcon}
	tx, err := dbcon.Begin()
	if err != nil {
		//dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}
	count, err := dataAccess.CheckDuplicateMstprocess(tz)
	if err != nil {
		logger.Log.Println(err)
		return 0, false, err, "Data insertion failure."
	}
	if count.Total == 0 {
		lastinsertedID, err := dao.InsertMstprocesswithtransaction(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			return 0, false, err, "Data insertion failure."
		}
		count, err := dataAccess.CheckDuplicateMstprocessrecordmap(tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			return 0, false, err, "Data insertion failure."
		}
		if count.Total == 0 {
			_, err := dao.InsertMstprocessrecordmapwithtransaction(tx, tz, lastinsertedID)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				return 0, false, err, "Data insertion failure."
			}
			count, err := dataAccess.CheckDuplicateMapprocesstoentity(tz, lastinsertedID)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				return 0, false, err, "Data insertion failure."
			}
			if count.Total == 0 {
				_, err := dao.InsertMapprocesstoentitywithtransaction(tx, tz, lastinsertedID)
				if err != nil {
					logger.Log.Println(err)
					tx.Rollback()
					return 0, false, err, "Data insertion failure."
				}
				tx.Commit()
				return lastinsertedID, true, err, ""
			} else {
				tx.Rollback()
				return 0, false, err, "Data Already Exist."
			}

		} else {
			tx.Rollback()
			return 0, false, err, "Data Already Exist."
		}

	} else {
		return 0, false, err, "Process Already Exist with this property type and value"
	}
}

func GetAllMstprocess(page *entities.MstprocessEntity) (entities.MstprocessEntities, bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	t := entities.MstprocessEntities{}
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
	values, err1 := dataAccess.GetAllMstprocess(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstprocessCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstprocess(tz *entities.MstprocessEntity) (bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstprocess(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func DeleteMstprocesswithtransaction(tz *entities.MstprocessEntity) (bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	tx, err := dbcon.Begin()
	if err != nil {
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}
	err3 := dao.DeleteMstprocesswithtransaction(tx, tz)
	if err3 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Data deletion failure."
	}
	err1 := dao.DeleteMstprocessrecordmapwithtransaction(tx, tz)
	if err1 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Data deletion failure."
	}
	err2 := dao.DeleteMapprocesstoentitywithtransaction(tx, tz)
	if err2 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		return false, err, "Data deletion failure."
	}
	pstate := entities.MapprocessstateEntity{}
	pstate.Processid = tz.Id
	err4 := dao.Deletestatebyprocessid(&pstate, tx)
	if err4 != nil {
		logger.Log.Println(err4)
		tx.Rollback()
		return false, err, "Data deletion failure."
	}
	wentity := entities.WorkflowUtilityEntity{}
	wentity.Processid = tz.Id
	err = dao.Deleteprocessdetails(&wentity, tx)
	if err != nil {
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	err = dao.Deletetprocesstransition(&wentity, tx)
	if err != nil {
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	err = dao.Deleteprocessgroupdetails(&wentity, tx)
	if err != nil {
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	tx.Commit()
	return true, nil, ""
}

func UpdateMstprocess(tz *entities.MstprocessEntity) (bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstprocess(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstprocess(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

func UpdateMstprocesswithtransaction(tz *entities.MstprocessEntity) (bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	//dbcon, err := config.ConnectMySqlDb()
	var updatecount int64
	//defer dbcon.Close()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	tx, err := dbcon.Begin()
	if err != nil {
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}
	count, err := dao.CheckDuplicateMstprocesswithtransaction(tx, tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dao.UpdateMstprocesswithtransaction(tx, tz)
		if err != nil {
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Something Went Wrong"
		}
	} else {
		//return false, nil, "Data Already Exist."
		updatecount++
	}

	count1, err1 := dao.CheckDuplicateMstprocessrecordmapwithtransaction(tx, tz, tz.Id)
	if err1 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Something Went Wrong"
	}
	if count1.Total == 0 {
		err := dao.UpdateMstprocessrecordmapwithtransaction(tx, tz)
		if err != nil {
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Something Went Wrong"
		}
	} else {
		//return false, nil, "Data Already Exist."
		updatecount++
	}

	count2, err2 := dao.CheckDuplicateMapprocesstoentitywithtransaction(tx, tz, tz.Id)
	if err2 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Something Went Wrong"
	}
	if count2.Total == 0 {
		err := dao.UpdateMapprocesstoentitywithtransaction(tx, tz)
		if err != nil {
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Something Went Wrong"
		}
	} else {
		//return false, nil, "Data Already Exist."
		updatecount++
	}

	if updatecount == 3 {
		return false, nil, "Data Already Exist."
	}
	tx.Commit()
	return true, err, ""
}
