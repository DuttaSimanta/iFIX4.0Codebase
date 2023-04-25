package model

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"src/config"
	"src/dao"
	"src/entities"
	"src/fileutils"
	Logger "src/logger"
	"strings"
	"time"
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

	var assignedUserName = props["AutomationUser"]
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
				Logger.Log.Println(supportGrpFetchErr)
				//return errors.New("ERROR: Unable to connect DB")
			}
			ticketTypeSeqNo, ticketTypeSeqNoErr := dao.GetTicketTypeSeqNo(db, ticketList.Values[i].TicketTypeID)
			if ticketTypeSeqNoErr != nil {
				Logger.Log.Println(ticketTypeSeqNoErr)
				//return errors.New("ERROR: Unable to connect DB")
			}
			//ticketTypeSeqNo = 5
			if ticketTypeSeqNo == 2 {

				sTaskIDList, sTaskIDListErr := dao.GetStaskIDList(db, &ticketList.Values[i])
				if sTaskIDListErr != nil {
					Logger.Log.Println(sTaskIDListErr)
					//return errors.New("ERROR: Unable to connect DB")
				}
				sTaskIDList = append(sTaskIDList, ticketList.Values[i].TicketID)

				Logger.Log.Println("Stast ========>", sTaskIDList)
				for j := 0; j < len(sTaskIDList); j++ {
					Logger.Log.Println("Stask TicketID==================>", sTaskIDList[j])
					if forwardingSupportGrpID != 0 {

						GroupForwarding.TicketID = sTaskIDList[j]
						GroupForwarding.MstGroupID = forwardingSupportGrpID
						GroupForwarding.Createdgroupid = ticketList.Values[i].MstGroupID
						GroupForwarding.Mstuserid = 0
						GroupForwarding.Samegroup = false
						GroupForwarding.UserID = botUserID

					} else if defaultGroupID != 0 {
						forwardingSupportGrpID = defaultGroupID
						GroupForwarding.TicketID = sTaskIDList[j]
						GroupForwarding.MstGroupID = forwardingSupportGrpID
						GroupForwarding.Createdgroupid = ticketList.Values[i].MstGroupID
						GroupForwarding.Mstuserid = 0
						GroupForwarding.Samegroup = false
						GroupForwarding.UserID = botUserID
					} else {
						Logger.Log.Println("No Group FOUND")
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
						Logger.Log.Println("Forwarding Successful For Stask===>", sTaskIDList[j])

						if sTaskIDList[j] != ticketList.Values[i].TicketID {
							autoResponseError := AutoResponse(db, sTaskIDList[j], ticketList.Values[i].ClientID, ticketList.Values[i].OrgID, assignedUserName, props)
							if autoResponseError != nil {

								continue
							}
						}

						//ticketID = result["response"].(string)
					}

				}
			} else if ticketTypeSeqNo == 3 {
				Logger.Log.Println("Stask found")
				continue
			} else {

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
					Logger.Log.Println("No Group FOUND")
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

					autoResponseError := AutoResponse(db, ticketList.Values[i].TicketID, ticketList.Values[i].ClientID, ticketList.Values[i].OrgID, assignedUserName, props)
					if autoResponseError != nil {

					}
				}

			}

		}
	}

}

func AutoResponse(db *sql.DB, TicketID int64, clientID int64, orgID int64, assignedUserName string, props fileutils.AppConfigProperties) error {

	Logger.Log.Println("============================= IN AUTO Response ===========================================>")
	Logger.Log.Println("TicketID==================>", TicketID)

	assignedUserID, assignedUserIDErr := dao.GetAssignedUserID(db, clientID, orgID, assignedUserName)
	if assignedUserIDErr != nil {
		Logger.Log.Println("User Not Found For self Assignment===>", assignedUserIDErr)
		//return ticketNo, errors.New("User Not Found For self Assignment")
		return errors.New("User Not Found For self Assignment")
	}
	if assignedUserID == 0 {
		Logger.Log.Println("User Not Found For self Assignment===>", assignedUserID)
		//return ticketNo, errors.New("User Not Found For self Assignment")
		return errors.New("User Not Found For self Assignment")
	}
	assignedSupportgroupID, assignedSupportgroupIDErr := dao.GetAssignedUserGrpID(db, clientID, orgID, assignedUserID, TicketID)

	if assignedSupportgroupIDErr != nil {
		Logger.Log.Println("User is not mapped with supportgroup=====>", assignedSupportgroupIDErr)
		//return ticketNo, errors.New("User is not mapped with supportgroup ")
		return errors.New("User is not mapped with supportgroup")
	}
	if assignedSupportgroupID == 0 {
		Logger.Log.Println("User Not Found For self Assignment===>", assignedSupportgroupID)
		//return ticketNo, errors.New("User Not Found For self Assignment")
		return errors.New("User is not mapped with supportgroup")

	}
	Logger.Log.Println("Assigned User ID============>", assignedUserID)

	Logger.Log.Println("Assigned SUPP GRP============>", assignedSupportgroupID)
	time.Sleep(10 * time.Millisecond)
	// First API get Record Details
	Logger.Log.Println("===================================First API get Record Details===============================")

	var recordDetails entities.GetRecordDetailsRequest
	recordDetails.ClientID = clientID
	recordDetails.Mstorgnhirarchyid = orgID
	recordDetails.RecordID = TicketID
	sendData, err := json.Marshal(recordDetails)
	if err != nil {
		Logger.Log.Println("Unable to marshal data ====== First API get Record Details=======>", err)
		//return ticketNo, errors.New("Unable to marshal data")
		return err
	}
	Logger.Log.Println(" Get Record Details Request", string(sendData))

	resp, err := http.Post(props["URLGetRecordDetails"], "application/json", bytes.NewBuffer(sendData))
	Logger.Log.Println("Request Sent To creat record===>", resp)
	var recordDetailsResponeData entities.RecordDetailsResponeData
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger.Log.Println("Unable to read response data========>", err)
		//	return ticketNo, errors.New("Unable to read response data")
		return err
	}
	err1 := json.Unmarshal(body, &recordDetailsResponeData)
	if err1 != nil {
		Logger.Log.Println(err1)
		//return ticketNo, errors.New("Unable to Unmarchal data")
		return err1
	}
	resp.Body.Close()
	time.Sleep(10 * time.Millisecond)
	//  json.NewDecoder(resp.Body).Decode(&res)
	if recordDetailsResponeData.Status == false {
		Logger.Log.Println("getting False")
		//return ticketNo, errors.New("Ticket Details Fetching failed intermittently. Please try again later")
		return errors.New("===================recordDetailsResponeData getting false========================")
	} else {
		Logger.Log.Println(recordDetailsResponeData)
		if len(recordDetailsResponeData.Details) == 0 {
			Logger.Log.Println("No TicketFound")
			return errors.New("No TicketFound")
		}
		workflowID := recordDetailsResponeData.Details[0].WorkFlowDetails.WorkFlowID
		//recordStageID := recordDetailsResponeData.Details[0].RecordStageID
		clientID = recordDetailsResponeData.Details[0].Clientid
		orgID = recordDetailsResponeData.Details[0].Mstorgnhirarchyid
		cattypeID := recordDetailsResponeData.Details[0].WorkFlowDetails.CatTypeID
		catID := recordDetailsResponeData.Details[0].WorkFlowDetails.CatID
		Logger.Log.Println("===================================Second API get State By Seq===============================")
		currentStateID, currentStateIDErr := dao.GetRecordCurrentStateID(db, clientID, orgID, TicketID)
		if currentStateIDErr != nil {
			Logger.Log.Println("Unable to fetch current state id====>", currentStateIDErr)
			//return ticketNo, errors.New("Unable to fetch current state id")
			return currentStateIDErr
		}
		if currentStateID == 0 {
			Logger.Log.Println("Unable to fetch current state id====>", currentStateID)
			return errors.New("getting state ID 0")
		}
		Logger.Log.Println("===================================Third API get State By Seq===============================")
		var stateBySeqRequest entities.GetStateBySeqRequest

		stateBySeqRequest.ClientID = clientID
		stateBySeqRequest.MstorgnhirarchyID = orgID
		stateBySeqRequest.Typeseqno = 2
		stateBySeqRequest.SeqNo = 2
		stateBySeqRequest.TransitionID = 0
		stateBySeqRequest.ProcessID = workflowID
		stateBySeqRequest.UserID = assignedUserID

		sendData, err := json.Marshal(stateBySeqRequest)
		if err != nil {
			Logger.Log.Println("Unable to third api marshal data====>", err)
			//return ticketNo, errors.New("Unable to marshal data")
			return errors.New("3rd api data marshar error")
		}
		Logger.Log.Println(" Get State By Seq Request", string(sendData))

		resp, err := http.Post(props["URLGetStateSeq"], "application/json", bytes.NewBuffer(sendData))
		Logger.Log.Println("Request Sent To creat record===>", resp)
		var stateSeqResponse = entities.StateSeqResponse{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Logger.Log.Println(err)
			//return ticketNo, errors.New("Unable to read response data")
			return errors.New("3rd api response body error")
		}
		err1 := json.Unmarshal(body, &stateSeqResponse)
		if err1 != nil {
			Logger.Log.Println(err1)
			//return ticketNo, errors.New("Unable to Unmarchal data")
			return errors.New("3rd api response unmarsharl error")
		}
		resp.Body.Close()
		time.Sleep(10 * time.Millisecond)
		//  json.NewDecoder(resp.Body).Decode(&res)
		if stateSeqResponse.Success == false {
			Logger.Log.Println("getting False")
			//return ticketNo, errors.New("State Sequence Fetching failed intermittently. Please try again later")
			return errors.New("3rd api getting respomnse false")
		} else {
			Logger.Log.Println(stateSeqResponse.Details[0])
			mststateID := stateSeqResponse.Details[0].Mststateid

			//Forth API Move workFlow
			Logger.Log.Println("===================================Forth API Move workFlow===============================")

			var moveWorkFlowRequest entities.MoveWorkFlowRequest
			moveWorkFlowRequest.ClientID = clientID
			moveWorkFlowRequest.MstorgnhirarchyID = orgID
			moveWorkFlowRequest.RecorddifftypeID = cattypeID
			moveWorkFlowRequest.RecordDiffID = catID
			moveWorkFlowRequest.TransitionID = 0
			moveWorkFlowRequest.PreviousstateID = currentStateID
			moveWorkFlowRequest.CurrentstateID = mststateID
			moveWorkFlowRequest.Manualstateselection = 0
			moveWorkFlowRequest.TransactionID = TicketID
			moveWorkFlowRequest.CreatedgroupID = assignedSupportgroupID
			moveWorkFlowRequest.Issrrequestor = 0
			moveWorkFlowRequest.UserID = assignedUserID
			moveWorkFlowRequest.MstgroupID = assignedSupportgroupID
			moveWorkFlowRequest.MstuserID = assignedUserID

			sendData, err := json.Marshal(moveWorkFlowRequest)
			if err != nil {
				Logger.Log.Println(err)
				//return ticketNo, errors.New("Unable to marshal data")
				return errors.New("4th api data marsharl error")
			}
			Logger.Log.Println(" Get Moveworkflow Request", string(sendData))

			resp, err := http.Post(props["URLMoveWorkFlow"], "application/json", bytes.NewBuffer(sendData))
			Logger.Log.Println("Request Sent To creat record===>", resp)
			var moveWorkFlowResult map[string]interface{}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Logger.Log.Println(err)
				//return ticketNo, errors.New("Unable to read response data")
				return errors.New("4th api  error")
			}
			err1 := json.Unmarshal(body, &moveWorkFlowResult)
			if err1 != nil {
				Logger.Log.Println(err1)
				//return ticketNo, errors.New("Unable to Unmarchal data")
				return errors.New("4th api  error")
			}
			resp.Body.Close()
			time.Sleep(10 * time.Millisecond)
			//  json.NewDecoder(resp.Body).Decode(&res)
			if moveWorkFlowResult["success"].(bool) == false {
				Logger.Log.Println("getting False")
				//return ticketNo, errors.New("Move WorkFlow failed intermittently. Please try again later")
				return errors.New("4th api  error")
			} else {
				Logger.Log.Println(moveWorkFlowResult)

				//Fifth API Move ChangeSupportGroup
				Logger.Log.Println("===================================Fifth API Move ChangeSupportGroup===============================")
				var changeRecordGroupRequest entities.ChangeRecordGroupRequest

				changeRecordGroupRequest.CreatedgroupID = assignedSupportgroupID
				changeRecordGroupRequest.MstgroupID = assignedSupportgroupID
				changeRecordGroupRequest.MstuserID = assignedUserID
				changeRecordGroupRequest.Samegroup = true
				changeRecordGroupRequest.TransactionID = TicketID
				changeRecordGroupRequest.UserID = assignedUserID
				sendData, err := json.Marshal(changeRecordGroupRequest)
				if err != nil {
					Logger.Log.Println(err)
					//return ticketNo, errors.New("Unable to marshal data")
					return errors.New("5rd api  error")
				}
				Logger.Log.Println(" changeRecordGroup Request", string(sendData))

				resp, err := http.Post(props["URLChangeRecordGroup"], "application/json", bytes.NewBuffer(sendData))
				Logger.Log.Println("Request Sent To creat record===>", resp)
				var changeRecordGroupResult map[string]interface{}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					Logger.Log.Println(err)
					//return ticketNo, errors.New("Unable to read response data")
					return errors.New("5th api  error")
				}
				err1 := json.Unmarshal(body, &changeRecordGroupResult)
				if err1 != nil {
					Logger.Log.Println(err1)
					//return ticketNo, errors.New("Unable to Unmarchal data")
					return errors.New("5th api  error")
				}
				resp.Body.Close()
				time.Sleep(10 * time.Millisecond)
				//  json.NewDecoder(resp.Body).Decode(&res)
				if changeRecordGroupResult["success"].(bool) == false {
					Logger.Log.Println("getting False")
					//return ticketNo, errors.New("changeRecordGroup failed intermittently. Please try again later")
					return errors.New("5th api  error")
				} else {
					Logger.Log.Println(changeRecordGroupResult)
					Logger.Log.Println("======================Self Assignment Done Successfully=============>")

				}
			}

		}

	}

	return nil
}
