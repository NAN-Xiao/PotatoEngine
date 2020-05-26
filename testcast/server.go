package main

import "potatoengine/src/server"

func main() {
	//sr:=gate.NewGateServer()
	//sp:=space.NewLoginSpace("LogIn")
	//sr.RegisterSpace(sp)
	//sr.Start()

	server.Serv()
	select {}
}
