package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMapexternalattributes(tz *entities.MapexternalattributesEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Mapexternalattributesmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	var id int64
	id = 0
	/* Starting Transaction*/
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	// dataAccess := dao.DbConn{DB: db}
	dataAccess1 := dao.TxConn{TX: tx}
	for i := 0; i < len(tz.Map); i++ {
		id, err = dataAccess1.InsertMapexternalattributes(tz, i)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
	}
	err = tx.Commit()
	if err != nil {
		//log.Print("MoveWorkflow  Statement Commit error", err)
		logger.Log.Print("Mapexternalattributes  Statement Commit error", err)
		return 0, false, err, ""
	}
	return id, true, err, ""
}
func GetMappedattributes(page *entities.MapexternalattributesEntity) ([]entities.Attr, bool, error, string) {
	logger.Log.Println("In side Mapexternalattributesmodel")
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetMappedattributes(page)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func GetAllMapexternalattributes(page *entities.MapexternalattributesEntity) (entities.MapexternalattributesEntities, bool, error, string) {
	logger.Log.Println("In side Mapexternalattributesmodel")
	t := entities.MapexternalattributesEntities{}
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
	values, err1 := dataAccess.GetAllMapexternalattributes(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMapexternalattributesCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMapexternalattributes(tz *entities.MapexternalattributesEntity) (bool, error, string) {
	logger.Log.Println("In side Mapexternalattributesmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMapexternalattributes(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}
