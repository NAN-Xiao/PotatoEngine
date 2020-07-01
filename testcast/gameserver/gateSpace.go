package main

import (
	"fmt"
	"net"
	"potatoengine/src/agent"
	"potatoengine/src/space"
)

type GateSpace struct {
	space.BaseSpace
}

func (this *GateSpace) Process() {
	addr := "0.0.0.0"
	post := 9000
	tcpAddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", addr, post))
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return
	}
	for {
		tcpcon, err := listener.AcceptTCP()
		if err != nil {
			break
		}
		go GateHandl(tcpcon)

	}

}
func (this *GateSpace) GetName() string {
	return this.Spacename
}
func (this *GateSpace) GetID() int32 {
	return this.SpaceID
}

func (this *GateSpace) LeaveSpace(ag *agent.Agent) {

}
func (this *GateSpace) EnterSpace(ag *agent.Agent) {

}
