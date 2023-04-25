package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

//AddUserRoleAction for implements business logic
func AddUserRoleAction(tz *entities.MstUserRoleActionEntity) (int64, bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateUserRoleAction(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total > 0 {
		return 0, false, nil, "Data Already Exist."
	}
	err = dataAccess.DeleteActionRoleUserWise(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	for i := 0; i < len(tz.Actionids); i++ {
		tz.ActionID = tz.Actionids[i]
		dataAccess.InsertUserRoleActionData(tz)
	}

	return 0, true, err, ""

}

//GetAllUserRoleActionForClient for implements business logic
func GetAllUserRoleActionForClient(tz *entities.MstUserRoleActionEntity) (entities.MstUserRoleActionEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstUserRoleActionEntities{}
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
	values, err1 := dataAccess.GetAllUserRoleActionForClient(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetUserRoleActionCountForClient(tz, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//GetAllUserRoleAction for implements business logic
func GetAllUserRoleAction(tz *entities.MstUserRoleActionEntity) (entities.MstUserRoleActionEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstUserRoleActionEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllUserRoleAction(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetUserRoleActionCount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//DeleteUserRoleAction for implements business logic
func DeleteUserRoleAction(tz *entities.MstUserRoleActionEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteUserRoleActionData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//UpdateUserRoleAction for implements business logic
func UpdateUserRoleAction(tz *entities.MstUserRoleActionEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdatUserRoleActionData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//GetRoleUserWiseAction for implements business logic
func GetRoleUserWiseAction(tz *entities.MstUserRoleActionEntity) ([]entities.MstActionEntity, bool, error, string) {
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
	val, err1 := dataAccess.GetRoleUserWiseAction(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return val, true, nil, ""
}
