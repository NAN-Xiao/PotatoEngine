package main

import "potatoengine/src/server"

func main() {
	server.Serv()
	select {}
}
