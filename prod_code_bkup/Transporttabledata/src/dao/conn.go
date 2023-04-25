package dao

import "database/sql"

//DbConn is used for initialized connection
type DbConn struct {
	DB *sql.DB
}
type TxConn struct {
	//DB *sql.DB
	TX *sql.Tx
}
