package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
	"strconv"
	"strings"
)

func Getcatelogrecordmodel(page *entities.CatalogEntity) (entities.RecordEntity, bool, error, string) {
	logger.Log.Println("In side Assetmodel")
	var t []entities.ParentCategoryEntity
	var v entities.ParentCategoryEntity
	var u entities.RecordEntity
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return u, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllParentCategoryDetails(page)
	if err1 != nil {
		return u, false, err1, "Something Went Wrong"
	}
	values1, err2 := dataAccess.GetCatalogTickettype(page)
	if len(values1) > 0 {
		if err2 != nil {
			return u, false, err1, "Something Went Wrong"
		}
		s := strings.Split(values.ID, "->")
		s = append(s, strconv.Itoa(int(page.Id)))
		var ids string = ""
		for i, id := range s {
			if i > 0 {
				ids += ","
			}
			ids += id
		}
		//ids += "," + strconv.Itoa(int(page.Id))
		log.Print("\n\n ids ", ids)
		page.Fromrecorddifftypeid = values1[0].TypeId
		page.Fromrecorddiffid = values1[0].Id
		values2, err3 := dataAccess.Getcatalogdetails(page, ids)
		if err3 != nil {
			return u, false, err3, "Something Went Wrong"
		}
		if len(values2) == 0 {
			return u, false, nil, "No Catalog mapped with category"
		}
		values3, err3 := dataAccess.Getmappedcatalogid(page)
		if err3 != nil {
			return u, false, err3, "Something Went Wrong"
		}
		if len(values3) == 0 {
			return u, false, nil, "No Catalog mapped with category"
		}
		s1 := strings.Split(values.NAME, "->")
		b := len(s1)
		for i := 0; i < b; i++ {
			if values3[0].ID < s[i] {
				v.ID = s[i]
				v.NAME = s1[i]
				t = append(t, v)
			}
		}
		u = entities.RecordEntity{t, values1[0], values2[0]}
		return u, true, err, ""
	} else {
		return u, false, nil, "No Ticket type mapped with category"
	}

}
func InsertCatalog(tz *entities.CatalogEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Catalogmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateCatalog(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertCatalog(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Module Name Already Exist."
	}
}

func GetAllCatalog(page *entities.CatalogEntity) (entities.CatalogEntities, bool, error, string) {
	logger.Log.Println("In side Catalogmodel")
	t := entities.CatalogEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(page.Clientid, page.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values, err1 := dataAccess.GetAllCatalog(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetCatalogCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
func Getcatalogorgwise(page *entities.CatalogEntity) ([]entities.CatalogEntitySingle, bool, error, string) {
	logger.Log.Println("In side Catalogmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getcatalogorgwise(page)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

func DeleteCatalog(tz *entities.CatalogEntity) (bool, error, string) {
	logger.Log.Println("In side Catalogmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteCatalog(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateCatalog(tz *entities.CatalogEntity) (bool, error, string) {
	logger.Log.Println("In side Catalogmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateCatalog(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}
