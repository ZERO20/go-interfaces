package dbservice

import "fmt"

// MySql
type mysqldb struct {
	user     string
	password string
	host     string
	port     uint16
	dbName   string
}

// cosnstructor for export
func NewMySqlDB(user string, password string, host string, port uint16, dbName string) *mysqldb {
	return &mysqldb{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		dbName:   dbName,
	}
}

// Check methods for implement
var _ DbConn = &mysqldb{}

// Connect Implement
func (mdb *mysqldb) Connect() {
	fmt.Printf("Connecting to mysql://%v:%v@%v:%v/%v \n",
		mdb.user,
		mdb.password,
		mdb.host,
		mdb.port,
		mdb.dbName,
	)
}

func (mdb *mysqldb) Close() {
	fmt.Printf("Disconnecting mysql")
}

func (mdb *mysqldb) CreateTable(name string) {
	fmt.Printf("Creating table %v in database %v \n", name, mdb.dbName)
}
