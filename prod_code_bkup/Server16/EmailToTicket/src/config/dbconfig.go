package config

import (
	"database/sql"
	"log"
	"os"
	ReadProperties "src/fileutils"
	Logger "src/logger"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var mutex = &sync.Mutex{}
var dbConn *sql.DB = nil

func GetDB() (db *sql.DB, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	if dbConn == nil {
		Logger.Log.Println("dbConn======> ", dbConn)
		//calling properties file
		wd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		//log.Println(wd)
		contextPath := strings.ReplaceAll(wd, "\\", "/")
		//log.Println(contextPath)
		props, err := ReadProperties.ReadPropertiesFile(contextPath + "/resource/application.properties")
		connectionString := props["DBUser"] + ":" + props["DBPassword"] + "@" + "tcp(" + props["DBUrl"] + ":" + props["DBPort"] + ")/" + props["DBName"]
		Logger.Log.Println(connectionString)

		db, err = sql.Open(props["DBDriver"], connectionString)
		Logger.Log.Println("Connection Object=====>", db)
		if err != nil {
			Logger.Log.Fatalln(err.Error())
			return nil, err
		} else {
			Logger.Log.Println("DB Connected!!!")
		}
		dbConn = db
	}
	Logger.Log.Println("DbConn in Else condition =======>", dbConn)
	return dbConn, nil
}

// package config

// import (
// 	"database/sql"
// 	"log"
// 	"os"
// 	ReadProperties "src/fileutils"
// 	"strings"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func GetDB() (db *sql.DB, err error) {

// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	//log.Println(wd)
// 	contextPath := strings.ReplaceAll(wd, "\\", "/")
// 	//log.Println(contextPath)
// 	props, err := ReadProperties.ReadPropertiesFile(contextPath + "/resource/application.properties")
// 	connectionString := props["DBUser"] + ":" + props["DBPassword"] + "@" + "tcp(" + props["DBUrl"] + ":" + props["DBPort"] + ")/" + props["DBName"]
// 	log.Println(connectionString)
// 	db, err = sql.Open(props["DBDriver"], connectionString)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	} else {
// 		log.Println("DB Connected!!!")
// 	}
// 	return
// }
