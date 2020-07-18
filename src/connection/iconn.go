package connection

type IConn interface {
	//SendMessage(msg netmessage.ServerMsgPackage)
	Receive()
	Read() interface{}
	Send(interface{})
	Write(interface{})
	WriteToNet()
	Close()
	IsClosed() bool
}
