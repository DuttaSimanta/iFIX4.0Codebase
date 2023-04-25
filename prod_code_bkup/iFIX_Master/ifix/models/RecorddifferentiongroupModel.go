package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertRecorddifferentiongroup(tz *entities.RecorddifferentiongroupEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiongroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}

	j := 0
	var id int64
	id = 0
	/* Starting Transaction*/
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	dataAccess1 := dao.TxConn{TX: tx}
	for i := 0; i < len(tz.Mstworkdifferentiationtypeids); i++ {
		count, err := dataAccess.CheckDuplicateRecorddifferentiongroupWithTx(tz, i)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				//log.Print("MoveWorkflow  Statement Commit error", err)
				logger.Log.Print("Recorddifferentiongroupmodel  Statement Rollback error", err)
				return 0, false, err, ""
			}
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total > 0 {
			j++
			// continue;
		} else {
			id, err = dataAccess1.InsertRecorddifferentiongroup(tz, i)
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					//log.Print("MoveWorkflow  Statement Commit error", err)
					logger.Log.Print("Recorddifferentiongroupmodel  Statement Rollback error", err)
					return 0, false, err, ""
				}
				return 0, false, err, "Something Went Wrong"
			}
		}

	}

	if j < len(tz.Mstworkdifferentiationtypeids) {
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Recorddifferentiongroupmodel  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Commit()
		if err != nil {
			// log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Recorddifferentiongroupmodel  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}

}

func GetAllRecorddifferentiongroup(page *entities.RecorddifferentiongroupEntity) (entities.RecorddifferentiongroupEntities, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiongroupmodel")
	t := entities.RecorddifferentiongroupEntities{}
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
	values, err1 := dataAccess.GetAllRecorddifferentiongroup(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetRecorddifferentiongroupCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteRecorddifferentiongroup(tz *entities.RecorddifferentiongroupEntity) (bool, error, string) {
	logger.Log.Println("In side Recorddifferentiongroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteRecorddifferentiongroup(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateRecorddifferentiongroup(tz *entities.RecorddifferentiongroupEntity) (bool, error, string) {
	logger.Log.Println("In side Recorddifferentiongroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRecorddifferentiongroup(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateRecorddifferentiongroup(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

func GetWorkinglevel(page *entities.RecorddifferentiongroupEntity) (entities.WorkinglevelEntities, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiongroupmodel")
	t := entities.WorkinglevelEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetWorkinglevel(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	t.Values = values
	return t, true, err, ""
}
