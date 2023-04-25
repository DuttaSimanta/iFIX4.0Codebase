package logger

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	logpath := contextPath + "/log/logFile_AutoCloseTicketScheduler.log"

	//ar logpath = "log/logfile.log"
	removeFilePath := contextPath + "/log"

	d, err := os.Open(removeFilePath)
	if err != nil {
		panic(err)
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		panic(err)
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(removeFilePath, name))
		if err != nil {
			panic(err)
		}
	}

	flag.Parse()
	var file, err1 = os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)

}
