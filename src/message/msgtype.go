package message

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

///为了将来容错。msgid从10001开始
var NetMessageType map[uint32]*proto.Message

func RegisteredNetMessage(msg *proto.Message) {
	if NetMessageType == nil {
		NetMessageType = make(map[uint32]*proto.Message)
	}
	v := reflect.ValueOf(msg)
	msgID := v.FieldByName("msgID")
	if msgID.Kind() == reflect.Uint32 {
		var id = msgID.Interface().(uint32)
		for ms := range NetMessageType {
			if ms == id {
				fmt.Println("current msg hava same id")
				return
			}
		}
		NetMessageType[id] = msg
	}
	fmt.Println("registered message is error bec not find msgid")
}

