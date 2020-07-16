package main

import (
	"potatoengine/src/agent"
	"potatoengine/src/connection"
	"potatoengine/src/engine"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
	"potatoengine/src/server"
)

func RegistServerInfo() {
	//注册msg
	netmessage.RegistePBNetMessage(&message.CheckToken{})
	//注册msghandle
	netmessage.RegistePBNetMessageHandl(&message.CheckTokenResult{}, CheckLoginToken)
}
func main() {
	RegistServerInfo()
	//new space
	game := server.NewServer(server.E_Game, connection.ETcp)
	gatasp := GateSpace{struct {
		SpaceID    int32
		Spacename  string
		Agents     map[uint32]*agent.Agent
		Spacechanl chan netmessage.ServerMsgPackage
	}{SpaceID: 0, Spacename: "GateSpace", Agents: make(map[uint32]*agent.Agent), Spacechanl: make(chan netmessage.ServerMsgPackage)}}

	game.RegisterSpace(&gatasp)


	engine.AddServer(game)
	engine.Start()
	println("game server started")
	select {}
}
