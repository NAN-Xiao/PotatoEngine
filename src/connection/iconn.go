package connection

type IConn interface {
	//SendMessage(msg netmessage.ServerMsgPackage)
	Read()
	Write(data []byte)
	Close() bool
	Listen()
}
