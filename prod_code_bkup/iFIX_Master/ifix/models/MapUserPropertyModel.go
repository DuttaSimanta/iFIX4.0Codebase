package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertUserRoleProperty(tz *entities.MapUserRolePropertyEntity) (int64, bool, error, string) {
	logger.Log.Println("In side InsertUserRoleProperty")
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
	for i := 0; i < len(tz.Roleid); i++ {
		count, err := dataAccess.CheckDuplicateRoleProperty(tz, i)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total > 0 {
			j++
			// continue;
		} else {
			id, err = dataAccess1.InsertRoleProperty(tz, i)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
		}

	}

	if j < len(tz.Roleid) {
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Banner  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Commit()
		if err != nil {
			// log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("Banner  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}

}

func GetAllUserRoleProperty(page *entities.MapUserRolePropertyEntity) (entities.MapUserRolePropertyEntities, bool, error, string) {
	logger.Log.Println("In side GetAllUserRoleProperty")

	t := entities.MapUserRolePropertyEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(page.Clientid, page.Mstorgnhirarchyid)
	logger.Log.Println(orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	values, err1 := dataAccess.GetAllUserRoleProperty(page, orgntype)
	// logger.Log.Println(values)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	if page.Offset == 0 {
		total, err1 := dataAccess.GetUserRolePropertyCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func GetUserPropertyName(page *entities.MapUserRolePropertyEntity) ([]entities.GetUserPropertyNameEntity, bool, error, string) {
	logger.Log.Println("In side GetUserPropertyName")
	t := []entities.GetUserPropertyNameEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetUserPropertyName(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t = values
	return t, true, err, ""
}

func UpdateUserPropertyName(tz *entities.MapUserRolePropertyEntity) (bool, error, string) {
	logger.Log.Println("In side UpdateUserPropertyName")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	//dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateRoleProperty(tz, 0)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateUserPropertyName(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}

		return true, err, ""
	} else {

		return false, nil, "Data Already Exist."
	}
}

func DeleteUserPropertyName(tz *entities.MapUserRolePropertyEntity) (bool, error, string) {
	logger.Log.Println("In side DeleteUserPropertyName")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteUserPropertyName(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}
