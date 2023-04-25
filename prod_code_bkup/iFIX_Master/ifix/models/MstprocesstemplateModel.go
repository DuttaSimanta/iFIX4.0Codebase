package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstprocesstemplatewithtransaction(tz *entities.MstprocessEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	// dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	// defer dbcon.Close()
	dataAccess := dao.DbConn{DB: dbcon}
	tx, err := dbcon.Begin()
	if err != nil {
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}
	count, err := dataAccess.CheckDuplicateMstprocesstemplate(tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return 0, false, err, "Data insertion failure."
	}
	if count.Total == 0 {
		lastinsertedID, err := dao.InsertMstprocesstemplatewithtransaction(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}

		/*count, err := dataAccess.CheckDuplicateMstprocesstemplaterecordmap( tz,lastinsertedID)
		if count.Total == 0 {
			_, err := dao.InsertMstprocesstemplaterecordmapwithtransaction(tx, tz, lastinsertedID)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				// dbcon.Close()
				return 0, false, err, "Data insertion failure."
			}*/
		count, err := dataAccess.CheckDuplicateMapprocesstemplatetoentity(tz, lastinsertedID)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}
		if count.Total == 0 {
			_, err := dao.InsertMapprocesstemplatetoentitywithtransaction(tx, tz, lastinsertedID)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				// dbcon.Close()
				return 0, false, err, "Data insertion failure."
			}
			tx.Commit()
			return lastinsertedID, true, err, ""
		} else {
			return 0, false, err, "Process template already mapped with table."
		}

		/*} else {
			return 0, false, err, "Process template already mapped with category"
		}*/
	} else {
		return 0, false, err, "Process Already Exist with this property type and value"
	}
}

func GetAllMstprocesstemplate(page *entities.MstprocessEntity) (entities.MstprocessEntities, bool, error, string) {
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
	values, err1 := dataAccess.GetAllMstprocesstemplate(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstprocesstemplateCount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstprocesstemplatewithtransaction(tz *entities.MstprocessEntity) (bool, error, string) {
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
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}
	err3 := dao.DeleteMstprocesstemplatewithtransaction(tx, tz)
	if err3 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		return false, err, "Data deletion failure."
	}

	err2 := dao.DeleteMapprocesstemplatetoentitywithtransaction(tx, tz)
	if err2 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		return false, err, "Data deletion failure."
	}
	pstate := entities.MapprocessstateEntity{}
	pstate.Processid = tz.Id
	err1 := dao.Deletestatebytemplateid(&pstate, tx)
	if err1 != nil {
		logger.Log.Println(err1)
		tx.Rollback()
		return false, err, "Data deletion failure."
	}
	wentity := entities.WorkflowUtilityEntity{}
	wentity.Processid = tz.Id
	err = dao.Deleteprocesstemplatedetails(&wentity, tx)
	if err != nil {
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	err = dao.Deletetprocesstemplatetransition(&wentity, tx)
	if err != nil {
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	err = dao.Deleteprocesstemplategroupdetails(&wentity, tx)
	if err != nil {
		tx.Rollback()
		return false, err, "Something Went Wrong"
	}
	tx.Commit()
	return true, nil, ""
}

func UpdateMstprocesstemplatewithtransaction(tz *entities.MstprocessEntity) (bool, error, string) {
	logger.Log.Println("In side Mstprocessmodel")
	// dbcon, err := config.ConnectMySqlDb()
	var updatecount int64
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	// defer dbcon.Close()
	tx, err := dbcon.Begin()
	if err != nil {
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: dbcon}
	count, err := dataAccess.CheckDuplicateMstprocesstemplate(tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dao.UpdateMstprocesstemplatewithtransaction(tx, tz)
		if err != nil {
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Something Went Wrong"
		}
	} else {
		updatecount++
	}
	/*count1, err1 := dataAccess.CheckDuplicateMstprocesstemplaterecordmap( tz,tz.Id)
	if err1 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Something Went Wrong"
	}
	if count1.Total == 0 {
		err := dao.UpdateMstprocesstemplaterecordmapwithtransaction(tx, tz)
		if err != nil {
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Something Went Wrong"
		}
	} else {
		//return false, nil, "Data Already Exist."
		updatecount++
	}*/

	count2, err2 := dataAccess.CheckDuplicateMapprocesstemplatetoentity(tz, tz.Id)
	if err2 != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Something Went Wrong"
	}
	if count2.Total == 0 {
		err := dao.UpdateMapprocesstemplatetoentitywithtransaction(tx, tz)
		if err != nil {
			tx.Rollback()
			// dbcon.Close()
			return false, err, "Something Went Wrong"
		}
	} else {
		updatecount++
	}
	if updatecount == 2 {
		return false, nil, "Data Already Exist."
	}
	tx.Commit()
	return true, err, ""
}
