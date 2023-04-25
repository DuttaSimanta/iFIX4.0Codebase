package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertBanner = "INSERT INTO mstbanner (clientid, mstorgnhirarchyid, groupid, message,starttime,endtime,sequence,color,size) VALUES (?,?,?,?,?,?,?,?,?)"
var duplicateBanner = "SELECT count(id) total FROM  mstbanner WHERE clientid = ? AND mstorgnhirarchyid = ? AND groupid = ? AND message = ? AND starttime=? AND endtime =?  AND deleteflg = 0 AND activeflg=1"

// var getBanner = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as groupid, a.message as Message,a.starttime as Starttime,a.endtime as Endtime, a.sequence as Sequence ,b.name as Mstorgnhirarchyname,c.name as Groupname, a.color as Color,a.size as Size FROM mstbanner a ,mstorgnhierarchy b,mstsupportgrp c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.groupid = c.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
// var getBannercount = "SELECT count(a.id) as total FROM mstbanner a ,mstorgnhierarchy b,mstsupportgrp c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.groupid = c.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0"
var updateBanner = "UPDATE mstbanner SET clientid = ?,mstorgnhirarchyid = ?, groupid = ?, message = ?,starttime=?,endtime=? WHERE id = ? "
var deleteBanner = "UPDATE mstbanner SET deleteflg = '1' WHERE id = ? "
var getBannermessage = "select a.message as Message,a.color as Color,a.size as Size from mstbanner a where  a.clientid=? and a.mstorgnhirarchyid=? and a.groupid = ? and UNIX_TIMESTAMP()>= a.starttime and UNIX_TIMESTAMP()<=a.endtime and activeflg=1 and deleteflg=0 ORDER BY a.sequence ASC LIMIT 3"
var sequenceupdate = "UPDATE mstbanner SET sequence=?,color=?,size=? WHERE id = ? "

func (dbc DbConn) CheckDuplicateBanner(tz *entities.BannerEntity, i int) (entities.BannerEntities, error) {
	logger.Log.Println("In side CheckDuplicateBanner")
	value := entities.BannerEntities{}
	err := dbc.DB.QueryRow(duplicateBanner, tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid[i], tz.Message, tz.Starttime, tz.Endtime).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateBanner Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc TxConn) InsertBanner(tz *entities.BannerEntity, i int) (int64, error) {
	logger.Log.Println("In side InsertBanner")
	logger.Log.Println("Query -->", insertBanner)
	stmt, err := dbc.TX.Prepare(insertBanner)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertBanner Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid[i], tz.Message, tz.Starttime, tz.Endtime, tz.Sequence)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid[i], tz.Message, tz.Starttime, tz.Endtime, tz.Sequence, tz.Color, tz.Size)
	if err != nil {
		logger.Log.Println("InsertBanner Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllBanner(tz *entities.BannerEntity, OrgnType int64) ([]entities.BannerEntity, error) {
	logger.Log.Println("In side dao GelAllBanner")
	values := []entities.BannerEntity{}
	var getBanner string
	var params []interface{}
	if OrgnType == 1 {
		getBanner = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as groupid, a.message as Message,a.starttime as Starttime,a.endtime as Endtime, a.sequence as Sequence ,b.name as Mstorgnhirarchyname,c.name as Groupname, a.color as Color,a.size as Size FROM mstbanner a ,mstorgnhierarchy b,mstsupportgrp c WHERE a.mstorgnhirarchyid = b.id and a.groupid = c.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getBanner = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as groupid, a.message as Message,a.starttime as Starttime,a.endtime as Endtime, a.sequence as Sequence ,b.name as Mstorgnhirarchyname,c.name as Groupname, a.color as Color,a.size as Size FROM mstbanner a ,mstorgnhierarchy b,mstsupportgrp c WHERE a.clientid=? and a.mstorgnhirarchyid = b.id and a.groupid = c.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getBanner = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid, a.groupid as groupid, a.message as Message,a.starttime as Starttime,a.endtime as Endtime, a.sequence as Sequence ,b.name as Mstorgnhirarchyname,c.name as Groupname, a.color as Color,a.size as Size FROM mstbanner a ,mstorgnhierarchy b,mstsupportgrp c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.groupid = c.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}

	rows, err := dbc.DB.Query(getBanner, params...)
	// rows, err := dbc.DB.Query(getBanner, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllBanner Get Statement Prepare Error", err)
		return values, err
	}
	var groupid int64
	for rows.Next() {
		value := entities.BannerEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &groupid, &value.Message, &value.Starttime, &value.Endtime, &value.Sequence, &value.Mstorgnhirarchyname, &value.Groupname, &value.Color, &value.Size)
		value.Groupid = append(value.Groupid, groupid)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateBanner(tz *entities.BannerEntity) error {
	logger.Log.Println("In side UpdateBanner")
	stmt, err := dbc.DB.Prepare(updateBanner)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateBanner Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Groupid[0], tz.Message, tz.Starttime, tz.Endtime, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateBanner Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteBanner(tz *entities.BannerEntity) error {
	logger.Log.Println("In side DeleteBanner")
	stmt, err := dbc.DB.Prepare(deleteBanner)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteBanner Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteBanner Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetBannerCount(tz *entities.BannerEntity, OrgnTypeID int64) (entities.BannerEntities, error) {
	logger.Log.Println("In side GetBannerCount")
	value := entities.BannerEntities{}
	var getBannercount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getBannercount = "SELECT count(a.id) as total FROM mstbanner a ,mstorgnhierarchy b,mstsupportgrp c WHERE  a.mstorgnhirarchyid = b.id and a.groupid = c.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getBannercount = "SELECT count(a.id) as total FROM mstbanner a ,mstorgnhierarchy b,mstsupportgrp c WHERE a.clientid=? and a.mstorgnhirarchyid = b.id and a.groupid = c.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0"
		params = append(params, tz.Clientid)
	} else {
		getBannercount = "SELECT count(a.id) as total FROM mstbanner a ,mstorgnhierarchy b,mstsupportgrp c WHERE a.clientid=? and a.mstorgnhirarchyid=? and a.mstorgnhirarchyid = b.id and a.groupid = c.id and a.activeflg=1 and a.deleteflg =0  and c.activeflg=1 and c.deleteflg=0"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getBannercount, params...).Scan(&value.Total)
	// err := dbc.DB.QueryRow(getBannercount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetBannerCount Get Statement Prepare Error", err)
		return value, err
	}
}

func (dbc DbConn) GetAllMessage(page *entities.BannerEntity) ([]entities.BannerMessageEntity, error) {
	logger.Log.Println("In side GelAllMessage")
	values := []entities.BannerMessageEntity{}
	rows, err := dbc.DB.Query(getBannermessage, page.Clientid, page.Mstorgnhirarchyid, page.Groupid[0])
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllBannermessage Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.BannerMessageEntity{}
		rows.Scan(&value.Message, &value.Color, &value.Size)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateBannerSequence(tz *entities.BannerEntity) error {
	logger.Log.Println("In side UpdateBannerSequence")
	stmt, err := dbc.DB.Prepare(sequenceupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateBannerSequence Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Sequence, tz.Color, tz.Size, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateBannerSequence Execute Statement  Error", err)
		return err
	}
	return nil
}
