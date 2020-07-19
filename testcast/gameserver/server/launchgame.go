package main

import (
	"potatoengine/src/agent"
	"potatoengine/src/netWork"
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

	//创建gateserver添加全局
	gate:=GateServer()
	engine.AddServer(gate)
	engine.Start()
	println("game server started")
	select {}
	println("engine out")
}

//创建gateserver
func GateServer() *server.BaseServer {
	gate := server.NewServer(server.E_Game, netWork.ETcp)
	gatasp := GateSpace{struct {
		SpaceID    int32
		Spacename  string
		Agents     map[uint32]*agent.Agent
		Spacechanl chan netmessage.ServerMsgPackage
	}{SpaceID: 0, Spacename: "GateSpace", Agents: make(map[uint32]*agent.Agent), Spacechanl: make(chan netmessage.ServerMsgPackage)}}

	gate.RegisterSpace(&gatasp)
	return gate
}