package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	ReadProperties "src/fileutils"
	"log"
	"os"
	"strings"
)

func GetDB() (db *sql.DB, err error) {

	wd, err := os.Getwd()
	if err != nil {
        log.Println(err)
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd,"\\","/")
	//log.Println(contextPath)
	props, err := ReadProperties.ReadPropertiesFile(contextPath+"/resource/application.properties")
	connectionString :=props["DBUser"] + ":" + props["DBPassword"] + "@" + "tcp(" + props["DBUrl"] + ":" + props["DBPort"] + ")/"+ props["DBName"]
	log.Println(connectionString)
	db, err = sql.Open(props["DBDriver"], connectionString)
	if err != nil {
        log.Fatalln(err.Error())
    }else{
		log.Println("DB Connected!!!")
	}
	return
}