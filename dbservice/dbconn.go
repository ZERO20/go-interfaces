package dbservice

type DbConn interface {
	Connect()
	Close()
	CreateTable(name string)
}
