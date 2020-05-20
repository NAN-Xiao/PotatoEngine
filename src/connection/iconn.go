package connection

type IConn interface {
	Read()
	Write()
	CloseConnection()
	NewConnection()
}
