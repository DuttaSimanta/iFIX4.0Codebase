package models

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

	Logger "src/logger"

	"strings"

	"sync"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB

	once sync.Once
)

func ConnectMySqlDb() (*sql.DB, error) {

	once.Do(func() {

		var err error

		// dbDriver := "mysql"           // Database Driver Name

		// dbUser := "gouser"            // Database Username

		// dbPassword := "TCSUAT@54321"  // Database  Password

		// dbUrl := "tcp(10.5.2.4:3306)" // Database ip/host with port

		// dbName := "iFIX"              // Database Name

		// dbDriver := "mysql"             // Database Driver Name

		// dbUser := "ifix"                // Database Username

		// dbPassword := "Staging@4321"    // Database  Password

		// dbUrl := "tcp(172.17.0.1:3306)" // Database ip/host with port

		// dbName := "iFIX"                // Database Name

		// db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@"+dbUrl+"/"+dbName)

		wd, err := os.Getwd()

		if err != nil {

			log.Println(err)

		}

		//log.Println(wd)

		contextPath := strings.ReplaceAll(wd, "\\", "/")

		//log.Println(contextPath)

		props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")

		connectionString := props["DBUser"] + ":" + props["DBPassword"] + "@" + "tcp(" + props["DBUrl"] + ":" + props["DBPort"] + ")/" + props["DBName"]

		Logger.Log.Println(connectionString)

		db, err = sql.Open(props["DBDriver"], connectionString)

		if err != nil {

			panic(err.Error())

			//return nil, err

		}

		db.SetMaxIdleConns(0)

		db.SetMaxOpenConns(150)

		db.SetConnMaxLifetime(time.Millisecond * 100)

	})

	return db, nil

}
