package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

func InsertCatalogwithcategory(tz *entities.CatalogwithcategoryEntity) (int64, bool, error, string) {
	logger.Log.Println("In side InsertCatalogwithcategory")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	j := 0
	var id int64
	id = 0
	/* Starting Transaction*/
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	dataAccess1 := dao.TxConn{TX: tx}
	vals, err2 := dataAccess.Getalltorecorddiffids(tz)
	if err2 != nil {
		return 0, false, err2, "Something Went Wrong"
	}
	x := 0
	m := 0
	tz.Torecorddiffids = vals.Torecorddiffids
	for j := 0; j < len(tz.Torecorddiffids); j++ {
		for k := 0; k < x; k++ {
			if tz.Torecorddiffids[k] == tz.Torecorddiffids[j] {
				m = 1
			}
		}
		if m == 0 {
			tz.Torecorddiffids[x] = tz.Torecorddiffids[j]
			x++
		}
		m = 0
	}
	for k := 0; k < x; k++ {
		count, err := dataAccess.CheckDuplicateCatalogwithcategory(tz, k)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total > 0 {
			j++
		} else {
			tz.Forrecorddiffid = tz.Torecorddiffid
			id, err = dataAccess1.InsertCatalogwithcategory(tz, k)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
		}
	}
	if j < x {
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("InsertCatalogwithcategory  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("InsertCatalogwithcategory  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllCatalogwithcategory(page *entities.CatalogwithcategoryEntity) (entities.CatalogwithcategoryEntities, bool, error, string) {
	logger.Log.Println("In side Catalogwithcategorymodel")
	t := entities.CatalogwithcategoryEntities{}
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
	values, err1 := dataAccess.GetAllCatalogwithcategory(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetCatalogwithcategoryCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
func Getcategorybycatalog(page *entities.CatalogwithcategoryEntity) ([]entities.CatalogwithsingleEntity, bool, error, string) {
	logger.Log.Println("In side Catalogwithcategorymodel")
	t := []entities.CatalogwithsingleEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getcategorybycatalog(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
func Getcategorybyparentname(page *entities.CatalogwithcategoryEntity) ([]entities.CatalogwithsingleEntity, bool, error, string) {
	logger.Log.Println("In side Getcategorybyparentname")
	t := []entities.CatalogwithsingleEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getcategorybyparentname(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
func Getfromtypebydiffname(page *entities.CatalogwithcategoryEntity) ([]entities.CatalogwithsingleEntity, bool, error, string) {
	logger.Log.Println("In side Getfromtypebydiffname")
	t := []entities.CatalogwithsingleEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getfromtypebydiffname(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

func DeleteCatalogwithcategory(tz *entities.CatalogwithcategoryEntity) (bool, error, string) {
	logger.Log.Println("In side Catalogwithcategorymodel")
	//db, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	//defer db.Close()
	//dataAccess := dao.DbConn{DB: db}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}
	err1 := dao.DeleteCatalogwithcategory(tx, tz)
	if err1 != nil {
		tx.Rollback()
		return false, err1, "Something Went Wrong"
	}
	err2 := dao.Deletemappedcategory(tx, tz)
	if err2 != nil {
		tx.Rollback()
		return false, err2, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		log.Print("DeleteCatalogwithcategory  Statement Commit error", err)
		logger.Log.Print("DeleteCatalogwithcategory  Statement Commit error", err)
		return false, err, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateCatalogwithcategory(tz *entities.CatalogwithcategoryEntity) (bool, error, string) {
	logger.Log.Println("In side Catalogwithcategorymodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateCatalogwithcategory(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}
