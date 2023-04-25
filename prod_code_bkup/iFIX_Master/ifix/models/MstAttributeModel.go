package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func AddMstAttribute(tz *entities.MstAttributeEntity) (int64, bool, error, string) {
	logger.Log.Println("In side AddMstAttributeCopymodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}

	dataAccess := dao.DbConn{DB: db}

	count, err := dataAccess.CheckDuplicateMstAttribute(tz)
	if err != nil {

		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {

		id, err := dataAccess.AddMstAttribute(tz)
		if err != nil {

			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	}
	return 0, false, err, "Data Already Exist"

}

func GetAllMstAttribute(page *entities.MstAttributeEntity) (entities.MstAttributeEntities, bool, error, string) {
	logger.Log.Println("In side MstAttribute model")
	t := entities.MstAttributeEntities{}
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
	values, err1 := dataAccess.GetAllMstAttribute(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstAttributeCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstAttribute(tz *entities.MstAttributeEntity) (bool, error, string) {
	logger.Log.Println("In side MstAttribute model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstAttribute(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstAttribute(tz *entities.MstAttributeEntity) (bool, error, string) {
	logger.Log.Println("In side MstAttribute model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	count, err := dataAccess.CheckDuplicateMstAttribute(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstAttribute(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}

}

func GetMstAttribute(page *entities.MstAttributeEntity) ([]entities.Attributes, bool, error, string) {
	logger.Log.Println("In side GetGetMstAttributemodel")
	t := []entities.Attributes{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	t, err1 := dataAccess.GetMstAttribute(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return t, true, err, ""
}
