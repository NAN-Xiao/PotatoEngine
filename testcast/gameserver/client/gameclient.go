package main

import (
	"github.com/golang/protobuf/proto"
	"net"
	message "potatoengine/src/netmessage/pbmessage"
)

func main() {
	post := "0.0.0.0:9000"
	addr,err:=net.ResolveTCPAddr("tcp",post)
	if err!=nil{
		return
	}
	conn,err:=net.DialTCP("tcp",nil,addr)
	defer conn.Close()
	if err!=nil{
		return
	}
	request:=&message.LoginResquest{
		Username: "xiaonan",
		Password: "123456",
	}
	data,err:=proto.Marshal(request)
	if err!=nil{
		return
	}
	conn.Write(data)
}
