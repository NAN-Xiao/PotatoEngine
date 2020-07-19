package main

import (
	"potatoengine/src/space"
)

type GateSpace struct {
	space.BaseSpace
	space.ISpace
}

func (this *GateSpace) GetSpace() *space.BaseSpace {
	return &this.BaseSpace
}
//Ispace
func (this *GateSpace)OnStart()  {
	
}
func (this *GateSpace) Process() {

}
func (this *GateSpace) Tick() {
	println("gate space tick")
}


