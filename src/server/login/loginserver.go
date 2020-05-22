package login

import (
	"net"
	"potatoengine/src/router"
	"potatoengine/src/server"
)

type LoginServer struct {
	_listener *net.TCPListener
	logrouter *router.IRouter
}

func (this *LoginServer) RegisterLoginRouter(rt *router.IRouter) {
	if rt == nil {
		return
	}
	this.logrouter = rt
}

func (this *LoginServer) Initialize() {

}

//内部调用
func (this *LoginServer) Begin() {

}

func (this *LoginServer) Start() {
	this.Begin()
}
func (this *LoginServer) Stop() {

}

func NewLoginServer() *server.IServer {
	ser := &LoginServer{logrouter: nil}
	return ser
}
