package models

import (
	"errors"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	Logger "iFIX/ifix/logger"
	"net/mail"
	"os"
	"strings"
)

func getEmailContextPath() (string, error) {

	wd, err := os.Getwd()
	if err != nil {
		return "", errors.New("ERROR: Unable to get WD")
	}
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	return contextPath, nil
}

func GetDelimiter(requestData map[string]interface{}) ([]string, error) {

	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	var delimeter []string
	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return delimeter, errors.New("ERROR: Unable to connect DB")
	}
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	var delimiterFetchError error
	delimeter, delimiterFetchError = dao.GetDelimiter(db, clientID, orgID)

	if delimiterFetchError != nil {
		Logger.Log.Println(delimiterFetchError)
		return delimeter, errors.New("delimiter Fetch Error")
	}

	return delimeter, nil
}

func GetServiceUser(requestData map[string]interface{}) (entities.ServiceUserEntities, error) {

	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	usersList := entities.ServiceUserEntities{}
	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return usersList, errors.New("ERROR: Unable to connect DB")
	}
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	//var serviceUsers map[int64]string

	var userListError error
	usersList, userListError = dao.GetSeviceUserList(db, clientID, orgID)
	if userListError != nil {
		Logger.Log.Println(userListError)
		return usersList, errors.New("No service User Found")
	}

	return usersList, nil
}

func GetLastCategoryList(requestData map[string]interface{}) (entities.CategoryList, error) {

	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	categoryLavelID := int64(requestData["categorylevelid"].(float64))
	categoryList := entities.CategoryList{}
	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return categoryList, errors.New("ERROR: Unable to connect DB")
	}
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	var caltegoryListError error
	categoryList, caltegoryListError = dao.GetCategoryList(db, clientID, orgID, categoryLavelID)
	if caltegoryListError != nil {
		Logger.Log.Println(caltegoryListError)
		return categoryList, errors.New("No Categories found Found")
	}

	return categoryList, nil
}

func UpdateEmailTicketConfiguration(requestData map[string]interface{}) error {
	var emailTicketObj entities.MstEmailTicket
	emailTicketObj.ID = int64(requestData["id"].(float64))
	emailTicketObj.ClientID = int64(requestData["clientid"].(float64))
	emailTicketObj.OrgID = int64(requestData["mstorgnhirarchyid"].(float64))
	emailTicketObj.TicketDiffTypeID = int64(requestData["mstrecorddifftypeid"].(float64))
	emailTicketObj.TicketDiffID = int64(requestData["mstrecorddiffid"].(float64))
	emailTicketObj.CategoryDiffTypeID = int64(requestData["categorydifftypeid"].(float64))
	emailTicketObj.CategoryLevelID = int64(requestData["categorylevelid"].(float64))
	emailTicketObj.LastCategoryID = int64(requestData["lastcategoryid"].(float64))
	emailTicketObj.LastCategoryName = requestData["lastcategoryname"].(string)
	emailTicketObj.CategoryIDList = requestData["categoryidlist"].(string)
	emailTicketObj.CategoryNameList = requestData["categorynamelist"].(string)
	emailTicketObj.CategoryWithPath = requestData["categorywithpath"].(string)
	emailTicketObj.ServiceUserID = int64(requestData["serviceuserid"].(float64))
	emailTicketObj.ServiceUserGroupID = int64(requestData["serviceusergroupid"].(float64))
	emailTicketObj.SenderTypeSeq = int64(requestData["sendertypeseq"].(float64))
	emailTicketObj.DefaultSeq = int64(requestData["defaultseq"].(float64))

	emailTicketObj.SenderDomain = requestData["senderdomain"].(string)
	emailTicketObj.CreatedByID = int64(requestData["createdbyid"].(float64))
	if emailTicketObj.DefaultSeq == 0 {
		emailTicketObj.EmailSubKeyword = requestData["emailsubkeyword"].(string)
		emailTicketObj.Delimiter = requestData["delimiter"].(string)
	}
	if emailTicketObj.SenderTypeSeq == 1 {
		emailTicketObj.SenderEmail = requestData["senderemail"].(string)
	} else {
		emailTicketObj.SenderDomain = requestData["senderdomain"].(string)
	}

	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return errors.New("ERROR: Unable to connect DB")
	}
	// defer Logger.Log.Println("DB STats=====> ", db.Stats())
	// defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	// defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	emailTicketObj.SenderEmail = strings.Trim(emailTicketObj.SenderEmail, " ")
	Logger.Log.Println("SenderEmail log", emailTicketObj.SenderEmail)
	if len(emailTicketObj.SenderEmail) > 0 {
		Logger.Log.Println(">>>>>>>>>>>>>", emailTicketObj.SenderEmail)
		var senderEmails []string
		//var pattern = "[a-zA-Z0-9.-_]{1,}@[a-zA-Z.-]{2,}[.]{1}[a-zA-Z]{2,}"
		senderEmails = strings.Split(emailTicketObj.SenderEmail, ",")
		Logger.Log.Println("email log", senderEmails)
		for i := 0; i < len(senderEmails); i++ {

			senderEmails[i] = strings.Trim(senderEmails[i], " ")
			senderEmails[i] = strings.Trim(senderEmails[i], "\n")
			Logger.Log.Println("senderEmails", senderEmails[i])

			match, err := mail.ParseAddress(senderEmails[i])
			Logger.Log.Println("email log", match)
			if err != nil {
				Logger.Log.Println(err)
				return errors.New("ERROR: Invalid Email: " + senderEmails[i])
			}
			// else if match {
			// 	continue
			// } else {
			// 	return errors.New("ERROR: Invalid Email: " + senderEmails[i])
			// }
		}
		for i := 0; i < len(senderEmails); i++ {
			senderEmails[i] = strings.Trim(senderEmails[i], " ")
			senderEmails[i] = strings.Trim(senderEmails[i], "\n")
			emailTicketObj.SenderEmail = senderEmails[i]
			updateError := dao.UpdateMstEmailTicket(db, emailTicketObj)

			if updateError != nil {
				Logger.Log.Println(updateError)
				return errors.New("ERROR: Unable To Update")
			}
		}
	} else if len(emailTicketObj.SenderDomain) > 0 {
		updateError := dao.UpdateMstEmailTicket(db, emailTicketObj)

		if updateError != nil {
			Logger.Log.Println(updateError)
			return errors.New("ERROR: Unable To Update")
		}
	}

	//log.Println(resultset)
	return nil
}

func SaveEmailTicketConfiguration(requestData map[string]interface{}) error {

	var emailTicketObj entities.MstEmailTicket
	emailTicketObj.ClientID = int64(requestData["clientid"].(float64))
	emailTicketObj.OrgID = int64(requestData["mstorgnhirarchyid"].(float64))
	emailTicketObj.TicketDiffTypeID = int64(requestData["mstrecorddifftypeid"].(float64))
	emailTicketObj.TicketDiffID = int64(requestData["mstrecorddiffid"].(float64))
	emailTicketObj.CategoryDiffTypeID = int64(requestData["categorydifftypeid"].(float64))
	emailTicketObj.CategoryLevelID = int64(requestData["categorylevelid"].(float64))
	emailTicketObj.LastCategoryID = int64(requestData["lastcategoryid"].(float64))
	emailTicketObj.LastCategoryName = requestData["lastcategoryname"].(string)
	emailTicketObj.CategoryIDList = requestData["categoryidlist"].(string)
	emailTicketObj.CategoryNameList = requestData["categorynamelist"].(string)
	emailTicketObj.CategoryWithPath = requestData["categorywithpath"].(string)
	emailTicketObj.ServiceUserID = int64(requestData["serviceuserid"].(float64))
	emailTicketObj.ServiceUserGroupID = int64(requestData["serviceusergroupid"].(float64))
	emailTicketObj.SenderTypeSeq = int64(requestData["sendertypeseq"].(float64))
	emailTicketObj.DefaultSeq = int64(requestData["defaultseq"].(float64))
	emailTicketObj.SenderEmail = requestData["senderemail"].(string)
	emailTicketObj.SenderDomain = requestData["senderdomain"].(string)
	emailTicketObj.CreatedByID = int64(requestData["createdbyid"].(float64))
	if emailTicketObj.DefaultSeq == 0 {
		emailTicketObj.EmailSubKeyword = requestData["emailsubkeyword"].(string)
		emailTicketObj.Delimiter = requestData["delimiter"].(string)
	}

	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return errors.New("ERROR: Unable to connect DB")
	}
	// defer Logger.Log.Println("DB STats=====> ", db.Stats())
	// defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	// defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	emailTicketObj.SenderEmail = strings.Trim(emailTicketObj.SenderEmail, " ")
	Logger.Log.Println("SenderEmail log", emailTicketObj.SenderEmail)
	if len(emailTicketObj.SenderEmail) > 0 {
		Logger.Log.Println(">>>>>>>>>>>>>", emailTicketObj.SenderEmail)
		var senderEmails []string
		//var pattern = "[a-zA-Z0-9.-_]{1,}@[a-zA-Z.-]{2,}[.]{1}[a-zA-Z]{2,}"
		senderEmails = strings.Split(emailTicketObj.SenderEmail, ",")
		Logger.Log.Println("email log", senderEmails)
		for i := 0; i < len(senderEmails); i++ {

			senderEmails[i] = strings.Trim(senderEmails[i], " ")
			senderEmails[i] = strings.Trim(senderEmails[i], "\n")
			Logger.Log.Println("senderEmails", senderEmails[i])

			match, err := mail.ParseAddress(senderEmails[i])
			Logger.Log.Println("email log", match)
			if err != nil {
				Logger.Log.Println(err)
				return errors.New("ERROR: Invalid Email: " + senderEmails[i])
			}
			// else if match {
			// 	continue
			// } else {
			// 	return errors.New("ERROR: Invalid Email: " + senderEmails[i])
			// }
		}
		for i := 0; i < len(senderEmails); i++ {
			senderEmails[i] = strings.Trim(senderEmails[i], " ")
			senderEmails[i] = strings.Trim(senderEmails[i], "\n")
			emailTicketObj.SenderEmail = senderEmails[i]
			insertError := dao.InsertMstEmailTicketConfig(db, emailTicketObj)

			if insertError != nil {
				Logger.Log.Println(insertError)
				return errors.New("ERROR: Unable To Insert Properly")
			}
		}
	} else if len(emailTicketObj.SenderDomain) > 0 {
		insertError := dao.InsertMstEmailTicketConfig(db, emailTicketObj)

		if insertError != nil {
			Logger.Log.Println(insertError)
			return errors.New("ERROR: Unable To Insert Properly")
		}
	}

	// insertError := dao.InsertMstEmailTicketConfig(db, emailTicketObj)

	// if insertError != nil {
	// 	Logger.Log.Println(insertError)
	// 	return errors.New("ERROR: Unable To Insert")
	// }

	return nil
}
func GetEmailTicketConfigurations(requestData map[string]interface{}) (entities.MstEmailTicketConfigtList, error) {

	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	limit := int64(requestData["limit"].(float64))
	offset := int64(requestData["offset"].(float64))
	var emailTicketViewList entities.MstEmailTicketConfigtList
	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return emailTicketViewList, errors.New("ERROR: Unable to connect DB")
	}
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	var emailTicketViewListError error
	dataAccess := dao.DbConn{DB: db}
	orgntype, err1 := dataAccess.GetOrgnType(clientID, orgID)
	if err1 != nil {
		return emailTicketViewList, errors.New("Something Went Wrong")
	}
	emailTicketViewList, emailTicketViewListError = dao.GetEmailTicketConfigurations(db, clientID, orgID, limit, offset, orgntype)

	if emailTicketViewListError != nil {
		Logger.Log.Println(emailTicketViewListError)
		return emailTicketViewList, errors.New("GetEmailTicketConfigurations List Error")
	}
	return emailTicketViewList, nil

}

func DeleteEmailTicketConfiguration(requestData map[string]interface{}) error {
	rowID := int64(requestData["id"].(float64))
	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return errors.New("ERROR: Unable to connect DB")
	}
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	deleteEmailTicketConfigurationError := dao.DeleteEmailTicketConfiguration(db, rowID)
	if deleteEmailTicketConfigurationError != nil {
		Logger.Log.Println(deleteEmailTicketConfigurationError)
		return errors.New("Unable to delete")
	}

	return nil
}

func DeleteEmailTicketConfigu(requestData map[string]interface{}) error {
	rowID := int64(requestData["id"].(float64))
	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return errors.New("ERROR: Unable to connect DB")
	}
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	deleteEmailTicketConfiguError := dao.DeleteEmailTicketConfigu(db, rowID)
	if deleteEmailTicketConfiguError != nil {
		Logger.Log.Println(deleteEmailTicketConfiguError)
		return errors.New("Unable to delete")
	}

	return nil
}

func AddEmailBaseConfig(tz *entities.EmailBaseConfig) error {

	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return errors.New("ERROR: Unable to connect DB")
	}
	if tz.OrgTypeID > 2 {

		return errors.New("ERROR: You are not allowed to add Email Config")
	}

	count, countErr := dao.CheckDuplicateEmailBaseConfig(db, tz)
	if countErr != nil {
		Logger.Log.Println(countErr)
		return errors.New("ERROR: Unable to Check Duplicate")
	}
	// if count == 0 {
	addConfigError := dao.AddEmailBaseConfig(db, tz, count)
	if addConfigError != nil {
		Logger.Log.Println(addConfigError)
		return addConfigError
	}
	// } else {
	// 	return errors.New("Data already exist")
	// }

	return nil
}

func GetDelimiterForAllClient(requestData map[string]interface{}) (entities.EmailTicketBaseConfigtList, error) {

	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	limit := int64(requestData["limit"].(float64))
	offset := int64(requestData["offset"].(float64))
	var emailTicketViewList entities.EmailTicketBaseConfigtList
	db, dBerr := config.ConnectMySqlDbSingleton()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return emailTicketViewList, errors.New("ERROR: Unable to connect DB")
	}
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	var emailTicketViewListError error
	emailTicketViewList, emailTicketViewListError = dao.GetDelimiterForAllClient(db, clientID, orgID, limit, offset)

	if emailTicketViewListError != nil {
		Logger.Log.Println(emailTicketViewListError)
		return emailTicketViewList, errors.New("GetEmailTicketConfigurations List Error")
	}
	return emailTicketViewList, nil

}
