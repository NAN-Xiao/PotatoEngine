package message

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)
//客户端和服务器端通讯消息
type Messsage struct {
	//消息id
	_msgid MsgID
	_pbmsg *proto.Message
}

func (msg *Messsage) GetData() *proto.Message {
	return msg._pbmsg
}

func (msg *Messsage) GetID() MsgID {
	return msg._msgid
}

func NewMessage(id MsgID, data *proto.Message) *Messsage {
	if id < 1000 || data == nil {
		fmt.Println("creat new message error becase the pb msg is nill")
		return nil
	}
	//todo
	//填充message数据
	msg := &Messsage{
		_msgid: id,
		_pbmsg: data,
	}
	//new一个新的PBmessage
	//msg._len = uint32(len(data))
	return msg
}
