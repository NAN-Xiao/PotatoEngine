package netmessage

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

var PBMessageMap map[int32]interface{}
var PBMesssageRouterMap map[int32]func()

func init() {
	PBMessageMap = make(map[int32]interface{})
	PBMesssageRouterMap = make(map[int32]func())
}

//注册消息到messageMap
func RegistePBNetMessageID(msg interface{}) {
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
//todo 应该还要注册消息处理函数到MesssageRouterMap


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
