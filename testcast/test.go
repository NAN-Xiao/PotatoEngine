package main

import "potatoengine/testcast/spacetest"

func main() {
	LaunchServer()
	sp := spacetest.NewLoginSpace("Login")
	RegistSpace("Login", sp)
	Serv()
	select {}
}
