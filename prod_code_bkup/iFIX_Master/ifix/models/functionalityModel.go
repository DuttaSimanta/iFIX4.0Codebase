package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

func Getfunctionality() ([]entities.FuncmasterEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.FuncmasterEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getfunctionality()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getfuncmappingbytype(tz *entities.FuncmappingEntity) ([]entities.FuncmappingsingleRespEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.FuncmappingsingleRespEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getfuncmappingbytype(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getfuncmappingbycatalogtype(tz *entities.FuncmappingEntity) ([]entities.FuncmappingsingleRespEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.FuncmappingsingleRespEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getfuncmappingbycatalogtype(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getfuncmappingbytypeforquery(tz *entities.FuncmappingEntity) ([]entities.FuncmappingsingleRespEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.FuncmappingsingleRespEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getfuncmappingbytypeforquery(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Insertfuncmapping(tz *entities.FuncmappingEntity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.Checkduplicatefuncmapping(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		if tz.Funcdescid == 0 {
			desc, err := dataAccess.Getlastfuncdescid(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			if len(desc) == 0 {
				tz.Funcdescid = 1
			} else {
				tz.Funcdescid = desc[0].Funcdescid + 1
			}
		}
		id, err := dataAccess.Insertfuncmapping(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}

		return id, true, err, ""
	} else {
		return 0, false, nil, "Module Name Already Exist."
	}
}

/**
Get Functinality mapping for a all client
*/
func Getfuncmappingdetails(tz *entities.FuncmappingEntity) (entities.FuncmappingEntitities, bool, error, string) {
	log.Println("In side model")
	t := entities.FuncmappingEntitities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getfuncmappingdetails(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.Getmappingdetailscount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

/**
Delete a functionality mapping
*/
func Deletefunctionmapping(tz *entities.FuncmappingEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.Deletefunctionmapping(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}
