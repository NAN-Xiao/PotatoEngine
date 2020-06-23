package main

import (
	"fmt"
	"potatoengine/src/server"
	"potatoengine/src/server/login"
	"potatoengine/src/space"
)

type ServerName string

const (
	Loging ServerName = "LoginServer"
	Gate   ServerName = "GateServer"
	Game   ServerName = "GameServer"
)

var GateServer server.IServer
var LoginServer server.IServer
var ServerMap map[string]server.IServer

//func init() {
//	LaunchServer()
//	Initialize()
//}

func LaunchServer() {
	//GateServer = gate.NewGateServer()
	LoginServer = login.NewLoginServer()
	ServerMap = make(map[string]server.IServer)
	ServerMap["Login"] = LoginServer
	fmt.Println("launchServer")
	//ServerMap["Gate"] = GateServer
}

//给server注册space
func RegistSpace(name ServerName, sp space.ISpace) {

	if ServerMap == nil {
		fmt.Printf("sermap :%s is null", name)
		return
	}
	for m := range ServerMap {
		if m == string(name) {
			ServerMap[m].RegisterSpace(sp)
		}
	}
}
func Initialize() {

}
func Serv() {
	if ServerMap == nil {
		fmt.Println("server  map is nill")
		return
	}
	for m := range ServerMap {
		ServerMap[m].Run()
	}
}

var serverMsg map[string] interface{}

func RegistMessage(msg interface{}) {
	//todo



}
