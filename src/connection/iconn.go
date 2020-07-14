package connection

type IConn interface {
	//SendMessage(msg netmessage.ServerMsgPackage)
	Receive() error
	ReadMsg() interface{}
	GetMsgQue() chan interface{}
	Write(data []byte)
	Close() bool
}
