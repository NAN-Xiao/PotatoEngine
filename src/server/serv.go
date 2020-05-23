package server

import (
	"potatoengine/src/server/gate"
	"potatoengine/src/server/login"
)

var _gateServer IServer
var _loginServer IServer

func NewServer(serType string) IServer {
	var sr *IServer
	switch serType {
	case "Login":
		sr = login.NewLoginServer()
	case "Gate":
		sr = gate.NewGateServer()
	}
	if sr != nil {
		return *sr
	}
	return nil
}

//func Serv() {
//	_gateServer = NewServer("Gate")
//	_loginServer = NewServer("Login")
//	Initialize()
//}
//
//func Initialize() {
//	_gateServer.Initialize()
//	_loginServer.Initialize()
//}
