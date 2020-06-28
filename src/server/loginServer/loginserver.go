package login

import (
	"potatoengine/src/server"
	"potatoengine/src/space"
)

type LoginServer struct {
	server.BaseServer
}

func NewServer() *LoginServer {

	sp :=&LoginServer{struct {
		//Listener *net.TCPListener
		Spaces   map[string]space.ISpace
		Name     server.E_ServerNames
	}{Spaces:make( map[string]space.ISpace) , Name: server.E_Loging}}
	return sp
}
