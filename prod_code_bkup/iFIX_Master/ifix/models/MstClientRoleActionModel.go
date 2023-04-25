package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

//AddRoleAction for implements business logic
func AddRoleAction(tz *entities.MstClientRoleActionEntity) (int64, bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRoleAction(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total > 0 {
		return 0, false, nil, "Data Already Exist."
	}
	err = dataAccess.DeleteActionRoleWise(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	for i := 0; i < len(tz.ActionIDs); i++ {
		tz.ActionID = tz.ActionIDs[i]
		dataAccess.InsertRoleActionData(tz)
	}
	return 0, true, err, ""

}

//GetAllRoleActionForClient for implements business logic
func GetAllRoleActionForClient(tz *entities.MstClientRoleActionEntity) (entities.MstClientRoleActionEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstClientRoleActionEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.ClientID, tz.MstorgnhirarchyID)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllRoleActionForClient(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetRoleActionCountForClient(tz, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//GetAllRoleAction for implements business logic
func GetAllRoleAction(tz *entities.MstClientRoleActionEntity) (entities.MstClientRoleActionEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstClientRoleActionEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllRoleAction(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetRoleActionCount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//DeleteRoleAction for implements business logic
func DeleteRoleAction(tz *entities.MstClientRoleActionEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteRoleActionData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//UpdateRoleAction for implements business logic
func UpdateRoleAction(tz *entities.MstClientRoleActionEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdatRoleActionData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//GetAllAction for implements business logic
func GetAllAction() ([]entities.MstActionEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstActionEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	val, err1 := dataAccess.GetMstAction()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return val, true, err, ""
}

//GetRoleWiseAction for implements business logic
func GetRoleWiseAction(tz *entities.MstClientRoleActionEntity) ([]entities.MstActionEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstActionEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	val, err1 := dataAccess.GetRoleWiseAction(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return val, true, err, ""
}
