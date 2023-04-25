package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/entities"
	"log"
	"iFIX/ifix/dao"
)

func InsertUrl(tz *entities.UrlEntity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	details,err :=dataAccess.CheckDuplicateUrl(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(details) == 0 {
		id, err := dataAccess.InsertUrl(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if id == -1{
			return 0, false, err, "URL already mapped with module"
		}else{
			return id, true, nil, ""
		}
	}else{
		tz.Id=details[0].Id
		id, err := dataAccess.InsertIntoModuleUrl(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if id == -1{
			return 0, false, err, "URL already mapped with module"
		}else{
			return id, true, nil, ""
		}
	}
}

func GetAllUrls(page *entities.PaginationEntity) (entities.UrlEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.UrlEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllUrls(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0{
		total, err1 := dataAccess.GetUrlsCount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total=total.Total
		t.Values=values
	}
	t.Values=values
	return t, true, nil, ""
}
func GetAllModuleUrls(page *entities.PaginationEntity) (entities.ModuleUrlEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.ModuleUrlEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllModuleUrls(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0{
		total, err1 := dataAccess.GetModuleUrlsCount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total=total.Total
		t.Values=values
	}
	t.Values=values
	return t, true, nil, ""
}
func DeleteUrl(tz *entities.UrlEntity) (bool,error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false,err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteUrl(tz)
	if err1 != nil {
		return false,err1, "Something Went Wrong"
	}
	return true,nil, ""
}
func DeleteModUrl(tz *entities.ModuleUrlEntity) (bool,error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false,err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteModUrl(tz)
	if err1 != nil {
		return false,err1, "Something Went Wrong"
	}
	return true,nil, ""
}
func UpdateUrl(tz *entities.UrlEntity) (bool,error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false,err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateUrl(tz)
	if err1 != nil {
		return false,err1, "Something Went Wrong"
	}
	return true,nil, ""
}

func GetDistinctUrl(tz *entities.UrlEntity) ([]entities.UrlRespEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.UrlRespEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetDistinctUrl(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, nil, ""
}
func GetRemainingUrl(tz *entities.UrlEntity) ([]entities.UrlRespEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.UrlRespEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetRemainingUrl(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, nil, ""
}


