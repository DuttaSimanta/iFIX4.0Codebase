package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertMstAttribute = "INSERT INTO mstadfsattributes ( clientid, mstorgnhirarchyid, adfsattribute) VALUES (?,?,?) "
var duplicateMstAttribute = "SELECT count(id) total FROM  mstadfsattributes WHERE clientid = ? AND mstorgnhirarchyid = ?  AND adfsattribute=? AND activeflg =1 AND deleteflg = 0 "

var getMstAttributecount = "SELECT count(a.id) as total FROM mstadfsattributes a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?  AND a.deleteflg =0 and a.activeflg=1 AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id ORDER BY a.id DESC LIMIT ?,?"
var updateMstAttribute = "UPDATE mstadfsattributes SET clientid=?,mstorgnhirarchyid = ?, adfsattribute = ? WHERE id = ? "

var getmstattrbutes = "select adfsattribute from mstadfsattributes where clientid = ? AND mstorgnhirarchyid = ?  AND activeflg =1 AND deleteflg = 0"
var deleteMstAttribute = "UPDATE mstadfsattributes SET deleteflg ='1' WHERE id = ? "

func (dbc DbConn) CheckDuplicateMstAttribute(tz *entities.MstAttributeEntity) (*entities.MstAttributeEntities, error) {
	logger.Log.Println("In side CheckDuplicateMstAttributeDao ")
	value := entities.MstAttributeEntities{}
	err := dbc.DB.QueryRow(duplicateMstAttribute, tz.Clientid, tz.Mstorgnhirarchyid, tz.Adfsattribute).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return &value, nil
	case nil:
		return &value, nil
	default:
		logger.Log.Println("CheckDuplicateMstAttributeDao Get Statement Prepare Error", err)
		return &value, err
	}
}

func (dbc DbConn) AddMstAttribute(tz *entities.MstAttributeEntity) (int64, error) {
	logger.Log.Println("In side AddMstAttributeDao")
	logger.Log.Println("Query -->", insertMstAttribute)
	stmt, err := dbc.DB.Prepare(insertMstAttribute)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddMstAttributeDao Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Adfsattribute)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Adfsattribute)
	if err != nil {
		logger.Log.Println("AddMstAttributeDao Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllMstAttribute(tz *entities.MstAttributeEntity, OrgnType int64) ([]entities.MstAttributeEntity, error) {
	logger.Log.Println("In side GetAllMstAttributeDao")
	values := []entities.MstAttributeEntity{}

	var getMstAttribute string
	var params []interface{}
	if OrgnType == 1 {
		getMstAttribute = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.adfsattribute as adfsattribute,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstadfsattributes a,mstclient b,mstorgnhierarchy c WHERE    a.clientid =b.id AND a.mstorgnhirarchyid = c.id and a.deleteflg =0 and a.activeflg=1  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getMstAttribute = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.adfsattribute as adfsattribute,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstadfsattributes a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ?   AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getMstAttribute = "SELECT a.id as Id, a.clientid as Clientid, a.mstorgnhirarchyid as Mstorgnhirarchyid,a.adfsattribute as adfsattribute,a.activeflg as Activeflg,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstadfsattributes a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?    AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.deleteflg =0 and a.activeflg=1 ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getMstAttribute, params...)

	//rows, err := dbc.DB.Query(getMstAttribute, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllMstAttributeDao Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstAttributeEntity{}
		rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.Adfsattribute, &value.Activeflg, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateMstAttribute(tz *entities.MstAttributeEntity) error {
	logger.Log.Println("In side UpdateMstAttributeDao")
	stmt, err := dbc.DB.Prepare(updateMstAttribute)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateMstAttributeDao Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Adfsattribute, tz.Id)
	if err != nil {
		logger.Log.Println("UpdateMstAttributeDao Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteMstAttribute(tz *entities.MstAttributeEntity) error {
	logger.Log.Println("In side DeleteMstAttributeDao", tz)
	stmt, err := dbc.DB.Prepare(deleteMstAttribute)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteMstAttributeDao Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteMstAttributeDao Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetMstAttributeCount(tz *entities.MstAttributeEntity, OrgnTypeID int64) (entities.MstAttributeEntities, error) {
	logger.Log.Println("In side GetMstAttributeCountdao")
	value := entities.MstAttributeEntities{}

	var getClientsupportgroupnewcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstadfsattributes a,mstclient b,mstorgnhierarchy c WHERE   a.clientid =b.id AND a.mstorgnhirarchyid = c.id  and a.deleteflg =0 and a.activeflg=1 "
	} else if OrgnTypeID == 2 {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstadfsattributes a,mstclient b,mstorgnhierarchy c WHERE  a.clientid = ? AND  a.clientid =b.id AND a.mstorgnhirarchyid = c.id and    a.deleteflg =0 and a.activeflg=1"
		params = append(params, tz.Clientid)
	} else {
		getClientsupportgroupnewcount = "SELECT count(a.id) as total FROM mstadfsattributes a,mstclient b,mstorgnhierarchy c WHERE a.clientid = ? AND a.mstorgnhirarchyid = ?   AND a.clientid =b.id AND a.mstorgnhirarchyid = c.id AND a.deleteflg =0 and a.activeflg=1 "
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getClientsupportgroupnewcount, params...).Scan(&value.Total)

	//err := dbc.DB.QueryRow(getMstAttributecount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetMstAttributeCountdao Get Statement Prepare Error", err)
		return value, err
	}
}

// func (dbc DbConn) GetSeq(tz *entities.MstAttributeEntity) (interface{}, error) {
// 	logger.Log.Println("In side GetSeq")
// 	var seq interface{}
// 	// value := entities.CountryEntities{}
// 	err := dbc.DB.QueryRow("select max(seqno) as seq FROM mstadfsattributes where activeflg=1 and deleteflg=0").Scan(&seq)
// 	switch err {
// 	case sql.ErrNoRows:
// 		return seq, nil
// 	case nil:
// 		logger.Log.Println(seq)
// 		return seq, nil
// 	default:
// 		logger.Log.Println("Getseq Get Statement Prepare Error", err)
// 		return seq, err
// 	}
// }
// func (mdao DbConn) GetSeq(tz *entities.MstAttributeEntity) ([]entities.MstAttributeEntity, error) {
// 	logger.Log.Println("In side GetSeqdao")
// 	values := []entities.MstAttributeEntity{}
// 	rows, err := mdao.DB.Query(getseq)

// 	if err != nil {
// 		logger.Log.Print("GetSeqdao Get Statement Prepare Error", err)
// 		return values, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		value := entities.MstAttributeEntity{}
// 		rows.Scan(&value.Sequence)
// 		values = append(values, value)
// 	}
// 	return values, nil
// }
// func (mdao DbConn) GetSeqcopy(tz *entities.MstAttributeEntity) ([]entities.MstAttributeEntity, error) {
// 	logger.Log.Println("In side GetSeqcopydao")
// 	values := []entities.MstAttributeEntity{}
// 	rows, err := mdao.DB.Query(getseqcopy, tz.Activitydesc)

// 	if err != nil {
// 		logger.Log.Print("GetSeqcopydao Get Statement Prepare Error", err)
// 		return values, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		value := entities.MstAttributeEntity{}
// 		rows.Scan(&value.Sequence)
// 		values = append(values, value)
// 	}
// 	return values, nil
// }

func (dbc DbConn) GetMstAttribute(tz *entities.MstAttributeEntity) ([]entities.Attributes, error) {
	logger.Log.Println("In side GetAttributeDao")
	values := []entities.Attributes{}

	rows, err := dbc.DB.Query(getmstattrbutes, tz.Clientid, tz.Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetOrgWiseActivitydescDao Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.Attributes{}

		rows.Scan(&value.Adfsattribute)
		//logger.Log.Println(values)
		values = append(values, value)
	}
	return values, nil
}
