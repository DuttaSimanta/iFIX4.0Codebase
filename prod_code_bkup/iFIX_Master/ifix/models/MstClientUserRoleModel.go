package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

//AddRoleModel for implements business logic
func AddRoleModel(tz *entities.MstClientUserRoleEntity) (int64, bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	tz.ID = 0
	count, err := dataAccess.CheckDuplicateCientRoles(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total > 0 {
		return 0, false, nil, "Role Name Already Exist."
	}

	id, err := dataAccess.InsertRoleData(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}

	return id, true, err, ""

}

//GetAllRoleModel for implements business logic
func GetAllRoleModel(page *entities.MstClientUserRoleEntity) (entities.MstClientUserRoleEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstClientUserRoleEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(page.ClientID, page.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllRoleData(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetRoleCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//Getrolebyorgid for implements business logic
func Getrolebyorgid(page *entities.MstClientUserRoleEntity) ([]entities.MstClientRoleEntityResp, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstClientRoleEntityResp{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getrolebyorgid(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

//DeleteRoleModel for implements business logic
func DeleteRoleModel(tz *entities.MstClientUserRoleEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteRoleData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//UpdateRoleModel for implements business logic
func UpdateRoleModel(tz *entities.MstClientUserRoleEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateCientRoles(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	count1, err := dataAccess.CheckDuplicateCientRolesforupdate(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 && count1.Total == 0 {

		err1 := dataAccess.UpdateRoleData(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		return true, nil, ""
	}
	return false, nil, "Data Already Exist."

}
