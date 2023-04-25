package fileutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"src/entities"
	Logger "src/logger"
	"strings"
)

func SendSMS(indiaDltContentTemplateId string, mobileNo string, smsContent string, smsType string,
	SMSUrlString string, IndiaDltPrincipalEntityIdVal string, smsUserName string, smsPassword string) error {
	client := &http.Client{}

	wd, err := os.Getwd()
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Open Directory")
	}
	//smsContent = strings.ReplaceAll(smsContent, " ", "%20")
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/")
	//log.Println(contextPath)
	props, err := ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Read Properties File")
	}
	// smsURL = props["SMSUrl"] + props["UsernameString"] + "=" + props["UsernameVal"] + "&" + props["PasswordString"] + "=" + props["PasswordVal"] +
	// 	"&" + props["IndiaDltPrincipalEntityIdString"] + "=" + props["IndiaDltPrincipalEntityIdVal"] + "&" + props["IndiaDltContentTemplateIdString"] + "=" +
	// 	indiaDltContentTemplateId + "&" + props["ToString"] + "=" + mobileNo + "&" + props["TextString"] + "=" + smsContent +
	// 	"&" + props["TypeString"] + "=" + "longSMS"

	// basicAuthorizationString := smsUserName + ":" + smsPassword
	// basicAuthorizationKey = base64.StdEncoding.EncodeToString([]byte(basicAuthorizationString))

	var smsServiceEntity = entities.SMSServiceEntity{}
	//	var toSlice []string = strings.Split(toEmailAddress, ",")
	smsServiceEntity.CustomerID = props["CustomerID"]
	smsServiceEntity.DestinationAddress = strings.Split(mobileNo, ",")
	smsServiceEntity.DltTemplateID = indiaDltContentTemplateId
	smsServiceEntity.EntityID = IndiaDltPrincipalEntityIdVal
	smsServiceEntity.MessageType = props["MessageType"]
	smsServiceEntity.Message = smsContent
	smsServiceEntity.SourceAddress = props["SourceAddress"]

	sendData, err := json.Marshal(smsServiceEntity)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("Unable to marshal data")
	}
	req, err := http.NewRequest("POST", SMSUrlString, bytes.NewBuffer(sendData))
	req.SetBasicAuth(smsUserName, smsPassword)
	req.Header.Set("Content-Type", "application/json")
	//resp, err := http.Post(props["URL"], "application/json", bytes.NewBuffer(sendData))

	// smsURL = SMSUrlString + props["UsernameString"] + "=" + smsUserName + "&" + props["PasswordString"] + "=" + smsPassword +
	// 	"&" + props["IndiaDltPrincipalEntityIdString"] + "=" + IndiaDltPrincipalEntityIdVal + "&" + props["IndiaDltContentTemplateIdString"] + "=" +
	// 	indiaDltContentTemplateId + "&" + props["ToString"] + "=" + mobileNo + "&" + props["TextString"] + "=" + smsContent +
	// 	"&" + props["TypeString"] + "=" + "longSMS"

	// Logger.Log.Println("smsURL===>>", smsURL)

	resp, err := client.Do(req)

	//resp, err := http.Get(smsURL)
	if err != nil {
		Logger.Log.Println(err)
		return errors.New("ERROR: Unable to Send SMS")
	}
	defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)
	var payLoad map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger.Log.Println(err)

		return errors.New("ERROR: Not able to fetch Request Data")
	}
	jsonErr := json.Unmarshal(body, &payLoad)
	Logger.Log.Println("Payload====>", payLoad)
	payLoadString := fmt.Sprintf("%v", payLoad)
	if jsonErr != nil {
		Logger.Log.Println(jsonErr)

		return errors.New("ERROR: Json Unmarshal error")
	}
	if strings.Contains(payLoadString, "messageRequestId") {
		Logger.Log.Println("SMS Sent Successfully")
	} else {
		Logger.Log.Println("ERROR: Something went Wrong")
		return errors.New("ERROR: Something went Wrong")
	}

	return nil
}
