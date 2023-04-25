package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

//AddClientUserRole for implements business logic
func AddClientUserRole(tz *entities.MapClientUserRoleUserEntity) (int64, bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMapRoleUser(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total > 0 {
		return 0, false, nil, "Mapping already exist."
	}

	id, err := dataAccess.InsertMapRoleUserData(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}

	return id, true, err, ""
}

//DeleteClientUserRole for implements business logic
func DeleteClientUserRole(tz *entities.MapClientUserRoleUserEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMapRoleUserData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//UpdateClientUserRole for implements business logic
func UpdateClientUserRole(tz *entities.MapClientUserRoleUserEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateMapRoleUserData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//GetAllClientUserRole for implements business logic
func GetAllClientUserRole(tz *entities.MapClientUserRoleUserEntity) (entities.MapClientUserRoleUserEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MapClientUserRoleUserEntities{}
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
	values, err1 := dataAccess.GetAllMapRoleUser(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetMapRoleUserCount(tz, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
