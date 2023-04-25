package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertClientdayofweek(tz *entities.ClientdayofweekEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Clientdayofweekmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("database transaction failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess1 := dao.TxConn{TX: tx}

	details := tz.Details
	var countday int
	for k := 0; k < len(details); k++ {
		count, err := dataAccess.CheckDuplicateClientdayofweek(tz, &details[k])
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			_, err := dataAccess1.InsertClientdayofweek(tz, &details[k])
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}

		} else {
			countday++
		}
	}
	if countday < len(details) {
		tx.Commit()
		return 0, true, err, ""
	} else {
		return 0, true, err, "Data Already Exist"
	}
}

func GetAllClientdayofweek(page *entities.ClientdayofweekEntity) (entities.ClientdayofweekEntities, bool, error, string) {
	logger.Log.Println("In side Clientdayofweekmodel")
	t := entities.ClientdayofweekEntities{}
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
	values, err1 := dataAccess.GetAllClientdayofweek(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetClientdayofweekCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteClientdayofweek(tz *entities.ClientdayofweekEntity) (bool, error, string) {
	logger.Log.Println("In side Clientdayofweekmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteClientdayofweek(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

// func UpdateClientdayofweek(tz *entities.ClientdayofweekEntity) (bool, error, string) {
// 	logger.Log.Println("In side Clientdayofweekmodel")
// 	db, err := config.ConnectMySqlDb()
// 	defer db.Close()
// 	if err != nil {
// 		logger.Log.Println("database connection failure", err)
// 		return false, err, "Something Went Wrong"
// 	}
// 	dataAccess := dao.DbConn{DB: db}
// 	err1 := dataAccess.UpdateClientdayofweek(tz)
// 	if err1 != nil {
// 		return false, err1, "Something Went Wrong"
// 	}
// 	return true, nil, ""
// }

func GetClientwisedayofweek(page *entities.ClientdayofweekEntity) ([]entities.ClientdayofweekresponseEntity, bool, error, string) {
	logger.Log.Println("In side Clientdayofweekmodel")
	t := []entities.ClientdayofweekresponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetClientwisedayofweek(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}
