package main

import (
	"fmt"
	"potatoengine/src/netmessage"
	"potatoengine/src/server"
	login "potatoengine/src/server/loginServer"
)
func RegistServerInfo()  {

	netmessage.RegistePBNetMessageID(&LoginRequest{})
	netmessage.RegistePBNetMessageID(&LoginResponse{})


}
func main() {
	RegistServerInfo()

	login:=login.NewServer()
	login.Name=server.E_Loging
	//ser,ok:=login.BaseServer
	//if ok{
		server.AddServer(login.BaseServer)
	//}

	sp := NewLoginSpace("Login")
	server.RegistSpace(server.E_Loging, sp)

	server.LaunchServer()
	//fmt.Println("child", &login)
	b:= &login.BaseServer
	fmt.Println("base", b)
	//fmt.Println("server started")
	//select {}
}
