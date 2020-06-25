package netmessage

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

var MessageMap map[int32]interface{}
var MesssageRouterMap map[int32]func()

func init() {
	MessageMap = make(map[int32]interface{})
	MesssageRouterMap = make(map[int32]func())
}

//注册消息到messageMap
func RegisteNetMessageID(msg interface{}) {
	id, ok := GetServerMsgID(msg)
	if ok != nil {
		return
	}
	_, has := MessageMap[id]
	if has {
		return
	}
	MessageMap[id] = msg
}
//todo 应该还要注册消息处理函数到MesssageRouterMap


//解析proto
func DeCodeNetMessage(id int32, data []byte) (interface{}, error) {
	i, ok := MessageMap[id]
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
func UnCodeNetMessage(msg interface{}) ([]byte ,error) {
	m,ok:=msg.(proto.Message)
	if ok{
		return  proto.Marshal(m)
	}
	return  nil,fmt.Errorf("%s type not  proto message",reflect.TypeOf(msg).Name())
}
