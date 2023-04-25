package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var termseq = "SELECT id,seq FROM mstrecordterms WHERE clientid=? AND mstorgnhirarchyid=? AND deleteflg=0 AND activeflg=1 AND seq is not null"

func (mdao DbConn) Termsequance(ClientID int64, Mstorgnhirarchyid int64) (map[int64]int64, error) {
	logger.Log.Println("In side Getfrequentissues---------------->", termseq)
	var t = make(map[int64]int64)
	var ID int64
	var Seq int64
	rows, err := mdao.DB.Query(termseq, ClientID, Mstorgnhirarchyid)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getfrequentissues Get Statement Prepare Error", err)
		return t, err
	}
	for rows.Next() {
		err = rows.Scan(&ID, &Seq)
		t[Seq] = ID
	}
	logger.Log.Println("Hashmap value is---------------->", t)
	return t, nil
}

func (dbc DbConn) Getstatetems(ClientID int64, OrgnID int64, TermID int64) ([]entities.MstStatetermEntity, error) {
	logger.Log.Println("In side Getstatetems")
	values := []entities.MstStatetermEntity{}
	var sql = "SELECT recorddifftypeid,recorddiffid,recordtermvalue,iscompulsory FROM mststateterm WHERE clientid=? AND mstorgnhirarchyid=? AND recordtermid=? AND deleteflg=0 AND activeflg=1"
	rows, err := dbc.DB.Query(sql, ClientID, OrgnID, TermID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("Getstatetems Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstStatetermEntity{}
		rows.Scan(&value.RecorddifftypeID, &value.RecorddiffID, &value.Recordtermvalue, &value.Iscompulsery)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetDiffID(FromclientID int64, FromorgnID int64, ToclinentID int64, ToorgnID int64, RecorddifftypeID int64, RecorddiffID int64) (int64, error) {
	logger.Log.Println("In side GetDiffID")
	var DiffID int64
	var sql = "Select id as DiffID FROM mstrecorddifferentiation WHERE clientid=? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND deleteflg=0 AND activeflg=1 AND seqno in (SELECT seqno FROM mstrecorddifferentiation WHERE clientid=? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND id=? AND deleteflg=0 AND activeflg=1)"
	rows, err := dbc.DB.Query(sql, ToclinentID, ToorgnID, RecorddifftypeID, FromclientID, FromorgnID, RecorddifftypeID, RecorddiffID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetDiffID Prepare Error", err)
		return DiffID, err
	}
	for rows.Next() {
		rows.Scan(&DiffID)

	}
	return DiffID, nil
}

func (dbc DbConn) Getgrpids(FromclientID int64, FromorgnID int64, TermID int64) ([]int64, error) {
	logger.Log.Println("In side GetDiffID")
	var GrpIDs []int64
	var sql = "SELECT grpid FROM mstsupportgrptermmap WHERE clientid=? AND mstorgnhirarchyid=? AND termid=? AND activeflg=1 AND deleteflg=0"
	rows, err := dbc.DB.Query(sql, FromclientID, FromorgnID, TermID)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetDiffID Prepare Error", err)
		return GrpIDs, err
	}
	for rows.Next() {
		var grpid int64
		err := rows.Scan(&grpid)
		logger.Log.Println("In side GetDiffID Error >>>>>>>>>>>>>>>>>>>>>", err)
		GrpIDs = append(GrpIDs, grpid)

	}
	return GrpIDs, nil
}

func (dbc DbConn) GetduplicateCheck(ToclientID int64, ToorgnID int64, Terseq int64) (int64, error) {
	logger.Log.Println("In side GetduplicateCheck")
	var Count int64
	var sql = "SELECT count(id) as Count FROM mstrecordterms WHERE clientid=? AND mstorgnhirarchyid=? AND seq=? AND activeflg=1 AND deleteflg=0"
	rows, err := dbc.DB.Query(sql, ToclientID, ToorgnID, Terseq)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetduplicateCheck Prepare Error", err)
		return Count, err
	}
	for rows.Next() {
		err := rows.Scan(&Count)
		logger.Log.Println("In side GetduplicateCheck Error >>>>>>>>>>>>>>>>>>>>>", err)
	}
	return Count, nil
}

func (dbc DbConn) GetTermdtls(FromclientID int64, FromorgnID int64, TermsSeq int64) (entities.MstRecordtermsmapEntity, error) {
	logger.Log.Println("In side GetDiffID")
	values := entities.MstRecordtermsmapEntity{}
	var sql = "SELECT id,termname,termtypeid,termvalue FROM mstrecordterms WHERE clientid=? AND mstorgnhirarchyid=? AND seq=? AND activeflg=1 AND deleteflg=0"
	rows, err := dbc.DB.Query(sql, FromclientID, FromorgnID, TermsSeq)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetDiffID Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		err := rows.Scan(&values.ID, &values.Termname, &values.Termtype, &values.Termvalue)
		logger.Log.Println("In side GetDiffID Error >>>>>>>>>>>>>>>>>>>>>", err)
	}
	return values, nil
}

func InsertRecordterms(tx *sql.Tx, ToclientID int64, ToorgnID int64, Termname string, Termvalue string, Termtype int64, Termseq int64) (int64, error) {
	logger.Log.Println("In side GetDiffID")
	var sql = "INSERT INTO mstrecordterms(clientid,mstorgnhirarchyid,termname,termtypeid,termvalue,seq) VALUES(?,?,?,?,?,?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		logger.Log.Println("GetDiffID Prepare Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ToclientID, ToorgnID, Termname, Termtype, Termvalue, Termseq)
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	return lastInsertedID, nil
}

func InsertRecordstateTerms(tx *sql.Tx, ToclientID int64, ToorgnID int64, DifftypeID int64, DiffID int64, TermID int64, Termvalue string, Iscompulsory int64) (int64, error) {
	logger.Log.Println("In side GetDiffID")
	var sql = "INSERT INTO mststateterm(clientid,mstorgnhirarchyid,recorddifftypeid,recorddiffid,recordtermid,recordtermvalue,iscompulsory) VALUES(?,?,?,?,?,?,?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		logger.Log.Println("GetDiffID Prepare Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ToclientID, ToorgnID, DifftypeID, DiffID, TermID, Termvalue, Iscompulsory)
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	return lastInsertedID, nil
}

func InsertSupportgrpTerms(tx *sql.Tx, ToclientID int64, ToorgnID int64, GrpID int64, TermID int64) (int64, error) {
	logger.Log.Println("In side GetDiffID")
	var sql = "INSERT INTO mstsupportgrptermmap(clientid,mstorgnhirarchyid,termid,grpid) VALUES(?,?,?,?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		logger.Log.Println("GetDiffID Prepare Error", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ToclientID, ToorgnID, TermID, GrpID)
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, err
	}
	return lastInsertedID, nil
}
