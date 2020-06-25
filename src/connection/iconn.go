package connection

import "potatoengine/src/message"

type IConn interface {
	SendMessage(msg *netmessage.Messsage)
	Read()
	Write(data []byte)
	CloseConnection() bool
}
