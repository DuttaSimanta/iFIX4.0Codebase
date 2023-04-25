package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMstslaresponsiblesupportgroup(tz *entities.MstslaresponsiblesupportgroupEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstslaresponsiblesupportgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstslaresponsiblesupportgroup(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertMstslaresponsiblesupportgroup(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMstslaresponsiblesupportgroup(page *entities.MstslaresponsiblesupportgroupEntity) (entities.MstslaresponsiblesupportgroupEntities, bool, error, string) {
	logger.Log.Println("In side Mstslaresponsiblesupportgroupmodel")
	t := entities.MstslaresponsiblesupportgroupEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllMstslaresponsiblesupportgroup(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstslaresponsiblesupportgroupCount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstslaresponsiblesupportgroup(tz *entities.MstslaresponsiblesupportgroupEntity) (bool, error, string) {
	logger.Log.Println("In side Mstslaresponsiblesupportgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstslaresponsiblesupportgroup(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstslaresponsiblesupportgroup(tz *entities.MstslaresponsiblesupportgroupEntity) (bool, error, string) {
	logger.Log.Println("In side Mstslaresponsiblesupportgroupmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMstslaresponsiblesupportgroup(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstslaresponsiblesupportgroup(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

func GetAllSlanames(page *entities.MstslaresponsiblesupportgroupEntity) ([]entities.Mstslanames, bool, error, string) {
	logger.Log.Println("In side Mstslaresponsiblesupportgroupmodel")
	t := []entities.Mstslanames{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllSlanames(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

func GetFullfillmentcriteriaid(page *entities.MstslaresponsiblesupportgroupEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mstbusinessmatrixmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	data, err1 := dataAccess.GetFullfillmentcriteriaid(page)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	return data, true, nil, ""
}
