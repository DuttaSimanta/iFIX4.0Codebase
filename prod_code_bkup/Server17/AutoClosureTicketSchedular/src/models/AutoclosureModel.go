package models

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
	"src/logger"

	Logger "src/logger"
	"strings"
)

func Autoclosure() {

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
	logger.Log.Println("In side Autoclosure model function")
	db, dBerr := config.GetDB()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
		//return errors.New("ERROR: Unable to connect DB")
	}
	autoclosureconfigList, autoclosureconfigListErr := dao.GetAutoClosureConfigList(db)
	if autoclosureconfigListErr != nil {
		logger.Log.Println("autoclosureconfigListErr is --------->", autoclosureconfigListErr)
	}
	for j := 0; j < len(autoclosureconfigList); j++ {
		if autoclosureconfigList[j].AutocloseFromdate == "" || autoclosureconfigList[j].AutocloseTodate == "" {
			logger.Log.Println("Date Is Wrong:", autoclosureconfigList[j].ClientID, autoclosureconfigList[j].MstorgnhirarchyID, autoclosureconfigList[j].AutocloseFromdate, autoclosureconfigList[j].AutocloseTodate)

			continue
		}
		//dao := dao.DbConn{DB: db}
		values, err := dao.GetResolvedRecordsInfo(db, autoclosureconfigList[j])
		if err != nil {
			logger.Log.Println("Error is --------->", err)
		}
		logger.Log.Println("Total Tickets ****************************", len(values))
		logger.Log.Println("\n\n")

		for i := 0; i < len(values); i++ {

			var ClientID int64 = values[i].ClientID
			var MstorgnhirarchyID int64 = values[i].MstorgnhirarchyID
			logger.Log.Println("Total Tickets  --------->", len(values))
			logger.Log.Println("===================================================Ticket ID is ===========================>", values[i].RecordID)

			nxtstateID, err := dao.GetNxtStateID(db, ClientID, MstorgnhirarchyID)
			if err != nil {
				logger.Log.Println("Error is --------->", err)
			}
			logger.Log.Println("nxtstateID is---->", nxtstateID)
			hashmap, err := dao.Termsequance(db, ClientID, MstorgnhirarchyID)
			if err != nil {
				logger.Log.Println("Error is --------->", err)
			}

			sequance, err := dao.GetCurrentStatusSeq(db, ClientID, MstorgnhirarchyID, values[i].RecordID)
			if err != nil {
				logger.Log.Println("Error is --------->", err)
			}
			if sequance == 3 {

				termreqbody := &entities.RecordmultiplecommonEntity{}

				termreqbody.ClientID = values[i].ClientID
				termreqbody.Mstorgnhirarchyid = values[i].MstorgnhirarchyID
				termreqbody.RecordID = values[i].RecordID
				termreqbody.RecordstageID = values[i].RecordStageID
				termreqbody.ForuserID = values[i].MstuserID
				termreqbody.Userid = values[i].MstuserID
				termreqbody.Usergroupid = values[i].CreatedgrpID
				termreqbody.Details = append(termreqbody.Details, entities.RecordTermnamesEntity{ID: hashmap[17], Insertedvalue: "Auto Close"})
				termreqbody.Details = append(termreqbody.Details, entities.RecordTermnamesEntity{ID: hashmap[18], Insertedvalue: "NULL"})
				termreqbody.Recorddifftypeid = values[i].TickettypeDiffTypeId
				termreqbody.Recorddiffid = values[i].TickettypeDiffId
				postdata, _ := json.Marshal(termreqbody)

				// logger.Log.Println("Record status request body -->", reqbd)

				logger.Log.Println("Insert RecordTerm request body -->", string(postdata))
				responseData := bytes.NewBuffer(postdata)
				response, err := http.Post(props["InsertTerm"], "application/json", responseData)
				if err != nil {
					logger.Log.Println("Error is --------->", err)
				}
				defer response.Body.Close()
				bodyy, err := ioutil.ReadAll(response.Body)
				if err != nil {
					logger.Log.Println("Error is --------->", err)
				}
				sbb := string(bodyy)
				result := entities.RecordcommonResponseInt{}
				json.Unmarshal([]byte(sbb), &result)
				var termflagflag = result.Success
				var termerrormsg = result.Message
				logger.Log.Println("Auto close Term message -->", termflagflag, termerrormsg)
				if termflagflag != true {

					err2 := dao.UpdateDefectiveFlag(db, values[i].ClientID, values[i].MstorgnhirarchyID, values[i].RecordID)
					if err2 != nil {
						logger.Log.Println("Errorr in UpdateDeefectiveflag --------->", err2)
					}
					logger.Log.Println("<***************************InsertTermerror API Error****************************>", values[i].RecordID)
				}
				//rec.ClientID, rec.Mstorgnhirarchyid, rec.RecordID, rec.RecordstageID, rec.TermID, rec.Termvalue, rec.ForuserID, rec.Userid
				// err := dao.InsertClosureComment(db, values[i].ClientID, values[i].MstorgnhirarchyID, values[i].RecordID, values[i].RecordStageID, hashmap[17], "Auto Close", values[i].MstuserID, values[i].CreatedgrpID)
				// if err != nil {
				// 	logger.Log.Println("Error is --------->", err)
				// }
				// commenttermnm, err := dao.Gettermnamebyid(db, hashmap[17], values[i].ClientID, values[i].MstorgnhirarchyID)
				// if err != nil {
				// 	logger.Log.Println("Error is --------->", err)
				// }
				// err = dao.InsertActivityLogs(db, values[i].ClientID, values[i].MstorgnhirarchyID, values[i].RecordID, 100, commenttermnm+" :: Auto Close", values[i].MstuserID, values[i].CreatedgrpID, hashmap[17])
				// if err != nil {
				// 	logger.Log.Println("Error is --------->", err)
				// }

				// err1 := dao.InsertNPSFeedback(db, values[i].ClientID, values[i].MstorgnhirarchyID, values[i].RecordID, values[i].RecordStageID, hashmap[18], "NULL", values[i].MstuserID, values[i].CreatedgrpID)
				// if err1 != nil {
				// 	logger.Log.Println("Error is --------->", err1)
				// }
				// feedbacktermnm, err := dao.Gettermnamebyid(db, hashmap[18], values[i].ClientID, values[i].MstorgnhirarchyID)
				// if err != nil {
				// 	logger.Log.Println("Error is --------->", err)
				// }
				// err = dao.InsertActivityLogs(db, values[i].ClientID, values[i].MstorgnhirarchyID, values[i].RecordID, 100, feedbacktermnm+" :: NULL", values[i].MstuserID, values[i].CreatedgrpID, hashmap[18])
				// if err != nil {
				// 	logger.Log.Println("Error is --------->", err)
				// }

				reqbd := &entities.RequestBody{}
				reqbd.ClientID = values[i].ClientID
				reqbd.MstorgnhirarchyID = values[i].MstorgnhirarchyID
				reqbd.RecorddifftypeID = values[i].WorkingDifftypeID
				reqbd.RecorddiffID = values[i].WorkingDiffID

				reqbd.PreviousstateID = values[i].PreviousStateID
				reqbd.CurrentstateID = nxtstateID
				reqbd.TransactionID = values[i].RecordID
				reqbd.CreatedgroupID = values[i].CreatedgrpID
				reqbd.MstgroupID = values[i].CreatedgrpID
				reqbd.MstuserID = values[i].MstuserID
				reqbd.UserID = values[i].MstuserID
				postBody, _ := json.Marshal(reqbd)
				logger.Log.Println("Record status request body -->", reqbd)
				responseBody := bytes.NewBuffer(postBody)
				resp, err := http.Post(props["MoveWorkflowURL"], "application/json", responseBody)
				//resp, err := http.Post("http://20.204.74.38:8082/api/moveWorkflow", "application/json", responseBody)
				if err != nil {
					logger.Log.Println("Error is --------->", err)
				}
				if resp != nil {
					Logger.Log.Println("NO Response From MoveWorkflowURL")
					defer resp.Body.Close()
					//					continue
					//}
					body, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						logger.Log.Println("Error is --------->", err)
					}
					sb := string(body)
					wfres := entities.WorkflowResponse{}
					json.Unmarshal([]byte(sb), &wfres)
					var workflowflag = wfres.Success
					var errormsg = wfres.Message
					logger.Log.Println("Auto close workflow message -->", workflowflag)
					logger.Log.Println("Auto close workflow response error message -->", errormsg, "======", values[i].RecordID)
					if workflowflag == true {
						err2 := dao.UpdateclosureFlag(db, values[i].ClientID, values[i].MstorgnhirarchyID, values[i].RecordID)
						if err2 != nil {
							logger.Log.Println("Errorr in UpdateclosureFlag --------->", err2)
						} else {
							logger.Log.Println("<================================Ticket IS closed===============================>", values[i].RecordID)
						}
						logger.Log.Println("")
					} else {
						err2 := dao.UpdateDefectiveFlag(db, values[i].ClientID, values[i].MstorgnhirarchyID, values[i].RecordID)
						if err2 != nil {
							logger.Log.Println("Errorr in UpdateDeefectiveflag --------->", err2)
						}
						logger.Log.Println("<***************************MoveWorkflow API Error****************************>", values[i].RecordID)
					}
				}
			}
		}
	}

}
