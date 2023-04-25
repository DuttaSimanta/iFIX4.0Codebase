package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"src/config"
	"src/dao"
	"src/entities"
	"src/fileutils"
	Logger "src/logger"
	"strings"
)

func TicketDispatcherBot() {

	wd, err := os.Getwd() // to get working directory
	if err != nil {
		Logger.Log.Println(err)
	}
	//log.Println(wd)
	contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	//log.Println(contextPath)
	props, err := fileutils.ReadPropertiesFile(contextPath + "/resource/application.properties")
	if err != nil {
		Logger.Log.Println(err)
	}
	var supportGroupName = props["BotGroupName"]
	var username = props["BotUser"]
	var defaultGrpName = props["DefaultGroupName"]

	db, dBerr := config.GetDB()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		//return errors.New("ERROR: Unable to connect DB")
	}
	botUserID, userFetchErr := dao.GetBotUserID(db, username)
	if userFetchErr != nil || botUserID == 0 {
		Logger.Log.Println("Unable to fetch bot user")
		return
	}
	defaultGroupID, defaultGroupIDFetchErr := dao.GetDefaultGroupID(db, defaultGrpName)
	if defaultGroupIDFetchErr != nil || defaultGroupID == 0 {
		Logger.Log.Println("Unable to fetch defaultGroupID")
		return
	}
	Logger.Log.Println("Bot UserID", botUserID)
	ticketList, fetchError := dao.GetTicketList(db, supportGroupName)
	if fetchError != nil {
		Logger.Log.Println(fetchError)
		//return errors.New("ERROR: Unable to connect DB")
	}

	Logger.Log.Println("********************Toltal No of tickets To Forward*********************", len(ticketList.Values))
	for i := 0; i < len(ticketList.Values); i++ {
		Logger.Log.Println("========================================================================>")
		Logger.Log.Println("TicketID==================>", ticketList.Values[i].TicketID)
		if botUserID != 0 {
			var GroupForwarding = entities.GroupForwarding{}
			forwardingSupportGrpID, supportGrpFetchErr := dao.GetForwardingSupportGrp(db, &ticketList.Values[i])
			if supportGrpFetchErr != nil {
				Logger.Log.Println(fetchError)
				//return errors.New("ERROR: Unable to connect DB")
			}

			if forwardingSupportGrpID != 0 {

				GroupForwarding.TicketID = ticketList.Values[i].TicketID
				GroupForwarding.MstGroupID = forwardingSupportGrpID
				GroupForwarding.Createdgroupid = ticketList.Values[i].MstGroupID
				GroupForwarding.Mstuserid = 0
				GroupForwarding.Samegroup = false
				GroupForwarding.UserID = botUserID

			} else if defaultGroupID != 0 {
				forwardingSupportGrpID = defaultGroupID
				GroupForwarding.TicketID = ticketList.Values[i].TicketID
				GroupForwarding.MstGroupID = forwardingSupportGrpID
				GroupForwarding.Createdgroupid = ticketList.Values[i].MstGroupID
				GroupForwarding.Mstuserid = 0
				GroupForwarding.Samegroup = false
				GroupForwarding.UserID = botUserID
			} else {
				continue
			}
			sendData, err := json.Marshal(GroupForwarding)
			if err != nil {
				Logger.Log.Println(err)
				//return
			}
			Logger.Log.Println(string(sendData))

			resp, err := http.Post(props["ForwadingURL"], "application/json", bytes.NewBuffer(sendData))
			Logger.Log.Println("Request Sent Forwarding API===>", resp)
			var result map[string]interface{}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Logger.Log.Println(err)
				//return
			}
			err1 := json.Unmarshal(body, &result)
			if err1 != nil {
				Logger.Log.Println(err1)
				//return ticketID, errors.New("Unable to Unmarchal data")
			}

			//  json.NewDecoder(resp.Body).Decode(&res)
			if result["success"].(bool) == false {
				Logger.Log.Println("getting False")
				//return ticketID, errors.New("Ticket creation failed intermittently. Please try again later")
			} else {
				Logger.Log.Println("Forwarding Successful For TicketID===>", ticketList.Values[i].TicketID)
				//ticketID = result["response"].(string)
			}

		}
	}

}
