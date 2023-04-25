package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/utility"
	"log"
)

func Insertuserticket(tz *entities.TicketUserEntity) (int64, bool, error, string) {
	logger.Log.Println("In side UidGenmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	users, err := dataAccess.CheckDuplicateticketuser(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(users) == 0 {
		delerr := dataAccess.Deleteuserticket(tz)
		if delerr != nil {
			return 0, false, err, "Something Went Wrong"
		}
		id, err := dataAccess.Insertuserticket(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		msg := "Ticket is already opened by " + users[0].Username +"."
		return 0, false, nil, msg
	}
}
func Deleteuserticket(tz *entities.TicketUserEntity) (int64, bool, error, string) {
	logger.Log.Println("In side UidGenmodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	delerr := dataAccess.Deleteuserticket(tz)
	if delerr != nil {
		return 0, false, err, "Something Went Wrong"
	}
	return 0, true, err, ""
}

func Getworklowutilitylist(tz *entities.WorkflowUtilityEntity) ([]entities.WorkflowSingleEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowSingleEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getworklowutilitylist(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getutilitydatabyfield(tz *entities.WorkflowUtilityEntity) ([]entities.WorkflowSingleEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowSingleEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getutilitydatabyfield(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getprocessbydiffid(tz *entities.WorkflowUtilityEntity) ([]entities.WorkflowSingleEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowSingleEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getprocessbydiffid(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getstatebyseq(tz *entities.WorkflowUtilityEntity) ([]entities.StateStatusEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.StateStatusEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getstatebyseq(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}

//Searchworkflowuser for implements business logic
func Searchworkflowuser(tz *entities.WorkflowUtilityEntity) ([]entities.MstUserSearchEntity, bool, error, string) {
	logger.Log.Println("In side model")
	t := []entities.MstUserSearchEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Searchworkflowuser(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getstatebyprocesstemplate(tz *entities.WorkflowUtilityEntity) (entities.StateCategory, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.StateCategory{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getstatebyprocesstemplate(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	var arr = [] int64{}
	stypearr := []entities.MapStateTypeEntity{}
	for _, val := range values {
		isExist, pos := utility.ItemExists(arr, val.Statetypeid)
		if isExist == false {
			stype := entities.MapStateTypeEntity{}
			stype.Statetypeid = val.Statetypeid
			stype.Statetypename = val.Statetypename
			state := entities.MapStateEntity{}
			state.Stateid = val.Stateid
			state.Statename = val.Statename
			stype.States = append(stype.States, state)
			stypearr = append(stypearr, stype)
			arr = append(arr, val.Statetypeid)
		} else {
			state := entities.MapStateEntity{}
			state.Stateid = val.Stateid
			state.Statename = val.Statename
			stypearr[pos].States = append(stypearr[pos].States, state)
		}
	}
	catarr := entities.StateCategory{}
	catarr.States = stypearr
	return catarr, true, err, ""
}
func Getstatebyprocess(tz *entities.WorkflowUtilityEntity) (entities.StateCategory, bool, error, string) {
	logger.Log.Println("In side model")
	t := entities.StateCategory{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getstatebyprocess(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	var arr = [] int64{}
	stypearr := []entities.MapStateTypeEntity{}
	for _, val := range values {
		isExist, pos := utility.ItemExists(arr, val.Statetypeid)
		if isExist == false {
			stype := entities.MapStateTypeEntity{}
			stype.Statetypeid = val.Statetypeid
			stype.Statetypename = val.Statetypename
			state := entities.MapStateEntity{}
			state.Stateid = val.Stateid
			state.Statename = val.Statename
			stype.States = append(stype.States, state)
			stypearr = append(stypearr, stype)
			arr = append(arr, val.Statetypeid)
		} else {
			state := entities.MapStateEntity{}
			state.Stateid = val.Stateid
			state.Statename = val.Statename
			stypearr[pos].States = append(stypearr[pos].States, state)
		}
	}
	catarr := entities.StateCategory{}
	cat, err1 := dataAccess.Getdiffidbyprocess(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	catarr.Recorddifftypeid = cat[0].Recorddifftypeid
	catarr.Recorddiffid = cat[0].Recorddiffid
	catarr.States = stypearr
	return catarr, true, err, ""
}
func Deleteprocessdetails(tw *entities.WorkflowUtilityEntity) (int64, bool, error, string) {
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		//dbcon.Close()
		logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}
	err = dao.Deleteprocessdetails(tw, tx)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = dao.Deletetprocesstransition(tw, tx)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = dao.Deleteprocessgroupdetails(tw, tx)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		log.Print("Deleteprocessdetails  Statement Commit error", err)
		logger.Log.Print("Deleteprocessdetails  Statement Commit error", err)
		return 0, false, err, ""
	}
	return 0, true, nil, ""
}
func Deleteprocesstemplatedetails(tw *entities.WorkflowUtilityEntity) (int64, bool, error, string) {
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		//dbcon.Close()
		logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}
	err = dao.Deleteprocesstemplatedetails(tw, tx)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = dao.Deletetprocesstemplatetransition(tw, tx)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = dao.Deleteprocesstemplategroupdetails(tw, tx)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		log.Print("Deleteprocessdetails  Statement Commit error", err)
		logger.Log.Print("Deleteprocessdetails  Statement Commit error", err)
		return 0, false, err, ""
	}
	return 0, true, nil, ""
}
