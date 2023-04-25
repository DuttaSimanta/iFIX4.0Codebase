package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertRecordTermAdditionalMap(tz *entities.RecordTermAdditionalMapEntity) (int64, bool, error, string) {
	logger.Log.Println("In side RecordTermAdditionalMapmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRecordTermAdditionalMap(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	seq, err := dataAccess.GetRecordTermAdditionalMapMaxSeq(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	tz.DisplaySeq = seq + 1
	if count.Total == 0 {
		id, err := dataAccess.InsertRecordTermAdditionalMap(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllRecordTermAdditionalMap(page *entities.RecordTermAdditionalMapEntity) (entities.RecordTermAdditionalMapEntities, bool, error, string) {
	logger.Log.Println("In side RecordTermAdditionalMapmodel")
	t := entities.RecordTermAdditionalMapEntities{}
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
	values, err1 := dataAccess.GetAllRecordTermAdditionalMap(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetRecordTermAdditionalMapCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteRecordTermAdditionalMap(tz *entities.RecordTermAdditionalMapEntity) (bool, error, string) {
	logger.Log.Println("In side RecordTermAdditionalMapmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteRecordTermAdditionalMap(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func GetAdditionalTab() ([]entities.AdditionalTabEntity, bool, error, string) {
	logger.Log.Println("In side AdditionalTabEntitymodel")
	t := []entities.AdditionalTabEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAdditionalTab()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
