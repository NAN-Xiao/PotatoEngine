package main

import (
	"potatoengine/src/agent"
	"potatoengine/src/connection"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
	"potatoengine/src/server"
	"potatoengine/src/space"
)

func RegistServerInfo() {

	netmessage.RegistePBNetMessage(&message.LoginResquest{})
	netmessage.RegistePBNetMessage(&message.LoginResponse{})
}

func main() {
	RegistServerInfo()
	//new server
	game := &server.BaseServer{
		SpacesMap: make(map[string]space.ISpace),
		Name:   server.E_Loging,
		Conn: &connection.TcpConnect{},
	}
	//new space
	gatasp :=GateSpace{struct {
		SpaceID    int32
		Spacename  string
		Agents     map[uint32]*agent.Agent
		Spacechanl chan netmessage.ServerMsgPackage
	}{SpaceID: 0, Spacename: "GateSpace", Agents: make(map[uint32]*agent.Agent), Spacechanl: make(chan netmessage.ServerMsgPackage)}}

	game.RegisterSpace(&gatasp)
	game.Run()
	println("game server started")
	select{}
}
