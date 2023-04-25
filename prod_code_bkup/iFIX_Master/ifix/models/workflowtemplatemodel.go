package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
	"strconv"
)

func Getprocesstemplatedetails(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowResponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getprocesstemplatedetails(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Getprocesstemplate(tz *entities.Workflowentity) ([]entities.WorkflowResponseEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.WorkflowResponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getprocesstemplate(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
func Createprocesstemplatetransition(tw *entities.Workflowentity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}

	dataAccess := dao.DbConn{DB: dbcon}

	id, err := dao.Createtemplatetransition(tw, tx)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	if len(tw.Transitionids) > 0 {
		tw.Transitionid = tw.Transitionids[0]
		details, err2 := dataAccess.Getalltemplatetransitionstatedetails(tw)
		if err2 != nil {
			return 0, false, err2, "Something Went Wrong"
		}
		if len(details) > 0 {
			log.Print("len(details):", len(details))
			for _, val := range details {
				tw.Mstgroupid = val.Mstgroupid
				tw.Mstuserid = val.Mstuserid
				tw.Transitionid = id
				_, err := dao.Inserttemplatetransitiondetails(tw, tx)
				if err != nil {
					tx.Rollback()
					return 0, false, err, "Something Went Wrong"
				}
			}
			err = tx.Commit()
			if err != nil {
				log.Print("Createprocesstemplatetransition  Statement Commit error", err)
				logger.Log.Print("Createprocesstemplatetransition  Statement Commit error", err)
				return 0, false, err, ""
			}

			return id, true, nil, ""
		} else {
			err = tx.Commit()
			if err != nil {
				log.Print("Createprocesstemplatetransition  Statement Commit error", err)
				logger.Log.Print("Createprocesstemplatetransition  Statement Commit error", err)
				return 0, false, err, ""
			}
			return id, true, nil, ""
		}
	} else {
		err = tx.Commit()
		if err != nil {
			log.Print("Createprocesstemplatetransition  Statement Commit error", err)
			logger.Log.Print("Createprocesstemplatetransition  Statement Commit error", err)
			return 0, false, err, ""
		}
		return id, true, nil, ""
	}

}
func Insertprocesstemplate(tz *entities.Workflowentity) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	savemst := dao.DbConn{DB: db}
	values, err1 := savemst.Getprocesstemplatedetails(tz)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	if len(values) == 0 {
		id, err := savemst.Insertprocesstemplate(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		tz.Id = values[0].Id
		err = savemst.Updateprocesstemplatedetails(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return 0, true, err, ""
	}
}
func Deletetemplatetransitionstate(tw *entities.Workflowentity) (int64, bool, error, string) {
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
	var ids string = ""
	for i, transition := range tw.Transitionids {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(transition))
	}
	err = dao.Deletetemplatetransitiondetails(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	/*err = dao.Deleteactivitydetails(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}*/
	err = dao.Deletetemplatetransition(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		log.Print("Deletetransitionstate  Statement Commit error", err)
		logger.Log.Print("Deletetransitionstate  Statement Commit error", err)
		return 0, false, err, ""
	}
	return 0, true, nil, ""
}
func Upserttemplatetransitiondetails(tw *entities.Workflowentity) (int64, bool, error, string) {
	log.Println("In side model")
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
	var ids string = ""
	for i, transition := range tw.Transitionids {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(transition))
	}
	err = dao.Deletetemplatetransitiondetails(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}
	/*err = dao.Deleteactivitydetails(tw, tx, ids)
	if err != nil {
		tx.Rollback()
		return 0, false, err, "Something Went Wrong"
	}*/
	for _, tid := range tw.Transitionids {
		tw.Transitionid = tid
		for _, user := range tw.Users {
			tw.Mstuserid = user.Mstuserid
			tw.Mstgroupid = user.Mstgroupid
			_, err := dao.Inserttemplatetransitiondetails(tw, tx)
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}
		}
		/*for _, activity := range tw.Activities {
			tw.Activity = activity
			_, err := dao.Insertstateactivity(tw, tx)
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}
		}*/
	}
	err = tx.Commit()
	if err != nil {
		log.Print("Upserttransitiondetails  Statement Commit error", err)
		logger.Log.Print("Upserttransitiondetails  Statement Commit error", err)
		return 0, false, err, ""
	}
	return 0, true, nil, ""
}
func Gettemplatetransitionstatedetails(tz *entities.Workflowentity) (entities.WorkflowStateResponseEntity, bool, error, string) {
	log.Println("In side model")
	t := entities.WorkflowStateResponseEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	if len(tz.Transitionids) > 0 {
		tz.Transitionid = tz.Transitionids[0]
		details, err2 := dataAccess.Gettemplatetransitionstatedetails(tz)
		if err2 != nil {
			return t, false, err2, "Something Went Wrong"
		}
		//activities, err3 := dataAccess.Getactivitybytransition(tz)
		//if err3 != nil {
		//	return t, false, err3, "Something Went Wrong"
		//}
		//for _, activity := range activities {
		//	t.Activityids = append(t.Activityids, activity.Id)
		//}
		if len(details) > 0 && details[0].Mstuserid > -1 {
			groupdetails, err2 := dataAccess.Gettemplatetransitiongroup(tz)
			if err2 != nil {
				return t, false, err2, "Something Went Wrong"
			}
			t.Groups = groupdetails
			return t, true, nil, ""
		} else {
			t.Groups = details
			return t, true, nil, ""
		}
	} else {
		return t, false, nil, "Problem with this state.Please remove and recreate this state"
	}
}
