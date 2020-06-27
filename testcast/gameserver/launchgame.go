package main

import "potatoengine/src/server"

func main() {

	server.LaunchServer()
	sp := NewLoginSpace("Login")
	RegistSpace("Login", sp)
	Serv()
	select {}
}