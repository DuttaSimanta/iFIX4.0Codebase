package models

import (
	"database/sql"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
	"strings"
)

func InsertRecorddifferentiation(tz *entities.RecorddifferentiationEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRecorddifferentiation(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertRecorddifferentiation(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Module Name Already Exist."
	}
}

func GetRecorddifferentiationbyrecursive(page *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentionSingle, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	t := []entities.RecorddifferentionSingle{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	values, err1 := GetRecorddifferentiationbyparent1(db, page, t)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
func GetRecorddifferentiationbyparent(page *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentionSingle, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	t := []entities.RecorddifferentionSingle{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetRecorddifferentiationbyparent(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getdiffdetailsbyseq(page *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentionSingle, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getdiffdetailsbyseq(page)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Searchcategory(page *entities.RecorddifferentiationEntity) ([]entities.RecorddifferentionSingle, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	t := []entities.RecorddifferentionSingle{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Searchcategory(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		for i, cat := range values {
			var category string = ""
			parents := strings.Split(cat.Parentcategorynames, "->")
			//if len(parents) > 4 {
			//	category = parents[3] + "->" + parents[4]
			//} else if len(parents) > 3 {
			//	category = parents[3]
			//} else {
			//	category = ""
			//}
			for i := 1; i < len(parents); i++ {
				category = category + "->" + parents[i]
			}
			values[i].Sortedcategorynames = category
		}
	}
	return values, true, err, ""
}
func GetRecorddifferentiationbyparent1(db *sql.DB, page *entities.RecorddifferentiationEntity, output []entities.RecorddifferentionSingle) ([]entities.RecorddifferentionSingle, error) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	//if err != nil {
	//	logger.Log.Println("database connection failure", err)
	//	return output, err
	//}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetRecorddifferentiationbyparent(page)
	if err1 != nil {
		return output, err1
	}
	if len(values) > 0 {
		output = append(output, values[0])
		log.Print(output)
		if values[0].Parentid > 0 {
			page.Id = values[0].Parentid
			return GetRecorddifferentiationbyparent1(db, page, output)
		} else {
			return output, err1
		}
	} else {
		return output, err1
	}
}

func GetAllRecorddifferentiation(page *entities.RecorddifferentiationEntity) (entities.RecorddifferentiationEntities, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	t := entities.RecorddifferentiationEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllRecorddifferentiation(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetRecorddifferentiationCount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteRecorddifferentiation(tz *entities.RecorddifferentiationEntity) (bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteRecorddifferentiation(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateRecorddifferentiation(tz *entities.RecorddifferentiationEntity) (bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	db, err := config.ConnectMySqlDb()
	defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	if strings.Contains(tz.Name, "->") {
		return false, err, "'->' Is Not Allowed In Name"
	}
	// new addition
	oldname, OldParentcategorynames, err := dataAccess.GetRecorddifferentiationoldname(tz.Id)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	split := strings.Split(OldParentcategorynames, "->")
	split[len(split)-1] = tz.Name
	newparentcategoryname := strings.Join(split, "->")
	//new addition
	count, err := dataAccess.CheckDuplicateParentcategoriesforupdate(tz, newparentcategoryname)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total > 0 {
		return false, err, "Data Already Exist"
	}
	err1 := dataAccess.UpdateRecorddifferentiation(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}

	if len(oldname) > 0 {
		err := dataAccess.Updateparentcategorynames(tz.Clientid, tz.Mstorgnhirarchyid, tz.Id, newparentcategoryname, OldParentcategorynames)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
	}
	return true, nil, ""
}

func GetRecorddifferentiationname(page *entities.RecorddifferentiationEntity) (entities.RecorddifferentiationnameEntities, bool, error, string) {
	logger.Log.Println("In side Recorddifferentiationmodel")
	t := entities.RecorddifferentiationnameEntities{}
	db, err := config.ConnectMySqlDb()
	defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetRecorddifferentiationname(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t.Values = values
	return t, true, err, ""
}
