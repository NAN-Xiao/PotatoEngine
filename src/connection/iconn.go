package connection

type IConn interface {
	//SendMessage(msg netmessage.ServerMsgPackage)
	Read()(l int,err error)
	Write(data []byte)
	Close() bool
	Listen()
}
