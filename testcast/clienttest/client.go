package main

import (
	"fmt"
	"github.com/golang/protobuf/descriptor"
	pb "github.com/golang/protobuf/proto"
	"net"
	message "potatoengine/src/message/msg"
)

func main() {


	var msg interface{}=&message.LoginResponse{

	}
	des, ol := msg.(descriptor.Message)
	if ol == false {
		 fmt.Errorf("msg is not a descriptor message")
	}
	_, md := descriptor.MessageDescriptorProto(des)
	if pb.HasExtension(md.GetOptions(), message.E_ServerMsgID) {
		ext, _ := pb.GetExtension(md.GetOptions(), message.E_ServerMsgID)
		fmt.Sprint(ext)
	}



	fmt.Println("begin")
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("addr is err")
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("conn is err:%v\n", err)
		return
	}
	go func() {
		for {
			conn.Write([]byte("hello"))
		}
	}()

	select {}
}
