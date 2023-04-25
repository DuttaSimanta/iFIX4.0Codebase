package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertClientmoduleurlroleuser(tz *entities.ClientmoduleurlroleuserEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Clientmoduleurlroleusermodel")
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
		logger.Log.Println("database transaction connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess1 := dao.TxConn{TX: tx}
	j := 0
	var id int64
	for i := 0; i < len(tz.Menuids); i++ {
		count, err := dataAccess.CheckDuplicateClientmoduleurlroleuser(tz, i)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				logger.Log.Print("Clientmoduleurlroleusermodel  Statement Rollback error", err)
				return 0, false, err, ""
			}
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err = dataAccess1.InsertClientmoduleurlroleuser(tz, i)
			if err != nil {
				err = tx.Rollback()
				if err != nil {
					logger.Log.Print("Clientmoduleurlroleusermodel  Statement Rollback error", err)
					return 0, false, err, ""
				}
				return 0, false, err, "Something Went Wrong"
			}
		} else {
			j++
		}
	}
	if j < len(tz.Menuids) {
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("Clientmoduleurlroleusermodel  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("Clientmoduleurlroleusermodel  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}
}
func GetAllClientmoduleurlroleuser(page *entities.ClientmoduleurlroleuserEntity) (entities.ClientmoduleurlroleuserEntities, bool, error, string) {
	logger.Log.Println("In side Clientmoduleurlroleusermodel")
	t := entities.ClientmoduleurlroleuserEntities{}
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
	values, err1 := dataAccess.GetAllClientmoduleurlroleuser(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetClientmoduleurlroleuserCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteClientmoduleurlroleuser(tz *entities.ClientmoduleurlroleuserEntity) (bool, error, string) {
	logger.Log.Println("In side Clientmoduleurlroleusermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteClientmoduleurlroleuser(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateClientmoduleurlroleuser(tz *entities.ClientmoduleurlroleuserEntity) (bool, error, string) {
	logger.Log.Println("In side Clientmoduleurlroleusermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateClientmoduleurlroleuserforupdate(tz)
	if err != nil {
		return false, err, "Something Went Wrong"

	}
	if count.Total == 0 {

		err1 := dataAccess.UpdateClientmoduleurlroleuser(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		return true, nil, ""
	} else {
		return false, nil, "Data Already Exist."

	}
}
