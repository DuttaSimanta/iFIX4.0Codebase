package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"strings"
)

func InsertTransporttable(tz *entities.TransporttableEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Transporttablemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	tz.Msttablename = strings.Trim(tz.Msttablename, " ")
	tz.Tabletypedescription = strings.Trim(tz.Tabletypedescription, " ")
	count, err := dataAccess.CheckDuplicateTransporttable(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	result, err1 := dataAccess.Gettype(tz)
	if err1 != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(result) == 0 {
		maxtype, err1 := dataAccess.Getmaxtype()
		if err1 != nil {
			return 0, false, err, "Something Went Wrong"
		}
		tz.Tabletype = maxtype + 1
	} else {
		tz.Tabletype = result[0].Tabletype
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertTransporttable(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllTransporttable(page *entities.TransporttableEntity) (entities.TransporttableEntities, bool, error, string) {
	logger.Log.Println("In side Transporttablemodel")
	t := entities.TransporttableEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	values, err1 := dataAccess.GetAllTransporttable(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetTransporttableCount(page)

		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
func DeleteTransporttable(tz *entities.TransporttableEntity) (bool, error, string) {
	logger.Log.Println("In side Transporttablemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteTransporttable(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateTransporttable(tz *entities.TransporttableEntity) (bool, error, string) {
	logger.Log.Println("In side Transporttablemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	tz.Msttablename = strings.Trim(tz.Msttablename, " ")
	tz.Tabletypedescription = strings.Trim(tz.Tabletypedescription, " ")
	count, err := dataAccess.CheckDuplicateTransporttableupdate(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	result, err1 := dataAccess.Gettype(tz)
	if err1 != nil {
		return false, err, "Something Went Wrong"
	}
	if len(result) == 0 {
		maxtype, err1 := dataAccess.Getmaxtype()
		if err1 != nil {
			return false, err, "Something Went Wrong"
		}
		tz.Tabletype = maxtype + 1
	} else {
		tz.Tabletype = result[0].Tabletype
	}
	if count.Total == 0 {
		err := dataAccess.UpdateTransporttable(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
func Gettypedescription(tz *entities.TransporttableEntity) ([]entities.GettableEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.GettableEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	val, err1 := dataAccess.Gettypedescription(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return val, true, nil, ""
}
func Gettable(tz *entities.TransporttableEntity) ([]entities.TableEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.TableEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	val, err1 := dataAccess.Gettable(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return val, true, nil, ""
}
func Gettypefortransport(tz *entities.TransporttableEntity) ([]entities.GettableEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.GettableEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	val, err1 := dataAccess.Gettypefortransport(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return val, true, nil, ""
}
