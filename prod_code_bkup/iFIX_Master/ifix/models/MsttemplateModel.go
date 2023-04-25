package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func InsertNotificationTemplate(tz *entities.MstNotificationTemplateEntity) ([]int64, bool, error, string) {
	logger.Log.Println("In side Msttemplatemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	var response []int64
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return response, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	var errCounter int64
	for _, v := range tz.WorkingCategories {
		tz.WorkingCategoryID = v
		count, err := dataAccess.CheckDuplicateNotificationTemplate(tz)
		if err != nil {
			return response, false, err, "Something Went Wrong"
		}
		errCounter = errCounter + count
	}
	if errCounter == 0 {
		for _, v := range tz.WorkingCategories {
			tz.WorkingCategoryID = v
			if tz.Isconverted != 1 {
				tz.Isconverted = 2
			}
			id, err := dataAccess.InsertNotificationTemplate(tz)
			response = append(response, id)
			if err != nil {
				return response, false, err, "Something Went Wrong"
			}
			for i, _ := range tz.Recipients {
				tz.Recipients[i].NotificationTemplateID = id
				_, err := dataAccess.InsertNotificationRecipient(&tz.Recipients[i])
				if err != nil {
					return response, false, err, "Something Went Wrong"
				}
			}
		}
		return response, true, err, ""
	} else {
		return response, false, nil, "Data Already Exist."
	}
}

func InsertMsttemplate(tz *entities.MsttemplateEntity) (int64, bool, error, string) {
	logger.Log.Println("In side Msttemplatemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMsttemplate(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.InsertMsttemplate(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		for _, s := range tz.MaptemplatediffEntities {
			_, err := dataAccess.InsertMsttemplatediff(tz, &s, id)
			if err != nil {
				return 0, false, err, "Something Went Wrong"
			}
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetNotificationEvents() ([]entities.MstNotificationEvent, bool, error, string) {
	logger.Log.Println("In side GetNotificationEvents")
	t := []entities.MstNotificationEvent{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetNotificationEvents()
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err, ""
}

func GetAllNotificationTemplates(page *entities.MstNotificationTemplateEntity) (entities.MstNotificationTemplateEntities, bool, error, string) {
	logger.Log.Println("In side GetAllNotificationTemplates")
	t := entities.MstNotificationTemplateEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(page.ClientID, page.MstOrgnHirarchyID)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values1, err1 := dataAccess.GetAllNotificationTemplates(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	for i, dt := range values1 {
		valdiff, _ := dataAccess.GetAllRecipientsByTemplateID(&dt)
		values1[i].Recipients = valdiff
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetNotificationTemplateCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total
		t.Values = values1
	}
	t.Values = values1
	return t, true, err, ""
}

func GetAllNotificationVariables(page *entities.MstNotificationTemplateEntity) ([]entities.MstNotificationVariable, bool, error, string) {
	logger.Log.Println("In side GetAllNotificationVariables")
	t := []entities.MstNotificationVariable{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetAllNotificationVariables(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	return values, true, err1, ""
}

func GetAllMsttemplate(page *entities.MsttemplateEntity) (entities.MsttemplateEntities, bool, error, string) {
	logger.Log.Println("In side Msttemplatemodel")
	t := entities.MsttemplateEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(page.Clientid, page.Mstorgnhirarchyid)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	values1, err1 := dataAccess.GetAllMsttemplate(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	for i, dt := range values1 {
		valdiff, _ := dataAccess.GetAllMsttemplatediff(&dt)
		values1[i].MaptemplatediffEntities = valdiff
	}
	if page.Offset == 0 {
		total, err1 := dataAccess.GetMsttemplateCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values1
	}
	t.Values = values1
	return t, true, err, ""
}

func DeleteMsttemplate(tz *entities.MsttemplateEntity) (bool, error, string) {
	logger.Log.Println("In side Msttemplatemodel")
	//t := entities.MsttemplateEntities{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMsttemplate(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	err2 := dataAccess.DeleteMsttemplatediff(tz)
	if err2 != nil {
		return false, err2, "Something Went Wrong"
	}
	return true, nil, ""
}

func DeleteNotificationTemplate(tz *entities.MstNotificationTemplateEntity) (bool, error, string) {
	logger.Log.Println("In side DeleteNotificationTemplate")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteNotificationTemplate(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	err2 := dataAccess.DeleteNotificationRecipients(tz)
	if err2 != nil {
		return false, err2, "Something Went Wrong"
	}
	return true, nil, ""
}
func UpdateNotificationTemplate(tz *entities.MstNotificationTemplateEntity) (bool, error, string) {
	logger.Log.Println("In side UpdateNotificationTemplate")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateNotificationTemplate(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count == 0 {
		err := dataAccess.UpdateNotificationTemplate(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		err1 := dataAccess.DeleteNotificationRecipients(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		for i, _ := range tz.Recipients {
			tz.Recipients[i].NotificationTemplateID = tz.ID
			dataAccess.InsertNotificationRecipient(&tz.Recipients[i])
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
func UpdateMsttemplate(tz *entities.MsttemplateEntity) (bool, error, string) {
	logger.Log.Println("In side Msttemplatemodel")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	count, err := dataAccess.CheckDuplicateMsttemplate(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMsttemplate(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		err1 := dataAccess.DeleteMsttemplatediff(tz)
		if err1 != nil {
			return false, err1, "Something Went Wrong"
		}
		for _, s := range tz.MaptemplatediffEntities {
			dataAccess.InsertMsttemplatediff(tz, &s, tz.ID)
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}
