package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"src/config"
	"src/dao"
	"src/entities"
	SendMailUtils "src/fileutils"
	Logger "src/logger"
	"strconv"
	"strings"
)

//var mutex = &sync.Mutex{}

// func TemplateVariableMappingUsingDynamicQueries(db *sql.DB, templateSubject string,
// 	templatebody string, requestData map[string]interface{}) (string, string, error) {

// 	clientID := int64(requestData["clientid"].(float64))
// 	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
// 	//recordID := int64(requestData["recordid"].(float64))
// 	//termSeq := int64(requestData["termseq"].(float64))
// 	var emailSubject string
// 	var emailBody string
// 	var templateVariable []string
// 	var dynamicQuery []string
// 	var dynamicQueryParam []string
// 	var getTemplateVariableQuery = "select templatename,query,params from msttemplatevariable where clientid=? and mstorgnhirarchyid=? and queryflag=1 and deleteflg=0 and activeflg=1"
// 	templateVariableResultset, err := db.Query(getTemplateVariableQuery, clientID, orgID)

// 	if err != nil {
// 		Logger.Log.Println(err)
// 		return emailSubject, emailBody, errors.New("ERROR: Unable to Fetch templateVariableResultset ")
// 	}
// 	defer templateVariableResultset.Close()
// 	for templateVariableResultset.Next() {
// 		var tempVar string
// 		var query string
// 		var param string
// 		err := templateVariableResultset.Scan(&tempVar, &query, &param)
// 		if err != nil {
// 			Logger.Log.Println(err)
// 			//return emailSubject, emailBody, errors.New("ERROR: Unable to Scan templateVariableResultset ")
// 		}
// 		templateVariable = append(templateVariable, tempVar)
// 		dynamicQuery = append(dynamicQuery, query)
// 		dynamicQueryParam = append(dynamicQueryParam, param)
// 	}
// 	//templateVariableMap := make(map[string]string)
// 	for i := 0; i < len(templateVariable); i++ {
// 		var templateVarValues string
// 		//templateVariable[i] = templateVariable[i]
// 		paramList := strings.Split(dynamicQueryParam[i], ",")
// 		mappedParam := []interface{}{}
// 		for _, v := range paramList {
// 			mappedParam = append(mappedParam, requestData[v])
// 		}
// 		//Logger.Log.Println(dynamicQuery[i], mappedParam)
// 		dynamicQuryResultSet, err := db.Query(dynamicQuery[i], mappedParam...)
// 		if err != nil {
// 			Logger.Log.Println(err)
// 			//return emailSubject, emailBody, errors.New("ERROR: Unable to fetch dynamicQuryResultSet")
// 		}

// 		//defer dynamicQuryResultSet.Close()
// 		for dynamicQuryResultSet.Next() {
// 			var values string
// 			err := dynamicQuryResultSet.Scan(&values)
// 			if err != nil {
// 				Logger.Log.Println(err)
// 				//return emailSubject, emailBody, errors.New("ERROR: Unable to scan dynamicQuryResultSet")
// 			}
// 			//Logger.Log.Println("values===>", values)
// 			//	templateVariableMap[templateVariable[i]] = values
// 			/* if strings.EqualFold(values,""){
// 				values=""
// 			} */
// 			templateVarValues = values

// 		}
// 		dynamicQuryResultSet.Close()
// 		templatebody = strings.Replace(templatebody, templateVariable[i], "<b>"+templateVarValues+"</b>", -2)
// 		templateSubject = strings.Replace(templateSubject, templateVariable[i], templateVarValues, -2)
// 	}
// 	emailBody = templatebody
// 	emailSubject = templateSubject

// 	return emailSubject, emailBody, nil

// }

func GetToEmails(creatorEmail string, originalCreatorEmail string, assigneeEmail string, assigneeSupportGrpEmail string,
	assigneeSupportGrpMemberEmail string) string {
	var userTOEmails string
	if len(creatorEmail) > 0 {
		userTOEmails = userTOEmails + "," + creatorEmail
	}
	if len(originalCreatorEmail) > 0 {
		userTOEmails = userTOEmails + "," + originalCreatorEmail
	}
	if len(assigneeEmail) > 0 {
		userTOEmails = userTOEmails + "," + assigneeEmail
	}
	if len(assigneeSupportGrpEmail) > 0 {
		userTOEmails = userTOEmails + "," + assigneeSupportGrpEmail
	}
	if len(assigneeSupportGrpMemberEmail) > 0 {
		userTOEmails = userTOEmails + "," + assigneeSupportGrpMemberEmail
	}
	// Logger.Log.Println("creatorEmail====>", creatorEmail)
	// Logger.Log.Println("originalCreatorEmail====>", originalCreatorEmail)
	// Logger.Log.Println("assigneeEmail====>", assigneeEmail)
	// Logger.Log.Println("assigneeSupportGrpEmail====>", assigneeSupportGrpEmail)
	// Logger.Log.Println("assigneeSupportGrpMemberEmail====>", assigneeSupportGrpMemberEmail)
	// Logger.Log.Println("userTOEmails====>", userTOEmails)

	return userTOEmails
}
func GetEmailIDs(db *sql.DB, clientID int64, orgID int64, recordID int64,
	additionalRecipient string, creatorID int64, originalCreatorID int64,
	sendToCreator int64, sendToOriginalCreator int64, sendToAssignee int64,
	sendToAssigneeGroup int64, sendToAssigneeGroupMember int64, notificationTemplateID int64) (string, string, error) {

	// Logger.Log.Println("<==============IN GetEmailIDs========================>")
	// Logger.Log.Print("recordID==> ", recordID)
	// Logger.Log.Print("creatorID==> ", creatorID)
	// Logger.Log.Print("originalCreatorID==> ", originalCreatorID)
	// Logger.Log.Print("sendToCreator==> ", sendToCreator)
	// Logger.Log.Print("sendToOriginalCreator==> ", sendToOriginalCreator)
	// Logger.Log.Print("sendToAssignee==> ", sendToAssignee)
	// Logger.Log.Print("sendToAssigneeGroup==> ", sendToAssigneeGroup)
	// Logger.Log.Print("sendToAssigneeGroupMember==> ", sendToAssigneeGroupMember)
	// Logger.Log.Print("notificationTemplateID==> ", notificationTemplateID)
	// Logger.Log.Print("additionalRecipient==> ", additionalRecipient)
	// Logger.Log.Println("<====================================================>")

	var userTOEmails string
	var userCCEmails string
	var creatorEmail string
	var originalCreatorEmail string
	var assigneeEmail string
	var assigneeSupportGrpEmail string
	var assigneeSupportGrpMemberEmail string
	var groupID int64
	//111
	if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		// creatorEmail, getEmailResultseterr = dao.GetUserEmail(db, clientID, orgID, creatorID)
		// if getEmailResultseterr != nil {
		// 	Logger.Log.Println(getEmailResultseterr)
		// 	//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		// }
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()
		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//originalCreatorEmail = creatorEmail
		//110
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//101
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()
		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//100
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//011
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()
		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		} //
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()
		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//000
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//111 for originalcreater!=1
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//110
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()
		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//100

	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//011
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//010
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		} //
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()
		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//000
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//111
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//110
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//101
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//100
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//011
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//010
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		} //
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//001
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//000
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//111
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {

		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		//originalCreatorEmail = creatorEmail
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//110
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {

		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//101
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {

		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//100
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//011
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {

		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//010
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		} //
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//001
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {

		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//000
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		// creator not sender
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//110
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//101
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//100
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//011
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//010
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		} //

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//000
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//111 for originalcreater!=1
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//110
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//101
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//100
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//011
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//010
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		} //
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//000
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//111
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//110
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)
		//101
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//100
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeEmailAddress := "select useremail from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getAssigneeEmailResultseterr := db.QueryRow(getAssigneeEmailAddress, clientID, orgID, recordID).Scan(&assigneeEmail)
		if getAssigneeEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}

		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//011
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//010
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getAssigneeSupportGrpEmailAddress := "select email,grpid from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpEmailResultseterr := db.QueryRow(getAssigneeSupportGrpEmailAddress, clientID, orgID, recordID).Scan(&assigneeSupportGrpEmail, &groupID)
		if getAssigneeSupportGrpEmailResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Assignee Supportgroup")
		} //
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//001
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getEmailAddOriginalCreatorquery := " select useremail from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getEmailOriginalCreatorResultseterr := db.QueryRow(getEmailAddOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalCreatorEmail)
		if getEmailOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getEmailOriginalCreatorResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflag=0 and activeflag=1"
		userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, recordID)
		if UserEmailResultsetError != nil {
			Logger.Log.Println(UserEmailResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userEmailResultset.Close()

		for userEmailResultset.Next() {
			var userEmailForAgroupMemer string
			scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMemberEmail = assigneeSupportGrpMemberEmail + "," + userEmailForAgroupMemer
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

		//000
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getEmailAddquery := "select useremail from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getEmailResultseterr := db.QueryRow(getEmailAddquery, creatorID).Scan(&creatorEmail)
		if getEmailResultseterr != nil {
			Logger.Log.Println(getEmailResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Creator")
		}
		userTOEmails = GetToEmails(creatorEmail, originalCreatorEmail, assigneeEmail, assigneeSupportGrpEmail, assigneeSupportGrpMemberEmail)

	}

	//CCGrp Emails
	var getCCGroupEmailAddquery = "select grpid,email from mstclientsupportgroup where grpid in(select groupid from mstnotificationrecipients where clientid=? and mstorgnhirarchyid=? and activeflg=1 and deleteflg=0 and notificationtemplateid=? and recipienttype='CC')"
	getCCGroupEmailResultset, getCCGroupEmailResultseErr := db.Query(getCCGroupEmailAddquery, clientID, orgID, notificationTemplateID)
	if getCCGroupEmailResultseErr != nil {
		Logger.Log.Println(getCCGroupEmailResultseErr)
		//return errors.New("ERROR: Unable to fetch EmailAddress of group")
	}
	defer getCCGroupEmailResultset.Close()
	for getCCGroupEmailResultset.Next() {
		var grpID int64
		var grpEmail string
		scanErr := getCCGroupEmailResultset.Scan(&grpID, &grpEmail)
		if scanErr != nil {
			Logger.Log.Println(scanErr)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		if len(grpEmail) == 0 {
			getUserEmailOfGrpQuery := "select useremail from mstclientuser where id in(select userid from mstgroupmember where clientid=? and mstorgnhirarchyid=? and groupid=? and activeflg=1 and deleteflg=0)"
			userEmailResultset, UserEmailResultsetError := db.Query(getUserEmailOfGrpQuery, clientID, orgID, grpID)
			if UserEmailResultsetError != nil {
				Logger.Log.Println(UserEmailResultsetError)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			defer userEmailResultset.Close()

			for userEmailResultset.Next() {
				var userEmailForAgroupMemer string
				scanErr := userEmailResultset.Scan(&userEmailForAgroupMemer)
				if scanErr != nil {
					Logger.Log.Println(scanErr)
					//return errors.New("ERROR: Unable to fetch EmailAddress of group")
				}
				userCCEmails = userCCEmails + "," + userEmailForAgroupMemer
			}

		}
		userCCEmails = grpEmail
	}
	// Logger.Log.Println("userCCEmails", userCCEmails)
	// Logger.Log.Println("Additional Receipint in getEmails", additionalRecipient)
	if !strings.EqualFold(additionalRecipient, "") {
		userCCEmails = userCCEmails + "," + additionalRecipient
	}
	var toUserEmail []string
	var toEmailAddress string
	toUserEmail = strings.Split(userTOEmails, ",")
	for i := 0; i < len(toUserEmail); i++ {
		if len(toUserEmail[i]) > 0 {
			toEmailAddress = toEmailAddress + "," + toUserEmail[i]
		}
	}
	// for i:=0 ;range toUserEmails

	Logger.Log.Println("<========================ExitingFrom Get emailids===================>")
	return toEmailAddress, userCCEmails, nil
}
func SendEMail(db *sql.DB, emailSubject string, emailBody string, userTOEmails string, userCCEmails string,
	clientID int64, orgID int64, recordID int64, eventNotificationID int64, notificationTemplateID int64) error {
	userTOEmails = strings.Trim(userTOEmails, ",")
	userCCEmails = strings.Trim(userCCEmails, ",")
	Logger.Log.Println("Subject===> ", emailSubject)
	Logger.Log.Println("Body====> ", emailBody)
	Logger.Log.Println("TO Emails==> ", userTOEmails)
	Logger.Log.Println("CC Emails==> ", userCCEmails)
	smtpHostForNotification, smtpPort, emailUserName, emailPassword, smtpError := dao.GetEmailSMTPDetails(db, clientID, orgID)
	if smtpError != nil {
		Logger.Log.Println(smtpError)
		return errors.New("ERROR: SMTP Configuration Not Found!!!")
	}
	sendMailErr := SendMailUtils.SendMaiL(emailSubject, emailBody, userTOEmails, userCCEmails, smtpHostForNotification, smtpPort, emailUserName, emailPassword)
	if sendMailErr != nil {
		Logger.Log.Println(sendMailErr)
		notificationLogQuery := "INSERT INTO `mstnotificationlog`(`clientid`,`mstorgnhirarchyid`,`recordid`,`notificationeventid`,`notificationtemplateid`,`processflag`,`deleteflg`,`activeflg`)VALUES(?,?,?,?,?,?,?,?)"
		Resultset, err := db.Query(notificationLogQuery, clientID, orgID, recordID, eventNotificationID, notificationTemplateID, "N", 0, 1)
		if err != nil {
			Logger.Log.Println(err)
			return err
		}
		Resultset.Close()
		return sendMailErr
	} else {
		notificationLogQuery := "INSERT INTO `mstnotificationlog`(`clientid`,`mstorgnhirarchyid`,`recordid`,`notificationeventid`,`notificationtemplateid`,`processflag`,`deleteflg`,`activeflg`)VALUES(?,?,?,?,?,?,?,?)"
		Resultset, err := db.Query(notificationLogQuery, clientID, orgID, recordID, eventNotificationID, notificationTemplateID, "Y", 0, 1)
		if err != nil {
			Logger.Log.Println(err)
			return err
		}
		Resultset.Close()
	}
	return nil
}
func StatusChangeEvent(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	channelType := int64(requestData["channeltype"].(float64))
	//termSeq := int64(requestData["termseq"].(float64))
	statusSeq := int64(requestData["statusid"].(float64))

	if channelType != 1 {
		return errors.New("Wrong Channel Type")
	}

	var convertedString string = ""
	var convertedInt int64 = 2
	getConvertedQuery := "select recordtrackvalue from trnreordtracking where recordtermid in (SELECT id FROM mstrecordterms where clientid=? and mstorgnhirarchyid=? and seq=79) and recordid=?"
	ConvertedErr := db.QueryRow(getConvertedQuery, clientID, orgID, recordID).Scan(&convertedString)
	if ConvertedErr != nil {
		Logger.Log.Println(ConvertedErr)
		//return errors.New("Unable To fetch Converted Val")
	}

	if strings.EqualFold(convertedString, "YES") {
		convertedInt = 1
	}
	Logger.Log.Println("convertedInt===>", convertedInt)
	var notificationTemplateID int64
	var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	//var eventParams interface{}

	var creatorID int64
	var originalCreatorID int64
	//var groupID int64
	getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
	creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
	if creatorIDresulsetErr != nil {
		Logger.Log.Println(creatorIDresulsetErr)
		errors.New("Unable To fetch CreatorID")
	}

	var eventParam entities.EventParam
	eventParam.StatusSeq = statusSeq
	eventParamBytesArray, jsonMarshallErr := json.Marshal(eventParam)
	if jsonMarshallErr != nil {
		Logger.Log.Println(jsonMarshallErr)
		return errors.New("Error in status Change Status Parsing")
	}

	Logger.Log.Println("working cat===>", workingCategoryID)
	eventParamstring := string(eventParamBytesArray)
	Logger.Log.Println("eventParamstring===> ", eventParamstring)
	getSubjectAndBodyQury := "select id,subjectortitle,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=? and eventparams=? and converted=? and activeflg=1 and deleteflg=0"
	resultset, getSubjectAndBodyErr := db.Query(getSubjectAndBodyQury, clientID, orgID, eventNotificationID, channelType, workingCategoryID, eventParamstring, convertedInt)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	defer resultset.Close()
	for resultset.Next() {

		scanErr := resultset.Scan(&notificationTemplateID, &subjectTitle, &body, &additionalRecipient, &sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember)
		if scanErr != nil {
			Logger.Log.Println(getSubjectAndBodyErr)
			//errors.New("Workin Category is Not mapped properly")
			//return errors.New("Workin Category is Not mapped properly")
		}
		emailSubject, emailBody, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForEmails(db, subjectTitle, body, requestData)

		if templateVariableMappingUsingDynamicQueriesError != nil {
			Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
			errors.New("Workin Category is Not mapped properly")
		}
		Logger.Log.Println("Subject===> ", emailSubject)
		Logger.Log.Println("Body====> ", emailBody)
		Logger.Log.Println("Notification Template id", notificationTemplateID)

		userTOEmails, userCCEmails, getEmailIDError := GetEmailIDs(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember, notificationTemplateID)
		if getEmailIDError != nil {
			Logger.Log.Println(getEmailIDError)
			//errors.New("Unable To fetch CreatorID")
		}

		sendEmailError := SendEMail(db, emailSubject, emailBody, userTOEmails, userCCEmails, clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
		if sendEmailError != nil {
			Logger.Log.Println(sendEmailError)
			//return sendEmailError
		}
	}

	// log.Println("ID===>", notificationTemplateID)
	// log.Println("Sub===>", subjectTitle)
	// log.Println("body", body)
	// log.Println("additional ===>", additionalRecipient)

	//var userTOEmails string
	//var userCCEmails string
	// var creatorEmail string
	// var originalCreatorEmail string
	// var assigneeEmail string
	// var assigneeSupportGrpEmail string
	// var assigneeSupportGrpMemberEmail string

	return nil

}
func CustomerVisibleWorkNoteEvent(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	channelType := int64(requestData["channeltype"].(float64))
	//termSeq := int64(requestData["termseq"].(float64))
	//	statusSeq := int64(requestData["statusid"].(float64))

	if channelType != 1 {
		return errors.New("Wrong Channel Type")
	}
	log.Println("working cat===>", workingCategoryID)
	//	eventParamstring := string(eventParamBytesArray)
	//Logger.Log.Println("eventParamstring===> ", eventParamstring)

	var notificationTemplateID int64
	var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	getSubjectAndBodyQury := "select id,subjectortitle,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=? and activeflg=1 and deleteflg=0"
	getSubjectAndBodyErr := db.QueryRow(getSubjectAndBodyQury, clientID, orgID, eventNotificationID, channelType, workingCategoryID).Scan(&notificationTemplateID, &subjectTitle, &body, &additionalRecipient, &sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	log.Println("ID===>", notificationTemplateID)
	log.Println("Sub===>", subjectTitle)
	log.Println("body", body)
	log.Println("additional ===>", additionalRecipient)
	emailSubject, emailBody, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForEmails(db, subjectTitle, body, requestData)

	if templateVariableMappingUsingDynamicQueriesError != nil {
		Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
		return errors.New("Workin Category is Not mapped properly")
	}
	Logger.Log.Println("Subject===> ", emailSubject)
	Logger.Log.Println("Body====> ", emailBody)
	Logger.Log.Println("Notification Template id", notificationTemplateID)
	//var userTOEmails string
	//var userCCEmails string
	// var creatorEmail string
	// var originalCreatorEmail string
	// var assigneeEmail string
	// var assigneeSupportGrpEmail string
	// var assigneeSupportGrpMemberEmail string
	var creatorID int64
	var originalCreatorID int64
	//var groupID int64
	getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
	creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
	if creatorIDresulsetErr != nil {
		Logger.Log.Println(creatorIDresulsetErr)
		return errors.New("Unable To fetch CreatorID")
	}
	userTOEmails, userCCEmails, getEmailIDError := GetEmailIDs(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember, notificationTemplateID)
	if getEmailIDError != nil {
		Logger.Log.Println(getEmailIDError)
		return errors.New("Unable To fetch CreatorID")
	}

	sendEmailError := SendEMail(db, emailSubject, emailBody, userTOEmails, userCCEmails, clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
	if sendEmailError != nil {
		Logger.Log.Println(sendEmailError)
		return sendEmailError
	}
	return nil
}
func UserFollowUpCountEvent(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	channelType := int64(requestData["channeltype"].(float64))
	//termSeq := int64(requestData["termseq"].(float64))
	//	statusSeq := int64(requestData["statusid"].(float64))

	if channelType != 1 {
		return errors.New("Wrong Channel Type")
	}
	log.Println("working cat===>", workingCategoryID)
	//	eventParamstring := string(eventParamBytesArray)
	//Logger.Log.Println("eventParamstring===> ", eventParamstring)
	var countVal int64
	var lastAssignmentGrp string
	var lastTolastAssignmentGrp string
	if eventNotificationID == 5 {
		countVal = int64(requestData["followupcount"].(float64))
	} else if eventNotificationID == 6 {
		countVal = int64(requestData["hopcount"].(float64))
		lastAssignmentGrp = requestData["lastgroupname"].(string)
		Logger.Log.Println(lastAssignmentGrp)
		lastTolastAssignmentGrp = requestData["lasttolastgroupname"].(string)
	}

	var notificationTemplateID int64
	var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	var eventParam string
	var eventParamObj entities.EventParam
	//eventParam.NoOfCount = countVal
	// eventParamBytesArray, jsonMarshallErr := json.Marshal(eventParam)
	// if jsonMarshallErr != nil {
	// 	Logger.Log.Println(jsonMarshallErr)
	// 	return errors.New("Error in  Parsing for no of count")
	// }
	log.Println("working cat===>", workingCategoryID)
	// eventParamstring := string(eventParamBytesArray)
	// Logger.Log.Println("eventParamstring===> ", eventParamstring)
	var count int64 = 0
	getSubjectAndBodyQury := "select id,subjectortitle,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember,eventparams from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=? and activeflg=1 and deleteflg=0"
	Resultset, getSubjectAndBodyErr := db.Query(getSubjectAndBodyQury, clientID, orgID, eventNotificationID, channelType, workingCategoryID)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	defer Resultset.Close()
	for Resultset.Next() {
		scanErr := Resultset.Scan(&notificationTemplateID, &subjectTitle, &body, &additionalRecipient,
			&sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember,
			&eventParam)
		if scanErr != nil {
			Logger.Log.Println(scanErr)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		unmarshalError := json.Unmarshal([]byte(eventParam), &eventParamObj)
		if unmarshalError != nil {
			Logger.Log.Println(unmarshalError)
			return errors.New("ERROR: Unable to unmarshal event param")
		}
		if eventParamObj.NoOfCount < countVal {
			count++
			break
		}
	}
	if count == 1 {

		log.Println("ID===>", notificationTemplateID)
		log.Println("Sub===>", subjectTitle)
		log.Println("body", body)
		log.Println("additional ===>", additionalRecipient)
		emailSubject, emailBody, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForEmails(db, subjectTitle, body, requestData)

		if templateVariableMappingUsingDynamicQueriesError != nil {
			Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
			return errors.New("Workin Category is Not mapped properly")
		}

		if eventNotificationID == 6 {
			Logger.Log.Println("In Hop Count Value Replacement")

			hopcountValue := strconv.Itoa(int(countVal))
			Logger.Log.Println("Hop Count Value=====>", hopcountValue)
			hopCount := "{{HopCount}}"
			lastToLastAssignmentGrpVar := "{{LastAssignmentTower}}"

			if strings.Contains(emailSubject, hopCount) {
				emailSubject = strings.ReplaceAll(emailSubject, hopCount, hopcountValue)
			}
			if strings.Contains(emailBody, hopCount) {
				emailBody = strings.ReplaceAll(emailBody, hopCount, hopcountValue)
			}

			if strings.Contains(emailSubject, lastToLastAssignmentGrpVar) {
				emailSubject = strings.ReplaceAll(emailSubject, lastToLastAssignmentGrpVar, lastTolastAssignmentGrp)
			}
			if strings.Contains(emailBody, lastToLastAssignmentGrpVar) {
				emailBody = strings.ReplaceAll(emailBody, lastToLastAssignmentGrpVar, lastTolastAssignmentGrp)
			}

		}
		Logger.Log.Println("Subject After HopCount===> ", emailSubject)
		Logger.Log.Println("Body After hopCount====> ", emailBody)
		Logger.Log.Println("Notification Template id", notificationTemplateID)
		//var userTOEmails string
		//var userCCEmails string
		// var creatorEmail string
		// var originalCreatorEmail string
		// var assigneeEmail string
		// var assigneeSupportGrpEmail string
		// var assigneeSupportGrpMemberEmail string
		var creatorID int64
		var originalCreatorID int64
		//var groupID int64
		getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
		creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
		if creatorIDresulsetErr != nil {
			Logger.Log.Println(creatorIDresulsetErr)
			return errors.New("Unable To fetch CreatorID")
		}
		userTOEmails, userCCEmails, getEmailIDError := GetEmailIDs(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember, notificationTemplateID)
		if getEmailIDError != nil {
			Logger.Log.Println(getEmailIDError)
			return errors.New("Unable To fetch CreatorID")
		}
		sendEmailError := SendEMail(db, emailSubject, emailBody, userTOEmails, userCCEmails, clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
		if sendEmailError != nil {
			Logger.Log.Println(sendEmailError)
			return sendEmailError
		}
	} else {
		Logger.Log.Println("Email not sent for noofcount as criteria not satisfies")
	}

	return nil
}
func PriorityChangeEvent(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	channelType := int64(requestData["channeltype"].(float64))
	priorityID := int64(requestData["priorityid"].(float64))
	var eventParam entities.EventParam
	eventParam.PriorityID = priorityID
	eventParamBytesArray, jsonMarshallErr := json.Marshal(eventParam)
	if jsonMarshallErr != nil {
		Logger.Log.Println(jsonMarshallErr)
		return errors.New("Error in  Parsing for no of count")
	}
	log.Println("working cat===>", workingCategoryID)
	eventParamstring := string(eventParamBytesArray)
	Logger.Log.Println("eventParamstring===> ", eventParamstring)
	var notificationTemplateID int64
	var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	getSubjectAndBodyQury := "select id,subjectortitle,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=? and eventparams=? and activeflg=1 and deleteflg=0"
	getSubjectAndBodyErr := db.QueryRow(getSubjectAndBodyQury, clientID, orgID, eventNotificationID, channelType, workingCategoryID, eventParamstring).Scan(&notificationTemplateID, &subjectTitle, &body, &additionalRecipient, &sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	log.Println("ID===>", notificationTemplateID)
	log.Println("Sub===>", subjectTitle)
	log.Println("body", body)
	log.Println("additional ===>", additionalRecipient)

	emailSubject, emailBody, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForEmails(db, subjectTitle, body, requestData)

	if templateVariableMappingUsingDynamicQueriesError != nil {
		Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
		return errors.New("Workin Category is Not mapped properly")
	}
	Logger.Log.Println("Subject===> ", emailSubject)
	Logger.Log.Println("Body====> ", emailBody)
	Logger.Log.Println("Notification Template id", notificationTemplateID)

	var currentPrioritySeq int64
	var previousPrioritySeq int64

	getCurrentPrioritySeq := "select b.seqno from maprecordtorecorddifferentiation a,mstrecorddifferentiation b where a.clientid=? and a.mstorgnhirarchyid=? and a.recordid=? and a.recorddifftypeid=5 and a.recorddiffid=b.id order by a.id desc limit 1"
	getCurrentPrioritySeqErr := db.QueryRow(getCurrentPrioritySeq, clientID, orgID, recordID).Scan(&currentPrioritySeq)
	if getCurrentPrioritySeqErr != nil {
		Logger.Log.Println(getCurrentPrioritySeqErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("get current priority Seq error")
	}

	getPreviousPrioritySeq := "select b.seqno from maprecordtorecorddifferentiation a,mstrecorddifferentiation b where a.clientid=? and a.mstorgnhirarchyid=? and a.recordid=? and a.recorddifftypeid=5 and a.recorddiffid=b.id order by a.id desc limit 1 offset 1"
	getPreviousPrioritySeqErr := db.QueryRow(getPreviousPrioritySeq, clientID, orgID, recordID).Scan(&previousPrioritySeq)
	if getPreviousPrioritySeqErr != nil {
		Logger.Log.Println(getPreviousPrioritySeqErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("get Previous priority Seq error")
	}
	var AdditionalInfo string
	if currentPrioritySeq > previousPrioritySeq {
		AdditionalInfo = "Downgrade"
	} else {
		AdditionalInfo = "Upgrade"
	}
	var additionalInfoVar string = "{{PriorityChangeType}}"

	if strings.Contains(emailSubject, additionalInfoVar) {
		emailSubject = strings.ReplaceAll(emailSubject, additionalInfoVar, AdditionalInfo)
	}
	if strings.Contains(emailBody, additionalInfoVar) {
		emailBody = strings.ReplaceAll(emailBody, additionalInfoVar, AdditionalInfo)
	}
	//var userTOEmails string
	//var userCCEmails string
	// var creatorEmail string
	// var originalCreatorEmail string
	// var assigneeEmail string
	// var assigneeSupportGrpEmail string
	// var assigneeSupportGrpMemberEmail string
	var creatorID int64
	var originalCreatorID int64
	//var groupID int64
	getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
	creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
	if creatorIDresulsetErr != nil {
		Logger.Log.Println(creatorIDresulsetErr)
		return errors.New("Unable To fetch CreatorID")
	}
	userTOEmails, userCCEmails, getEmailIDError := GetEmailIDs(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember, notificationTemplateID)
	if getEmailIDError != nil {
		Logger.Log.Println(getEmailIDError)
		return errors.New("Unable To fetch CreatorID")
	}

	sendEmailError := SendEMail(db, emailSubject, emailBody, userTOEmails, userCCEmails, clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
	if sendEmailError != nil {
		Logger.Log.Println(sendEmailError)
		return sendEmailError
	}

	return nil
}
func SLAResponseEvent(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	//channelType := int64(requestData["channeltype"].(float64))
	notificationTemplateID := int64(requestData["notificationtemplateid"].(float64))
	slaPercentage := requestData["percentage"].(float64)
	slaPercentageString := fmt.Sprintf("%f", slaPercentage)

	//var notificationTemplateID int64
	var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	var eventParamstring string
	getSubjectAndBodyQury := "select subjectortitle,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember,eventparams from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and id=? and activeflg=1 and deleteflg=0"
	getSubjectAndBodyErr := db.QueryRow(getSubjectAndBodyQury, clientID, orgID, notificationTemplateID).Scan(&subjectTitle, &body, &additionalRecipient, &sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember, &eventParamstring)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	var eventParamObj entities.EventParam
	unmarshalError := json.Unmarshal([]byte(eventParamstring), &eventParamObj)
	if unmarshalError != nil {
		Logger.Log.Println(unmarshalError)
		return errors.New("ERROR: Unable to unmarshal event param")
	}
	log.Println("ID===>", notificationTemplateID)
	log.Println("Sub===>", subjectTitle)
	log.Println("body", body)
	log.Println("additional ===>", additionalRecipient)
	emailSubject, emailBody, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForEmails(db, subjectTitle, body, requestData)

	if templateVariableMappingUsingDynamicQueriesError != nil {
		Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
		return errors.New("Workin Category is Not mapped properly")
	}
	slaTemplateVarForProcessComplete := "{{SLAPercentage}}"
	//ProcessComplete := strconv.Itoa(123)
	if strings.Contains(emailSubject, slaTemplateVarForProcessComplete) {
		emailSubject = strings.ReplaceAll(emailSubject, slaTemplateVarForProcessComplete, slaPercentageString)
	}
	if strings.Contains(emailBody, slaTemplateVarForProcessComplete) {
		emailBody = strings.ReplaceAll(emailBody, slaTemplateVarForProcessComplete, slaPercentageString)
	}
	Logger.Log.Println("Subject===> ", emailSubject)
	Logger.Log.Println("Body====> ", emailBody)
	Logger.Log.Println("Notification Template id", notificationTemplateID)
	//var userTOEmails string
	//var userCCEmails string
	// var creatorEmail string
	// var originalCreatorEmail string
	// var assigneeEmail string
	// var assigneeSupportGrpEmail string
	// var assigneeSupportGrpMemberEmail string
	var creatorID int64
	var originalCreatorID int64
	//var groupID int64
	getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
	creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
	if creatorIDresulsetErr != nil {
		Logger.Log.Println(creatorIDresulsetErr)
		return errors.New("Unable To fetch CreatorID")
	}
	userTOEmails, userCCEmails, getEmailIDError := GetEmailIDs(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember, notificationTemplateID)
	if getEmailIDError != nil {
		Logger.Log.Println(getEmailIDError)
		return errors.New("Unable To fetch CreatorID")
	}

	sendEmailError := SendEMail(db, emailSubject, emailBody, userTOEmails, userCCEmails, clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
	if sendEmailError != nil {
		Logger.Log.Println(sendEmailError)
		return sendEmailError
	}

	return nil
}
func MailBodyFormationFromTemplate(requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	channelType := int64(requestData["channeltype"].(float64))
	//termSeq := requestData["termseq"]

	if channelType != 1 {
		return errors.New("Wrong Channel Type")
	}
	//var templateVariable []string
	//var dynamicQuery []string
	//var dynamicQueryParam []string

	// mutex.Lock()
	// defer mutex.Unlock()
	db, dBerr := config.GetDB()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		return dBerr
	}
	//db.SetConnMaxIdleTime()
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	//defer db.Close()
	notificationFlag, notificationFlagErr := dao.GetNotificationFlag(db, clientID, orgID)
	if notificationFlagErr != nil {
		Logger.Log.Println(notificationFlagErr)
		return errors.New("Notification Flag Error")
	}
	if notificationFlag != 1 {
		return errors.New("Notification is Disabled")
	}

	var workingCategoryID int64
	getWorkingCategoryIDQuery := "select recorddiffid from maprecordtorecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recordid=? and isworking=1 and islatest=1"
	Logger.Log.Println("Record ID ===>", recordID)
	Logger.Log.Println(getWorkingCategoryIDQuery)
	Logger.Log.Println("clientID", clientID)
	Logger.Log.Println("Orgid", orgID)
	WorkingCategoryErr := db.QueryRow(getWorkingCategoryIDQuery, clientID, orgID, recordID).Scan(&workingCategoryID)
	if WorkingCategoryErr != nil {
		Logger.Log.Println(WorkingCategoryErr)
		return errors.New("Unable To fetch Working Category ID")
	}

	switch eventNotificationID {
	case 1:
		statusChangeEventErr := StatusChangeEvent(db, workingCategoryID, requestData)
		if statusChangeEventErr != nil {
			Logger.Log.Println(statusChangeEventErr)
			return errors.New("Error in sending mail for Status change")
		}
		break

	case 2:
		customerVisibleWorkNoteError := CustomerVisibleWorkNoteEvent(db, workingCategoryID, requestData)
		if customerVisibleWorkNoteError != nil {
			Logger.Log.Println(customerVisibleWorkNoteError)
			return errors.New("Error in sending mail for Customer Visible WorkNote")
		}
		break
	case 3:
		attachmentsAddedError := CustomerVisibleWorkNoteEvent(db, workingCategoryID, requestData)
		if attachmentsAddedError != nil {
			Logger.Log.Println(attachmentsAddedError)
			return errors.New("Error in sending mail  For Attachments Added")
		}
		break
	case 4:
		attachmentsDeletedError := CustomerVisibleWorkNoteEvent(db, workingCategoryID, requestData)
		if attachmentsDeletedError != nil {
			Logger.Log.Println(attachmentsDeletedError)
			return errors.New("Error in sending mail For Attachments Delete")
		}
		break

	case 5:
		userFollowUpCountError := UserFollowUpCountEvent(db, workingCategoryID, requestData)
		if userFollowUpCountError != nil {
			Logger.Log.Println(userFollowUpCountError)
			return errors.New("Error in sending mail in User Follow Up Count")
		}
		break
	case 6:
		hopCountError := UserFollowUpCountEvent(db, workingCategoryID, requestData)
		if hopCountError != nil {
			Logger.Log.Println(hopCountError)
			return errors.New("Error in sending mail for Hop Count")
		}
		break
	case 7:
		priorityChangeError := PriorityChangeEvent(db, workingCategoryID, requestData)
		if priorityChangeError != nil {
			Logger.Log.Println(priorityChangeError)
			return errors.New("Error in sending mail for PrioritypriorityChangeErrorChange")
		}
		break
	case 8:
		SLAResponseError := SLAResponseEvent(db, workingCategoryID, requestData)
		if SLAResponseError != nil {
			Logger.Log.Println(SLAResponseError)
			return errors.New("Error in sending mail for SLA - Response")
		}
		break
	case 9:
		SLAResolutionError := SLAResponseEvent(db, workingCategoryID, requestData)
		if SLAResolutionError != nil {
			Logger.Log.Println(SLAResolutionError)
			return errors.New("Error in sending mail for SLA - Resolution")
		}
		break
	case 10:
		forwardedToOwnGrpError := CustomerVisibleWorkNoteEvent(db, workingCategoryID, requestData)
		if forwardedToOwnGrpError != nil {
			Logger.Log.Println(forwardedToOwnGrpError)
			return errors.New("Error in sending mail for Forwared To Own Group")
		}
		break
	case 11:
		forwardedToDiffGrpError := CustomerVisibleWorkNoteEvent(db, workingCategoryID, requestData)
		if forwardedToDiffGrpError != nil {
			Logger.Log.Println(forwardedToDiffGrpError)
			return errors.New("Error in sending mail for Forwared To Different Group")
		}
		break
	default:
		break
	}

	return nil
}
