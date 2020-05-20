package connection

import "potatoengine/src/message"

type IConn interface {
	SendMessage(msg *message.Messsage)
	Read()
	Write(data []byte)
	CloseConnection() bool
}
