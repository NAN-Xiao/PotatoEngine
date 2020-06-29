package netmessage

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/descriptor"
	pb "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
	message "potatoengine/src/netmessage/pbmessage"
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
//打包pbmessage到网络传输的byte数组
func PackageNetMessage(m interface{}) []byte {
	id, err := GetServerMsgID(m)
	fmt.Println(id)
	if err == nil {
		//序列化pb

		msgdata, _ := proto.Marshal(m.(proto.Message))
		//序列化id
		iddata := new(bytes.Buffer)
		binary.Write(iddata, binary.BigEndian, id)

		//序列化长度
		msglen := len(msgdata) + 8
		fmt.Println(msglen)
		len := new(bytes.Buffer)
		binary.Write(len, binary.BigEndian, int32(msglen))

		//组成buff
		buff := len.Bytes()
		buff = append(buff, iddata.Bytes()...)
		buff = append(buff, msgdata...)
		return buff
	}
	return nil
}
