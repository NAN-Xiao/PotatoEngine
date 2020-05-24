package main

import (
	"potatoengine/src/server/gate"
	"potatoengine/src/server/login"
)

func main() {
	sr:=gate.NewGateServer()
	sp:=login.NewLoginSpace("Login")
	sr.RegisterLoginRouter()
	sr.Start()
	select {}
}
