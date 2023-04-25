package config

import (
	"database/sql"
	"log"
	"os"
	"src/fileutils"
	Logger "src/logger"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var lock = &sync.Mutex{}
var dbConn *sql.DB = nil

func GetDB() (db *sql.DB, err error) {
	// mutex.Lock()
	// defer mutex.Unlock()
	if fileutils.MutexLocked(lock) == false {
		lock.Lock()
		defer lock.Unlock()
	}
	log.Println("DB===?>", dbConn)
	if dbConn == nil {
		log.Println("DB===test", dbConn)
		Logger.Log.Println("dbConn======> ", dbConn)
		//calling properties file
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
		Logger.Log.Println("Connection Object=====>", db)
		if err != nil {
			Logger.Log.Println(err.Error())
			return nil, err
		} else {
			Logger.Log.Println("DB Connected!!!")
		}
		dbConn = db
	}
	Logger.Log.Println("DbConn in Else condition =======>", dbConn)
	return dbConn, nil
}
