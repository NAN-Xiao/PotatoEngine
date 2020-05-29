package message

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
)

/////////////////////////
//用于服务器模块之间的通讯//
/////////////////////////
type MsgPackage struct {
	//用户id
	_uid uint32
	//角色id
	_playerID uint32
	//消息id
	_mid uint32
	//网络消息
	_msg Messsage
}

func (this *MsgPackage) GetMsgID() uint32 {
	return this._mid
}
func (this *MsgPackage)GetMessage()*Messsage  {
	return &this._msg
}

//打包成服务器间的消息模块
func PackMessagePackage(uid uint32, pid uint32, msg Messsage) *MsgPackage {
	pack := &MsgPackage{
		_uid:      uid,
		_playerID: pid,
		_msg:      msg,
	}
	return pack
}

//////////////////////////
//客户端和服务器端通讯消息//
/////////////////////////
type Messsage struct {
	//消息id
	_msgid uint32
	//消息体
	_pbmsg *proto.Message
}

//message的函数
func (msg *Messsage) GetData() *proto.Message {
	return msg._pbmsg
}

//获取消息ID
func (msg *Messsage) GetID() uint32 {
	return msg._msgid
}

//把当前message打包成网络消息
func (this *Messsage) Pack() []byte {

	b, err := proto.Marshal(*this._pbmsg)
	if err != nil {
		fmt.Println("message pack is error::%s", err)
		return nil
	}
	var id, ln, data []byte
	binary.BigEndian.PutUint32(id, uint32(this._msgid))
	binary.BigEndian.PutUint32(ln, uint32(len(b)))
	data = append(id, ln...)
	data = append(data, b...)
	return data
}

//新建一个message
func NewMessage(mid uint32, data []byte) *Messsage {
	if mid < 1000 || data == nil {
		fmt.Println("creat new message error becase the pb msg is nill")
		return nil
	}
	//todo
	//填充message数据
	msg := &Messsage{
		_msgid: mid,
		_pbmsg: Unmarshal(mid, data),
	}
	return msg
}

///处理消息
func Unmarshal(id uint32, data []byte) *proto.Message {
	v, ok := NetMessageType[id]
	if ok != true {
		return nil
	}
	obj := proto.Clone(*v)
	err := proto.Unmarshal(data, obj)
	if err != nil {
		return nil
	}
	return &obj
}
