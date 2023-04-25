package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertBanner(tz *entities.BannerEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Bannermodel")
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
	for i := 0; i < len(tz.Groupid); i++ {
		count, err := dataAccess.CheckDuplicateBanner(tz, i)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		if count.Total > 0 {
			j++
			// continue;
		} else {
			id, err = dataAccess1.InsertBanner(tz, i)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
		}

	}

	if j < len(tz.Groupid) {
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

func GetAllBanner(page *entities.BannerEntity) (entities.BannerEntities, bool, error, string) {
	logger.Log.Println("In side Bannermodel")

	t := entities.BannerEntities{}
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

	tz := entities.UtilityEntity{}
	tz.Clientid = page.Clientid
	tz.Mstorgnhirarchyid = page.Mstorgnhirarchyid
	err2, timediff := dataAccess.Gettimediff(&tz)
	if err2 != nil {
		return t, false, err2, "Something went wrong"
	}

	values, err1 := dataAccess.GetAllBanner(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	/* convert */
	for i := 0; i < len(values); i++ {
		values[i].ActualStarttime = dao.Convertdate(values[i].Starttime, timediff[0].Timediff, timediff[0].Timeformat)
		values[i].ActualEndtime = dao.Convertdate(values[i].Endtime, timediff[0].Timediff, timediff[0].Timeformat)
	}

	if page.Offset == 0 {
		total, err1 := dataAccess.GetBannerCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteBanner(tz *entities.BannerEntity) (bool, error, string) {
	logger.Log.Println("In side Bannermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteBanner(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateBanner(tz *entities.BannerEntity) (bool, error, string) {
	logger.Log.Println("In side Bannermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	//dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateBanner(tz, 0)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateBanner(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}

		return true, err, ""
	} else {

		return false, nil, "Data Already Exist."
	}
}

func GetAllMessage(page *entities.BannerEntity) ([]entities.BannerMessageEntity, bool, error, string) {
	logger.Log.Println("In side Bannermodel")

	t := []entities.BannerMessageEntity{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}

	values, err1 := dataAccess.GetAllMessage(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""

}

func UpdateBannerSequence(tz *entities.BannerEntity) (bool, error, string) {
	logger.Log.Println("In side Bannermodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	// count,err :=dataAccess.CheckDuplicateBanner(tz)
	// if err != nil {
	//    return false, err, "Something Went Wrong"
	// }
	// if count.Total == 0 {
	err = dataAccess.UpdateBannerSequence(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	return true, err, ""
	// }else{
	//     return false, nil, "Data Already Exist."
	// }
}
