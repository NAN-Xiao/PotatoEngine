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
func (this *GateSpace) OnStart() {
	println("gate space started")
	println( this.BaseSpace.Spacename)
}
func (this *GateSpace) Process() {
	//for {
	//	println("gate space processing")
	//}
}
func (this *GateSpace) Tick() {

	//println("gate space tick;entitys len ",len(this.Entitys))
}
