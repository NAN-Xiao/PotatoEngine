package game

import (
	"potatoengine/src/router"
)

type GameServer struct {
	_router *router.IRouter

}

func (this *GameServer) RegisterSpace(rt *router.IRouter) {
	this.router = rt
}
func (this *GameServer) RegisterRouter(rt *router.IRouter) {
	this.router = rt
}
func (this *GameServer) Initialize() {

}
func (this *GameServer) Begin() {

}
func (this *GameServer) Stop() {

}
func (this *GameServer) Start() {

}
