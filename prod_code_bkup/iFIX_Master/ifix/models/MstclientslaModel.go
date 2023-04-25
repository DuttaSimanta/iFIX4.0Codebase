package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstclientsla(tz *entities.MstclientslaEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstclientslamodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstclientsla(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertMstclientsla(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMstclientsla(page *entities.MstclientslaEntity) (entities.MstclientslaEntities, bool, error, string) {
	logger.Log.Println("In side Mstclientslamodel")
	t := entities.MstclientslaEntities{}
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
	values, err1 := dataAccess.GetAllMstclientsla(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstclientslaCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstclientsla(tz *entities.MstclientslaEntity) (bool, error, string) {
	logger.Log.Println("In side Mstclientslamodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstclientsla(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstclientsla(tz *entities.MstclientslaEntity) (bool, error, string) {
	logger.Log.Println("In side Mstclientslamodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstclientsla(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstclientsla(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

func GetSlanames(page *entities.MstclientslaEntity) ([]entities.Mstslaname, bool, error, string) {
	logger.Log.Println("In side Mstclientslamodel")
	t := []entities.Mstslaname{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetSlanames(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
