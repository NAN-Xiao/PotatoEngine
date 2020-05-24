package main

import (
	"potatoengine/src/server/gate"
	"potatoengine/src/space"
)

func main() {
	sr:=gate.NewGateServer()
	sp:=space.NewLoginSpace("LogIn")
	sr.RegisterSpace(sp)
	sr.Start()
	select {}
}
