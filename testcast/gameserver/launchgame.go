package main

import (
	"net"
	"potatoengine/src/agent"
	"potatoengine/src/connection"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/msg"
	"potatoengine/src/server"
	"potatoengine/src/space"
)

func RegistServerInfo() {

	//netmessage.RegistePBNetMessageID(&message.LoginResquest{})
	//netmessage.RegistePBNetMessageID(&message.LoginResponse{})
}

func main() {

	RegistServerInfo()
	game := &server.BaseServer{
		Spaces: make(map[string]space.ISpace),
		Name:   server.E_Loging,
	}
	//new space
	sp := GateSpace{struct {
		SpaceID    int32
		Spacename  string
		Agents     map[uint32]*agent.Agent
		Spacechanl chan netmessage.ServerMsgPackage
	}{SpaceID: 0, Spacename: "GateSpace", Agents: make(map[uint32]*agent.Agent), Spacechanl: make(chan netmessage.ServerMsgPackage)}}

	game.RegisterSpace(&sp)
	game.Run()
	select {}
}
