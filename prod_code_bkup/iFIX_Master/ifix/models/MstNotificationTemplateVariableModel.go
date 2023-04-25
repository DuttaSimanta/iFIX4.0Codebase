package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func AddMstTemplateVariablecopy(tz *entities.MstTemplateVariableEntity) (int64, bool, error, string) {
	logger.Log.Println("In side MstTemplateVariable model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	dataAccess1 := dao.TxConn{TX: tx}
	var id int64
	var k int
	for i := 0; i < len(tz.TemplateNames); i++ {
		values, err := dataAccess.GetTemplatevariableCopy(tz, i)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				logger.Log.Print("MstTemplateVariable  Statement Rollback error", err)
				return 0, false, err, ""
			}
			return 0, false, err, "Something Went Wrong"
		}
		values[0].Clientid = tz.ToClientid
		values[0].Mstorgnhirarchyid = tz.ToMstorgnhirarchyid
		count, err := dataAccess.CheckDuplicateMstTemplateVariable(&values[0])
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				logger.Log.Print("MstTemplateVariable  Statement Rollback error", err)
				return 0, false, err, ""
			}
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err = dataAccess1.AddMstTemplateVariablecopy(&values[0])
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					logger.Log.Print("MstTemplateVariable  Statement Rollback error", err)
					return 0, false, err, ""
				}
				return 0, false, err, "Something Went Wrong"
			}

		} else {
			k++
		}
	}
	if k < len(tz.TemplateNames) {
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("MstTemplateVariable  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Rollback()
		if err != nil {
			logger.Log.Print("MstTemplateVariable  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}

}
func AddMstTemplateVariable(tz *entities.MstTemplateVariableEntity) (int64, bool, error, string) {
	logger.Log.Println("In side MstTemplateVariable model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstTemplateVariable(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.AddMstTemplateVariable(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}
func GetAllMstTemplateVariable(page *entities.MstTemplateVariableEntity) (entities.MstTemplateVariableEntities, bool, error, string) {
	logger.Log.Println("In side MstTemplateVariable model")
	t := entities.MstTemplateVariableEntities{}
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
	values, err1 := dataAccess.GetAllMstTemplateVariable(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstTemplateVariableCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func GetAllMstTemplateVariableList(page *entities.MstTemplateVariableEntity) ([]entities.MstTemplateVariableEntityList, bool, error, string) {
	logger.Log.Println("In side MstTemplateVariable model")
	t := []entities.MstTemplateVariableEntityList{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllMstTemplateVariableList(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t = values
	return t, true, err, ""
}

func DeleteMstTemplateVariable(tz *entities.MstTemplateVariableEntity) (bool, error, string) {
	logger.Log.Println("In side MstTemplateVariable model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstTemplateVariable(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstTemplateVariable(tz *entities.MstTemplateVariableEntity) (bool, error, string) {
	logger.Log.Println("In side MstTemplateVariable model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstTemplateVariable(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstTemplateVariable(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
