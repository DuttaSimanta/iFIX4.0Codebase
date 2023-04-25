package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"src/dao"
	"src/entities"
	"src/fileutils"
	"src/logger"
	Logger "src/logger"
	"strings"
	"time"
)

func AutoResponse() {

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

	if db == nil {
		dbcon, err := ConnectMySqlDb()
		if err != nil {
			logger.Log.Println("Error in DBConnection")
			return
		}
		db = dbcon
	}
	srouceType := props["Source"]
	ticketList, fetchError := dao.GetTicketList(db, srouceType)
	if fetchError != nil {
		Logger.Log.Println(fetchError)
		//return errors.New("ERROR: Unable to connect DB")
	}

	Logger.Log.Println("********************Toltal No of tickets To Forward*********************", len(ticketList.Values))
	for i := 0; i < len(ticketList.Values); i++ {
		Logger.Log.Println("========================================================================>")
		Logger.Log.Println("TicketID==================>", ticketList.Values[i].TicketID)
		Logger.Log.Println("TicketID==================>", ticketList.Values[i].TicketCode)
		clientID := ticketList.Values[i].ClientID
		orgID := ticketList.Values[i].OrgID
		assignedUserName := props["AutomationUser"]
		TicketID := ticketList.Values[i].TicketID

		assignedUserID, assignedUserIDErr := dao.GetAssignedUserID(db, clientID, orgID, assignedUserName)
		if assignedUserIDErr != nil {
			Logger.Log.Println("User Not Found For self Assignment===>", assignedUserIDErr)
			//return ticketNo, errors.New("User Not Found For self Assignment")
			continue
		}
		if assignedUserID == 0 {
			Logger.Log.Println("User Not Found For self Assignment===>", assignedUserID)
			//return ticketNo, errors.New("User Not Found For self Assignment")
			continue
		}
		assignedSupportgroupID, assignedSupportgroupIDErr := dao.GetAssignedUserGrpID(db, clientID, orgID, assignedUserID, TicketID)

		if assignedSupportgroupIDErr != nil {
			Logger.Log.Println("User is not mapped with supportgroup=====>", assignedSupportgroupIDErr)
			//return ticketNo, errors.New("User is not mapped with supportgroup ")
			continue
		}
		if assignedSupportgroupID == 0 {
			Logger.Log.Println("User Not Found For self Assignment===>", assignedSupportgroupID)
			//return ticketNo, errors.New("User Not Found For self Assignment")
			continue
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
			Logger.Log.Println("Unable to marshal data=======>", err)
			//return ticketNo, errors.New("Unable to marshal data")
			continue
		}
		Logger.Log.Println(" Get Record Details Request", string(sendData))

		resp, err := http.Post(props["URLGetRecordDetails"], "application/json", bytes.NewBuffer(sendData))
		Logger.Log.Println("Request Sent To creat record===>", resp)
		var recordDetailsResponeData entities.RecordDetailsResponeData
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Logger.Log.Println("Unable to read response data========>", err)
			//	return ticketNo, errors.New("Unable to read response data")
			continue
		}
		err1 := json.Unmarshal(body, &recordDetailsResponeData)
		if err1 != nil {
			Logger.Log.Println(err1)
			//return ticketNo, errors.New("Unable to Unmarchal data")
			continue
		}
		resp.Body.Close()
		time.Sleep(10 * time.Millisecond)
		//  json.NewDecoder(resp.Body).Decode(&res)
		if recordDetailsResponeData.Status == false {
			Logger.Log.Println("getting False")
			//return ticketNo, errors.New("Ticket Details Fetching failed intermittently. Please try again later")
			continue
		} else {
			Logger.Log.Println(recordDetailsResponeData)
			if len(recordDetailsResponeData.Details) == 0 {
				continue
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
				continue
			}
			if currentStateID == 0 {
				Logger.Log.Println("Unable to fetch current state id====>", currentStateID)
				continue
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
				Logger.Log.Println("Unable to marshal data====>", err)
				//return ticketNo, errors.New("Unable to marshal data")
				continue
			}
			Logger.Log.Println(" Get State By Seq Request", string(sendData))

			resp, err := http.Post(props["URLGetStateSeq"], "application/json", bytes.NewBuffer(sendData))
			Logger.Log.Println("Request Sent To creat record===>", resp)
			var stateSeqResponse = entities.StateSeqResponse{}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				Logger.Log.Println(err)
				//return ticketNo, errors.New("Unable to read response data")
				continue
			}
			err1 := json.Unmarshal(body, &stateSeqResponse)
			if err1 != nil {
				Logger.Log.Println(err1)
				//return ticketNo, errors.New("Unable to Unmarchal data")
				continue
			}
			resp.Body.Close()
			time.Sleep(10 * time.Millisecond)
			//  json.NewDecoder(resp.Body).Decode(&res)
			if stateSeqResponse.Success == false {
				Logger.Log.Println("getting False")
				//return ticketNo, errors.New("State Sequence Fetching failed intermittently. Please try again later")
				continue
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
					continue
				}
				Logger.Log.Println(" Get Moveworkflow Request", string(sendData))

				resp, err := http.Post(props["URLMoveWorkFlow"], "application/json", bytes.NewBuffer(sendData))
				Logger.Log.Println("Request Sent To creat record===>", resp)
				var moveWorkFlowResult map[string]interface{}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					Logger.Log.Println(err)
					//return ticketNo, errors.New("Unable to read response data")
					continue
				}
				err1 := json.Unmarshal(body, &moveWorkFlowResult)
				if err1 != nil {
					Logger.Log.Println(err1)
					//return ticketNo, errors.New("Unable to Unmarchal data")
					continue
				}
				resp.Body.Close()
				time.Sleep(10 * time.Millisecond)
				//  json.NewDecoder(resp.Body).Decode(&res)
				if moveWorkFlowResult["success"].(bool) == false {
					Logger.Log.Println("getting False")
					//return ticketNo, errors.New("Move WorkFlow failed intermittently. Please try again later")
					continue
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
						continue
					}
					Logger.Log.Println(" changeRecordGroup Request", string(sendData))

					resp, err := http.Post(props["URLChangeRecordGroup"], "application/json", bytes.NewBuffer(sendData))
					Logger.Log.Println("Request Sent To creat record===>", resp)
					var changeRecordGroupResult map[string]interface{}
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						Logger.Log.Println(err)
						//return ticketNo, errors.New("Unable to read response data")
						continue
					}
					err1 := json.Unmarshal(body, &changeRecordGroupResult)
					if err1 != nil {
						Logger.Log.Println(err1)
						//return ticketNo, errors.New("Unable to Unmarchal data")
						continue
					}
					resp.Body.Close()
					time.Sleep(10 * time.Millisecond)
					//  json.NewDecoder(resp.Body).Decode(&res)
					if changeRecordGroupResult["success"].(bool) == false {
						Logger.Log.Println("getting False")
						//return ticketNo, errors.New("changeRecordGroup failed intermittently. Please try again later")
						continue
					} else {
						Logger.Log.Println(changeRecordGroupResult)
						Logger.Log.Println("======================Self Assignment Done Successfully=============>")

					}
				}

			}

		}

	}

}
