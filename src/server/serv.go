package server

import (
	"potatoengine/src/server/gate"
	"potatoengine/src/server/login"
	"potatoengine/src/space"
)

type ServerName string

const (
	Loging ServerName = "LoginServer"
	Gate   ServerName = "GateServer"
	Game   ServerName = "GameServer"
)

var GateServer IServer
var LoginServer IServer
var ServerMap map[string]IServer

func init() {
	LaunchServer()
	Initialize()
}

func LaunchServer() {
	GateServer = gate.NewGateServer()
	LoginServer = login.NewLoginServer()
	ServerMap := make(map[string]IServer)
	ServerMap["Loging"] = LoginServer
	ServerMap["Gate"] = GateServer
}

//给server注册space
func RegistSpace(name ServerName, sp space.ISpace) {

	if ServerMap == nil {
		return
	}
	for m := range ServerMap {
		if m == string(name) {
			ServerMap[m].RegisterSpace(sp)
		}
	}
}
func Initialize() {
	if ServerMap == nil {

		return
	}
	for m := range ServerMap {
		ServerMap[m].Initialize()
	}
}
func Serv() {
	if ServerMap == nil {
		return
	}
	for m := range ServerMap {
		ServerMap[m].Begin()
	}
}

func RegistMessage() {
	//todo
	//注册各种消息类型
	//message.RegisteredNetMessage("")
}
