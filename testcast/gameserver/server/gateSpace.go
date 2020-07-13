package main

import (
	"potatoengine/src/agent"
	"potatoengine/src/space"
)

type GateSpace struct {
	space.BaseSpace
}

func (this *GateSpace) Process() {

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
func (this *GateSpace)Tick()  {
	println("gate space tick")
}
