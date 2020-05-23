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

func NewMessage(pbmsg *proto.Message) *Messsage {
	if pbmsg == nil {
		fmt.Println("creat new message error becase the pb msg is nill")
		return nil
	}
	//pb:=*pb
	msg := &Messsage{
		_pbmsg: pbmsg,
		//_body: data,
	}
	//msg._len = uint32(len(data))
	return msg
}
