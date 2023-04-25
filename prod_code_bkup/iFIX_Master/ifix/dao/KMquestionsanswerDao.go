package dao

import (
	"database/sql"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
)

var insertquestionanswer = "INSERT INTO mstrecorddifferentiation(clientid,mstorgnhirarchyid,recorddifftypeid,name) values (?,?,?,?)"
var insertrecordtype = "INSERT INTO mstrecordtype(clientid,mstorgnhirarchyid,fromrecorddifftypeid,fromrecorddiffid,torecorddifftypeid,torecorddiffid) values(?,?,?,?,?,?)"
var checkduplicatequestions = "SELECT count(id) total from mstrecorddifferentiation WHERE clientid=? AND mstorgnhirarchyid=? AND recorddifftypeid=? AND name=? AND deleteflg = 0 AND activeflg=1"
var deletequestionanswer = "UPDATE mstrecorddifferentiation SET deleteflg=1 WHERE id=?"
var deleterecordmap = "UPDATE mstrecordtype SET deleteflg=1 WHERE id=?"
var getquestionanswer = "SELECT a.id AS Id,a.clientid AS Clientid,a.mstorgnhirarchyid AS Mstorgnhirarchyid, a.fromrecorddifftypeid AS Fromrecorddifftypeid, a.fromrecorddiffid AS Fromrecorddiffid, a.torecorddifftypeid AS Torecorddifftypeid, a.torecorddiffid AS Torecorddiffid, a.activeflg AS Activeflg , d.typename AS Fromrecorddifftypename, (select name from mstrecorddifferentiation where id = a.fromrecorddiffid) AS FromrecorddiffName, (select typename from mstrecorddifferentiationtype where id = a.torecorddifftypeid) AS Torecorddifftypename, (select name from mstrecorddifferentiation where id = a.torecorddiffid) AS TorecorddiffName,b.name as Clientname,c.name as Mstorgnhirarchyname FROM mstrecordtype a , mstclient b,mstorgnhierarchy c ,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg = 0 AND a.activeflg = 1 AND d.seqno = 8 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.fromrecorddifftypeid = d.id ORDER BY a.id DESC LIMIT ?,?"
var questionanswercount = "SELECT count(a.id) total FROM mstrecordtype a , mstclient b,mstorgnhierarchy c ,mstrecorddifferentiationtype d WHERE a.clientid = ? AND a.mstorgnhirarchyid = ? AND a.deleteflg = 0 AND a.activeflg = 1 AND d.seqno = 8 AND a.clientid = b.id AND a.mstorgnhirarchyid = c.id AND a.fromrecorddifftypeid = d.id"

func CheckDuplicateQuestions(tx *sql.Tx, tz *entities.KMquestionsanswerEntity) (entities.KMquestionsanswerEntities, error) {
	logger.Log.Println("In side CheckDuplicateQuestions")
	value := entities.KMquestionsanswerEntities{}
	err := tx.QueryRow(checkduplicatequestions, tz.Clientid, tz.Mstorgnhirarchyid, tz.QuestiondifftypeID, tz.Questions).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("CheckDuplicateQuestions Get Statement Prepare Error", err)
		return value, err
	}
}

func InsertQuestions(tx *sql.Tx, tz *entities.KMquestionsanswerEntity) (int64, error) {
	logger.Log.Println("In side InsertQuestions")
	logger.Log.Println("Query -->", insertquestionanswer)
	stmt, err := tx.Prepare(insertquestionanswer)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertQuestions Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.QuestiondifftypeID, tz.Questions)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.QuestiondifftypeID, tz.Questions)
	if err != nil {
		logger.Log.Println("InsertQuestions Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func InsertAnswers(tx *sql.Tx, tz *entities.KMquestionsanswerEntity) (int64, error) {
	logger.Log.Println("In side InsertAnswers")
	logger.Log.Println("Query -->", insertquestionanswer)
	stmt, err := tx.Prepare(insertquestionanswer)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertAnswers Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.AnswerdifftypeID, tz.Answer)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.AnswerdifftypeID, tz.Answer)
	if err != nil {
		logger.Log.Println("InsertAnswers Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func InsertRecordmap(tx *sql.Tx, tz *entities.KMquestionsanswerEntity, lastinsertedQID int64, lastinsertedAID int64) (int64, error) {
	logger.Log.Println("In side InsertAnswers")
	logger.Log.Println("Query -->", insertrecordtype)
	stmt, err := tx.Prepare(insertrecordtype)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("InsertAnswers Prepare Statement  Error", err)
		return 0, err
	}
	logger.Log.Println("Parameter -->", tz.Clientid, tz.Mstorgnhirarchyid, tz.QuestiondifftypeID, lastinsertedQID, tz.AnswerdifftypeID, lastinsertedAID)
	res, err := stmt.Exec(tz.Clientid, tz.Mstorgnhirarchyid, tz.QuestiondifftypeID, lastinsertedQID, tz.AnswerdifftypeID, lastinsertedAID)
	if err != nil {
		logger.Log.Println("InsertAnswers Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return lastInsertedId, nil
}

func DeleteQuestions(tx *sql.Tx, tz *entities.KMquestionsanswerEntity) error {
	logger.Log.Println("In side DeleteQuestions")
	stmt, err := tx.Prepare(deletequestionanswer)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteQuestions Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.QuestiondiffID)
	if err != nil {
		logger.Log.Println("DeleteQuestions Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteAnswers(tx *sql.Tx, tz *entities.KMquestionsanswerEntity) error {
	logger.Log.Println("In side DeleteQuestions")
	stmt, err := tx.Prepare(deletequestionanswer)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteAnswers Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.AnswerdiffID)
	if err != nil {
		logger.Log.Println("DeleteAnswers Execute Statement  Error", err)
		return err
	}
	return nil
}

func DeleteRecordmap(tx *sql.Tx, tz *entities.KMquestionsanswerEntity) error {
	logger.Log.Println("In side DeleteRecordmap")
	stmt, err := tx.Prepare(deleterecordmap)
	defer stmt.Close()
	if err != nil {
		logger.Log.Println("DeleteRecordmap Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Println("DeleteRecordmap Execute Statement  Error", err)
		return err
	}
	return nil
}

func (dbc DbConn) GetAllQuestionanswer(page *entities.KMquestionsanswerEntity) ([]entities.KMquestionsanswerEntity, error) {
	logger.Log.Println("In side GelAllClientsupportgroup")
	logger.Log.Println(getquestionanswer)
	values := []entities.KMquestionsanswerEntity{}
	rows, err := dbc.DB.Query(getquestionanswer, page.Clientid, page.Mstorgnhirarchyid, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		logger.Log.Println("GetAllClientsupportgroup Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.KMquestionsanswerEntity{}
		rows.Scan(&value.ID, &value.Clientid, &value.Mstorgnhirarchyid, &value.QuestiondifftypeID, &value.QuestiondiffID, &value.AnswerdifftypeID, &value.AnswerdiffID, &value.Activeflg, &value.Questionheader, &value.Questions, &value.Answerheader, &value.Answer, &value.Clientname, &value.Mstorgnhirarchyname)
		values = append(values, value)
	}
	return values, nil
}

func (dbc DbConn) GetQuestionanswerscount(tz *entities.KMquestionsanswerEntity) (entities.KMquestionsanswerEntities, error) {
	logger.Log.Println("In side GetClientsupportgroupCount")
	value := entities.KMquestionsanswerEntities{}
	err := dbc.DB.QueryRow(questionanswercount, tz.Clientid, tz.Mstorgnhirarchyid).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Println("GetClientsupportgroupCount Get Statement Prepare Error", err)
		return value, err
	}
}
