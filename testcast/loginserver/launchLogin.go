package main

import (
	"potatoengine/src/entity"
	"potatoengine/src/netmessage"
	message "potatoengine/src/netmessage/pbmessage"
	"potatoengine/src/server"
	"potatoengine/src/space"
)
func RegistServerInfo()  {
	//注册消息
	netmessage.RegistePBNetMessage(&message.LoginResquest{})


	//注册消息处理句柄
	netmessage.RegistePBNetMessageHandl(&message.LoginResquest{}, ProcessLoginRequest)
}
func main() {
	RegistServerInfo()
	login:=&server.BaseServer{
		SpacesMap: make(map[string]space.ISpace),
		Name:   server.E_Loging,
	}
	sp:=new(LoginSpace)
	sp.BaseSpace=space.BaseSpace{
		GameID:    0,
		SpaceID:   0,
		Spacename: "LoginSpace",
		Entitys:   make(map[int32] entity.IEntity,0),
	}
	login.RegisterSpace(sp)
	login.Run()
	select {}
}
