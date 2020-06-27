package gameServer

import (
	"potatoengine/src/server"
	"potatoengine/src/space"
)

type GameServer struct {
	server.BaseServer
	serverName server.E_ServerNames
}


func (this *GameServer) Run() {
	ok := this.RunSpace()
	if ok == false {
		return
	}
}
func (this *GameServer) Stop() {

}

//启动所有space的go
func NewGameServer() server.IServer {
	sr := &GameServer{
		server.BaseServer{
			Listener: nil,
			Spaces:   make(map[string]space.ISpace),
		},
	}
	return sr
}
