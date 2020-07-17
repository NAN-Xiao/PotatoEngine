package connection

type IConn interface {
	//SendMessage(msg netmessage.ServerMsgPackage)
	Receive() error
	Read() interface{}
	Write(interface{})
	WriteToNet()
	Close()
	IsClosed() bool
}
