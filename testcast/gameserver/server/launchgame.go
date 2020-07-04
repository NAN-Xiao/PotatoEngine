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
	//注册msg
	netmessage.RegistePBNetMessage(&message.CheckToken{})
	//注册msghandle
	netmessage.RegistePBNetMessageHandl(&message.CheckTokenResult{}, CheckLoginToken)
}
func main() {
	RegistServerInfo()
	//new server
	game := &server.BaseServer{
		SpacesMap: make(map[string]space.ISpace),
		Name:      server.E_Loging,
		Conn:      &connection.TcpConnect{},
	}
	//new space
	gatasp := GateSpace{struct {
		SpaceID    int32
		Spacename  string
		Agents     map[uint32]*agent.Agent
		Spacechanl chan netmessage.ServerMsgPackage
	}{SpaceID: 0, Spacename: "GateSpace", Agents: make(map[uint32]*agent.Agent), Spacechanl: make(chan netmessage.ServerMsgPackage)}}

	game.RegisterSpace(&gatasp)
	game.Run()
	println("game server started")
	select {}
}
