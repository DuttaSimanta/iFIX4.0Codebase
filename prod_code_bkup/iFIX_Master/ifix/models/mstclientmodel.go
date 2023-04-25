//***************************//
// Package models
// Date Of Creation: 18/12/2020
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: This file is used to do crud operation of mstclient table. It is used as model. All DB operation is defiend here.
// Functions: InsertMstClient,GetMstClientByID,GetMstClientAll,DelMstClientByID, UpdateMstClientByID
// InsertMstClient() Parameter:  (<*entities.MstClient>)
// GetMstClientByID() Parameter:  (<*entities.MstClient>)
// GetMstClientAll() Parameter:  (<*entities.MstClient>)
// DelMstClientByID() Parameter:  (<*entities.MstClient>)
// UpdateMstClientByID() Parameter:  (<*entities.MstClient>)
// Global Variable: N/A
// Version: 1.0.0
//***************************//

package models

import (
	"errors"
	"iFIX/ifix/config"
	"iFIX/ifix/entities"
	"log"
)

// InsertMstClient function is used to insert data in mstclient table using mysql query
func InsertMstClient(mstClient *entities.MstClient) (entities.MstClient, error) {
	var mstClientData entities.MstClient
	//dbCon, dbErr := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbCon, dbErr := config.ConnectMySqlDbSingleton()
	if dbErr != nil {
		log.Println(dbErr)
		return mstClientData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("Connection Error")
	}

	stmt, err := tx.Prepare("INSERT INTO mstclient (code, name, description, onboarddate, clientauditflg) VALUES (?,?,?,NOW(),?)")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(mstClient.Code, mstClient.Name, mstClient.Description, mstClient.ClientAuditFlg)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("SQL Execution Error")
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("SQL Commit Error")
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return mstClientData, errors.New("SQL Last ID fetch Error")
	}

	stmt.Close()
	//dbCon.Close()
	mstClientData.Id = lastInsertedID
	return mstClientData, nil

}

// GetMstClientByID function is used to fetch perticular row data by id from mstclient table  using mysql query
func GetMstClientByID(mstClient *entities.MstClient) (entities.MstClient, error) {
	var mstClientData entities.MstClient
	//dbCon, dbErr := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbCon, dbErr := config.ConnectMySqlDbSingleton()
	if dbErr != nil {
		log.Println(dbErr)
		return mstClientData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("Connection Error")
	}

	stmt, err := tx.Prepare("SELECT id,code, name, description, onboarddate, clientauditflg FROM mstclient WHERE id=?")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("SQL Prepare Error")
	}

	rows, err := stmt.Query(mstClient.Id)
	if err != nil {
		log.Println(err)
		return mstClientData, errors.New("SQL Execution Error")
	}

	for rows.Next() {
		rows.Scan(&mstClientData.Id, &mstClientData.Code, &mstClientData.Name, &mstClientData.Description, &mstClientData.OnboardDate, &mstClientData.ClientAuditFlg)
	}

	stmt.Close()
	//dbCon.Close()
	return mstClientData, nil

}

// GetMstClientAll function is used to all rows data from mstclient table  using mysql query
func GetMstClientAll() ([]entities.MstClient, error) {
	var mstClientData entities.MstClient
	var mstClientAllData []entities.MstClient

	//dbCon, dbErr := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbCon, dbErr := config.ConnectMySqlDbSingleton()
	if dbErr != nil {
		log.Println(dbErr)
		return mstClientAllData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientAllData, errors.New("Connection Error")
	}

	rows, err := tx.Query("SELECT id,code, name, description, onboarddate, clientauditflg FROM mstclient")
	if err != nil {
		log.Println(err)
		return mstClientAllData, errors.New("SQL Execution Error")
	}

	for rows.Next() {
		rows.Scan(&mstClientData.Id, &mstClientData.Code, &mstClientData.Name, &mstClientData.Description, &mstClientData.OnboardDate, &mstClientData.ClientAuditFlg)
		mstClientAllData = append(mstClientAllData, mstClientData)
	}

	//dbCon.Close()
	return mstClientAllData, nil

}

// DelMstClientByID function is used to delete perticular row data by id in mstclient table  using mysql query
func DelMstClientByID(mstClient *entities.MstClient) (entities.MstClient, error) {
	var mstClientData entities.MstClient
	//dbCon, dbErr := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbCon, dbErr := config.ConnectMySqlDbSingleton()
	if dbErr != nil {
		log.Println(dbErr)
		return mstClientData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("Connection Error")
	}

	stmt, err := tx.Prepare("DELETE FROM mstclient WHERE id=?")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(mstClient.Id)
	if err != nil {
		log.Println(err)
		return mstClientData, errors.New("SQL Execution Error")
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("SQL Commit Error")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return mstClientData, errors.New("SQL affected row fetch Error")
	}

	log.Println("Affected rows :", rowsAffected)

	stmt.Close()
	//dbCon.Close()
	return mstClientData, nil

}

// UpdateMstClientByID function is used to modify  data by id in mstclient table  using mysql query
func UpdateMstClientByID(mstClient *entities.MstClient) (entities.MstClient, error) {

	var mstClientData entities.MstClient

	//dbCon, dbErr := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbCon, dbErr := config.ConnectMySqlDbSingleton()
	if dbErr != nil {
		log.Println(dbErr)
		return mstClientData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("Connection Error")
	}

	stmt, err := tx.Prepare("UPDATE mstclient SET code=?,name=?,description=?,clientauditflg=? WHERE id=?")
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(mstClient.Code, mstClient.Name, mstClient.Description, mstClient.ClientAuditFlg, mstClient.Id)
	if err != nil {
		log.Println(err)
		return mstClientData, errors.New("SQL Execution Error")
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return mstClientData, errors.New("SQL Commit Error")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return mstClientData, errors.New("SQL affected row fetch Error")
	}

	log.Println("Affected rows :", rowsAffected)

	stmt.Close()
	//dbCon.Close()
	return mstClientData, nil

}
