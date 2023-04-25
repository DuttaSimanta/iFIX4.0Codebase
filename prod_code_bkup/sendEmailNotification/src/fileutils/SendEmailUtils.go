package fileutils

import (
	"errors"
	Logger "src/logger"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
	//"html"
	//ReadProperties "src/fileutils"
)

func SendMaiL(subject string, mailBody string, toEmailAddress string, ccEmailAddress string,
	smtpHostForNotification string, smtpPort string, emailUserName string, emailPassword string) error {
	Logger.Log.Println("SendMail Starting...")
	// wd, err := os.Getwd()
	// if err != nil {
	// 	Logger.Log.Println(err)
	// 	return errors.New("ERROR: Unable to Open Directory")
	// }
	// //log.Println(wd)
	// contextPath := strings.ReplaceAll(wd, "\\", "/")
	// //log.Println(contextPath)
	// //props, err := ReadPropertiesFile(contextPath + "/resource/application.properties")
	// if err != nil {
	// 	Logger.Log.Println(err)
	// 	return errors.New("ERROR: Unable to Read Properties File")
	// }

	smtpport, err := strconv.Atoi(smtpPort)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Read smtpPort")
	}
	dialer := gomail.NewDialer(smtpHostForNotification, smtpport, emailUserName, emailPassword)
	message := gomail.NewMessage()
	message.SetHeader("From", emailUserName)
	/*if strings.EqualFold(toEmailAddress, "") {
		var toSlice []string = strings.Split(toEmailAddress, ",")
		message.SetHeader("To", toSlice...)
	} else {
		message.SetHeader("To", toEmailAddress)
	}*/
	//

	var toSlice []string = strings.Split(toEmailAddress, ",")
	message.SetHeader("To", toSlice...)
	if !strings.EqualFold(ccEmailAddress, "") {
		var ccSlice []string = strings.Split(ccEmailAddress, ",")
		message.SetHeader("Cc", ccSlice...)
	}
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", mailBody)
	Logger.Log.Println(toEmailAddress)
	if err := dialer.DialAndSend(message); err != nil {

		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Send Mail")
	}
	Logger.Log.Println("Email Sent Successfully")
	return nil
}
