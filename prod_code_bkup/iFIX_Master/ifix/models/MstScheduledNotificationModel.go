package models

import (
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"strconv"
	"strings"
)

func AddMstScheduledNotification(tz *entities.MstScheduledNotificationEntity) (int64, bool, error, string) {
	logger.Log.Println("In side MstScheduledNotification model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	valuesText := []string{}
	for i := range tz.SendToUseridsArray {
		number := tz.SendToUseridsArray[i]
		s := strconv.FormatInt(number, 10)
		valuesText = append(valuesText, s)
	}
	tz.SendToUserids = strings.Join(valuesText, ",")

	Text := []string{}
	for i := range tz.SendToGroupidsArray {
		number := tz.SendToGroupidsArray[i]
		s := strconv.FormatInt(number, 10)
		Text = append(Text, s)
	}
	tz.SendToGroupids = strings.Join(Text, ",")
	// tz.AdditionalRecipint= strings.Join(tz.AdditionalRecipintArray, ",")

	count, err := dataAccess.CheckDuplicateMstScheduledNotification(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		id, err := dataAccess.AddMstScheduledNotification(tz)
		if err != nil {
			return 0, false, err, "Something Went Wrong"
		}
		return id, true, err, ""
	} else {
		return 0, false, nil, "Data Already Exist."
	}
}

func GetAllMstScheduledNotification(page *entities.MstScheduledNotificationEntity) (entities.MstScheduledNotificationEntities, bool, error, string) {
	logger.Log.Println("In side MstScheduledNotification model")
	t := entities.MstScheduledNotificationEntities{}
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
	values, err1 := dataAccess.GetAllMstScheduledNotification(page, orgntype)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}

	//logger.Log.Println("client",values[0].Clientname)
	users := []string{}
	for i := range values {
		if values[i].SendToUserids != "" {
			res1 := strings.Split(values[i].SendToUserids, ",")

			for i := range res1 {
				n, err := strconv.ParseInt(res1[i], 10, 64)
				if err != nil {
					logger.Log.Println("ParseInt error in GetAllMstScheduledNotificationuser in Model", err)
					return t, false, err, "Something Went Wrong"
				}

				user, err := dataAccess.GetUser(n)
				logger.Log.Println(n, user, users)
				users = append(users, user)
			}
			values[i].SendToUserNames = strings.Join(users, ",")
		}
	}

	groups := []string{}
	for i := range values {
		if values[i].SendToGroupids != "" {
			res1 := strings.Split(values[i].SendToGroupids, ",")
			logger.Log.Println("length", len(res1))
			for i := range res1 {
				n, err := strconv.ParseInt(res1[i], 10, 64)
				if err != nil {
					logger.Log.Println("ParseInt error in GetAllMstScheduledNotificationgroup in Model", err)
					return t, false, err, "Something Went Wrong"
				}
				group, err := dataAccess.GetGroup(n)
				if err != nil {
					return t, false, err1, "Something Went Wrong"
				}
				logger.Log.Println(n, group, groups)
				groups = append(groups, group)
			}
			values[i].SendToGroupNames = strings.Join(groups, ",")
		}
	}

	if page.Offset == 0 {
		total, err1 := dataAccess.GetMstScheduledNotificationCount(page, orgntype)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		}
		t.Total = total.Total
		t.Values = values
	}
	t.Values = values
	return t, true, err, ""
}

func DeleteMstScheduledNotification(tz *entities.MstScheduledNotificationEntity) (bool, error, string) {
	logger.Log.Println("In side MstScheduledNotification model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.DeleteMstScheduledNotification(tz)
	if err1 != nil {
		return false, err1, "Something Went Wrong"
	}
	return true, nil, ""
}

func UpdateMstScheduledNotification(tz *entities.MstScheduledNotificationEntity) (bool, error, string) {
	logger.Log.Println("In side MstScheduledNotification model", tz.AdditionalRecipint)
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	valuesText := []string{}
	for i := range tz.SendToUseridsArray {
		number := tz.SendToUseridsArray[i]
		s := strconv.FormatInt(number, 10)
		valuesText = append(valuesText, s)
	}
	tz.SendToUserids = strings.Join(valuesText, ",")

	Text := []string{}
	for i := range tz.SendToGroupidsArray {
		number := tz.SendToGroupidsArray[i]
		s := strconv.FormatInt(number, 10)
		Text = append(Text, s)
	}
	tz.SendToGroupids = strings.Join(Text, ",")
	//tz.AdditionalRecipint= strings.Join(tz.AdditionalRecipintArray, ",")

	count, err := dataAccess.CheckDuplicateMstScheduledNotification(tz)
	if err != nil {
		return false, err, "Something Went Wrong"
	}
	if count.Total == 0 {
		err := dataAccess.UpdateMstScheduledNotification(tz)
		if err != nil {
			return false, err, "Something Went Wrong"
		}
		return true, err, ""
	} else {
		return false, nil, "Data Already Exist."
	}
}

func GetClientAndOrgWiseclientuser(page *entities.MstScheduledNotificationEntity) ([]entities.GetClientAndOrgWiseclientuserEntity, bool, error, string) {
	logger.Log.Println("In side MstScheduledNotification model")
	t := []entities.GetClientAndOrgWiseclientuserEntity{}
	db, err := config.ConnectMySqlDb()
	defer db.Close()
	if err != nil {
		logger.Log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetClientAndOrgWiseclientuser(page)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	t = values
	return t, true, err, ""
}
