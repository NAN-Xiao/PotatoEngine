package login

import (
	"net"
	"potatoengine/src/router"
	"potatoengine/src/server"
	"potatoengine/src/space"
)

type LoginServer struct {
	_listener *net.TCPListener
	_space    map[string]*space.BaseSpace
}

func (this *LoginServer) RegisterSpace(sp *space.BaseSpace) {
	if sp == nil {
		return
	}

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

func NewLoginServer() *LoginServer {
	ser := &LoginServer{
		_space: make(map[string]*space.BaseSpace),
	}
	return ser
}
