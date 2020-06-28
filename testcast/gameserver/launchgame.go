package main

import (
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/msg"
	"potatoengine/src/server"
)

//func RegistServerInfo()  {
//
//	netmessage.RegistePBNetMessageID(&message.LoginResquest{})
//	netmessage.RegistePBNetMessageID(&message.LoginResponse{})
//
//
//}

func main() {

	server.LaunchServer()
	sp := NewLoginSpace("Login")
	RegistSpace("Login", sp)
	Serv()
	select {}
}