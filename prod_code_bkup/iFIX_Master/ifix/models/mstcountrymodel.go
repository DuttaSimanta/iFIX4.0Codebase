//***************************//
// Package models
// Date Of Creation: 18/12/2020
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: This file is used to do crud operation of mstcountry table. It is used as model. All DB operation is defiend here.
// Functions: InsertMstCountry,GetMstCountryByID,GetMstCountryAll,DelMstCountryByID, UpdateMstCountryByID
// InsertMstCountry() Parameter:  (<*entities.MstCountry>)
// GetMstCountryByID() Parameter:  (<*entities.MstCountry>)
// GetMstCountryAll() Parameter:  (<*entities.MstCountry>)
// DelMstCountryByID() Parameter:  (<*entities.MstCountry>)
// UpdateMstCountryByID() Parameter:  (<*entities.MstCountry>)
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

// InsertMstCountry function is used to insert data in mstcountry table using mysql query
func InsertMstCountry(mstCountry *entities.MstCountry) (entities.MstCountry, error) {
	var mstCountryData entities.MstCountry
	//dbCon, dbErr := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbCon, dbErr := config.ConnectMySqlDbSingleton()
	if dbErr != nil {
		log.Println(dbErr)
		return mstCountryData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("Connection Error")
	}

	stmt, err := tx.Prepare("INSERT INTO mstcountry (countrycode, countryname) VALUES (?,?)")
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(mstCountry.CountryCode, mstCountry.CountryName)
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Execution Error")
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Commit Error")
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Last ID fetch Error")
	}

	stmt.Close()
	// dbCon.Close()
	mstCountryData.Id = lastInsertedID
	return mstCountryData, nil

}

// GetMstCountryByID function is used to fetch perticular row data by id from mstcountry table  using mysql query
func GetMstCountryByID(mstCountry *entities.MstCountry) (entities.MstCountry, error) {
	var mstCountryData entities.MstCountry

	//dbCon, dbErr := config.ConnectMySqlDb()
	lock.Lock()
	defer lock.Unlock()
	dbCon, dbErr := config.ConnectMySqlDbSingleton()
	if dbErr != nil {
		log.Println(dbErr)
		return mstCountryData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()

	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("Connection Error")
	}

	stmt, err := tx.Prepare("SELECT id,countrycode, countryname FROM mstcountry WHERE id=?")
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Prepare Error")
	}

	rows, err := stmt.Query(mstCountry.Id)
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Execution Error")
	}

	for rows.Next() {
		rows.Scan(&mstCountryData.Id, &mstCountryData.CountryCode, &mstCountryData.CountryName)
	}

	stmt.Close()
	// dbCon.Close()
	return mstCountryData, nil

}

// GetMstCountryAll function is used to all rows data from mstcountry table  using mysql query
func GetMstCountryAll() ([]entities.MstCountry, error) {
	var mstCountryData entities.MstCountry
	var mstCountryAllData []entities.MstCountry

	dbCon, dbErr := config.ConnectMySqlDb()
	if dbErr != nil {
		log.Println(dbErr)
		return mstCountryAllData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)
		return mstCountryAllData, errors.New("Connection Error")
	}

	rows, err := tx.Query("SELECT id,countrycode,countryname FROM mstcountry")
	if err != nil {
		log.Println(err)
		return mstCountryAllData, errors.New("SQL Execution Error")
	}

	for rows.Next() {
		rows.Scan(&mstCountryData.Id, &mstCountryData.CountryCode, &mstCountryData.CountryName)
		mstCountryAllData = append(mstCountryAllData, mstCountryData)
	}

	// dbCon.Close()
	return mstCountryAllData, nil

}

// DelMstCountryByID function is used to delete perticular row data by id in mstcountry table  using mysql query
func DelMstCountryByID(mstCountry *entities.MstCountry) (entities.MstCountry, error) {
	var mstCountryData entities.MstCountry
	dbCon, dbErr := config.ConnectMySqlDb()
	if dbErr != nil {
		log.Println(dbErr)
		return mstCountryData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)

		return mstCountryData, errors.New("Connection Error")
	}

	stmt, err := tx.Prepare("DELETE FROM mstcountry WHERE id=?")
	if err != nil {
		log.Println(err)

		return mstCountryData, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(mstCountry.Id)
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Execution Error")
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Commit Error")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL affected row fetch Error")
	}

	log.Println("Affected rows :", rowsAffected)

	stmt.Close()
	// dbCon.Close()
	return mstCountryData, nil

}

// UpdateMstCountryByID function is used to modify  data by id in mstcountry table  using mysql query
func UpdateMstCountryByID(mstCountry *entities.MstCountry) (entities.MstCountry, error) {
	var mstCountryData entities.MstCountry

	dbCon, dbErr := config.ConnectMySqlDb()
	if dbErr != nil {
		log.Println(dbErr)
		return mstCountryData, errors.New("Connection Intialization Error")
	}
	tx, err := dbCon.Begin()
	if err != nil {
		log.Println(err)

		return mstCountryData, errors.New("Connection Error")
	}

	stmt, err := tx.Prepare("UPDATE mstcountry SET code=?,name=?,description=?,clientauditflg=? WHERE id=?")
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Prepare Error")
	}

	res, err := stmt.Exec(mstCountry.CountryCode, mstCountry.CountryName)
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Execution Error")
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL Commit Error")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return mstCountryData, errors.New("SQL affected row fetch Error")
	}

	log.Println("Affected rows :", rowsAffected)

	stmt.Close()
	// dbCon.Close()
	return mstCountryData, nil

}
