package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"strconv"
	"strings"
)

//Inserthighkeyvalue method is used to insert data differentiation & recordtype tables
func Inserthighkeyvalue(tz *entities.RecorddifferentiationhigherkeyEntity) (int64, bool, error, string) {
	logger.Log.Println("Inside Recorddifferentiationhihgerkey Models")
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return 0, false, err, "Something Went Wrong"
	}
	if strings.Contains(tz.Name, "->") {
		return 0, false, err, "'->' Is Not Allowed"

	}
	tx, err := dbcon.Begin()
	if err != nil {
		//dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return 0, false, err, "Something Went Wrong"
	}

	differentiationtotal, err := dao.CheckDuplicateRecorddifferentiationkey(tx, tz)
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		//dbcon.Close()
		return 0, false, err, "Already data exist.Please verify the data."
	}
	if differentiationtotal == 0 { //mstrecorddifferentiation count value check here
		// mstrecordtypetotal, err := dao.CheckDuplicateMstrecordtypedata(tx, tz)
		// if err != nil {
		// 	logger.Log.Println(err)
		// 	tx.Rollback()
		// 	dbcon.Close()
		// 	return 0, false, err, "Already data exist.Please verify the data."
		// }
		// if mstrecordtypetotal == 0 { //mstrecordtype count value check here

		lastinsertedID, err := dao.InsertRecorddifferentiation(tx, tz)
		if err != nil {
			logger.Log.Println(err)
			tx.Rollback()
			//dbcon.Close()
			return 0, false, err, "Data insertion failure."
		}
		if lastinsertedID > 0 {
			lastinsertedmapID, err := dao.InsertMstRecordtype(tx, tz, lastinsertedID)
			if err != nil {
				logger.Log.Println(err)
				tx.Rollback()
				//dbcon.Close()
				return 0, false, err, "Data insertion failure."
			}
			if lastinsertedmapID > 0 {
				err = tx.Commit()
				if err != nil {
					logger.Log.Println(err)
					tx.Rollback()
					//dbcon.Close()
					return 0, false, err, "Data insertion failure."
				}
				//add new code here
				dataAccess := dao.DbConn{DB: dbcon}
				logger.Log.Println("Parent details --------- tz.Parentid ------------------- tz.Parentid -------------------------------->", tz.Parentid)
				if tz.Parentid == 0 {
					err := dataAccess.UpdateParentPath(strconv.FormatInt(tz.Parentid, 10), tz.Name, lastinsertedID)
					if err != nil {
						logger.Log.Println(err)
						//dbcon.Close()
						return 0, false, err, "Data insertion failure."
					}
				} else {
					parentIDs, parentnames, err := dataAccess.GetParentdetails(tz)
					logger.Log.Println("Parent details ---------------------------11--------------------------------->", parentIDs, parentnames)
					if err != nil {
						logger.Log.Println(err)
						//dbcon.Close()
						return 0, false, err, "Data insertion failure."
					}
					if parentIDs == "0" {
						var parent = parentnames + "->" + tz.Name
						var substr = "->"
						var pos = strings.LastIndex(parent, substr)
						var parentcats = parent[0:pos]
						logger.Log.Println("Parent details ------------------------22------------------------------------>", parent)
						logger.Log.Println("Parent details ------------------------33------------------------------------>", parentcats)
						err := dataAccess.UpdateParentPath(strconv.FormatInt(tz.Parentid, 10), parent, lastinsertedID)
						if err != nil {
							logger.Log.Println(err)
							//dbcon.Close()
							return 0, false, err, "Data insertion failure."
						}
					} else {
						// var parent = parentnames + "->" + tz.Name
						// var substr = "->"
						// var pos = strings.LastIndex(parent, substr)
						// var parentcats = parent[0:pos]
						//pnames, err := dataAccess.GetParentnamesbyID(lastinsertedID)
						// if err != nil {
						// 	logger.Log.Println(err)
						// 	dbcon.Close()
						// 	return 0, false, err, "Data insertion failure."
						// }
						logger.Log.Println("Parent details ------------------------44------------------------------------>", tz.Name)
						err = dataAccess.UpdateParentPath(parentIDs+"->"+strconv.FormatInt(tz.Parentid, 10), parentnames+"->"+tz.Name, lastinsertedID)
						if err != nil {
							logger.Log.Println(err)
							//dbcon.Close()
							return 0, false, err, "Data insertion failure."
						}
					}
				}

				//End new code here
				//dbcon.Close()
				return lastinsertedID, true, err, ""
			}
		} // lastinsertedID if condition
		//} //mstrecordtypetotal if condition
	} // total value check here.
	return 0, false, err, "Data insertion failure."
}

//Deletehigherkey method is used to delete data differentiation & recordtype tables
func Deletehigherkey(tz *entities.RecorddifferentiationhigherkeyEntity) (bool, error, string) {
	logger.Log.Println("In side Workdifferentiationmodel")
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return false, err, "Something Went Wrong"
	}

	tx, err := dbcon.Begin()
	if err != nil {
		tx.Rollback()
		//dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}

	err1 := dao.DeleteRecorddifferentiation(tx, tz)
	if err1 != nil {
		tx.Rollback()
		//dbcon.Close()
		logger.Log.Println("recorddifferentiation table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err2 := dao.DeleteMstrecordtype(tx, tz)
	if err2 != nil {
		tx.Rollback()
		//dbcon.Close()
		logger.Log.Println("mstrecordtype table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		//dbcon.Close()
		return false, err, "Data deletion failure."
	}
	//dbcon.Close()
	return true, nil, ""
}

func GetAllRecorddifferentiationHighkey(page *entities.RecorddifferentiationhigherkeyEntity) (entities.RecorddifferentiationhigherkeyEntities, bool, error, string) {
	logger.Log.Println("In side Catalogmodel")
	t := entities.RecorddifferentiationhigherkeyEntities{}
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
	values, err1 := dataAccess.GetAllRecorddifferentiationHighkey(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetRecorddifferentiationhighkeyCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

//Updatehigherkey method is used to delete data differentiation & recordtype tables
func Updatehigherkey(tz *entities.RecorddifferentiationhigherkeyEntity) (bool, error, string) {
	logger.Log.Println("In side Workdifferentiationmodel")
	lock.Lock()
	defer lock.Unlock()
	dbcon, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("Database connection failure.", err)
		return false, err, "Something Went Wrong"
	}

	tx, err := dbcon.Begin()
	if err != nil {
		tx.Rollback()
		//dbcon.Close()
		logger.Log.Println("Transaction creation error.", err)
		return false, err, "Something Went Wrong"
	}

	err1 := dao.UpdateRecorddifferentiationData(tx, tz)
	if err1 != nil {
		tx.Rollback()
		//dbcon.Close()
		logger.Log.Println("recorddifferentiation table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err2 := dao.UpdateMstrecordtypeData(tx, tz)
	if err2 != nil {
		tx.Rollback()
		//dbcon.Close()
		logger.Log.Println("mstrecordtype table data delete error.", err1)
		return false, err1, "Something Went Wrong"
	}
	err = tx.Commit()
	if err != nil {
		logger.Log.Println(err)
		tx.Rollback()
		//dbcon.Close()
		return false, err, "Data deletion failure."
	}
	//dbcon.Close()
	return true, nil, ""
}
