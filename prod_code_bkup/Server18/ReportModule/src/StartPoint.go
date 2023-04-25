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
	router.NewRouter()
	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	Logger.Log.Println(contextPath)
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
