package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

//GetAllClientsnames for implements business logic
func GetAllClientsnames() ([]entities.AllMstClientEntity, bool, error, string) {
	logger.Log.Println("In side model GetAllClientsnames")
	t := []entities.AllMstClientEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllClientsnames()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

//AddClients for implements business logic
func AddClients(tz *entities.MstClientEntity) (int64, bool, error, string) {
	logger.Log.Println("In side model")
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	tx, err := dbcon.Begin()
	if err != nil {
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}

	count, err := dao.CheckDuplicateCientwithTX(tx, tz)
	if err != nil {
		tx.Rollback()
		// dbcon.Close()
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total > 0 {
		tx.Rollback()
		// dbcon.Close()
		return 0, false, nil, "Client Already Exist."
	}
	id, err := dao.InsertClientDatawithTX(tx, tz)
	if err != nil {
		tx.Rollback()
		// dbcon.Close()
		return 0, false, err, "Something Went Wrong"
	}
	if id > 0 {
		count, err := dao.CheckDuplicateOrganizationwithTX(tx, id, 2, tz.Name, tz.Code, 1, 1, 2)
		if err != nil {
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total > 0 {
			tx.Rollback()
			// dbcon.Close()
			return 0, false, nil, "Data Already Exist."
		}

		_, err1 := dao.InsertOrganizationwithTX(tx, id, 2, tz.Name, tz.Code, 1, 1, 2, 1)
		if err1 != nil {
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Something Went Wrong"
		}
		err = tx.Commit()
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}
		// dbcon.Close()
		return id, true, err, ""
	} else {
		tx.Rollback()
		// dbcon.Close()
		return 0, false, err, "Client creation failed"
	}
}

//GetAllClients for implements business logic
func GetAllClients(tz *entities.MstClientEntity) (entities.MstClientEntities, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.MstClientEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllClients(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if tz.Offset == 0 {
		total, err1 := dataAccess.GetClientCount()
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//DeleteClients for implements business logic
func DeleteClients(tz *entities.MstClientEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteClientData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

//UpdateClients for implements business logic
func UpdateClients(tz *entities.MstClientEntity) (bool, error, string) {
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.UpdateClientData(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}
