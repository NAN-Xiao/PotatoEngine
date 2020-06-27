package server

import (
	"fmt"
	login "potatoengine/src/server/loginServer"

	//"net"
	"potatoengine/src/space"
)

var GateServer IServer
var LoginServer IServer
var ServerMap map[E_ServerNames]IServer
var serverMsg map[string] interface{}


func Init() {


	LoginServer,ok:=login.NewServer().(IServer)
	if ok==false{
		return
	}
	ServerMap = make(map[E_ServerNames]IServer)
	ServerMap[E_Loging] = LoginServer
	fmt.Println("launchServer")
}
//启动服务
func LaunchServer() {
	for m := range ServerMap {
		ServerMap[m].Run()
	}
}

//给server注册space
func RegistSpace(name E_ServerNames, sp space.ISpace) {

	if ServerMap == nil {
		fmt.Printf("sermap :%s is null", name)
		return
	}
	for m := range ServerMap {
		if m == name {
			ServerMap[m].RegisterSpace(sp)
		}
	}
}


