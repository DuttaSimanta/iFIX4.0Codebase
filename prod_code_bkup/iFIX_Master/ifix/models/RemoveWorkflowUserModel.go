package models

import (
	"database/sql"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	Logger "iFIX/ifix/logger"
	"log"
	"time"
)

func RemoveWorkflowUsersWithapi(tw *entities.MstClientUserEntity) ( bool, error) {
	log.Println("In side model")

	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		//dbcon.Close()
		Logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return false, err
	}
	success, err:= RemoveWorkflowUsers(tw, dbcon)
	return success, err
}
func RemoveWorkflowUsers(tz *entities.MstClientUserEntity, db *sql.DB) (bool, error) {
	log.Println("RemoveWorkflowUsers:",tz.ID)
	dataAccess := dao.DbConn{DB: db}

	/*val := entities.LoginEntityReq{}
	val.ID = tz.ID
	groups, err := dataAccess.Getgroupbyuserid(&val)
	if err != nil {
		return  false, err
	}
	if len(groups)>0{
		isResolver:=false
		if len(groups)>1{
			isResolver=true
		}else{
			if groups[0].Levelid !=1{
				isResolver=true
			}
		}*/
	tickets, err1 := dataAccess.FetchAssignedTicketByUser(tz)
	if err1 != nil {
		return false, err1
	}
	//log.Println(len("Number of deleted tickets:",tickets))
	Logger.Log.Println("Number of deleted tickets:",len(tickets))
	//log.Println(tickets)
	if len(tickets) > 0 {
		for _, ticket := range tickets {
			history, err2 := dataAccess.Gethistorydetails(ticket.Mstrequestid)
			if err2 != nil {
				return false, err2
			}
			latestTime := time.Now().Unix()
			if len(history) > 0 {
				tx, err := db.Begin()
				if err != nil {
					Logger.Log.Println("Transaction creation error.", err)
					return false, err
				}
				isResolver := false
				if history[0].Levelid > 1 {
					isResolver = true
				}
				//log.Println(latestTime,isResolver,ticket.Mstrequestid,history[0].Mstrequestid)
				var mstuserid int64 = 0
				var mstgroupid int64 = 0
				if isResolver {
					history[0].Mstuserid = 0
					err3 := dao.InsertProcessHistory(&history[0], tx, latestTime, "N")
					if err3 != nil {
						tx.Rollback()
						return false, err3
					}
					mstgroupid = history[0].Mstgroupid
				} else {
					if len(history) > 1 {
						history[1].Mstuserid = 0
						mstgroupid = history[1].Mstgroupid
						err4 := dao.InsertProcessHistory(&history[1], tx, latestTime, "N")
						if err4 != nil {
							tx.Rollback()
							return false, err4
						}
					} else {
						Logger.Log.Println("No last Assigned group for ", tz.ID, ticket.Mstrequestid)
						log.Println("No last Assigned group for  ", tz.ID, ticket.Mstrequestid)
					}
				}
				err4 := dao.UpdateAssignedUser(mstgroupid,mstuserid, tx, ticket.Mstrequestid)
				if err4 != nil {
					tx.Rollback()
					return false, err4
				}
				err = tx.Commit()
				if err != nil {
					Logger.Log.Println(err)
					tx.Rollback()
					return false, err
				}
			} else {
				Logger.Log.Println("No Workflow ticket for ", tz.ID, ticket.Mstrequestid)
				log.Println("No Workflow ticket for  ", tz.ID, ticket.Mstrequestid)
			}
		}
		return true, nil
	} else {
		Logger.Log.Println("No Assigned ticket for ", tz.ID)
		log.Println("No Assigned ticket for ", tz.ID)
		return false, nil
	}
	
}
