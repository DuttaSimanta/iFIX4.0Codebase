package main

import (
	"fmt"
	"net/http"
	"os"
	ReadProperties "src/fileutils"
	Logger "src/logger"
	"src/router"
	"strings"
	//Utils "src/fileutils"
	//"log"
)

func main() {
	Logger.Log.Println("Main Started")
	// db, dBerr := config.ConnectMySqlDb()
	// if dBerr != nil {
	// 	Logger.Log.Fatal(dBerr)
	// 	//return ticketNo, errors.New("Unable to connect")
	// }
	// // Maximum Idle Connections
	// db.SetMaxIdleConns(100)
	// // Maximum Open Connections
	// db.SetMaxOpenConns(200)
	// // Idle Connection Timeout
	// db.SetConnMaxIdleTime(1 * time.Second)
	// // Connection Lifetime
	// db.SetConnMaxLifetime(30 * time.Second)
	router.NewRouter()
	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
	}

	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	//log.Println(contextPath)
	props, err := ReadProperties.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
	}
	Logger.Log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", props["SERVERPORT"]), nil))
}

/* func main() {
	Logger.Log.Println("Main Started")
	/*err := Utils.SendMail("hi","test","kaustubh@ifixtechglobal.com")
		if err != nil {
	        log.Println(err)
		}

	Routes.Handle()
} */
