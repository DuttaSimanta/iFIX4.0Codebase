package models

import (
	"database/sql"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

/**
Create  menu for client
*/
func InsertMenuwithapi(tz *entities.MenuEntity) (int64, bool, error, string) {
	log.Println("In side model")
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	_, success, err, msg := InsertMenu(tz, db)
	return 0, success, err, msg
}
func InsertMenu(tz *entities.MenuEntity, db *sql.DB) (int64, bool, error, string) {
	log.Println("In side model")
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	//lock.Lock()
	//defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	count, err := dataAccess.Checkduplicatemenu(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		tx, err := db.Begin()
		if err != nil {
			logger.Log.Println("database transaction connection failure", err)
			return 0, false, err, "Something Went Wrong"
		}
		dataAccess1 := dao.TxConn{TX: tx}
		_, err1 := dataAccess1.IncrementMenuseq(tz)
		if err1 != nil {
			tx.Rollback()
			return 0, false, err1, "Something Went Wrong"
		}
		id, err := dataAccess1.InsertMenu(tz)
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Something Went Wrong"
		}
		tx.Commit()
		return id, true, err, ""
	} else {
		return 0, false, nil, "Menu Already Exist."
	}
}
func GetMenuByUser(page *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MenuHierarchyEntity{}
	t1 := []entities.MenuHierarchyEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	t, err1 := dataAccess.GetMenuByUserID(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	if len(t) > 0 {
		t1 = entities.GetRootMenu(t)
		return t1, true, err1, ""
	} else {
		t2, err1 := dataAccess.GetMenuByUserIDNRole(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t1 = entities.GetRootMenu(t2)
		return t1, true, err1, ""
	}
	return t, true, err1, ""
}

/**
Search menu by keyword for a particular user using permission
*/
func SearchMenuByUser(page *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, bool, error, string) {
	log.Println("-------------In side model SearchMenuByUser  --------------")
	logger.Log.Println("-------------In side model SearchMenuByUser  --------------")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	t, err1 := dataAccess.SearchMenuByUserID(page)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	if len(t) > 0 {
		logger.Log.Println("-------------Sending SearchMenuByUser 1 --------------")
		return t, true, err1, ""
	} else {
		t2, err1 := dataAccess.SearchMenuByUserIDNRole(page)
		if err1 != nil {
			return nil, false, err1, "Something Went Wrong"
		}
		logger.Log.Println("-------------Sending SearchMenuByUser 2 --------------")
		return t2, true, err1, ""
	}
}

//func Checkmenupermissionforuser(page *entities.MenuByUserRequest) ([]entities.MenuHierarchyEntity, bool, error, string) {
//	log.Println("In side model")
//	lock.Lock()
//	defer lock.Unlock()
//	db, err := config.ConnectMySqlDbSingleton()
//	if err != nil {
//		log.Println("database connection failure", err)
//		return nil, false, err, "Something Went Wrong"
//	}
//	dataAccess := dao.DbConn{DB: db}
//	t, err1 := dataAccess.Getsinglemenubyuserid(page)
//	if err1 != nil {
//		return nil, false, err1, "Something Went Wrong"
//	}
//	if len(t)>0{
//		return t, true, err1, ""
//	} else {
//		t2, err1 := dataAccess.SearchMenuByUserIDNRole(page)
//		if err1 != nil {
//			return nil, false, err1, "Something Went Wrong"
//		}
//		return t2, true, err1, ""
//	}
//}

/**
get menus with no child menu
*/
func Getparentmenu(tz *entities.MenuEntity) ([]entities.MenuSingleEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MenuSingleEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getparentmenu(tz)
	log.Println("\n\n-----------", err1)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	} else {
		return values, true, nil, ""
	}

}

/**
get menus by module id
*/
func Getmenubymodule(tz *entities.MenuEntity) ([]entities.MenuSingleEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.MenuSingleEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getmenubymodule(tz)
	log.Println("\n\n-----------", err1)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	} else {
		return values, true, nil, ""
	}

}

/**
Get all url menu for base client
*/
func Getmenudetails(page *entities.PaginationEntity) (entities.MenuEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.MenuEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getmenudetails(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.Getmenucount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err1, ""
}

/**
Delete a menu from base client
*/
func DeleteUrlFromMenu(tz *entities.MenuEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteUrlFromMenu(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

/**
Delete a menu from base client
*/
func DeleteMenu(tz *entities.MenuEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMenu(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

/**
Update a menu from base client
*/
func UpdateMenu(tz *entities.MenuEntity) (bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	if tz.Urlid == 0 {
		err1 := dataAccess.UpdateMenu(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
	} else {
		err1 := dataAccess.Addurltomenu(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
	}

	return true, nil, ""
}

/**
Get all url menu for base client
*/
func Geturlmenudetails(page *entities.PaginationEntity) (entities.MenuEntities, bool, error, string) {
	log.Println("In side model")
	t := entities.MenuEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Geturlmenudetails(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.Geturlmenucount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err1, ""
}
