package server

import (
	"fmt"
	//"net"
	"potatoengine/src/space"
)

var GateServer IServer
var loginServer IServer
var ServerMap map[E_ServerNames]BaseServer
var serverMsg map[string] interface{}


func init() {

	//LoginServer,ok:=login.NewServer().(IServer)
	//if ok==false{
	//	return
	//}

	//ServerMap[E_Loging] = LoginServer
	//fmt.Println("launchServer")
}

func AddServer(server BaseServer)  {
	if ServerMap==nil|| len(serverMsg)<=0{
		ServerMap = make(map[E_ServerNames]BaseServer)
	}
	if _,ok:=ServerMap[server.Name];ok==true{
		return
	}
	ServerMap[server.Name]=server
}

//启动服务
func LaunchServer() {
	if ServerMap==nil|| len(ServerMap)<=0{
		print("not have any server")
		return
	}
	for m := range ServerMap {
		ser,ok:=ServerMap[m]
		if ok{
			ser.Run()
		}
	}
}

//给server注册space
func RegistSpace(name E_ServerNames, sp space.ISpace) {

	if ServerMap == nil {
		fmt.Printf("sermap :%s is null", name)
		return
	}
	 serv,ok:=ServerMap[name]
	 if ok==false{
		return
	}
	serv.RegisterSpace(sp)
}


