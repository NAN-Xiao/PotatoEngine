package netmessage

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	message "potatoengine/src/netmessage/pbmessage"
	"reflect"
)

var PBMessageMap map[int32]interface{}
var PBMesssageHandleMap map[int32] func(interface{})(interface{},interface{})

func init() {
	PBMessageMap = make(map[int32]interface{})
	PBMesssageHandleMap = make(map[int32] func(interface{})(interface{},interface{}))
}

//注册消息到messageMap
func RegistePBNetMessage(msg interface{}) {
	id, ok := GetServerMsgID(msg)
	if ok != nil {
		return
	}
	_, has := PBMessageMap[id]
	if has {
		return
	}
	PBMessageMap[id] = msg
}
// 应该还要注册消息处理函数到MesssageRouterMap
func RegistePBNetMessageHandl(msg interface{},f func(interface{})(interface{},interface{})) {
	id, ok := GetServerMsgID(msg)
	if ok != nil {
		return
	}
	_, has := PBMesssageHandleMap[id]
	if has {
		return
	}
	PBMesssageHandleMap[id] = f
}

func GetProcessFuction(id int32) func(interface{})(interface{},interface{})  {
	f,ok:=PBMesssageHandleMap[id]
	if ok{
		return f
	}
	return nil
}

func GetPbMessage(i int32)  (interface{},bool){
	f,ok:=PBMessageMap[i]
	if ok{
		return f,true
	}
	return nil,false
}

//解析proto
func DeCodePBNetMessage(id int32, data []byte) (interface{}, error) {
	i, ok := PBMessageMap[id]
	if ok {
		msg, ok := i.(proto.Message)
		if ok {
			err := proto.Unmarshal(data, msg)
			if err == nil {
				return msg, nil
			}
		}
	}
	return nil, fmt.Errorf("decode msg :%d fail", id)
}
//序列化proto
func UnCodePBNetMessage(msg interface{}) ([]byte ,error) {
	m,ok:=msg.(proto.Message)
	if ok{
		return  proto.Marshal(m)
	}
	return  nil,fmt.Errorf("%s type not  proto message",reflect.TypeOf(msg).Name())
}
func DefaultNetErrorData() ([]byte,error){
	netError:=&message.NetError{
		ErrorCode: message.EMsg_Error_Unknown,
	}
	return proto.Marshal(netError)
}