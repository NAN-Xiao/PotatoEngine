package main

import (
	"potatoengine/src/agent"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
	"potatoengine/src/server"
	"potatoengine/src/space"
)
func RegistServerInfo()  {
	//注册消息
	netmessage.RegistePBNetMessage(&message.LoginResquest{})
	//注册消息处理句柄
	netmessage.RegistePBNetMessageHandl(&message.LoginResquest{},ProcessLoginRequest)
}
func main() {
	RegistServerInfo()


	login:=&server.BaseServer{
		SpacesMap: make(map[string]space.ISpace),
		Name:   server.E_Loging,
	}
	//new space
	sp :=LoginSpace{struct {
		SpaceID    int32
		Spacename  string
		Agents     map[uint32]*agent.Agent
		Spacechanl chan *netmessage.ServerMsgPackage
	}{SpaceID:0 , Spacename: "login", Agents: nil, Spacechanl:nil }}

	login.RegisterSpace(&sp)
	login.Run()
	select {}
}
