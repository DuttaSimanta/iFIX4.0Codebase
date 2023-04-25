package fileutils

import (
	"log"
	"errors"
	"strings"
	"os"
	"gopkg.in/gomail.v2"
	"strconv"
Logger 	"src/logger"
//"html"
//ReadProperties "src/fileutils"

)

func SendMail(subject string,mailBody string,toEmailAddress string) error{
	Logger.Log.Println("SendMail Starting...")
	wd, err := os.Getwd()
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Open Directory")
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd,"\\","/")
	//log.Println(contextPath)
	props, err := ReadPropertiesFile(contextPath+"/resource/application.properties")
	if err != nil {
        Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Read Properties File")
	}
	smtpPort ,err := strconv.Atoi (props["SmtpPort"])
	if err != nil {
        Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Read read smtpPort")
	}
	dialer := gomail.NewDialer(props["SmtpEmail"],smtpPort,props["UserName"], props["password"])
	message := gomail.NewMessage()
	message.SetHeader("From", props["UserName"])
	message.SetHeader("To", toEmailAddress)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", mailBody)
	log.Println(toEmailAddress)
	if err := dialer.DialAndSend(message); err != nil {
		
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Send Mail")
	} 
	Logger.Log.Println("Email Sent Successfully")
	return nil
}