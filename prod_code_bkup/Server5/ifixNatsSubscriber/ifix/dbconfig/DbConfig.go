package dbconfig

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//ConnectMySqlDb is used for db connection
// func ConnectMySqlDb() (db *sql.DB, err error) {
//      dbDriver := "mysql"             // Database Driver Name
//      dbUser := "ifix"                // Database Username
//      dbPassword := "Staging@4321"    // Database  Password
//      dbUrl := "tcp(172.17.0.1:3306)" // Database ip/host with port
//      dbName := "iFIX"                // Database Name
//      db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@"+dbUrl+"/"+dbName)
//      return
// }

// func ConnectMySqlDb() (db *sql.DB, err error) {
//         dbDriver := "mysql"           // Database Driver Name
//         dbUser := "gouser"            // Database Username
//         dbPassword := "TCSUAT@54321"  // Database  Password
//         dbUrl := "tcp(10.5.2.4:3306)" // Database ip/host with port
//         dbName := "iFIX"              // Database Name
//         db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@"+dbUrl+"/"+dbName)
//         return

// }

/* func ConnectMySqlDb() (db *sql.DB, err error) {
         dbDriver := "mysql"                   // Database Driver Name
	 dbUser := "gouser"                    // Database Username
	 dbPassword := "#TCSICCiFIXProd@65243" // Database  Password
	 dbUrl := "tcp(10.5.3.10:3306)"        // Database ip/host with port
	 dbName := "iFIX"                      // Database Name
         db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@"+dbUrl+"/"+dbName)
         return

}*/

var db *sql.DB = nil

func ConnectMySqlDb() (*sql.DB, error) {

	if db == nil {
		d, err := sql.Open(DBDRIVER, DBUSER+":"+DBPASWORD+"@"+DBURL+"/"+DBNAME)
		if err != nil {
			// panic(err.Error())
			return nil, err
		}
		db = d
	}
	return db, nil

}
