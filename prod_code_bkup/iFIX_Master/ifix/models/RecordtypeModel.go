package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertRecordtype(tz *entities.RecordtypeEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRecordtype(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertRecordtype(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Mapping Already Exist."
	}
}

func GetAllRecordtype(page *entities.RecordtypeEntity) (entities.RecordtypeEntities, bool, error, string) {
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
	values, err1 := dataAccess.GetAllRecordtype(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetRecordtypeCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
func Getlabelbydiffid(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	t := []entities.Recordtypesingleentity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getlabelbydiffid(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getlabelbydiffseq(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	t := []entities.Recordtypesingleentity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getlabelbydiffseq(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getmappeddiffbyseq(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	t := []entities.Recordtypesingleentity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getmappeddiffbyseq(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getlablelmappingbydifftype(page *entities.RecordtypeEntity) ([]entities.Recordtypesingleentity, bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	t := []entities.Recordtypesingleentity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getparentid(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		page.Parentid = values[0].Id
		values1, err1 := dataAccess.Getlablelmappingbydifftype(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		return values1, true, nil, ""
	} else {
		return t, false, nil, "No Parent Mapping Found"
	}

}

func DeleteRecordtype(tz *entities.RecordtypeEntity) (bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteRecordtype(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateRecordtype(tz *entities.RecordtypeEntity) (bool, error, string) {
	logger.Log.Println("In side Recordtypemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRecordtype(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	logger.Log.Println("In side Update Recordtypemodel Count value ------------>", count.Total)
	if count.Total == 0 {
		err1 := dataAccess.UpdateRecordtype(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		return true, nil, ""
	} else {
		return false, nil, "Mapping Already Exist."
	}
}
