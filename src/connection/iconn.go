package connection

import (
	"potatoengine/src/netmessage"
)

type IConn interface {
	SendMessage(msg *netmessage.MsgPackage)
	Read()
	Write(data []byte)
	CloseConnection() bool
}
