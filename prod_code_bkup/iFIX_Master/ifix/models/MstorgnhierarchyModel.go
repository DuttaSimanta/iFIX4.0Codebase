package models

import (
	"encoding/json"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func UpdateOrganizationField(tz *entities.MstorgnhierarchyFieldEntity) (bool, error, string) {
	input, _ := json.Marshal(tz)
	logger.Log.Println(string(input))
	logger.Log.Println("In side UpdateOrganizationField model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateOrganizationField(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func GetAllOrganizationField(tz *entities.MstorgnhierarchyFieldEntity) (entities.MstorgnhierarchyFieldEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstorgnhierarchyFieldEntities{}
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.ClientID, tz.ID)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllOrganizationField(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	logger.Log.Println(values)

	if tz.Offset == 0 {
		total, err1 := dataAccess.GetOrganizationFieldCount(tz, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//AddOrganizations for implements business logic
func AddOrganizations(tz *entities.MstorgnhierarchyEntity) (int64, bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateOrganization(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total > 0 {
		return 0, false, nil, "Client Already Exist."
	}

	id, err := dataAccess.InsertOrganization(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}

	return id, true, err, ""
}

//UpdateOrganizations for implements business logic
func UpdateOrganizations(tz *entities.MstorgnhierarchyEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateOrganization(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//GetAllOrganizations for implements business logic
func GetAllOrganizations(tz *entities.MstorgnhierarchyEntity) (entities.MstorgnhierarchyEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstorgnhierarchyEntities{}
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllOrganization(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	logger.Log.Println(values)

	if tz.Offset == 0 {
		total, err1 := dataAccess.GetOrganizationCount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//GetAllOrganizationsClientWise for implements business logic
func GetAllOrganizationsClientWise(tz *entities.MstorgnhierarchyEntity) ([]entities.MstorgnhierarchyEntityResp, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstorgnhierarchyEntityResp{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllOrganizationClientWise(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

//GetAllOrganizationsClientWise for implements business logic
func GetAllOrganizationsClientWisenew(tz *entities.MstorgnhierarchyEntity) ([]entities.MstorgnhierarchyEntityResp, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstorgnhierarchyEntityResp{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(tz.ClientID, tz.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllOrganizationClientWisenew(tz, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

func Gettimeformat() ([]entities.MstorgnhierarchyEntityResp, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstorgnhierarchyEntityResp{}
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Gettimeformat()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func GetLogintype() ([]entities.LogintypeEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.LogintypeEntity{}
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetLogintype()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
