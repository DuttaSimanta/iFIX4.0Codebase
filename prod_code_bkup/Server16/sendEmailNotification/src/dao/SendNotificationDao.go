package dao

import (
	"database/sql"
	"errors"
	Logger "src/logger"
	"strings"
)

func GetNotificationFlag(db *sql.DB, clientID int64, orgID int64) (int64, error) {
	var notificationFlag int64

	getNotificationFlagQuery := "select notification from mstorgnhierarchy where clientid=? and id=? and activeflg=1 and deleteflg=0"
	err := db.QueryRow(getNotificationFlagQuery, clientID, orgID).Scan(&notificationFlag)
	if err != nil {
		Logger.Log.Println(err)
		return notificationFlag, errors.New("ERROR: Notification Flag Error")
	}

	return notificationFlag, nil

}
func GetSMSConfiGDetails(db *sql.DB, clientID int64, orgID int64) (string, string, string, string, error) {
	var SMSUrl string
	var IndiaDltPrincipalEntityIdVal string
	var smsUserName string
	var smsPassword string

	getSMTPDetailsQuery := "select credentialkey,credentialendpoint,credentialaccount,credentialpassword from mstclientcredential where clientid=? and mstorgnhirarchyid=? and credentialtypeid=4 and activeflg=1 and deleteflg=0 "
	smtpError := db.QueryRow(getSMTPDetailsQuery, clientID, orgID).Scan(&SMSUrl, &IndiaDltPrincipalEntityIdVal, &smsUserName, &smsPassword)
	if smtpError != nil {
		Logger.Log.Println(smtpError)
		return SMSUrl, IndiaDltPrincipalEntityIdVal, smsUserName, smsPassword, errors.New("ERROR: SMS Configuration Not Found!!!")
	}
	return SMSUrl, IndiaDltPrincipalEntityIdVal, smsUserName, smsPassword, nil

}
func GetEmailSMTPDetails(db *sql.DB, clientID int64, orgID int64) (string, string, string, string, error) {

	var smtpHostForNotification string
	var emailUserName string
	var emailPassword string
	var smtpPort string
	getSMTPDetailsQuery := "select credentialkey,credentialendpoint,credentialaccount,credentialpassword from mstclientcredential where clientid=? and mstorgnhirarchyid=? and credentialtypeid=2 and activeflg=1 and deleteflg=0 "
	smtpError := db.QueryRow(getSMTPDetailsQuery, clientID, orgID).Scan(&smtpHostForNotification, &smtpPort, &emailUserName, &emailPassword)
	if smtpError != nil {
		Logger.Log.Println(smtpError)
		return smtpHostForNotification, smtpPort, emailUserName, emailPassword, errors.New("ERROR: SMTP Confu=iguration Not Found!!!")
	}
	return smtpHostForNotification, smtpPort, emailUserName, emailPassword, nil
}
func TemplateVariableMappingUsingDynamicQueriesForEmails(db *sql.DB, templateSubject string,
	templatebody string, requestData map[string]interface{}) (string, string, error) {

	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	//recordID := int64(requestData["recordid"].(float64))
	//termSeq := int64(requestData["termseq"].(float64))
	var emailSubject string
	var emailBody string
	var templateVariable []string
	var dynamicQuery []string
	var dynamicQueryParam []string
	var getTemplateVariableQuery = "select templatename,query,params from msttemplatevariable where clientid=? and mstorgnhirarchyid=? and queryflag=1 and deleteflg=0 and activeflg=1"
	templateVariableResultset, err := db.Query(getTemplateVariableQuery, clientID, orgID)

	if err != nil {
		Logger.Log.Println(err)
		return emailSubject, emailBody, errors.New("ERROR: Unable to Fetch templateVariableResultset ")
	}
	defer templateVariableResultset.Close()
	for templateVariableResultset.Next() {
		var tempVar string
		var query string
		var param string
		err := templateVariableResultset.Scan(&tempVar, &query, &param)
		if err != nil {
			Logger.Log.Println(err)
			//return emailSubject, emailBody, errors.New("ERROR: Unable to Scan templateVariableResultset ")
		}
		templateVariable = append(templateVariable, tempVar)
		dynamicQuery = append(dynamicQuery, query)
		dynamicQueryParam = append(dynamicQueryParam, param)
	}
	//templateVariableMap := make(map[string]string)
	for i := 0; i < len(templateVariable); i++ {
		var templateVarValues string
		//templateVariable[i] = templateVariable[i]
		paramList := strings.Split(dynamicQueryParam[i], ",")
		mappedParam := []interface{}{}
		for _, v := range paramList {
			mappedParam = append(mappedParam, requestData[v])
		}
		Logger.Log.Println(dynamicQuery[i], mappedParam)
		dynamicQuryResultSet, err := db.Query(dynamicQuery[i], mappedParam...)
		if err != nil {
			Logger.Log.Println(err)
			//return emailSubject, emailBody, errors.New("ERROR: Unable to fetch dynamicQuryResultSet")
		}

		//defer dynamicQuryResultSet.Close()
		for dynamicQuryResultSet.Next() {
			var values string
			err := dynamicQuryResultSet.Scan(&values)
			if err != nil {
				Logger.Log.Println(err)
				//return emailSubject, emailBody, errors.New("ERROR: Unable to scan dynamicQuryResultSet")
			}
			//Logger.Log.Println("values===>", values)
			//	templateVariableMap[templateVariable[i]] = values
			/* if strings.EqualFold(values,""){
				values=""
			} */
			templateVarValues = values

		}
		dynamicQuryResultSet.Close()
		templatebody = strings.Replace(templatebody, templateVariable[i], templateVarValues, -2)
		templateSubject = strings.Replace(templateSubject, templateVariable[i], templateVarValues, -2)
	}
	// if strings.Contains(templatebody, "\n") {
	// 	templatebody = strings.ReplaceAll(templatebody, "\n", "<br/>")
	// }

	emailBody = templatebody
	emailSubject = templateSubject

	return emailSubject, emailBody, nil

}
func TemplateVariableMappingUsingDynamicQueriesForSMS(db *sql.DB,
	templatebody string, requestData map[string]interface{}) (string, error) {

	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	//recordID := int64(requestData["recordid"].(float64))
	//termSeq := int64(requestData["termseq"].(float64))
	//var emailSubject string
	var smsContent string
	var templateVariable []string
	var dynamicQuery []string
	var dynamicQueryParam []string
	var getTemplateVariableQuery = "select templatename,query,params from msttemplatevariable where clientid=? and mstorgnhirarchyid=? and queryflag=1 and deleteflg=0 and activeflg=1"
	templateVariableResultset, err := db.Query(getTemplateVariableQuery, clientID, orgID)
	if err != nil {
		Logger.Log.Println(err)
		return smsContent, errors.New("ERROR: Unable to Fetch templateVariableResultset ")
	}
	defer templateVariableResultset.Close()
	for templateVariableResultset.Next() {
		var tempVar string
		var query string
		var param string
		err := templateVariableResultset.Scan(&tempVar, &query, &param)
		if err != nil {
			Logger.Log.Println(err)
			//return emailSubject, emailBody, errors.New("ERROR: Unable to Scan templateVariableResultset ")
		}
		templateVariable = append(templateVariable, tempVar)
		dynamicQuery = append(dynamicQuery, query)
		dynamicQueryParam = append(dynamicQueryParam, param)
	}
	//templateVariableMap := make(map[string]string)
	for i := 0; i < len(templateVariable); i++ {
		var templateVarValues string
		//templateVariable[i] = templateVariable[i]
		paramList := strings.Split(dynamicQueryParam[i], ",")
		mappedParam := []interface{}{}
		for _, v := range paramList {
			mappedParam = append(mappedParam, requestData[v])
		}
		//Logger.Log.Println(dynamicQuery[i], mappedParam)
		dynamicQuryResultSet, err := db.Query(dynamicQuery[i], mappedParam...)
		if err != nil {
			Logger.Log.Println(err)
			//return emailSubject, emailBody, errors.New("ERROR: Unable to fetch dynamicQuryResultSet")
		}

		//defer dynamicQuryResultSet.Close()
		for dynamicQuryResultSet.Next() {
			var values string
			err := dynamicQuryResultSet.Scan(&values)
			if err != nil {
				Logger.Log.Println(err)
				//return emailSubject, emailBody, errors.New("ERROR: Unable to scan dynamicQuryResultSet")
			}
			//Logger.Log.Println("values===>", values)
			//	templateVariableMap[templateVariable[i]] = values
			/* if strings.EqualFold(values,""){
				values=""
			} */
			templateVarValues = values

		}
		dynamicQuryResultSet.Close()
		templatebody = strings.Replace(templatebody, templateVariable[i], templateVarValues, -2)
		//templateSubject = strings.Replace(templateSubject, templateVariable[i], templateVarValues, -2)
	}
	smsContent = templatebody
	///emailSubject = templateSubject

	return smsContent, nil

}

// func GetUserEmail(db *sql.DB, clientID int64, orgID int64, creatorID int64) error {

// }
