package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstScheduledNotification = "INSERT INTO mstschedulednotification (clientid, mstorgnhirarchyid, scheduledeventid, channeltype, emailsub, emailbody, sendtousersid, sendtogroupsid, additionalrecipint, triggerconditiondays, scheduledtime, recorddifftypeid, recorddiffid, priorityseqno) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
var duplicateMstScheduledNotification = "SELECT count(id) total FROM  mstschedulednotification WHERE clientid = ? AND mstorgnhirarchyid = ?  AND scheduledeventid=? AND channeltype=? AND emailsub=? AND emailbody=? AND sendtousersid=? AND sendtogroupsid=? AND additionalrecipint =? AND triggerconditiondays=? AND scheduledtime=? AND recorddifftypeid=? AND recorddiffid=? AND priorityseqno=? AND activeflg =1 AND deleteflg = 0 "

//var getMstScheduledNotification= "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.scheduledeventid as ScheduledEventid, a.channeltype as ChannelType, a.emailsub as EmailSub,a.emailbody as EmailBody,a.sendtousersid as SendToUserids,coalesce(a.sendtogroupsid,'') as SendToGroupids,a.additionalrecipint as AdditionalRecipint,a.triggerconditiondays as TriggerConditionDays,coalesce(a.scheduledtime,'') as ScheduleTime,a.recorddifftypeid as RecordDiffTypeid,a.recorddiffid as RecordDiffid,a.priorityseqno as PrioritySeqNo,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.eventname as ScheduledEventName,e.typename as RecordDiffTypeName,f.name as RecordDiffName,g.name as PrioritySeqName FROM mstschedulednotification a,mstclient b,mstorgnhierarchy c,mstnotificationevents d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g  WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.scheduledeventid=d.id AND a.recorddifftypeid=e.id AND a.recorddiffid=f.id AND g.recorddifftypeid=5 AND g.seqno=a.priorityseqno AND a.clientid=g.clientid AND a.mstorgnhirarchyid=g.mstorgnhirarchyid ORDER BY a.id DESC LIMIT ?,?"
//var getMstScheduledNotificationcount = "SELECT count(a.id) as total FROM mstschedulednotification a,mstclient b,mstorgnhierarchy c,mstnotificationevents d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g  WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.scheduledeventid=d.id AND a.recorddifftypeid=e.id AND a.recorddiffid=f.id AND g.recorddifftypeid=5 AND g.seqno=a.priorityseqno AND a.clientid=g.clientid AND a.mstorgnhirarchyid=g.mstorgnhirarchyid "
var updateMstScheduledNotification = "UPDATE mstschedulednotification SET clientid=?,mstorgnhirarchyid = ?, scheduledeventid = ?, channeltype = ?, emailsub = ?,emailbody=?,sendtousersid=?, sendtogroupsid=?,additionalrecipint=?,triggerconditiondays=?,scheduledtime=?,recorddifftypeid=?,recorddiffid=?,priorityseqno=? WHERE id = ? "
var deleteMstScheduledNotification = "UPDATE mstschedulednotification SET deleteflg ='1' WHERE id = ? "
var getUser = "SELECT a.useremail as user from mstclientuser a where a.id=?"
var getGroup = "SELECT a.name as groupname from mstsupportgrp a where a.id=?"
var getClientAndOrgWiseclientuser = "SELECT a.id as clientid,a.name as clientname from mstclientuser a where clientid=? AND mstorgnhirarchyid=?"

func (dbc DbConn) CheckDuplicateMstScheduledNotification(tz *entities.MstScheduledNotificationEntity) (entities.MstScheduledNotificationEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstScheduledNotification ")
	value := entities.MstScheduledNotificationEntities{}
	err := dbc.DB.QueryRow(duplicateMstScheduledNotification, tz.Clientid, tz.Mstorgnhirarchyid, tz.ScheduledEventid, tz.ChannelType, tz.EmailSub, tz.EmailBody, tz.SendToUserids, tz.SendToGroupids, tz.AdditionalRecipint, tz.TriggerConditionDays, tz.ScheduleTime, tz.RecordDiffTypeid, tz.RecordDiffid, tz.PrioritySeqNo).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstScheduledNotification Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) AddMstScheduledNotification(tz *entities.MstScheduledNotificationEntity) (int64, error) {
	logger.Log.Println("In side AddMstScheduledNotification")
	logger.Log.Println("Query -->", insertMstScheduledNotification)
	stmt, err := dbc.DB.Prepare(insertMstScheduledNotification)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstScheduledNotification Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.ScheduledEventid, tz.ChannelType, tz.EmailSub, tz.EmailBody, tz.SendToUserids, tz.SendToGroupids, tz.AdditionalRecipint, tz.TriggerConditionDays, tz.ScheduleTime, tz.RecordDiffTypeid, tz.RecordDiffid, tz.PrioritySeqNo)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.ScheduledEventid, tz.ChannelType, tz.EmailSub, tz.EmailBody, tz.SendToUserids, tz.SendToGroupids, tz.AdditionalRecipint, tz.TriggerConditionDays, tz.ScheduleTime, tz.RecordDiffTypeid, tz.RecordDiffid, tz.PrioritySeqNo)
	if err != nil {
		logger.Log.Println("AddMstScheduledNotification Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstScheduledNotification(page *entities.MstScheduledNotificationEntity, OrgnType int64) ([]entities.MstScheduledNotificationEntity, error) {
	logger.Log.Println("In side GetAllMstScheduledNotification")
	values := []entities.MstScheduledNotificationEntity{}

	var getMstScheduledNotification string
	var params []interface{}
	if OrgnType == 1 {
		getMstScheduledNotification = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.scheduledeventid as ScheduledEventid, a.channeltype as ChannelType, a.emailsub as EmailSub,a.emailbody as EmailBody,a.sendtousersid as SendToUserids,coalesce(a.sendtogroupsid,'') as SendToGroupids,a.additionalrecipint as AdditionalRecipint,a.triggerconditiondays as TriggerConditionDays,coalesce(a.scheduledtime,'') as ScheduleTime,a.recorddifftypeid as RecordDiffTypeid,a.recorddiffid as RecordDiffid,a.priorityseqno as PrioritySeqNo,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.eventname as ScheduledEventName,e.typename as RecordDiffTypeName,f.name as RecordDiffName,g.name as PrioritySeqName FROM mstschedulednotification a,mstclient b,mstorgnhierarchy c,mstnotificationevents d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g  WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.scheduledeventid=d.id AND a.recorddifftypeid=e.id AND a.recorddiffid=f.id AND g.recorddifftypeid=5 AND g.seqno=a.priorityseqno AND a.clientid=g.clientid AND a.mstorgnhirarchyid=g.mstorgnhirarchyid ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else if OrgnType == 2 {
		getMstScheduledNotification = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.scheduledeventid as ScheduledEventid, a.channeltype as ChannelType, a.emailsub as EmailSub,a.emailbody as EmailBody,a.sendtousersid as SendToUserids,coalesce(a.sendtogroupsid,'') as SendToGroupids,a.additionalrecipint as AdditionalRecipint,a.triggerconditiondays as TriggerConditionDays,coalesce(a.scheduledtime,'') as ScheduleTime,a.recorddifftypeid as RecordDiffTypeid,a.recorddiffid as RecordDiffid,a.priorityseqno as PrioritySeqNo,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.eventname as ScheduledEventName,e.typename as RecordDiffTypeName,f.name as RecordDiffName,g.name as PrioritySeqName FROM mstschedulednotification a,mstclient b,mstorgnhierarchy c,mstnotificationevents d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g  WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.scheduledeventid=d.id AND a.recorddifftypeid=e.id AND a.recorddiffid=f.id AND g.recorddifftypeid=5 AND g.seqno=a.priorityseqno AND a.clientid=g.clientid AND a.mstorgnhirarchyid=g.mstorgnhirarchyid ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	} else {
		getMstScheduledNotification = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.scheduledeventid as ScheduledEventid, a.channeltype as ChannelType, a.emailsub as EmailSub,a.emailbody as EmailBody,a.sendtousersid as SendToUserids,coalesce(a.sendtogroupsid,'') as SendToGroupids,a.additionalrecipint as AdditionalRecipint,a.triggerconditiondays as TriggerConditionDays,coalesce(a.scheduledtime,'') as ScheduleTime,a.recorddifftypeid as RecordDiffTypeid,a.recorddiffid as RecordDiffid,a.priorityseqno as PrioritySeqNo,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname,d.eventname as ScheduledEventName,e.typename as RecordDiffTypeName,f.name as RecordDiffName,g.name as PrioritySeqName FROM mstschedulednotification a,mstclient b,mstorgnhierarchy c,mstnotificationevents d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g  WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.scheduledeventid=d.id AND a.recorddifftypeid=e.id AND a.recorddiffid=f.id AND g.recorddifftypeid=5 AND g.seqno=a.priorityseqno AND a.clientid=g.clientid AND a.mstorgnhirarchyid=g.mstorgnhirarchyid ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, page.Clientid)
		params = append(params, page.Mstorgnhirarchyid)
		params = append(params, page.Offset)
		params = append(params, page.Limit)
	}

	rows, err := dbc.DB.Query(getMstScheduledNotification, params...)

	//rows, err := dbc.DB.Query(getMstScheduledNotification, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstScheduledNotification Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstScheduledNotificationEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.ScheduledEventid, &value.ChannelType, &value.EmailSub, &value.EmailBody, &value.SendToUserids, &value.SendToGroupids, &value.AdditionalRecipint, &value.TriggerConditionDays, &value.ScheduleTime, &value.RecordDiffTypeid, &value.RecordDiffid, &value.PrioritySeqNo, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname, &value.ScheduledEventName, &value.RecordDiffTypeName, &value.RecordDiffName, &value.PrioritySeqName)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstScheduledNotification(tz *entities.MstScheduledNotificationEntity) error {
	logger.Log.Println("In side UpdateMstScheduledNotification")
	stmt, err := dbc.DB.Prepare(updateMstScheduledNotification)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstScheduledNotification Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.ScheduledEventid, tz.ChannelType, tz.EmailSub, tz.EmailBody, tz.SendToUserids, tz.SendToGroupids, tz.AdditionalRecipint, tz.TriggerConditionDays, tz.ScheduleTime, tz.RecordDiffTypeid, tz.RecordDiffid, tz.PrioritySeqNo, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstScheduledNotification Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstScheduledNotification(tz *entities.MstScheduledNotificationEntity) error {
	logger.Log.Println("In side DeleteMstScheduledNotification", tz)
	stmt, err := dbc.DB.Prepare(deleteMstScheduledNotification)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstScheduledNotification Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstScheduledNotification Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstScheduledNotificationCount(tz *entities.MstScheduledNotificationEntity, OrgnTypeID int64) (entities.MstScheduledNotificationEntities, error) {
	logger.Log.Println("In side GetMstScheduledNotificationCount")
	value := entities.MstScheduledNotificationEntities{}
	var getMstScheduledNotificationcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getMstScheduledNotificationcount = "SELECT count(a.id) as total FROM mstschedulednotification a,mstclient b,mstorgnhierarchy c,mstnotificationevents d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g  WHERE a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.scheduledeventid=d.id AND a.recorddifftypeid=e.id AND a.recorddiffid=f.id AND g.recorddifftypeid=5 AND g.seqno=a.priorityseqno AND a.clientid=g.clientid AND a.mstorgnhirarchyid=g.mstorgnhirarchyid "
	} else if OrgnTypeID == 2 {
		getMstScheduledNotificationcount = "SELECT count(a.id) as total FROM mstschedulednotification a,mstclient b,mstorgnhierarchy c,mstnotificationevents d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g  WHERE a.clientid = ? AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.scheduledeventid=d.id AND a.recorddifftypeid=e.id AND a.recorddiffid=f.id AND g.recorddifftypeid=5 AND g.seqno=a.priorityseqno AND a.clientid=g.clientid AND a.mstorgnhirarchyid=g.mstorgnhirarchyid "
		params = append(params, tz.Clientid)
	} else {
		getMstScheduledNotificationcount = "SELECT count(a.id) as total FROM mstschedulednotification a,mstclient b,mstorgnhierarchy c,mstnotificationevents d,mstrecorddifferentiationtype e,mstrecorddifferentiation f,mstrecorddifferentiation g  WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.scheduledeventid=d.id AND a.recorddifftypeid=e.id AND a.recorddiffid=f.id AND g.recorddifftypeid=5 AND g.seqno=a.priorityseqno AND a.clientid=g.clientid AND a.mstorgnhirarchyid=g.mstorgnhirarchyid "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getMstScheduledNotificationcount, params...).Scan(&value.Total)

	//err := dbc.DB.QueryRow(getMstScheduledNotificationcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstScheduledNotificationCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetUser(n int64) (string, error) {
	logger.Log.Println("In side GetUser")
	var user string
	err := dbc.DB.QueryRow(getUser, n).Scan(&user)
	switch err {
	case sql.ErrNoRows:
		user = "null"
		return user, nil
	case nil:
		return user, nil
	default:
		logger.Log.Println("GetUser Get Statement Prepare Error", err)
		return user, err
	}
}

func (dbc DbConn) GetGroup(n int64) (string, error) {
	logger.Log.Println("In side GetGroup")
	var groupname string
	err := dbc.DB.QueryRow(getGroup, n).Scan(&groupname)
	switch err {
	case sql.ErrNoRows:
		groupname = "null"
		return groupname, nil
	case nil:
		return groupname, nil
	default:
		logger.Log.Println("GetGroup Get Statement Prepare Error", err)
		return groupname, err
	}
}

func (dbc DbConn) GetClientAndOrgWiseclientuser(page *entities.MstScheduledNotificationEntity) ([]entities.GetClientAndOrgWiseclientuserEntity, error) {
	logger.Log.Println("In side GetAllMstScheduledNotification")
	values := []entities.GetClientAndOrgWiseclientuserEntity{}

	rows, err := dbc.DB.Query(getClientAndOrgWiseclientuser, page.Clientid, page.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstScheduledNotification Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.GetClientAndOrgWiseclientuserEntity{}
		rows.Scan(&value.Clientid, &value.Clientname)
		values = append(values, value)
	}
	return values, nil
}
