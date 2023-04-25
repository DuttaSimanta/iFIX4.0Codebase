package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertRecordConfigWithoutTx = "INSERT INTO mstrecordconfig (clientid, mstorgnhirarchyid,recorddifftypeid,recorddiffid, isclient) VALUES (?,?,?,?,?)"
var insertRecordConfig = "INSERT INTO mstrecordconfig (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, prefix, year, month, day, configurezero, isclient) VALUES (?,?,?,?,?,?,?,?,?,?)"
var insertRecordIncrement = "INSERT INTO mstrecordautoincreament (clientid, mstorgnhirarchyid, recorddifftypeid, recorddiffid, number,isclient) VALUES (?,?,?,?,?,?)"

var duplicateRecordConfigIncrement = "SELECT count(id) total FROM  mstrecordconfig WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid = ? AND recorddiffid = ? AND deleteflg = 0 AND activeflg=1"
var duplicateRecordConfig = "SELECT count(id) total FROM  mstrecordconfig WHERE clientid = ? AND mstorgnhirarchyid = ? AND recorddifftypeid=? AND recorddiffid=? AND deleteflg = 0 AND activeflg=1"

// var getRecordConfigIncrement = "SELECT a.id as id,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid,a.isclient as isclient,coalesce(a.recorddifftypeid,0) as recorddifftypeid,coalesce(a.recorddiffid,0) as recorddiffid,coalesce(a.prefix,'') as prefix,coalesce(a.year,'') as year,coalesce(a.month,'') as month,coalesce(a.day,'') as day, coalesce(a.configurezero,'') as configurezero ,coalesce(b.typename,'') as RecorddifftypeName,coalesce(c.name,'') as RecorddiffName,coalesce(d.name,'') as Clientname,coalesce(e.name,'') as Mstorgnhirarchyname,coalesce(f.number,0) as number FROM  mstrecordconfig a  left join mstrecorddifferentiationtype b on a.recorddifftypeid=b.id AND b.activeflg=1 AND b.deleteflg=0 left join mstrecorddifferentiation c on a.recorddiffid=c.id AND c.activeflg=1 AND c.deleteflg=0 join mstclient d on a.clientid=d.id join mstorgnhierarchy e on a.mstorgnhirarchyid=e.id left join mstrecordautoincreament f on a.clientid=f.clientid AND a.mstorgnhirarchyid=f.mstorgnhirarchyid AND a.recorddifftypeid=f.recorddifftypeid AND a.recorddiffid=f.recorddiffid AND f.activeflg=1 AND f.deleteflg=0 where a.activeflg=1 and a.deleteflg=0 AND a.clientid=? AND a.mstorgnhirarchyid=? ORDER BY a.id DESC LIMIT ?,?"
// var getRecordConfigIncrementcount = "SELECT count(a.id) as total FROM mstrecordconfig a  left join mstrecorddifferentiationtype b on a.recorddifftypeid=b.id AND b.activeflg=1 AND b.deleteflg=0 left join mstrecorddifferentiation c on a.recorddiffid=c.id AND c.activeflg=1 AND c.deleteflg=0 join mstclient d on a.clientid=d.id join mstorgnhierarchy e on a.mstorgnhirarchyid=e.id left join mstrecordautoincreament f on a.clientid=f.clientid AND a.mstorgnhirarchyid=f.mstorgnhirarchyid AND a.recorddifftypeid=f.recorddifftypeid AND a.recorddiffid=f.recorddiffid AND f.activeflg=1 AND f.deleteflg=0 where a.activeflg=1 and a.deleteflg=0 AND a.clientid=? AND a.mstorgnhirarchyid=?"

var updateRecordConfigWithoutTx = "UPDATE mstrecordconfig SET clientid=?, mstorgnhirarchyid = ?,recorddifftypeid=?,recorddiffid=? WHERE id = ? AND isclient=? "
var updateRecordConfig = "UPDATE mstrecordconfig SET prefix = ?,year = ?,month=?,day=?, configurezero=? WHERE id = ? AND isclient=?"
var updateRecordConfigincrement = "UPDATE mstrecordautoincreament SET number = ? WHERE clientid=? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND recorddiffid=? AND deleteflg=0 AND isclient=?"

var deleteRecordConfig = "UPDATE mstrecordconfig SET deleteflg = '1' WHERE id = ? "
var deleteRecordIncrement = "UPDATE mstrecordautoincreament SET deleteflg = '1' WHERE clientid = ? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND recorddiffid=?"

func (dbc DbConn) CheckDuplicateRecordConfigIncrement(tz *entities.RecordConfigIncrementEntity) (entities.RecordConfigIncrementEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecordConfigIncrement")
	value := entities.RecordConfigIncrementEntities{}
	err := dbc.DB.QueryRow(duplicateRecordConfigIncrement, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateMstExcelTemplate Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) CheckDuplicateRecordConfig(tz *entities.RecordConfigIncrementEntity) (entities.RecordConfigIncrementEntities, error) {
	logger.Log.Println("In side CheckDuplicateRecordConfigIncrement ")
	value := entities.RecordConfigIncrementEntities{}
	err := dbc.DB.QueryRow(duplicateRecordConfig, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateRecordConfigIncrement Get Statement Prepare Error", err)
		return value, err
	}
}
func (dbc DbConn) AddRecordConfigWithoutTx(tz *entities.RecordConfigIncrementEntity) (int64, error) {
	logger.Log.Println("In side AddRecordConfigIncrement")
	logger.Log.Println("Query -->", insertRecordConfigWithoutTx)
	stmt, err := dbc.DB.Prepare(insertRecordConfigWithoutTx)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("AddRecordConfigIncrement Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.IsClient)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.IsClient)
	if err != nil {
		logger.Log.Println("AddRecordConfigIncrement Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc TxConn) AddRecordConfig(tz *entities.RecordConfigIncrementEntity) (int64, error) {
	logger.Log.Println("In side addRecordConfigIncrement")
	logger.Log.Println("Query -->", insertRecordConfig)
	stmt, err := dbc.TX.Prepare(insertRecordConfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordConfigIncrement Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Prefix, tz.Year, tz.Month, tz.Day, tz.Configurezero, tz.IsClient)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Prefix, tz.Year, tz.Month, tz.Day, tz.Configurezero, tz.IsClient)
	if err != nil {
		logger.Log.Println("InsertRecordConfigIncrement Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}
func (dbc TxConn) AddRecordIncrement(tz *entities.RecordConfigIncrementEntity) (int64, error) {
	logger.Log.Println("In side InsertRecordConfigIncrement")
	logger.Log.Println("Query -->", insertRecordIncrement)
	stmt, err := dbc.TX.Prepare(insertRecordIncrement)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertRecordConfigIncrement Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Number, tz.IsClient)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Number, tz.IsClient)
	if err != nil {
		logger.Log.Println("InsertRecordConfigIncrement Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func (dbc DbConn) GetAllRecordConfigIncrement(tz *entities.RecordConfigIncrementEntity, OrgnType int64) ([]entities.RecordConfigIncrementEntity, error) {
	logger.Log.Println("In side dao GelAllRecordConfigIncrement")
	values := []entities.RecordConfigIncrementEntity{}
	var getRecordConfigIncrement string
	var params []interface{}
	if OrgnType == 1 {
		getRecordConfigIncrement = "SELECT a.id as id,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid,a.isclient as isclient,coalesce(a.recorddifftypeid,0) as recorddifftypeid,coalesce(a.recorddiffid,0) as recorddiffid,coalesce(a.prefix,'') as prefix,coalesce(a.year,'') as year,coalesce(a.month,'') as month,coalesce(a.day,'') as day, coalesce(a.configurezero,'') as configurezero ,coalesce(b.typename,'') as RecorddifftypeName,coalesce(c.name,'') as RecorddiffName,coalesce(d.name,'') as Clientname,coalesce(e.name,'') as Mstorgnhirarchyname,coalesce(f.number,0) as number FROM  mstrecordconfig a  left join mstrecorddifferentiationtype b on a.recorddifftypeid=b.id AND b.activeflg=1 AND b.deleteflg=0 left join mstrecorddifferentiation c on a.recorddiffid=c.id AND c.activeflg=1 AND c.deleteflg=0 join mstclient d on a.clientid=d.id join mstorgnhierarchy e on a.mstorgnhirarchyid=e.id left join mstrecordautoincreament f on a.clientid=f.clientid AND a.mstorgnhirarchyid=f.mstorgnhirarchyid AND a.recorddifftypeid=f.recorddifftypeid AND a.recorddiffid=f.recorddiffid AND f.activeflg=1 AND f.deleteflg=0 where a.activeflg=1 and a.deleteflg=0  ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else if OrgnType == 2 {
		getRecordConfigIncrement = "SELECT a.id as id,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid,a.isclient as isclient,coalesce(a.recorddifftypeid,0) as recorddifftypeid,coalesce(a.recorddiffid,0) as recorddiffid,coalesce(a.prefix,'') as prefix,coalesce(a.year,'') as year,coalesce(a.month,'') as month,coalesce(a.day,'') as day, coalesce(a.configurezero,'') as configurezero ,coalesce(b.typename,'') as RecorddifftypeName,coalesce(c.name,'') as RecorddiffName,coalesce(d.name,'') as Clientname,coalesce(e.name,'') as Mstorgnhirarchyname,coalesce(f.number,0) as number FROM  mstrecordconfig a  left join mstrecorddifferentiationtype b on a.recorddifftypeid=b.id AND b.activeflg=1 AND b.deleteflg=0 left join mstrecorddifferentiation c on a.recorddiffid=c.id AND c.activeflg=1 AND c.deleteflg=0 join mstclient d on a.clientid=d.id join mstorgnhierarchy e on a.mstorgnhirarchyid=e.id left join mstrecordautoincreament f on a.clientid=f.clientid AND a.mstorgnhirarchyid=f.mstorgnhirarchyid AND a.recorddifftypeid=f.recorddifftypeid AND a.recorddiffid=f.recorddiffid AND f.activeflg=1 AND f.deleteflg=0 where a.activeflg=1 and a.deleteflg=0 AND a.clientid=? ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	} else {
		getRecordConfigIncrement = "SELECT a.id as id,a.clientid as Clientid,a.mstorgnhirarchyid as Mstorgnhirarchyid,a.isclient as isclient,coalesce(a.recorddifftypeid,0) as recorddifftypeid,coalesce(a.recorddiffid,0) as recorddiffid,coalesce(a.prefix,'') as prefix,coalesce(a.year,'') as year,coalesce(a.month,'') as month,coalesce(a.day,'') as day, coalesce(a.configurezero,'') as configurezero ,coalesce(b.typename,'') as RecorddifftypeName,coalesce(c.name,'') as RecorddiffName,coalesce(d.name,'') as Clientname,coalesce(e.name,'') as Mstorgnhirarchyname,coalesce(f.number,0) as number FROM  mstrecordconfig a  left join mstrecorddifferentiationtype b on a.recorddifftypeid=b.id AND b.activeflg=1 AND b.deleteflg=0 left join mstrecorddifferentiation c on a.recorddiffid=c.id AND c.activeflg=1 AND c.deleteflg=0 join mstclient d on a.clientid=d.id join mstorgnhierarchy e on a.mstorgnhirarchyid=e.id left join mstrecordautoincreament f on a.clientid=f.clientid AND a.mstorgnhirarchyid=f.mstorgnhirarchyid AND a.recorddifftypeid=f.recorddifftypeid AND a.recorddiffid=f.recorddiffid AND f.activeflg=1 AND f.deleteflg=0 where a.activeflg=1 and a.deleteflg=0 AND a.clientid=? AND a.mstorgnhirarchyid=? ORDER BY a.id DESC LIMIT ?,?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
		params = append(params, tz.Offset)
		params = append(params, tz.Limit)
	}
	rows, err := dbc.DB.Query(getRecordConfigIncrement, params...)

	//rows, err := dbc.DB.Query(getRecordConfigIncrement, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllRecordConfigIncrement Get Statement Prepare Error", err)
		return values, err
	}
	//var groupid int64
	for rows.Next() {
		value := entities.RecordConfigIncrementEntity{}
		err = rows.Scan(&value.Id, &value.Clientid, &value.Mstorgnhirarchyid, &value.IsClient, &value.Recorddifftypeid, &value.Recorddiffid, &value.Prefix, &value.Year, &value.Month, &value.Day, &value.Configurezero, &value.RecorddifftypeName, &value.RecorddiffName, &value.ClientName, &value.Mstorgnhirarchyname, &value.Number)
		//value.Groupid=append(value.Groupid,groupid)
		if err != nil {
			logger.Log.Println("GetAllRecordConfigIncrement Get Statement Scan Error", err)
			return values, err
		}
		// logger.Log.Println(rows)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) UpdateRecordConfigWithoutTx(tz *entities.RecordConfigIncrementEntity) error {
	logger.Log.Println("In side UpdateRecordConfigIncrement")
	stmt, err := dbc.DB.Prepare(updateRecordConfigWithoutTx)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordConfigIncrement Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.Id, tz.IsClient)
	if err != nil {
		logger.Log.Println("UpdateRecordConfigIncrement Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc TxConn) UpdateRecordConfig(tz *entities.RecordConfigIncrementEntity) error {
	logger.Log.Println("In side UpdateRecordConfigIncrement")
	stmt, err := dbc.TX.Prepare(updateRecordConfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordConfigIncrement Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Prefix, tz.Year, tz.Month, tz.Day, tz.Configurezero, tz.Id, tz.IsClient)
	if err != nil {
		logger.Log.Println("UpdateRecordConfigIncrement Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc TxConn) UpdateRecordIncrement(tz *entities.RecordConfigIncrementEntity) error {
	logger.Log.Println("In side UpdateRecordConfigIncrement")
	stmt, err := dbc.TX.Prepare(updateRecordConfigincrement)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("UpdateRecordConfigIncrement Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Number, tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid, tz.IsClient)
	if err != nil {
		logger.Log.Println("UpdateRecordConfigIncrement Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) DeleteRecordConfigWithoutTx(tz *entities.RecordConfigIncrementEntity) error {
	logger.Log.Println("In side DeleteRecordConfigIncrement")
	stmt, err := dbc.DB.Prepare(deleteRecordConfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordConfigIncrement Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecordConfigIncrement Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc TxConn) DeleteRecordConfig(tz *entities.RecordConfigIncrementEntity) error {
	logger.Log.Println("In side DeleteRecordConfigIncrement")
	stmt, err := dbc.TX.Prepare(deleteRecordConfig)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordConfigIncrement Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Id)
	if err != nil {
		logger.Log.Println("DeleteRecordConfigIncrement Execute Statement  Error", err)
		return err
	}
	return nil
}
func (dbc TxConn) DeleteRecordIncrement(tz *entities.RecordConfigIncrementEntity) error {
	logger.Log.Println("In side DeleteRecordConfigIncrement")
	stmt, err := dbc.TX.Prepare(deleteRecordIncrement)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordConfigIncrement Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.Recorddifftypeid, tz.Recorddiffid)
	if err != nil {
		logger.Log.Println("DeleteRecordConfigIncrement Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetRecordConfigIncrementCount(tz *entities.RecordConfigIncrementEntity, OrgnTypeID int64) (entities.RecordConfigIncrementEntities, error) {
	logger.Log.Println("In side GetRecordConfigIncrementCount")
	value := entities.RecordConfigIncrementEntities{}
	var getRecordConfigIncrementcount string
	var params []interface{}
	if OrgnTypeID == 1 {
		getRecordConfigIncrementcount = "SELECT count(a.id) as total FROM mstrecordconfig a  left join mstrecorddifferentiationtype b on a.recorddifftypeid=b.id AND b.activeflg=1 AND b.deleteflg=0 left join mstrecorddifferentiation c on a.recorddiffid=c.id AND c.activeflg=1 AND c.deleteflg=0 join mstclient d on a.clientid=d.id join mstorgnhierarchy e on a.mstorgnhirarchyid=e.id left join mstrecordautoincreament f on a.clientid=f.clientid AND a.mstorgnhirarchyid=f.mstorgnhirarchyid AND a.recorddifftypeid=f.recorddifftypeid AND a.recorddiffid=f.recorddiffid AND f.activeflg=1 AND f.deleteflg=0 where a.activeflg=1 and a.deleteflg=0"
	} else if OrgnTypeID == 2 {
		getRecordConfigIncrementcount = "SELECT count(a.id) as total FROM mstrecordconfig a  left join mstrecorddifferentiationtype b on a.recorddifftypeid=b.id AND b.activeflg=1 AND b.deleteflg=0 left join mstrecorddifferentiation c on a.recorddiffid=c.id AND c.activeflg=1 AND c.deleteflg=0 join mstclient d on a.clientid=d.id join mstorgnhierarchy e on a.mstorgnhirarchyid=e.id left join mstrecordautoincreament f on a.clientid=f.clientid AND a.mstorgnhirarchyid=f.mstorgnhirarchyid AND a.recorddifftypeid=f.recorddifftypeid AND a.recorddiffid=f.recorddiffid AND f.activeflg=1 AND f.deleteflg=0 where a.activeflg=1 and a.deleteflg=0 AND a.clientid=? "
		params = append(params, tz.Clientid)
	} else {
		getRecordConfigIncrementcount = "SELECT count(a.id) as total FROM mstrecordconfig a  left join mstrecorddifferentiationtype b on a.recorddifftypeid=b.id AND b.activeflg=1 AND b.deleteflg=0 left join mstrecorddifferentiation c on a.recorddiffid=c.id AND c.activeflg=1 AND c.deleteflg=0 join mstclient d on a.clientid=d.id join mstorgnhierarchy e on a.mstorgnhirarchyid=e.id left join mstrecordautoincreament f on a.clientid=f.clientid AND a.mstorgnhirarchyid=f.mstorgnhirarchyid AND a.recorddifftypeid=f.recorddifftypeid AND a.recorddiffid=f.recorddiffid AND f.activeflg=1 AND f.deleteflg=0 where a.activeflg=1 and a.deleteflg=0 AND a.clientid=? AND a.mstorgnhirarchyid=?"
		params = append(params, tz.Clientid)
		params = append(params, tz.Mstorgnhirarchyid)
	}
	err := dbc.DB.QueryRow(getRecordConfigIncrementcount, params...).Scan(&value.Total)

	//err := dbc.DB.QueryRow(getRecordConfigIncrementcount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetRecordConfigIncrementCount Get Statement Prepare Error", err)
		return value, err
	}
}
