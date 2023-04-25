package logger

import (
	"flag"
	"log"
	"os"
//	"time"
	"strings"
	//"path/filepath"
)

var (
	Log *log.Logger
)

func init() {
	// set location of log file
	//currentTime := time.Now() 
	//var presentDT string= currentTime.Format("02_Feb_2006_03_04_PM")
	log.Println("In Init Method of Logger")
	wd, err := os.Getwd() // to get working directory
	if err != nil {
        log.Println(err)
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd,"\\","/") // replacing backslash by  forwardslash
	logpath := contextPath+"/log/logFile_TCSICCIntegrationAPI.log"

	

	flag.Parse()
	var file, err1 = os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
	
}
