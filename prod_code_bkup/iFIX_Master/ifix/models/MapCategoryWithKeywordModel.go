package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertMapCategoryWithKeyword(tz *entities.MapCategoryWithKeywordEntity) (int64, bool, error, string) {
	logger.Log.Println("In side MapCategoryWithKeywordmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMapCategoryWithKeyword(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess1 := dao.TxConn{TX: tx}

	if count.Total == 0 {
		count1, err1 := dataAccess.CheckDuplicateMapCategoryWithKeywordforstupa(tz)
		if err1 != nil {
			return 0, false, err, "Something Went Wrong"
		}
		clientid := tz.Clientid
		orgnid := tz.Mstorgnhirarchyid
		if count1.Total == 0 {
			tz.Clientid = 1
			tz.Mstorgnhirarchyid = 1
			_, err := dataAccess1.InsertMapCategoryWithKeyword(tz)
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					//log.Print("MoveWorkflow  Statement Commit error", err)
					logger.Log.Print("MapCategoryWithKeywordmodel  Statement Commit error", err)
					return 0, false, err, ""
				}
				return 0, false, err, "Something Went Wrong"
			}
		}
		tz.Clientid = clientid
		tz.Mstorgnhirarchyid = orgnid
		id, err := dataAccess1.InsertMapCategoryWithKeyword(tz)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				//log.Print("MoveWorkflow  Statement Commit error", err)
				logger.Log.Print("MapCategoryWithKeywordmodel  Statement Commit error", err)
				return 0, false, err, ""
			}
			return 0, false, err, "Something Went Wrong"
		}
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("MapCategoryWithKeywordmodel  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("MapCategoryWithKeywordmodel  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMapCategoryWithKeyword(page *entities.MapCategoryWithKeywordEntity) (entities.MapCategoryWithKeywordEntities, bool, error, string) {
	logger.Log.Println("In side MapCategoryWithKeywordmodel")
	t := entities.MapCategoryWithKeywordEntities{}
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
	values, err1 := dataAccess.GetAllMapCategoryWithKeyword(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMapCategoryWithKeywordCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMapCategoryWithKeyword(tz *entities.MapCategoryWithKeywordEntity) (bool, error, string) {
	logger.Log.Println("In side MapCategoryWithKeywordmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMapCategoryWithKeyword(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMapCategoryWithKeyword(tz *entities.MapCategoryWithKeywordEntity) (bool, error, string) {
	logger.Log.Println("In side MapCategoryWithKeywordmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMapCategoryWithKeyword(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess1 := dao.TxConn{TX: tx}

	if count.Total == 0 {
		count1, err1 := dataAccess.CheckDuplicateMapCategoryWithKeywordforstupa(tz)
		if err1 != nil {
			return false, err, "Something Went Wrong"
		}
		clientid := tz.Clientid
		orgnid := tz.Mstorgnhirarchyid
		if count1.Total == 0 {
			tz.Clientid = 1
			tz.Mstorgnhirarchyid = 1
			_, err := dataAccess1.InsertMapCategoryWithKeyword(tz)
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					// log.Print("MoveWorkflow  Statement Commit error", err)
					logger.Log.Print("MapCategoryWithKeywordmodel  Statement Commit error", err)
					return false, err, ""
				}
				return false, err, "Something Went Wrong"
			}
		}
		tz.Clientid = clientid
		tz.Mstorgnhirarchyid = orgnid
		err := dataAccess1.UpdateMapCategoryWithKeyword(tz)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				// log.Print("MoveWorkflow  Statement Commit error", err)
				logger.Log.Print("MapCategoryWithKeywordmodel  Statement Commit error", err)
				return false, err, ""
			}
			return false, err, "Something Went Wrong"
		}
		err = tx.Commit()
		if err != nil {
			// log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("MapCategoryWithKeywordmodel  Statement Commit error", err)
			return false, err, ""
		}
		return true, err, ""
	} else {
		err = tx.Commit()
		if err != nil {
			// log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("MapCategoryWithKeywordmodel  Statement Commit error", err)
			return false, err, ""
		}
		return false, nil, "Data Already Exist."
	}
}

// func GetAllkeyword(page *entities.MapCategoryWithKeywordEntity) ([]entities.Getkeyword, bool, error, string) {
// 	logger.Log.Println("In side MapCategoryWithKeywordmodel")
// 	t := []entities.Getkeyword{}
// 	db, err := config.ConnectMySqlDb()
// 	defer db.Close()
// 	if err != nil {
// 		logger.Log.Println("database connection failure", err)
// 		return t, false, err, "Something Went Wrong"
// 	}
// 	dataAccess := dao.DbConn{DB: db}
// 	values, err1 := dataAccess.GetAllkeyword(page)
// 	if err1 != nil {
// 		return t, false, err1, "Something Went Wrong"
// 	}
// 	t = values
// 	return t, true, err, ""
// }
func GetAllCategoryvalue(page *entities.MapCategoryWithKeywordEntity) ([]entities.Getkeyword, bool, error, string) {
	logger.Log.Println("In side MapCategoryWithKeywordmodel")
	t := []entities.Getkeyword{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllCategoryvalue(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t = values
	return t, true, err, ""
}
