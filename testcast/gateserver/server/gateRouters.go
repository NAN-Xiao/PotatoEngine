package main

import (
	"net"
	"potatoengine/src/db"
	"potatoengine/src/netmessage/pbmessage"
)

//import "net"

func GateHandl(conn *net.TCPConn) {

}

func CheckLoginToken(m interface{}) (interface{}, error) {
	msg, ok := m.(message.CheckToken)
	if ok == false {
		return nil, nil
	}
	redisclient, err :=  db.GetRedisManager().GetDB()
	if err != nil||redisclient==nil {
		return nil,nil
	}
	serverToken := redisclient.Get(string(msg.Userid)).Val()
	if serverToken == msg.Token {
		response := &message.CheckTokenResult{
			Result: true,
		}
		return response, nil
	}
	return nil, nil
}
