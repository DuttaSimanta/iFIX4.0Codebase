package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"log"
)

/**
Create  module for base client
*/
func InsertModule(tz *entities.ModuleEntity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateModule(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertModule(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}

		return id, true, err, ""
	} else {
		return 0, false, nil, "Module Name Already Exist."
	}
}

/**
Get all modules for base client
*/
func GetAllModules(page *entities.PaginationEntity) (entities.ModuleEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.ModuleEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllModules(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetModuleCount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err1, ""
}

/**
Delete a module from base client
*/
func DeleteModule(tz *entities.ModuleEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteModule(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

/**
Update a particular module for base client
*/
func UpdateModule(tz *entities.ModuleEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateModule(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}
