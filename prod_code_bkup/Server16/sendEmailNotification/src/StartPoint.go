package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"src/config"
	ReadProperties "src/fileutils"
	Logger "src/logger"
	"src/router"
	"strings"
	//"log"
)

var dbConn *sql.DB = nil

func main() {
	Logger.Log.Println("Main Started")
	dbConn, dBerr := config.GetDB()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		//return dBerr
	}
	Logger.Log.Println(dbConn)
	// SendSMS(indiaDltContentTemplateId string, mobileNo string, smsContent string, smsType string,
	// 	SMSUrlString string, IndiaDltPrincipalEntityIdVal string, smsUserName string, smsPassword string)

	// indiaDltContentTemplateId := "1107163282950818969"
	// mobileNo := "8013225433,8961329244,9875642426"
	// IndiaDltPrincipalEntityIdVal := "1201158065364369018"
	// smsContent := "Dear User,The Incident IN00001245 is created Regards,TCS ICCM"
	// smsType := ""
	// smsUserName := "iFIX_SMS"
	// smsPassword := "eWyuk7V\\yw5JL@UT"
	// Logger.Log.Println("smsPass====>", smsPassword)
	// SMSUrlString := "https://openapi.airtel.in/gateway/airtel-iq-sms-utility/sendSms"
	// ReadProperties.SendSMS(indiaDltContentTemplateId, mobileNo, smsContent, smsType,
	// 	SMSUrlString, IndiaDltPrincipalEntityIdVal, smsUserName, smsPassword)
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

// func main() {
// 	Logger.Log.Println("Main Started")
// 	err := ReadProperties.SendMail("hi", "test", "kaustubh@ifixtechglobal.com", "alam@ifixtechglobal.com")
// 	if err != nil {
// 		// log.Println(err)
// 	}

// 	//Routes.Handle()
// }
