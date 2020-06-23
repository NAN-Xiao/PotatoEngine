package message

import (
	"fmt"
	"github.com/golang/protobuf/descriptor"
	pb "github.com/golang/protobuf/proto"
	"potatoengine/src/message/msg"
	"reflect"
)

func GetServerMsgID(msg interface{}) (string, error) {
	des, ol := msg.(descriptor.Message)
	if ol == false {
		s := fmt.Errorf("msg is not a descriptor message")
		return "", s
	}
	_, md := descriptor.MessageDescriptorProto(des)
	if pb.HasExtension(md.GetOptions(), message.E_ServerMsgID) {
		ext, _ := pb.GetExtension(md.GetOptions(), message.E_ServerMsgID)
		return fmt.Sprint(ext), nil
	}
	s := fmt.Errorf("%s not find extension for msgID", reflect.TypeOf(msg).Name())
	return "", s
}
