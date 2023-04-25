package config

var SR_SEQ int64 = 2
var STASK_SEQ int64 = 3
var CR_SEQ int64 = 4
var CTASK_SEQ int64 = 5
var INCIDENT_SEQ int64 = 1

var CANCEL_SEQ int64 = 11
var RESOLVE_SEQ int64 = 3
var NEW_SEQ int64 = 1
var CLOSE_SEQ int64 = 8
var PENDING_USER_STATUS_SEQ int64 = 5
var PENDING_REQUESTER_INPUT int64 = 13
var REJECTED_STATUS_SEQ int64 = 14
var PENDING_APPROVAL_STATUS_SEQ int64 = 12

var STATUS_SEQ int64 = 2

// LOCAL

//var DBDRIVER = "mysql"
//var DBUSER = "root"
//var DBPASWORD = "admin"
//var DBURL = "tcp(127.0.0.1:3306)"
//var DBNAME = "iFIX"
//var RECORD_URL = "http://localhost:8083/recordapi"
//var EMAIL_URL = "http://localhost:8089/iFIXNotification"
//var TOKEN_EXPIRE_TIME = 3600
//var JWT_KEY = []byte("ifix4@secret#token%^@")
//var LDAP_URL = "http://20.198.64.145:8084/ldapapi"

// UAT TCS

//var DBDRIVER = "mysql"
//var DBUSER = "gouser"
//var DBPASWORD = "TCSUAT@54321"
//var DBURL= "tcp(10.5.2.4:3306)"
//var DBNAME ="iFIX"
//var RECORD_URL="http://localhost:8083/recordapi"
//var EMAIL_URL="http://localhost:8089/iFIXNotification"
//var TOKEN_EXPIRE_TIME = 3600
//var JWT_KEY = []byte("ifix4@secret#token%^@")
//var LDAP_URL = "http://20.198.64.145:8084/ldapapi"

//STAGING TCS

//var DBDRIVER = "mysql"
//var DBUSER = "ifix"
//var DBPASWORD = "Staging@4321"
//var DBURL = "tcp(172.17.0.1:3306)"
//var DBNAME = "iFIX"
//var RECORD_URL = "http://localhost:8083/recordapi"
//var EMAIL_URL = "http://localhost:8089/iFIXNotification"
//var TOKEN_EXPIRE_TIME = 3600
//var JWT_KEY = []byte("ifix4@secret#token%^@")
//var LDAP_URL = "http://20.198.64.145:8084/ldapapi"


// PROD TCS

var DBDRIVER = "mysql"
var DBUSER = "gouser"
var DBPASWORD = "#TCSICCiFIXProd@65243"
var DBURL = "tcp(10.5.3.10:3306)"
var DBNAME = "iFIX"
var RECORD_URL = "http://localhost:8083/recordapi"
var EMAIL_URL = "http://localhost:8089/iFIXNotification"
var TOKEN_EXPIRE_TIME = 3600
var JWT_KEY = []byte("ifix4@secret#token%^@")
var LDAP_URL = "http://10.5.3.14:8084/ldapapi"

// UAT WISEL

//var DBDRIVER = "mysql"
//var DBUSER = "wisel"
//var DBPASWORD = "Wisel@54321"
//var DBURL= "tcp(10.210.1.7:3306)"
//var DBNAME ="iFIX"
//var RECORD_URL="http://localhost:8083/recordapi"
//var EMAIL_URL="http://localhost:8089/iFIXNotification"
//var TOKEN_EXPIRE_TIME = 3600
//var JWT_KEY = []byte("ifix4@secret#token%^@")
//var LDAP_URL = "http://20.198.64.145:8084/ldapapi"

//STAGING WISEL

//var DBDRIVER = "mysql"
//var DBUSER = "wisel"
//var DBPASWORD = "Wisel@54321"
//var DBURL = "tcp(10.210.2.14:3306)"
//var DBNAME = "iFIX"
//var RECORD_URL = "http://localhost:8083/recordapi"
//var EMAIL_URL = "http://localhost:8089/iFIXNotification"
//var TOKEN_EXPIRE_TIME = 3600
//var JWT_KEY = []byte("ifix4@secret#token%^@")
//var LDAP_URL = "http://20.198.64.145:8084/ldapapi"

// PROD WISEL

//var DBDRIVER = "mysql"
//var DBUSER = "wiselprod"
//var DBPASWORD = "WISELProd@4321"
//var DBURL= "tcp(10.210.2.9:3306)"
//var DBNAME ="iFIX"
//var RECORD_URL="http://localhost:8083/recordapi"
//var EMAIL_URL="http://localhost:8089/iFIXNotification"
//var TOKEN_EXPIRE_TIME = 3600
//var JWT_KEY = []byte("ifix4@secret#token%^@")
//var LDAP_URL = "http://20.198.64.145:8084/ldapapi"

// STAGING L&T

//var DBDRIVER = "mysql"
//var DBUSER = "ifixpoc"
//var DBPASWORD = "ifixpoc@4321"
//var DBURL= "tcp(10.7.0.7:3306)"
//var DBNAME ="iFIX"
//var RECORD_URL="http://localhost:8083/recordapi"
//var EMAIL_URL="http://localhost:8089/iFIXNotification"
//var TOKEN_EXPIRE_TIME = 3600
//var JWT_KEY = []byte("ifix4@secret#token%^@")
//var LDAP_URL = "http://20.198.64.145:8084/ldapapi"


var GetRecordfulldetailsUrl = "http://localhost:8082/api/getrecordfulldetails"

var GetjsonUrl = "http://localhost:8083/recordapi/recordfullresult"
var GridResultUrl = "http://localhost:8083/recordapi/recordgridresult"
var FileUploadUrl = "http://localhost:8082/api/fileupload"
var DownloadFileURL = "http://localhost:8082/api/filedownload"
