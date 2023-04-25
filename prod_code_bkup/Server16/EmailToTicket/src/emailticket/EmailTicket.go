package emailticket

import (
	"log"
	// "os"
	"src/config"
	// ReadProperties "src/fileutils"
	Logger "src/logger"
	"strings"
	"time"

	"github.com/BrianLeishman/go-imap"
)

func CreateTicketByEmail() {
	var baseClientid int64 = 2
	var baseorgID int64 = 2
	Logger.Log.Println("In CreateTicketByEmail")
	// wd, err := os.Getwd() // to get working directory
	// if err != nil {
	// 	Logger.Log.Println(err)
	// }
	// //log.Println(wd)
	// // contextPath := strings.ReplaceAll(wd, "\\", "/") // replacing backslash by  forwardslash
	// //log.Println(contextPath)
	// // props, err := ReadProperties.ReadPropertiesFile(contextPath + "/resource/application.properties")
	// if err != nil {
	// 	Logger.Log.Println(err)
	// }

	//imap.Verbose = true
	// clientID, err := strconv.ParseInt(props["CLIENTID"], 10, 64)
	// if err != nil {
	// 	Logger.Log.Println(err)
	// }
	// orgID, err := strconv.ParseInt(props["ORGID"], 10, 64)
	// if err != nil {
	// 	Logger.Log.Println(err)
	// }
	db, dBerr := config.GetDB()
	if dBerr != nil {
		Logger.Log.Println(dBerr)
	}
	//defer db.Close()
	defer Logger.Log.Println("DB STats=====> ", db.Stats())
	defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
	defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
	defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
	Logger.Log.Println("Login Done")
	IMAPEmailDomain, IMAPPort, UserName, password, err := GetCredential(db, baseClientid, baseorgID)
	// im, err := imap.New(props["UserName"], props["password"], props["IMAPEmailDomain"], 993)
        Logger.Log.Println("Before Login Done")
	im, err := imap.New(UserName, password, IMAPEmailDomain, IMAPPort)
        Logger.Log.Println("After Login Done")
	if err != nil {
		Logger.Log.Println(err)
	}
	Logger.Log.Println(im)
	defer im.Close()

	for {
		Logger.Log.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

		_, err2 := im.GetFolders()
		//log.Println()
		if err2 != nil {
			CreateTicketByEmail()
			//continue
		}
		err = im.SelectFolder("INBOX")
		if err != nil {
			Logger.Log.Println(err)
			CreateTicketByEmail()
			//continue
		}

		uids, err := im.GetUIDs("ALL")
		if err != nil {
			Logger.Log.Println(err)
			CreateTicketByEmail()
			//continue
		}

		//log.Println(uids)
		// db, dBerr := config.GetDB()
		// if dBerr != nil {
		// 	Logger.Log.Println(dBerr)
		// }
		// //defer db.Close()
		// defer Logger.Log.Println("DB STats=====> ", db.Stats())
		// defer Logger.Log.Println("DB STats INUSe conn=====> ", db.Stats().InUse)
		// defer Logger.Log.Println("DB STatsn Idle Conn=====> ", db.Stats().Idle)
		// defer Logger.Log.Println("DB STatsn Open Conn=====> ", db.Stats().OpenConnections)
		// Logger.Log.Println("Login Done")
		var lastUid int = uids[len(uids)-1]
		//log.Println("LastUid==>", lastUid)
		lastUpdatedUID, err := GetLastUpdatedUID(db, baseClientid, baseorgID)
		Logger.Log.Println("LastUid==>", lastUid)
		Logger.Log.Println("lastUpdatedUID==>", lastUpdatedUID)
		if err != nil {
			Logger.Log.Println(err)
		}

		var unseenUID []int
		//var j int
		for i := 0; i < len(uids); i++ {
			if uids[i] > lastUpdatedUID {
				unseenUID = append(unseenUID, uids[i])
			}

		}
		if lastUid > lastUpdatedUID {
			Logger.Log.Println("New Email Received")
			emails, err := im.GetEmails(unseenUID...)
			if err != nil {
				Logger.Log.Println(err)
			}

			Logger.Log.Println("LastUid==>", lastUid)
			for email := range emails {
				tx, err := db.Begin()
				if err != nil {
					Logger.Log.Println(err)
				}
				// senderEmail := emails[email].From.String()
				// if strings.Contains(senderEmail, "<") {
				// 	senderEmail = senderEmail[strings.Index(senderEmail, "<")+1 : len(senderEmail)-1]
				// }
				//Logger.Log.Println("Email Body====>", emails[email].Text)

				//emailbody:= strings.Index(emails[email].Text,"From:")
				Logger.Log.Println("emails[email].Subject===>", emails[email].Subject)
				createdUserID, createdGrpID, clientID, orgID, getClientIDAndOrgIDErr := GetClientIDOrgID(db, emails[email].From.String())
				if getClientIDAndOrgIDErr != nil {
					Logger.Log.Println(getClientIDAndOrgIDErr)
					//  presentUID int = emails[email].UID
					// UpdateUiDvar(db, tx, clientID, orgID, presentUID)
					// continue
				} else {
					Logger.Log.Println("emails[email].Text ===>===============BEFORE================>", emails[email].Text)
					if len(emails[email].Text) > 4900 {
						emails[email].Text = emails[email].Text[0:4900]
					}
					Logger.Log.Println("emails[email].Text=================AFTER=========>", emails[email].Text)
					Logger.Log.Println("Clientid===>", clientID)
					Logger.Log.Println("orgID===>", orgID)

					var emailbody string
					if strings.Contains(emails[email].Text, "---------- Forwarded message ---------") && strings.Contains(emails[email].Subject, "Fwd:") {
						Logger.Log.Println("<===================GMAIL DOMAIN FOR FORWARDED EMAIL===================>")
						emailbody = emails[email].Text[:strings.Index(emails[email].Text, "---------- Forwarded message ---------")-1]

					} else if strings.Contains(emails[email].Text, "\nOn") && strings.Contains(emails[email].Subject, "Re:") {
						Logger.Log.Println("<===================GMAIL DOMAIN FOR REPLY EMAIL===================>")
						emailbody = emails[email].Text[:strings.Index(emails[email].Text, "\nOn")-1]

					} else if strings.Contains(emails[email].Text, "\nFrom:") {

						Logger.Log.Println("<===================OUTLOOK DOMAIN FOR REPLY OR FORWARDED EMAIL===================>")
						emailbody = emails[email].Text[:strings.Index(emails[email].Text, "\nFrom:")-1]

					}
					// if (strings.Contains(emails[email].Subject, "RE:") || strings.Contains(emails[email].Subject, "FW:") {
					// 	Logger.Log.Println("<===================OUTLOOK DOMAIN FOR REPLY OR FORWARDED EMAIL===================>")
					// 	emailbody = emails[email].Text[:strings.Index(emails[email].Text, "From:")-1]

					// }
					// if strings.EqualFold(senderEmail, "kaustubh@ifixtechglobal.com") || strings.EqualFold(senderEmail, "kaustubhsays@gmail.com") {
					// 	Logger.Log.Println("Email Body===>", emailbody)
					// }
					if !strings.EqualFold(emailbody, "") {
						emails[email].Text = emailbody
						//	Logger.Log.Println("<================================Email Body=======================>", emailbody)
						Logger.Log.Println("<======================Email Subject=================>", emails[email].Subject)
						Logger.Log.Println("<======================Email Body=================>", emailbody)
					}
					//For Updated Case:=
					Logger.Log.Println("Sent Timing=====>", emails[email].Sent)
					Logger.Log.Println("Received Timing=====>", emails[email].Received)
					if strings.Contains(emails[email].Subject, "(") {
						countRow, countRowErr := GetDuplicateEmailCount(db, clientID, orgID, emails[email].Subject, emails[email].From.String(), emails[email].Sent.String(), emails[email].Received.String())
						if countRowErr != nil {
							Logger.Log.Println(countRowErr)
						}
						if countRow == 0 {
							startIndex := strings.Index(emails[email].Subject, "(")
							endIndex := strings.Index(emails[email].Subject, ")")
							ticketNo := emails[email].Subject[startIndex+1 : endIndex]
							log.Println(ticketNo)
							recordID, ticketClientID, ticketOrgID, userID, userGrpID, recordStagedID, foundTicket, invalidSender, checkForTicketNumberError := ValidateTicketNowithSenderEmail(db, clientID, orgID, ticketNo, emails[email].From.String())
							if checkForTicketNumberError != nil {
								Logger.Log.Println(checkForTicketNumberError)
							}
							if foundTicket != false && invalidSender != false {
								Logger.Log.Println("<===========================================Update Option On Comment Ticket=====================>")

								Logger.Log.Println("Ticket found with validEmail, recordID==> ", recordID)
								Logger.Log.Println("Ticket found with validEmail, SenderEmail==> ", emails[email].From)
								updateTicketError := UpdateTicket(db, tx, ticketClientID, ticketOrgID, recordID, userID, userGrpID, recordStagedID, emails[email].Subject, emails[email].Text, emails[email].Attachments, emails[email].From.String(), emails[email].Sent.String(), emails[email].Received.String())
								if updateTicketError != nil {
									Logger.Log.Println("Unable To update Ticket===>error==>", updateTicketError)
									//return
								}
								//time.Sleep(2 * time.Second)
							} else if foundTicket != true && invalidSender != false {
								//create new ticket
								Logger.Log.Println("<===========================================Update Option Create Ticket=====================>")
								Logger.Log.Println("Ticket No not found recordID==> ", recordID)

								Logger.Log.Println("CountRow for Ticket Creation===>: ", countRow)
								Logger.Log.Println("Subject: ", emails[email].Subject)
								Logger.Log.Println("Attachments==>", emails[email].Attachments)
								Logger.Log.Println("senderEmail===>", emails[email].From)
								rowID, senderTypeSeq, defaultSeq, validateEmail, _ := ValidateEmail(db, clientID, orgID, emails[email].Subject, emails[email].From.String(), baseClientid, baseorgID)
								//emailvalidationError
								Logger.Log.Println("RowID===>", rowID)
								Logger.Log.Println("senderTypeSeq===>", senderTypeSeq)
								Logger.Log.Println("defaultSeq===>", defaultSeq)
								Logger.Log.Println("validateEmail flag===>", validateEmail)
								if rowID == 0 || senderTypeSeq == 0 || validateEmail == false {
									Logger.Log.Println("<<<<<<====Not A valid Email======>>>>>>")
								} else {
									Logger.Log.Println("To email TO", emails[email].To.String())
									Logger.Log.Println("To email CC", emails[email].CC.String())
									//ccEmailID := emails[email].CC.String()
									ToEmailID := emails[email].To.String()
									toEmaillist := strings.Split(ToEmailID, ",")

									//if len(toEmaillist) == 1 {
									Logger.Log.Println("To email length is having one mailId", len(toEmaillist))

									Logger.Log.Println("To email length is having one mailId", len(toEmaillist))

									if rowID != 0 && senderTypeSeq != 0 && validateEmail != false && defaultSeq != 2 {
										ticketID, createTicketError := CreateTicket(db, tx, emails[email].Attachments, emails[email].Subject, emails[email].Text, rowID, senderTypeSeq, defaultSeq, clientID, orgID, emails[email].From.String(), emails[email].Sent.String(), emails[email].Received.String(), createdUserID, createdGrpID)
										if createTicketError != nil {
											Logger.Log.Println(err)
										} else {
											Logger.Log.Println("Created TicketID===> ", ticketID)
										}

									} else {
										Logger.Log.Println("<<<<<<====Not A valid Email======>>>>>>")
									}
									//}
								}
								time.Sleep(2 * time.Second)

							} else {
								Logger.Log.Println("")
								Logger.Log.Println("")
								Logger.Log.Println("============================Invalid Sender==================>")
								Logger.Log.Println("")
								Logger.Log.Println("")
							}
						} else {
							Logger.Log.Println("Duplicate Email for Update Ticket Section")
						}

					} else {
						countRow, countRowErr := GetDuplicateEmailCount(db, clientID, orgID, emails[email].Subject, emails[email].From.String(), emails[email].Sent.String(), emails[email].Received.String())
						if countRowErr != nil {
							Logger.Log.Println(countRowErr)
						}
						if countRow == 0 {
							Logger.Log.Println("To email TO", emails[email].To.String())
							Logger.Log.Println("To email CC", emails[email].CC.String())
							//ccEmailID := emails[email].CC.String()
							ToEmailID := emails[email].To.String()
							toEmaillist := strings.Split(ToEmailID, ",")

							//
							Logger.Log.Println("To email length is having one mailId", len(toEmaillist))

							Logger.Log.Println("To email length is having one mailId", len(toEmaillist))
							Logger.Log.Println("<===========================================Fresh  Create Ticket=====================>")

							Logger.Log.Println("CountRow for Ticket Creation===>: ", countRow)
							Logger.Log.Println("Subject: ", emails[email].Subject)
							Logger.Log.Println("Attachments==>", emails[email].Attachments)
							Logger.Log.Println("senderEmail===>", emails[email].From)
							rowID, senderTypeSeq, defaultSeq, validateEmail, _ := ValidateEmail(db, clientID, orgID, emails[email].Subject, emails[email].From.String(), baseClientid, baseorgID)
							//emailvalidationError
							Logger.Log.Println("RowID===>", rowID)
							Logger.Log.Println("senderTypeSeq===>", senderTypeSeq)
							Logger.Log.Println("defaultSeq===>", defaultSeq)
							Logger.Log.Println("validateEmail flag===>", validateEmail)
							if rowID == 0 || senderTypeSeq == 0 || validateEmail == false {
								Logger.Log.Println("<<<<<<====Not A valid Email======>>>>>>")
							} else {
								if rowID != 0 && senderTypeSeq != 0 && validateEmail != false && defaultSeq != 2 {
									ticketID, createTicketError := CreateTicket(db, tx, emails[email].Attachments, emails[email].Subject, emails[email].Text, rowID, senderTypeSeq, defaultSeq, clientID, orgID, emails[email].From.String(), emails[email].Sent.String(), emails[email].Received.String(), createdUserID, createdGrpID)
									if createTicketError != nil {
										Logger.Log.Println(err)
									} else {
										Logger.Log.Println("Created TicketID===> ", ticketID)
									}

								} else {
									Logger.Log.Println("<<<<<<====Not A valid Email======>>>>>>")
								}
							}
							time.Sleep(2 * time.Second)
							//}
						} else {
							Logger.Log.Println("Duplicate Email For Create Ticket section......")
						}

					}
					emailbody = ""
				}
				var presentUID int = emails[email].UID
				UpdateUiD(db, tx, presentUID, baseClientid, baseorgID)
				commitErr := tx.Commit()
				if commitErr != nil {
					Logger.Log.Println(err)
				}
			}
			Logger.Log.Println("###########################################")
		} else {
			Logger.Log.Println("No New Email Received")
			Logger.Log.Println("###########################################")

		}

		//log.Println("LastUid==>", lastUid)
		time.Sleep(5 * time.Second)
	}
	// c, err := client.DialTLS(props["IMAPEmailDomain"]+":"+props["IMAPPort"], nil)
	// if err != nil {
	// 	Logger.Log.Println(err)
	// }
	// log.Println("DOMAIN CONNECTED ==> ", props["IMAPEmailDomain"])
	// Logger.Log.Println("DOMAIN CONNECTED ==> ", props["IMAPEmailDomain"])
	// defer c.Logout()
	// for {
	// 	if err := c.Login(props["UserName"], props["password"]); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println("Inside EMAIL Account")

	// }

}
