package dao

import (
	"database/sql"
	"encoding/json"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var insertMsttemplate = "INSERT INTO msttemplate(clientid, mstorgnhirarchyid, templatename, templatetype, templatecontent, autoclosuretime, activeflg, deleteflg) VALUES (?,?,?,?,?,?,1,0) "
var insetMsttemplatediff = "INSERT INTO maptemplatediff(clientid, mstorgnhirarchyid, templateid, recorddifftypeid, recorddiffid, activeflg, deleteflg) VALUES (?,?,?,?,?,1,0) "

var duplicateMsttemplate = "SELECT count(id) total FROM  msttemplate WHERE clientid = ? AND mstorgnhirarchyid = ? AND templatename=? AND templatetype = ? AND deleteflg = 0"

// var getMsttemplate = "SELECT msttemplate.id, msttemplate.clientid, msttemplate.mstorgnhirarchyid, msttemplate.templatename, msttemplate.templatetype, msttemplate.templatecontent, msttemplate.autoclosuretime, msttemplate.activeflg,mstclient.name clientname,mstorgnhierarchy.name orgname FROM msttemplate,mstclient,mstorgnhierarchy WHERE msttemplate.clientid=? AND msttemplate.mstorgnhirarchyid=?  AND msttemplate.activeflg=1 AND msttemplate.deleteflg=0 AND msttemplate.clientid=mstclient.id AND msttemplate.clientid=mstorgnhierarchy.clientid AND msttemplate.mstorgnhirarchyid=mstorgnhierarchy.id ORDER BY msttemplate.id DESC LIMIT ?,? "
var getMsttemplatediff = "SELECT maptemplatediff.id,maptemplatediff.clientid,maptemplatediff.mstorgnhirarchyid,maptemplatediff.templateid,maptemplatediff.recorddifftypeid, maptemplatediff.recorddiffid,maptemplatediff.activeflg,mstrecorddifferentiationtype.typename,mstrecorddifferentiation.name,mstrecorddifferentiationtype.parentid FROM maptemplatediff,mstrecorddifferentiationtype, mstrecorddifferentiation WHERE maptemplatediff.activeflg = 1 AND maptemplatediff.deleteflg = 0 AND mstrecorddifferentiationtype.deleteflg = 0 AND mstrecorddifferentiationtype.activeflg = 1 AND mstrecorddifferentiation.deleteflg = 0 AND mstrecorddifferentiation.activeflg=1 AND maptemplatediff.recorddifftypeid=mstrecorddifferentiationtype.id AND maptemplatediff.recorddifftypeid=mstrecorddifferentiation.recorddifftypeid AND  maptemplatediff.recorddiffid = mstrecorddifferentiation.id  AND maptemplatediff.clientid = ? AND maptemplatediff.mstorgnhirarchyid = ? AND maptemplatediff.templateid = ? "

// var getMsttemplatecount = "SELECT count(msttemplate.id) total FROM msttemplate,mstclient,mstorgnhierarchy WHERE msttemplate.clientid=? AND msttemplate.mstorgnhirarchyid=?  AND msttemplate.activeflg=1 AND msttemplate.deleteflg=0 AND msttemplate.clientid=mstclient.id AND msttemplate.clientid=mstorgnhierarchy.clientid AND msttemplate.mstorgnhirarchyid=mstorgnhierarchy.id"

var updateMsttemplate = "UPDATE msttemplate SET templatename = ?, templatetype = ?,templatecontent=?,autoclosuretime=?  WHERE id = ? "

var deleteMsttemplate = "UPDATE msttemplate SET deleteflg = '1',activeflg=0 WHERE id = ? "

var deleteMsttemplatediff = "UPDATE maptemplatediff SET deleteflg = '1',activeflg=0 WHERE templateid = ? "

func (dbc DbConn) CheckDuplicateMsttemplate(tz *entities.MsttemplateEntity) (entities.MsttemplateEntities, error) {
	logger.Log.Println("In side CheckDuplicateMsttemplate")
	value := entities.MsttemplateEntities{}
	err := dbc.DB.QueryRow(duplicateMsttemplate, tz.Clientid, tz.Mstorgnhirarchyid, tz.Templatename, tz.Templatetype).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMsttemplate Get Statement Prepare Error", err)
		return value, err
	}
}

//check duplicate template
func (dbc DbConn) CheckDuplicateNotificationTemplate(tz *entities.MstNotificationTemplateEntity) (int64, error) {
	logger.Log.Println("In side CheckDuplicateNotificationTemplate")
	var value int64
	if tz.ChannelTypeID == 1 {
		var query = "SELECT COUNT(id) Total FROM mstnotificationtemplate WHERE clientid=? AND mstorgnhirarchyid=? AND recordtypeid=? AND workingcategoryid=? AND channeltype=? AND subjectortitle=? AND eventtype=? AND eventparams=? AND converted=? AND activeflg=1 AND deleteflg=0"
		logger.Log.Println("Query: ", query)
		b, err := json.Marshal(tz.EventParams)
		if err != nil {
			logger.Log.Println("CheckDuplicateNotificationTemplate JSON Conversion Error", err)
			return 0, err
		}
		logger.Log.Println("Params: ", tz.ClientID, tz.MstOrgnHirarchyID, tz.RecordTypeID, tz.WorkingCategoryID, tz.ChannelTypeID, tz.SubjectOrTitle, tz.EventTypeID, string(b))
		err = dbc.DB.QueryRow(query, tz.ClientID, tz.MstOrgnHirarchyID, tz.RecordTypeID, tz.WorkingCategoryID, tz.ChannelTypeID, tz.SubjectOrTitle, tz.EventTypeID, string(b), tz.Isconverted).Scan(&value)
		switch err {
		case sql.ErrNoRows:
			value = 0
			return value, nil
		case nil:
			return value, nil
		default:
			logger.Log.Println("CheckDuplicateNotificationTemplate Get Statement Prepare Error", err)
			return value, err
		}
	} else {
		var query = "SELECT COUNT(id) Total FROM mstnotificationtemplate WHERE clientid=? AND mstorgnhirarchyid=? AND recordtypeid=? AND workingcategoryid=? AND channeltype=? AND body=? AND eventtype=? AND eventparams=? AND activeflg=1 AND deleteflg=0"
		logger.Log.Println("Query: ", query)
		b, err := json.Marshal(tz.EventParams)
		if err != nil {
			logger.Log.Println("CheckDuplicateNotificationTemplate JSON Conversion Error", err)
			return 0, err
		}
		logger.Log.Println("Params: ", tz.ClientID, tz.MstOrgnHirarchyID, tz.RecordTypeID, tz.WorkingCategoryID, tz.ChannelTypeID, tz.SubjectOrTitle, tz.EventTypeID, string(b))
		err = dbc.DB.QueryRow(query, tz.ClientID, tz.MstOrgnHirarchyID, tz.RecordTypeID, tz.WorkingCategoryID, tz.ChannelTypeID, tz.Body, tz.EventTypeID, string(b)).Scan(&value)
		switch err {
		case sql.ErrNoRows:
			value = 0
			return value, nil
		case nil:
			return value, nil
		default:
			logger.Log.Println("CheckDuplicateNotificationTemplate Get Statement Prepare Error", err)
			return value, err
		}
	}

}

//insert Notification Template for save data
func (dbc DbConn) InsertNotificationTemplate(tz *entities.MstNotificationTemplateEntity) (int64, error) {
	logger.Log.Println("In side InsertNotificationTemplate")
	var query string
	var lastInsertedId int64
	if tz.ChannelTypeID == 1 {
		query = "INSERT INTO mstnotificationtemplate(clientid, mstorgnhirarchyid, recordtypeid, workingcategoryid, channeltype, subjectortitle, body, additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember, eventtype, eventparams,converted,activeflg, deleteflg) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1,0)"
		logger.Log.Println("Query -->", query)
		stmt, err := dbc.DB.Prepare(query)

		if err != nil {
			logger.Log.Println("InsertNotificationTemplate Prepare Statement  Error", err)
			return 0, err
		}
		defer stmt.Close()
		b, err := json.Marshal(tz.EventParams)
		if err != nil {
			logger.Log.Println("InsertNotificationTemplate JSON Conversion Error", err)
			return 0, err
		}
		logger.Log.Println("Parameter -->", tz.ClientID, tz.MstOrgnHirarchyID, tz.RecordTypeID, tz.WorkingCategoryID, tz.ChannelTypeID, tz.SubjectOrTitle, tz.Body, tz.AdditionalRecipient, tz.SendToCreator, tz.SendToOrgCreator, tz.SendToAssignee, tz.SendToAssigneeGroup, tz.SendToAssigneeGroupMember, tz.EventTypeID, string(b), tz.Isconverted)
		res, err := stmt.Exec(tz.ClientID, tz.MstOrgnHirarchyID, tz.RecordTypeID, tz.WorkingCategoryID, tz.ChannelTypeID, tz.SubjectOrTitle, tz.Body, tz.AdditionalRecipient, tz.SendToCreator, tz.SendToOrgCreator, tz.SendToAssignee, tz.SendToAssigneeGroup, tz.SendToAssigneeGroupMember, tz.EventTypeID, string(b), tz.Isconverted)
		if err != nil {
			logger.Log.Println("InsertNotificationTemplate Execute Statement  Error", err)
			return 0, err
		}
		lastInsertedId, err = res.LastInsertId()
		if err != nil {
			logger.Log.Println("InsertNotificationTemplate Get Insert ID  Error", err)
			return 0, err
		}

	} else {
		log.Println("Channel type ID===>", tz.ChannelTypeID)
		query = "INSERT INTO mstnotificationtemplate(clientid, mstorgnhirarchyid, recordtypeid, workingcategoryid, channeltype,body,additionalrecipient,sendtocreator,sendtooriginalcreator,sendtoassignee,sendtoassigneegroup,sendtoassigneegroupmember, eventtype, eventparams,smstemplateid,smstype,converted,activeflg, deleteflg) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1,0)"
		logger.Log.Println("Query -->", query)
		stmt, err := dbc.DB.Prepare(query)

		if err != nil {
			logger.Log.Println("InsertNotificationTemplate Prepare Statement  Error", err)
			return 0, err
		}
		defer stmt.Close()
		b, err := json.Marshal(tz.EventParams)
		if err != nil {
			logger.Log.Println("InsertNotificationTemplate JSON Conversion Error", err)
			return 0, err
		}
		logger.Log.Println("Parameter -->", tz.ClientID, tz.MstOrgnHirarchyID, tz.RecordTypeID, tz.WorkingCategoryID, tz.ChannelTypeID, tz.Body, tz.AdditionalRecipient, tz.SendToCreator, tz.SendToOrgCreator, tz.SendToAssignee, tz.SendToAssigneeGroup, tz.SendToAssigneeGroupMember, tz.EventTypeID, string(b), tz.SmsTemplateID, tz.SmsTypeID, tz.Isconverted)
		res, err := stmt.Exec(tz.ClientID, tz.MstOrgnHirarchyID, tz.RecordTypeID, tz.WorkingCategoryID, tz.ChannelTypeID, tz.Body, tz.AdditionalRecipient, tz.SendToCreator, tz.SendToOrgCreator, tz.SendToAssignee, tz.SendToAssigneeGroup, tz.SendToAssigneeGroupMember, tz.EventTypeID, string(b), tz.SmsTemplateID, tz.SmsTypeID, tz.Isconverted)
		if err != nil {
			logger.Log.Println("InsertNotificationTemplate Execute Statement  Error", err)
			return 0, err
		}
		lastInsertedId, err = res.LastInsertId()
		if err != nil {
			logger.Log.Println("InsertNotificationTemplate Get Insert ID  Error", err)
			return 0, err
		}

	}

	return lastInsertedId, nil
}

func (dbc DbConn) InsertNotificationRecipient(tz *entities.MstNotificationRecipientEntity) (int64, error) {
	logger.Log.Println("In side InsertNotificationRecipient")
	var query = "INSERT INTO mstnotificationrecipients(notificationtemplateid, recipienttype, groupid, userid, activeflg, deleteflg) VALUES (?,?,?,?,1,0)"
	logger.Log.Println("Query -->", query)
	stmt, err := dbc.DB.Prepare(query)
	if err != nil {
		logger.Log.Println("InsertNotificationRecipient Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.NotificationTemplateID, tz.RecipientType, tz.GroupID, tz.UserID)
	res, err := stmt.Exec(tz.NotificationTemplateID, tz.RecipientType, tz.GroupID, tz.UserID)
	if err != nil {
		logger.Log.Println("InsertNotificationRecipient Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println("InsertNotificationRecipient Get Insert ID  Error", err)
		return 0, err
	}
	return lastInsertedId, nil
}
func (dbc DbConn) InsertMsttemplate(tz *entities.MsttemplateEntity) (int64, error) {
	logger.Log.Println("In side InsertMsttemplate")
	logger.Log.Println("Query -->", insertMsttemplate)
	stmt, err := dbc.DB.Prepare(insertMsttemplate)

	if err != nil {
		logger.Log.Println("InsertMsttemplate Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Templatename, tz.Templatetype, tz.Templatecontent, tz.Autoclosuretime)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Templatename, tz.Templatetype, tz.Templatecontent, tz.Autoclosuretime)
	if err != nil {
		logger.Log.Println("InsertMsttemplate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println("InsertMsttemplate Get Insert ID  Error", err)
		return 0, err
	}
	return lastInsertedId, nil
}

func (dbc DbConn) InsertMsttemplatediff(tz *entities.MsttemplateEntity, tz1 *entities.MaptemplatediffEntity, templateid int64) (int64, error) {
	logger.Log.Println("In side InsertMsttemplatediff")
	logger.Log.Println("Query -->", insetMsttemplatediff)
	stmt, err := dbc.DB.Prepare(insetMsttemplatediff)
	if err != nil {
		logger.Log.Println("InsertMsttemplatediff Prepare Statement  Error", err)
		return 0, err
	}
	defer stmt.Close()
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, templateid, tz1.Recorddifftypeid, tz1.Recorddiffid)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, templateid, tz1.Recorddifftypeid, tz1.Recorddiffid)
	if err != nil {
		logger.Log.Println("InsertMsttemplate Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println("InsertMsttemplate Get Insert ID  Error", err)
		return 0, err
	}
	return lastInsertedId, nil
}

func (dbc DbConn) GetNotificationEvents() ([]entities.MstNotificationEvent, error) {
	logger.Log.Println("In side GetNotificationEvents")
	values := []entities.MstNotificationEvent{}
	var query = "SELECT id, eventname FROM mstnotificationevents"
	rows, err := dbc.DB.Query(query)
	if err != nil {
		logger.Log.Println("GetNotificationEvents Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MstNotificationEvent{}
		rows.Scan(&value.ID, &value.Name)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllNotificationVariables(page *entities.MstNotificationTemplateEntity) ([]entities.MstNotificationVariable, error) {
	logger.Log.Println("In side GetAllNotificationVariables")
	values := []entities.MstNotificationVariable{}
	var query = "SELECT msttemplatevariable.id,msttemplatevariable.templatename name FROM msttemplatevariable WHERE msttemplatevariable.activeflg=1 AND msttemplatevariable.deleteflg=0 AND msttemplatevariable.clientid=? AND msttemplatevariable.mstorgnhirarchyid=?"
	rows, err := dbc.DB.Query(query, page.ClientID, page.MstOrgnHirarchyID)
	if err != nil {
		logger.Log.Println("GetAllNotificationVariables Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MstNotificationVariable{}
		rows.Scan(&value.ID, &value.Name)
		values = append(values, value)
	}
	return values, nil
}

//grid API
func (dbc DbConn) GetAllNotificationTemplates(page *entities.MstNotificationTemplateEntity, OrgnType int64) ([]entities.MstNotificationTemplateEntity, error) {
	logger.Log.Println("In side GetAllNotificationTemplates")
	values := []entities.MstNotificationTemplateEntity{}
	var query = "SELECT mstnotificationtemplate.id,mstnotificationtemplate.clientid,mstclient.name clientname, " +
		" mstnotificationtemplate.mstorgnhirarchyid,mstorgnhierarchy.name orgname,mstrecorddifferentiationtype.id recordtypedifftypeid," +
		" mstrecorddifferentiationtype.typename recordtypedifftype,mstnotificationtemplate.recordtypeid,mstrecorddifferentiation.name recordtype," +
		" difftype.id workingcategorytypeid,difftype.typename workingcategorytype,mstnotificationtemplate.workingcategoryid,workcatdiff.name workingcategory," +
		" mstnotificationtemplate.channeltype channeltypeid,IF(mstnotificationtemplate.channeltype=1,'Email','SMS') channeltype,coalesce( mstnotificationtemplate.subjectortitle,'') subjectortitle," +
		" coalesce( mstnotificationtemplate.body,'') body,coalesce( mstnotificationtemplate.additionalrecipient,'') additionalrecipient,mstnotificationtemplate.sendtocreator,mstnotificationtemplate.sendtooriginalcreator," +
		" mstnotificationtemplate.sendtoassignee,mstnotificationtemplate.sendtoassigneegroup,mstnotificationtemplate.sendtoassigneegroupmember," +
		" mstnotificationtemplate.eventtype eventtypeid,mstnotificationevents.eventname eventtype,mstnotificationtemplate.eventparams," +
		" coalesce( mstnotificationtemplate.smstemplateid,'') smstemplateid,coalesce(mstnotificationtemplate.smstype,0)smstypeid," +
		" IF(mstnotificationtemplate.smstype=1,'LongSMS','N/A') smstype,mstnotificationtemplate.converted,IF(mstnotificationtemplate.converted=1,'Converted','Normal') smstype,mstnotificationtemplate.activeflg FROM mstnotificationtemplate JOIN mstclient ON mstnotificationtemplate.clientid=mstclient.id JOIN mstorgnhierarchy ON" +
		" mstnotificationtemplate.clientid=mstorgnhierarchy.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecorddifferentiation ON " +
		" mstnotificationtemplate.recordtypeid=mstrecorddifferentiation.id AND mstnotificationtemplate.clientid=mstrecorddifferentiation.clientid AND" +
		" mstnotificationtemplate.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid  JOIN mstrecorddifferentiationtype ON" +
		" mstrecorddifferentiation.recorddifftypeid=mstrecorddifferentiationtype.id   JOIN mstrecorddifferentiation workcatdiff ON " +
		" mstnotificationtemplate.workingcategoryid=workcatdiff.id AND mstnotificationtemplate.clientid=workcatdiff.clientid AND" +
		" mstnotificationtemplate.mstorgnhirarchyid=workcatdiff.mstorgnhirarchyid  JOIN mstrecorddifferentiationtype difftype ON" +
		" workcatdiff.recorddifftypeid=difftype.id JOIN mstnotificationevents ON mstnotificationtemplate.eventtype=mstnotificationevents.id WHERE " +
		" mstnotificationtemplate.activeflg=1 AND mstnotificationtemplate.deleteflg=0 AND mstnotificationtemplate.clientid=? AND" +
		" mstnotificationtemplate.mstorgnhirarchyid=? order by mstnotificationtemplate.id desc LIMIT ?,?"

	// rows, err := dbc.DB.Query(query, page.ClientID, page.MstOrgnHirarchyID, page.Offset, page.Limit)

	var params []interface{}
	if OrgnType == 1 {
		query = "SELECT mstnotificationtemplate.id,mstnotificationtemplate.clientid,mstclient.name clientname, " +
			" mstnotificationtemplate.mstorgnhirarchyid,mstorgnhierarchy.name orgname,mstrecorddifferentiationtype.id recordtypedifftypeid," +
			" mstrecorddifferentiationtype.typename recordtypedifftype,mstnotificationtemplate.recordtypeid,mstrecorddifferentiation.name recordtype," +
			" difftype.id workingcategorytypeid,difftype.typename workingcategorytype,mstnotificationtemplate.workingcategoryid,workcatdiff.name workingcategory," +
			" mstnotificationtemplate.channeltype channeltypeid,IF(mstnotificationtemplate.channeltype=1,'Email','SMS') channeltype,coalesce( mstnotificationtemplate.subjectortitle,'') subjectortitle," +
			" coalesce( mstnotificationtemplate.body,'') body,coalesce( mstnotificationtemplate.additionalrecipient,'') additionalrecipient,mstnotificationtemplate.sendtocreator,mstnotificationtemplate.sendtooriginalcreator," +
			" mstnotificationtemplate.sendtoassignee,mstnotificationtemplate.sendtoassigneegroup,mstnotificationtemplate.sendtoassigneegroupmember," +
			" mstnotificationtemplate.eventtype eventtypeid,mstnotificationevents.eventname eventtype,mstnotificationtemplate.eventparams," +
			" coalesce( mstnotificationtemplate.smstemplateid,'') smstemplateid,coalesce(mstnotificationtemplate.smstype,0)smstypeid," +
			" IF(mstnotificationtemplate.smstype=1,'LongSMS','N/A') smstype,mstnotificationtemplate.converted,IF(mstnotificationtemplate.converted=1,'Converted','Normal') smstype,mstnotificationtemplate.activeflg FROM mstnotificationtemplate JOIN mstclient ON mstnotificationtemplate.clientid=mstclient.id JOIN mstorgnhierarchy ON" +
			" mstnotificationtemplate.clientid=mstorgnhierarchy.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecorddifferentiation ON " +
			" mstnotificationtemplate.recordtypeid=mstrecorddifferentiation.id AND mstnotificationtemplate.clientid=mstrecorddifferentiation.clientid AND" +
			" mstnotificationtemplate.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid  JOIN mstrecorddifferentiationtype ON" +
			" mstrecorddifferentiation.recorddifftypeid=mstrecorddifferentiationtype.id   JOIN mstrecorddifferentiation workcatdiff ON " +
			" mstnotificationtemplate.workingcategoryid=workcatdiff.id AND mstnotificationtemplate.clientid=workcatdiff.clientid AND" +
			" mstnotificationtemplate.mstorgnhirarchyid=workcatdiff.mstorgnhirarchyid  JOIN mstrecorddifferentiationtype difftype ON" +
			" workcatdiff.recorddifftypeid=difftype.id JOIN mstnotificationevents ON mstnotificationtemplate.eventtype=mstnotificationevents.id WHERE " +
			" mstnotificationtemplate.activeflg=1 AND mstnotificationtemplate.deleteflg=0 " +
			"  order by mstnotificationtemplate.id desc LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		query = "SELECT mstnotificationtemplate.id,mstnotificationtemplate.clientid,mstclient.name clientname, " +
			" mstnotificationtemplate.mstorgnhirarchyid,mstorgnhierarchy.name orgname,mstrecorddifferentiationtype.id recordtypedifftypeid," +
			" mstrecorddifferentiationtype.typename recordtypedifftype,mstnotificationtemplate.recordtypeid,mstrecorddifferentiation.name recordtype," +
			" difftype.id workingcategorytypeid,difftype.typename workingcategorytype,mstnotificationtemplate.workingcategoryid,workcatdiff.name workingcategory," +
			" mstnotificationtemplate.channeltype channeltypeid,IF(mstnotificationtemplate.channeltype=1,'Email','SMS') channeltype,coalesce( mstnotificationtemplate.subjectortitle,'') subjectortitle," +
			" coalesce( mstnotificationtemplate.body,'') body,coalesce( mstnotificationtemplate.additionalrecipient,'') additionalrecipient,mstnotificationtemplate.sendtocreator,mstnotificationtemplate.sendtooriginalcreator," +
			" mstnotificationtemplate.sendtoassignee,mstnotificationtemplate.sendtoassigneegroup,mstnotificationtemplate.sendtoassigneegroupmember," +
			" mstnotificationtemplate.eventtype eventtypeid,mstnotificationevents.eventname eventtype,mstnotificationtemplate.eventparams," +
			" coalesce( mstnotificationtemplate.smstemplateid,'') smstemplateid,coalesce(mstnotificationtemplate.smstype,0)smstypeid," +
			" IF(mstnotificationtemplate.smstype=1,'LongSMS','N/A') smstype,mstnotificationtemplate.converted,IF(mstnotificationtemplate.converted=1,'Converted','Normal') smstype,mstnotificationtemplate.activeflg FROM mstnotificationtemplate JOIN mstclient ON mstnotificationtemplate.clientid=mstclient.id JOIN mstorgnhierarchy ON" +
			" mstnotificationtemplate.clientid=mstorgnhierarchy.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecorddifferentiation ON " +
			" mstnotificationtemplate.recordtypeid=mstrecorddifferentiation.id AND mstnotificationtemplate.clientid=mstrecorddifferentiation.clientid AND" +
			" mstnotificationtemplate.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid  JOIN mstrecorddifferentiationtype ON" +
			" mstrecorddifferentiation.recorddifftypeid=mstrecorddifferentiationtype.id   JOIN mstrecorddifferentiation workcatdiff ON " +
			" mstnotificationtemplate.workingcategoryid=workcatdiff.id AND mstnotificationtemplate.clientid=workcatdiff.clientid AND" +
			" mstnotificationtemplate.mstorgnhirarchyid=workcatdiff.mstorgnhirarchyid  JOIN mstrecorddifferentiationtype difftype ON" +
			" workcatdiff.recorddifftypeid=difftype.id JOIN mstnotificationevents ON mstnotificationtemplate.eventtype=mstnotificationevents.id WHERE " +
			" mstnotificationtemplate.activeflg=1 AND mstnotificationtemplate.deleteflg=0 AND mstnotificationtemplate.clientid=? " +
			"  order by mstnotificationtemplate.id desc LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		query = "SELECT mstnotificationtemplate.id,mstnotificationtemplate.clientid,mstclient.name clientname, " +
			" mstnotificationtemplate.mstorgnhirarchyid,mstorgnhierarchy.name orgname,mstrecorddifferentiationtype.id recordtypedifftypeid," +
			" mstrecorddifferentiationtype.typename recordtypedifftype,mstnotificationtemplate.recordtypeid,mstrecorddifferentiation.name recordtype," +
			" difftype.id workingcategorytypeid,difftype.typename workingcategorytype,mstnotificationtemplate.workingcategoryid,workcatdiff.name workingcategory," +
			" mstnotificationtemplate.channeltype channeltypeid,IF(mstnotificationtemplate.channeltype=1,'Email','SMS') channeltype,coalesce( mstnotificationtemplate.subjectortitle,'') subjectortitle," +
			" coalesce( mstnotificationtemplate.body,'') body,coalesce( mstnotificationtemplate.additionalrecipient,'') additionalrecipient,mstnotificationtemplate.sendtocreator,mstnotificationtemplate.sendtooriginalcreator," +
			" mstnotificationtemplate.sendtoassignee,mstnotificationtemplate.sendtoassigneegroup,mstnotificationtemplate.sendtoassigneegroupmember," +
			" mstnotificationtemplate.eventtype eventtypeid,mstnotificationevents.eventname eventtype,mstnotificationtemplate.eventparams," +
			" coalesce( mstnotificationtemplate.smstemplateid,'') smstemplateid,coalesce(mstnotificationtemplate.smstype,0)smstypeid," +
			" IF(mstnotificationtemplate.smstype=1,'LongSMS','N/A') smstype,mstnotificationtemplate.converted,IF(mstnotificationtemplate.converted=1,'Converted','Normal') smstype,mstnotificationtemplate.activeflg FROM mstnotificationtemplate JOIN mstclient ON mstnotificationtemplate.clientid=mstclient.id JOIN mstorgnhierarchy ON" +
			" mstnotificationtemplate.clientid=mstorgnhierarchy.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecorddifferentiation ON " +
			" mstnotificationtemplate.recordtypeid=mstrecorddifferentiation.id AND mstnotificationtemplate.clientid=mstrecorddifferentiation.clientid AND" +
			" mstnotificationtemplate.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid  JOIN mstrecorddifferentiationtype ON" +
			" mstrecorddifferentiation.recorddifftypeid=mstrecorddifferentiationtype.id   JOIN mstrecorddifferentiation workcatdiff ON " +
			" mstnotificationtemplate.workingcategoryid=workcatdiff.id AND mstnotificationtemplate.clientid=workcatdiff.clientid AND" +
			" mstnotificationtemplate.mstorgnhirarchyid=workcatdiff.mstorgnhirarchyid  JOIN mstrecorddifferentiationtype difftype ON" +
			" workcatdiff.recorddifftypeid=difftype.id JOIN mstnotificationevents ON mstnotificationtemplate.eventtype=mstnotificationevents.id WHERE " +
			" mstnotificationtemplate.activeflg=1 AND mstnotificationtemplate.deleteflg=0 AND mstnotificationtemplate.clientid=? AND" +
			" mstnotificationtemplate.mstorgnhirarchyid=? order by mstnotificationtemplate.id desc LIMIT ?,?"
		params = append(params, page.ClientID)
		params = append(params, page.MstOrgnHirarchyID)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}

	rows, err := dbc.DB.Query(query, params...)

	if err != nil {
		logger.Log.Println("GetAllNotificationTemplates Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MstNotificationTemplateEntity{}
		var eparams string
		err = rows.Scan(&value.ID, &value.ClientID, &value.ClientName, &value.MstOrgnHirarchyID, &value.MstOrgnHirarchyName, &value.RecordTypeTypeID, &value.RecordTypeType, &value.RecordTypeID, &value.RecordType, &value.WorkingCategoryTypeID, &value.WorkingCategoryType, &value.WorkingCategoryID, &value.WorkingCategory, &value.ChannelTypeID, &value.ChannelType, &value.SubjectOrTitle, &value.Body, &value.AdditionalRecipient, &value.SendToCreator, &value.SendToOrgCreator, &value.SendToAssignee, &value.SendToAssigneeGroup, &value.SendToAssigneeGroupMember, &value.EventTypeID, &value.EventType, &eparams, &value.SmsTemplateID, &value.SmsTypeID, &value.SmsType, &value.Isconverted, &value.Convertedtype, &value.Activeflg)
		if err != nil {
			log.Println("err", err)
			//logger.Log.Println("GetAllNotificationTemplates JSON conversion Error", err)
			return values, err
		}
		log.Println("notification id ====>", value.ID)
		log.Println("eventParam ====>", eparams)

		err := json.Unmarshal([]byte(eparams), &value.EventParams)
		if err != nil {
			logger.Log.Println("GetAllNotificationTemplates JSON conversion Error", err)
			return values, err
		}
		//This is Written by josimoddin to show status in grid
		var statusname string
		if value.EventParams.StatusID > 0 {
			var query = "SELECT name FROM mstrecorddifferentiation where id=? and activeflg=1 and deleteflg=0"
			err := dbc.DB.QueryRow(query, value.EventParams.StatusID).Scan(&statusname)
			if err != nil && err != sql.ErrNoRows {
				logger.Log.Println("GetAllNotificationTemplates staus getting error", err)
				return values, err
			}
		}
		value.StatusName = statusname
		//ends here
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllRecipientsByTemplateID(page *entities.MstNotificationTemplateEntity) ([]entities.MstNotificationRecipientEntity, error) {
	logger.Log.Println("In side GetAllRecipientsByTemplateID")
	values := []entities.MstNotificationRecipientEntity{}
	var query = "SELECT mstnotificationrecipients.notificationtemplateid,mstnotificationrecipients.recipienttype,mstnotificationrecipients.groupid,COALESCE(mstclientsupportgroup.supportgroupname,'') groupname,mstnotificationrecipients.userid,COALESCE(mstclientuser.name,'') username FROM mstnotificationrecipients LEFT JOIN mstclientsupportgroup ON mstclientsupportgroup.id=mstnotificationrecipients.groupid AND mstclientsupportgroup.activeflg=1 AND mstclientsupportgroup.deleteflg=0 LEFT JOIN mstclientuser ON mstclientuser.id=mstnotificationrecipients.userid AND mstclientuser.activeflag=1 AND mstclientuser.deleteflag=0 WHERE mstnotificationrecipients.activeflg = 1 AND mstnotificationrecipients.deleteflg = 0 AND mstnotificationrecipients.notificationtemplateid=?"
	rows, err := dbc.DB.Query(query, page.ID)

	if err != nil {
		logger.Log.Println("GetAllRecipientsByTemplateID Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MstNotificationRecipientEntity{}
		rows.Scan(&value.NotificationTemplateID, &value.RecipientType, &value.GroupID, &value.GroupName, &value.UserID, &value.UserName)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllMsttemplate(page *entities.MsttemplateEntity, OrgnType int64) ([]entities.MsttemplateEntity, error) {
	logger.Log.Println("In side GelAllMsttemplate")
	values := []entities.MsttemplateEntity{}
	var getMsttemplate string
	var params []interface{}
	if OrgnType == 1 {
		getMsttemplate = "SELECT msttemplate.id, msttemplate.clientid, msttemplate.mstorgnhirarchyid, msttemplate.templatename, msttemplate.templatetype, msttemplate.templatecontent, msttemplate.autoclosuretime, msttemplate.activeflg,mstclient.name clientname,mstorgnhierarchy.name orgname FROM msttemplate,mstclient,mstorgnhierarchy WHERE msttemplate.activeflg=1 AND msttemplate.deleteflg=0 AND msttemplate.clientid=mstclient.id AND msttemplate.clientid=mstorgnhierarchy.clientid AND msttemplate.mstorgnhirarchyid=mstorgnhierarchy.id ORDER BY msttemplate.id DESC LIMIT ?,? "
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMsttemplate = "SELECT msttemplate.id, msttemplate.clientid, msttemplate.mstorgnhirarchyid, msttemplate.templatename, msttemplate.templatetype, msttemplate.templatecontent, msttemplate.autoclosuretime, msttemplate.activeflg,mstclient.name clientname,mstorgnhierarchy.name orgname FROM msttemplate,mstclient,mstorgnhierarchy WHERE msttemplate.clientid=? AND msttemplate.activeflg=1 AND msttemplate.deleteflg=0 AND msttemplate.clientid=mstclient.id AND msttemplate.clientid=mstorgnhierarchy.clientid AND msttemplate.mstorgnhirarchyid=mstorgnhierarchy.id ORDER BY msttemplate.id DESC LIMIT ?,? "
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMsttemplate = "SELECT msttemplate.id, msttemplate.clientid, msttemplate.mstorgnhirarchyid, msttemplate.templatename, msttemplate.templatetype, msttemplate.templatecontent, msttemplate.autoclosuretime, msttemplate.activeflg,mstclient.name clientname,mstorgnhierarchy.name orgname FROM msttemplate,mstclient,mstorgnhierarchy WHERE msttemplate.clientid=? AND msttemplate.mstorgnhirarchyid=?  AND msttemplate.activeflg=1 AND msttemplate.deleteflg=0 AND msttemplate.clientid=mstclient.id AND msttemplate.clientid=mstorgnhierarchy.clientid AND msttemplate.mstorgnhirarchyid=mstorgnhierarchy.id ORDER BY msttemplate.id DESC LIMIT ?,? "
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}

	rows, err := dbc.DB.Query(getMsttemplate, params...)

	// rows, err := dbc.DB.Query(getMsttemplate, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)

	if err != nil {
		logger.Log.Println("GetAllMsttemplate Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MsttemplateEntity{}
		rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Templatename, &value.Templatetype, &value.Templatecontent, &value.Autoclosuretime, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllMsttemplatediff(page *entities.MsttemplateEntity) ([]entities.MaptemplatediffEntity, error) {
	logger.Log.Println("In side GetAllMsttemplatediff")
	values := []entities.MaptemplatediffEntity{}
	rows, err := dbc.DB.Query(getMsttemplatediff, page.Clientid, page.Mstorgnhirarchyid, page.ID)

	if err != nil {
		logger.Log.Println("GetAllMsttemplatediff Get Statement Prepare Error", err)
		return values, err
	}
	defer rows.Close()
	for rows.Next() {
		value := entities.MaptemplatediffEntity{}
		rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.Templateid, &value.Recorddifftypeid, &value.Recorddiffid, &value.Activeflg, &value.Mstrecorddifferentiationtypename, &value.Mstrecorddifferentiationname, &value.RecorddifftypeParentid)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMsttemplate(tz *entities.MsttemplateEntity) error {
	logger.Log.Println("In side UpdateMsttemplate")
	stmt, err := dbc.DB.Prepare(updateMsttemplate)

	if err != nil {
		logger.Log.Println("UpdateMsttemplate Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.Templatename, tz.Templatetype, tz.Templatecontent, tz.Autoclosuretime, tz.ID)
	if err != nil {
		logger.Log.Println("UpdateMsttemplate Execute Statement  Error", err)
		return err
	}
	return nil
}

//update Api for notification
func (dbc DbConn) UpdateNotificationTemplate(tz *entities.MstNotificationTemplateEntity) error {
	logger.Log.Println("In side UpdateNotificationTemplate")

	if tz.ChannelTypeID == 1 {
		var query = "UPDATE mstnotificationtemplate SET channeltype=?,subjectortitle=?,body=?,additionalrecipient=?,sendtocreator=?,sendtooriginalcreator=?,sendtoassignee=?,sendtoassigneegroup=?,sendtoassigneegroupmember=?,eventtype=?,eventparams=?,converted=? WHERE id=? AND clientid=? AND mstorgnhirarchyid=?"
		stmt, err := dbc.DB.Prepare(query)

		if err != nil {
			logger.Log.Println("UpdateNotificationTemplate Prepare Statement  Error", err)
			return err
		}
		defer stmt.Close()
		b, err := json.Marshal(tz.EventParams)
		if err != nil {
			logger.Log.Println("UpdateNotificationTemplate JSON Conversion Error", err)
			return err
		}
		_, err = stmt.Exec(tz.ChannelTypeID, tz.SubjectOrTitle, tz.Body, tz.AdditionalRecipient, tz.SendToCreator, tz.SendToOrgCreator, tz.SendToAssignee, tz.SendToAssigneeGroup, tz.SendToAssigneeGroupMember, tz.EventTypeID, string(b), tz.Isconverted, tz.ID, tz.ClientID, tz.MstOrgnHirarchyID)
		if err != nil {
			logger.Log.Println("UpdateMsttemplate Execute Statement  Error", err)
			return err
		}

	} else {
		var query = "UPDATE mstnotificationtemplate SET channeltype=?,additionalrecipient=?,body=?,sendtocreator=?,sendtooriginalcreator=?,sendtoassignee=?,sendtoassigneegroup=?,sendtoassigneegroupmember=?,eventtype=?,eventparams=?,smstemplateid=?,smstype=?,converted=? WHERE id=? AND clientid=? AND mstorgnhirarchyid=?"
		stmt, err := dbc.DB.Prepare(query)

		if err != nil {
			logger.Log.Println("UpdateNotificationTemplate Prepare Statement  Error", err)
			return err
		}
		defer stmt.Close()
		b, err := json.Marshal(tz.EventParams)
		if err != nil {
			logger.Log.Println("UpdateNotificationTemplate JSON Conversion Error", err)
			return err
		}
		_, err = stmt.Exec(tz.ChannelTypeID, tz.AdditionalRecipient, tz.Body, tz.SendToCreator, tz.SendToOrgCreator, tz.SendToAssignee, tz.SendToAssigneeGroup, tz.SendToAssigneeGroupMember, tz.EventTypeID, string(b), tz.SmsTemplateID, tz.SmsTypeID, tz.Isconverted, tz.ID, tz.ClientID, tz.MstOrgnHirarchyID)
		if err != nil {
			logger.Log.Println("UpdateMsttemplate Execute Statement  Error", err)
			return err
		}
	}

	return nil
}

func (dbc DbConn) DeleteMsttemplate(tz *entities.MsttemplateEntity) error {
	logger.Log.Println("In side DeleteMsttemplate")
	stmt, err := dbc.DB.Prepare(deleteMsttemplate)

	if err != nil {
		logger.Log.Println("DeleteMsttemplate Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println("DeleteMsttemplate Execute Statement  Error", err)
		return err
	}
	return nil
}

//delete notification Template
func (dbc DbConn) DeleteNotificationTemplate(tz *entities.MstNotificationTemplateEntity) error {
	logger.Log.Println("In side DeleteNotificationTemplate")
	var query = "UPDATE mstnotificationtemplate SET activeflg=1,deleteflg=1   WHERE id=?"
	stmt, err := dbc.DB.Prepare(query)

	if err != nil {
		logger.Log.Println("DeleteNotificationTemplate Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println("DeleteNotificationTemplate Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMsttemplatediff(tz *entities.MsttemplateEntity) error {
	logger.Log.Println("In side DeleteMsttemplatediff")
	stmt, err := dbc.DB.Prepare(deleteMsttemplatediff)

	if err != nil {
		logger.Log.Println("DeleteMsttemplatediff Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println("DeleteMsttemplatediff Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteNotificationRecipients(tz *entities.MstNotificationTemplateEntity) error {
	logger.Log.Println("In side DeleteNotificationRecipients")
	var query = "UPDATE mstnotificationrecipients SET activeflg=1,deleteflg=1  WHERE notificationtemplateid=?"
	stmt, err := dbc.DB.Prepare(query)

	if err != nil {
		logger.Log.Println("DeleteMsttemplatediff Prepare Statement  Error", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println("DeleteMsttemplatediff Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetNotificationTemplateCount(tz *entities.MstNotificationTemplateEntity, OrgnTypeID int64) (int64, error) {
	logger.Log.Println("In side GetNotificationTemplateCount")
	var value int64
	// var query = "SELECT COUNT(mstnotificationtemplate.id) total FROM mstnotificationtemplate JOIN mstclient ON mstnotificationtemplate.clientid=mstclient.id JOIN mstorgnhierarchy ON mstnotificationtemplate.clientid=mstorgnhierarchy.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecorddifferentiation ON mstnotificationtemplate.recordtypeid=mstrecorddifferentiation.id AND mstnotificationtemplate.clientid=mstrecorddifferentiation.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid JOIN mstrecorddifferentiation workcatdiff ON mstnotificationtemplate.workingcategoryid=workcatdiff.id AND mstnotificationtemplate.clientid=workcatdiff.clientid AND mstnotificationtemplate.mstorgnhirarchyid=workcatdiff.mstorgnhirarchyid JOIN mstnotificationevents ON mstnotificationtemplate.eventtype=mstnotificationevents.id WHERE mstnotificationtemplate.activeflg=1 AND mstnotificationtemplate.deleteflg=0 AND mstnotificationtemplate.clientid=? AND mstnotificationtemplate.mstorgnhirarchyid=?"
	var query string
	var params []interface{}
	if OrgnTypeID == 1 {
		query = "SELECT COUNT(mstnotificationtemplate.id) total FROM mstnotificationtemplate JOIN mstclient ON mstnotificationtemplate.clientid=mstclient.id JOIN mstorgnhierarchy ON mstnotificationtemplate.clientid=mstorgnhierarchy.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecorddifferentiation ON mstnotificationtemplate.recordtypeid=mstrecorddifferentiation.id AND mstnotificationtemplate.clientid=mstrecorddifferentiation.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid JOIN mstrecorddifferentiation workcatdiff ON mstnotificationtemplate.workingcategoryid=workcatdiff.id AND mstnotificationtemplate.clientid=workcatdiff.clientid AND mstnotificationtemplate.mstorgnhirarchyid=workcatdiff.mstorgnhirarchyid JOIN mstnotificationevents ON mstnotificationtemplate.eventtype=mstnotificationevents.id WHERE mstnotificationtemplate.activeflg=1 AND mstnotificationtemplate.deleteflg=0 "
	} else if OrgnTypeID == 2 {
		query = "SELECT COUNT(mstnotificationtemplate.id) total FROM mstnotificationtemplate JOIN mstclient ON mstnotificationtemplate.clientid=mstclient.id JOIN mstorgnhierarchy ON mstnotificationtemplate.clientid=mstorgnhierarchy.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecorddifferentiation ON mstnotificationtemplate.recordtypeid=mstrecorddifferentiation.id AND mstnotificationtemplate.clientid=mstrecorddifferentiation.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid JOIN mstrecorddifferentiation workcatdiff ON mstnotificationtemplate.workingcategoryid=workcatdiff.id AND mstnotificationtemplate.clientid=workcatdiff.clientid AND mstnotificationtemplate.mstorgnhirarchyid=workcatdiff.mstorgnhirarchyid JOIN mstnotificationevents ON mstnotificationtemplate.eventtype=mstnotificationevents.id WHERE mstnotificationtemplate.activeflg=1 AND mstnotificationtemplate.deleteflg=0 AND mstnotificationtemplate.clientid=?"
		params = append(params, tz.ClientID)
	} else {
		query = "SELECT COUNT(mstnotificationtemplate.id) total FROM mstnotificationtemplate JOIN mstclient ON mstnotificationtemplate.clientid=mstclient.id JOIN mstorgnhierarchy ON mstnotificationtemplate.clientid=mstorgnhierarchy.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstorgnhierarchy.id JOIN mstrecorddifferentiation ON mstnotificationtemplate.recordtypeid=mstrecorddifferentiation.id AND mstnotificationtemplate.clientid=mstrecorddifferentiation.clientid AND mstnotificationtemplate.mstorgnhirarchyid=mstrecorddifferentiation.mstorgnhirarchyid JOIN mstrecorddifferentiation workcatdiff ON mstnotificationtemplate.workingcategoryid=workcatdiff.id AND mstnotificationtemplate.clientid=workcatdiff.clientid AND mstnotificationtemplate.mstorgnhirarchyid=workcatdiff.mstorgnhirarchyid JOIN mstnotificationevents ON mstnotificationtemplate.eventtype=mstnotificationevents.id WHERE mstnotificationtemplate.activeflg=1 AND mstnotificationtemplate.deleteflg=0 AND mstnotificationtemplate.clientid=? AND mstnotificationtemplate.mstorgnhirarchyid=?"
		params = append(params, tz.ClientID)
		params = append(params, tz.MstOrgnHirarchyID)
	}
	err := dbc.DB.QueryRow(query, params...).Scan(&value)

	// err := dbc.DB.QueryRow(query, tz.ClientID, tz.MstOrgnHirarchyID).Scan(&value)
	switch err {
	case sql.ErrNoRows:
		value = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetNotificationTemplateCount Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) GetMsttemplateCount(tz *entities.MsttemplateEntity, OrgnTypeID int64) (entities.MsttemplateEntities, error) {
	logger.Log.Println("In side GetMsttemplateCount")
	value := entities.MsttemplateEntities{}
	var getMsttemplatecount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMsttemplatecount = "SELECT count(msttemplate.id) total FROM msttemplate,mstclient,mstorgnhierarchy WHERE msttemplate.activeflg=1 AND msttemplate.deleteflg=0 AND msttemplate.clientid=mstclient.id AND msttemplate.clientid=mstorgnhierarchy.clientid AND msttemplate.mstorgnhirarchyid=mstorgnhierarchy.id"
	} else if OrgnTypeID == 2 {
		getMsttemplatecount = "SELECT count(msttemplate.id) total FROM msttemplate,mstclient,mstorgnhierarchy WHERE msttemplate.clientid=? AND msttemplate.activeflg=1 AND msttemplate.deleteflg=0 AND msttemplate.clientid=mstclient.id AND msttemplate.clientid=mstorgnhierarchy.clientid AND msttemplate.mstorgnhirarchyid=mstorgnhierarchy.id"
		params = append(params, tz.Clientid)
	} else {
		getMsttemplatecount = "SELECT count(msttemplate.id) total FROM msttemplate,mstclient,mstorgnhierarchy WHERE msttemplate.clientid=? AND msttemplate.mstorgnhirarchyid=?  AND msttemplate.activeflg=1 AND msttemplate.deleteflg=0 AND msttemplate.clientid=mstclient.id AND msttemplate.clientid=mstorgnhierarchy.clientid AND msttemplate.mstorgnhirarchyid=mstorgnhierarchy.id"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMsttemplatecount, params...).Scan(&value.Total)

	// err := dbc.DB.QueryRow(getMsttemplatecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMsttemplateCount Get Statement Prepare Error", err)
		return value, err
	}
}
