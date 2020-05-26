package login

import (
	"fmt"
	"net"
	"potatoengine/src/server"
	"potatoengine/src/space"
)

type LoginServer struct {
	server.BaseServer
}

//为当前服务注册space


func (this *LoginServer) Initialize() {

}

//内部调用
func (this *LoginServer) Begin() {
	ok := this.RunSpace()
	if ok == false {
		fmt.Println("not have any space")
		return
	}
	fmt.Println("GaterServer Started")
}

func (this *LoginServer) Run() {
	this.Begin()
}
func (this *LoginServer) Stop() {

}

//启动所有space的go

///新建LoginServer
func NewLoginServer() server.IServer {
	sr := &LoginServer{struct {
		Listener *net.TCPListener
		Spaces   map[string]space.ISpace
	}{Listener: nil, Spaces: make(map[string]space.ISpace)}}
	return sr
}
