package dbservice

import "fmt"

type sqlitedb struct {
	file string
}

func NewSqliteDB(file string) *sqlitedb {
	return &sqlitedb{
		file: file,
	}
}

var _ DbConn = &sqlitedb{}

func (hsqldb *sqlitedb) Connect() {
	hsqldb.file = "/opt/users.sqlite"
	fmt.Println("Connecting to ", hsqldb.file)
}

func (hsqldb *sqlitedb) Close() {
	fmt.Println("Disconnecting from ", hsqldb.file)
}

func (hsqldb *sqlitedb) CreateTable(name string) {
	fmt.Printf("Creating table %v in database %v \n", name, hsqldb.file)
}
