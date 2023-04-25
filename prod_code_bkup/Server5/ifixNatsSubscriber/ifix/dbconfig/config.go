package dbconfig

// var MASTER_URL = "http://localhost:8082/api"
var MASTER_URL = "http://localhost:8082/api"

//var MASTER_URL = "http://52.172.198.186:8082/api"

// var SlaNotificationURL = "http://localhost:8089/ /iFIXNotification/sendnotification"
var SlaNotificationURL = "https://iccmuat.ifixcloud.io/iFIXNotification/sendnotification"

// UAT
var RECORD_URL = "http://localhost:8083/recordapi"

//var RECORD_URL = "https://10.5.2.6:8083/recordapi"

// var DBDRIVER = "mysql"
// var DBUSER = "gouser"
// var DBPASWORD = "TCSUAT@54321"
// var DBURL = "tcp(10.5.2.4:3306)"
// var DBNAME = "iFIX"

//local
// var DBDRIVER = "mysql"            // Database Driver Name
// var DBUSER = "root"               // Database Username
// var DBPASWORD = "7980161455"      // Database  Password
// var DBURL = "tcp(localhost:3306)" // Database ip/host with port
// var DBNAME = "iFIX"               // Database Name

// Staging
//var DBDRIVER = "mysql"
//var DBUSER = "ifix"
//var DBPASWORD = "Staging@4321"
//var DBURL = "tcp(172.17.0.1:3306)"
//var DBNAME = "iFIX"

//Production
var DBDRIVER = "mysql"
var DBUSER = "gouser"
var DBPASWORD = "#TCSICCiFIXProd@65243"
var DBURL = "tcp(10.5.3.10:3306)"
var DBNAME = "iFIX"
