package models

import (
	"encoding/json"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"log"
	"strconv"
)

func Mapprocesstemplate(tz *entities.MapprocesstemplateEntity) (int64, bool, error, string) {
	log.Println("In side model",tz)
	dbcon, err := config.ConnectMySqlDb()
	if err != nil {
		dbcon.Close()
		logger.Log.Println("Database connection failure", err)
		log.Println("Database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	defer dbcon.Close()

	dataAccess := dao.DbConn{DB: dbcon}
	/**
	Get template name by template id ,client id and org id
	*/
	processnames, prerr := dataAccess.Gettemplatename(tz)
	if prerr != nil {
		return 0, false, prerr, "Something Went Wrong"
	}
	if len(processnames) == 0 {
		return 0, false, nil, "Template not found"
	}
	/**
	Get template table and column details by template id ,client id and org id
	*/
	entitityfields, ererr := dataAccess.Gettemplateentity(tz)
	if ererr != nil {
		return 0, false, ererr, "Something Went Wrong"
	}
	if len(entitityfields) == 0 {
		return 0, false, nil, "Entity not found"
	}
	/**
	Get template transition details by template id ,client id and org id
	*/
	transitions, terr := dataAccess.Gettemplatetransition(tz)
	if ererr != nil {
		return 0, false, terr, "Something Went Wrong"
	}
	if len(transitions) == 0 {
		return 0, false, nil, "Transition path not found"
	}

	wentity := entities.Workflowentity{}
	wentity.Processid = tz.Processid
	wentity.Clientid = tz.Clientid
	wentity.Mstorgnhirarchyid = tz.Loggedinmstorgnhirarchyid
	/**
	Get full template xml and json details
	*/
	details, err1 := dataAccess.Getprocesstemplatedetails(&wentity)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	if len(details) == 0 {
		return 0, false, nil, "Template details not defined"
	}
	/**
	Get all state seq which are mapped with the template
	*/
	tstates, err1 := dataAccess.Gettemplatestate(tz)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	if len(tstates) == 0 {
		return 0, false, nil, "Template state not defined"
	}
	var ids string = ""
	for i, state := range tstates {
		if i > 0 {
			ids += ","
		}
		ids += strconv.Itoa(int(state.Id))
	}
	/**
	Get state ids of the organization which are mapped with the template
	*/
	pstates, err1 := dataAccess.Getprocessstatebyseq(tz, ids)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	if len(pstates) == 0 {
		return 0, false, nil, " State not defined for organization"
	}
	processEntity := entities.MstprocessEntity{}
	processEntity.Clientid = tz.Clientid
	processEntity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid

	processEntity.Mstdatadictionaryfieldid = entitityfields[0].Mstdatadictionaryfieldid
	//var lastinsertedID int64
	for _, diff := range tz.Recorddiffids {
		tx, err := dbcon.Begin()
		if err != nil {
			logger.Log.Println("Transaction creation error.", err)
			return 0, false, err, "Something Went Wrong"
		}
		diffname, derr := dataAccess.Getdiffnamebyid(diff.Id)
		if derr != nil {
			return 0, false, derr, "Data insertion failure."
		}
		if len(diffname) > 0 {
			processEntity.Processname = diffname[0] + " " + processnames[0].Processname
		} else {
			processEntity.Processname = processnames[0].Processname
		}
		processEntity.Recorddiffid = diff.Id
		processEntity.Recorddifftypeid = diff.Type
		wentity.Recorddiffid = diff.Id
		wentity.Recorddifftypeid = diff.Type
		/**
		Duplicate checking for process name
		*/
		count, err := dataAccess.CheckDuplicateMstprocess(&processEntity)
		if err != nil {
			return 0, false, err, "Data insertion failure."
		}
		if count.Total > 0 {
			logger.Log.Println("Process Already Exist with the name.")
			log.Println("Process Already Exist with the name.")
			break
			//return 0, false, err, "Process Already Exist with the name"
		}

		/**
		Insert Process details in mstprocess table
		*/
		lastinsertedID, err := dao.InsertMstprocesswithtransaction(tx, &processEntity)
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Data insertion failure."
		}
		/**
		Duplicate checking for category mapped with process.
		One category is mapped with only one process.
		It is a one to one mapping
		*/
		recount, err := dataAccess.CheckDuplicateMstprocessrecordmap(&processEntity)
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Data insertion failure."
		}
		if recount.Total > 0 {
			logger.Log.Println("Process already mapped with category.")
			log.Println("Process already mapped with category.")
			tx.Rollback()
			break
			//return 0, false, err, "Process already mapped with category"
		}
		/**
		Insert the category which are mapped with the process
		*/
		_, reerr := dao.InsertMstprocessrecordmapwithtransaction(tx, &processEntity, lastinsertedID)
		if reerr != nil {
			tx.Rollback()
			return 0, false, reerr, "Data insertion failure."
		}
		/**
		Check whether one table and column is mapped with one process
		*/
		encount, err := dataAccess.CheckDuplicateMapprocesstoentity(&processEntity, lastinsertedID)
		if err != nil {
			tx.Rollback()
			return 0, false, err, "Data insertion failure."
		}
		if encount.Total > 0 {
			logger.Log.Println("Process already mapped with table.")
			log.Println("Process already mapped with table.")
			tx.Rollback()
			break
			//return 0, false, err, "Process already mapped with table"
		}
		/**
		Insert table and column which are mapped for process
		*/
		_, enerr := dao.InsertMapprocesstoentitywithtransaction(tx, &processEntity, lastinsertedID)
		if enerr != nil {
			tx.Rollback()
			return 0, false, enerr, "Data insertion failure."
		}
		wentity.Processid = lastinsertedID
		wentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
		values, err1 := dataAccess.Getprocessdetails(&wentity)
		if err1 != nil {
			tx.Rollback()
			return 0, false, err1, "Something Went Wrong"
		}
		if len(values) > 0 {
			logger.Log.Println("Process details already mapped.")
			log.Println("Process details already mapped.")
			tx.Rollback()
			break
			//return 0, false, err, "Process already mapped with table"
		}
		pstateentity := entities.MapprocessstateEntity{}
		pstateentity.Clientid = tz.Clientid
		pstateentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
		pstateentity.Processid = lastinsertedID

		/**
		Map states with the process,which are already mapped with the template
		*/
		for _, state := range pstates {
			pstateentity.Statetid = state
			stcount, err := dataAccess.CheckDuplicateMapprocessstate(&pstateentity)
			if err != nil {
				tx.Rollback()
				return 0, false, err, "Something Went Wrong"
			}
			if stcount.Total == 0 {
				_, err := dao.InsertMapprocessstatewithtransaction(&pstateentity, tx)
				if err != nil {
					tx.Rollback()
					return 0, false, err, "Something Went Wrong"
				}
			}
		}

		wentity.Details = details[0].Details
		wentity.Detailsjson = details[0].Detailsjson
		wentity.Iscomplete = details[0].Iscomplete

		in:=[]byte(details[0].Detailsjson)
		var detailsEntity []entities.ProcessdetailsEntity
		err = json.Unmarshal(in, &detailsEntity)
		if err != nil {
			log.Print(err)
		}

		for _, transition := range transitions {
			tz.Templatetransitionid = transition.Templatetransitionid
			log.Println("template transition before:",tz.Templatetransitionid)
			/**
			Get each state details(group and user) of a particular template transition
			*/
			groups, gerr := dataAccess.Gettemplategroupbytemplatetransition(tz)
			if gerr != nil {
				tx.Rollback()
				return 0, false, gerr, "Something Went Wrong"
			}
			if len(groups) == 0 {
				tx.Rollback()
				return 0, false, nil, "Transition state details not defined for the template"
			}
			//log.Print("Transition state: ",transition.Currentseq,transition.Previousseq,lastinsertedID)
			isProceed := true
			if transition.Previousseq == 0 {
				wentity.Previousstateid = -1
			} else {
				prevs, err := dataAccess.Getstateidbyseq(tz, transition.Previousseq)
				if err != nil {
					tx.Rollback()
					return 0, false, err, "Something Went Wrong"
				}
				if len(prevs) > 0 {
					wentity.Previousstateid = prevs[0].Id
				} else {
					isProceed = false
				}
			}
			if isProceed {
				currs, err := dataAccess.Getstateidbyseq(tz, transition.Currentseq)
				if err != nil {
					tx.Rollback()
					return 0, false, err, "Something Went Wrong"
				}
				if len(currs) > 0 {
					wentity.Currentstateid = currs[0].Id
					states, err := dataAccess.Checkduplicatestate(&wentity)
					if err != nil {
						tx.Rollback()
						return 0, false, err, "Something Went Wrong"
					}
					if len(states) > 0 {

					} else {
						transitionid, err := dao.Createtransition(&wentity, tx)
						if err != nil {
							tx.Rollback()
							return 0, false, err, "Something Went Wrong"
						}
						log.Println("transitionid: ",transitionid)
						log.Println("template transition after:",tz.Templatetransitionid)
						for i, ent := range detailsEntity {
							for j, in := range ent.Instate {
								log.Println("in :",in)
								if in == tz.Templatetransitionid {
									log.Println(" matched in state:",in,transitionid)
									detailsEntity[i].Instate[j] = transitionid
								}
							}
							for j, out := range ent.OutState {
								log.Println("out :",out)
								if out == tz.Templatetransitionid {
									log.Println(" matched out state:",out,transitionid)
									detailsEntity[i].OutState[j] = transitionid
								}
							}
						}
						wentity.Transitionid = transitionid
						for _, group := range groups {
							wentity.Mstuserid = group.Mstuserid
							wentity.Mstgroupid = group.Mstgroupid
							_, err := dao.Inserttransitiondetails(&wentity, tx)
							if err != nil {
								tx.Rollback()
								return 0, false, err, "Something Went Wrong"
							}

						}
					}
				}
			}

		}
		out, err := json.Marshal(detailsEntity)
		if err != nil {
			panic (err)
		}
		wentity.Detailsjson = string(out)
		/**
		Insert template json and xml details for process
		*/
		_, deerr := dao.Insertprocesswithtransaction(&wentity, tx)
		if deerr != nil {
			tx.Rollback()
			return 0, false, deerr, "Something Went Wrong"
		}
		tx.Commit()
	}

	return 0, true, err, ""
	//return 0, false, err, "Something Went Wrong"
}
