package dao

import (
	"database/sql"
	"errors"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
)

var clientinsert = "INSERT INTO mstclient(code,name,description,keyperson,keyemail,keymobile,baseflag,spocname,spocemail,spocnumber) VALUES(?,?,?,?,?,?,?,?,?,?)"
var clientdelete = "UPDATE mstclient SET baseflag=1 WHERE id=?"
var clientduplicate = "SELECT count(id) total FROM mstclient WHERE code=?"
var getclient = "SELECT id as ID,code as Code,name as Name,description as Description,keyperson as Keyperson,keyemail as Keyemail,keymobile as Keymobile,baseflag as Baseflag,spocname as Spocname,spocemail as Spocemail,spocnumber as Spocnumber FROM mstclient ORDER BY id DESC LIMIT ?,?"
var clientupdate = "UPDATE mstclient SET code=?,name=?,description=?,keyperson=?,keyemail=?,keymobile=?,baseflag=?,spocname=?,spocemail=?,spocnumber=? WHERE id=?"
var clientcount = "SELECT count(id) total FROM mstclient"

func (mdao DbConn) GetAllClientsnames() ([]entities.AllMstClientEntity, error) {
	values := []entities.AllMstClientEntity{}
	var sql = "SELECT a.id,a.name,b.id FROM mstclient a,mstorgnhierarchy b WHERE a.id=b.clientid AND b.parentid=0 order by a.name"
	rows, err := mdao.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllClientsnames Get Statement Prepare Error", err)
		logger.Log.Print("GetAllClientsnames Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.AllMstClientEntity{}
		rows.Scan(&value.ID, &value.Name, &value.OrgnID)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	//logger.Log.Println("values -->", values)
	return values, nil
}
//CheckDuplicateCient check duplicate record
func (mdao DbConn) CheckDuplicateCient(tz *entities.MstClientEntity) (entities.MstClientEntities, error) {
	logger.Log.Println("clientduplicate Query -->", clientduplicate)
	logger.Log.Println("parameters -->", tz.Code)
	value := entities.MstClientEntities{}
	err := mdao.DB.QueryRow(clientduplicate, tz.Code).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("clientduplicate Get Statement Prepare Error", err)
		return value, err
	}
}

//CheckDuplicateCient check duplicate record
func CheckDuplicateCientwithTX(tx *sql.Tx, tz *entities.MstClientEntity) (entities.MstClientEntities, error) {
	logger.Log.Println("clientduplicate Query -->", clientduplicate)
	logger.Log.Println("parameters -->", tz.Code)
	value := entities.MstClientEntities{}
	err := tx.QueryRow(clientduplicate, tz.Code).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("clientduplicate Get Statement Prepare Error", err)
		return value, err
	}
}

//InsertClientData data insertd in mstclientuser table
func (mdao DbConn) InsertClientData(data *entities.MstClientEntity) (int64, error) {
	logger.Log.Println("clientinsert query -->", clientinsert)
	logger.Log.Println("parameters -->", data.Code, data.Name, data.Description, data.Keyperson, data.Keyemail, data.Keymobile, data.Baseflag, data.Spocname, data.Spocemail, data.Spocnumber)
	stmt, err := mdao.DB.Prepare(clientinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertClientData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(data.Code, data.Name, data.Description, data.Keyperson, data.Keyemail, data.Keymobile, data.Baseflag, data.Spocname, data.Spocemail, data.Spocnumber)
	if err != nil {
		logger.Log.Print("InsertClientData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//InsertClientData data insertd in mstclientuser table
func InsertClientDatawithTX(tx *sql.Tx, data *entities.MstClientEntity) (int64, error) {
	logger.Log.Println("clientinsert query -->", clientinsert)
	logger.Log.Println("parameters -->", data.Code, data.Name, data.Description, data.Keyperson, data.Keyemail, data.Keymobile, data.Baseflag, data.Spocname, data.Spocemail, data.Spocnumber)
	stmt, err := tx.Prepare(clientinsert)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("InsertClientData Prepare Statement  Error", err)
		return 0, err
	}
	res, err := stmt.Exec(data.Code, data.Name, data.Description, data.Keyperson, data.Keyemail, data.Keymobile, data.Baseflag, data.Spocname, data.Spocemail, data.Spocnumber)
	if err != nil {
		logger.Log.Print("InsertClientData Execute Statement  Error", err)
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		logger.Log.Println(err)
		return 0, errors.New("SQL Last ID fetch Error")
	}
	return lastInsertedID, nil
}

//UpdateClientData update mstclientuser table
func (mdao DbConn) UpdateClientData(data *entities.MstClientEntity) error {
	logger.Log.Println("userupdate Query -->", clientupdate)
	logger.Log.Println("parameters -->", data.Code, data.Name, data.Description, data.Keyperson, data.Keyemail, data.Keymobile, data.Baseflag, data.Spocname, data.Spocemail, data.Spocnumber, data.ID)
	stmt, err := mdao.DB.Prepare(clientupdate)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Update Client Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(data.Code, data.Name, data.Description, data.Keyperson, data.Keyemail, data.Keymobile, data.Baseflag, data.Spocname, data.Spocemail, data.Spocnumber, data.ID)
	if err != nil {
		logger.Log.Print("Update Client Execute Statement  Error", err)
		return err
	}
	return nil
}

//DeleteClientData update mstclientuser table
func (mdao DbConn) DeleteClientData(tz *entities.MstClientEntity) error {
	logger.Log.Println("clientdelete Query -->", clientdelete)
	logger.Log.Println("parameters -->", tz.ID)

	stmt, err := mdao.DB.Prepare(clientdelete)
	defer stmt.Close()
	if err != nil {
		logger.Log.Print("Delete Client Prepare Statement  Error", err)
		return err
	}
	_, err = stmt.Exec(tz.ID)
	if err != nil {
		logger.Log.Print("Delete Client Execute Statement  Error", err)
		return err
	}
	return nil
}

//GetClientCount get user count with condition
func (mdao DbConn) GetClientCount() (entities.MstClientEntities, error) {
	logger.Log.Println("usergetcount Query -->", clientcount)
	value := entities.MstClientEntities{}
	err := mdao.DB.QueryRow(clientcount).Scan(&value.Total)
	switch err {
	case sql.ErrNoRows:
		value.Total = 0
		return value, nil
	case nil:
		return value, nil
	default:
		logger.Log.Print("GetClientCount Get Statement Prepare Error", err)
		log.Print("GetClientCount Get Statement Prepare Error", err)
		return value, err
	}
}

//GetAllClients get user count with condition
func (mdao DbConn) GetAllClients(page *entities.MstClientEntity) ([]entities.MstClientEntity, error) {
	logger.Log.Println("getclient Query -->", getclient)
	logger.Log.Println("parameters -->", page.Offset, page.Limit)
	values := []entities.MstClientEntity{}
	rows, err := mdao.DB.Query(getclient, page.Offset, page.Limit)
	defer rows.Close()
	if err != nil {
		log.Print("GetAllClients Get Statement Prepare Error", err)
		logger.Log.Print("GetAllClients Get Statement Prepare Error", err)
		return values, err
	}
	for rows.Next() {
		value := entities.MstClientEntity{}
		rows.Scan(&value.ID, &value.Code, &value.Name, &value.Description, &value.Keyperson, &value.Keyemail, &value.Keymobile, &value.Baseflag, &value.Spocname, &value.Spocemail, &value.Spocnumber)
		logger.Log.Println("value -->", value)
		values = append(values, value)
	}
	logger.Log.Println("values -->", values)
	return values, nil
}
