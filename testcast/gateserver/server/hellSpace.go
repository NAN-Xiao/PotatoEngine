package main

import (
	"potatoengine/src/space"
)
//角色大厅
type HellSpace struct {
	space.BaseSpace
	space.ISpace
}

func (this *HellSpace) GetSpace() *space.BaseSpace {
	return &this.BaseSpace
}

//Ispace
func (this *HellSpace) OnStart() {
	println("gate space started")
	println(this.BaseSpace.Spacename)
}
func (this *HellSpace) Process() {

}
func (this *HellSpace) Tick() {

	//println("gate space tick;entitys len ",len(this.Entitys))
}
