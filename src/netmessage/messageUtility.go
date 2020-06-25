package netmessage

import (
	"fmt"
	"github.com/golang/protobuf/descriptor"
	pb "github.com/golang/protobuf/proto"
	message "potatoengine/src/netmessage/msg"
	"reflect"
)

func GetServerMsgID(msg interface{}) (int32, error) {
	des, ol := msg.(descriptor.Message)
	if ol == false {
		s := fmt.Errorf("msg is not a descriptor message")
		return -1, s
	}
	_, md := descriptor.MessageDescriptorProto(des)
	if pb.HasExtension(md.GetOptions(), message.E_ServerMsgID) {
		ext, _ := pb.GetExtension(md.GetOptions(), message.E_ServerMsgID)
		s := fmt.Sprint(ext)
		mid, ok := message.ServerMsg_ID_value[s]
		if ok {
			return mid, nil
		}
	}
	s := fmt.Errorf("%s not find extension for msgID", reflect.TypeOf(msg).Name())
	return -1, s
}
