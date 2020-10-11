package main

import (
	"potatoengine/src/entity"
	"potatoengine/src/netWork/connect"
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
	//创建gateserver添加全局
	gate := GateServer()
	server.AddServer(gate)
	server.Start()
	println("game server started")
	select {}
	println("engine out")
}

//创建gateserver
func GateServer() *server.BaseServer {
	gate := server.NewServer(server.E_Game, connect.ETcp)
	gatesp := new(GateSpace)
	gatesp.BaseSpace = space.BaseSpace{
		GameID:    0,
		SpaceID:   1,
		Spacename: "GateSpace",
		Entitys:   make(map[int32]entity.IEntity),
	}
	gate.RegisterSpace(gatesp)
	return gate
}