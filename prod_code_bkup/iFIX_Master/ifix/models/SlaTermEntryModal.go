package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func AddSlaTermEntry(tz *entities.SlaTermEntryEntity) (int64, bool, error, string) {
	//logger.Log.Println("In side SlaTermEntry model", tz)
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
	for i := 0; i < len(tz.MeterNames); i++ {
		values, err := dataAccess.GetRow(tz, i)
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Something Went Wrong"
		}
		// tz.Clientid = tz.ToClientid
		// tz.Mstorgnhirarchyid = tz.ToMstorgnhirarchyid
		tz.Seqno = values.Seqno
		count, err := dataAccess.CheckDuplicateSlaTermEntry(tz, i)
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total > 0 {
			j++
			// continue;
		} else {
			id, err = dataAccess1.AddSlaTermEntry(tz, i)
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}
		}

	}

	if j < len(tz.MeterNames) {
		err = tx.Commit()
		if err != nil {
			//log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("SlaTermEntry  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, err, ""
	} else {
		err = tx.Commit()
		if err != nil {
			// log.Print("MoveWorkflow  Statement Commit error", err)
			logger.Log.Print("SlaTermEntry  Statement Commit error", err)
			return 0, false, err, ""
		}
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllSlaTermEntry(page *entities.SlaTermEntryEntity) (entities.SlaTermEntryEntities, bool, error, string) {
	logger.Log.Println("In side SlaTermEntry model")
	t := entities.SlaTermEntryEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllSlaTermEntry(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetSlaTermEntryCount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteSlaTermEntry(tz *entities.SlaTermEntryEntity) (bool, error, string) {
	logger.Log.Println("In side SlaTermEntity model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteSlaTermEntry(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}
