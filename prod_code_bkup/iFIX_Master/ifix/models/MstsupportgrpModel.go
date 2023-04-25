package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func Insertmstsupportgrp(tz *entities.MstsupportgrpEntity) (int64, bool, error, string) {
	logger.Log.Println("In side mstsupportgrpmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicatemstsupportgrp(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.Insertmstsupportgrp(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllmstsupportgrp(page *entities.MstsupportgrpEntity) (entities.MstsupportgrpEntities, bool, error, string) {
	logger.Log.Println("In side mstsupportgrpmodel")
	t := entities.MstsupportgrpEntities{}
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
	values, err1 := dataAccess.GetAllmstsupportgrp(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetmstsupportgrpCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func Deletemstsupportgrp(tz *entities.MstsupportgrpEntity) (bool, error, string) {
	logger.Log.Println("In side mstsupportgrpmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.Deletemstsupportgrp(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func Updatemstsupportgrp(tz *entities.MstsupportgrpEntity) (bool, error, string) {
	logger.Log.Println("In side mstsupportgrpmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicatemstsupportgrp(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.Updatemstsupportgrp(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

func GetAllmstsupportgrpbycopyable(page *entities.MstsupportgrpEntity) ([]entities.MstsupportgrpbycopyableEntity, bool, error, string) {
	logger.Log.Println("In side mstsupportgrpmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllmstsupportgrpbycopyable(page)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getallcreatedsupportgrp(page *entities.MstsupportgrpEntity) ([]entities.MstsupportgrpbycopyableEntity, bool, error, string) {
	logger.Log.Println("In side mstsupportgrpmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getallcreatedsupportgrp(page)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
