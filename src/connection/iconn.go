package connection

import "net"

type IConn interface {
	//SendMessage(msg netmessage.ServerMsgPackage)
	Read(buf []byte)
	Write(data []byte)
	Close() bool
	Listen()
	GetConnection() net.Conn
}
