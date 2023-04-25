package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

func (mdao DbConn) GetDifferentiationname(DiffID int64) (string, string, error) {
	logger.Log.Println("In side GetDifferentiationname")
	var name string
	var seq string
	var sql = "SELECT name,seqno FROM mstrecorddifferentiation WHERE id=?"
	rows, err := mdao.DB.Query(sql, DiffID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetDifferentiationname Get Statement Prepare Error", err)
		return name, seq, err
	}
	for rows.Next() {
		err = rows.Scan(&name, &seq)
		logger.Log.Println("GetDifferentiationname rows.next() Error", err)
	}
	return name, seq, nil
}

func (mdao DbConn) DuplicateChecking(ClientID int64, OrgnID int64, DifftypeID int64, Name string) (int64, error) {
	logger.Log.Println("In side DuplicateChecking")
	var total int64
	var sql = "SELECT count(*) total FROM mstrecorddifferentiation WHERE clientid=? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND name=? AND deleteflg=0 AND activeflg=1"
	rows, err := mdao.DB.Query(sql, ClientID, OrgnID, DifftypeID, Name)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("DuplicateChecking Get Statement Prepare Error", err)
		return total, err
	}
	for rows.Next() {
		err = rows.Scan(&total)
		logger.Log.Println("DuplicateChecking rows.next() Error", err)
	}
	return total, nil
}

func InsertDifferentiationTBL(tx *sql.Tx, ClientID int64, OrgnID int64, DifftypeID int64, Name string, Seq string) (int64, error) {
	logger.Log.Println("parameters -->", ClientID, OrgnID, DifftypeID, Name, Seq)
	var sql = "INSERT INTO mstrecorddifferentiation(clientid,mstorgnhirarchyid,recorddifftypeid,name,seqno) VALUES(?,?,?,?,?)"
	stmt, err := tx.Prepare(sql)

	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Prepare Error")
	}
	defer stmt.Close()
	res, err := stmt.Exec(ClientID, OrgnID, DifftypeID, Name, Seq)
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Execution Error")
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}

	return lastInsertedID, nil
}

func InsertDifferentiationMAPTBL(tx *sql.Tx, FromClientID int64, FromOrgnID int64, FromDifftypeID int64, FromDiffID int64, ToClientID int64, ToOrgnID int64, ToDifftypeID int64, ToDiffID int64) error {
	logger.Log.Println("parameters -->", FromClientID, FromOrgnID, FromDifftypeID, FromDiffID, ToClientID, ToOrgnID, ToDifftypeID, ToDiffID)
	var sql = "INSERT INTO mstrecorddifferentiationmap(fromclientid,frommstorgnhirarchyid,fromrecorddifftypeid,fromrecorddiffid,toclientid,tomstorgnhirarchyid,torecorddifftypeid,torecorddiffid) VALUES(?,?,?,?,?,?,?,?)"
	stmt, err := tx.Prepare(sql)

	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Prepare Error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(FromClientID, FromOrgnID, FromDifftypeID, FromDiffID, ToClientID, ToOrgnID, ToDifftypeID, ToDiffID)
	if err != nil {
		logger.Log.Println(err)
		return errors.New("SQL Execution Error")
	}
	return nil
}

func DeleteFromDifferentiation(tx *sql.Tx, DiffID int64) error {
	logger.Log.Println("parameters -->", DiffID)
	var sql = "Update mstrecorddifferentiation SET deleteflg=1 WHERE id=?"
	stmt, err := tx.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("DeleteFromDifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(DiffID)
	if err != nil {
		logger.Log.Print("DeleteFromDifferentiation Statement  Error", err)
		return err
	}
	return nil
}

func DeleteFromDifferentiationMap(tx *sql.Tx, MapID int64) error {
	logger.Log.Println("parameters -->", MapID)
	var sql = "Update mstrecorddifferentiationmap SET deleteflg=1 WHERE id=?"
	stmt, err := tx.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("DeleteFromDifferentiation Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(MapID)
	if err != nil {
		logger.Log.Print("DeleteFromDifferentiation Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetAllDifferentiationMapDtls(page *entities.MstDifferentiationmapEntity) ([]entities.MstDifferentiationmapResponseFields, error) {
	logger.Log.Println("In side GetAllDifferentiationMapDtls")
	values := []entities.MstDifferentiationmapResponseFields{}
	var sql = "SELECT a.torecorddiffid as ID,a.id as MapID,b.name as Fromclientname,c.name as Fromorgnname,d.typename as Fromdifftypename,e.name as Fromdiffname,f.name as Toclientname,g.name as Toorgnname,h.typename as Todifftypename,i.name as Todiffname  FROM mstrecorddifferentiationmap a,mstclient b,mstorgnhierarchy c,mstrecorddifferentiationtype d,mstrecorddifferentiation e,mstclient f,mstorgnhierarchy g,mstrecorddifferentiationtype h,mstrecorddifferentiation i WHERE a.fromclientid=? AND a.deleteflg=0 AND a.activeflg=1 AND a.fromclientid=b.id AND a.frommstorgnhirarchyid=c.id AND a.fromrecorddifftypeid=d.id AND a.fromrecorddiffid=e.id AND a.toclientid=f.id AND a.tomstorgnhirarchyid =g.id AND a.torecorddifftypeid=h.id AND a.torecorddiffid=i.id ORDER BY a.id DESC LIMIT ?,?"
	rows, err := dbc.DB.Query(sql, page.FromclientID, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllDifferentiationMapDtls Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstDifferentiationmapResponseFields{}
		rows.Scan(&value.ID, &value.MapID, &value.Fromclientname, &value.Fromorgnname, &value.Fromdifftypename, &value.Fromdiffname, &value.Toclientname, &value.Toorgnname, &value.Todifftypename, &value.Todiffname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetAllDifferentiationMapDtlsCount(tz *entities.MstDifferentiationmapEntity) (entities.MstDifferentiationmaEntities, error) {
	logger.Log.Println("In side GetAllDifferentiationMapDtlsCount")
	value := entities.MstDifferentiationmaEntities{}
	var query = "SELECT count(id) as count FROM mstrecorddifferentiationmap WHERE fromclientid=? AND deleteflg=0 AND activeflg=1"
	err := dbc.DB.QueryRow(query, tz.FromclientID).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetAllDifferentiationMapDtlsCount Get Statement Prepare Error", err)
		return value, err
	}
}
