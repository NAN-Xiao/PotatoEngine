package main

import (
	"net"
	"potatoengine/src/db"
	"potatoengine/src/netmessage/pbmessage"
)

//import "net"

func GateHandl(conn *net.TCPConn) {

}

func CheckLoginToken(m interface{}) (interface{}, interface{}) {
	msg, ok := m.(message.CheckToken)
	if ok == false {
		return nil, &message.NetError{ErrorCode: message.EMsg_Error_CheckTokenFail}
	}
	redis := db.GetRedisManager()
	redisclient, err := redis.GetDB()
	if err == nil {
		serverToken := redisclient.Get(string(msg.Userid)).Val()
		if serverToken == msg.Token {
			response := &message.CheckTokenResult{
				Result: true,
			}
			return response, nil
		}
	}
	return nil, &message.NetError{ErrorCode: message.EMsg_Error_CheckTokenFail}
}
