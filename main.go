package main

import (
	"interfaces/dbservice"
)

func main() {
	var dbConn dbservice.DbConn
	// MySql
	/*dbConn = dbservice.NewMySqlDB(
		"user1",
		"Password123",
		"localhost",
		5432,
		"users",
	)*/
	// SQLite
	dbConn = dbservice.NewSqliteDB("/var/wwww/users.sqlite")
	dbConn.Connect()
	dbConn.CreateTable("users")
	dbConn.Close()
}
