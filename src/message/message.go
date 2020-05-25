package message

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

type Messsage struct {
	//消息id
	_mid uint32
	//用户id
	_uid uint32
	//playerid
	_pid   uint32
	_body  []byte
	_pbmsg *proto.Message
}

func (msg *Messsage) GetData() []byte {
	return msg._body
}

func (msg *Messsage) GetID() uint32 {
	return msg._mid
}

func NewMessage(id int, data []byte) *Messsage {
	if id<1000||data == nil {
		fmt.Println("creat new message error becase the pb msg is nill")
		return nil
	}
	//todo
	//填充message数据
	msg:=&Messsage{
		_mid:   0,
		_uid:   0,
		_pid:   0,
		_body:  nil,
		_pbmsg: nil,
	}
	//new一个新的PBmessage
	//msg._len = uint32(len(data))
	return msg
}
