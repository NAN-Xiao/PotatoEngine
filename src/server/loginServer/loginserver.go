package login

import (
	"potatoengine/src/server"
)

type LoginServer struct {
	serverName server.E_ServerNames
	base server.BaseServer
}

//启动space
func (this *LoginServer) Run() {
	this.base.Run()
}
func (this *LoginServer) Stop() {
	this.base.Stop()
}

func NewServer() interface{} {

	sp:=&LoginServer{
		serverName: server.E_Loging,
	}
	return sp
}



