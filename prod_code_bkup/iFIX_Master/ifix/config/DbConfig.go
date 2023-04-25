package config
//***************************//
// Package Name: config
// Date Of Creation: 17/12/2020
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: Database configuration file with connection
// Functions: ConnectMySqlDb
// Inputs: <*sql.DB>, <error>
// Global Variable: N/A
// Version: 1.0.0
//***************************//

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


//ConnectMySqlDb is used for db connection
func ConnectMySqlDb() (db *sql.DB, err error) {
	db, err = sql.Open(DBDRIVER, DBUSER+":"+DBPASWORD+"@"+DBURL+"/"+DBNAME)
	return
}
var dbconn *sql.DB = nil

func ConnectMySqlDbSingleton() (*sql.DB, error) {
	dbDriver := DBDRIVER          // Database Driver Name
	dbUser := DBUSER            // Database Username
	dbPassword := DBPASWORD       // Database  Password
	dbUrl := DBURL // Database ip/host with port
	dbName := DBNAME             // Database Name
	if dbconn == nil {
		d, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@"+dbUrl+"/"+dbName)
		if err != nil {
			// panic(err.Error())
			return nil, err
		}
		dbconn = d
	}
	return dbconn, nil
}
//ConnectMySqlDb is used for db connection
//func ConnectMySqlDb() (db *sql.DB, err error) {
//	dbDriver := "mysql" // Database Driver Name
//	dbUser := "ifix" // Database Username
//	dbPassword := "Staging@4321" // Database  Password
//	dbUrl := "tcp(172.17.0.1:3306)" // Database ip/host with port
//	dbName := "iFIX" // Database Name
//	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@"+dbUrl+"/"+dbName)
//	return
//}
