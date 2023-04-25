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

func GetMobileNumber(creatorMobileNo string, originalcreatorMobileNo string, assigneeMobileNo string,
	assigneeSupportobileNo string, assigneeSupportGrpMembersMobileNo string) string {
	var usersMobileNo string
	if len(creatorMobileNo) > 0 {
		usersMobileNo = usersMobileNo + ",91" + creatorMobileNo
	}
	if len(originalcreatorMobileNo) > 0 {
		usersMobileNo = usersMobileNo + ",91" + originalcreatorMobileNo
	}
	if len(assigneeMobileNo) > 0 {
		usersMobileNo = usersMobileNo + ",91" + assigneeMobileNo

	}
	if len(assigneeSupportobileNo) > 0 {
		usersMobileNo = usersMobileNo + "," + assigneeSupportobileNo
	}
	if len(assigneeSupportGrpMembersMobileNo) > 0 {
		usersMobileNo = usersMobileNo + "," + assigneeSupportGrpMembersMobileNo
	}
	// Logger.Log.Println("creatorEmail====>", creatorEmail)
	// Logger.Log.Println("originalCreatorEmail====>", originalCreatorEmail)
	// Logger.Log.Println("assigneeEmail====>", assigneeEmail)
	// Logger.Log.Println("assigneeSupportGrpEmail====>", assigneeSupportGrpEmail)
	// Logger.Log.Println("assigneeSupportGrpMemberEmail====>", assigneeSupportGrpMemberEmail)
	// Logger.Log.Println("userTOEmails====>", userTOEmails)

	return usersMobileNo
}
func GetSMSNumbers(db *sql.DB, clientID int64, orgID int64, recordID int64,
	additionalRecipient string, creatorID int64, originalCreatorID int64,
	sendToCreator int64, sendToOriginalCreator int64, sendToAssignee int64,
	sendToAssigneeGroup int64, sendToAssigneeGroupMember int64) (string, error) {

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

	var usersMobileNo string
	//var userCCEmails string
	var creatorMobileNo string
	var originalcreatorMobileNo string
	var assigneeMobileNo string
	var assigneeSupportobileNo string
	var assigneeSupportGrpMembersMobileNo string
	var groupID int64
	//111
	if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()
		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//originalcreatorMobileNo = creatorMobileNo
		//110
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//101
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()
		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//100
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//011
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()
		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		} //
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()
		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//000

	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//111 for originalcreater!=1
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//110
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()
		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//100

	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//011
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//010
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		} //
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()
		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//000
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//111
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//110
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//	return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//101
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//100
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//011
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//010
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		} //
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//001
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//000
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//111
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {

		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		//originalcreatorMobileNo = creatorMobileNo
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//110
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && creatorID == originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {

		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//101
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {

		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//100
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//011
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {

		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//010
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		} //
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//001
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {

		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//000
	} else if sendToCreator != 1 && sendToOriginalCreator != 1 && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		// creator not sender
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//110
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//101
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//100
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//011
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//010
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		} //

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//000
	} else if sendToCreator == 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//111 for originalcreater!=1
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//110
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//101
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//100
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//011
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//010
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		} //
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//000
	} else if sendToCreator == 1 && sendToOriginalCreator != 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//111
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//	return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//110
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)
		//101
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//100
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee == 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeMobileNo := "select usermobileno from mstclientuser where id in(select mstuserid from mstrequest where id in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and activeflag=1 and deleteflag=0"
		getassigneeMobileNoResultseterr := db.QueryRow(getAssigneeMobileNo, clientID, orgID, recordID).Scan(&assigneeMobileNo)
		if getassigneeMobileNoResultseterr != nil {
			Logger.Log.Println(getassigneeMobileNoResultseterr)
			//return userTOEmails, userCCEmails, errors.New("ERROR: Unable to fetch EmailAddress of Resolver")
		}

		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//011
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//010
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID == originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup == 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getAssigneeSupportGrpMobileNo := "select groupmobileno,grp from mstclientsupportgroup where grpid in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 "
		getAssigneeSupportGrpMobileNoResultseterr := db.QueryRow(getAssigneeSupportGrpMobileNo, clientID, orgID, recordID).Scan(&assigneeSupportobileNo, &groupID)
		if getAssigneeSupportGrpMobileNoResultseterr != nil {
			Logger.Log.Println(getAssigneeSupportGrpMobileNoResultseterr)
			//return  usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Assignee Supportgroup")
		} //
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//001
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember == 1 {
		getMobileNoOriginalCreatorquery := " select usermobileno from mstclientuser where  id in (select userid from mstgroupmember where groupid  in(select id from mstclientsupportgroup where clientid=? and mstorgnhirarchyid=? and supportgrouplevelid=1) and userid = ?)"
		getMobileNoOriginalCreatorResultseterr := db.QueryRow(getMobileNoOriginalCreatorquery, clientID, orgID, originalCreatorID).Scan(&originalcreatorMobileNo)
		if getMobileNoOriginalCreatorResultseterr != nil {
			Logger.Log.Println(getMobileNoOriginalCreatorResultseterr)
			//	return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		getUserMobileNoGrpQuery := "select usermobileno from mstclientuser where id in (select userid from mstgroupmember where groupid in ( select id from mstclientsupportgroup where id in( select mstgroupid from mstrequest where id  in ( select mstrequestid from maprequestorecord where clientid=? and mstorgnhirarchyid=? and recordid=? and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1) and deleteflg=0 and activeflg=1 )) and activeflag=1 and deleteflag=0"
		userMobileNoResultset, userMobileNoResultsetError := db.Query(getUserMobileNoGrpQuery, clientID, orgID, recordID)
		if userMobileNoResultsetError != nil {
			Logger.Log.Println(userMobileNoResultsetError)
			//return errors.New("ERROR: Unable to fetch EmailAddress of group")
		}
		defer userMobileNoResultset.Close()

		for userMobileNoResultset.Next() {
			var userMobileNoForAgroupMemer string
			scanErr := userMobileNoResultset.Scan(&userMobileNoForAgroupMemer)
			if scanErr != nil {
				Logger.Log.Println(scanErr)
				//return errors.New("ERROR: Unable to fetch EmailAddress of group")
			}
			assigneeSupportGrpMembersMobileNo = assigneeSupportGrpMembersMobileNo + "," + userMobileNoForAgroupMemer
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

		//000
	} else if sendToCreator != 1 && sendToOriginalCreator == 1 && creatorID != originalCreatorID && sendToAssignee != 1 && sendToAssigneeGroup != 1 && sendToAssigneeGroupMember != 1 {
		getMobileNoquery := "select usermobileno from mstclientuser where id=? and activeflag=1 and deleteflag=0"
		getMobilrNoResultseterr := db.QueryRow(getMobileNoquery, creatorID).Scan(&creatorMobileNo)
		if getMobilrNoResultseterr != nil {
			Logger.Log.Println(getMobilrNoResultseterr)
			//return usersMobileNo, errors.New("ERROR: Unable to fetch MobileNo of Creator")
		}
		usersMobileNo = GetMobileNumber(creatorMobileNo, originalcreatorMobileNo, assigneeMobileNo, assigneeSupportobileNo, assigneeSupportGrpMembersMobileNo)

	}

	// Logger.Log.Println("userCCEmails", userCCEmails)
	// Logger.Log.Println("Additional Receipint in getEmails", additionalRecipient)
	var tempAdditionalNo string
	if !strings.EqualFold(additionalRecipient, "") {
		additionalRecipientList := strings.Split(additionalRecipient, ",")
		for i := 0; i < len(additionalRecipientList); i++ {
			additionalRecipientList[i] = "91" + additionalRecipientList[i]
			tempAdditionalNo = tempAdditionalNo + "," + additionalRecipientList[i]
		}
	}
	tempAdditionalNo = strings.Trim(tempAdditionalNo, ",")
	toNumber := strings.Split(usersMobileNo, ",")
	//usersMobileNo = strings.Split(usersMobileNo, ",")
	tempToNo := ""
	for i := 0; i < len(toNumber); i++ {
		if len(toNumber[i]) > 0 {
			tempToNo = tempToNo + "," + toNumber[i]
		}
	}
	tempToNo = strings.Trim(tempToNo, ",")
	usersMobileNo = tempToNo + "," + tempAdditionalNo
	// for i:=0 ;range toUserEmails
	Logger.Log.Println("UsersMObileNumber=======>", usersMobileNo)
	Logger.Log.Println("<========================ExitingFrom Get MobileNo===================>")
	return usersMobileNo, nil
}
func SendSMS(db *sql.DB, smsContent string, usersMobileNo string, smsTemplateID string, smsType string,
	clientID int64, orgID int64, recordID int64, eventNotificationID int64, notificationTemplateID int64) error {
	usersMobileNo = strings.Trim(usersMobileNo, ",")

	if !strings.EqualFold(smsType, "1") {
		Logger.Log.Println("Invalid smsType==> ", smsType)
		return errors.New("Invalid smsType")
	}
	Logger.Log.Println("\n\n=============================================IN Send SMS Service=======================>\n\n")
	Logger.Log.Println("eventNotificationID===> ", eventNotificationID)
	Logger.Log.Println("notificationTemplateID==> ", notificationTemplateID)
	Logger.Log.Println("Body====> ", smsContent)
	Logger.Log.Println("smsTemplateID==> ", smsTemplateID)
	Logger.Log.Println("smsType==> ", smsType)

	SMSUrl, IndiaDltPrincipalEntityIdVal, smsUserName, smsPassword, smsError := dao.GetSMSConfiGDetails(db, clientID, orgID)
	if smsError != nil {
		Logger.Log.Println(smsError)
		return errors.New("ERROR: SMS Configuration Not Found!!!")
	}
	sendSMSErr := SendMailUtils.SendSMS(smsTemplateID, usersMobileNo, smsContent, smsType, SMSUrl, IndiaDltPrincipalEntityIdVal, smsUserName, smsPassword)
	if sendSMSErr != nil {
		Logger.Log.Println(sendSMSErr)
		notificationLogQuery := "INSERT INTO `mstnotificationlog`(`clientid`,`mstorgnhirarchyid`,`recordid`,`notificationeventid`,`notificationtemplateid`,`processflag`,`deleteflg`,`activeflg`)VALUES(?,?,?,?,?,?,?,?)"
		Resultset, err := db.Query(notificationLogQuery, clientID, orgID, recordID, eventNotificationID, notificationTemplateID, "N", 0, 1)
		if err != nil {
			Logger.Log.Println(err)
			return err
		}
		Resultset.Close()
		return sendSMSErr
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

//func SMSFormationFromTemplate
func StatusChangeEventForSMS(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	channelType := int64(requestData["channeltype"].(float64))
	//termSeq := int64(requestData["termseq"].(float64))
	statusSeq := int64(requestData["statusid"].(float64))

	if channelType != 2 {
		return errors.New("Wrong Channel Type")
	}
	var notificationTemplateID int64
	//var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	var indiaDltContentTemplateId string
	var smsType string
	//var eventParams interface{}

	var eventParam entities.EventParam
	eventParam.StatusSeq = statusSeq
	eventParamBytesArray, jsonMarshallErr := json.Marshal(eventParam)
	if jsonMarshallErr != nil {
		Logger.Log.Println(jsonMarshallErr)
		return errors.New("Error in status Change Status Parsing")
	}
	log.Println("working cat===>", workingCategoryID)
	eventParamstring := string(eventParamBytesArray)
	Logger.Log.Println("eventParamstring===> ", eventParamstring)
	getSubjectAndBodyQury := "select id,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember,smstemplateid,smstype from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=? and eventparams=? and activeflg=1 and deleteflg=0"
	getSubjectAndBodyErr := db.QueryRow(getSubjectAndBodyQury, clientID, orgID, eventNotificationID, channelType, workingCategoryID, eventParamstring).Scan(
		&notificationTemplateID, &body, &additionalRecipient, &sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember, &indiaDltContentTemplateId, &smsType)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	smsContent, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForSMS(db, body, requestData)
	if templateVariableMappingUsingDynamicQueriesError != nil {
		Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
		return errors.New("Workin Category is Not mapped properly")
	}
	//Logger.Log.Println("Subject===> ", emailSubject)
	Logger.Log.Println("smsContent====> ", smsContent)
	Logger.Log.Println("Notification Template id===>", notificationTemplateID)

	var creatorID int64
	var originalCreatorID int64
	//var groupID int64
	getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
	creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
	if creatorIDresulsetErr != nil {
		Logger.Log.Println(creatorIDresulsetErr)
		return errors.New("Unable To fetch CreatorID")
	}
	getMobileNumbers, getgetMobileNumberError := GetSMSNumbers(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember)
	if getgetMobileNumberError != nil {
		Logger.Log.Println(getgetMobileNumberError)
		return errors.New("Unable To fetch Mobile Numbers")
	}
	sendSMSErr := SendSMS(db, smsContent, getMobileNumbers, indiaDltContentTemplateId, smsType,
		clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
	if sendSMSErr != nil {
		Logger.Log.Println(sendSMSErr)
		return errors.New("Unable to Send SMS")
	}
	return nil
}
func CustomerVisibleWorkNoteEventForSMS(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
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
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	var indiaDltContentTemplateId string
	var smsType string
	getSubjectAndBodyQury := "select id,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember,smstemplateid,smstype from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=? and activeflg=1 and deleteflg=0"
	getSubjectAndBodyErr := db.QueryRow(getSubjectAndBodyQury, clientID, orgID, eventNotificationID, channelType, workingCategoryID).Scan(&notificationTemplateID, &body, &additionalRecipient, &sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember, &indiaDltContentTemplateId, &smsType)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	log.Println("ID===>", notificationTemplateID)
	log.Println("smstemplateid,==>", indiaDltContentTemplateId)
	log.Println("body", body)
	log.Println("additional ===>", additionalRecipient)
	smsContent, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForSMS(db, body, requestData)
	if templateVariableMappingUsingDynamicQueriesError != nil {
		Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
		return errors.New("Workin Category is Not mapped properly")
	}
	//Logger.Log.Println("Subject===> ", emailSubject)
	Logger.Log.Println("smsContent====> ", smsContent)
	Logger.Log.Println("Notification Template id===>", notificationTemplateID)

	var creatorID int64
	var originalCreatorID int64
	//var groupID int64
	getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
	creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
	if creatorIDresulsetErr != nil {
		Logger.Log.Println(creatorIDresulsetErr)
		return errors.New("Unable To fetch CreatorID")
	}
	getMobileNumbers, getgetMobileNumberError := GetSMSNumbers(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember)
	if getgetMobileNumberError != nil {
		Logger.Log.Println(getgetMobileNumberError)
		return errors.New("Unable To fetch Mobile Numbers")
	}
	sendSMSErr := SendSMS(db, smsContent, getMobileNumbers, indiaDltContentTemplateId, smsType,
		clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
	if sendSMSErr != nil {
		Logger.Log.Println(sendSMSErr)
		return errors.New("Unable to Send SMS")
	}
	return nil
}
func UserFollowUpCountEventForSMS(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	channelType := int64(requestData["channeltype"].(float64))
	//termSeq := int64(requestData["termseq"].(float64))
	//	statusSeq := int64(requestData["statusid"].(float64))

	if channelType != 2 {
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
	//var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	var indiaDltContentTemplateId string
	var smsType string
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
	getSubjectAndBodyQury := "select id,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember,eventparams,smstemplateid,smstype from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=?  and activeflg=1 and deleteflg=0"
	Resultset, getSubjectAndBodyErr := db.Query(getSubjectAndBodyQury, clientID, orgID, eventNotificationID, channelType, workingCategoryID)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	defer Resultset.Close()
	for Resultset.Next() {
		scanErr := Resultset.Scan(&notificationTemplateID, &body, &additionalRecipient,
			&sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember,
			&eventParam, &indiaDltContentTemplateId, &smsType)
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
		log.Println("smstemplateid,==>", indiaDltContentTemplateId)
		log.Println("body", body)
		log.Println("additional ===>", additionalRecipient)
		smsContent, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForSMS(db, body, requestData)
		if templateVariableMappingUsingDynamicQueriesError != nil {
			Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
			return errors.New("Workin Category is Not mapped properly")
		}
		//Logger.Log.Println("Subject===> ", emailSubject)
		Logger.Log.Println("smsContent====> ", smsContent)
		Logger.Log.Println("Notification Template id===>", notificationTemplateID)

		if eventNotificationID == 6 {
			Logger.Log.Println("In Hop Count Value Replacement")

			hopcountValue := strconv.Itoa(int(countVal))
			Logger.Log.Println("Hop Count Value=====>", hopcountValue)
			hopCount := "{{HopCount}}"
			lastToLastAssignmentGrpVar := "{{LastAssignmentTower}}"

			if strings.Contains(smsContent, hopCount) {
				smsContent = strings.ReplaceAll(smsContent, hopCount, hopcountValue)
			}

			if strings.Contains(smsContent, lastToLastAssignmentGrpVar) {
				smsContent = strings.ReplaceAll(smsContent, lastToLastAssignmentGrpVar, lastTolastAssignmentGrp)
			}

		}

		Logger.Log.Println("smsContent====> ", smsContent)
		Logger.Log.Println("Notification Template id===>", notificationTemplateID)

		var creatorID int64
		var originalCreatorID int64
		//var groupID int64
		getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
		creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
		if creatorIDresulsetErr != nil {
			Logger.Log.Println(creatorIDresulsetErr)
			return errors.New("Unable To fetch CreatorID")
		}
		getMobileNumbers, getgetMobileNumberError := GetSMSNumbers(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember)
		if getgetMobileNumberError != nil {
			Logger.Log.Println(getgetMobileNumberError)
			return errors.New("Unable To fetch Mobile Numbers")
		}
		sendSMSErr := SendSMS(db, smsContent, getMobileNumbers, indiaDltContentTemplateId, smsType,
			clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
		if sendSMSErr != nil {
			Logger.Log.Println(sendSMSErr)
			return errors.New("Unable to Send SMS")
		}
	} else {
		Logger.Log.Println("SMS not sent for noofcount as criteria not satisfies")
	}

	return nil
}
func PriorityChangeEventForSMS(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
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
	//var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	var indiaDltContentTemplateId string
	var smsType string
	getSubjectAndBodyQury := "select id,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember,smstemplateid,smstype from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=? and eventparams=? and activeflg=1 and deleteflg=0"
	getSubjectAndBodyErr := db.QueryRow(getSubjectAndBodyQury, clientID, orgID, eventNotificationID, channelType, workingCategoryID, eventParamstring).Scan(
		&notificationTemplateID, &body, &additionalRecipient, &sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember, &indiaDltContentTemplateId, &smsType)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	smsContent, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForSMS(db, body, requestData)
	if templateVariableMappingUsingDynamicQueriesError != nil {
		Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
		return errors.New("Workin Category is Not mapped properly")
	}
	//Logger.Log.Println("Subject===> ", emailSubject)
	log.Println("ID===>", notificationTemplateID)
	log.Println("smstemplateid,==>", indiaDltContentTemplateId)
	log.Println("body", body)
	log.Println("additional ===>", additionalRecipient)
	Logger.Log.Println("smsContent====> ", smsContent)
	Logger.Log.Println("Notification Template id===>", notificationTemplateID)

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

	if strings.Contains(smsContent, additionalInfoVar) {
		smsContent = strings.ReplaceAll(smsContent, additionalInfoVar, AdditionalInfo)
	}

	var creatorID int64
	var originalCreatorID int64
	//var groupID int64
	getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
	creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
	if creatorIDresulsetErr != nil {
		Logger.Log.Println(creatorIDresulsetErr)
		return errors.New("Unable To fetch CreatorID")
	}
	getMobileNumbers, getgetMobileNumberError := GetSMSNumbers(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember)
	if getgetMobileNumberError != nil {
		Logger.Log.Println(getgetMobileNumberError)
		return errors.New("Unable To fetch Mobile Numbers")
	}
	sendSMSErr := SendSMS(db, smsContent, getMobileNumbers, indiaDltContentTemplateId, smsType,
		clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
	if sendSMSErr != nil {
		Logger.Log.Println(sendSMSErr)
		return errors.New("Unable to Send SMS")
	}
	return nil
}
func SLAResponseEventForSMS(db *sql.DB, workingCategoryID int64, requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	//channelType := int64(requestData["channeltype"].(float64))
	notificationTemplateID := int64(requestData["notificationtemplateid"].(float64))
	slaPercentage := requestData["percentage"].(float64)
	slaPercentageString := fmt.Sprintf("%f", slaPercentage)

	//var notificationTemplateID int64
	//var subjectTitle string
	var body string
	var additionalRecipient string
	var sendToOriginalCreator int64
	var sendToCreator int64
	var sendToAssignee int64
	var sendToAssigneeGroup int64
	var sendToAssigneeGroupMember int64
	var indiaDltContentTemplateId string
	var smsType string
	var eventParam string
	var eventParamObj entities.EventParam

	getSubjectAndBodyQury := "select id,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember,eventparams,smstemplateid,smstype from mstnotificationtemplate where clientid=? and mstorgnhirarchyid=? and eventtype=? and channeltype=? and workingcategoryid=?  and activeflg=1 and deleteflg=0"
	getSubjectAndBodyErr := db.QueryRow(getSubjectAndBodyQury, clientID, orgID, notificationTemplateID).Scan(&notificationTemplateID, &body, &additionalRecipient,
		&sendToCreator, &sendToOriginalCreator, &sendToAssignee, &sendToAssigneeGroup, &sendToAssigneeGroupMember,
		&eventParam, &indiaDltContentTemplateId, &smsType)
	if getSubjectAndBodyErr != nil {
		Logger.Log.Println(getSubjectAndBodyErr)
		//errors.New("Workin Category is Not mapped properly")
		return errors.New("Workin Category is Not mapped properly")
	}
	unmarshalError := json.Unmarshal([]byte(eventParam), &eventParamObj)
	if unmarshalError != nil {
		Logger.Log.Println(unmarshalError)
		return errors.New("ERROR: Unable to unmarshal event param")
	}
	log.Println("ID===>", notificationTemplateID)
	log.Println("smstemplateid,==>", indiaDltContentTemplateId)
	log.Println("body", body)
	log.Println("additional ===>", additionalRecipient)
	smsContent, templateVariableMappingUsingDynamicQueriesError := dao.TemplateVariableMappingUsingDynamicQueriesForSMS(db, body, requestData)
	if templateVariableMappingUsingDynamicQueriesError != nil {
		Logger.Log.Println(templateVariableMappingUsingDynamicQueriesError)
		return errors.New("Workin Category is Not mapped properly")
	}
	slaTemplateVarForProcessComplete := "{{SLAPercentage}}"
	//ProcessComplete := strconv.Itoa(123)
	if strings.Contains(smsContent, slaTemplateVarForProcessComplete) {
		smsContent = strings.ReplaceAll(smsContent, slaTemplateVarForProcessComplete, slaPercentageString)
	}

	var creatorID int64
	var originalCreatorID int64
	//var groupID int64
	getCreatorIDQuery := "select userid,originaluserid from trnrecord where clientid=? and mstorgnhirarchyid=? and id=?"
	creatorIDresulsetErr := db.QueryRow(getCreatorIDQuery, clientID, orgID, recordID).Scan(&creatorID, &originalCreatorID)
	if creatorIDresulsetErr != nil {
		Logger.Log.Println(creatorIDresulsetErr)
		return errors.New("Unable To fetch CreatorID")
	}
	getMobileNumbers, getgetMobileNumberError := GetSMSNumbers(db, clientID, orgID, recordID, additionalRecipient, creatorID, originalCreatorID, sendToCreator, sendToOriginalCreator, sendToAssignee, sendToAssigneeGroup, sendToAssigneeGroupMember)
	if getgetMobileNumberError != nil {
		Logger.Log.Println(getgetMobileNumberError)
		return errors.New("Unable To fetch Mobile Numbers")
	}
	sendSMSErr := SendSMS(db, smsContent, getMobileNumbers, indiaDltContentTemplateId, smsType,
		clientID, orgID, recordID, eventNotificationID, notificationTemplateID)
	if sendSMSErr != nil {
		Logger.Log.Println(sendSMSErr)
		return errors.New("Unable to Send SMS")
	}
	return nil
}
func SMSFormationFromTemplate(requestData map[string]interface{}) error {
	clientID := int64(requestData["clientid"].(float64))
	orgID := int64(requestData["mstorgnhirarchyid"].(float64))
	recordID := int64(requestData["recordid"].(float64))
	eventNotificationID := int64(requestData["eventnotificationid"].(float64))
	channelType := int64(requestData["channeltype"].(float64))
	//termSeq := requestData["termseq"]

	if channelType != 2 {
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
	getWorkingCategoryIDQuery := "select recorddiffid from maprecordtorecorddifferentiation where clientid=? and mstorgnhirarchyid=? and recordid=? and isworking=1"
	resulsetErr := db.QueryRow(getWorkingCategoryIDQuery, clientID, orgID, recordID).Scan(&workingCategoryID)
	if resulsetErr != nil {
		Logger.Log.Println(resulsetErr)
		return errors.New("Unable To fetch Working Category ID")
	}

	switch eventNotificationID {
	case 1:
		statusChangeEventErr := StatusChangeEventForSMS(db, workingCategoryID, requestData)
		if statusChangeEventErr != nil {
			Logger.Log.Println(statusChangeEventErr)
			return errors.New("Error in sending SMS for Status change")
		}
		break

	case 2:
		customerVisibleWorkNoteError := CustomerVisibleWorkNoteEventForSMS(db, workingCategoryID, requestData)
		if customerVisibleWorkNoteError != nil {
			Logger.Log.Println(customerVisibleWorkNoteError)
			return errors.New("Error in sending SMS for Customer Visible WorkNote")
		}
		break
	case 3:
		attachmentsAddedError := CustomerVisibleWorkNoteEventForSMS(db, workingCategoryID, requestData)
		if attachmentsAddedError != nil {
			Logger.Log.Println(attachmentsAddedError)
			return errors.New("Error in sending SMS  For Attachments Added")
		}
		break
	case 4:
		attachmentsDeletedError := CustomerVisibleWorkNoteEventForSMS(db, workingCategoryID, requestData)
		if attachmentsDeletedError != nil {
			Logger.Log.Println(attachmentsDeletedError)
			return errors.New("Error in sending SMS For Attachments Delete")
		}
		break

	case 5:
		userFollowUpCountError := UserFollowUpCountEventForSMS(db, workingCategoryID, requestData)
		if userFollowUpCountError != nil {
			Logger.Log.Println(userFollowUpCountError)
			return errors.New("Error in sending SMS in User Follow Up Count")
		}
		break
	case 6:
		hopCountError := UserFollowUpCountEventForSMS(db, workingCategoryID, requestData)
		if hopCountError != nil {
			Logger.Log.Println(hopCountError)
			return errors.New("Error in sending SMS for Hop Count")
		}
		break
	case 7:
		priorityChangeError := PriorityChangeEventForSMS(db, workingCategoryID, requestData)
		if priorityChangeError != nil {
			Logger.Log.Println(priorityChangeError)
			return errors.New("Error in sending SMS for PrioritypriorityChangeErrorChange")
		}
		break
	case 8:
		SLAResponseError := SLAResponseEventForSMS(db, workingCategoryID, requestData)
		if SLAResponseError != nil {
			Logger.Log.Println(SLAResponseError)
			return errors.New("Error in sending SMS for SLA - Response")
		}
		break
	case 9:
		SLAResolutionError := SLAResponseEventForSMS(db, workingCategoryID, requestData)
		if SLAResolutionError != nil {
			Logger.Log.Println(SLAResolutionError)
			return errors.New("Error in sending SMS for SLA - Resolution")
		}
		break
	case 10:
		forwardedToOwnGrpError := CustomerVisibleWorkNoteEventForSMS(db, workingCategoryID, requestData)
		if forwardedToOwnGrpError != nil {
			Logger.Log.Println(forwardedToOwnGrpError)
			return errors.New("Error in sending SMS for Forwared To Own Group")
		}
		break
	case 11:
		forwardedToDiffGrpError := CustomerVisibleWorkNoteEventForSMS(db, workingCategoryID, requestData)
		if forwardedToDiffGrpError != nil {
			Logger.Log.Println(forwardedToDiffGrpError)
			return errors.New("Error in sending SMS for Forwared To Different Group")
		}
		break
	default:
		break
	}

	return nil
}
