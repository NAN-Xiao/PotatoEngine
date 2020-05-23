package main

import (
	"potatoengine/src/server/gate"
)

func main() {
	sr:=gate.NewGateServer()
	sr.Start()
	select {}
}
