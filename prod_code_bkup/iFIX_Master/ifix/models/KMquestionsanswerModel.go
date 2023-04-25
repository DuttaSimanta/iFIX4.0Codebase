package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertQuestionanswer(tz *entities.KMquestionsanswerEntity) (int64, bool, error, string) {

	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	tx, err := dbcon.Begin()
	if err != nil {
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}
	count, err := dao.CheckDuplicateQuestions(tx, tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		// dbcon.Close()
		return 0, false, err, "Data insertion failure."
	}
	if count.Total == 0 {
		lastinsertedQID, err := dao.InsertQuestions(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}
		lastinsertedAID, err := dao.InsertAnswers(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}
		lastinsertedMID, err := dao.InsertRecordmap(tx, tz, lastinsertedQID, lastinsertedAID)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			// dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}
		tx.Commit()
		return lastinsertedMID, true, err, ""

	} else {
		tx.Rollback()
		// dbcon.Close()
		return 0, false, err, "Already data exist.Please verify the data."
	}

}

func DeleteQuetionanswer(tz *entities.KMquestionsanswerEntity) (bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupmodel")
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	tx, err := dbcon.Begin()
	if err != nil {
		// dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}
	err3 := dao.DeleteQuestions(tx, tz)
	if err3 != nil {
		logger.Log.Println(err3)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Data deletion failure."
	}

	err1 := dao.DeleteAnswers(tx, tz)
	if err1 != nil {
		logger.Log.Println(err1)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Data deletion failure."
	}
	err2 := dao.DeleteRecordmap(tx, tz)
	if err2 != nil {
		logger.Log.Println(err2)
		tx.Rollback()
		// dbcon.Close()
		return false, err, "Data deletion failure."
	}
	tx.Commit()
	return true, nil, ""

}

func GetAllQuestionAnswers(page *entities.KMquestionsanswerEntity) (entities.KMquestionsanswerEntities, bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupmodel")
	t := entities.KMquestionsanswerEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllQuestionanswer(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetQuestionanswerscount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
