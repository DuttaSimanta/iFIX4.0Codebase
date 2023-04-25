package fileutils

import (
	"errors"
	"log"
	"os"
	Logger "src/logger"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
	//"html"
	//ReadProperties "src/fileutils"
)

func SendMail(subject string, mailBody string, toEmailAddress string, ccEmailAddress string) error {
	Logger.Log.Println("SendMail Starting...")
	wd, err := os.Getwd()
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Open Directory")
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/")
	//log.Println(contextPath)
	props, err := ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Read Properties File")
	}
	smtpPort, err := strconv.Atoi(props["SmtpPort"])
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Read read smtpPort")
	}
	dialer := gomail.NewDialer(props["SmtpEmail"], smtpPort, props["UserName"], props["password"])
	message := gomail.NewMessage()
	message.SetHeader("From", props["UserName"])
	if !strings.EqualFold(toEmailAddress, "") {
		var toSlice []string = strings.Split(toEmailAddress, ",")
		message.SetHeader("To", toSlice...)
	} else {
		message.SetHeader("To", toEmailAddress)
	}
	//
	if !strings.EqualFold(ccEmailAddress, "") {
		var ccSlice []string = strings.Split(ccEmailAddress, ",")
		message.SetHeader("Cc", ccSlice...)
	}
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
