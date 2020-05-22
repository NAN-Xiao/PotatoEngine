package server

import "potatoengine/src/router"

type IServer interface {
	RegisterLoginRouter(rt *router.IRouter)
	Initialize()
	Begin()
	Stop()
	Start()
}
