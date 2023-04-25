//SearchUser for implements business logic
package models

import (
	"src/config"
	"src/dao"
	"src/entities"
	"src/logger"
)

func SearchLocation(tz *entities.LocationPriorityEntity) ([]entities.LocationSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.LocationSearchEntity{}
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	db, err := config.GetDB()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//dataAccess := dao.DbConn{DB: db}
	values, err1 := dao.SearchLocation(tz, db)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

func SelectLocation(tz *entities.LocationPriorityEntity) ([]entities.LocationSelectEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.LocationSelectEntity{}
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	db, err := config.GetDB()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//dataAccess := dao.DbConn{DB: db}
	values, err1 := dao.SelectLocation(tz, db)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func AddLocation(tz *entities.LocationPriorityEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	db, err := config.GetDB()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	// dataAccess := dao.DbConn{DB: db}
	count, err := dao.CheckDuplicateLocation(tz, db)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dao.AddLocation(tz, db)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllLocation(page *entities.LocationPriorityEntity) (entities.LocationEntities, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	t := entities.LocationEntities{}
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	db, err := config.GetDB()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	// dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dao.GetOrgnType(page.ClientID, page.MstorgnhirarchyID, db)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dao.GetAllLocation(page, orgntype, db)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dao.GetLocationCount(page, orgntype, db)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
func DeleteLocation(tz *entities.LocationPriorityEntity) (bool, error, string) {
	logger.Log.Println("In side Locationmodel")
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	db, err := config.GetDB()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	// dataAccess := dao.DbConn{DB: db}
	err1 := dao.DeleteLocation(tz, db)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateLocation(tz *entities.LocationPriorityEntity) (bool, error, string) {
	logger.Log.Println("In side Locationmodel")
	// lock.Lock()
	// defer lock.Unlock()
	// db, err := config.ConnectMySqlDbSingleton()
	db, err := config.GetDB()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	// dataAccess := dao.DbConn{DB: db}
	count, err := dao.CheckDuplicateLocationforupdate(tz, db)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dao.UpdateLocation(tz, db)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
