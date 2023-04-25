package models

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/pquerna/otp/totp"
	"iFIX/ifix/config"
	"iFIX/ifix/dao"
	"iFIX/ifix/entities"
	"iFIX/ifix/logger"
	"iFIX/ifix/utility"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
	"strings"
)

var lock = &sync.Mutex{}

func Changepassword(tz *entities.LoginEntityReq) (int64, bool, error, string) {
	log.Println("In side model")
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Fetchpasswordbyid(tz)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		pwdMatch := utility.ComparePasswords(values[0].Password, []byte(tz.Oldpassword))
		if pwdMatch {
			log.Print("first: ", tz.Password)
			tz.Password = utility.HashAndSalt([]byte(tz.Password))
			log.Print("second: ", tz.Password)
			errup := dataAccess.Updatepassword(tz)
			if errup != nil {
				return 0, false, errup, "Something Went Wrong"
			} else {
				return 1, true, nil, ""
			}
		} else {
			return 0, false, nil, "Old password not matched."
		}
	} else {
		return 0, false, nil, "User details not found"
	}
}
func Validateusertoken(tz *entities.LoginEntityResp) (int64, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	err, values := dataAccess.Validateusertoken(tz)
	if err != nil {
		return 0, false, err, "Something Went Wrong"
	}
	if len(values) > 0 {
		return 1, true, err, ""
	} else {
		return 0, false, nil, "You session has expired, please re-login."
	}
}
func Generatetoken(tz *entities.LoginEntityReq) ([]entities.LoginEntityResp, bool, error, string) {
	//t := []entities.LoginEntityResp{}
	//db, err := config.ConnectMySqlDb()
	//defer db.Close()
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Checkuser(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		pwdMatch := true
		//log.Println("condition --> ", (tz.Password != "" || tz.Type != "internal"))
		if tz.Password != "" {
			pwdMatch = utility.ComparePasswords(values[0].Password, []byte(tz.Password))
		}
		if pwdMatch {
			tx, err := db.Begin()

			value := []entities.LoginEntityReq{}
			val := entities.LoginEntityReq{}
			val.ID = values[0].Userid
			value = append(value, val)
			tokenString, err := Createandinserttoken(value, tx)
			if err != nil {
				tx.Rollback()
				return nil, false, err, "Something Went Wrong"
			}
			err = tx.Commit()
			if err != nil {
				logger.Log.Print("Token  Statement Commit error", err)
				log.Print("Token  Statement Commit error", err)
				return nil, false, err, ""
			}
			values[0].Token = tokenString
			values[0].Password = ""
			return values, true, nil, ""
		} else {
			return nil, false, nil, "Invalid Credentials"
		}
	} else {
		return nil, false, nil, "User details not found"
	}
}

func Createandinserttoken(values []entities.LoginEntityReq, tx *sql.Tx) (string, error) {
	tokenString, tokerr := utility.CreateToken(values[0].ID)
	if tokerr != nil {
		log.Print("token create error", tokerr)
		logger.Log.Print("token create error", tokerr)
		return "", tokerr
	}
	//logger.Log.Print(" is Locked :", utility.MutexLocked(lock))
	//log.Print(" is Locked :", utility.MutexLocked(lock))
	err := dao.Deletetoken(tx, values[0].ID)
	if err != nil {

		tx.Rollback()
		return "", err
	}
	err = dao.Inserttoken(tx, values[0].ID, tokenString)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return tokenString, nil
}
func VerifyTOTP(tz *entities.LoginEntityReq) ([]entities.LoginEntityResp, bool, error, string) {
	logger.Log.Println("====================================================================VerifyTOTP========================================")
	log.Println("====================================================================VerifyTOTP========================================")
	firstTimeTOPTObject := []entities.LoginEntityResp{}
	val := entities.LoginEntityResp{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()

	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, keyerror := dataAccess.Login(tz)
	if keyerror != nil {
		log.Println("Unable to fetch secretkey", keyerror)
		return nil, false, err, "Something Went Wrong"
	}
	logger.Log.Println("===========values==============", values)
	logger.Log.Println("TOTP=====>", tz.Totp)
	logger.Log.Println("SecretKey====>", values[0].Secretkey)
	log.Println("===========values==============", values)
	log.Println("TOTP=====>", tz.Totp)
	log.Println("SecretKey====>", values[0].Secretkey)
	ok := totp.Validate(string(tz.Totp), values[0].Secretkey)
	if ok {
		logger.Log.Print("TOTP Validated Successfully")
		log.Print("TOTP Validated Successfully")

	} else {
		return nil, false, nil, "Sorry......Invalid OTP"
	}
	updateErr := dataAccess.UpdateUserMFA(tz)
	if updateErr != nil {
		logger.Log.Println(updateErr)
		log.Println(updateErr)
		return firstTimeTOPTObject, false, nil, "Something Went Wrong With Update MFA"
	}
	urls, err1 := dataAccess.Geturlbytype(tz, "logout")
	if err1 != nil {
		logger.Log.Println(err1)
		log.Println(err1)
		return nil, false, err1, "Something Went Wrong"
	}
	val.Url = urls[0].Url
	firstTimeTOPTObject = append(firstTimeTOPTObject, val)

	return firstTimeTOPTObject, true, nil, "You have registered for Two Factor Authentication Successfully. Please Login again..."
}
func FetchMFA(values []entities.LoginEntityReq, db *sql.DB) ([]entities.LoginEntityResp, bool, error, string) {
	firstTimeTOPTObject := []entities.LoginEntityResp{}
	val := entities.LoginEntityResp{}
	val.Clientid = values[0].Clientid
	val.Userid = values[0].ID
	val.OrgnTypeId = values[0].OrgnTypeId
	val.Mstorgnhirarchyid = values[0].Mstorgnhirarchyid
	dataAccess := dao.DbConn{DB: db}
	urls, err1 := dataAccess.Geturlbytype(&values[0], "mfavalidation")
	if err1 != nil {
		logger.Log.Println(err1)
		log.Println(err1)
		return nil, false, err1, "Something Went Wrong"
	}
	val.Url = urls[0].Url
	firstTimeTOPTObject = append(firstTimeTOPTObject, val)
	logger.Log.Println("===========firstTimeTOPTObject==============", firstTimeTOPTObject)
	log.Println("===========firstTimeTOPTObject==============", firstTimeTOPTObject)

	return firstTimeTOPTObject, true, nil, "You Organisation Has Enabled Two Factor Authentication...Please Register for Two Factor Authentication"

}
func EnableTOTEenable(values []entities.LoginEntityReq, db *sql.DB) ([]entities.LoginEntityResp, bool, error, string) {
	logger.Log.Println("====================================================================EnableTOTEenable========================================")
	log.Println("====================================================================EnableTOTEenable========================================")
	firstTimeTOPTObject := []entities.LoginEntityResp{}
	val := entities.LoginEntityResp{}
	val.Clientid = values[0].Clientid
	val.Userid = values[0].ID
	val.OrgnTypeId = values[0].OrgnTypeId
	val.Mstorgnhirarchyid = values[0].Mstorgnhirarchyid

	dataAccess := dao.DbConn{DB: db}
	authURI, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "iFIX",
		AccountName: values[0].UserEmail,
	})
	if err != nil {
		return firstTimeTOPTObject, false, nil, "Something Went Wrong With MFA"
	}
	contextPath, contextPatherr := os.Getwd()
	props, err := utility.ReadPropertiesFile(contextPath + "/ifix/resource/application.properties")
	logger.Log.Println(props)
	log.Println(props)
	if err != nil {
		logger.Log.Println(err)
		log.Println(err)
		return firstTimeTOPTObject, false, err, "Unable to Get URL From utility.ReadPropertiesFile"
	}
	if contextPatherr != nil {
		logger.Log.Println(contextPatherr)
		log.Println(contextPatherr)
		return firstTimeTOPTObject, false, contextPatherr, "Contextpath ERROR"
	}
	//tx, err := db.Begin()
	secretkey := authURI.Secret()
	updateErr := dataAccess.UpdateMstClientUser(&values[0], authURI.String(), secretkey)
	if updateErr != nil {
		//logger.Log.Println(updateErr)
		return firstTimeTOPTObject, false, nil, "Something Went Wrong With Update MFA"
	}
	var buf bytes.Buffer
	img, err := authURI.Image(300, 300)
	imgerr := png.Encode(&buf, img)

	if imgerr != nil {
		logger.Log.Println(imgerr)
		log.Println(imgerr)
		return firstTimeTOPTObject, false, nil, "Something Went Wrong With Update MFA"
	}
	filePath := contextPath + "/ifix/resource/downloads/QRImgGA.png"
	imageSaveErr := imaging.Save(img, filePath)
	if imageSaveErr != nil {
		logger.Log.Println(imgerr)
		log.Println(imgerr)
		return firstTimeTOPTObject, false, nil, "Something Went Wrong With QR MFA"
	}
	var data = entities.FileuploadEntity{}
	data.Clientid = values[0].Clientid
	data.Mstorgnhirarchyid = values[0].Mstorgnhirarchyid
	imgbytes := buf.Bytes()
	// OriginalFileName, UploadedFileName, err := utility.FileUploadAPICall(values[0].Clientid, values[0].Mstorgnhirarchyid, props["fileUploadUrl"], filePath)
	dataDetails, success, err, msg := UploadFileWithConn(&data, imgbytes, filePath, "image/png", db) //utility.FileUploadAPICall(values[0].Clientid, values[0].Mstorgnhirarchyid, props["fileUploadUrl"], filePath)
	if err != nil {
		logger.Log.Println("Error while downloading", "-", err)
		log.Println("Error while downloading", "-", err)
		return firstTimeTOPTObject, false, nil, "Something Went Wrong With QR MFA"
	}
	logger.Log.Println("===========FileUploadMessage==============", msg)
	logger.Log.Println("===========FileUploadSuccss==============", success)
	log.Println("===========FileUploadMessage==============", msg)
	log.Println("===========FileUploadSuccss==============", success)

	OriginalFileName := dataDetails.Originalfile
	UploadedFileName := dataDetails.Filename

	logger.Log.Println("===========OriginalFileName==============", OriginalFileName)
	logger.Log.Println("===========UploadedFileName==============", UploadedFileName)
	log.Println("===========OriginalFileName==============", OriginalFileName)
	log.Println("===========UploadedFileName==============", UploadedFileName)
	val.OriginalFileName = OriginalFileName
	val.UploadedFileName = UploadedFileName
	urls, err1 := dataAccess.Geturlbytype(&values[0], "mfaregistration")
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	logger.Log.Println("===========urls==============", urls[0].Url)
	log.Println("===========urls==============", urls[0].Url)
	val.Url = urls[0].Url
	//fmt.Println(t.OriginalFileName, t.UploadedFileName)
	//return OriginalFileName, UploadedFileName, true, nil, ""

	firstTimeTOPTObject = append(firstTimeTOPTObject, val)
	logger.Log.Println("===========firstTimeTOPTObject==============", firstTimeTOPTObject)
	log.Println("===========firstTimeTOPTObject==============", firstTimeTOPTObject)
	return firstTimeTOPTObject, true, nil, "You Organisation Has Enabled Two Factor Authentication...Please Register for Two Factor Authentication"
}
func fetchlogindetails(values []entities.LoginEntityReq, db *sql.DB, tokenString string) ([]entities.LoginEntityResp, bool, error, string) {
	val := entities.LoginEntityResp{}
	t := []entities.LoginEntityResp{}
	dataAccess := dao.DbConn{DB: db}

	if values[0].Clientid == 1 {
		val.Clientid = values[0].Clientid
		val.Mstorgnhirarchyid = values[0].Mstorgnhirarchyid
		val.OrgnTypeId = values[0].OrgnTypeId
		val.Userid = values[0].ID
		val.Token = tokenString
		t = append(t, val)
		return t, true, nil, ""
	} else {
		roles, err1 := dataAccess.GetRolebyUserId(&values[0])
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		} else {
			if len(roles) == 0 {
				return t, false, nil, "User is not mapped with any role"
			} else {
				roles[0].Clientid = values[0].Clientid
				roles[0].Mstorgnhirarchyid = values[0].Mstorgnhirarchyid
				roles[0].OrgnTypeId = values[0].OrgnTypeId
				roles[0].Token = tokenString
				urls, err1 := dataAccess.Geturlbytype(&values[0], "dashboard")
				if err1 != nil {
					return nil, false, err1, "Something Went Wrong"
				} else {
					if len(urls) > 0 {
						roles[0].Dashboardurl = urls[0].Url
					} else {
						roles[0].Dashboardurl = ""
					}
					urls, err1 := dataAccess.Geturlbytype(&values[0], "ExternalTicket")
					if err1 != nil {
						return nil, false, err1, "Something Went Wrong"
					} else {
						if len(urls) > 0 {
							roles[0].Externalurl = urls[0].Url
						} else {
							roles[0].Externalurl = ""
						}
						return roles, true, nil, ""
					}
					return roles, true, nil, ""
				}

			}
		}
	}
}
func Getorgname(tz *entities.LoginEntityReq) ([]entities.LoginEntityResp, bool, error, string) {
	log.Println("In side Getorgname model")
	logger.Log.Println("In side Getorgname model")
	//logger.Log.Print(" is Locked :", utility.MutexLocked(lock))
	//if utility.MutexLocked(lock) == false {
	lock.Lock()
	defer lock.Unlock()
	//}
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()

	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}

	err1, values := dataAccess.Getorgdetailsbycode(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		t := []entities.LoginEntityResp{}
		v := entities.LoginEntityResp{}
		v.Org = values[0].Org
		t = append(t, v)
		return t, true, nil, ""
	} else {
		return nil, false, nil, "Sorry...Invalid Credentials"
	}
}
func SendOTP(tz *entities.LoginEntityReq) ([]entities.LoginEntityResp, bool, error, string) {
	log.Println("In side model Loginchecking")
	logger.Log.Println("In side model")
	logger.Log.Print(" is Locked :", utility.MutexLocked(lock), "  ", tz.Loginname)
	log.Print(" is Locked :", utility.MutexLocked(lock), "  ", tz.Loginname)
	//if utility.MutexLocked(lock) == false {
	lock.Lock()
	defer lock.Unlock()
	//}
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()

	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Login(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		otp := GenerateCode()
		err := dataAccess.SaveOTP(&values[0], otp)
		if err != nil {
			return nil, false, err1, "Something Went Wrong"
		}
		credEntity := entities.FileuploadEntity{}
		credEntity.Clientid = values[0].Clientid
		credEntity.Mstorgnhirarchyid = values[0].Mstorgnhirarchyid
		credEntity.Credentialtype = "2"
		creds, err := dataAccess.GetCredential(&credEntity)
		if err != nil {
			return nil, false, err1, "Something Went Wrong"
		}
		if len(creds) > 0 {
			var smtpHostForNotification string = creds[0].Credentialkey
			var emailUserName string = creds[0].Credentialaccount
			var emailPassword string = creds[0].Credentialpassword
			var smtpPort string = "587"
			emailSubject := "OTP for Forget Password"
			emailBody := "Dear Sir,\n Your OTP is " + otp + ". This OTP is valid for only 10 mins"
			userTOEmails := values[0].UserEmail
			var userCCEmails string = ""
			sendMailErr := utility.SendMail(emailSubject, emailBody, userTOEmails, userCCEmails, smtpHostForNotification, smtpPort, emailUserName, emailPassword)
			log.Println(sendMailErr)
			logger.Log.Println(sendMailErr)
			return nil, true, nil, ""
		} else {
			return nil, false, nil, "Mail is  not configured for this organization."

		}

	} else {
		return nil, false, nil, "Username not found for this organization."
	}
}
func ForgetPassword(tz *entities.LoginEntityReq) (int64, bool, error, string) {
	log.Println("In side model ForgetPassword")
	logger.Log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()

	if err != nil {
		log.Println("database connection failure", err)
		return 0, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.Login(tz)
	if err1 != nil {
		return 0, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		err, otpdetails := dataAccess.Fetchotpbyid(&values[0])
		if err != nil {
			return 0, false, err1, "Something Went Wrong"
		}
		if len(otpdetails) > 0 {
			if otpdetails[0].OTP == tz.OTP {
				value := entities.LoginEntityReq{}
				value.ID = values[0].ID
				value.Password = utility.HashAndSalt([]byte(tz.Password))
				errup := dataAccess.Updatepassword(&value)
				if errup != nil {
					return 0, false, errup, "Something Went Wrong"
				}
				errup1 := dataAccess.DeleteOTP(otpdetails[0].ID)
				if errup1 != nil {
					return 0, false, errup1, "Something Went Wrong"
				}
				return 1, true, nil, ""
			} else {
				return 0, false, nil, "OTP not found for this user."
			}

		} else {
			return 1, false, nil, "OTP not found for this user."
		}

	} else {
		return 1, false, nil, "Username not found for this organization."
	}
}
func GenerateCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}
func Loginchecking(tz *entities.LoginEntityReq) ([]entities.LoginEntityResp, bool, error, string) {
	log.Println("In side model Loginchecking")
	logger.Log.Println("In side model")
	logger.Log.Print(" is Locked :", utility.MutexLocked(lock), "  ", tz.Loginname)
	log.Print(" is Locked :", utility.MutexLocked(lock), "  ", tz.Loginname)
	//if utility.MutexLocked(lock) == false {
	lock.Lock()
	defer lock.Unlock()
	//}
	db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()

	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}

	err1, values := dataAccess.Getorgdetailsbycode(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	if len(values) > 0 {
		if values[0].Logintypeid == 1 {
			log.Println("LDAP Login")
			logger.Log.Println("LDAP Login")
			tz.Clientid = values[0].Clientid
			tz.Mstorgnhirarchyid = values[0].Mstorgnhirarchyid
			tz.OrgnTypeId = values[0].OrgnTypeId
			details, success, err, msg := Ldaplogin(tz, db)
			log.Println("success: ", success)
			logger.Log.Println("success: ", success)
			if err != nil || !success {
				log.Print(msg)
				logger.Log.Println(msg)
				if msg != "CONNECTION_REFUSED" && msg != "User is not mapped with any role" && values[0].Islocallogin == 2 {
					log.Println("Local1 Login")
					logger.Log.Println("Local1 Login")
					details, success, err, msg := Login(tz, db)
					if err != nil || !success {
						return nil, false, err, msg
					} else {
						details[0].Logintypeid = values[0].Logintypeid
						return details, true, nil, ""
					}
				} else {
					//log.Print(" is Locked :", utility.MutexLocked(lock))
					//logger.Log.Print(" is Locked :", utility.MutexLocked(lock))
					return nil, false, err, msg
				}
			} else {
				logger.Log.Println(details)
				log.Println(details)
				details[0].Logintypeid = values[0].Logintypeid
				logger.Log.Println("sending..... ")
				log.Println("sending..... ")
				return details, true, nil, ""
			}
		} else {
			log.Println("Local Login")
			logger.Log.Println("Local Login")
			details, success, err, msg := Login(tz, db)
			if err != nil || !success {
				return nil, false, err, msg
			} else {
				details[0].Logintypeid = values[0].Logintypeid
				return details, true, nil, ""
			}
			//return nil, false, nil, "No login type mapped.."
		}
	} else {
		return nil, false, nil, "Sorry...Invalid Credentials"
	}
}

func Inpersonatelogin(tz *entities.LoginEntityReq) ([]entities.LoginEntityResp, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//
	//defer db.Close()
	//dataAccess := dao.DbConn{DB: db}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		log.Println("Transaction creation error.", err)
		return nil, false, err, "Something Went Wrong"
	}
	values := []entities.LoginEntityReq{}
	if tz.ID > 0 {
		value := entities.LoginEntityReq{}
		value.Clientid = tz.Clientid
		value.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
		value.ID = tz.ID
		values = append(values, value)
		token, err := Createandinserttoken(values, tx)
		if err != nil {
			tx.Rollback()
			return nil, false, err, "Something Went Wrong"
		}
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("Token  Statement Commit error", err)
			log.Print("Token  Statement Commit error", err)
			return nil, false, err, ""
		}
		details, success, err, msg := fetchlogindetails(values, db, token)
		if err != nil {
			return nil, false, err, "Something Went Wrong"
		} else {
			if success {
				return details, true, nil, ""
			} else {
				return nil, false, nil, msg
			}
		}
	} else {
		return nil, false, nil, "User not found"
	}
}

func Login(tz *entities.LoginEntityReq, db *sql.DB) ([]entities.LoginEntityResp, bool, error, string) {
	log.Println("In side model")
	//lock.Lock()
	//defer lock.Unlock()
	//db, err := config.ConnectMySqlDbSingleton()
	//db, err := config.ConnectMySqlDb()
	//if err != nil {
	//	log.Println("database connection failure", err)
	//	return nil, false, err, "Something Went Wrong"
	//}
	//defer db.Close()
	dataAccess := dao.DbConn{DB: db}
	tx, err := db.Begin()
	if err != nil {
		logger.Log.Println("Transaction creation error.", err)
		return nil, false, err, "Something Went Wrong"
	}
	values, err1 := dataAccess.Login(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}

	//logger.Log.Println("===========values==============", values)
	if len(values) > 0 {
		values[0].Loginname = tz.Loginname
		pwdMatch := utility.ComparePasswords(values[0].Password, []byte(tz.Password))
		if pwdMatch || (tz.OrgMFA == 1 && tz.UserMFA == 1) {
			//after passwordmatches we have to write the code for authienticator

			token, err := Createandinserttoken(values, tx)
			if err != nil {
				tx.Rollback()
				return nil, false, err, "Something Went Wrong"
			}
			err = tx.Commit()
			if err != nil {
				logger.Log.Print("Token  Statement Commit error", err)
				return nil, false, err, ""
			}

			if tz.OrgMFA == 1 && tz.UserMFA == 1 {
				ok := totp.Validate(string(tz.Totp), values[0].Secretkey)
				if ok {
					logger.Log.Print("OTP Validated Successfully")
					log.Print("OTP Validated Successfully")

				} else {
					return nil, false, nil, "Sorry......Invalid OTP"
				}
			} else if values[0].UserMFA == 1 && values[0].OrgMFA == 1 {
				//ok := totp.Validate(string(tz.Totp), values[0].Secretkey)
				details, success, err, msg := FetchMFA(values, db)
				details[0].UserMFA = values[0].UserMFA
				details[0].OrgMFA = values[0].OrgMFA
				details[0].Token = token
				if err != nil {
					return nil, false, err, "Something Went Wrong"
				} else {
					if success {
						return details, true, nil, ""
					} else {
						return nil, false, nil, msg
					}
				}

			} else if values[0].UserMFA == 2 && values[0].OrgMFA == 1 {
				details, success, err, msg := EnableTOTEenable(values, db)
				details[0].OrgMFA = 1
				details[0].UserMFA = 2
				details[0].Token = token
				if err != nil {
					return nil, false, err, "Something Went Wrong"
				} else {
					if success {
						return details, true, nil, ""
					} else {
						return nil, false, nil, msg
					}
				}
			}

			details, success, err, msg := fetchlogindetails(values, db, token)
			if err != nil {
				return nil, false, err, "Something Went Wrong"
			} else {
				if success {
					return details, true, nil, ""
				} else {
					return nil, false, nil, msg
				}
			}
		} else {
			return nil, false, nil, "Sorry......Invalid Credentials"
		}
	} else {
		return nil, false, nil, "Sorry...Invalid Credentials"
	}
}

func Adfslogin(result map[string]interface{}) ([]entities.LoginEntityResp, bool, error, string) {
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	var issuer string
	var nameID string
	dataAccess := dao.DbConn{DB: db}
	for _, value := range result["details"].([]interface{}) {
		if value.(map[string]interface{})["key"] == "issuer" {
			issuer = value.(map[string]interface{})["value"].(string)

		}
		if value.(map[string]interface{})["key"] == "nameID" {
			nameID = value.(map[string]interface{})["value"].(string)

		}
	}
	log.Println(issuer, nameID)
	err, orgdetails := dataAccess.Getorgdetailsfromadfs(issuer)
	if err != nil {
		return nil, false, err, "Something Went Wrong"
	} else {
		if len(orgdetails) > 0 {
			tz := entities.LoginEntityReq{}
			tz.Clientid = orgdetails[0].Clientid
			tz.Mstorgnhirarchyid = orgdetails[0].Mstorgnhirarchyid
			tz.Code = orgdetails[0].Code
			tz.Loginname = nameID
			tz.Password = nameID + "$12345"
			details, success, err, msg := upsertexternaluser(&tz, db, result, 3)
			return details, success, err, msg
			return nil, false, nil, "No organization is mapped"
		} else {
			return nil, false, nil, "No organization is mapped"
		}
	}
}
func Ldaplogin(tz *entities.LoginEntityReq, db *sql.DB) ([]entities.LoginEntityResp, bool, error, string) {
	log.Println("In side model")
	dataAccess := dao.DbConn{DB: db}

	fileuploadentity := entities.FileuploadEntity{}
	fileuploadentity.Clientid = tz.Clientid
	fileuploadentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
	x, err := dataAccess.GetCredentialById(&fileuploadentity)
	if err != nil {
		return nil, false, err, "Something Went Wrong"
	}
	//log.Print(x.Credentialaccount, x.Credentialpassword, x.Credentialkey)
	//logger.Log.Print(x.Credentialaccount, x.Credentialpassword, x.Credentialkey)

	ldapentity := entities.MstldapEntity{}
	ldapentity.Clientid = tz.Clientid
	ldapentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
	ldapentity.Offset = 0
	ldapentity.Limit = 1
	ldapdetails, err := dataAccess.GetAllMstldap(&ldapentity)
	if err != nil {
		return nil, false, err, "Something Went Wrong"
	} else {
		if len(ldapdetails) > 0 {
			postBody, _ := json.Marshal(map[string]string{"username": tz.Loginname, "password": tz.Password, "servername": ldapdetails[0].ServerName, "serverurl": ldapdetails[0].ServerUrl, "binddn": ldapdetails[0].Binddn, "basedn": ldapdetails[0].Basedn, "bindpassword": ldapdetails[0].Password, "filterdn": ldapdetails[0].Filterdn, "chncertificate": ldapdetails[0].Chn_Certificate, "credentialaccount": x.Credentialaccount, "credentialpassword": x.Credentialpassword, "credentialkey": x.Credentialkey})

			responseBody := bytes.NewBuffer(postBody)
			logger.Log.Println("postBody       --->", tz.Loginname, ldapdetails[0].ServerName, ldapdetails[0].ServerUrl, ldapdetails[0].Binddn)
			log.Println("postBody       --->", tz.Loginname, ldapdetails[0].ServerName, ldapdetails[0].ServerUrl, ldapdetails[0].Binddn)
			resp, err := http.Post(config.LDAP_URL+"/verifyldapuser", "application/json", responseBody)
			if err != nil {
				logger.Log.Println("An Error Occured --->", err)
				log.Println("An Error Occured --->", err)
				return nil, false, err, "Login Server Error"
			}
			defer resp.Body.Close()
			//Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logger.Log.Println("response body ------> ", err)
				log.Println("response body ------> ", err)
				return nil, false, err, "Login Server Fetch Error"
			}
			sb := string(body)
			log.Println(sb)
			logger.Log.Println(sb)
			//var result entities.LdapAttrEntityResponse
			var result map[string]interface{}
			json.Unmarshal(body, &result)
			log.Println("Success : ", result["success"])
			logger.Log.Println("Success : ", result["success"])

			details, ok := result["success"]
			if !ok {
				log.Println("Not Present")
			}
			var success = details.(bool)
			if success {
				details, success, err, msg := upsertexternaluser(tz, db, result, 1)
				return details, success, err, msg

			} else {
				return nil, false, err, result["message"].(string)
			}
		} else {
			return nil, false, err, "LDAP Configuration is not done yet"
		}
	}

}
func upsertexternaluser(tz *entities.LoginEntityReq, db *sql.DB, result map[string]interface{}, systemid int64) ([]entities.LoginEntityResp, bool, error, string) {
	logger.Log.Println("inside upsertexternaluser")
	log.Println("inside upsertexternaluser")
	dataAccess := dao.DbConn{DB: db}
	attr := entities.MapexternalattributesEntity{}
	attr.Systemid = systemid
	attr.Clientid = tz.Clientid
	attr.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
	attrs, attrerr := dataAccess.GetMappedattributes(&attr)
	user := entities.MstClientUserEntity{}
	user.Vipuser = "N"
	if attrerr != nil {
		return nil, false, attrerr, "Something Went Wrong"
	}
	logger.Log.Println("mapped attr len:", len(attrs))
	log.Println("mapped attr len:", len(attrs))
	if len(attrs) > 0 {
		for _, mapped := range attrs {
			if mapped.Sysattr == "firstname" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Firstname = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "lastname" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Lastname = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "usermobileno" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Usermobileno = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "division" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Division = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "useremail" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Useremail = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "secondaryno" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Secondaryno = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "designation" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Designation = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "branch" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Branch = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "brand" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.Brand = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "city" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == mapped.Extattr {
						user.City = extattrs.(map[string]interface{})["value"].(string)
						break
					}
				}
			}
			if mapped.Sysattr == "vipuser" {
				for _, extattrs := range result["details"].([]interface{}) {
					if extattrs.(map[string]interface{})["key"] == "description" {
						var value = extattrs.(map[string]interface{})["value"].(string)
						val1 := strings.ToUpper(value[0:1])
						val, _ := strconv.Atoi(value[1:2])
						if val1 == "M" {
						if val < 5 && val > 0 {
							user.Vipuser = "Y"
						}
				  		}
						log.Println("----> ", user.Vipuser)
						logger.Log.Println("----> ", user.Vipuser)
						break
					}
				}
			}

		}
		//log.Println(user.Lastname, user.Usermobileno)
		values, err1 := dataAccess.Login(tz)
		if err1 != nil {
			return nil, false, err1, "Something Went Wrong"
		}
		user.Loginname = tz.Loginname

		user.Usertype = "NA"
		logger.Log.Println("new user:", len(values))
		log.Println("new user:", len(values))
		if len(values) > 0 && values[0].UserMFA == 1 {
			user.Mfa = 1
		} else {
			user.Mfa = 2
		}
		user.Password = utility.HashAndSalt([]byte(tz.Password))
		user.ClientID = tz.Clientid
		user.MstorgnhirarchyID = tz.Mstorgnhirarchyid
		user.Createtype = systemid
		tx, err := db.Begin()
		logger.Log.Println("----------Transaction Begin :------------")
		log.Println("----------Transaction Begin :------------")
		if err != nil {
			logger.Log.Println("Transaction creation error.", err)
			return nil, false, err, "Something Went Wrong"
		}
		if len(values) > 0 {
			logger.Log.Println("values length :  ", values[0].ID)
			log.Println("values length :  ", values[0].ID)
			user.ID = values[0].ID
			err = dao.UpdateClientUserData(tx, &user)
			if err != nil {
				tx.Rollback()
				return nil, false, err, "Something Went Wrong"
			}
			logger.Log.Println("----MstClientUser data updated------")
			log.Println("----MstClientUser data updated------")
			err = dao.UpdateMstUserData(tx, &user)
			if err != nil {
				tx.Rollback()
				return nil, false, err, "Something Went Wrong"
			}
			logger.Log.Println("----Mstuser data updated------")
			log.Println("----Mstuser data updated------")
			err = dao.Updatepasswordtransaction(tx, user.Password, user.ID)
			if err != nil {
				tx.Rollback()
				return nil, false, err, "Something Went Wrong"
			}
			logger.Log.Println("----MstClientUser password updated------")
			log.Println("----MstClientUser password updated------")
			err = dao.Updateuserpasswordtransaction(tx, user.Password, user.ID)
			if err != nil {
				tx.Rollback()
				return nil, false, err, "Something Went Wrong"
			}
			logger.Log.Println("----MstUser password updated------")
			log.Println("----MstUser password updated------")
		} else {

			count, err := dao.CheckDuplicateCientUser(tx, &user)
			if err != nil {
				tx.Rollback()
				return nil, false, err, "Something Went Wrong"
			}
			if count == 0 {
				userid, err := dao.InsertClientUserData(tx, &user)
				if err != nil {
					tx.Rollback()
					return nil, false, err, "Something Went Wrong"
				}
				if userid > 0 {
					log.Print("user added in transaction")
					logger.Log.Print("user added in transaction")
					count1, err := dao.CheckDuplicateMstUser(tx, &user)
					if err != nil {
						tx.Rollback()
						return nil, false, err, "Something Went Wrong"
					}
					if count1 == 0 {
						mstuserid, err := dao.InsertMstUserData(tx, &user, userid)
						if err != nil {
							tx.Rollback()
							return nil, false, err, "Something Went Wrong"
						}
						if mstuserid > 0 {
							log.Print("user added in workflow")
							logger.Log.Print("user added in workflow")
							groupentity := entities.MapldapgrouproleEntity{}
							groupentity.Clientid = tz.Clientid
							groupentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
							groupentity.Limit = 1
							groupentity.Offset = 0
							mapping, err := dataAccess.GetAllmapldapgrouprole(&groupentity, systemid)
							if err != nil {
								tx.Rollback()
								return nil, false, err, "Something went wrong"
							}
							if len(mapping) > 0 {
								memberentity := entities.GroupmemberEntity{}
								memberentity.Clientid = tz.Clientid
								memberentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
								memberentity.Groupid = mapping[0].Groupid
								memberentity.Refuserid = userid

								count, err := dataAccess.CheckDuplicateGroupmember(&memberentity)
								if err != nil {
									tx.Rollback()
									return nil, false, err, "Something Went Wrong"
								}
								if count.Total == 0 {
									log.Print("\nGroup Added:")
									logger.Log.Print("\nGroup Added:")
									_, err := dataAccess.InsertGroupmembertransaction(&memberentity, tx)
									if err != nil {
										tx.Rollback()
										return nil, false, err, "Something Went Wrong"
									}
									roleentity := entities.MapClientUserRoleUserEntity{}
									roleentity.ClientID = tz.Clientid
									roleentity.MstorgnhirarchyID = tz.Mstorgnhirarchyid
									roleentity.RoleID = mapping[0].Roleid
									roleentity.Refuserid = userid
									count, err := dataAccess.CheckDuplicateMapRoleUser(&roleentity)
									if err != nil {
										tx.Rollback()
										return nil, false, err, "Something Went Wrong"
									}
									if count.Total > 0 {
										tx.Rollback()
										return nil, false, nil, "Mapping already exist."
									}
									roleid, err := dataAccess.InsertMapRoleUserDataTransaction(&roleentity, tx)
									if err != nil {
										tx.Rollback()
										return nil, false, err, "Something Went Wrong"
									}
									log.Print("\nRole Added:", roleid)
									logger.Log.Print("\nRole Added:", roleid)

									loginvalues := []entities.LoginEntityReq{}
									loginval := entities.LoginEntityReq{}
									loginval.Clientid = tz.Clientid
									loginval.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
									loginval.OrgnTypeId = tz.OrgnTypeId
									loginval.ID = userid
									loginvalues = append(loginvalues, loginval)
									values = loginvalues
								} else {
									tx.Rollback()
									return nil, false, nil, "Data Already Exist."
								}
							} else {
								tx.Rollback()
								return nil, false, nil, "Support Group Not Found "
							}
						} else {
							tx.Rollback()
							return nil, false, nil, "User Already Exist."
						}
					} else {
						tx.Rollback()
						return nil, false, nil, "User Already Exist.."
					}
				} else {
					tx.Rollback()
					return nil, false, nil, "User Already Exist..."
				}

			} else {
				tx.Rollback()
				return nil, false, nil, "User Already Exist...."
			}
		}
		log.Println(tz.Code, tz.Loginname)

		token, err := Createandinserttoken(values, tx)
		if err != nil {
			tx.Rollback()
			return nil, false, err, "Something Went Wrong"
		}
		logger.Log.Println("Token Creation Successful:")
		log.Println("Token Creation Successful:")
		err = tx.Commit()
		if err != nil {
			logger.Log.Print("Token  Statement Commit error", err)
			return nil, false, err, ""
		}

		mfaresp, _, _, msg := LoginMFA(tz, db)
		if err != nil {
			return nil, false, err, msg
		}
		if len(mfaresp) > 0 {
			mfaresp[0].Token = token
			urls, err1 := dataAccess.Geturlbytype(&values[0], "ExternalTicket")
			if err1 != nil {
				return nil, false, err1, "Something Went Wrong"
			} else {
				if len(urls) > 0 {
					mfaresp[0].Externalurl = urls[0].Url
				} else {
					mfaresp[0].Externalurl = ""
				}
				return mfaresp, true, nil, ""
			}

		} else {
			details, success, err, msg := fetchlogindetails(values, db, token)
			logger.Log.Println("------------------------------", msg)
			logger.Log.Println(details)
			log.Println("------------------------------", msg)
			log.Println(details)
			if err != nil {
				return nil, false, err, msg
			} else {
				if !success {
					return nil, false, err, msg
				}
				logger.Log.Println("----Transaction Committed------")
				log.Println("----Transaction Committed------")
				return details, true, nil, ""
			}
		}
	} else {
		log.Println("No attributes is mapped with system")
		logger.Log.Println("No attributes is mapped with system")
		return nil, false, nil, "No attributes is mapped with system"
	}

}
func Getldapattributes(tz *entities.LoginEntityReq) (entities.LdapAttrEntityResponse, bool, error, string) {
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	t := entities.LdapAttrEntityResponse{}
	if err != nil {
		log.Println("database connection failure", err)
		return t, false, err, "Something Went Wrong"
	}
	//defer db.Close()
	//
	dataAccess := dao.DbConn{DB: db}
	fileuploadentity := entities.FileuploadEntity{}
	fileuploadentity.Clientid = tz.Clientid
	fileuploadentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
	x, err := dataAccess.GetCredentialById(&fileuploadentity)
	if err != nil {
		return t, false, err, "Something Went Wrong"
	}
	log.Print(x.Credentialaccount, x.Credentialpassword, x.Credentialkey)
	logger.Log.Print(x.Credentialaccount, x.Credentialpassword, x.Credentialkey)

	ldapentity := entities.MstldapEntity{}
	ldapentity.Clientid = tz.Clientid
	ldapentity.Mstorgnhirarchyid = tz.Mstorgnhirarchyid
	ldapentity.Offset = 0
	ldapentity.Limit = 1
	ldapdetails, err := dataAccess.GetAllMstldap(&ldapentity)
	if err != nil {
		return t, false, err, "Something Went Wrong"
	} else {
		if len(ldapdetails) > 0 {
			if err != nil {
				logger.Log.Println("Transaction creation error.", err)
				return t, false, err, "Something Went Wrong"
			}
			postBody, _ := json.Marshal(map[string]string{"username": tz.Loginname, "password": tz.Password, "servername": ldapdetails[0].ServerName, "serverurl": ldapdetails[0].ServerUrl, "binddn": ldapdetails[0].Binddn, "basedn": ldapdetails[0].Basedn, "bindpassword": ldapdetails[0].Password, "filterdn": ldapdetails[0].Filterdn, "chncertificate": ldapdetails[0].Chn_Certificate, "credentialaccount": x.Credentialaccount, "credentialpassword": x.Credentialpassword, "credentialkey": x.Credentialkey})

			responseBody := bytes.NewBuffer(postBody)
			//logger.Log.Println("postBody       --->", responseBody)
			resp, err := http.Post(config.LDAP_URL+"/getldapattributes", "application/json", responseBody)
			if err != nil {
				logger.Log.Println("An Error Occured --->", err)
				log.Println("An Error Occured --->", err)
				return t, false, err, "Login Server Error"
			}
			defer resp.Body.Close()
			//Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logger.Log.Println("response body ------> ", err)
				log.Println("response body ------> ", err)
				return t, false, err, "Login Server Fetch Error"
			}
			sb := string(body)
			logger.Log.Println("sb change group body value is --->", sb)
			log.Println("sb change group body value is --->", sb)
			var result entities.LdapAttrEntityResponse
			json.Unmarshal(body, &result)
			if result.Success {
				//log.Print(result.Details)
				return result, true, nil, ""
			} else {
				return t, false, err, result.Message
			}
		} else {
			return t, false, err, "LDAP Configuration is not done yet"
		}
	}
}
func LoginMFA(tz *entities.LoginEntityReq, db *sql.DB) ([]entities.LoginEntityResp, bool, error, string) {
	log.Println("In side model")
	details := []entities.LoginEntityResp{}
	dataAccess := dao.DbConn{DB: db}

	values, err1 := dataAccess.Login(tz)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}

	if len(values) > 0 {
		values[0].Loginname = tz.Loginname
		logger.Log.Println("===========id==============", values[0].ID)
		logger.Log.Println("===========clientid==============", values[0].Clientid)
		logger.Log.Println("===========mstorgnhirarchyid==============", values[0].Mstorgnhirarchyid)
		logger.Log.Println("===========orgnTypeId==============", values[0].OrgnTypeId)
		logger.Log.Println("===========password==============", values[0].Password)
		logger.Log.Println("===========useremail==============", values[0].UserEmail)
		logger.Log.Println("===========usermfa==============", values[0].UserMFA)
		logger.Log.Println("===========orgmfa==============", values[0].OrgMFA)
		logger.Log.Println("===========Secretkey==============", values[0].Secretkey)
		logger.Log.Println("===========loginname==============", values[0].Loginname)
		logger.Log.Println("===========USER==============", tz.UserMFA)
		logger.Log.Println("===========ORG==============", tz.OrgMFA)
		logger.Log.Println(">>>>>>>>>>>>>>>>", tz)
		log.Println(">>>>>>>>>>>>>>>>", tz)
		log.Println(">>>>>>>>>>>1>>>>>", values[0])
		if tz.OrgMFA == 1 && tz.UserMFA == 1 {
			ok := totp.Validate(string(tz.Totp), values[0].Secretkey)
			if ok {
				logger.Log.Print("OTP Validated Successfully")
				log.Print("OTP Validated Successfully")

				return details, true, nil, "OTP Validated Successfully"

			} else {
				return nil, false, nil, "Sorry......Invalid OTP"
			}
		} else if values[0].UserMFA == 1 && values[0].OrgMFA == 1 {
			//ok := totp.Validate(string(tz.Totp), values[0].Secretkey)
			detail, success, err, msg := FetchMFA(values, db)
			details = detail
			details[0].UserMFA = values[0].UserMFA
			details[0].OrgMFA = values[0].OrgMFA
			details[0].Loginname = tz.Loginname
			// details[0].Token = token
			if err != nil {
				return nil, false, err, "Something Went Wrong"
			} else {
				if success {
					return details, true, nil, ""
				} else {
					return nil, false, nil, msg
				}
			}

		} else if values[0].UserMFA == 2 && values[0].OrgMFA == 1 {
			logger.Log.Println("===========usermfa==============", values[0].UserMFA)
			logger.Log.Println("===========orgmfa==============", values[0].OrgMFA)
			logger.Log.Println("===========Secretkey==============", values[0].Secretkey)
			logger.Log.Println("===========loginname==============", values[0].Loginname)
			logger.Log.Println("===========USER==============", tz.UserMFA)
			logger.Log.Println("===========ORG==============", tz.OrgMFA)
			detail, success, err, msg := EnableTOTEenable(values, db)
			details = detail
			details[0].OrgMFA = 1
			details[0].UserMFA = 2
			details[0].Loginname = tz.Loginname
			// details[0].Token = token
			if err != nil {
				return nil, false, err, "Something Went Wrong"
			} else {
				if success {
					return details, true, nil, ""
				} else {
					return nil, false, nil, msg
				}
			}
		} else {
			return details, true, nil, "MFA is not enable for this Organization"
		}
	} else {
		return nil, false, nil, "Sorry...Invalid Credentials"
	}
	//return nil, false, nil, "Something Went Wrong"
}
func GetUserDetailsById(tz *entities.UserEntity) ([]entities.UserEntity, bool, error, string) {
	log.Println("In side model")
	t := []entities.UserEntity{}
	val := entities.LoginEntityReq{}
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//defer db.Close()

	dataAccess := dao.DbConn{DB: db}
	values, err1 := dataAccess.GetUserDetailsById(tz)
	if err1 != nil {
		return t, false, err1, "Something Went Wrong"
	}
	//log.Println("values:",len(values),tz.Userid)
	if len(values) > 0 {
		val.ID = tz.Userid
		roles, err1 := dataAccess.GetRolebyUserId(&val)
		if err1 != nil {
			return t, false, err1, "Something Went Wrong"
		} else {
			if len(roles) == 0 {
				return t, false, nil, "User is not mapped with any role"
			} else {
				values[0].Rolename = roles[0].Rolename
				values[0].Roleid = roles[0].Roleid
				values[0].IsAdmin = roles[0].IsAdmin
				properties, err3 := dataAccess.Getpropertybyroleid(&values[0])
				//log.Println("length---->",len(properties))
				if err3 != nil {
					return nil, false, err3, "Something Went Wrong"
				} else {
					values[0].Propertyids = properties
				}
				groups, err2 := dataAccess.Getgroupbyuserid(&val)
				if err2 != nil {
					return t, false, err2, "Something Went Wrong"
				}
				values[0].Group = groups
				if len(groups) > 0 {
					def, err3 := dataAccess.Getdefaultgroupbyid(&values[0], tz.Userid)
					if err3 != nil {
						return t, false, err3, "Something Went Wrong"
					}
					log.Print("---> ", len(def))
					if len(def) > 0 {
						values[0].Deafultgroup = def[0].Deafultgroup
					} else {
						values[0].Deafultgroup = 0
					}
				}
				urls, err2 := dataAccess.Geturlbyuser(&values[0])
				if err2 != nil {
					return t, false, err2, "Something Went Wrong"
				}
				values[0].Urls = urls
				acn := entities.UserroleactionnameEntity{}
				acn.Clientid = values[0].Clientid
				acn.Mstorgnhirarchyid = values[0].Mstorgnhirarchyid
				acn.UserID = tz.Userid
				actions, _, err2, _ := GetUserActionname(db, &acn)
				if err2 != nil {
					return t, false, err2, "Something Went Wrong"
				} else {
					if len(actions) > 0 {
						for i := 0; i < len(actions); i++ {
							if actions[i] == 1 {
								values[0].Add = true
							} else if actions[i] == 2 {
								values[0].Delete = true
							} else if actions[i] == 3 {
								values[0].View = true
							} else {
								values[0].Edit = true
							}
						}
						return values, true, nil, ""
					} else {
						return values, true, nil, ""
					}
				}

			}
		}
	} else {
		return t, false, nil, "User Details Not Found"
	}
}
func Logout(tz *entities.UserEntity) ([]entities.UserEntity, bool, error, string) {
	log.Println("In side model")
	lock.Lock()
	defer lock.Unlock()
	db, err := config.ConnectMySqlDbSingleton()
	if err != nil {
		log.Println("database connection failure", err)
		return nil, false, err, "Something Went Wrong"
	}
	//defer db.Close()

	dataAccess := dao.DbConn{DB: db}
	err1 := dataAccess.Deleteuserticketbyuserid(tz.Userid)
	if err1 != nil {
		return nil, false, err1, "Something Went Wrong"
	}
	return nil, true, nil, ""
}
