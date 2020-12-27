package main

import (
	"potatoengine/src/entity"
	"potatoengine/src/netWork/connect"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
	"potatoengine/src/server"
	"potatoengine/src/space"
)
var gate_server *server.BaseServer
func RegistServerInfo() {
	//注册msg
	netmessage.RegistePBNetMessage(&message.CheckToken{})
	//注册msghandle
	netmessage.RegistePBNetMessageHandl(&message.CheckTokenResult{}, CheckLoginToken)
}
func RegistSpace()  {
	gatesp := new(GateSpace)
	gatesp.BaseSpace = space.BaseSpace{
		GameID:    0,
		SpaceID:   1,
		Spacename: "GateSpace",
		Entitys:   make(map[int32]entity.IEntity),
	}
	gate_server.RegisterSpace(gatesp)
}

func main() {
	gate_server = server.NewServer(server.E_Game, connect.ETcp)
	RegistServerInfo()
	//创建gateserver添加全局
	server.AddServer(gate_server)
	server.Start()
	println("game server started")
	select {}
	println("engine out")
}
