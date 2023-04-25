package model

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
	"log"
	"os"
	"src/fileutils"
	"src/logger"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = nil

func GetDB() (*sql.DB, error) {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/")
	//log.Println(contextPath)
	props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	connectionString := props["DBUser"] + ":" + props["DBPassword"] + "@" + "tcp(" + props["DBUrl"] + ":" + props["DBPort"] + ")/" + props["DBName"]
	logger.Log.Println(connectionString)
	if db == nil {
		logger.Log.Println("DB is Nil ", db)
		d, err := sql.Open(props["DBDriver"], connectionString)
		if err != nil {
			//logger.Log.Println(err.Error())
			return nil, err
		}
		db = d
	}
	logger.Log.Println("Database Open Connection Count is  ------------------>", db.Stats().OpenConnections)

	return db, nil
}
