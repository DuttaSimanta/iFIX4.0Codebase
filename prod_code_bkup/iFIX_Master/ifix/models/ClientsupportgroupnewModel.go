package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"strconv"
)

func InsertClientsupportgroupnew(tz *entities.ClientsupportgroupnewEntity) (int64, bool, error, string) {

	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	if tz.Isworkflow == "Y" {
		tx, err := dbcon.Begin()
		if err != nil {
			dbcon.Close()
			logger.Log.Println("Transaction creation error.", err)
			return 0, false, err, "Something Went Wrong"
		}
		count, err := dao.CheckDuplicateClientsupportgroupnewwithtransaction(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}

		if count.Total == 0 {
			lastinsertedID, err := dao.InsertClientsupportgroupnewwithtransaction(tx, tz)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				dbcon.Close()
				return 0, false, err, "Data insertion failure."
			}
			count1, err := dao.CheckDuplicateMstgroupnewwithtransaction(tx, tz)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				dbcon.Close()
				return 0, false, err, "Data insertion failure."
			}
			if count1.Total == 0 {
				_, err := dao.InsertMstgroupnewwithtransaction(tx, tz, lastinsertedID)
				if err != nil {
					logger.Log.Println(err)
					tx.Rollback()
					dbcon.Close()
					return 0, false, err, "Data insertion failure."
				}
				tx.Commit()
				return lastinsertedID, true, err, ""
			} else {
				tx.Rollback()
				dbcon.Close()
				return 0, false, err, "Already data exist.Please verify the data."
			}
		} else {
			tx.Rollback()
			dbcon.Close()
			return 0, false, err, "Already data exist.Please verify the data."
		}
	} else {
		dataAccess := dao.DbConn{DB: dbcon}
		count, err := dataAccess.CheckDuplicateClientsupportgroupnew(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			id, err := dataAccess.InsertClientsupportgroupnew(tz)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
			return id, true, err, ""
		} else {
			return 0, false, nil, "Data Already Exist."
		}
	} //else part end here....

}

func GetAllClientsupportgroupnew(page *entities.ClientsupportgroupnewEntity) (entities.ClientsupportgroupnewEntities, bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupnewmodel")
	t := entities.ClientsupportgroupnewEntities{}
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
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
	values, err1 := dataAccess.GetAllClientsupportgroupnew(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetClientsupportgroupnewCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteClientsupportgroupnew(tz *entities.ClientsupportgroupnewEntity) (bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupnewmodel")
	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	if tz.Isworkflow == "Y" {
		tx, err := dbcon.Begin()
		if err != nil {
			logger.Log.Println("Transaction creation error.", err)
			dbcon.Close()
			return false, err, "Something Went Wrong"
		}
		err1 := dao.DeleteMstgroupnewwithtransaction(tx, tz)
		if err1 != nil {
			logger.Log.Println(err1)
			dbcon.Close()
			tx.Rollback()
			return false, err, "Data deletion failure."
		}

		err2 := dao.DeleteClientsupportgroupnewwithtransaction(tx, tz)
		if err2 != nil {
			logger.Log.Println(err2)
			dbcon.Close()
			tx.Rollback()
			return false, err, "Data deletion failure."
		}
		tx.Commit()
		return true, nil, ""
	} else {
		dataAccess := dao.DbConn{DB: dbcon}
		err1 := dataAccess.DeleteClientsupportgroupnew(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		return true, nil, ""
	}

}

func UpdateClientsupportgroupnew(tz *entities.ClientsupportgroupnewEntity) (bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupnewmodel")
	// dbcon, err := config.ConnectMySqlDb()
	// defer dbcon.Close()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}

	//if tz.Isworkflow == "Y" {
	tx, err := dbcon.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		dbcon.Close()
		return false, err, "Something Went Wrong"
	}
	count, err := dao.CheckDuplicateMstgroupupdatenewwithtransaction(tx, tz)
	if err != nil {
		logger.Log.Println(err)

		tx.Rollback()
		dbcon.Close()
		return false, err, "Data updation failure."
	}
	if count.Total == 0 {
		err := dao.UpdateMstgroupnewwithtransaction(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			dbcon.Close()
			return false, err, "Data updation failure."
		}
		//}  remove to214
		count1, err1 := dao.CheckDuplicateClientsupportgroupnewforupdatewithtransaction(tx, tz)
		if err1 != nil {
			logger.Log.Println(err)
			tx.Rollback()
			dbcon.Close()
			return false, err, "Data updation failure."
		}
		if count1.Total == 0 {
			err := dao.UpdateClientsupportgroupnewwithtransaction(tx, tz)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				dbcon.Close()
				return false, err, "Data updation failure."
			}
			tx.Commit()
			return true, err, ""
		} else {
			tx.Rollback()
			dbcon.Close()
			return false, nil, "Data Already Exist."
		}
	} else {
		tx.Rollback()
		dbcon.Close()
		return false, err, "Already data exist.Please verify the data."
	}
	/*} else {
		dataAccess := dao.DbConn{DB: dbcon}
		count, err := dataAccess.CheckDuplicateClientsupportgroupnewupdate(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		if count.Total == 0 {
			err := dataAccess.UpdateClientsupportgroupnew(tz)
			if err != nil {
				return false, err, "Something Went Wrong"
			}
			return true, err, ""
		} else {
			return false, nil, "Data Already Exist."
		}
	}*/
}

func InsertClientsupportgroupfromto(tz *entities.ClientsupportgroupnewEntity) (int64, bool, error, string) {

	//dbcon, err := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer dbcon.Close()
	var m int64
	var idd int64
	k := 0
	m = 0
	n := 0
	idd = 0
	dataAccess := dao.DbConn{DB: dbcon}

	for i := 0; i < len(tz.FromGroupids); i++ {
		for j := 0; j < len(tz.ToMstorgnhirarchyids); j++ {
			tz.Clientid = tz.FromClientid
			tz.Mstorgnhirarchyid = tz.FromMstorgnhirarchyid
			tz.Supportgroupid = tz.FromGroupids[i]
			values, err1 := dataAccess.Getrowbygrpid(tz)
			if err1 != nil {
				return 0, false, err1, "Something Went Wrong"
			}
			values.Clientid = tz.ToClientid
			values.Mstorgnhirarchyid = tz.ToMstorgnhirarchyids[j]

			//}
			//}

			logger.Log.Println("In side Clientsupportgroupmodel", tz)

			if values.Id == 0 {
				n++
			}
			if values.Isworkflow == "Y" {
				tx, err := dbcon.Begin()
				if err != nil {
					dbcon.Close()
					logger.Log.Println("Transaction creation error.", err)
					return 0, false, err, "Something Went Wrong"
				}
				count, err := dao.CheckDuplicateClientsupportgroupnewwithtransaction(tx, values)
				if err != nil {
					logger.Log.Println(err)
					tx.Rollback()
					dbcon.Close()
					return 0, false, err, "Data insertion failure."
				}

				if count.Total == 0 {
					lastinsertedID, err := dao.InsertClientsupportgroupnewwithtransaction(tx, values)
					m = lastinsertedID
					if err != nil {
						logger.Log.Println(err)
						tx.Rollback()
						dbcon.Close()
						return 0, false, err, "Data insertion failure."
					}
					count1, err := dao.CheckDuplicateMstgroupnewwithtransaction(tx, values)
					if err != nil {
						logger.Log.Println(err)
						tx.Rollback()
						dbcon.Close()
						return 0, false, err, "Data insertion failure."
					}
					if count1.Total == 0 {
						_, err := dao.InsertMstgroupnewwithtransaction(tx, values, lastinsertedID)
						if err != nil {
							logger.Log.Println(err)
							tx.Rollback()
							dbcon.Close()
							return 0, false, err, "Data insertion failure."
						}
						k++
						tx.Commit()
						//return lastinsertedID, true, err, ""
					} /*else {
						tx.Rollback()
						dbcon.Close()
						return 0, false, err, "Already data exist.Please verify the data."
					}*/
				} /* else {
					tx.Rollback()
					dbcon.Close()
					return 0, false, err, "Already data exist.Please verify the data."
				}*/
			} //else {
			if values.Isworkflow == "N" {
				dataAccess := dao.DbConn{DB: dbcon}
				count, err := dataAccess.CheckDuplicateClientsupportgroupnew(values)
				if err != nil {
					return 0, false, err, "Something Went Wrong"
				}
				if count.Total == 0 {
					k++
					id, err := dataAccess.InsertClientsupportgroupnew(values)
					idd = id
					if err != nil {
						return 0, false, err, "Something Went Wrong"
					}
					//return id, true, err, ""
				} /*else {
					return 0, false, nil, "Data Already Exist."

				}*/
			} //else part end here....
		}
	}
	if k == 0 {
		if n == (len(tz.ToMstorgnhirarchyids) * len(tz.FromGroupids)) {
			return 0, false, nil, "Either all FromGroupids or FromMstorgnhirarchyid or both does not Exist."

		} else {
			return 0, false, nil, "Data Already Exist."
		}

	} else if idd > m {
		return idd, true, err, ""

	} else {
		return m, true, err, ""

	}
}

func GetAllClientsupportgroupbyclient(page *entities.ClientsupportgroupnewEntity) (entities.ClientsupportgroupnewEntities, bool, error, string) {
	logger.Log.Println("In side Clientsupportgroupnewmodel")
	t := entities.ClientsupportgroupnewEntities{}
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllClientsupportgroupbyclient(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetAllClientsupportgroupbyclientcount(page)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}
func Getsupportgroupbyorg(page *entities.ClientsupportgroupnewEntity) ([]entities.GetsupportgroupbyorgEntity, bool, error, string) {
	//logger.Log.Println("In side Clientsupportgroupnewmodel")
	t := []entities.GetsupportgroupbyorgEntity{}
	// db, err := config.ConnectMySqlDb()
	// defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	var ids string
	for j, i := range page.Mstorgnhirarchyids {
		if j > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(i))
	}
	//logger.Log.Println(ids)
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Getsupportgroupbyorg(page, ids)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	return values, true, err, ""
}
