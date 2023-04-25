package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertTaskmapping(tz *entities.RecordtypeEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error in AddRecordModelAction", err)
		return 0, false, err, "Something Went Wrong"
	}
	var id int64
	//dataAccess := dao.DbConn{DB: db}
	// count, err := dataAccess.CheckDuplicateRecordtype(tz)
	// if err != nil {
	// 	tx.Rollback()
	// 	db.Close()
	// 	return 0, false, err, "Something Went Wrong"
	// }
	//if count.Total == 0 {for
	for i := 0; i < len(tz.Fromrecorddiffids); i++ {

		id, err := dao.InsertRecordtypetran(tx, tz, i)
		if err != nil {
			tx.Rollback()
			//db.Close()
			return 0, false, err, "Something Went Wrong"
		}
		if id > 0 {
			_, err := dao.Insertextrafieldvalue(tx, tz.Clientid, tz.Mstorgnhirarchyid, id, tz.Title, tz.Description)
			if err != nil {
				tx.Rollback()
				//db.Close()
				return 0, false, err, "Something Went Wrong"
			}

			//err = tx.Commit()
			// if err != nil {
			// 	logger.Log.Println(err)
			// 	tx.Rollback()
			// 	db.Close()
			// 	return 0, false, err, "Something Went Wrong"
			// }
		} else {
			tx.Rollback()
			//db.Close()
			return 0, false, err, "Something Went Wrong"
		}
	}
	err = tx.Commit()
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		//db.Close()
		return 0, false, err, "Something Went Wrong"
	}
	return id, true, err, ""
	// } else {
	// 	tx.Rollback()
	// 	db.Close()
	// 	return 0, false, nil, "Data Already Exist."
	// }
}

func GetAllTaskmap(page *entities.RecordtypeEntity) (entities.RecordtypeEntities, bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	t := entities.RecordtypeEntities{}
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
	values, err1 := dataAccess.GetAllTaskmapvalues(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GettaskmapCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteTaskmap(tz *entities.RecordtypeEntity) (bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Error is ----------->", err)
		return false, err, "Something Went Wrong"
	}
	err1 := dao.DeleteTaskmap(tx, tz)
	if err1 != nil {
		tx.Rollback()
		//db.Close()
		return false, err1, "Something Went Wrong"
	}
	err2 := dao.DeleteTaskProperty(tx, tz)
	if err2 != nil {
		tx.Rollback()
		//db.Close()
		return false, err2, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		//db.Close()
		return false, err, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateTaskmap(tz *entities.RecordtypeEntity) (bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Error is >>>>>>>>>>", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRecordtype(tz)
	if err != nil {
		tx.Rollback()
		//db.Close()
		return false, err, "Something Went Wrong"
	}
	count1, err := dataAccess.CheckDuplicatetaskproperty(tz)
	if err != nil {
		tx.Rollback()
		//db.Close()
		return false, err, "Something Went Wrong"
	}

	logger.Log.Println("In side Update Recordtypemodel Count value ------------>", count.Total, count1.Total)
	if count.Total == 0 && count1.Total == 0 {
		err1 := dao.UpdateRecordtype(tx, tz)
		if err1 != nil {
			tx.Rollback()
			//db.Close()
			return false, err1, "Something Went Wrong"
		}
		err2 := dao.UpdateTaskFieldValue(tx, tz)
		if err2 != nil {
			tx.Rollback()
			//db.Close()
			return false, err2, "Something Went Wrong"
		}
		err = tx.Commit()
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			//db.Close()
			return false, err2, "Something Went Wrong"
		}
		//db.Close()
		return true, nil, ""
	} else {
		tx.Rollback()
		//db.Close()
		return false, nil, "Data Already Exist."
	}
}
