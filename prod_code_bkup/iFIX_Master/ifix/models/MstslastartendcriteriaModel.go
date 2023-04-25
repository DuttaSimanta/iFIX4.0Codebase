package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstslastartendcriteria(tz *entities.MstslastartendcriteriaEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstslastartendcriteriamodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstslastartendcriteria(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertMstslastartendcriteria(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMstslastartendcriteria(page *entities.MstslastartendcriteriaEntity) (entities.MstslastartendcriteriaEntities, bool, error, string) {
	logger.Log.Println("In side Mstslastartendcriteriamodel")
	t := entities.MstslastartendcriteriaEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllMstslastartendcriteria(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstslastartendcriteriaCount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstslastartendcriteria(tz *entities.MstslastartendcriteriaEntity) (bool, error, string) {
	logger.Log.Println("In side Mstslastartendcriteriamodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstslastartendcriteria(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstslastartendcriteria(tz *entities.MstslastartendcriteriaEntity) (bool, error, string) {
	logger.Log.Println("In side Mstslastartendcriteriamodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstslastartendcriteria(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstslastartendcriteria(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

func GetSlanameagainstworkflowid(page *entities.MstslastartendcriteriaEntity) ([]entities.MstslanameagaistworkflowEntity, bool, error, string) {
	logger.Log.Println("In side GetSlanameagainstworkflowid")
	t := []entities.MstslanameagaistworkflowEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetSlanameagainstworkflowid(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
